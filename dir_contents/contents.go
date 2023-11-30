package dir_contents


var EnvFile = `
NODE_ENV=development
PORT=
DEV_DATABASE_URL=postgres://<db_owner_name>:<password>PgAdmin@localhost:5432/<db_name>
SECRET=
`

func ReadMeFile(project_name string) string {

	var fileContents = project_name + ` -Project

run

npm install 
npm run start

on the terminal to start express js project

`

	return fileContents
}

var GitIgnore = `
node_modules
.env
.nyc_output
coverage
`

var DatabaseJson = `
{
    "dev": {
        "ENV": "DEV_DATABASE_URL"
    },
    "test": {
        "ENV": "TEST_DATABASE_URL"
    },
    "prod": {
        "ENV": "PROD_DATABASE_URL"
    },
    "sql-file": true
}
`

var IndexFile = `

require('./src/config/database.config')

const express = require('express');
const ap1Version1 = require('./src/config/versioning/v1')
const envConfig = require('./src/config/env/index')
const { notFound, appErrorHandler, genericErrorHandler } = require('./src/middlewares/error.middleware');

const app = express();

app.use(express.json())

const PORT = envConfig.APP_PORT || 7006;

app.listen(PORT, () => {
    console.log("Application running on port", PORT)
})

app.use('/api/v1', ap1Version1);
app.use(appErrorHandler);
app.use(genericErrorHandler);
app.use(notFound)

module.exports = app;


`

func PackageJson(project_name string) string{


var package_json = `
{
	"name":` + `"` + project_name + `"` + `,` +
`
	"version": "1.0.0",
	"description": "",
	"main": "index.js",
	"directories": {
	  "test": "tests"
	},
	"scripts": {
	  "test": "mocha tests/**/*.js",
	  "coverage": "nyc mocha tests/**/*.js",
	  "start": "nodemon index.js",
	  "migrate:create": "db-migrate create --config database.json -e dev",
	  "migrate:up": "db-migrate up --config database.json -e dev",
	  "migrate:down": "db-migrate down -c 2000 --config database.json -e dev",
	  "migrate-test:up": "db-migrate up --config database.json -e test",
	  "migrate-test:down": "db-migrate down -c 200 -e test"
	},
	"repository": {
	  "type": "",
	  "url": ""
	},
	"keywords": [],
	"author": "",
	"license": "ISC",
	"bugs": {
	  "url": ""
	},
	"homepage": "",
  }
  

`

return package_json

}


var Services = ``

var Routes = ` `

var Controllers = ` `

var Middleware = ` `

// index.js
// README.md npm install npm run start
// services
// routes
// controller
// middleware
// queries
// database.json

var ExpressFiles = []string{
	"README.md",
	"package.json",
	"index.js",
	"database.json",
	".env",
	".gitignore",
}

var ExpressDirectories = []string{
	"/src/config/env/",
	"/src/config/versioning/",
	"/src/controllers/",
	"/src/middlewares/",
	"/src/queries/",
	"/src/routes/",
	"/src/services/",
	"/db/",
	"/helper/",
}
