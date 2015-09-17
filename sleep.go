package main

import (
	"fmt"
	"time"
)

func mySleepFunction(x int) {
	<-time.After(time.Second * time.Duration(x))
}

func main() {
	fmt.Println(time.Now())
	mySleepFunction(10)
	fmt.Println(time.Now())
}
