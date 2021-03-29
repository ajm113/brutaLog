package main

import "flag"

type (
	rainbowTableConfig struct {
		UserAgents string
		Emails     string
		Passwords  string
	}

	requestConfig struct {
		Type          string
		Method        string
		ContentType   string
		UserField     string
		PasswordField string
		Timeout       int
	}

	bruteLogConfig struct {
		URL                string
		WorkerCount        int
		MaxDelayPerRequest int
		RequestCount       int
		QuitOnErrorCount   int
		VeboseMode         bool

		Request      requestConfig
		RainbowTable rainbowTableConfig
	}
)

func loadFlagIntoConfig() (c *bruteLogConfig) {
	targetURLPtr := flag.String("u", "", "Target URL to send requests")
	workerCount := flag.Int("w", 1, "Number of workers to send requests.")
	maxDelayPerRequest := flag.Int("d", 1, "Delay between requests")
	numberOfRequests := flag.Int("c", 1, "Number of requests to target. 0 = infinite")
	quitOnNumberOfErrors := flag.Int("e", 3, "Number of errors to occure on via connection or HTTP code to quit. 0 = infinit")

	HTTPMethod := flag.String("X", "POST", "POST method")
	HTTPContentType := flag.String("C", "application/x-www-form-urlencoded", "Content-Type header field.")
	HTTPTimeout := flag.Int("T", 2, "Number of seconds to timeout HTTP requests.")

	userFieldName := flag.String("U", "user", "User field name to send to server.")
	passwordFieldName := flag.String("P", "password", "Password field name to send to server.")

	userAgentList := flag.String("RA", "", "User agent lists to randomly choose from.")
	loginList := flag.String("RE", "", "Emails/Logins to use from.")
	passwordList := flag.String("RP", "", "Passwords to use from. If non supplied random strings will be generated.")
	verboseMode := flag.Bool("v", false, "Outputs valueable debugging information when cordinating attacks.")

	flag.Parse()

	c = &bruteLogConfig{
		URL:                *targetURLPtr,
		WorkerCount:        *workerCount,
		MaxDelayPerRequest: *maxDelayPerRequest,
		RequestCount:       *numberOfRequests,
		QuitOnErrorCount:   *quitOnNumberOfErrors,

		Request: requestConfig{
			Method:        *HTTPMethod,
			ContentType:   *HTTPContentType,
			UserField:     *userFieldName,
			PasswordField: *passwordFieldName,
			Timeout:       *HTTPTimeout,
		},
		RainbowTable: rainbowTableConfig{
			UserAgents: *userAgentList,
			Emails:     *loginList,
			Passwords:  *passwordList,
		},
		VeboseMode: *verboseMode,
	}

	return
}
