package shell_commands

import (
	"fmt"
	execute "github.com/alexellis/go-execute/pkg/v1"
)

func ExecuteSh(file, projectName string) {
	cmd := execute.ExecTask{
		Command:     "sh",
		Args:        []string{file},
		StreamStdio: false,
		Cwd:         projectName,
	}

	res, err := cmd.Execute()
	if err != nil {
		panic(err)
	}

	fmt.Printf("output: %s", res.Stderr)
}

func ExecuteShellCommand(command string, args ...string)  error {
	cmd := execute.ExecTask{
		Command: command,
		Args:    args,
		StreamStdio: true,
	}


	res, err := cmd.Execute()
	if err != nil {
		return err
	}

	fmt.Printf("output: %s", res.Stdout)
	fmt.Printf("output: %s", res.Stderr)

	return nil
}
