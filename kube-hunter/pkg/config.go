package pkg

type KubeHunterResults struct {
    Nodes   []Node   `json:"nodes"`
    Services []Service  `json:"service"`
    Vulnerabilities []Vulnerabilitie  `json:"vulnerabilities"`
}

type Node struct {
    Type string `json:"type"`
    Location string `json:"location"`
}

type Service struct {
    Service string `json:"service"`
    Location string `json:"location"`
}

type Vulnerabilitie struct {
    Location string `json:"location"`
    Vid string      `json:"vid"`
    Category    string  `json:"category"`
    Severity    string  `json:"severity"`
    Vulnerability string    `json:"vulnerability"`
    Description string  `json:"description"`
    Evidence    string  `json:"evidence"`
    Avd_reference string    `json:"avd_reference"`
}

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