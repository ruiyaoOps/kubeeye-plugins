package pkg

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