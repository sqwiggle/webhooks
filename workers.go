package main

import (
	"net/http"
	"github.com/awsmsrc/llog"
	"strings"
	"time"
	"sync"
)

func EventDispatcher(ch chan Event, worker_count int) {
	jobs := make(chan Event)
	for i := 0; i < worker_count; i++ {
		go Worker(jobs)
	}
	for {
		jobs <- <- ch
	}
}

func Worker(jobs chan Event) {
	for {
		go Work(<-jobs)
	}
}

func Work(e Event) {
	var wg sync.WaitGroup
	for _, registration := range e.Registrations() {
		wg.Add(1)
		go func (registration_id Registration, e Event) {
			defer wg.Done()
			i := 0
			for i < 10 {
				llog.Debugf("POSTING %v to %s", e.Data, registration.Url)
				resp, err := http.Post(registration.Url, "test/html", strings.NewReader(e.Data))
				if err != nil {
					llog.Error(err)
				} else {
					llog.Debugf("%v", resp)
				}
				db.Create(&Attempt{
					EventId:  e.Id,
					Status:   resp.StatusCode,
					Response: "TODO",
				})
				if (resp.StatusCode >= 00 && resp.StatusCode< 300){
					llog.Successf("WEBHOOK SUCCESS for %d to %s", e.Id, registration.Url)
					return
				}
				<- time.After(time.Duration(10) * time.Second)
				i++
			}
			llog.Errorf("WEBHOOK FAILURE for %d to %s", e.Id, registration.Url)
		}(registration, e)
	}
	wg.Wait()
	result, err  := db.Exec(
		"UPDATE events SET state = ? WHERE id =	?",
		"processed",
		e.Id,
	)
	if err != nil {
		llog.FATAL(err)
	}
	rows_affected, err := result.RowsAffected()
	llog.Debugf("ev update %v + %v", rows_affected, err) 
}


