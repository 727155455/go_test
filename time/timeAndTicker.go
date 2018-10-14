package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	timer1 := time.NewTimer(time.Second * 2);
	ticker1 := time.NewTicker(time.Second * 2);
	var wg sync.WaitGroup
	wg.Add(2)


	go func(t *time.Ticker) {
		defer wg.Done()
		for {
			<-t.C
			fmt.Println("get ticker1", time.Now().Format("2006-01-02 15:04:05"))
		}
	}(ticker1)

	go func(t *time.Timer) {
		defer wg.Done()
		for {
			<-t.C
			fmt.Println("get timer", time.Now().Format("2006-01-02 15:04:05"))
			//Reset 使 t 重新开始计时，（本方法返回后再）等待时间段 d 过去后到期。如果调用时t
			//还在等待中会返回真；如果 t已经到期或者被停止了会返回假。
			t.Reset(2 * time.Second)
		}
	}(timer1)

	wg.Wait()

}
