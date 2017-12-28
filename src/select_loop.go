package main

import "time"
import "github.com/cihub/seelog"

func main() {
	ticker := time.NewTicker(60000 * time.Millisecond)
	bufferChan := make(chan int, 20)

	go func() {
		for {
			time.Sleep(10 * time.Minute)
			bufferChan <- 1
		}
	}()

	for {
		select {
		case now := <-ticker.C:
			seelog.Debugf("now:%v", now)
		case <-bufferChan:
			seelog.Debugf("buffer chan")
			//default:
			//	seelog.Debugf("default")
		}
	}
}
