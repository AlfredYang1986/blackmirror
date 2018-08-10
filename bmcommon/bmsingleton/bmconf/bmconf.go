package bmconf

import (
	"encoding/json"
	"fmt"
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
	// read open bracket
	t, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)

	// while the array contains values
	for dec.More() {
		var conf BMBrickConf
		// decode an array value (Message)
		err := dec.Decode(&conf)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(conf)
		brickconf[conf.Name] = conf
	}

	// read closing bracket
	t, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)
}
