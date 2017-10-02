package utils

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestWaitSingleJobDone(t *testing.T) {
	status := false
	var wg sync.WaitGroup

	WaitSingleJobDone(&wg, func(wg *sync.WaitGroup) {
		status = true
		wg.Done()
	})

	if false == status {
		t.Error("testing failed")
	}
}

func TestWaitJobsQueueDone(t *testing.T) {
	for i := 0; i < 100; i++ {
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(1234567890)
		rr := r
		var wg sync.WaitGroup
		WaitJobsQueueDone(&wg, func(wg *sync.WaitGroup) {
			r = r + 1
			wg.Done()
		}, func(wg *sync.WaitGroup) {
			r = r * 2
			wg.Done()
		}, func(wg *sync.WaitGroup) {
			r = r - 3
			wg.Done()
		}, func(wg *sync.WaitGroup) {
			r = r * 4
			wg.Done()
		}, func(wg *sync.WaitGroup) {
			r = r / 5
			wg.Done()
		})

		expected := ((((rr + 1) * 2) - 3) * 4) / 5
		if expected != r {
			t.Error("testing failed")
		}
	}
}

func TestWaitJobsDone(t *testing.T) {
	r1, r2, r3, r4, r5 := 0, 0, 0, 0, 0
	var wg sync.WaitGroup
	WaitJobsDone(&wg, func(wg *sync.WaitGroup) {
		r1++
		wg.Done()
	}, func(wg *sync.WaitGroup) {
		r2++
		wg.Done()
	}, func(wg *sync.WaitGroup) {
		r3++
		wg.Done()
	}, func(wg *sync.WaitGroup) {
		r4++
		wg.Done()
	}, func(wg *sync.WaitGroup) {
		r5++
		wg.Done()
	})

	if 1 != r1 {
		t.Error("testing failed")
	}
	if 1 != r2 {
		t.Error("testing failed")
	}
	if 1 != r3 {
		t.Error("testing failed")
	}
	if 1 != r4 {
		t.Error("testing failed")
	}
	if 1 != r5 {
		t.Error("testing failed")
	}
}
