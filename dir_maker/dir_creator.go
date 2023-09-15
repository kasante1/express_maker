package dir_maker



import (
	"fmt"
	"os"

	// "log"
	"errors"
	"path/filepath"
)




// create project directory/folders
func CreateProjectDirectory(cli_argument string) error {

	err := os.Mkdir(cli_argument, os.ModePerm)

	if err != nil {
		return errors.New("[X] Failed! check directory permission")
	}

	fmt.Println("[OK] project created here :", cli_argument)
	return nil
}






// create project files
func CreateProjectFiles(SubDirectories, fileName, file_contents string) {

	// file directory
	file_path := filepath.Join(SubDirectories, fileName)

	// check if file exist or otherwise
	_, err := os.Stat(file_path)

	// if file does not exits, create file.
	if errors.Is(err, os.ErrNotExist) {

		// write content to files
		WriteProjectFiles(file_path, file_contents)

		fmt.Println("[OK]", fileName, "created succesfully")

	} else {

		fmt.Println("[X]", fileName, "failed. already exits!")
	}
}





// write content into files
func WriteProjectFiles(fileName, file_contents string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	if _, err := file.WriteString(file_contents); err != nil {
		fmt.Println(err)
	}
}




// create test directory

func SubdirectoryMaker(cwd_dir string, dir_paths []string) {

	// get the defined subdirectories eg: src/controllers/user.controller
	for _, path := range dir_paths {

		// join the subdirectories to cwd of the project
		var sub_dir = filepath.Join(cwd_dir, path)

		
		// check if folder exists
		_, err := os.Stat(sub_dir)

		if errors.Is(err, os.ErrNotExist){
	
			err := os.MkdirAll(sub_dir, 0755)

			if err != nil {
				fmt.Println("[ x ] failed to create", sub_dir, "directory")
			}
	
			fmt.Println("[ OK ] -", path, " --directory created")

		}else{
			fmt.Println("[X]", path, " -- failed! directory already exits")
		}

	
	}

}
