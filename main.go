package main

import (
	"flag"
	"fmt"
)

var (
	token = flag.String("token", "", "github token")
	user  = flag.String("user", "", "user to be invited")
	org   = flag.String("org", "", "org")
	team  = flag.String("team", "", "team")
)

func main() {
	flag.Parse()

	if checkParameter() != 0 {
		return
	}

	initGithub(*token)

	if err := invite(*user, *org, *team); err != nil {
		fmt.Println(err)
	}
}

func checkParameter() int {

	if *token == "" {
		fmt.Println("\"token\" parameter is required")
		return 1
	}

	if *user == "" {
		fmt.Println("\"user\" parameter is required")
		return 2
	}

	if *org == "" {
		fmt.Println("\"org\" parameter is required")
		return 3
	}

	if *team == "" {
		fmt.Println("\"team\" parameter is required")
		return 4
	}
	return 0
}
