package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

type (
	Worker struct {
		ID                      int
		MaxRequests             int
		RequestOptions          sendRequestOptions
		MaxDelayBetweenRequests int
	}

	Manager struct {
		workers       []*Worker
		WG            *sync.WaitGroup
		chErrorCount  chan int
		errorCount    int
		maxErrorCount int

		mutex sync.Mutex
		h     *handle
	}
)

func (w *Worker) Do(m *Manager) {
	defer m.WG.Done()

	requestsSent := 0
	for {
		delay := rand.Intn(w.MaxDelayBetweenRequests)

		// Since we may not want to float all the requets at once.
		if w.MaxDelayBetweenRequests != 0 {
			time.Sleep(time.Duration(delay) * time.Second)
		}

		requestsSent++

		requestOptions := w.RequestOptions
		requestOptions.UserAgent = m.h.UserAgents.GetRandomElement()

		requestOptions.FollowRedirects = m.h.Config.Request.FollowRedirects

		if m.h.Logins == nil {
			requestOptions.UserName = generateEmail()
		} else {
			requestOptions.UserName = m.h.Logins.GetRandomElement()
		}

		if m.h.Passwords == nil {
			requestOptions.Password = generatePassword()
		} else {
			requestOptions.Password = m.h.Passwords.GetRandomElement()
		}

		if m.h.Config.VeboseMode {
			fmt.Println("sending login:", requestOptions.UserName, "/", requestOptions.Password)
		}

		resp, err := sendRequest(requestOptions)

		if m.h.Config.VeboseMode {
			fmt.Println("request response: ", resp.StatusCode)
			for h, v := range resp.Header {
				fmt.Println(h, ":", v)
			}
		}

		m.mutex.Lock()
		m.h.Stats.RequestsMade++
		if err != nil {
			fmt.Println("failed sending request: ", err)
			m.errorCount++
			m.h.Stats.RequestsFailed++
		} else {
			m.h.Stats.RequestsSuccessed++
			m.errorCount = 0
		}
		m.mutex.Unlock()

		m.chErrorCount <- m.errorCount

		fmt.Printf("[%s] request %s sent to: %s %d[%d:%d] elapsed time: %s | errors: %d\n",
			time.Now().Format("01/02 03:04:05PM"),
			requestOptions.Method,
			requestOptions.URL,
			w.ID,
			w.MaxRequests,
			requestsSent,
			time.Since(m.h.Stats.StartTime),
			m.h.Stats.RequestsFailed,
		)

		if w.MaxRequests > 0 && requestsSent >= w.MaxRequests {
			break
		}
	}

}

func newManager(maxErrorCount int, h *handle) *Manager {
	return &Manager{
		chErrorCount:  make(chan int, 1),
		maxErrorCount: maxErrorCount,
		workers:       make([]*Worker, 0),
		WG:            &sync.WaitGroup{},
		h:             h,
	}
}

func (m *Manager) AddWorker(w *Worker) {
	m.workers = append(m.workers, w)
}

func (m *Manager) Start() {
	m.WG.Add(len(m.workers))
	for _, w := range m.workers {
		go w.Do(m)
	}

	// Supervisor that determins if we should kill the process.
	if m.maxErrorCount > 0 {
		go func() {
			for {
				ec := <-m.chErrorCount

				if ec >= m.maxErrorCount {
					fmt.Println("max error count reached! Quiting executation!")
					os.Exit(1)
				}
			}
		}()
	}

	m.WG.Wait()
}
