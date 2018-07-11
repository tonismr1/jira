package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/tonismr1/jira/Entidades"
)

var config entidades.Configuracion

func configuracion() {
	file, e := ioutil.ReadFile("./config.json")
	if e != nil {
		config = entidades.Configuracion{"localhost", 1994, "admin", "admin"}
		test, _ := json.Marshal(config)
		ioutil.WriteFile("./config.json", test, 0644)
	}

	json.Unmarshal(file, &config)
}

func main() {

	configuracion()

	fmt.Println(config.Url)

}
