package bmredis

import (
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"os"
	"testing"
	"time"
)

func TestGetRedisClient(t *testing.T) {

	os.Setenv("BM_REDIS_CONF_HOME", "../resource/redisconfig.json")

	c := GetRedisClient()
	c.Set("HELLO", "3Q VERY MUCH", 10 * time.Minute)
	result, err := c.Get("HELLO").Result()
	bmerror.PanicError(err)
	fmt.Println(result)

}
