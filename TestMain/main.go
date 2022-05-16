package main

import (
	"fmt"
	"time"

	"./nettop" // change your mod name
)

func main() {

	bandwitch := nettop.NewNetTop()
	go bandwitch.Start()

	done := make(chan bool)
	go func() {
		for {
			time.Sleep(time.Second * 1)

			if bandwitch != nil {
				a1 := bandwitch.Getlast()
				if a1 != nil {
					if len(a1.Stat) != 0 {
						for k, v := range a1.Stat {
							fmt.Println(k, v.Name, v.Rx, v.Tx)
						}
					} else {
						// fmt.Println("a1.Stat is null")
					}
				} else {
					// fmt.Println("a1 is null")
				}
			}
		}
	}()
	<-done

}
