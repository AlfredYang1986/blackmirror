package brand

import "fmt"

type brand struct {
	name      string            `json:"name"`
	slogan    string            `json:"slogan"`
	highlight []string          `json:"highlight"`
	about     string            `json:"about"`
	awards    map[string]string `json:"awards"`
	attends   map[string]string `json:attends`
	qualifier map[string]string `json:"qualifier"`
}

func init() {
	// initialization code here
	fmt.Println("go lib")
}
