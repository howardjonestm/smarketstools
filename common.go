package smarketstools

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Client struct {
	ApiToken string
}

//Read token from home directory
func ReadToken() string {

	homeDir := os.Getenv("HOME")
	tokenDirectory := fmt.Sprintf("%s/.SMARKETS_TOKEN", homeDir)

	t, err := ioutil.ReadFile(tokenDirectory)
	token := string(t)
	checkErr(err)
	return token
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
