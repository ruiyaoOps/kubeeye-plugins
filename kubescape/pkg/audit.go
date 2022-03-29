package pkg

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
	"strings"

	"reflect"

	"github.com/armosec/opa-utils/reporthandling"
	"github.com/pkg/errors"
)

func KubescapeAudit() (err error, auditResults []reporthandling.FrameworkReport) {
	cmd := exec.Command("kubescape", "scan", "-e", "kube-system", "-f", "json")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err, auditResults
	}
	if err := cmd.Start(); err != nil {
		return err, nil
	}

	reader := bufio.NewReader(stdout)

	for {
		line, err := reader.ReadString('\n')
		if strings.Contains(line, "controlReports") && strings.Contains(line, "services") {
			err := json.NewDecoder(strings.NewReader(line)).Decode(&auditResults)
			if err != nil {
				return errors.Wrap(err, "decode result failed"), auditResults
			}
		}
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			return errors.Wrap(err, "get results failed"), auditResults
		}

	}

	if err := cmd.Wait(); err != nil {
		return errors.Wrap(err, "the command kube-hunter exec failed"), auditResults
	}
	return nil, auditResults
}

func FormatResult(results []reporthandling.FrameworkReport) (pluginResults PluginsResults) {
	for _, result := range results {
		if result.FailedResources == 0 && result.WarningResources == 0 {
			continue
		}

		for _, reports := range result.ControlReports {
			for _, ruleReport := range reports.RuleReports {
				for _, ruleRespons := range ruleReport.RuleResponses {
					//failedResources := ruleRespons.GetFailedResources()
					//for _, failedResource := range failedResources {
					//	kind := failedResource["kind"]
					//	name := failedResource["name"]
					//	if kind == "Group" && name == "system:masters"{
					//		kind = "ClusterRoleBinding"
					//
					//	}
					//	namespace := failedResource["namespace"]
					//
					//	fmt.Printf("get result kind: %+v\n", kind)
					//	fmt.Printf("get result name: %+v\n", name)
					//	fmt.Printf("get result namespace: %+v\n", namespace)
					//
					//}
					k8SApiObjects := ruleRespons.AlertObject.K8SApiObjects
					//g := ruleRespons.AlertObject.ExternalObjects
					for _, k8SApiObject := range k8SApiObjects {
						if k8SApiObject["relatedObjects"] == nil {
							continue
						}
						abc := reflect.TypeOf(k8SApiObject["relatedObjects"]).Kind()
						fmt.Printf("abc is : %+v", abc)
						bcd := reflect.ValueOf(k8SApiObject["relatedObjects"])
						fmt.Printf("bcd is : %+v",  bcd)
						for i := 0; i < bcd.Len(); i++ {
							fmt.Printf("abcdefg is : %+v \n",bcd.Index(i))
						}
					}
				}
			}
			level := "warning"
			message := reports.Name
			reason := reports.Description
			fmt.Printf("get result level: %+v\n", level)
			fmt.Printf("get result message: %+v\n", message)
			fmt.Printf("get result reason: %+v\n", reason)
		}
	}
	return pluginResults
}
