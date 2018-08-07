package main //authser

import (
	"github.com/alfredyang1986/blackmirror/bmser/authser"
	"net/http"
)

func main() {
	r := authser.GetRouter()
	http.ListenAndServe(":8080", r)
}
