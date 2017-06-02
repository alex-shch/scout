package main

import (
	"fmt"
	"time"
)

func GameLoop() {
	for {
		msg := fmt.Sprintf("%v", time.Now())

		for _, conn := range connections {
			if err := conn.send([]byte(msg)); err != nil {
				// remove connection
			}
		}

		time.Sleep(1000 * time.Millisecond)
	}
}
