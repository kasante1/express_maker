package install_packages

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// create the project.json file
func CreatePackageJson(project_dir string) {

	npm := "npm"

	init := "init"

	// options := "-"

	auto_yes := "-y"

	cmd := exec.Command(npm, init, auto_yes)

	cmd.Dir = project_dir

	err := cmd.Run()
	
	if err != nil {
		fmt.Printf("[ X ] %v create package json", err)
		return 
	}

	fmt.Println("[OK] create package json")
}

// install npm packages

func InstallPackages() {

	packages, err := os.ReadFile("/packages.txt")

	if err != nil {
		panic(err)
	}

	package_contents := string(packages)

	install_command := "npm install " + package_contents

	cmd := exec.Command(install_command)

	output, run_err := cmd.Output()

	if run_err != nil {
		log.Fatal(run_err)
	}

	fmt.Println(output)
}
