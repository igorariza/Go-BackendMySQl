# Go RESTful API

[![GoDoc](https://godoc.org/github.com/qiangxue/go-rest-api?status.png)](http://godoc.org/github.com/qiangxue/go-rest-api)


This starter kit is designed to get you up and running with a project structure optimized for developing
RESTful API services in Go. It promotes the best practices that follow the [SOLID principles](https://en.wikipedia.org/wiki/SOLID)
and [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html). 
It encourages writing clean and idiomatic Go code. 

The kit provides the following features right out of the box:

* RESTful endpoints in the widely accepted format
* Standard CRUD operations of a database table
* JWT-based authentication
* Environment dependent application configuration management
* Structured logging with contextual information
* Error handling with proper error response generation
* Database migration
* Data validation
* Full test coverage
* Live reloading during development
 
The kit uses the following Go packages which can be easily replaced with your own favorite ones
since their usages are mostly localized and abstracted. 

<!-- * Routing: [ozzo-routing](https://github.com/go-ozzo/ozzo-routing)
* Database access: [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx)
* Database migration: [golang-migrate](https://github.com/golang-migrate/migrate)
* Data validation: [ozzo-validation](https://github.com/go-ozzo/ozzo-validation)
* Logging: [zap](https://github.com/uber-go/zap) -->
* JWT: [jwt-go](https://github.com/dgrijalva/jwt-go)

## Getting Started

If this is your first time encountering Go, please follow [the instructions](https://golang.org/doc/install) to
install Go on your computer. The kit requires **Go 1.13 or above**.

After installing Go , run the following commands to start experiencing this starter kit:

```shell
# download the starter kit
git clone https://github.com/qiangxue/go-rest-api.git

cd go-rest-api

# start a PostgreSQL database server in a Docker container
make db-start

# seed the database with some test data
make testdata

# run the RESTful API server
make run

# or run the API server with live reloading, which is useful during development
# requires fswatch (https://github.com/emcrisostomo/fswatch)
make run-live
```

At this time, you have a RESTful API server running at `http://127.0.0.1:8080`. It provides the following endpoints:

* `GET /healthcheck`: a healthcheck service provided for health checking purpose (needed when implementing a server cluster)
* `POST /v1/login`: authenticates a user and generates a JWT
* `GET /v1/albums`: returns a paginated list of the albums
* `GET /v1/albums/:id`: returns the detailed information of an album
* `POST /v1/albums`: creates a new album
* `PUT /v1/albums/:id`: updates an existing album
* `DELETE /v1/albums/:id`: deletes an album

Try the URL `http://localhost:8080/healthcheck` in a browser, and you should see something like `"OK v1.0.0"` displayed.

If you have `cURL` or some API client tools (e.g. [Postman](https://www.getpostman.com/)), you may try the following 
more complex scenarios:

```shell
# authenticate the user via: POST /v1/login
curl -X POST -H "Content-Type: application/json" -d '{"username": "demo", "password": "pass"}' http://localhost:8080/v1/login
# should return a JWT token like: {"token":"...JWT token here..."}

# with the above JWT token, access the album resources, such as: GET /v1/albums
curl -X GET -H "Authorization: Bearer ...JWT token here..." http://localhost:8080/v1/albums
# should return a list of album records in the JSON format
```

To use the starter kit as a starting point of a real project whose package name is `github.com/abc/xyz`, do a global 
replacement of the string `github.com/qiangxue/go-rest-api` in all of project files with the string `github.com/abc/xyz`.


## Project Layout

The starter kit uses the following project layout:
 
```
.
├── cmd                  main applications of the project
│   └── server           the API server application
├── config               configuration files for different environments
├── internal             private application and library code
│   ├── album            album-related features
│   ├── auth             authentication feature
│   ├── config           configuration library
│   ├── entity           entity definitions and domain logic
│   ├── errors           error types and handling
│   ├── healthcheck      healthcheck feature
│   └── test             helpers for testing purpose
├── migrations           database migrations
├── pkg                  public library code
│   ├── accesslog        access log middleware
│   ├── graceful         graceful shutdown of HTTP server
│   ├── log              structured and context-aware logger
│   └── pagination       paginated list
└── testdata             test data scripts
```

### Good URL examples
* Login of users:
    * POST {{Ruta}}/login
    Request body:
    {
      "email": "adminsige@gmail.com",
      "password": "11223344"
    }
    * Response body:
    {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...."
    }
    {
        "success": true,
        "status": 200,
        "result": {
            "idUser": 3,
            "document_id": "11223344",
            "first_name": "Admin",
            "last_name": "SIGE",
            "email": "adminsige@gmail.com",
            "phone": "",
            "address": "",
            "photo": "",
            "created_at": "2021-01-19 16:40:30.750144 -0500 -05 m=+4.78115315",
            "type_id": "0",
            "date_birth": "0000-00-00",
            "last_access": "",
            "rh": "",
            "idSede": "1",
            "is_active": "0",
            "name_sede": "MANUELA BELTRAN"
        }
    }
<!-- * List of users:
    * GET {{Ruta}}/api/users/users.json
* Filtering is a query:
    * GET {{Ruta}}/api/v1/magazines.json?year=2011&sort=desc
    * GET {{Ruta}}/api/v1/magazines.json?topic=economy&year=2011
* A single magazine in JSON format:
    * GET {{Ruta}}/api/v1/magazines/1234.json
* All articles in (or belonging to) this magazine:
    * GET {{Ruta}}/api/v1/magazines/1234/articles.json
* All articles in this magazine in XML format:
    * GET {{Ruta}}/api/v1/magazines/1234/articles.xml
* Specify optional fields in a comma separated list:
    * GET {{Ruta}}/api/v1/magazines/1234.json?fields=title,subtitle,date
* Add a new article to a particular magazine:
    * POST {{Ruta}}/api/v1/magazines/1234/articles -->

<!-- ### Bad URL examples
* Non-plural noun:
    * {{Ruta}}/magazine
    * {{Ruta}}/magazine/1234
    * {{Ruta}}/publisher/magazine/1234
* Verb in URL:
    * {{Ruta}}/magazine/1234/create
* Filter outside of query string
    * {{Ruta}}/magazines/2011/desc

## Request & Response Examples

### API Resources

  - [GET /magazines](#get-magazines)
  - [GET /magazines/[id]](#get-magazinesid)
  - [POST /magazines/[id]/articles](#post-magazinesidarticles)

### GET /magazines

Example: {{Ruta}}/api/v1/magazines.json

Response body:

    {
        "metadata": {
            "resultset": {
                "count": 123,
                "offset": 0,
                "limit": 10
            }
        },
        "results": [
            {
                "id": "1234",
                "type": "magazine",
                "title": "Public Water Systems",
                "tags": [
                    {"id": "125", "name": "Environment"},
                    {"id": "834", "name": "Water Quality"}
                ],
                "created": "1231621302"
            },
            {
                "id": 2351,
                "type": "magazine",
                "title": "Public Schools",
                "tags": [
                    {"id": "125", "name": "Elementary"},
                    {"id": "834", "name": "Charter Schools"}
                ],
                "created": "126251302"
            }
            {
                "id": 2351,
                "type": "magazine",
                "title": "Public Schools",
                "tags": [
                    {"id": "125", "name": "Pre-school"},
                ],
                "created": "126251302"
            }
        ]
    }

### GET /magazines/[id]

Example: {{Ruta}}/api/v1/magazines/[id].json

Response body:

    {
        "id": "1234",
        "type": "magazine",
        "title": "Public Water Systems",
        "tags": [
            {"id": "125", "name": "Environment"},
            {"id": "834", "name": "Water Quality"}
        ],
        "created": "1231621302"
    }



### POST /magazines/[id]/articles

Example: Create – POST  {{Ruta}}/api/v1/magazines/[id]/articles

Request body:

    [
        {
            "title": "Raising Revenue",
            "author_first_name": "Jane",
            "author_last_name": "Smith",
            "author_email": "jane.smith@example.gov",
            "year": "2012",
            "month": "August",
            "day": "18",
            "text": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam eget ante ut augue scelerisque ornare. Aliquam tempus rhoncus quam vel luctus. Sed scelerisque fermentum fringilla. Suspendisse tincidunt nisl a metus feugiat vitae vestibulum enim vulputate. Quisque vehicula dictum elit, vitae cursus libero auctor sed. Vestibulum fermentum elementum nunc. Proin aliquam erat in turpis vehicula sit amet tristique lorem blandit. Nam augue est, bibendum et ultrices non, interdum in est. Quisque gravida orci lobortis... "
        }
    ]
