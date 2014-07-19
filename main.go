package main

import (
	"github.com/awsmsrc/llog"
	"net/http"
)

var (
	db           = new(DB)
	queue        = make(chan Event, 5)
	worker_limit = 100
)

func init() {
	StartTestServers()
	db.Init()
}

func main() {
	llog.Info("main")
	defer db.Close()

	go EventDispatcher(queue, worker_limit)

	http.Handle("/", InitRouter())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		llog.FATAL(err)
	}
}
