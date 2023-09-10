## Playing around with SQLC and Go

Explore and experiment with SQLC in a Go environment, along with a Docker-managed PostgreSQL database.

## Install and Run Project:

Setup the Database with Docker Compose:

If you don't have Docker and docker-compose installed, follow the instructions on the official Docker documentation to install them.

This command will pull the necessary images (if not already present), create containers, and start the database. Check that the database container is running with:

```bash
docker-compose up -d
```

This command will pull the necessary images (if not already present), create containers, and start the database. Check that the database container is running with:

Run the Application:

```bash
go run main.go
```

This will start the server on port 4747.

## Test the Application:

```bash
 curl -X GET http://localhost:4747/authors
```

```bash
 curl -X POST http://localhost:4747/authors \
-H "Content-Type: application/json" \
-d '{"Name":"John Doe","Bio":{"String":"Some bio about John.", "Valid":true}}'
```

```bash
curl -X GET http://localhost:4747/authors/1
```
