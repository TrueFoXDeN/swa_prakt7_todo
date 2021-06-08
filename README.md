# SWA -- Damn buggy todo web application

A simple web-based application, that allows you to log in and manage your todos.

## Introduction

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. Enjoy experimenting!

**Your help is required.**

Can you please harden this web app. We have learned that we do not handle security properly. Please check and repair anything you notice.

## Prerequisites

What things you need to have in the first place:

* [Go](https://golang.org/) (Version 1.12 or above)
* [PostgreSQL](https://www.postgresql.org/)
* [PostgreSQL - Go Driver](https://github.com/lib/pq)
* [Gorilla Toolkit - Sessions](https://www.gorillatoolkit.org/pkg/sessions)
* [Gorilla Toolkit - Multiplexer](https://www.gorillatoolkit.org/pkg/mux)

## Installation

Follow these steps:

1. Install a PostgreSQL database server.

2. Setup your database:
   * Add a database user `postgres` and provide the user with a password.
   * Ensure that the `postgres` user is able to create databases and tables.
   * Create an empty database named `todo`.
   * Use the SQL DDL scripts in the `sql-ddl` folder to setup the required tables.

3. The folder in which you unpacked the web app is denoted as "$TODOPATH" in the following (something like "/path/to/swa__prakt2_todo-02" for example).

* Enter the web application's directory

```CLI
cd $TODOPATH
```

4. Get the PostgreSQL driver and Gorilla Toolkit components

```CLI
go get github.com/lib/pq
go get github.com/gorilla/sessions
go get github.com/gorilla/mux
```

5. Build the application from the sources

```CLI
go build
```

6. Start the server (you could also start the executable "swa__prakt2_todo-02" directly)

```CLI
go run todoServer.go
```

7. Access the web application using your browser: [http://localhost:8383/](http://localhost:8383/)

## Built With

* [Golang](https://www.golang.org/) - The Go Programming Language

## Author

* **Luigi Lo Iacono**

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
