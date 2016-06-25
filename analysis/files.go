package analysis

// This file parse files

func (ana *Analysis) ParseFiles2Package() {
	for _, t := range *ana.tags {
		if t.Type == "p" {
			ana.File2Package[t.File] = t.Name
		}
	}
}
