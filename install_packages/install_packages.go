package install_packages


import (
	"fmt"
	"log"
	"os"
	"os/exec"
)




// create the project.json file
func CreatePackageJson(project_dir string) {

	dir := "cd" + project_dir

	create_package_json := "npm"

	packages_args := "init"
	
	auto_all := "--yes"

	cmd := exec.Command(dir,"&&" ,create_package_json, packages_args, auto_all)

	// fmt.Println(cmd)

	output, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("[OK]", string(output))
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
