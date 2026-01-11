// package main is the entry point for Lumna.
package main

import (
	"os"

	"gitlab.com/shaninalex/lumna/app2/cmd"
)

func main() {
	os.Exit(cmd.Execute())
}
