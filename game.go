package main

import (
	"fmt"
	"time"
)

func GameLoop() {
	for {
		msg := fmt.Sprintf("%v", time.Now())

		for _, conn := range connections {
			conn.send([]byte(msg))
		}

		time.Sleep(1000 * time.Millisecond)
	}
}
