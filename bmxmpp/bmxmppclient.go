package bmxmpp

import (
	"crypto/md5"
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmconfighandle"
	"github.com/mattn/go-xmpp"
	"io"
	"log"
	"os"
	"strings"
)
var bmXmppConfig bmconfig.BmXmppConfig

var status = flag.String("status", "xa", "status")
var statusMessage = flag.String("status-msg", "status-msg", "status message")
var notls = flag.Bool("notls", true, "No TLS")
var starttls = flag.Bool("starttls", true, "Start TLS")
var debug = flag.Bool("debug", false, "debug output")
var session = flag.Bool("session", true, "use server session")

func serverName(host string) string {
	return strings.Split(host, ":")[0]
}

func Forward(userjid string, msg string) error {
	bmXmppConfig.GenerateConfig()
	var server = bmXmppConfig.Host + ":" + bmXmppConfig.Port
	var username = bmXmppConfig.LoginUser + "@" + bmXmppConfig.HostName
	var password = bmXmppConfig.LoginUserPwd

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: example [options]\n")
		flag.PrintDefaults()
		os.Exit(2)
	}
	flag.Parse()
	if username == "" || password == "" {
		if *debug && username == "" && password == "" {
			fmt.Fprintf(os.Stderr, "no username or password were given; attempting ANONYMOUS auth\n")
		} else if username != "" || password != "" {
			flag.Usage()
		}
	}

	xmpp.DefaultConfig = tls.Config{
		ServerName:         serverName(server),
		InsecureSkipVerify: true,
	}

	h := md5.New()
	io.WriteString(h, userjid)
	resource := fmt.Sprintf("%x", h.Sum(nil))
	//resource,_ := uuid.GenerateUUID()

	options := xmpp.Options{
		Host:          server,
		User:          username,
		Password:      password,
		NoTLS:         *notls,
		StartTLS:      *starttls,
		Debug:         *debug,
		Session:       *session,
		Resource:      resource,
		Status:        *status,
		StatusMessage: *statusMessage,
	}

	var talk *xmpp.Client
	var err error
	talk, err = options.NewClient()
	if err != nil {
		log.Fatal(err)
		return err
	}
	_, err = talk.Send(xmpp.Chat{Remote: userjid, Type: "chat", Text: msg})
	return err
}
