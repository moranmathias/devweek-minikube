package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("message")
	log.Printf("new message: %v\n", msg)
	w.Write([]byte(msg))
}

func main() {
	http.HandleFunc("/echo", echoHandler)
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
