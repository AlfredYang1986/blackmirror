package bmredis

import (
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"os"
	"testing"
	"time"
)

func TestGetRedisClient(t *testing.T) {

	_ = os.Setenv("BM_REDIS_HOST", "192.168.100.176")
	_ = os.Setenv("BM_REDIS_PORT", "6379")
	_ = os.Setenv("BM_REDIS_PASS", "")
	_ = os.Setenv("BM_REDIS_DB", "0")

	c := GetRedisClient()
	c.Set("HELLO", "3Q VERY MUCH", 10 * time.Minute)
	result, err := c.Get("HELLO").Result()
	bmerror.PanicError(err)
	fmt.Println(result)

}
