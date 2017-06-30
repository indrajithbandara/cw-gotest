package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq" // postgres driver
)

// People - database
type People struct {
    id   int
    name string
    age  int
}

type appContext struct {
    db *sql.DB
}

// ConnectDB connect specify database
func connectDB(driverName string, dbName string) (c *appContext, errorMessage string) {
    db, err := sql.Open(driverName, dbName)
    if err != nil {
        return nil, err.Error()
    }
    if err = db.Ping(); err != nil {
        return nil, err.Error()
    }
    return &appContext{db}, ""
}

// Create
func (c *appContext) Create() {
    // get insert id
    lastInsertId := 0
    err := c.db.QueryRow("INSERT INTO users(name,age) VALUES($1,$2) RETURNING id", "jack", 22).Scan(&lastInsertId)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("inserted id is ", lastInsertId)
}

// Read
func (c *appContext) Read() {
    rows, err := c.db.Query("SELECT * FROM users")

    if err != nil {
        fmt.Println(err.Error())
        return
    }
    defer rows.Close()

    for rows.Next() {
        p := new(People)
        err := rows.Scan(&p.id, &p.name, &p.age)
        if err != nil {
            fmt.Println(err)
        }
        fmt.Println(p.id, p.name, p.age)
    }
}

// UPDATE
func (c *appContext) Update() {
    stmt, err := c.db.Prepare("UPDATE users SET age = $1 WHERE id = $2")
    if err != nil {
        log.Fatal(err)
    }
    result, err := stmt.Exec(10, 1)
    if err != nil {
        log.Fatal(err)
    }
    affectNum, err := result.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("update affect rows is ", affectNum)
}

// DELETE
func (c *appContext) Delete() {
    stmt, err := c.db.Prepare("DELETE FROM users WHERE id = $1")
    if err != nil {
        log.Fatal(err)
    }
    result, err := stmt.Exec(1)
    if err != nil {
        log.Fatal(err)
    }
    affectNum, err := result.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("delete affect rows is ", affectNum)
}

// Mypg - Postgresql CRUD
func main() {
    c, err := connectDB("postgres", "user=user1 password=password1 dbname=exampledb")
    defer c.db.Close()

    if err != "" {
        print(err)
    }

    c.Create()
    fmt.Println("add action done!")

    c.Read()
    fmt.Println("get action done!")

    c.Update()
    fmt.Println("update action done!")

    c.Delete()
    fmt.Println("delete action done!")
}