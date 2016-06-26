package analysis

import (
	"container/list"
	"github.com/lijianying10/GoClassGraph/tag"
)

// This file parse interfaces

func ParseInterfaceMethod(tagLst *list.List,typename string) []string {
	res := []string{}
	for e := tagLst.Front(); e != nil; e = e.Next() {
		if e.Value.(tag.Tag).Type == "m" {
			if e.Value.(tag.Tag).Fields["ntype"] == typename {
				res = append(res, e.Value.(tag.Tag).Name)
			}
		}
	}
	return res
}
