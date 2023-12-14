package pikachu

import (
	"encoding/json"
	"time"
)

type dataReport struct {
	Token     string           `json:"token"`
	Timestamp string           `json:"timestamp"`
	DataType  int              `json:"datatype"`
	Body      []DataReportBody `json:"body"`
}

type DataReportBody struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Val       string `json:"val"`
	Unit      string `json:"unit"`
	Quality   string `json:"quality"`
	Timestamp string `json:"timestamp"`
}

func DataChangedReport(deviceType string, deviceIdent string, dataType int, body []DataReportBody) error {
	token := GenerateStr()
	report := dataReport{
		Token:     token,
		Timestamp: time.Now().Format(time.RFC3339),
		DataType:  dataType,
		Body:      body,
	}

	b, err := json.Marshal(report)
	if err != nil {
		return err
	}

	return Publish(App+"/notify/spont/*/"+deviceType+"/"+deviceIdent, b)
}
