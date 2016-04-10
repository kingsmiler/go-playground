package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"bytes"
	"strings"
	"os"
	"os/exec"
)

func main() {
	//Gathers information to create a cli
	config, host := prompt()
	//Connects to the server
	client := connect2(config, host)
	//Handles interactions until "exit" is inputted
	runCommands(client)

}

//Retrieve information to build a client connection struct
func prompt() (*ssh.ClientConfig, string) {
	var host string
	var user string
	var password string
	fmt.Print("Enter your server domain, or ip address: ")
	fmt.Scanf("%s", &host)
	fmt.Print("Enter your username: ")
	fmt.Scanf("%s", &user)
	fmt.Print("Enter your password: ")
	fmt.Scanf("%s", &password)

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
	}
	return config, host
}

//Dial and establish connection to the server
func connect2(config *ssh.ClientConfig, host string) (*ssh.Client) {
	client, err := ssh.Dial("tcp", host + ":22", config)
	if err != nil {
		panic("Fail at dial: " + err.Error())
	}
	fmt.Println("... Connection on " + host + " successfully established ...")
	return client
}

//Execute ssh commands until "exit" is entered
func runCommands(client *ssh.Client) {
	var cmd string
	for strings.ToLower(cmd) != "exit" {
		//Creates a new session.  Only one command per session
		session, err := client.NewSession()
		if err != nil {
			panic("Failed to create session: " + err.Error())
		}

		defer session.Close()

		fmt.Scanf("%s", &cmd)

		var b bytes.Buffer
		session.Stdout = &b
		err1 := session.Run(cmd);
		if err1 != nil {
			fmt.Print("You used an invalid command.")
			err1 = nil
		}
		fmt.Println(b.String())
	}
	//clear the terminal and display conn closed
	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	clear.Run()
	fmt.Println("\n\nConnection Closed")
}