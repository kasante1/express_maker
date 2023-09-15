 package dir_maker


 import (
"fmt"
"log"
"os"
)




var ExpressDirectories = []string{
"src/config/database.config.js",
"src/env/development.js",
"src/env/index.js",
"src/env/production.js",
"src/env/test.js",
"src/env/versioning/v1.js",
"src/env/versioning/v2.js",
"src/controllers/",
"src/middlewares/auth.middleware.js",
"src/middlewares/error.middleware.js",
"src/middlewares/validation.middleware.js",
"src/queries/",
"src/routes/",
"src/services/",
"/db/",
"/helper/",
}




func dir_maker(path string){

	err := os.MkdirAll(path, 0755)

	if err != nil {
	log.Fatal(err)
	}

	fmt.Println("directory created")

}

for i range := ExpressDirectories{
	_ := dir_maker(i)
}