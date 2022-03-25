package pkg

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func ResultReporter(results PluginsResults) error {
	jsonResults, err := json.MarshalIndent(results,"","    ")
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(jsonResults)
	req, err := http.NewRequest("POST","http://kubeeye-server-manager.kubeeye-system.svc:8888/plugins", buf)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
