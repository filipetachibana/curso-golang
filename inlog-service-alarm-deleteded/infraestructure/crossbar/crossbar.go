package crossbar

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"

	"github.com/mitchellh/mapstructure"
	turnpike "gopkg.in/jcelliott/turnpike.v2"
)

// Crossbar struct client connection
type Crossbar struct {
	Conn turnpike.Client
}

type configuration struct {
	ServerAddress string
	ServerPort    int
	SubServerName string
}

var _config configuration

// SetConfiguration conection crossbar
func SetConfiguration(serverAddress string, serverPort int, subServerName string) {

	_config.ServerAddress = serverAddress
	_config.ServerPort = serverPort
	_config.SubServerName = subServerName

	crAdreess := "ws://" + _config.ServerAddress + ":" + strconv.Itoa(_config.ServerPort) + "/ws"
	c, err := turnpike.NewWebsocketClient(turnpike.JSON, crAdreess, nil, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = c.JoinRealm(_config.SubServerName, nil)
	conn = Crossbar{*c}
	if err != nil {
		log.Fatal(err)
	}
}

// EventHandler função de recebimento de informações
type EventHandler func(args Message)

// BasicMethodHandler função de retorno do rpc
type BasicMethodHandler func(args Message) (Args interface{})

var conn Crossbar

// Message mensagem string crossbar
type Message map[string]interface{}

// Load converto interface to struct domain
func (m Message) Load(p interface{}) error {
	err := mapstructure.Decode(m, p)
	return err
}

// Start crossbar
func (c *Crossbar) Start() {
	// s := turnpike.NewBasicWebsocketServer("sbe")
	// server := &http.Server{
	// 	Handler: s,
	// 	Addr:    ":8000",
	// }
	// log.Println("turnpike server starting on port 8000")
	// log.Fatal(server.ListenAndServe())
	quit := make(chan interface{})
	<-quit
	<-quit
}

// Get crossbar instancia
func Get() Crossbar {
	return conn
}

// Subscribe register event subscribe crossbar
func (c *Crossbar) Subscribe(procedure string, fn EventHandler) {
	go func() {
		if err := c.Conn.Subscribe(procedure, nil, func(args []interface{}, kwargs map[string]interface{}) {

			for _, msg := range args {
				str := castMessage(msg)
				fn(Message(str))
			}

		}); err != nil {
			log.Fatalln("Error subscribing to chat channel:", err)
		}
	}()
}

// Register call rpc
func (c *Crossbar) Register(procedure string, fn BasicMethodHandler) {
	go func() {
		var call = func(args []interface{}, kwargs map[string]interface{}, details map[string]interface{}) (result *turnpike.CallResult) {
			t := &turnpike.CallResult{}

			if len(args) > 0 {
				str := castMessage(args[0])
				rs := fn(Message(str))
				t = &turnpike.CallResult{Args: []interface{}{rs}}
			} else {
				t = &turnpike.CallResult{Args: []interface{}{nil}}
			}

			return t
		}

		if err := c.Conn.Register(procedure, call, map[string]interface{}{"invoke": "roundrobin"}); err != nil {
			log.Fatalln("Error register call:", err)
		}
	}()
}

func castMessage(arg interface{}) map[string]interface{} {
	str := make(map[string]interface{})

	var a string
	var b map[string]interface{}

	// compare string
	if reflect.TypeOf(a) == reflect.TypeOf(arg) {
		a = arg.(string)
		json.Unmarshal([]byte(a), &str)
	} else
	// compare map[string]interface{}
	if reflect.TypeOf(b) == reflect.TypeOf(arg) {
		str = arg.(map[string]interface{})
	}

	return str
}

// Call chamada rpc crossbar
func (c *Crossbar) Call(procecure string, arg interface{}) <-chan Message {
	messages := make(chan Message)
	go func() {
		c1, err := turnpike.NewWebsocketClient(turnpike.JSON, "ws://192.168.231.183:8080/ws", nil, nil, nil)
		defer c1.Close()

		if err != nil {
			fmt.Println(err)
		}
		_, err = c1.JoinRealm("sbe", nil)
		rs, err := c1.Call(procecure, map[string]interface{}{"disclose_me": true}, []interface{}{arg}, nil)

		if err != nil {
			fmt.Println("error setting alarm:", err)
			messages <- nil
		} else {
			messages <- rs.Arguments[0].(map[string]interface{})
		}
	}()
	return messages
}
