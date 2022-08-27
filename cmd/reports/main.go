package main

import (
	"io/fs"
	"io/ioutil"
	"path"
	"strings"

	"github.com/DaanV2/High-Performance-Cache/util/reports"
)

func main() {
	var files []fs.FileInfo
	var err error

	reportFolder := path.Join("..", "..", "reports")
	if files, err = ioutil.ReadDir(reportFolder); err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if path.Ext(file.Name()) == ".txt" {
			filepath := path.Join(reportFolder, file.Name())
			if err = convertTextToMD(filepath); err != nil {
				panic(err)
			}
		}
	}
}

func convertTextToMD(filepath string) error {
    report := reports.NewMarkDownReport()
    
    if err := report.ParseTextFile(filepath); err != nil {
        return err
    }

    filepath = strings.ReplaceAll(filepath, ".txt", ".md")
    return report.WriteTo(filepath)
}


