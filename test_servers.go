package main

import "github.com/sqwiggle/webhooks/test_servers"

func StartTestServers() {
	go test_servers.TestServer200(8200)
	go test_servers.TestServer404(8404)
	go test_servers.TestServer405(8405)
}
