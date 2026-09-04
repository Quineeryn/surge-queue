package worker

import (
	"log"
	"time"
)

type EmailJob struct {
	Email, Name string
}

func StartEmailWorker(workerID int, jobs <-chan EmailJob) {
	for job := range jobs {
		time.Sleep(1 * time.Second)
		log.Printf("[EMAIL SENT] by workerID: %d | Welcome to %s <%s>", workerID, job.Name, job.Email)
	}
}
