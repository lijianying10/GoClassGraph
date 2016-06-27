package analysis

import (
	"container/list"
	"fmt"

	"github.com/lijianying10/GoGraph/tag"
)

// This file parse interfaces

func ParseInterfaceMethod(tagLst *list.List, typename string) []string {
	res := []string{}
	for e := tagLst.Front(); e != nil; e = e.Next() {
		if e.Value.(tag.Tag).Type == "m" {
			if e.Value.(tag.Tag).Fields["ntype"] == typename {
				var access = AnalysisAccess(e.Value.(tag.Tag).Fields["access"])
				res = append(res, fmt.Sprintf("%s %s%s:%s", access, e.Value.(tag.Tag).Name, e.Value.(tag.Tag).Fields["signature"], e.Value.(tag.Tag).Fields["ntype"]))
			}
		}
	}
	return res
}
