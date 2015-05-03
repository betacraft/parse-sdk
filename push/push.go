package push

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type ParseSdk struct {
	applicationId string
	restApiKey    string
}

func Init(applicationId string, restApiKey string) *ParseSdk {
	return &ParseSdk{applicationId: applicationId, restApiKey: restApiKey}
}

func (p *ParseSdk) send(url *string, content *[]byte) error {
	client := &http.Client{}
	log.Println("Sending to", *url)
	req, err := http.NewRequest("POST", *url, bytes.NewBuffer(*content))
	req.Header.Add("X-Parse-Application-Id", p.applicationId)
	req.Header.Add("X-Parse-REST-API-Key", p.restApiKey)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var data map[string]interface{}
	err = json.Unmarshal(contents, &data)
	if err != nil {
		return err
	}
	log.Println(data)
	return nil
}

func (p *ParseSdk) SendTo(message map[string]interface{}) error {
	content, err := json.Marshal(message)
	if err != nil {
		return err
	}
	url := "https://api.parse.com/1/push"
	return p.send(&url, &content)
}

func (p *ParseSdk) SendDataToObjectId(objectId string, data map[string]interface{}) error {
	parseData := map[string]interface{}{"data": data, "where": map[string]interface{}{"objectId": objectId}}
	content, err := json.Marshal(parseData)
	if err != nil {
		return err
	}
	url := "https://api.parse.com/1/push"
	return p.send(&url, &content)
}

func (p *ParseSdk) SendAlertToObjectId(objectId string, message string) error {
	data := map[string]interface{}{"data": map[string]interface{}{"alert": message}, "where": map[string]interface{}{"objectId": objectId}}
	content, err := json.Marshal(data)
	if err != nil {
		return err
	}
	url := "https://api.parse.com/1/push"
	return p.send(&url, &content)
}
