package appsettings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// AppSettings estrutura config
type AppSettings struct {
	configuration []byte
}

var _app AppSettings

func init() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	loadSettings(dir)
}

// SetConfigurationDesenv set config de desenv
func SetConfigurationDesenv(dir string, app interface{}) {
	if _app.configuration == nil {
		loadSettings(dir)
	}
	get(app)
}

// Get appsettings
func get(app interface{}) {
	json.Unmarshal(_app.configuration, app)
}

func loadSettings(dir string) {
	app, err := ioutil.ReadFile(dir + "\\appsettings.json")
	if err != nil {
		fmt.Println("Error load appsettings:", err)
	} else {
		_app.configuration = app
	}
}
