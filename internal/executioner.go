/*
 id: executioner.go
 This file is part of Florentino (C) 2020 0xsha
 me[at]0xsha.io
 @0xsha
*/

package internal

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"
)


func SafeExec(command string,args ...string) string {

	cmd := exec.Command(command , args... )

	var buf bytes.Buffer
	cmd.Stdout = &buf

	cmd.Start()


	//https://github.com/golang/go/issues/9580
	done := make(chan error)
	go func() { done <- cmd.Wait() }()

	timeout := time.After(60 * time.Second)

	select {
	case <-timeout:
		cmd.Process.Kill()
		fmt.Println("[-] "+command+":Command timed out")
	case err := <-done:

		if err != nil {
			fmt.Println("[+] Non-zero exit code:", err)
		}
		return buf.String()
	}
return "Timeout"
}


