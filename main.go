package main

import (
    "database/sql"
    "fmt"

    _ "modernc.org/sqlite"
)

func main() {
    fmt.Println("Opening in-memory sqlite database")
    db, err := sql.Open("sqlite", ":memory:")
    if err != nil {
        fmt.Println(err)
        return
    }

    if _, err = db.Exec(`
DROP TABLE IF EXISTS sqlite;
CREATE TABLE sqlite(i);
    `); err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Done")
}
