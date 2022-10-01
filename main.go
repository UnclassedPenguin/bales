package main

import (
  "fmt"
  "os"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
  "log"
  "time"
  "flag"
)


// Create database file
func createDatabase(db string) {
  _, err := os.Stat(db)
  if os.IsNotExist(err) {
    fmt.Println("Database doesn't exist. Creating...")
    file, err := os.Create("database.db")
    if err != nil {
        log.Fatal(err)
    }
    file.Close()
  }
}


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
}


func addRecord(db *sql.DB, Date string, AnimalGroup string, NumOfBales int) {
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


func deleteRecord(db *sql.DB, id int) {
  records := "DELETE FROM bales where id = ?"
  query, err := db.Prepare(records)
  if err != nil {
    log.Fatal("prepare: ",  err)
  }
  _, err = query.Exec(id)
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

    fmt.Printf("Bales: ID | Date | Group | NumOfBales\n")
    for record.Next() {
        var id int
        var Date string
        var AnimalGroup string
        var NumOfBales int
        record.Scan(&id, &Date, &AnimalGroup, &NumOfBales)
        fmt.Printf("Bales: %d %s %s %d\n", id, Date, AnimalGroup, NumOfBales)
    }
}


func s() {
  fmt.Print("\n")
}


func exit(status int) {
  s()
  fmt.Println("Thanks, Bye!")
  os.Exit(status)
  s()
}


func printInfo() {
  fmt.Println("UnclassedPenguin Bale Tracker")
  fmt.Println("")
  fmt.Println("Groups: Sheep, Goats, Horse, Bulls, Cows")
}


const fileName = "database.db"
const testDb = "test-database.db"


func main() {

  var databaseToUse string

  // Flags
  var info bool
  var list bool
  var test bool

  flag.BoolVar(&info, "i", false, "Prints some information you might need to remember.")
  flag.BoolVar(&list, "l", false, "Prints the Database to terminal.")
  flag.BoolVar(&test, "t", false, "If set, uses the test database.")

  flag.Parse()

  if info {
    printInfo()
    os.Exit(0)
  }


  // Get Current Date 
  t := time.Now()
  timeStr := t.Format("2006-01-02")
  fmt.Println("Date: ", timeStr)

  // Create database file if it doesn't exist
  if test {
    databaseToUse = testDb
  } else {
    databaseToUse = fileName
  }

  createDatabase(databaseToUse)


  // Initialize database
  db, err := sql.Open("sqlite3", databaseToUse)
    if err != nil {
        log.Fatal(err)
    }

  // Creates the table initially. "IF NOT EXISTS"
  createTable(db)

  // How to add entry:
  // addRecord(db, timeStr, "Goats", 2)

  // How to delete entry:
  // deleteRecord(db, 2) // where 2 is id number of entry

  // How to query entire database
  // fetchRecords(db)

  if list {
    fetchRecords(db)
    os.Exit(0)
  }

  // User interaction starts here
  var userChoice int
  fmt.Println("What would you like to do? (1, 2, 3)")
  fmt.Println("1) Add Record")
  fmt.Println("2) Delete Record")
  fmt.Println("3) Print Records")
  fmt.Print(" > ")
  fmt.Scan(&userChoice)

  switch userChoice {
    case 1:
      for true {
        var group string
        var numOfBales int
        var cont string

        s()
        fmt.Println("What group is this for?(Sheep, Goats, Horse, Bulls, Cows)")
        fmt.Print(" > ")
        fmt.Scan(&group)

        fmt.Println("How many bales?")
        fmt.Print(" > ")
        fmt.Scan(&numOfBales)

        s()
        fmt.Println("Adding Record...")
        addRecord(db, timeStr, group, numOfBales)
        fmt.Println("Record Added!")
        s()

        fmt.Println("Would you like to add another record? (Y or n)")
        fmt.Print(" > ")
        fmt.Scan(&cont)

        if cont == "n" {
          exit(0)
        } else if cont == "" {
          continue
        }

      }

      exit(0)

    case 2:
      for true {
        var recordToDelete int
        var cont string

        fmt.Println("Which record would you like to Delete?")
        fmt.Print(" > ")
        fmt.Scan(&recordToDelete)

        s()
        fmt.Println("Deleting Record...")
        deleteRecord(db, recordToDelete)
        fmt.Println("Record Deleted!")
        s()

        fmt.Println("Would you like to delete another record? (Y or n)")
        fmt.Print(" > ")
        fmt.Scan(&cont)
        
        if cont == "n" {
          exit(0)
        }
      }

    case 3:
      fetchRecords(db)
      exit(0)
    default:
      fmt.Println("I guess Try again...")
      exit(0)
  }
}
