package analysis

import (
	"container/list"

	"github.com/lijianying10/GoClassGraph/tag"
)

// This file parse types

func ParseTypeField(tagLst *list.List, typename string) []string {
	res := []string{}
	for e := tagLst.Front(); e != nil; e = e.Next() {
		if e.Value.(tag.Tag).Type == "w" {
			if e.Value.(tag.Tag).Fields["ctype"] == typename {
				res = append(res, e.Value.(tag.Tag).Name)
			}
		}
	}
	return res
}

func ParseTypeMethod(tagLst *list.List, typename string) []string {
	res := []string{}
	for e := tagLst.Front(); e != nil; e = e.Next() {
		if e.Value.(tag.Tag).Type == "m" {
			if e.Value.(tag.Tag).Fields["ctype"] == typename {
				res = append(res, e.Value.(tag.Tag).Name)
			}
		}
	}
	return res
}
