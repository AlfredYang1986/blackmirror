package bmalioss

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmredis"
	"github.com/hashicorp/go-uuid"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

const appid = "LTAINO7wSDoWJRfN"
const appsec = "PcDzLSOE86DsnjQn8IEgbaIQmyBzt6"

func QuerySTSToken() (BmSTS, error) {
	reval, err := querySTSToken()
	if err == nil {
		return reval, err
	} else {
		return queryRemoteSTSToken()
	}
}

func queryRemoteSTSToken() (BmSTS, error){

	t := time.Now()
	tm := t.UTC()
	tmp := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02dZ", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second())
	fmt.Println(tmp)
	sn, err := uuid.GenerateUUID()

	cm := map[string]string{
		"Format": "JSON",
		"SignatureVersion": "1.0",
		"AccessKeyId": "LTAINO7wSDoWJRfN",
		"Version": "2015-04-01",
		"Action": "AssumeRole",
		"SignatureNonce": url.QueryEscape(sn),
		"SignatureMethod": "HMAC-SHA1",
		"RoleSessionName": "blackmirror",
		"RoleArn": url.QueryEscape("acs:ram::1696525727043311:role/aliyunosstokengeneratorrole"),
		"Timestamp": url.QueryEscape(tmp),
	}

	/**
	 * sort map by key
	 */
	var keys []string
	for k := range cm {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	cqs := ""
	for idx, key := range keys {
		if idx != 0 {
			cqs += "&"
		}
		cqs += key + "=" + cm[key]
	}

	//cqs = url.QueryEscape(cqs)
	string2sign := "GET" + "&" + url.QueryEscape("/") + "&" + url.QueryEscape(cqs)
	fmt.Println(cqs)
	fmt.Println(string2sign)

	sign := computeHmac256(string2sign, appsec + "&")
	fmt.Println(sign)

	cqs += "&Signature=" + url.QueryEscape(sign)

	url := strings.Join([]string{"https://sts.aliyuncs.com/?", cqs}, "")

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Read failed:", err)
		return BmSTS{}, err
	}
	resp.Body.Close()

	fmt.Println(string(content))

	client := bmredis.GetRedisClient()
	defer client.Close()

	var reval BmSTS
	dnc := json.NewDecoder(strings.NewReader(string(content)))
	err = dnc.Decode(&reval)
	fmt.Println(reval)

	client.HSet("dongda-oss-key", "AccessKeyId", reval.GetAccessKeyId())
	client.HSet("dongda-oss-key", "AccessKeySecret", reval.GetAccessKeySecret())
	client.HSet("dongda-oss-key", "SecurityToken", reval.GetSecurityToken())

	client.Expire("dongda-oss-key", 59 * time.Minute)

	return reval, err
}

func computeHmac256 (message string, secret string) string {
	hmac := hmac.New(sha1.New, []byte(secret))
	hmac.Write([]byte(message))
	encoded := base64.StdEncoding.EncodeToString([]byte(hmac.Sum(nil)))
	return encoded
}

func querySTSToken () (BmSTS, error) {
	client := bmredis.GetRedisClient()
	defer client.Close()

	var reval BmSTS
	tmp, err := client.HGetAll("dongda-oss-key").Result()
	if err == nil && len(tmp) != 0 {
		reval.ResetSecurityProp(tmp)
		return reval, nil
	}

	return reval, errors.New("redis get dongda-oss-key error")
}