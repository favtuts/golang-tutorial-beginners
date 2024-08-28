# Structuring your Golang app: Flat structure vs. layered architecture
* https://tuts.heomi.net/structuring-your-golang-app-flat-structure-vs-layered-architecture/

# Building a simple API using a flat structure

Create a new directory for this project by running:
```sh
mkdir notes_api_flat
```

Now, initialize the project:
```sh
go mod init github.com/username/notes_api_flat
```

We’ll use [SQLite3](https://www.sqlite.org/index.html) for storing the notes and [Gin](https://gopkg.in/gin-gonic/gin.v1) for routing.
```sh
go get github.com/mattn/go-sqlite3
go get github.com/gin-gonic/gin
```