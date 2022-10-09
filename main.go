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
  "os/exec"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
  "github.com/jedib0t/go-pretty/v6/table"
  "log"
  "time"
  "flag"
  "path/filepath"
  "strings"
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

// Fetches all records from database and prints to screen
// cmd line option -l (no other arguments)
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
  t.SetStyle(table.StyleLight)
  // This separates rows...Not sure I like it, leave it for now.
  //t.Style().Options.SeparateRows = true
  t.Render()
}

// Fetches all records for a specific group. 
// Requires -l and -g [groupname]
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
  t.SetStyle(table.StyleLight)
  t.Render()
}

// Fetches all records for a specific bale type. 
// Requires -l and -r or -s
func fetchBaleType(db *sql.DB,  TypeOfBale string) {
  record, err := db.Query("SELECT * FROM bales WHERE TypeOfBale = ?", TypeOfBale)
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
  t.SetStyle(table.StyleLight)
  t.Render()
}

// Fetches all records for a specific year. 
// Requires -l and -y [year]
func fetchRecordYear(db *sql.DB, year string) {
  record, err := db.Query("SELECT * FROM bales WHERE strftime('%Y', date) = ?", year)
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
  t.SetStyle(table.StyleLight)
  t.Render()
}

// Fetches all records for a specific year and month (either single 10 or range 10-12).
// Requires -l and -y [year] -m [month]
func fetchRecordMonth(db *sql.DB, year string, month string) {
  var record *sql.Rows
  var err error
  contains := strings.Contains(month, "-")

  if contains {
    months := strings.Split(month, "-")
    month1 := months[0]
    month2 := months[1]

    record, err = db.Query("SELECT * FROM bales WHERE strftime('%Y', date) = ? and (strftime('%m', date) between ? and ?)", year, month1, month2)
    if err != nil {
      log.Fatal(err)
    }
  } else {
     record, err = db.Query("SELECT * FROM bales WHERE strftime('%Y', date) = ? and strftime('%m', date) = ?", year, month)
    if err != nil {
      log.Fatal(err)
    }
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
  t.SetStyle(table.StyleLight)
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
  fmt.Printf("bales: exit (%d)\n", status)
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

  programName := os.Args[0]
  var databaseToUse string

  // Flags
  var info bool
  var list bool
  var test bool
  var add bool
  var del bool
  var push bool
  var pull bool
  var status bool
  var square bool
  var round bool
  var version bool
  var number int
  var group string
  var year string
  var month string

  flag.BoolVar(&info, "i", false, "Prints some information you might need to remember.")
  flag.BoolVar(&list, "l", false, "Prints the Database to terminal. Can add -g [group] to only list the records for a specific group.")
  flag.BoolVar(&test, "t", false, "If set, uses the test database.")
  flag.BoolVar(&add, "a", false, "Adds a record to the database. If set, requires -g (group) and -n (number of bales).")
  flag.BoolVar(&del, "d", false, "Deletes a record from the database. If set, requires -n (id number of entry to delete).")
  flag.BoolVar(&square, "s", false, "If set, indicates that the bale is square. Round is the default. This can be used when adding (-a) a record, or when listing (-l) to specify that you only want to see square bales.")
  flag.BoolVar(&round, "r", false, "If set, indicates that the bale is round. Round is the default. This can be used when adding (-a) a record, or when listing (-l) to specify that you only want to see round bales.")
  flag.BoolVar(&push, "push", false, "Pushes the databases with git.")
  flag.BoolVar(&pull, "pull", false, "Pulls the databases with git.")
  flag.BoolVar(&status, "status", false, "Checks the git status on project.")
  flag.BoolVar(&version, "v", false, "Print the version number and exit.")
  flag.StringVar(&group, "g", "", "The name of the group to add to database.")
  flag.StringVar(&year, "y", "", "Year to list from database.")
  flag.StringVar(&month, "m", "", "Month to list from database. Can be a single month(10) or a range (10-12). Requires year (-y).")
  flag.IntVar(&number, "n", 0, "The number of bales to add/ or the id of the record to delete .")

  // This changes the help/usage info when -h is used.
  flag.Usage = func() {
      w := flag.CommandLine.Output() // may be os.Stderr - but not necessarily
      fmt.Fprintf(w, "Description of %s:\n\nThis is a program to use to keep track of bales that have been fed.\nIts useful to have the data to see how many bales you go through for the winter.\n\nUsage:\n\nbales [-t] [-l [-g group | -s | -r]] [-a -g group [-s | -r] -n num] [-d -n num]\n\nAvailable arguments:\n", os.Args[0])
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

  // I use this directory in the git section near the end
  directory := filepath.Join(home, "git/bales")

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
  // addRecord(db, timeStr, "Goats", "round", 2)
  // How to delete entry:
  // deleteRecord(db, 2) // where 2 is id number of entry
  // How to query entire database
  // fetchRecords(db)

  var typeOfBale string

  // Handles the command line way to add record
  if add && group != "" && number != 0 {
    if square && !round{
      typeOfBale = "square"
    } else if round && !square {
      typeOfBale = "round"
    } else if square && round {
      fmt.Println("You can't use -s and -r together! How can a bale be a round and square?")
      exit(db, 1)
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
    if group != "" && !round && !square {
      fmt.Println("Date: ", timeStr)
      fetchGroup(db, group)
      exit(db, 0)
    } else if round && !square && group == "" {
      fmt.Println("Date: ", timeStr)
      fetchBaleType(db, "round")
      exit(db, 0)
    } else if square && !round && group == "" {
      fmt.Println("Date: ", timeStr)
      fetchBaleType(db, "square")
      exit(db, 0)
    } else if !square && !round && group == "" && year == "" && month == "" {
      fmt.Println("Date: ", timeStr)
      fetchRecords(db)
      exit(db, 0)
    } else if year != "" && month == "" {
      fmt.Println("Date: ", timeStr)
      fetchRecordYear(db, year)
      exit(db, 0)
    } else if year != "" && month != "" {
      fmt.Println("Date: ", timeStr)
      fetchRecordMonth(db, year, month)
      exit(db, 0)
    } else {
      fmt.Println("You may have used the options wrong. If using -l you can specify group, or baletype. So only one option of -g animal, -s, or -r")
      exit(db, 1)
    }
  }

  // Handles the github push command.
  if push {
    // git add --all
    cmd, stdout := exec.Command("git", "add", "--all"), new(strings.Builder)
    cmd.Dir = directory
    cmd.Stdout = stdout
    err := cmd.Run()
    if err != nil {
      fmt.Println("ERR:", err)
      exit(db, 1)
    }
    fmt.Println(stdout.String())

    // git commit -m 'update database'
    cmd, stdout = exec.Command("git", "commit", "-m", "'update database'"), new(strings.Builder)
    cmd.Dir = directory
    cmd.Stdout = stdout
    err = cmd.Run()
    if err != nil {
      fmt.Println("ERR:", err)
      exit(db, 1)
    }
    fmt.Println(stdout.String())

    // git push
    cmd, stdout, stderr := exec.Command("git", "push"), new(strings.Builder), new(strings.Builder)
    cmd.Dir = directory
    cmd.Stdout = stdout
    cmd.Stderr = stderr
    err = cmd.Run()
    if err != nil {
      fmt.Println("ERR:", err)
      exit(db, 1)
    }
    fmt.Println(stdout.String())
    fmt.Println(stderr.String())

    // Unsatisfactory confirmation message
    fmt.Println("You probably pushed it to git...")
    // Exit
    exit(db, 0)
  }

  // Handles the github pull command.
  if pull {
    // git pull 
    cmd, stdout := exec.Command("git", "pull"), new(strings.Builder)
    cmd.Dir = directory
    cmd.Stdout = stdout
    err := cmd.Run()
    if err != nil {
      fmt.Println("ERR:", err)
      exit(db, 1)
    }
    fmt.Println(stdout.String())

    // exit
    exit(db, 0)
  }

  // Handles the github status command.
  if status {
    // git status 
    cmd, stdout := exec.Command("git", "status"), new(strings.Builder)
    cmd.Dir = directory
    cmd.Stdout = stdout
    err := cmd.Run()
    if err != nil {
      fmt.Println("ERR:", err)
      exit(db, 1)
    }
    fmt.Println(stdout.String())

    // exit
    exit(db, 0)
  }

  // This runs if no arguments are specified. 
  fmt.Printf("%s: No arguments specified. Try -h for help.\n", programName)
}
