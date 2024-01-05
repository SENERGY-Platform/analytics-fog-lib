package mqtt

import (
	"crypto/tls"

	"os"
	"strconv"
	"time"
	"errors"
	log_level "github.com/y-du/go-log-level"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type MQTTClient struct {
	Client      MQTT.Client
	Retained    *bool
	Broker      BrokerConfig
	TopicConfig TopicConfig
	Logger      *log_level.Logger
	Relay 		RelayController
	ReconnectHandler *MQTT.ReconnectHandler
}

func (client *MQTTClient) ConnectMQTTBroker(username, password *string) {
	//MQTT.DEBUG = log.New(os.Stdout, "", 0)

	hostname, _ := os.Hostname()

	server := "tcp://"+client.Broker.Host+":"+client.Broker.Port
	retained := false
	client.Retained = &retained
	clientId := hostname+strconv.Itoa(time.Now().Second())

	connOpts := MQTT.NewClientOptions().
		AddBroker(server).
		SetClientID(clientId).
		SetCleanSession(true).

	
	if client.ReconnectHandler != nil {
		connOpts.SetReconnectingHandler(client.ReconnectHandler)
	}

	if username != nil && *username != "" {
		connOpts.SetUsername(*username)
		if *password != "" {
			connOpts.SetPassword(*password)
		}
	}
	// TODO insecure skip ?
	tlsConfig := &tls.Config{InsecureSkipVerify: true, ClientAuth: tls.NoClientCert}
	connOpts.SetTLSConfig(tlsConfig)

	connOpts.OnConnect = func(c MQTT.Client) {
		if(len(client.TopicConfig) != 0) {
			client.Logger.Debugf("Subscribed to topics: %v", client.TopicConfig)
			if token := c.SubscribeMultiple(client.TopicConfig, client.Relay.OnMessageReceived); token.Wait() && token.Error() != nil {
				panic(token.Error())
			}
		}
	}
	
	client.Client = MQTT.NewClient(connOpts)

	loopCounter := 0
	for {
		client.Logger.Debugf("Try to connect to: %s [%d/240]", server, loopCounter)

		if loopCounter > 240 {
			panic("Could not connect with broker")
		}

		if token := client.Client.Connect(); token.Wait() && token.Error() != nil {
			client.Logger.Errorf("Could not connect to %s : %s\n", server, token.Error())
			time.Sleep(5 * time.Second)
		} else {
			client.Logger.Debugf("Connected to %s\n", server)
			break
		}
		loopCounter += 1
	}

}

func (client *MQTTClient) CloseConnection() {
	client.Client.Disconnect(250)
	time.Sleep(1 * time.Second)
}

func (client *MQTTClient) Publish(topic string, message string, qos int) error {
	if !client.Client.IsConnected() {
		client.Logger.Error("WARNING: mqtt client not connected")
		return errors.New("mqtt client not connected")
	}

	token := client.Client.Publish(topic, byte(qos), *client.Retained, message)
	if token.Wait() && token.Error() != nil {
		client.Logger.Errorf("Error on Publish: ", token.Error())
		return token.Error()
	}
	return nil
}

func (client *MQTTClient) Subscribe(topic string, qos int) error {
	token := client.Client.Subscribe(topic, 2, client.Relay.OnMessageReceived)
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (client *MQTTClient) Unsubscribe(topic string) error {
	token := client.Client.Unsubscribe(topic)
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (client *MQTTClient) SetRelayController(relay RelayController) {
	client.Relay = relay
}