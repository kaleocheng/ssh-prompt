package main

import (
	"flag"
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"strings"

	prompt "github.com/c-bata/go-prompt"
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
		t := template.New("ssh config template")
		t, _ = t.Parse(`Host {{ index .Host 0 }}{{ if .HostName }}
  Hostname {{ .HostName }}{{ end }}{{ if .Port }}
  Port {{ .Port }}{{ end }}{{ if .User }}
  User {{ .User }}{{ end }}{{ if .ProxyCommand }}
  ProxyCommand {{ .ProxyCommand }}{{ end }}{{ if .HostKeyAlgorithms }}
  HostKeyAlgorithms {{ .HostKeyAlgorithms }}{{ end }}{{ if .IdentityFile }}
  IdentityFile {{ .IdentityFile }}{{ end }}
`)
		t.Execute(os.Stdout, m[h])
	} else {
		cmd := exec.Command("ssh", h)
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
}
