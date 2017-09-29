package utils

import (
	"fmt"
	"sync"
)

type Job func()

type Done func()

func Step(wg *sync.WaitGroup, jobs ...Job) {
	wg.Add(len(jobs))
	for _,job := range jobs {
		go func(j Job) {
			defer func() {
				fmt.Println("Job is down: ")
				wg.Done()
			}()
			j()
		}(job)
	}

	wg.Wait()
}

func Last() {

}
