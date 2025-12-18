// agent-cli is a command-line tool for managing agent rules.
// It downloads rule files from a GitHub repository and places them in .agent/rules.
package main

import "github.com/odinnordico/antigravity-agents/cmd"

func main() {
	cmd.Execute()
}
