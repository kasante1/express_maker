package install_packages

import (
	"fmt"
	"os/exec"
)

// is npm available in the dev env
func CheckNpm(project_dir string) error {

	npm := "npm"

	version := "-v"

	cmd := exec.Command(npm, version)

	cmd.Dir = project_dir

	err := cmd.Run()

	errMsg := "visit -> https://docs.npmjs.com/downloading-and-installing-node-js-and-npm"

	if err != nil {
		fmt.Printf(
			"[ X ] %v | %s ", err, errMsg)
		return err
	}

	return nil
}

// install npm packages

func InstallPackages(project_dir string) {

	npm := "npm"
	install := "install"

	cmd := exec.Command(npm, install)

	cmd.Dir = project_dir

	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(err)

	sanitize_packages := `

[ ! ] run npm audit
[ ! ] npm update --save
`

	fmt.Println(sanitize_packages)
}
