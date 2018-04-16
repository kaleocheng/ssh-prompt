package main

import (
	"fmt"
	"os"
	"os/exec"

	prompt "github.com/c-bata/go-prompt"
	"github.com/kaleocheng/sshconfig"
)

var promptSuggest []prompt.Suggest

func completer(d prompt.Document) []prompt.Suggest {
	return prompt.FilterHasPrefix(promptSuggest, d.GetWordBeforeCursor(), true)
}

func main() {
	c, err := sshconfig.ParseSSHConfig("")
	if err != nil {
		fmt.Println(err)
	}

	var hosts []string
	for _, item := range c {
		hosts = append(hosts, item.Host...)
	}

	for _, h := range hosts {
		s := prompt.Suggest{
			Text: h,
		}
		promptSuggest = append(promptSuggest, s)
	}

	h := prompt.Input("host> ", completer)

	cmd := exec.Command("ssh", h)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()

}
