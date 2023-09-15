package install_packages


import (
	"fmt"
	"log"
	"os"
	"os/exec"
)




// create the project.json file
func CreatePackageJson(){

	// create_package_json := "npm init --y"

	cmd := exec.Command("npm", "init --y")

	// fmt.Println(cmd)

	output, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[OK]", string(output))
}



// install npm packages

func InstallPackages(){

	packages, err := os.ReadFile("./packages.txt")

	if err != nil{
		panic(err)
	}
	
 	package_contents := string(packages)

	install_command := "npm install" + package_contents

	cmd := exec.Command(install_command)

	output, run_err := cmd.Output()

	if run_err != nil {
		log.Fatal(run_err)
	}

	fmt.Println(output)
}

