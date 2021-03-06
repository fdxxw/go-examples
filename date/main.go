package main

import (
	"fmt"
	"time"

	"github.com/spf13/cast"
)

func main() {
	fmt.Printf("今天是周几: %v \n", cast.ToInt(time.Now().Weekday()))

	now := time.Now()
	toDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Local().Location())
	fmt.Println(toDay)
	fmt.Println(toDay.Add(time.Hour * -48))
}
