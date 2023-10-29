package configs

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

var (
	//Config configuration of app
	Config AppConfigs
)

// init initiale config from ./conf.d/app.toml
func init() {

	fileName, err := filepath.Abs("./conf.d/app.toml")
	if err != nil {
		log.Println("file Error", err)

	}

	log.Println("file Name", fileName)

	buff, readFileError := os.ReadFile(fileName)

	if readFileError != nil {
		log.Println("Read file Error", readFileError)

	}
	log.Println("buff: ", toml.NewDecoder(strings.NewReader(string(buff))))

	configError := toml.NewDecoder(strings.NewReader(string(buff))).Decode(&Config)
	if configError != nil {
		log.Println("configError: ", err)

	}

}

type AppConfigs struct {
	Addr          string `toml:"addr"`
	Name          string `toml:"name"`
	Server        string `toml:"server"`
	Ports         []int  `toml:"ports"`
	ConnectionMax int    `toml:"connection_max"`
}
