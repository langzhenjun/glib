package utils

import (
	"time"
)

func doTick(job func(), ticker *time.Ticker) {
	go job()
	<-ticker.C
}

func tickDo(ticker *time.Ticker, job func()) {
	<-ticker.C
	go job()
}

func DoTick(job func(), d time.Duration, count int) {
	if d <= 0 {
		return
	}

	ticker := time.NewTicker(d)

	if count > 0 {
		for count > 0 {
			doTick(job, ticker)
			count--
		}
	} else {
		for {
			doTick(job, ticker)
		}
	}
}

func TickDo(job func(), d time.Duration, count int) {
	if d <= 0 {
		return
	}

	ticker := time.NewTicker(d)

	if count > 0 {
		for count > 0 {
			tickDo(ticker, job)
			count--
		}
	} else {
		for {
			tickDo(ticker, job)
		}
	}
}
