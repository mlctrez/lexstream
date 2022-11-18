package main

import (
	"bytes"
	"fmt"
	"github.com/mlctrez/lexstream/catalog"
	"github.com/mlctrez/lexstream/internal/jutil"
	"github.com/mlctrez/lexstream/internal/settings"
	"github.com/mlctrez/lexstream/smapi"
	v0catalog "github.com/mlctrez/lexstream/smapiv0/catalog"
	"github.com/mlctrez/lexstream/smapiv0/catalog/upload"
	"io"
	"log"
	"net/http"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"time"
)

// refresh catalogs currently stored

type CatalogRefresh struct {
	smapi        *smapi.Smapi
	settings     *settings.Settings
	vendorId     string
	catalogIdMap map[v0catalog.CatalogType]*v0catalog.CatalogSummary
}

func main() {
	r := &CatalogRefresh{}
	execute(
		r.setup,
		r.retrieveCatalogs,
		//r.createMissingCatalogs,
		//r.associateCatalogsWithSkill,
		//r.uploadCatalogs,
		r.dumpUploadStatus,
	)
}

func (c *CatalogRefresh) setup() (err error) {
	c.smapi = smapi.New()
	if c.settings, err = settings.Load(); err != nil {
		return
	}
	if c.vendorId, err = c.smapi.GetSingleVendor(); err != nil {
		return
	}
	return nil
}

func (c *CatalogRefresh) retrieveCatalogs() (err error) {
	// clear map
	c.catalogIdMap = map[v0catalog.CatalogType]*v0catalog.CatalogSummary{}
	// refresh it with catalogs matching this skill
	var catalogsResponse *v0catalog.ListCatalogsResponse
	catalogsResponse, err = c.smapi.V0.ListCatalogsForVendorV0("", 50, c.vendorId)
	for _, summary := range catalogsResponse.Catalogs {
		if strings.HasPrefix(summary.Title, c.settings.SkillName) {
			c.catalogIdMap[*summary.Type_] = summary
		}
	}
	return
}

func (c *CatalogRefresh) createMissingCatalogs() (err error) {

	createdCatalog := false
	for _, catalogType := range supportedCatalogTypes() {
		if c.catalogIdMap[catalogType] != nil {
			continue
		}
		createRequest := &v0catalog.CreateCatalogRequest{
			Title:    catalogTitle(catalogType, c.settings),
			Type_:    &catalogType,
			Usage:    usageForCatalogType(catalogType),
			VendorId: c.vendorId,
		}
		if _, err = c.smapi.V0.CreateCatalogV0(createRequest); err != nil {
			return err
		}
		createdCatalog = true
	}

	if createdCatalog {
		return c.retrieveCatalogs()
	}

	return nil
}

func (c *CatalogRefresh) associateCatalogsWithSkill() (err error) {

	var skillId string
	if skillId, err = c.smapi.GetSkillIdForName(c.settings.SkillName); err != nil {
		return
	}
	for _, summary := range c.catalogIdMap {
		foundAssociation := false
		for _, id := range summary.AssociatedSkillIds {
			if skillId == id {
				foundAssociation = true
			}
		}
		if !foundAssociation {
			if err = c.smapi.V0.AssociateCatalogWithSkillV0(skillId, summary.Id); err != nil {
				return
			}
			fmt.Println("associated", summary.Title, "with skill", skillId)
		}
	}
	return nil

}

func (c *CatalogRefresh) uploadCatalogs() (err error) {
	for _, summary := range c.catalogIdMap {

		fmt.Println(summary.Id, summary.Title, summary.LastUpdatedDate.Format(time.RFC822))

		var response *upload.ListUploadsResponse
		if response, err = c.smapi.V0.ListUploadsForCatalogV0("", 50, summary.Id); err != nil {
			return
		}

		sort.SliceStable(response.Uploads, func(i, j int) bool {
			return response.Uploads[i].LastUpdatedDate.Before(response.Uploads[j].LastUpdatedDate)
		})

		lenUploads := len(response.Uploads)
		goForUpload := lenUploads == 0
		var lastState upload.UploadStatus
		if lenUploads > 0 {
			lastState = *response.Uploads[lenUploads-1].Status
			goForUpload = lastState == upload.UploadStatus_SUCCEEDED() ||
				lastState == upload.UploadStatus_FAILED()
		}
		if !goForUpload {
			fmt.Printf("catalog %s latest upload is status %s, skipping\n", summary.Id, lastState)
			continue
		}
		fmt.Printf("go for upload catalog %s latest upload is status is %s\n", summary.Id, lastState)

		err = c.uploadFromStaging(summary)
	}
	return
}

func (c *CatalogRefresh) dumpUploadStatus() (err error) {
	for _, summary := range c.catalogIdMap {

		fmt.Println("**************************************************************")
		fmt.Println(summary.Id, summary.Title, summary.LastUpdatedDate.Format(time.RFC822))

		var response *upload.ListUploadsResponse
		if response, err = c.smapi.V0.ListUploadsForCatalogV0("", 50, summary.Id); err != nil {
			return
		}

		sort.SliceStable(response.Uploads, func(i, j int) bool {
			return response.Uploads[i].LastUpdatedDate.Before(response.Uploads[j].LastUpdatedDate)
		})

		if len(response.Uploads) == 0 {
			fmt.Println("  No uploads found")
			continue
		}

		for i, uploadSummary := range response.Uploads {
			// remove to dump all upload summaries, this just displays the last one
			if i != len(response.Uploads)-1 {
				continue
			}

			fmt.Println(i, uploadSummary.Id, uploadSummary.LastUpdatedDate.Format(time.RFC822), *uploadSummary.Status)
			var uploadResponse *upload.GetContentUploadResponse
			if uploadResponse, err = c.smapi.V0.GetContentUploadByIdV0(summary.Id, uploadSummary.Id); err != nil {
				return
			}

			sort.SliceStable(uploadResponse.IngestionSteps, func(i, j int) bool {
				return (*uploadResponse.IngestionSteps[i].Name) < (*uploadResponse.IngestionSteps[j].Name)
			})

			for _, step := range uploadResponse.IngestionSteps {
				fmt.Println("  ", *step.Name, *step.Status, step.LogUrl)
			}

		}

	}
	return
}

