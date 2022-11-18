package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/fs"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func main() {

	totalFiles := 0
	totalDuration := time.Duration(0)
	directories := make(map[string]int)
	var totalSize int64

	err := filepath.Walk("/home/mattman/Music", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		if strings.HasSuffix(info.Name(), ".flac") {
			totalSize += info.Size()
			totalFiles++
			output, err := exec.Command("metaflac", "--show-total-samples", "--show-sample-rate", path).CombinedOutput()
			if err != nil {
				return err
			}
			scanner := bufio.NewScanner(bytes.NewBuffer(output))
			scanner.Scan()
			samples, err := strconv.ParseFloat(scanner.Text(), 32)
			if err != nil {
				return err
			}
			scanner.Scan()
			rate, err := strconv.ParseFloat(scanner.Text(), 32)
			if err != nil {
				return err
			}
			ms := int64((samples / rate) * 1000)
			totalDuration += time.Duration(ms) * time.Millisecond
			directories[filepath.Dir(path)]++
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("CDs", len(directories))
	fmt.Println("Tracks", totalFiles)
	fmt.Println("Total Length", fmt.Sprintf("%3.2f", totalDuration.Hours()), "hours")
	totalGb := fmt.Sprintf("%3.2f", float64(totalSize)/float64(1000000000))
	fmt.Println("Total size of all flac files", totalGb, "GB")

}
