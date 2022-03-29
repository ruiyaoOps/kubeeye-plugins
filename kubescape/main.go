package main

import (
	"encoding/json"
	"fmt"

	log "github.com/golang/glog"
	"github.com/ruiyaoOps/kubeeye-plugins/kubscape/pkg"
)

func main()  {
	err, auditResults := pkg.KubescapeAudit()
	if err != nil {
		log.Error(err)
	}

	fmResults := pkg.FormatResult(auditResults)
	jsonResults, err := json.MarshalIndent(fmResults,"","    ")
	if err != nil {
		log.Error(err)
	}
	fmt.Println(string(jsonResults))
}