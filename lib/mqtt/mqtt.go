package mqtt

import (
	"crypto/tls"
	"flag"

	//"log"
	"os"
	"strconv"
	"time"

	log_level "github.com/y-du/go-log-level"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type MQTTClient struct {
	Client      MQTT.Client
	Retained    *bool
	Broker      BrokerConfig
	TopicConfig TopicConfig
	Logger      *log_level.Logger
}

func (client *MQTTClient) ConnectMQTTBroker(relay RelayController) {
	//MQTT.DEBUG = log.New(os.Stdout, "", 0)

	hostname, _ := os.Hostname()

	server := flag.String("server", "tcp://"+client.Broker.Host+":"+client.Broker.Port, "The full url of the MQTT server to connect to ex: tcp://127.0.0.1:1883")
	client.Retained = flag.Bool("retained", false, "Are the messages sent with the retained flag")
	clientId := flag.String("clientid", hostname+strconv.Itoa(time.Now().Second()), "A clientid for the connection")
	username := flag.String("username", "", "A username to authenticate to the MQTT server")
	password := flag.String("password", "", "Password to match username")
	flag.Parse()

	connOpts := MQTT.NewClientOptions().AddBroker(*server).SetClientID(*clientId).SetCleanSession(true)
	if *username != "" {
		connOpts.SetUsername(*username)
		if *password != "" {
			connOpts.SetPassword(*password)
		}
	}
	tlsConfig := &tls.Config{InsecureSkipVerify: true, ClientAuth: tls.NoClientCert}
	connOpts.SetTLSConfig(tlsConfig)

	connOpts.OnConnect = func(c MQTT.Client) {
		if token := c.SubscribeMultiple(client.TopicConfig, relay.OnMessageReceived); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}
	client.Client = MQTT.NewClient(connOpts)

	loopCounter := 0
	for {
		client.Logger.Debugf("Try to connect to: %s [%d/240]", *server, loopCounter)

		if loopCounter > 240 {
			panic("Could not connect with broker")
		}

		if token := client.Client.Connect(); token.Wait() && token.Error() != nil {
			client.Logger.Errorf("Could not connect to %s : %s\n", *server, token.Error())
			time.Sleep(5 * time.Second)
		} else {
			client.Logger.Debugf("Connected to %s\n", *server)
			break
		}
		loopCounter += 1
	}

}

func (client *MQTTClient) CloseConnection() {
	client.Client.Disconnect(250)
	time.Sleep(1 * time.Second)
}

func (client *MQTTClient) Publish(topic string, message string, qos int) {
	client.Client.Publish(topic, byte(qos), *client.Retained, message)
}
