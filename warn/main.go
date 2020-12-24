package main

import (
	"fmt"
	"github.com/rz1226/prototypes/warn/tickwarn"
	"time"
)

func main() {
	f := func() {
		fmt.Println("warn......", time.Now().Format("2006-01-02 15:04:05"))
	}
	warn := tickwarn.NewTickWarn(20, 10, f)

	for i := 0; i < 1000; i++ {
		fmt.Println(i)
		if i == 5 {
			warn.Tick()
			time.Sleep(time.Second * 21)
		} else {
			warn.Tick()
			time.Sleep(time.Second * 19)
		}

	}

}
