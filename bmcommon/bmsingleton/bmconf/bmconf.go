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
	Name string `json:"name"`
	Host string `json:"host"`
	Port int    `json:"Port"`
}

var brickconf map[string]BMBrickConf = make(map[string]BMBrickConf)
var once sync.Once

func GetBMBrickConf(n string) BMBrickConf {
	once.Do(initedConf)
	return brickconf[n]
}

func GetBMBrick(n string) (bmpipe.BMBrickFace, error) {
	once.Do(initedConf)

	fac := bmsingleton.GetFactoryInstance()

	name := n
	bks, err := fac.ReflectPointer(name)
	if err != nil {
		panic(err)
	}

	face, ok := bks.(bmpipe.BMBrickFace)
	if !ok {
		panic(ok)
	}

	return face, err
}

func initedConf() {
	fmt.Println("start of init conf")
	//TODO: 配置文件路径 待 用脚本指定dev路径和deploy路径
	b, _ := ioutil.ReadFile("resource/conf.json")
	jsonStream := string(b)
	dec := json.NewDecoder(strings.NewReader(jsonStream))

	_, err := dec.Token()
	if err != nil {
		log.Println("error token")
		log.Fatal(err)
	}

	for dec.More() {
		var conf BMBrickConf
		err := dec.Decode(&conf)
		if err != nil {
			log.Println("error loading")
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
