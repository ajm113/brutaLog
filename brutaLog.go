package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	c := loadFlagIntoConfig()

	var (
		userAgentTable = defaultUserAgents
		emailTable     = defaultUserEmails
		passwordTable  []string
		err            error
	)

	// We call this so BrutaLog can generate unique requests each run.
	rand.Seed(time.Now().Unix())

	if c.RainbowTable.UserAgents != "" {
		userAgentTable, err = loadRainbowTableFromFile(c.RainbowTable.UserAgents)

		if err != nil {
			fmt.Printf("failed opening user agent rainbow table for following reason: %s", err)
			os.Exit(1)
		}
	}

	if c.RainbowTable.Emails != "" {
		emailTable, err = loadRainbowTableFromFile(c.RainbowTable.Emails)

		if err != nil {
			fmt.Printf("failed opening email rainbow table for following reason: %s", err)
			os.Exit(1)
		}
	}

	if c.RainbowTable.Passwords != "" {
		emailTable, err = loadRainbowTableFromFile(c.RainbowTable.Passwords)

		if err != nil {
			fmt.Printf("failed opening password table for following reason: %s", err)
			os.Exit(1)
		}
	}

	manager := newManager(c.QuitOnErrorCount, userAgentTable, emailTable, passwordTable)

	for i := 0; i < c.WorkerCount; i++ {
		manager.AddWorker(&Worker{
			MaxRequests:             c.RequestCount,
			MaxDelayBetweenRequests: c.MaxDelayPerRequest,
			RequestOptions: sendRequestOptions{
				URL:               c.URL,
				Method:            c.Request.Method,
				ContentType:       c.Request.ContentType,
				Timeout:           time.Duration(c.Request.Timeout),
				PasswordFieldName: c.Request.PasswordField,
				UserFieldName:     c.Request.UserField,
			},
		})
	}

	fmt.Printf("starting brutaLog on %s w/ %d worker(s)...\n", c.URL, c.WorkerCount)
	manager.Start()

	fmt.Println("job completed succesfully!")
}
