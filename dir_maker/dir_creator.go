 package dir_maker


 import (
"fmt"
"log"
"os"
)


func dir_paths(path string){

	err := os.MkdirAll(path, 0755)

	if err != nil {
	log.Fatal(err)
	}

	fmt.Println("directory created")

}