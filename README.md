### Good URL examples
* Login of users:
    * POST {{Ruta}}/login
    
    * Request body:
    {
      "email": "adminsige@gmail.com",
      "password": "11223344"
    }
    * Response body:
    {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcGVsbGlkb3MiOiJTSUdFIiwiZW1haWwiOiJhZG1pbnNpZ2VAZ21haWwuY29tIiwiZXhwIjoxNjEyMDM5NDYzLCJmZWNoYV9uYWNpbWllbnRvIjoiMDAwMC0wMC0wMCIsIm5vbWJyZSI6IkFkbWluIn0.etd-lUbx02zNK1IWLDf0keLMv7co1fxHC0YJIAfsmgY"
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
http://stackoverflow.com -->