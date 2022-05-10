package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type clock struct {
	name, port string
}

func readTime(port string) string {
	conn, err := net.Dial("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	var time string
	_, err = fmt.Fscanln(conn, &time)
	if err != nil {
		log.Fatal(err)
	}

	conn.Close()
	return time
}

func main() {
	timezones := make([]*clock, 0)

	for _, arg := range os.Args[1:] {
		args := strings.Split(arg, "=")
		if len(args) != 2 || args[0] == "" {
			fmt.Println("Usage: clockWall [Timezone]=localhost:<port>")
			os.Exit(1)
		}

		port := strings.Split(args[1], ":")
		if port[0] != "localhost" {
			fmt.Println("Usage: clockWall [Timezone]=localhost:<port>")
			os.Exit(1)
		}

		if _, err := strconv.Atoi(port[1]); err != nil {
			fmt.Println("Error with port : port should be a number")
			os.Exit(1)
		}

		timezones = append(timezones, &clock{args[0], args[1]})

	}

	for {
		fmt.Printf("\033[H\033[2J")
		for _, timezone := range timezones {
			fmt.Println(timezone.name, ":", readTime(timezone.port))
		}
		time.Sleep(1 * time.Second)

	}

}
