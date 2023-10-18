package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	go runServer(7001, "a")
	go runServer(7002, "b")

	go runProxyServer()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-c
}

type ServerHandler struct {
	Name string
}

func (s *ServerHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/api" {
		byt, _ := json.Marshal(map[string]interface{}{
			"msg": fmt.Sprintf("server name: %s", s.Name),
			"data": map[string]interface{}{
				"headers":    req.Header,
				"method":     req.Method,
				"host":       req.Host,
				"requestURI": req.RequestURI,
				"url":        req.URL,
				"remoteAddr": req.RemoteAddr,
			},
		})
		resp.Write(byt)
		return
	}
	resp.WriteHeader(400)
	resp.Write([]byte("404"))
}

func runServer(port int, name string) {
	h := &ServerHandler{
		Name: name,
	}
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), h)
	if err != nil {
		panic(err)
	}
}

type ProxyServer struct {
	*httputil.ReverseProxy
}

func (s *ProxyServer) run() {
	err := http.ListenAndServe(":7000", s)
	if err != nil {
		panic(err)
	}
}

func runProxyServer() {
	reverseProxy := &httputil.ReverseProxy{}
	reverseProxy.Director = func(req *http.Request) {
		log.Println(req.RequestURI)
		rs := ""
		if strings.HasPrefix(req.RequestURI, "/a") {
			rs = "http://127.0.0.1:7001/api"
		} else if strings.HasPrefix(req.RequestURI, "/b") {
			rs = "http://127.0.0.1:7002/api"
		}
		targetURL, _ := url.Parse(rs)
		req.URL.Scheme = targetURL.Scheme
		req.URL.Host = targetURL.Host
		req.URL.Path = targetURL.Path
		req.Header.Set("User-Agent", "from-proxy")
	}
	serv := &ProxyServer{
		reverseProxy,
	}
	serv.run()
}
