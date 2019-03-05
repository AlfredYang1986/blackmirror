package bmconfig

import (
	"bufio"
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"os"
)

func GetConfigMap(configPath string) map[string]interface{} {

	var profile string
	flag.StringVar(&profile, "profile", configPath, "Full path of the profile.")
	flag.Parse()

	var configMap map[string]interface{}

	profileFD, err := os.Open(profile)
	if err != nil {
		panic(err)
	}
	defer profileFD.Close()

	buffer := bufio.NewReader(profileFD)
	var profileLines string
	for {
		line, err := buffer.ReadString('\n')
		if err == io.EOF {
			break
		} else if line[0] == '#' || line[0] == ';' {
			continue
		} else if err != nil {
			panic(err)
		}

		profileLines += line
	}

	jsonLines := []byte(profileLines)

	if err := json.Unmarshal(jsonLines, &configMap); err != nil {
		panic(err)
	}
	return configMap
}

func BMGetConfigMap(configPath string) map[string]interface{} {
	var configMap map[string]interface{}
	b, _ := ioutil.ReadFile(configPath)
	if err := json.Unmarshal(b, &configMap); err != nil {
		errstr := configPath + " => " + err.Error()
		panic(errstr)
	}
	return configMap
}
