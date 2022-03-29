package pkg

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

func KubescapeAudit() (err error, auditResults []FrameworkReport) {
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

func FormatResult(results []FrameworkReport) (pluginResults PluginsResults) {
	for _, result := range results {
		if result.FailedResources == 0 && result.WarningResources == 0 {
			continue
		}
		for _, reports := range result.ControlReports {
			for _, ruleReport := range reports.RuleReports {
				for _, ruleRespons := range ruleReport.RuleResponses {
					k8SApiObjects := ruleRespons.AlertObject.K8SApiObjects
					for _, object := range k8SApiObjects {
						fmt.Println(object)
						a := object.relatedObjects.GetName()
						b := object.relatedObjects.GetNamespace()
						c := object.relatedObjects.GetKind()
						fmt.Printf("name: %+v\n", a)
						fmt.Printf("namespace: %+v\n", b)
						fmt.Printf("kind: %+v\n", c)
					}
				}
			}
			//level := "warning"
			//message := reports.Name
			//reason := reports.Description
			//fmt.Printf("get result level: %+v\n", level)
			//fmt.Printf("get result message: %+v\n", message)
			//fmt.Printf("get result reason: %+v\n", reason)
		}
	}
	return pluginResults
}
