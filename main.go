package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

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

func commandos() {
	if len(os.Args) == 1 {
		fmt.Println("imprimir ayuda")
		os.Exit(2)
	}
	switch strings.ToLower(os.Args[1]) {
	case "post":
		fmt.Println("POST")
	case "get":
		fmt.Println("GET")
	case "update":

	case "delete":

	default:
		fmt.Println("ayuda")
	}

}

func main() {

	configuracion()

	commandos()

}
