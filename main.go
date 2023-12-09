package main

import (
	"fmt"
	"express_maker/cli_args"
	"os"
)

func main() {
	fmt.Println("express_maker!\n	Usage: express_maker! <project_name>")
	fmt.Println("")

	if len(os.Args) == 1 {
		fmt.Println("[X] enter a name for your express js project.")
		return
	}

	// get fetch cli_arguments
	cli_arguments := os.Args

	// get cli project argument
	cli_project_argument := cli_arguments[1]

	// checked if the cli_argument supplied is a directory
	var fileOrDirectory bool = cli_args.IsArgumentDirectory(cli_project_argument)

	// if cli is a directory notify user or
	// create project in the existing directory
	if fileOrDirectory {
		var dir_status error = cli_args.DirectoryExists(cli_project_argument)

		if dir_status != nil {
			fmt.Println(dir_status)
		}

	} else {
		// if cli_project_argument is not a directory
		// create new directory in CWD
		cli_args.NewProjectDirectry(cli_project_argument)
	}

}


