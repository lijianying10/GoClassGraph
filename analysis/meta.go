package analysis

import "github.com/lijianying10/GoClassGraph/tag"

type Pkg struct {
	Name       string
	Files      []string
	Import     []string
	Consts     []string
	Variables  []string
	Types      []AType
	Interfaces []AInterface
}

type AType struct {
	Name    string
	Fields  []string
	Methods []string
}

type AInterface struct {
	Name    string
	Methods []string
}

type Analysis struct {
	Pkgs         map[string]Pkg
	ATypes       map[string]AType
	AInterfaces  map[string]AInterface
	File2Package map[string]string
	tags         *[]tag.Tag
}
