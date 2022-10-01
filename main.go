package main

import (
  "fmt"
  "os"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
  "log"
  "time"
)

func createTable(db *sql.DB) {
  balesTable := `CREATE TABLE IF NOT EXISTS bales(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "Date" TEXT,
        "AnimalGroup" TEXT,
        "NumOfBales" INT);`
  query, err := db.Prepare(balesTable)
  if err != nil {
      log.Fatal(err)
  }
  query.Exec()
  fmt.Println("Table created successfully!")
}

func addEntry(db *sql.DB, Date string, AnimalGroup string, NumOfBales int) {
  records := "INSERT INTO bales(Date, AnimalGroup, NumOfBales) VALUES (?, ?, ?)"
  query, err := db.Prepare(records)
  if err != nil {
    log.Fatal("prepare: ",  err)
  }
  _, err = query.Exec(Date, AnimalGroup, NumOfBales)
  if err != nil {
    log.Fatal("exec:", err)
  }
}

func fetchRecords(db *sql.DB) {
    record, err := db.Query("SELECT * FROM bales")
    if err != nil {
        log.Fatal(err)
    }
    defer record.Close()
    for record.Next() {
        var id int
        var Date string
        var AnimalGroup string
        var NumOfBales int
        record.Scan(&id, &Date, &AnimalGroup, &NumOfBales)
        fmt.Printf("Bales: %d %s %s %d\n", id, Date, AnimalGroup, NumOfBales)
    }
}

const fileName = "database.db"

func main() {
  // Get Current Data 
  t := time.Now()
  timeStr := t.Format("2006-01-02")

  // Create database file if it doesn't exist
  _, err := os.Stat("database.db")
  if os.IsNotExist(err) {
    fmt.Println("Database doesn't exist. Creating...")
    file, err := os.Create("database.db")
    if err != nil {
        log.Fatal(err)
    }
    file.Close()
  }

  // Initialize database
  db, err := sql.Open("sqlite3", fileName)
    if err != nil {
        log.Fatal(err)
    }

  createTable(db)
  addEntry(db, timeStr, "Goats", 2)
  fetchRecords(db)
}
