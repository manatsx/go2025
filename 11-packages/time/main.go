package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	fmt.Println(now)

	then := time.Date(2019, 11, 17, 20, 34, 58, 651387237, time.UTC)

	fmt.Println(then)

	then = then.Add(time.Hour * 1)
	fmt.Println(then)

	fmt.Println(then.Year())
	fmt.Println(then.Month())
	fmt.Println(then.Day())
	fmt.Println(then.Hour())
	fmt.Println(then.Minute())
	fmt.Println(then.Second())
	fmt.Println(then.Nanosecond())
	fmt.Println(then.Location())
	fmt.Println(then.Weekday())

	fmt.Println("Then es antes que Now:", then.Before(now))
	fmt.Println("Then es desfmt.Printlnues que Now:", then.After(now))
	fmt.Println("Then es igual que Now:", then.Equal(now))

	diff := now.Sub(then)
	fmt.Println(diff)

	fmt.Println(diff.Hours())
	fmt.Println(diff.Minutes())
	fmt.Println(diff.Seconds())
	fmt.Println(diff.Nanoseconds())
}
