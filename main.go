//-------------------------------------------------------------------------------
//-------------------------------------------------------------------------------
//  
// Tyler(UnclassedPenguin) bales 2022
//  
// Author: Tyler(UnclassedPenguin)
//    URL: https://unclassed.ca
// GitHub: https://github.com/UnclassedPenguin
//   Desc: A program to keep track of how many bales have been fed to animals.
//
//-------------------------------------------------------------------------------
//-------------------------------------------------------------------------------


package main

import (
  "fmt"
  "os"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
  "github.com/jedib0t/go-pretty/v6/table"
  "log"
  "time"
  "flag"
  "path/filepath"
)

// Create database file if doesn't exist
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

// Creates table in database if doesn't exist
func createTable(db *sql.DB) {
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
func addRecord(db *sql.DB, Date string, AnimalGroup string, TypeOfBale string, NumOfBales int) {
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
func deleteRecord(db *sql.DB, id int) {
  records := "DELETE FROM bales where id = ?"
  query, err := db.Prepare(records)
  if err != nil {
    log.Fatal(err)
  }
  _, err = query.Exec(id)
  if err != nil {
    log.Fatal(err)
  }
}

// This is old way of doing fetchRecords, without go-pretty tables. Will
// Probably just end up removing it but leaving it for now...
// Fetches all records from database and prints to screen
//func fetchRecords(db *sql.DB) {
    //record, err := db.Query("SELECT * FROM bales")
    //if err != nil {
        //log.Fatal(err)
    //}
    //defer record.Close()

    //totalSlice := []int{}
    //var total int

    //fmt.Printf("Bales: ID | Date | Group | TypeOfBale | NumOfBales\n")
    //fmt.Println("-----------------------------------------------")
    //for record.Next() {
        //var id int
        //var Date string
        //var AnimalGroup string
        //var TypeOfBale string
        //var NumOfBales int
        //record.Scan(&id, &Date, &AnimalGroup, &TypeOfBale, &NumOfBales)
        //totalSlice = append(totalSlice, NumOfBales)
        //fmt.Printf("Bales: %d | %s | %s | %s | %d\n", id, Date, AnimalGroup, TypeOfBale, NumOfBales)
    //}

    //// adds up the slice to tell you the total number of bales
    //for _, num := range totalSlice {
      //total += num
    //}

    //fmt.Println("-----------------------------------------------")
    //fmt.Println("Bales:                           Total: ", total)
    //fmt.Println("-----------------------------------------------")
//}

// Fetches all records from database and prints to screen
func fetchRecords(db *sql.DB) {
  record, err := db.Query("SELECT * FROM bales")
  if err != nil {
    log.Fatal(err)
  }
  defer record.Close()

  totalSlice := []int{}
  var total int

  t := table.NewWriter()
  t.SetOutputMirror(os.Stdout)

  t.AppendHeader(table.Row{"id", "Date", "Group", "TypeOfBale", "NumOfBale"})
  for record.Next() {
    var id int
    var Date string
    var AnimalGroup string
    var TypeOfBale string
    var NumOfBales int
    record.Scan(&id, &Date, &AnimalGroup, &TypeOfBale, &NumOfBales)
    totalSlice = append(totalSlice, NumOfBales)
    t.AppendRows([]table.Row{{id, Date, AnimalGroup, TypeOfBale, NumOfBales}})
  }

  // adds up the slice to tell you the total number of bales
  for _, num := range totalSlice {
    total += num
  }

  t.AppendFooter(table.Row{"", "", "", "Total:", total})
  t.Render()
}

// This is old way of doing fetchGroup, without go-pretty tables. Will
// Probably just end up removing it but leaving it for now...
// Fetches all records for a specific group. Requires -l and -g [groupname]
//func fetchGroup(db *sql.DB, AnimalGroup string) {
    //record, err := db.Query("SELECT * FROM bales WHERE AnimalGroup = ?", AnimalGroup)
    //if err != nil {
      //log.Fatal(err)
    //}

    //defer record.Close()

    //totalSlice := []int{}
    //var total int

    //fmt.Printf("Bales: ID | Date | Group | TypeOfBale | NumOfBales\n")
    //fmt.Println("-----------------------------------------------")
    //for record.Next() {
        //var id int
        //var Date string
        //var AnimalGroup string
        //var TypeOfBale string
        //var NumOfBales int
        //record.Scan(&id, &Date, &AnimalGroup, &TypeOfBale, &NumOfBales)
        //totalSlice = append(totalSlice, NumOfBales)
        //fmt.Printf("Bales: %d | %s | %s | %s | %d\n", id, Date, AnimalGroup, TypeOfBale, NumOfBales)
    //}

    //// adds up the slice to tell you the total number of bales
    //for _, num := range totalSlice {
      //total += num
    //}

    //fmt.Println("-----------------------------------------------")
    //fmt.Println("Bales:                           Total: ", total)
    //fmt.Println("-----------------------------------------------")
//}

// Fetches all records for a specific group. Requires -l and -g [groupname]
func fetchGroup(db *sql.DB, AnimalGroup string) {
  record, err := db.Query("SELECT * FROM bales WHERE AnimalGroup = ?", AnimalGroup)
  if err != nil {
    log.Fatal(err)
  }

  defer record.Close()

  totalSlice := []int{}
  var total int
  t := table.NewWriter()
  t.SetOutputMirror(os.Stdout)

  t.AppendHeader(table.Row{"id", "Date", "Group", "TypeOfBale", "NumOfBale"})
  for record.Next() {
    var id int
    var Date string
    var AnimalGroup string
    var TypeOfBale string
    var NumOfBales int
    record.Scan(&id, &Date, &AnimalGroup, &TypeOfBale, &NumOfBales)
    totalSlice = append(totalSlice, NumOfBales)
    t.AppendRows([]table.Row{{id, Date, AnimalGroup, TypeOfBale, NumOfBales}})
  }

  // adds up the slice to tell you the total number of bales
  for _, num := range totalSlice {
    total += num
  }

  t.AppendFooter(table.Row{"", "", "", "Total:", total})
  t.Render()
}

// s for give me some (s)pace
func s() {
  fmt.Print("\n")
}

// Exits. Obvious,  yeah?
func exit(db *sql.DB, status int) {
  db.Close()
  s()
  fmt.Println("Thanks, Bye!")
  os.Exit(status)
  s()
}

// for flag -i. Should add some more useful (i)nfo here,
// but this is helpful for now.
func printInfo() {
  fmt.Println("UnclassedPenguin Bale Tracker")
  fmt.Println("")
  fmt.Println("Groups: sheep, bgoats, lgoats, horse, bulls, cows")
  fmt.Println("Types of bales: square, round")
  os.Exit(0)
}

// for flag -v. Print version info
func printVersion() {
  fmt.Println("UnclassedPenguin Bale Tracker")
  fmt.Println("v0.1.1")
  os.Exit(0)
}


// Global variable for databases. One for real, and one to test 
// things with, that has garbage data in it.
const fileName = "database.db"
const testDb = "test-database.db"


// Main Function
func main() {

  var databaseToUse string

  // Flags
  var info bool
  var list bool
  var test bool
  var add bool
  var del bool
  var push bool
  var pull bool
  var square bool
  var version bool
  var group string
  var number int

  flag.BoolVar(&info, "i", false, "Prints some information you might need to remember.")
  flag.BoolVar(&list, "l", false, "Prints the Database to terminal. Can add -g [group] to only list the records for a specific group.")
  flag.BoolVar(&test, "t", false, "If set, uses the test database.")
  flag.BoolVar(&add, "a", false, "Adds a record to the database. If set, requires -g (group) and -n (number of bales).")
  flag.BoolVar(&del, "d", false, "Deletes a record from the database. If set, requires -n (id number of entry to delete).")
  flag.BoolVar(&square, "s", false, "Wether it is a square bale or round bale. If set, indicates that the bale is square, else it is round.")
  flag.BoolVar(&push, "push", false, "Pushes the databases with git")
  flag.BoolVar(&pull, "pull", false, "Pulls the databases with git")
  flag.BoolVar(&version, "v", false, "Print the version number and exit.")
  flag.StringVar(&group, "g", "", "The name of the group to add to database.")
  flag.IntVar(&number, "n", 0, "The number of bales to add/ or the id of the record to delete .")

  // This changes the help/usage info when -h is used.
  flag.Usage = func() {
      w := flag.CommandLine.Output() // may be os.Stderr - but not necessarily
      fmt.Fprintf(w, "Description of %s:\n\nThis is a program to use to keep track of bales that have been fed.\nIts useful to have the data to see how many bales you go through for the winter.\n\nUsage:\n\nbales [-t] [-l [-g group]] [-a -g group [-s] -n num] [-d -n num]\n\nAvailable arguments:\n", os.Args[0])
      flag.PrintDefaults()
      //fmt.Fprintf(w, "...custom postamble ... \n")
  }

  flag.Parse()

  // Handles cmd line flag -i 
  if info {
    printInfo()
  }

  // Handles cmd line flag -v 
  if version {
    printVersion()
  }

  // Get Current Date 
  t := time.Now()
  timeStr := t.Format("2006-01-02")

  // Change dir to project directory
  // This is needed so a database isn't created where you execute from 
  // (I have the executable soft linked to to a command in ~/.bin)
  // Keeps the database in the project directory
  home, _ := os.UserHomeDir()
  err := os.Chdir(filepath.Join(home, "git/bales/"))
  if err != nil {
      panic(err)
  }

  // Says whether to use the test database or the real database. 
  // Set with -t 
  if test {
    databaseToUse = testDb
  } else {
    databaseToUse = fileName
  }

  // Creates database if it hasn't been created yet.
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


  var typeOfBale string

  // Handles the command line way to add record
  if add && group != "" && number != 0 {
    if square {
      typeOfBale = "square"
    } else {
      typeOfBale = "round"
    }
    fmt.Println("        Date: ", timeStr)
    fmt.Println("       Group: ", group)
    fmt.Println("Type of Bale: ", typeOfBale)
    fmt.Println("Num of Bales: ", number)
    s()
    fmt.Println("Adding record...")
    addRecord(db, timeStr, group, typeOfBale, number)
    fmt.Println("Record added!")
    exit(db, 0)
  } else if add {
    fmt.Println("Requires -g and -n! Try again, or try -h for help.")
    exit(db, 1)
  }

  // Handles the command line way to delete record
  if del && number != 0 {
    fmt.Print("Deleting record ", number , "...\n")
    deleteRecord(db, number)
    fmt.Println("Record deleted!")
    exit(db, 0)
  } else if del {
    fmt.Println("Requires -n (ID number of record to delete)! Try again, or try -h for help.")
    exit(db, 1)
  }

  // Handles the command line way to list records
  if list {
    if group != "" {
      fmt.Println("Date: ", timeStr)
      fetchGroup(db, group)
      exit(db, 0)
    } else {
      fmt.Println("Date: ", timeStr)
      fetchRecords(db)
      exit(db, 0)
    }
  }

  // Handles the github push command.
  if push {
    fmt.Println("This will eventually push to git repo.")
    exit(db, 0)
  }

  // Handles the github pull command.
  if pull {
    fmt.Println("This will eventually handle pull from git repo.")
    exit(db, 0)
  }

  // User interaction starts here
  var userChoice int

  // Main loop starts here if no command line options are used (except -t.)
  for true {
    fmt.Println("Date: ", timeStr)
    s()
    fmt.Println("What would you like to do? (1, 2, 3, 4)")
    fmt.Println("1) Add Record")
    fmt.Println("2) Delete Record")
    fmt.Println("3) Print Records")
    fmt.Println("4) exit")
    fmt.Print(" > ")
    fmt.Scan(&userChoice)

    switch userChoice {
      case 1:
        for true {
          var group string
          var numOfBales int
          var cont string
          var baletype string

          s()
          fmt.Println("What group is this for?(sheep, bgoats, lgoats, horse, bulls, cows)")
          fmt.Print(" > ")
          fmt.Scan(&group)

          s()
          fmt.Println("What type of bale? (s or r)")
          fmt.Println("s) square")
          fmt.Println("r) round")
          fmt.Print(" > ")
          fmt.Scan(&baletype)

          if baletype == "s" {
            typeOfBale = "square"
          } else {
            typeOfBale = "round"
          }

          s()
          fmt.Println("How many bales?")
          fmt.Print(" > ")
          fmt.Scan(&numOfBales)

          s()
          fmt.Println("Adding Record...")
          addRecord(db, timeStr, group, typeOfBale, numOfBales)
          fmt.Println("Record Added!")
          s()

          fmt.Println("Would you like to add another record? (Y or n)")
          fmt.Print(" > ")
          fmt.Scan(&cont)

          if cont == "n" {
            exit(db, 0)
          } else if cont == "" {
            continue
          }

        }

        exit(db, 0)

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
            exit(db, 0)
          }
        }

      case 3:
        s()
        fetchRecords(db)
        exit(db, 0)

      case 4:
        exit(db, 0)

      default:
        s()
        fmt.Println("Please enter a valid option...")
        s()
    }
  }
}
