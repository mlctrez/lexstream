package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/mlctrez/lexstream/catalog"
	"github.com/mlctrez/lexstream/internal/jutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}
	dd := dynamodb.NewFromConfig(cfg)

	tableName := "lexstream"

	var tableNames []string

	lti := &dynamodb.ListTablesInput{}
	first := true
	for first || lti.ExclusiveStartTableName != nil {
		first = false
		var lto *dynamodb.ListTablesOutput
		if lto, err = dd.ListTables(ctx, lti); err != nil {
			log.Fatal(err)
		}
		tableNames = append(tableNames, lto.TableNames...)
		lti.ExclusiveStartTableName = lto.LastEvaluatedTableName
	}

	fmt.Println(tableNames)
	tableExists := false
	for _, name := range tableNames {
		if name == tableName {
			tableExists = true
		}
	}

	fmt.Println("table exists =", tableExists)
	os.Exit(0)

	if !tableExists {
		_, err = dd.CreateTable(ctx, &dynamodb.CreateTableInput{
			AttributeDefinitions: []types.AttributeDefinition{
				{AttributeName: aws.String("id"), AttributeType: types.ScalarAttributeTypeS},
			},
			KeySchema: []types.KeySchemaElement{
				{AttributeName: aws.String("id"), KeyType: types.KeyTypeHash},
			},
			TableName: aws.String(tableName),
			ProvisionedThroughput: &types.ProvisionedThroughput{
				ReadCapacityUnits:  aws.Int64(5),
				WriteCapacityUnits: aws.Int64(5),
			},
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("created table, inserts may fail")
	}

	builder := catalog.New()

	for _, header := range []catalog.HeaderType{catalog.MusicGroup, catalog.MusicAlbum, catalog.MusicRecording} {
		n := strings.ToLower(strings.TrimPrefix(string(header), "AMAZON."))
		p := filepath.Join("catalog_staging", n+".json")
		switch header {
		case catalog.MusicGroup:
			err = jutil.DecodePath(p, builder.ArtistCatalog)
		case catalog.MusicAlbum:
			err = jutil.DecodePath(p, builder.AlbumCatalog)
		case catalog.MusicRecording:
			err = jutil.DecodePath(p, builder.TrackCatalog)
		}
		if err != nil {
			log.Fatal(err)
		}
	}

	batchPut := func(requests []types.WriteRequest) (*dynamodb.BatchWriteItemOutput, error) {
		return dd.BatchWriteItem(ctx, &dynamodb.BatchWriteItemInput{
			RequestItems:           map[string][]types.WriteRequest{tableName: requests},
			ReturnConsumedCapacity: types.ReturnConsumedCapacityTotal,
		})
	}

	l := len(builder.ArtistCatalog.Entities)
	var requests []types.WriteRequest
	for it, entity := range builder.ArtistCatalog.Entities {

		fmt.Println(fmt.Sprintf("%4d/%4d", it+1, l), entity.Id, entity.Names[0].Value)

		var value types.AttributeValue
		value, err = attributevalue.Marshal(entity)
		if err != nil {
			log.Fatal(err)
		}
		requests = append(requests, types.WriteRequest{PutRequest: &types.PutRequest{
			Item: map[string]types.AttributeValue{
				"id":    &types.AttributeValueMemberS{Value: entity.Id},
				"value": value,
			},
		}})

		if len(requests) > 20 {
			put, err := batchPut(requests)
			requests = []types.WriteRequest{}
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("unprocessed", len(put.UnprocessedItems[tableName]))
			fmt.Println("consumed", *put.ConsumedCapacity[0].CapacityUnits)
			time.Sleep(1000 * time.Millisecond)
		}

	}
	if len(requests) > 0 {
		put, err := batchPut(requests)
		requests = []types.WriteRequest{}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("unprocessed", len(put.UnprocessedItems[tableName]))
		fmt.Println("consumed", *put.ConsumedCapacity[0].CapacityUnits)
		time.Sleep(1000 * time.Millisecond)
	}

	os.Exit(0)
	l = len(builder.AlbumCatalog.Entities)
	for it, entity := range builder.AlbumCatalog.Entities {
		fmt.Println(fmt.Sprintf("%4d/%4d", it+1, l), entity.Id, entity.Names[0].Value)
		var value types.AttributeValue
		value, err = attributevalue.Marshal(entity)
		if err != nil {
			log.Fatal(err)
		}
		item, err := dd.PutItem(ctx, &dynamodb.PutItemInput{
			TableName: aws.String(tableName),
			Item: map[string]types.AttributeValue{
				"id":    &types.AttributeValueMemberS{Value: entity.Id},
				"value": value,
			},
		})
		if err != nil {
			log.Fatal(err)
		}
		_ = item
		time.Sleep(500 * time.Millisecond)
	}
	l = len(builder.TrackCatalog.Entities)
	for it, entity := range builder.TrackCatalog.Entities {
		fmt.Println(fmt.Sprintf("%4d/%4d", it+1, l), entity.Id, entity.Names[0].Value)
		var value types.AttributeValue
		value, err = attributevalue.Marshal(entity)
		if err != nil {
			log.Fatal(err)
		}
		item, err := dd.PutItem(ctx, &dynamodb.PutItemInput{
			TableName: aws.String(tableName),
			Item: map[string]types.AttributeValue{
				"id":    &types.AttributeValueMemberS{Value: entity.Id},
				"value": value,
			},
		})
		if err != nil {
			log.Fatal(err)
		}
		_ = item
		time.Sleep(500 * time.Millisecond)
	}

}
