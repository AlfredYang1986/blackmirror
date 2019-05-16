package bmxmpp

import (
	"crypto/md5"
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/mattn/go-xmpp"
	"io"
	"log"
	"os"
)

var server string
var username string
var password string
var status = flag.String("status", "xa", "status")
var statusMessage = flag.String("status-msg", "status-msg", "status message")
var notls = flag.Bool("notls", true, "No TLS")
var starttls = flag.Bool("starttls", true, "Start TLS")
var debug = flag.Bool("debug", false, "debug output")
var session = flag.Bool("session", true, "use server session")

func (bxc *BmXmppConfig) Forward(userjid string, msg string) error {
	server = bxc.Host + ":" + bxc.Port
	username = bxc.LoginUser + "@" + bxc.HostName
	password = bxc.LoginUserPwd

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
		ServerName:         bxc.HostName,
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
