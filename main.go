package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	pp "github.com/armon/go-proxyproto"
)

func main() {
	list, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	proxyList := &pp.Listener{
		Listener:           list,
		ProxyHeaderTimeout: time.Millisecond * 50,
		SourceCheck: pp.SourceChecker(func(addr net.Addr) (bool, error) {
			if addr == nil {
				return false, fmt.Errorf("no addr")
			}
			fmt.Printf("received request from %s\n", addr)
			return true, nil
		}),
	}
	http.Serve(proxyList, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("received request from %s\n", r.RemoteAddr)
		if r.Method != http.MethodGet {
			fmt.Printf("received invalid method, %s, request\n", r.Method)
			w.WriteHeader(500)
			return
		}
		w.Write([]byte("hello!"))
		w.WriteHeader(200)
	}))
}
