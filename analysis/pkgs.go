package analysis

import (
	"container/list"

	"github.com/lijianying10/GoClassGraph/tag"
)

// This file parse package
func (ana *Analysis) NewPkgsAndFiles() {
	for filename, pkgname := range ana.File2Package {
		if _, ok := ana.Pkgs[pkgname]; !ok {
			ana.Pkgs[pkgname] = Pkg{
				Name:  pkgname,
				Files: []string{filename},
			}
		} else {
			op := ana.Pkgs[pkgname]
			op.Files = append(op.Files, filename)
			ana.Pkgs[pkgname] = op
		}
	}
}

func (ana *Analysis) ParsePkg2Tags() {
	for pkgName, pkg := range ana.Pkgs {
		pkgList := list.New()
		for _, file := range pkg.Files {
			for e := ana.tagList.Front(); e != nil; e = e.Next() {
				if e.Value.(tag.Tag).File == file {
					pkgList.PushBack(e.Value)
				}
			}
		}
		ana.pkgTagList[pkgName] = pkgList
	}
}

func (ana *Analysis) ParsePkgImports() {
	for pkgName, pkg := range ana.Pkgs {
		imports := make(map[string]bool)
		for e := ana.pkgTagList[pkgName].Front(); e != nil; e = e.Next() {
			if e.Value.(tag.Tag).Type == "i" {
				imports[e.Value.(tag.Tag).Name] = true
			}
		}

		for imp, _ := range imports {
			pkg.Import = append(pkg.Import, imp)
		}

		ana.Pkgs[pkgName] = pkg

	}
}

func (ana *Analysis) ParseConstant() {
	for pkgName, pkg := range ana.Pkgs {
		for e := ana.pkgTagList[pkgName].Front(); e != nil; e = e.Next() {
			if e.Value.(tag.Tag).Type == "c" {
				pkg.Consts = append(pkg.Consts, e.Value.(tag.Tag).Name)
			}
		}

		ana.Pkgs[pkgName] = pkg

	}
}

func (ana *Analysis) ParseVariables() {
	for pkgName, pkg := range ana.Pkgs {
		for e := ana.pkgTagList[pkgName].Front(); e != nil; e = e.Next() {
			if e.Value.(tag.Tag).Type == "v" {
				pkg.Variables = append(pkg.Variables, e.Value.(tag.Tag).Name)
			}
		}

		ana.Pkgs[pkgName] = pkg

	}
}
func (ana *Analysis) ParseType() {
	for pkgName, pkg := range ana.Pkgs {
		for e := ana.pkgTagList[pkgName].Front(); e != nil; e = e.Next() {
			if e.Value.(tag.Tag).Type == "t" {
				typename:=e.Value.(tag.Tag).Name
				NewType:= AType{
					Name:typename,
					Fields:ParseTypeField(ana.pkgTagList[pkgName],typename),
					Methods:ParseTypeMethod(ana.pkgTagList[pkgName],typename),
				}
				pkg.Types = append(pkg.Types, NewType)
			}
		}
		ana.Pkgs[pkgName] = pkg
	}
}
func (ana *Analysis) ParseInterface() {
	for pkgName, pkg := range ana.Pkgs {
		for e := ana.pkgTagList[pkgName].Front(); e != nil; e = e.Next() {
			if e.Value.(tag.Tag).Type == "n" {
				typename:=e.Value.(tag.Tag).Name
				pkg.Interfaces= append(pkg.Interfaces, AInterface{
					Name:typename,
					Methods:ParseInterfaceMethod(ana.pkgTagList[pkgName],typename),
				})
			}
		}
		ana.Pkgs[pkgName] = pkg
	}
}
