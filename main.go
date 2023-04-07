package main

import (
	"github.com/869413421/chatgpt-web/bootstarp"
	"github.com/869413421/chatgpt-web/config"
	"github.com/alecthomas/kong"
	"os/exec"
)

func main() {
	kong.Parse(&config.CLI)
	bootstrap.StartWebServer()
	startFe()
}

func startFe() {
	cmd := exec.Command("cd chat-new && npm run dev")

	cmd.Run()
}
