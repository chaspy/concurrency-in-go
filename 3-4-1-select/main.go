package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	c := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(c) // 5秒待ったあとチャネルを閉じる
	}()

	fmt.Println("Blocking on read...")
	select {
	case <-c: // チャネルの読み込みを試みる。実際には select は必要ない
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	}
}
