package pkg

import (
    "bufio"
    "encoding/json"
    "io"
    "os/exec"
    "strings"

    "github.com/pkg/errors"
)

func KubeHunterAudit() (err error, result KubeHunterResults) {
    cmd := exec.Command("kube-hunter", "--pod", "--report", "json")
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        return errors.Wrap(err, "the command kube-hunter create pipe failed"), result
    }
    if err := cmd.Start(); err != nil {
        return errors.Wrap(err, "the command kube-hunter start failed"), result
    }

    reader := bufio.NewReader(stdout)

    for {
        line, err := reader.ReadString('\n')
        if err != nil && err == io.EOF {
            break
        } else if err != nil {
            return errors.Wrap(err, "get results failed"), result
        }
        if strings.Contains(line, "nodes") && strings.Contains(line, "services") {
            err := json.NewDecoder(strings.NewReader(line)).Decode(&result)
            if err != nil {
                return errors.Wrap(err, "decode result failed"), result
            }
        }
    }

    if err := cmd.Wait(); err != nil {
        return errors.Wrap(err, "the command kube-hunter exec failed"), result
    }
    return nil, result
}

func FormatResult(result KubeHunterResults) (pluginResults PluginsResults) {
    var auditResults []AuditResults
    var auditResult AuditResults

    for _, vulnerability := range result.Vulnerabilities {
        var resultInfos ResultInfos
        resultInfos.ResourceType = "Node"
        var resourceInfos ResourceInfos
        var resultItems ResultItems
        if vulnerability.Severity == "high" {
            resultItems.Level = "danger"
            resultItems.Message = vulnerability.Description
            resourceInfos.Name = vulnerability.Vulnerability
        } else if vulnerability.Severity == "medium" {
            resultItems.Level = "warning"
            resultItems.Message = vulnerability.Description
            resourceInfos.Name = vulnerability.Vulnerability
        } else {
            continue
        }
        resourceInfos.ResultItems = append(resourceInfos.ResultItems, resultItems)
        resultInfos.ResourceInfos = resourceInfos
        auditResult.ResultInfos = append(auditResult.ResultInfos, resultInfos)
    }
    auditResults = append(auditResults, auditResult)
    pluginResults.PluginName = "kubehunter"
    pluginResults.AuditResults = auditResults
    return pluginResults
}