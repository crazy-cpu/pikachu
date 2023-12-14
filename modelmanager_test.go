package pikachu

import (
	"testing"
)

func TestModelQuery(t *testing.T) {
	InitSync("mqtt://broker.emqx.io:1883", "ap323p")
	b, err := SendAndWaiting("App/get/request/database/modelschema", "123456", "hello,world ")
	if err != nil {
		t.Error(err)
	}

	t.Log(string(b))
}
