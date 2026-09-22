package main

import (
	"fmt"
	"time"
)

const (
	UU  = "Asia/Irkutsk"
	MSK = "Europe/Moscow"
)

func main() {
	now := time.Now()

	mskLoc, err := time.LoadLocation(MSK)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("Moscow time: %s\n", now.In(mskLoc).Format(time.RFC3339))

	uuLoc, err := time.LoadLocation(UU)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("Ulan-Ude time: %s\n", now.In(uuLoc).Format(time.RFC1123))
}
