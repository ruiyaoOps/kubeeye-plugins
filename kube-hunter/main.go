package main

import (
    log "github.com/golang/glog"

    "github.com/ruiyaoOps/kubeeye-plugins/kube-hunter/pkg"
)

func main()  {
    err, result := pkg.KubeHunterAudit()
    if err != nil {
        log.Error(err)
    }

    fmResult := pkg.FormatResult(result)
    //jsonOutput, err := json.MarshalIndent(fmResult, "", "    ")
    //if err != nil {
    //    log.Error(err)
    //}
    //fmt.Println(string(jsonOutput))
    err = pkg.ResultReporter(fmResult)
    if err != nil {
        log.Error(err)
    }
}