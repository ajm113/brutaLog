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
		MaxRequests             int
		RequestOptions          sendRequestOptions
		MaxDelayBetweenRequests int
	}

	Manager struct {
		workers       []*Worker
		WG            *sync.WaitGroup
		chErrorCount  chan int
		ErrorCount    int
		maxErrorCount int

		UserAgents RainbowTable
		Emails     RainbowTable
		Passwords  RainbowTable
	}
)

func (w *Worker) Do(m *Manager) {
	defer m.WG.Done()

	requestsSent := 0
	for {
		requestsSent++
		requestOptions := w.RequestOptions

		requestOptions.UserAgent = m.UserAgents.GetRandomElement()
		requestOptions.UserName = m.Emails.GetRandomElement()

		if m.Passwords == nil {
			requestOptions.Password = generatePassword()
		} else {
			requestOptions.Password = m.Passwords.GetRandomElement()
		}

		fmt.Println("sending request", time.Now().Format(time.RFC822), requestOptions.Method, requestOptions.URL)
		_, err := sendRequest(requestOptions)

		// @TODO: Make this work with verbose mode.
		// fmt.Println("response code:", resp.StatusCode)
		// for h, v := range resp.Header {
		// 	fmt.Println(h, ":", v)
		// }

		if err != nil {
			fmt.Println("failed sending request for the following: ", err)
			m.ErrorCount++
			m.chErrorCount <- m.ErrorCount
		} else {
			m.chErrorCount <- m.ErrorCount
		}

		if w.MaxRequests > 0 && requestsSent >= w.MaxRequests {
			break
		}

		delay := rand.Intn(w.MaxDelayBetweenRequests)

		time.Sleep(time.Duration(delay) * time.Second)
	}

}

func newManager(maxErrorCount int, userAgents RainbowTable, emails RainbowTable, passwords RainbowTable) *Manager {
	return &Manager{
		chErrorCount:  make(chan int, 1),
		maxErrorCount: maxErrorCount,
		workers:       make([]*Worker, 0),
		WG:            &sync.WaitGroup{},
		UserAgents:    userAgents,
		Emails:        emails,
		Passwords:     passwords,
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
