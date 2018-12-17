package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/striversity/glft/shared/cal"
)

const userDb = "./user.db"

func main() {
	fmt.Print("What is your name: ")
	var username string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		username = scanner.Text()
	}

	// clean up input
	username = strings.TrimSpace(username)
	username = strings.Title(username)
	printGreeting(username)

	if !userExist(username) {
		addUser(userDb, username)
	}
}

func addUser(fn string, username string) {
	// try to open existing user db file, if that failes, then create a new file
	f, err := os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if nil != err {
		return
	}

	defer f.Close()
	f.WriteString(username + "\n")
}

func userExist(username string) bool {
	// load users from file
	users := getUsers(userDb)

	for _, s := range users {
		if s == username {
			return true
		}
	}

	return false
}

func getUsers(fn string) (users []string) {
	f, err := os.Open(fn)
	if nil != err {
		//log.Println("INFO: No users in db")
		return
	}

	defer f.Close()

	// if no error opening file, try to read each line into a slice
	var s string
	reader := bufio.NewReader(f)
	s, err = reader.ReadString('\n')
	for nil == err {
		users = append(users, strings.TrimSpace(s))
		s, err = reader.ReadString('\n')
	}

	return
}

func printGreeting(username string) {
	fmt.Printf("Hi, %s\n", username)

	if userExist(username) {
		fmt.Println("Welcome back!")
	} else {
		fmt.Println("Welcome first time user!")
	}

	fmt.Printf("Today is %s. The current time is %s.\n", cal.DateNow(), cal.TimeNow())
}
