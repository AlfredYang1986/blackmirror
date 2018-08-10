package bmconf

import (
	"encoding/json"
	//"fmt"
	"io/ioutil"
	"log"
	"strings"
	"sync"
)

type BMBrickConf struct {
	Name   string `json:"name"`
	Host   string `json:"host"`
	Port   int    `json:"Port"`
	Router string `json:"router"`
}

var brickconf map[string]BMBrickConf = make(map[string]BMBrickConf)
var once sync.Once

func GetBMBrickConf(n string) BMBrickConf {
	once.Do(initedConf)
	return brickconf[n]
}

func initedConf() {
	b, _ := ioutil.ReadFile("resource/conf.json")
	jsonStream := string(b)
	dec := json.NewDecoder(strings.NewReader(jsonStream))

	_, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	for dec.More() {
		var conf BMBrickConf
		err := dec.Decode(&conf)
		if err != nil {
			log.Fatal(err)
		}
		brickconf[conf.Name] = conf
	}

	_, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
}
