package main

import "flag"

func main() {

	var folder string
	var email string

	flag.StringVar(&folder, "add", "", "add a new folder for git scan")
	flag.StringVar(&email, "email", "your@email.com", "email for scanning")
	flag.Parse()

	if folder != "" {
		scan(folder)
		return
	}

	stats(email)
}