func (c *CatalogRefresh) uploadFromStaging(catSummary *v0catalog.CatalogSummary) (err error) {

	uploadRequest := &upload.CreateContentUploadRequest{NumberOfUploadParts: 1}

	var uploadResponse *upload.CreateContentUploadResponse
	if uploadResponse, err = c.smapi.V0.CreateContentUploadV0(catSummary.Id, uploadRequest); err != nil {
		return
	}
	uploadUrl := uploadResponse.PresignedUploadParts[0].Url

	var data any
	if data, err = catalog.ReadFromStaging(*catSummary.Type_); err != nil {
		return err
	}
	var buf *bytes.Buffer
	if buf, err = jutil.Marshal(data, false); err != nil {
		return err
	}

	var etag string
	if etag, err = putSignedUrl(uploadUrl, buf); err != nil {
		return
	}

	completeUpload := &upload.CompleteUploadRequest{
		PartETags: []*upload.PreSignedUrlItem{{PartNumber: 1, ETag: etag}},
	}
	return c.smapi.V0.CompleteCatalogUploadV0(catSummary.Id, uploadResponse.Id, completeUpload)

}

func putSignedUrl(url string, buf *bytes.Buffer) (etag string, err error) {
	var request *http.Request
	if request, err = http.NewRequest(http.MethodPut, url, buf); err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")

	var response *http.Response
	if response, err = http.DefaultClient.Do(request); err != nil {
		return
	}

	if response.StatusCode != 200 {
		err = fmt.Errorf("non 200 response code %d : %s", response.StatusCode, response.Status)
		return
	}
	_, _ = io.ReadAll(response.Body)
	_ = response.Body.Close()

	etag = response.Header.Get("Etag")
	return
}

func catalogTitle(id v0catalog.CatalogType, settings *settings.Settings) string {
	switch id {
	case v0catalog.CatalogType_AMAZONMusicGroup():
		return settings.SkillName + " Artists"
	case v0catalog.CatalogType_AMAZONMusicAlbum():
		return settings.SkillName + " Albums"
	case v0catalog.CatalogType_AMAZONMusicRecording():
		return settings.SkillName
	case v0catalog.CatalogType_AMAZONMusicPlaylist():
		return settings.SkillName + " Playlists"
	default:
		panic("unhandled catalog type " + id)
	}
}

func usageForCatalogType(id v0catalog.CatalogType) (usage *v0catalog.CatalogUsage) {
	var result v0catalog.CatalogUsage
	switch id {
	case v0catalog.CatalogType_AMAZONMusicGroup():
		result = v0catalog.CatalogUsage_AlexaMusicCatalogMusicGroup()
	case v0catalog.CatalogType_AMAZONMusicAlbum():
		result = v0catalog.CatalogUsage_AlexaMusicCatalogMusicAlbum()
	case v0catalog.CatalogType_AMAZONMusicRecording():
		result = v0catalog.CatalogUsage_AlexaMusicCatalogMusicRecording()
	case v0catalog.CatalogType_AMAZONMusicPlaylist():
		result = v0catalog.CatalogUsage_AlexaMusicCatalogMusicPlaylist()
	default:
		panic("unhandled catalog type " + id)
	}
	return &result
}

func showFailedUploads(sc *smapi.Smapi, found *v0catalog.CatalogSummary) {
	listUploadsResponse, err := sc.V0.ListUploadsForCatalogV0("", 50, found.Id)
	failErr(err)
	for _, upl := range listUploadsResponse.Uploads {

		if *upl.Status == upload.UploadStatus_SUCCEEDED() {
			continue
		}
		fmt.Println("CatalogID: ", found.Id)
		fmt.Printf("  Type, Title = %s %s\n", *found.Type_, found.Title)
		fmt.Printf("    Status, UploadID, Created = %s %s %s\n", *upl.Status, upl.Id, upl.CreatedDate)

		contentUploadResponse, err := sc.V0.GetContentUploadByIdV0(found.Id, upl.Id)
		failErr(err)

		fmt.Println("      File:", *contentUploadResponse.File.Status)
		//if contentUploadResponse.File.PresignedDownloadUrl != "" {
		//	fmt.Println("      Doownload URL:", contentUploadResponse.File.PresignedDownloadUrl)
		//}

		for _, step := range contentUploadResponse.IngestionSteps {
			fmt.Printf("        Step %s = %s\n", *step.Name, *step.Status)
		}
	}
}

func supportedCatalogTypes() []v0catalog.CatalogType {
	return []v0catalog.CatalogType{
		v0catalog.CatalogType_AMAZONMusicGroup(),
		v0catalog.CatalogType_AMAZONMusicAlbum(),
		v0catalog.CatalogType_AMAZONMusicRecording(),
		v0catalog.CatalogType_AMAZONMusicPlaylist(),
	}
}

func failErr(err error) {
	if err != nil {
		panic(err)
	}
}

func execute(functions ...func() error) {
	for _, step := range functions {
		name := runtime.FuncForPC(reflect.ValueOf(step).Pointer()).Name()
		fmt.Println(strings.TrimPrefix(name, "main."))
		err := step()
		if err != nil {
			log.Fatal(err)
		}
	}
}
