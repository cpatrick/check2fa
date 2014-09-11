package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"code.google.com/p/goauth2/oauth"
	"github.com/google/go-github/github"
)

func main() {

	var org string

	tokenPtr := flag.String("token", "<your token here>",
		"OAuth2 token for accessing Github.")

	flag.Parse()
	if len(flag.Args()) < 1 {
		fmt.Println("Please provide an organization to audit for two-factor users.")
		os.Exit(1)
	} else {
		org = flag.Args()[0]
	}

	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: *tokenPtr},
	}

	fmt.Println("Printing two-factor status for:", org)
	client := github.NewClient(t.Client())
	opts := &github.ListMembersOptions{
		Filter: "2fa_disabled",
	}
	users, _, err := client.Organizations.ListMembers(org, opts)
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range users {
		fmt.Println(*user.Login)
	}
}
