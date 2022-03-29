package pkg

import (
	"github.com/armosec/armoapi-go/armotypes"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type PluginsResults struct {
	PluginName   string         `json:"pluginName"`
	AuditResults []AuditResults `json:"auditResults,omitempty"`
}

type AuditResults struct {
	NameSpace   string        `json:"namespace"`
	ResultInfos []ResultInfos `json:"resultInfos,omitempty"`
}

type ResultInfos struct {
	ResourceType  string `json:"resourceType"`
	ResourceInfos `json:"resourceInfos"`
}

type ResourceInfos struct {
	Name        string        `json:"name,omitempty"`
	ResultItems []ResultItems `json:"items"`
}

type ResultItems struct {
	Level   string `json:"level,omitempty"`
	Message string `json:"message,omitempty"`
	Reason  string `json:"reason,omitempty"`
}

type FrameworkReport struct {
	Name                  string          `json:"name"`
	ControlReports        []ControlReport `json:"controlReports"`
	Score                 float32         `json:"score,omitempty"`
	ARMOImprovement       float32         `json:"ARMOImprovement,omitempty"`
	WCSScore              float32         `json:"wcsScore,omitempty"`
	ResourceUniqueCounter `json:",inline"`
}

type ResourceUniqueCounter struct {
	TotalResources   int `json:"totalResources"`
	FailedResources  int `json:"failedResources"`
	WarningResources int `json:"warningResources"`
}

type ControlReport struct {
	armotypes.PortalBase  `json:",inline"`
	Control_ID            string       `json:"id,omitempty"`
	ControlID             string       `json:"controlID"`
	Name                  string       `json:"name"`
	RuleReports           []RuleReport `json:"ruleReports"`
	Remediation           string       `json:"remediation"`
	Description           string       `json:"description"`
	Score                 float32      `json:"score"`
	BaseScore             float32      `json:"baseScore,omitempty"`
	ARMOImprovement       float32      `json:"ARMOImprovement,omitempty"`
	ResourceUniqueCounter `json:",inline"`
}

type RuleReport struct {
	Name                  string         `json:"name"`
	Remediation           string         `json:"remediation"`
	RuleStatus            RuleStatus     `json:"ruleStatus"`
	RuleResponses         []RuleResponse `json:"ruleResponses"`
	ListInputKinds        []string       `json:"listInputIDs"`
	ResourceUniqueCounter `json:",inline"`
}

type RuleStatus struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type RuleResponse struct {
	AlertMessage string                            `json:"alertMessage"`
	FailedPaths  []string                          `json:"failedPaths"`
	FixPaths     []armotypes.FixPath               `json:"fixPaths"`
	RuleStatus   string                            `json:"ruleStatus"`
	PackageName  string                            `json:"packagename"`
	AlertScore   AlertScore                        `json:"alertScore"`
	AlertObject  AlertObject                       `json:"alertObject"`
	Context      []string                          `json:"context,omitempty"`
	Rulename     string                            `json:"rulename,omitempty"`
	Exception    *armotypes.PostureExceptionPolicy `json:"exception,omitempty"`
}

type AlertScore float32


type AlertObject struct {
	K8SApiObjects   []K8SObject `json:"k8sApiObjects,omitempty"`
	ExternalObjects map[string]interface{}   `json:"externalObjects,omitempty"`
}

type K8SObject struct {
	apiGroup string `json:"apiGroup"`
	kind string `json:"kind"`
	name string `json:"name"`
	relatedObjects unstructured.Unstructured `json:"relatedObjects"`
}