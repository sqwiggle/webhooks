package test_servers

import(
	"fmt"
	"net/http"
	"github.com/awsmsrc/llog"
)

func TestServer200(port int) {
	llog.Debugf("Starting 200 test server on %d", port)
	sm := http.NewServeMux()
	sm.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Success")
	})
	http.ListenAndServe(fmt.Sprintf(":%d", port), sm)
}

func TestServer204(port int) {

}

func TestServer400(port int) {

}

func TestServer404(port int) {
	llog.Debugf("Starting 404 test server on %d", port)
	sm := http.NewServeMux()
	sm.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", 404)
	})
	http.ListenAndServe(fmt.Sprintf(":%d", port), sm)
}

func TestServer405(port int) {
	llog.Debugf("Starting 405 test server on %d", port)
	sm := http.NewServeMux()
	sm.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Method Not Allowed", 405)
	})
	http.ListenAndServe(fmt.Sprintf(":%d", port), sm)
}

func TestServer500(port int) {

}

func testServer () {

}
