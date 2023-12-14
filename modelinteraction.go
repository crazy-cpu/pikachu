package pikachu

import (
	"encoding/json"
	"time"
)

type model struct {
	Token     string      `json:"token"`
	TimeStamp string      `json:"timeStamp"`
	Model     string      `json:"model"`
	Body      []ModelBody `json:"body"`
}

type ModelBody struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Unit       string `json:"unit"`
	DeadZone   string `json:"deadzone"`
	Ratio      string `json:"ratio"`
	IsReport   string `json:"isReport"`
	UserDefine string `json:"userdefine"`
}

func SetModel(name string, body []ModelBody) ([]byte, error) {
	token := GenerateStr()
	m := model{
		Token:     token,
		TimeStamp: time.Now().Format(time.RFC3339),
		Model:     name,
		Body:      body,
	}

	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	return SendAndWaiting(App+"/set/request/database/model", token, b)
}

type guidQuery struct {
	Token     string     `json:"token"`
	Timestamp string     `json:"timestamp"`
	Body      []GuidBody `json:"body"`
}

type GuidBody struct {
	Model string `json:"model"`
	Port  string `json:"port"`
	Addr  string `json:"addr"`
	Desc  string `json:"desc"`
}

func QueryGuid(body []GuidBody) ([]byte, error) {
	token := GenerateStr()
	m := guidQuery{
		Token:     token,
		Timestamp: time.Now().Format(time.RFC3339),
		Body:      body,
	}
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	res, err := SendAndWaiting(App+"/get/request/database/guid", token, b)
	if err != nil {
		return nil, err
	}
	return res, nil
}
