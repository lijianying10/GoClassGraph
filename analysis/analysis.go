package analysis

import (
	"fmt"

	"github.com/lijianying10/GoClassGraph/tag"
)

func NewAnalysis(tag *[]tag.Tag) Analysis {
	res := Analysis{}
	res.AInterfaces = make(map[string]AInterface)
	res.ATypes = make(map[string]AType)
	res.Pkgs = make(map[string]Pkg)
	res.File2Package = make(map[string]string)
	res.tags = tag
	return res
}

func (ana *Analysis) Analysis() {
	fmt.Println("start analysis")
	ana.ParseFiles2Package()
	ana.NewPkgsAndFiles()
}
