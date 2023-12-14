package pikachu

import "testing"

func TestSetModel(t *testing.T) {
	InitSync("mqtt://broker.emqx.io:1883", "ap323p")
	body := []ModelBody{
		{
			Name:       "name",
			Type:       "Type",
			Unit:       "Unit",
			DeadZone:   "deadZone",
			Ratio:      "ratio",
			IsReport:   "1",
			UserDefine: "UserDefine",
		},
	}
	b, err := SetModel("WD", body)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(b)
}
