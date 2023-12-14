package pikachu

import (
	"encoding/json"
	"time"
)

type deviceReg struct {
	Token     string        `json:"token"`
	Timestamp string        `json:"timestamp"`
	Body      DeviceReqBody `json:"body"`
}

type DeviceReqBody struct {
	Model     string `json:"model"`
	Port      string `json:"port"`
	Addr      string `json:"addr"`
	Desc      string `json:"desc"`
	ManuId    string `json:"manuId"`
	IsReport  string `json:"isreport"`
	NodeId    string `json:"nodeID"`
	ProductId string `json:"productID"`
}

func DeviceRegister(body DeviceReqBody) ([]byte, error) {
	token := GenerateStr()
	req := deviceReg{
		Token:     token,
		Timestamp: time.Now().Format(time.RFC3339),
		Body:      body,
	}

	Body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	res, err := SendAndWaiting(App+"/set/request/database/register", token, Body)
	if err != nil {
		return nil, err
	}
	return res, nil

}
