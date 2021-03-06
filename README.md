# Algorea Backend

## Running the app

Installing the dependencies:
```
make deps
```

Compile the app:
```
make
```

You can then run the app: (call `./AlgoreaBackend` to print the list of available commands)
```
./AlgoreaBackend <cmd> <opt>
```
For instance, you can launch the web server using `./AlgoreaBackend serve`.

## Database Configuration

Database configuration currently goes in `conf/default.yml` file or using environment variables (higher priority)
The empty dump (schema with data in it) can be loaded using the `./AlgoreaBackend db-restore` followed by `./AlgoreaBackend db-migrate`.

## Testing

Run all tests (unit and bdd):
```
make test
```
Only unit:
```
make test-unit
```
Only bdd (cucumber using `godog`):
```
make test-bdd <anyflag for godog>
```

## Style

A `.editorconfig` file defines the basic editor style configuration to use. Check the "editorconfig" support for your favorite editor if it is not installed by default.

For the Go coding styles, we use the standard linters (many). You can install and run them with:
```
go lint
```

## Software Walkthrough

### Routing a request

* The web app is defined in `app.go` which loads all the middlewares and routes. The routing part consists in mounting the API on `/` and giving a context to it (i.e., the app database connection)
* The API routing (`app/api/api.go`) does the same for mounting all group of services.
* A service group (e.g., `app/api/groups/groups.go`.) mounts all its services and pass the context again.
* Each service has its dedicated file (e.g., `app/api/groups/get-all.go`). We try to separate the actual HTTP request parsing and response generation from the actual business logic and the call to the database.
