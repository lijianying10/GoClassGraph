package tool

import (
	"encoding/json"

	"github.com/lijianying10/log"
)

func Dump(v interface{}) {
	r, err := json.MarshalIndent(v, " ", "  ")
	if err != nil {
		log.Error("err make json", err.Error())
	}
	log.Info(string(r))
}
