package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	prompt "github.com/c-bata/go-prompt"
	"github.com/k0kubun/pp"
	"github.com/kaleocheng/sshconfig"
)

func main() {
	showPtr := flag.Bool("show", false, "show ssh config")
	flag.Parse()

	c, err := sshconfig.ParseSSHConfig([]string{})
	if err != nil {
		fmt.Println(err)
	}

	m := make(map[string]*sshconfig.SSHHost)

	var hosts []string
	for _, item := range c {
		hosts = append(hosts, item.Host...)
		m[item.Host[0]] = item
	}

	h := strings.TrimSpace(prompt.Choose("host> ", hosts))
	if h == "quit" || h == "exit" {
		os.Exit(0)
		return
	}

	if *showPtr {
		pp.Print(m[h])
	} else {
		cmd := exec.Command("ssh", h)
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
}
