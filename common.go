package pikachu

import (
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"github.com/tidwall/gjson"
	"math/rand"
	"sync"
	"time"
)

var (
	tokenCache sync.Map
	App        string
)

func subscribersAll() {
	var subscribers = []string{
		"database/get/response/" + App + "/modelschema",
		"database/set/response/" + App + "/model",
		"database/set/response/" + App + "/register",
		"database/get/response/" + App + "/guid",
	}

	for _, topic := range subscribers {
		Conn.Subscribe(topic, 0, func(client paho.Client, message paho.Message) {
			token := gjson.GetBytes(message.Payload(), "token").String()
			matchMsg(token, message.Payload())
		})
	}
}

func InitSync(broker string, appName string) {
	App = appName
	initMqtt(broker)
	subscribersAll()

}

func matchMsg(token string, msg []byte) {
	syn, exist := tokenCache.Load(token)
	if exist {
		syn.(Sync).Response <- msg
	}
}

func GenerateInt() (int, error) {
	// 使用当前时间种子化随机数生成器
	rand.Seed(time.Now().UnixNano())

	// 生成一个1000000到9999999之间的随机数字
	randomNumber := rand.Intn(9000000) + 1000000
	return randomNumber, nil
}

func GenerateStr() string {
	return uuid.NewString()[0:7]
}
