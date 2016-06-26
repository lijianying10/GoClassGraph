package dot

import (
	"fmt"

	"strings"

	"github.com/lijianying10/GoClassGraph/analysis"
	"github.com/satori/go.uuid"
)

type DotOutput struct {
	ana *analysis.Analysis
}

func NewDotOutput(ana *analysis.Analysis) DotOutput {
	res := DotOutput{}
	res.ana = ana
	return res
}

func (dot *DotOutput) OutputClassDiagram() {
	fmt.Printf(`digraph G {
		fontname = "Bitstream Vera Sans"
		fontsize = 8
	    node [
                fontname = "Bitstream Vera Sans"
                fontsize = 8
                shape = "record"
        ]

        edge [
                fontname = "Bitstream Vera Sans"
                fontsize = 8
        ]
		%s
	}`, dot.DotParsePkg())
}

func (dot *DotOutput) DotParsePkg() string {
	res := ""
	for pkgName, pkg := range dot.ana.Pkgs {
		res = res + fmt.Sprintf(`

	subgraph cluster_%s {
                label = "%s"

                %s [
                        label = "{%s|%s|%s|%s}"
                ]

				%s

        }`,
			RandomName(),
			pkgName,
			RandomName(),
			StringArrayToDotLines(pkg.Files),
			StringArrayToDotLines(pkg.Imports),
			StringArrayToDotLines(pkg.Consts),
			StringArrayToDotLines(pkg.Variables),
			dot.DotParseType(&pkg)+dot.DotParseInterface(&pkg),
		)
	}
	return res
}

func (dot *DotOutput) DotParseType(pkg *analysis.Pkg) string {
	res := ""
	for _, t := range pkg.Types {
		res = res + fmt.Sprintf(`
		%s [
			label = "{%s|%s|%s}"
		]`, RandomName(), t.Name, StringArrayToDotLines(t.Fields), StringArrayToDotLines(t.Methods))
	}
	return res
}

func (dot *DotOutput) DotParseInterface(pkg *analysis.Pkg) string {
	res := ""
	for _, t := range pkg.Interfaces {
		res = res + fmt.Sprintf(`
		%s [
			label = "{%s|%s}"
		]`, RandomName(), t.Name, StringArrayToDotLines(t.Methods))
	}
	return res
}

func StringArrayToDotLines(inp []string) string {
	if len(inp) == 0 {
		return ""
	}
	res := ""
	for _, i := range inp {
		i = strings.Replace(i, "{", "\\{", -1)
		i = strings.Replace(i, "}", "\\}", -1)
		res = res + i + "\\l"
	}
	return res
}

func RandomName() string {
	var res string
	res = "T" + uuid.NewV4().String()
	res = strings.Replace(res, "-", "", -1)
	return res
}
