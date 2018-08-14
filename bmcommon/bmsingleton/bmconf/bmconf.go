package bmconf

import (
	"encoding/json"
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"io/ioutil"
	"log"
	//"reflect"
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

func GetBMBrick(n string) (bmpipe.BMBrickFace, bmpipe.BMBrickExtends) {
	once.Do(initedConf)

	fac := bmsingleton.GetFactoryInstance()

	var name string
	var extends string
	sp := strings.Split(n, ":")
	if len(sp) == 0 {
		name = n
		extends = ""
	} else {
		name = sp[0]
		extends = sp[1]
	}

	bks, err := fac.ReflectPointer(name)
	if err != nil {
		panic(err)
	}

	face, ok := bks.(bmpipe.BMBrickFace)
	if !ok {
		panic(ok)
	}

	bke, err := fac.ReflectPointer(extends)
	ext, ok := bke.(bmpipe.BMBrickExtends)
	if err != nil || !ok {
		panic(err)
	}
	if !ok {
		panic(ok)
	}

	return face, ext
}

func initedConf() {
	fmt.Println("start of init conf")
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
	fmt.Println("init conf success")
}
