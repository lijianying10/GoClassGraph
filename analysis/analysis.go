package analysis

import (
	"fmt"

	"container/list"

	"github.com/lijianying10/GoClassGraph/tag"
	"github.com/lijianying10/GoClassGraph/tool"
)

func NewAnalysis(tags *[]tag.Tag) Analysis {
	res := Analysis{}
	res.AInterfaces = make(map[string]AInterface)
	res.ATypes = make(map[string]AType)
	res.Pkgs = make(map[string]Pkg)
	res.File2Package = make(map[string]string)
	res.pkgTagList = make(map[string]*list.List)
	res.tags = tags
	res.tagList = list.New()
	for _, tag := range *tags {
		res.tagList.PushBack(tag)
	}
	return res
}

func (ana *Analysis) Analysis() {
	fmt.Println("start analysis")
	ana.ParseFiles2Package()
	ana.NewPkgsAndFiles()
	ana.ParsePkg2Tags()
	ana.ParsePkgImports()
	ana.ParseConstant()
	ana.ParseVariables()
	ana.ParseType()
	ana.ParseInterface()
	tool.Dump(ana.Pkgs)
}
