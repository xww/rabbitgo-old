package conf

import (
	"github.com/BurntSushi/toml"

	"fmt"
)

func InitConfig() map[string]interface{} {
	var config map[string]interface{}
	if _, err := toml.DecodeFile("conf\\rabbitgo.conf", &config); err != nil {
		//err := errors.New("cannot find config file")
		fmt.Println(err.Error())

	}

	return config
}
