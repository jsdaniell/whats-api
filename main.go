package main

import (
	"fmt"
	"github.com/jsdaniell/whats_api/api"
	"github.com/jsdaniell/whats_api/api/utils/file_utility"
	"github.com/jsdaniell/whats_api/api/utils/shell_commands"
)

func main() {
	err := file_utility.DownloadFile("whats-cli.tar.gz", "https://github.com/jsdaniell/whats-cli/releases/download/v1.1.4/whats-cli_1.1.4_linux_amd64.tar.gz")
	if err != nil {
		fmt.Println(err)
	}

	err = shell_commands.ExecuteShellCommand("tar", "-zxvf", "whats-cli.tar.gz")
	if err != nil {
		fmt.Println(err)
	}

	api.Run()
}
