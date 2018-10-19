# Algorea Backend

## Running the app

You can run the application commands (implemented in `cmd`) using:
```
go run main.go <cmd> <opt>
```
or by compiling it and then executing the binary:
```
go build
./AlgoreaBackend <cmd> <opt>
```

You can call `./AlgoreaBackend` alone to print the list of available commands.

For instance, you can launch the web server using `./AlgoreaBackend serve`.

## Software Walkthrough

### Routing a request

* The web app is defined in `app.go` which loads all the middlewares and routes. The routing part consists in mounting the API on `/` and giving a context to it (i.e., the app database connection)
* The API routing (`app/service/api/api.go`) does the same for mounting all group of services.
* A service group (e.g., `app/service/api/groups/groups.go`.) mounts all its services and pass the context again.
* Each service has its dedicated file (e.g., `app/service/api/groups/get-all.go`). We try to separate the actual HTTP request parsing and response generation from the actual business logic and the call to the database.
