# Go Coffee API

This is an API for coffee shops in Guadalajara. See details in the [Go Coffee API documentation](https://go-coffee-api.vercel.app/)

## Instructions for development

### Environmental variables

This project requires the following environmental variables set:

``` bash
JWT_SECRET=<random_value_with_high_entropy>
DB_USER=<db_user>
DB_PASSWORD=<user_db_password>
DB_HOST=<host>:<port>
DB_PATH=<db_name>
MODE=<dev|prod>
GO_COFFEE_API_DSN=postgres://<db_user>:<db_user_password>@<host>/<database_name>
```

### Migrations

You need three things for the development process
- Postgres database
- [Golang Migrate tool](https://github.com/golang-migrate/migrate)
- An environmental variable called GO_COFFEE_API_DSN equal to the following postgres://<db_user>:<db_user_password>@<host>/<database_name>

#### Migrate

Once you got those things in place just run the following command.

``` bash
make migrate/up
```

If you want to do it manually you can run the following command, replacing the appropiate values

``` bash
migrate -path=./migrations -database postgres://<db_user>:<db_user_password>@<host>/<database_name> up
```

#### Undo migrations

If you want to return the database to its initial state run the following command.

``` bash
make migrate/down
```

If you want to do it manually you can run the following command, replacing the appropiate values

``` bash
migrate -path=./migrations -database postgres://<db_user>:<db_user_password>@<host>/<database_name> down
```

### Update docs

You need to have swag installed. The docs are installed in the docs directory, but you don't need to configurate anything, just run the following command.

``` bash
docs/generate
```

### Install npm packages

Install the required packages for vercel to work

``` bash
npm install
```

Install vercel package globally in case you don't have it

``` bash
npm i -g vercel
```

### Run Go and javascript server

Run golang and javascript server using vercel command.

``` bash
vercel dev
```

The previous command will run javascript and go server