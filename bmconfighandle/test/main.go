package main

import (
	"fmt"
	"github.com/colinmarc/hdfs"
	"strings"
)

var ProfileItems map[string]interface{} //obviously,this var will be used in many files

func main() {

	fmt.Println("start")

	//println(os.Getenv("GOETC"))
	//
	//configPath := "/home/jeorch/github/jeorch/go/src/github.com/alfredyang1986/blackmirror/bmconfighandle/resource/test.json"
	////ProfileItems = bmconfig.GetConfigMap(configPath)
	//ProfileItems = bmconfig.GetConfigMap2(configPath)
	//
	//httpPort := int(ProfileItems["httpPort"].(float64))
	//httpsPort := int(ProfileItems["httpsPort"].(float64))
	//tcpPort := ProfileItems["tcpPort"].(float64)
	//others := ProfileItems["others"].(map[string]interface{})
	//item1 := others["item1"].(string)
	//item2 := others["item2"].(bool)
	//
	//println(httpPort)
	//println(httpsPort)
	//println(tcpPort)
	//println(item1)
	//println(item2)

	client, _ := hdfs.New("192.168.100.137:9000")
	originPath := "/workData/Export"
	destPath := "/home/jeorch/github/jeorch/go/src/github.com/alfredyang1986/blackmirror/bmconfighandle/resource"

	//err := client.CopyToLocal(originPath, destPath)

	p, err := client.ReadDir(originPath)

	if err != nil {
		fmt.Println("error1")
		fmt.Println(err.Error())
	}
	for _,f := range p {
		if strings.HasSuffix(f.Name(), ".csv.gz") {
			err = client.CopyToLocal(originPath + "/" + f.Name(), destPath + "/" + f.Name())
		}
	}
	if err != nil {
		fmt.Println("error2")
		fmt.Println(err.Error())
	}



}