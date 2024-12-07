package main

import (
	"errors"
	"fmt"
	"runtime"
	"time"

	"github.com/Jeffail/tunny"
)

func SendEmail(email string, subject string, body string) {
	fmt.Printf("Sending email to %s\n", email)
	fmt.Printf("Subject %s\n Body: %s\n", subject, body)
	// Simulate sending email
	time.Sleep(2 * time.Second)
}

func main() {
	numCPUs := runtime.NumCPU()
	fmt.Printf("Number of CPUs: %d\n\n", numCPUs)

	pool := tunny.NewFunc(numCPUs, func(payload any) any {
		m, ok := payload.(map[string]any)
		if !ok {
			return errors.New("unable to extract map")
		}

		// Extract the fields
		email, ok := m["email"].(string)
		if !ok {
			return errors.New("email field is missing or not a string")
		}

		subject, ok := m["subject"].(string)
		if !ok {
			return errors.New("subject field is missing or not a string")
		}

		body, ok := m["body"].(string)
		if !ok {
			return errors.New("body field is missing or not a string")
		}

		SendEmail(email, subject, body)
		return nil
	})
	defer pool.Close()

	for i := 0; i < 100; i++ {
		var data any = map[string]any{
			"email":   fmt.Sprintf("email%d@sample.io", i+1),
			"subject": "Welcome",
			"body":    "Thank you for signing up",
		}
		go func() {
			result := pool.Process(data)
			if result == nil {
				fmt.Println("Mail sent!")
			}
		}()
	}

	for {
		qLen := pool.QueueLength()
		fmt.Printf("----------------- Queue Length: %d\n", qLen)
		if qLen == 0 {
			break
		}
		time.Sleep(1 * time.Second)
	}
	time.Sleep(3 * time.Second)
}
