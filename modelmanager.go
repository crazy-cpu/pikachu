package pikachu

import "pikachu/mqtt"

func ModelQuery(token string, payload any) ([]byte, error) {
	err := mqtt.Publish(App+"/get/request/database/modelschema", payload)
	if err != nil {
		return nil, err
	}

	return newSyncAndWaiting(token)

}
