package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

var echoAddr string

func echoHandler(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	msg := r.URL.Query().Get("message")
	log.Printf("new message: %v\n", msg)
	req, err := http.NewRequest("GET", fmt.Sprintf("%v/echo", echoAddr), nil)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	q := req.URL.Query()
	q.Add("message", msg)
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	defer resp.Body.Close()
	response, _ := ioutil.ReadAll(resp.Body)
	log.Printf("echo server response: %v\n", response)
	w.Write([]byte(strings.ToUpper(string(response))))
}

func main() {
	flag.StringVar(&echoAddr, "echo-addr", "", "HTTP listen address")
	flag.Parse()
	log.Printf("Echo addr: %v\n", echoAddr)
	http.HandleFunc("/upper", echoHandler)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal(err)
		}
	}()
	go func() {
		defer wg.Done()
		c := make(chan os.Signal, 1)
		defer close(c)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
		sig := <-c
		log.Printf("Shutdown complete: %v signal", sig)
	}()
	wg.Wait()
}
