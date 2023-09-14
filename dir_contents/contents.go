package dir_contents

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)



//html file
// file contents
// get project name from cli
// use it to link the css and js files
// to the html


func HtmlFileContents(project_name string) string {
	var firstSection string = `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
	`
	var midSection string =	`<link rel="stylesheet" href="`+ project_name +`.css" />
		<script src="`+ project_name +`.js" defer></script>

		<title>`+ project_name +`</title>
	`

	var lastSection string = `
	</head>
	<body>
		<nav>
		</nav>

		<footer>
	
		</footer>
	</body>
	</html> 
	
`
	var fileContents string = firstSection + midSection + lastSection

	return fileContents
}


func CssFileContents() string{
	cssFileContents := ` 
	*{
		margin: 0px;
		padding: 0px;
		box-sizing: border-box;
	}
	
	`
	return cssFileContents

}

func JsFileContents() string{
	jsFileContents := ``
	return jsFileContents

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

// create file, folder, tests and test directories