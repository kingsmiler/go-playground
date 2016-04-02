package main

import (
	"log"
	"fmt"
	"github.com/pkg/sftp"
	"os"
	"os/exec"
)

func main() {

}

func testCopyFile()  {
	client, cmd:= getSftpClient()
	defer cmd.Wait()

	fmt.Println(client)
}

func listDir() {
	client, cmd:= getSftpClient()
	defer cmd.Wait()

	// read a directory
	list, err := client.ReadDir("/root/gits")
	if err != nil {
		log.Fatal(err)
	}

	// print contents
	for _, item := range list {
		fmt.Println(item.Name())
	}

	// close the connection
	client.Close()
}

func getSftpClient() (*sftp.Client, *exec.Cmd){
	// Connect to a remote host and request the sftp subsystem via the 'ssh'
	// command.  This assumes that passwordless login is correctly configured.
	cmd := exec.Command("ssh", "root@10.211.55.88", "-s", "sftp")

	// send errors from ssh to stderr
	cmd.Stderr = os.Stderr

	// get stdin and stdout
	wr, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	rd, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	// start the process
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	// open the SFTP session
	client, err := sftp.NewClientPipe(rd, wr)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return client, cmd
}
