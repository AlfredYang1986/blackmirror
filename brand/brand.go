package brand

import (
	"fmt"
)

type Alfred struct {
	Name      string            `json:"name"`
	Slogan    string            `json:"slogan"`
	Highlight []string          `json:"highlight"`
	About     string            `json:"about"`
	Awards    map[string]string `json:"awards"`
	Attends   map[string]string `json:attends`
	Qualifier map[string]string `json:"qualifier"`
}

func init() {
	// initialization code here
	fmt.Println("go lib")
}

func GetAlfred() (Alfred, error) {
	rst := Alfred{
		Name: "alfred"}

	return rst, nil
}
