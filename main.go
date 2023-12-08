package main

import (
	"express_maker/cliarguments"
	"fmt"
	"os"
)

func main() {
	fmt.Println("express_maker!\n	Usage: express_maker! <project_name>")
	fmt.Println("")

	if len(os.Args) == 1 {
		fmt.Println("[X] enter a name for your express js project.")
		return
	}

	// get cli project argument
	cliArgument := os.Args[1]

	// checked if the cli_argument supplied is a directory
	var fileOrDirectory bool = cliarguments.IsArgumentDirectory(cliArgument)

	// if cli is a directory notify user or
	// create project in the existing directory
	if fileOrDirectory {
		var dir_status error = cliarguments.AlreadyExist(cliArgument)

		if dir_status != nil {
			fmt.Println(dir_status)
		}

	} else {
		// if cliArgument is not a directory
		// create new directory in CWD
		cliarguments.NewProjectDirectry(cliArgument)
	}

}
