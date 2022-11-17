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

	err := filepath.Walk("/home/mattman/Music", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		if strings.HasSuffix(info.Name(), ".flac") {
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

		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(totalFiles)
	fmt.Println(totalDuration)

}
