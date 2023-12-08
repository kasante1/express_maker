package cliarguments

import (
	"fmt"
	"os"
	//"log"
	"errors"
	"express_maker/directorycontents"
	"express_maker/directorymaker"
	"express_maker/installpackages"
	"path/filepath"
	"strings"
)

// if cli_argument is a non exiting directory
// get CWD and create project there
func NewProjectDirectry(cli_argument string) {

	// get current directory
	path, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
	}

	// join CWD and project name to get projectDirectory
	projectDirectory := filepath.Join(path, cli_argument)

	// check if the directory does not exits
	file_status := AlreadyExist(projectDirectory)

	if file_status != nil {
		fmt.Println(file_status)
		os.Exit(0)
	}

	// create express js project directories/ folders
	dir_maker.SubdirectoryMaker(projectDirectory, dir_contents.ExpressDirectories)

	// create express js project files
	dir_maker.CreateProjectFiles(projectDirectory, "README.md", dir_contents.ReadMeFile(cli_argument))
	dir_maker.CreateProjectFiles(projectDirectory, "package.json", dir_contents.PackageJson(cli_argument))
	dir_maker.CreateProjectFiles(projectDirectory, ".gitignore", dir_contents.GitIgnore)
	dir_maker.CreateProjectFiles(projectDirectory, ".env", dir_contents.EnvFile)
	dir_maker.CreateProjectFiles(projectDirectory, "database.json", dir_contents.DatabaseJson)
	dir_maker.CreateProjectFiles(projectDirectory, "index.js", dir_contents.IndexFile)

	// npm and  package.json file things
	check_npm_err := install_packages.CheckNpm(projectDirectory)

	if check_npm_err != nil {
		return
	}

	// npm_init_err := install_packages.NpmInit(projectDirectory)

	// if npm_init_err != nil {
	// 	return
	// }

	for _, dependency := range install_packages.Dependencies{
		fmt.Println("\r installing  %s ... ", dependency)
		install_packages.InstallPackages(projectDirectory, dependency, "--save-prod")


	}


	for _, dependency := range install_packages.DevDependencies{
		fmt.Println("\r installing  %s ... ", dependency)
		install_packages.InstallPackages(projectDirectory, dependency, "--save-dev")

	}

	fmt.Println("Done ... ")



	sanitize_packages := `

[ ! ] run npm audit
[ ! ] npm update --save
`

	fmt.Println(sanitize_packages)
	
}

// is the provided cli_argument a string
// or a directory

func IsArgumentDirectory(cli_argument string) bool {

	// split cli_argument and test if its a directory minus
	// the cli_argument is a directory

	if strings.Contains(cli_argument, "/") {

		dirArgument, _ := SplitCliArgument(cli_argument)
		// check if the dirArgument is a directory
		if _, err := os.Stat(dirArgument); errors.Is(err, os.ErrExist) {
			return true
		}
		//check if the cli_argument is an existing directory
	} else if _, err := os.Stat(cli_argument); errors.Is(err, os.ErrExist) {
		return true
	}

	return false
}


// split cli_argument if it contains a slash
func SplitCliArgument(cli_argument string) (string, string) {

	// split cli_argument by the slash delimiter
	splitArgs := strings.Split(cli_argument, "/")

	//get string slice length
	splitArgs_length := len(splitArgs) - 1

	// get last cli_argument string
	argument_remainder := strings.Join(splitArgs[splitArgs_length:], "")

	return splitArgs[splitArgs_length], argument_remainder
}


// check if the cli_argument provided an existing directory
// project name
func AlreadyExist(cli_argument string) error {

	// is the cli_argument provided an existing directory
	if _, err := os.Stat(cli_argument); errors.Is(err, os.ErrNotExist) {

		// create directory/folder
		create_project := dir_maker.CreateProjectDirectory(cli_argument)

		// stop if directory creation fails
		if create_project != nil {
			return create_project
		}

		return nil

	} else {

		error_message := "[X] project [" + cli_argument + "] already exits!"

		return errors.New(error_message)
	}
}
