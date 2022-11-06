package database

import (
  "fmt"
  "os"
  "log"
  "github.com/jedib0t/go-pretty/v6/table"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
)

// Create database file if doesn't exist
func CreateDatabase(db string) {
  _, err := os.Stat(db)
  if os.IsNotExist(err) {
    fmt.Println("Database doesn't exist. Creating...")
    file, err := os.Create(db)
    if err != nil {
        log.Fatal(err)
    }
    file.Close()
  }
}

// Creates table in database if doesn't exist
func CreateTable(db *sql.DB) {
  balesTable := `CREATE TABLE IF NOT EXISTS bales(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "Date" TEXT,
        "AnimalGroup" TEXT,
        "TypeOfBale" TEXT,
        "NumOfBales" INT);`
  query, err := db.Prepare(balesTable)
  if err != nil {
      log.Fatal(err)
  }
  query.Exec()
}

// Adds a record to database
func AddRecord(db *sql.DB, Date string, AnimalGroup string, TypeOfBale string, NumOfBales int) {
  records := "INSERT INTO bales(Date, AnimalGroup, TypeOfBale, NumOfBales) VALUES (?, ?, ?, ?)"
  query, err := db.Prepare(records)
  if err != nil {
    log.Fatal(err)
  }
  _, err = query.Exec(Date, AnimalGroup, TypeOfBale, NumOfBales)
  if err != nil {
    log.Fatal(err)
  }
}

// Deletes a record from database
func DeleteRecord(db *sql.DB, str string) {
  query, err := db.Prepare(str)
  if err != nil {
    log.Fatal(err)
  }
  _, err = query.Exec()
  if err != nil {
    log.Fatal(err)
  }
}

// Fetches a record from database
// Uses github.com/jedib0t/go-pretty/v6/table tables to print it out pretty.
func FetchRecord(db *sql.DB, record *sql.Rows, err error) {
  if err != nil {
    log.Fatal(err)
  }
  defer record.Close()

  //totalSlice := []int{}
  var (
    totalSlice []int
    total int
    id int
    Date string
    AnimalGroup string
    TypeOfBale string
    NumOfBales int
    recordCount int
  )

  t := table.NewWriter()
  t.SetOutputMirror(os.Stdout)

  t.AppendHeader(table.Row{"id", "Date", "Group", "TypeOfBale", "NumOfBale"})

  // Loop over records and add them to the table
  for record.Next() {
    recordCount++
    record.Scan(&id, &Date, &AnimalGroup, &TypeOfBale, &NumOfBales)
    totalSlice = append(totalSlice, NumOfBales)
    t.AppendRows([]table.Row{{id, Date, AnimalGroup, TypeOfBale, NumOfBales}})
  }

  // adds up the slice to tell you the total number of bales
  for _, num := range totalSlice {
    total += num
  }

  t.AppendFooter(table.Row{"", "", "", "Total:", total})
  t.SetStyle(table.StyleLight)

  // To have a line separate all rows, uncomment next line.
  //t.Style().Options.SeparateRows = true

  if recordCount > 0 {
    t.Render()
  } else {
    fmt.Println("\nIt appears there are no entires matching that query. Please try again.\n")
  }
}

