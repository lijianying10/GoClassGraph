package main

import (
	"fmt"
	"os"

	"github.com/lijianying10/GoClassGraph/file"
	"github.com/lijianying10/GoClassGraph/tag"
	"github.com/lijianying10/log"
	"encoding/json"
)

func main() {
	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatal("get cwd error")
	}
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

	b,_:=json.Marshal(tags)
	fmt.Println(string(b))

}
