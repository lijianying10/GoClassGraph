package analysis

import "github.com/lijianying10/GoClassGraph/tool"

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
			op.Files= append(op.Files, filename)
			ana.Pkgs[pkgname] = op
		}
	}
	tool.Dump(ana.Pkgs)
}
