package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	prompt "github.com/c-bata/go-prompt"
	"github.com/kaleocheng/sshconfig"
)

func main() {
	c, err := sshconfig.ParseSSHConfig("")
	if err != nil {
		fmt.Println(err)
	}

	var hosts []string
	for _, item := range c {
		hosts = append(hosts, item.Host...)
	}

	h := strings.TrimSpace(prompt.Choose("host> ", hosts))
	cmd := exec.Command("ssh", h)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()

}
