package main

import (
	"fmt"
	"testing"
	"time"
)

func TestSleep(t *testing.T) {
	var expTime int64 = 80
	start := time.Now().Round(time.Second)
	mySleepFunction(80)
	end := time.Now().Round(time.Second)
	var diff = end.Sub(start)
	var diffAsInt64 = int64(diff) / 1000000000
	fmt.Println(expTime)
	if diffAsInt64 != expTime {
		t.Error("Sleep function failed")
	}
}
