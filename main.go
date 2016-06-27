package main

import (
	"os"

	"github.com/lijianying10/GoClassGraph/analysis"
	"github.com/lijianying10/GoClassGraph/dot"
	"github.com/lijianying10/GoClassGraph/file"
	"github.com/lijianying10/GoClassGraph/tag"
	"github.com/lijianying10/log"
)

func main() {
	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatal("get cwd error")
	}
	//debug
	//currentPath = "/mnt/idea/gopath/src/github.com/eleme/esm-agent"
	files, err := file.WalkDir(currentPath)
	if err != nil {
		log.Fatal("get sourcecode error")
	}

	if len(files) == 0 {
		log.Fatal("There is no files under current dir")
	}

	tags := []tag.Tag{}
	for _, file := range files {
		ts, err := tag.Parse(file, true, currentPath)
		if err != nil {
			log.Errorf("parse error: %s\n\n", err)
			continue
		}
		tags = append(tags, ts...)
	}

	analysing := analysis.NewAnalysis(&tags)
	analysing.Analysis()

	dot := dot.NewDotOutput(&analysing)
	dot.OutputClassDiagram()

}
