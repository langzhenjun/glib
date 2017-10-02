package utils

import (
	"sync"
)

// WaitJob a job to be waiting to done. if done, the job should call wg.Done()
type WaitJob func(wg *sync.WaitGroup)

// WaitSingleJobDone wait a job done
func WaitSingleJobDone(wg *sync.WaitGroup, job WaitJob) {
	wg.Add(1)
	go job(wg)
	wg.Wait()
}

// WaitJobsQueueDone wait jobs in queue done one by one
func WaitJobsQueueDone(wg *sync.WaitGroup, queue ...WaitJob) {
	for _, job := range queue {
		WaitSingleJobDone(wg, job)
	}
}

// WaitJobsDone wait all jobs done in no order
func WaitJobsDone(wg *sync.WaitGroup, jobs ...WaitJob) {
	wg.Add(len(jobs))
	for _, job := range jobs {
		go job(wg)
	}
	wg.Wait()
}
