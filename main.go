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
  "strings"
  "regexp"
  "io/ioutil"
  "gopkg.in/yaml.v2"
  "path/filepath"
)


// Create database file if doesn't exist
func createDatabase(db string) {
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

// Fetches a record from database
// Uses go-pretty tables to print it out pretty.
func fetchRecord(db *sql.DB, record *sql.Rows, err error) {
  if err != nil {
    log.Fatal(err)
  }
  defer record.Close()

  totalSlice := []int{}
  var total int

  t := table.NewWriter()
  t.SetOutputMirror(os.Stdout)

  t.AppendHeader(table.Row{"id", "Date", "Group", "TypeOfBale", "NumOfBale"})

  // I don't remember why I declared the variables in the for loop?
  // Is this needed? It would probably be more efficient to declare them
  // outside the loop, if possible. Look into it, and see if they can be moved.
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

// s for give me some (s)pace
func s() {
  fmt.Print("\n")
}

// Exits. Obvious,  yeah?
// Closes the database
func exit(db *sql.DB, status int) {
  db.Close()
  fmt.Printf("bales: exit (%d)\n", status)
  os.Exit(status)
}

// For flag -i. Should add some more useful (i)nfo here,
// but this is helpful for now.
func printInfo() {
  fmt.Println("UnclassedPenguin Bale Tracker")
  fmt.Println("")
  fmt.Println("Groups: sheep, bgoats, lgoats, horses, bulls, cows")
  fmt.Println("Types of bales: square (-s), round (-r)(default)")
  fmt.Println("Sql Table: bales")
  fmt.Println("Sql Columns: id, Date, AnimalGroup, TypeOfBale, NumOfBales")
  fmt.Println("Sql Dates: can use strftime('%Y', date) for year. can use '%m' for month")
  fmt.Println("           and '%d' for day.")
  os.Exit(0)
}

// Function to use for debugging or things
func debugFunction() {
  fmt.Println("Nothing here for now...")
  os.Exit(0)
}

// For flag -v. Print version info
func printVersion() {
  fmt.Println("UnclassedPenguin Bale Tracker")
  fmt.Println("v0.2.1")
  os.Exit(0)
}

// Struct for configuration.
type Configuration struct {
  DatabaseDir  string`yaml:"DatabaseDir"`
  RealDatabase string`yaml:"RealDatabase"`
  TestDatabase string`yaml:"TestDatabase"`
}

// Global variable for databases. One for real, and one to test 
// things with, that has garbage data in it.
//const realDb = "database.db"
//const testDb = "test-database.db"
var (
  realDb string
  testDb string
)


// Main Function
func main() {

  // Flags
  var info    bool
  var list    bool
  var test    bool
  var add     bool
  var del     bool
  var push    bool
  var pull    bool
  var status  bool
  var square  bool
  var round   bool
  var version bool
  var debug   bool
  var number  int
  var group   string
  var year    string
  var month   string
  var date    string
  var custom  string

  flag.BoolVar(&info,     "i",      false, "Prints some information you might need to remember.")
  flag.BoolVar(&list,     "l",      false, "Prints the Database to terminal. Can add -g [group] to only list the records for a specific group.")
  flag.BoolVar(&test,     "t",      false, "If set, uses the test database.")
  flag.BoolVar(&add,      "a",      false, "Adds a record to the database. If set, requires -g (group) and -n (number of bales).")
  flag.BoolVar(&del,      "d",      false, "Deletes a record from the database. If set, requires -n (id number of entry to delete).")
  flag.BoolVar(&square,   "s",      false, "If set, indicates that the bale is square. Round is the default. This can be used when adding (-a) a record, or when listing (-l) to specify that you only want to see square bales.")
  flag.BoolVar(&round,    "r",      false, "If set, indicates that the bale is round. Round is the default. This can be used when adding (-a) a record, or when listing (-l) to specify that you only want to see round bales.")
  flag.BoolVar(&push,     "push",   false, "Pushes the databases with git.")
  flag.BoolVar(&pull,     "pull",   false, "Pulls the databases with git.")
  flag.BoolVar(&status,   "status", false, "Checks the git status on project.")
  flag.BoolVar(&version,  "v",      false, "Print the version number and exit.")
  flag.BoolVar(&debug,    "debug",  false, "Execute function for debugging.")

  flag.StringVar(&group,  "g",      "",    "The name of the group to add to database.")
  flag.StringVar(&year,   "y",      "",    "Year to list from database. Can be a single year(ie 2019) or a range (ie 2019-2022)")
  flag.StringVar(&month,  "m",      "",    "Month to list from database. Can be a single month(ie 09) or a range (ie 09-12). Single digit months require a leading 0.")
  flag.StringVar(&date,   "date",   "",    "The date to put into the database, if not today. yyyy-mm-dd")
  flag.StringVar(&custom, "c",      "",    "Custom SQL request. Requires -l. Example:\nbales -t -l -c \"SELECT * FROM bales WHERE strftime('%d', date) BETWEEN '01' AND '03'\"")

  flag.IntVar(&number,    "n",       0,    "The number of bales to add/ or the id of the record to delete .")

  // This changes the help/usage info when -h is used.
  flag.Usage = func() {
      w := flag.CommandLine.Output() // may be os.Stderr - but not necessarily
      fmt.Fprintf(w, "Description of %s:\n\nThis is a program to use to keep track of bales that have been fed.\nIts useful to have the data to see how many bales you go through for the winter.\n\nUsage:\n\nbales [-t] [-l [-g group] [-s | -r] [-y year] [-m month]] [-a [-date YYYY-MM-DD] -g group [-s | -r] -n num] [-d -n num]\n\nAvailable arguments:\n", os.Args[0])
      flag.PrintDefaults()
      //fmt.Fprintf(w, "...custom postamble ... \n")
  }

  // Parse the flags :p
  flag.Parse()

  // Handles cmd line flag -i 
  // Prints info and exits
  if info {
    printInfo()
  }

  // Handles cmd line flag -v 
  // Prints version and exits
  if version {
    printVersion()
  }

  if debug {
    debugFunction()
  }

  // Variable to hold the date
  var timeStr string

  // Get either Current Date or a date entered as a command line option
  if date == "" {
    t := time.Now()
    timeStr = t.Format("2006-01-02")
  } else {
    timeStr = date
  }

  // Use regexp to check date to make sure it is a valid yyyy-mm-dd date
  dateCheck, err := regexp.MatchString("^\\d{4}-\\d{2}-\\d{2}$", timeStr)
  if err != nil {
    fmt.Println("Error in dateCheck: ", err)
    os.Exit(1)
  }

  // If Regexp check fails, print error and exit,
  // prompting user to use a proper format for date
  if !dateCheck {
    fmt.Println("Error:")
    fmt.Println("It seems your date isn't the proper format. Please enter date as YYYY-MM-DD ie 2022-01-12")
    os.Exit(1)
  }

  // Read Config file and setup databases
  home, _ := os.UserHomeDir()
  configFile, err := ioutil.ReadFile(filepath.Join(home, ".config/bales/config.yaml"))
  if err != nil {
    fmt.Println("Error reading config file:\n", err)
    os.Exit(1)
  }

  var configData Configuration
  err = yaml.Unmarshal(configFile, &configData)
  if err != nil {
    fmt.Println("Error Unmarshal-ling yaml config file:\n", err)
  }

  // I use this directory in the git section near the end
  dbDir := configData.DatabaseDir

  // This sets the database based on the config file
  realDb = configData.RealDatabase
  testDb = configData.TestDatabase

  // Change dir to database directory
  // This is needed so a database isn't created where you execute from 
  // (I have the executable soft linked to to a command in ~/.bin)
  // Keeps the database in the database directory
  err = os.Chdir(dbDir)
  if err != nil {
    fmt.Println("Error changing to directory:\n", err)
  }

  // Var that holds the current working database.
  var databaseToUse string

  // Says whether to use the test database or the real database. 
  // Set with -t 
  if test {
    databaseToUse = testDb
  } else {
    databaseToUse = realDb
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

  // Var to hold the type of bale.
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

    fmt.Println("Date        : ", timeStr)
    fmt.Println("Group       : ", group)
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

  // Handles command line way to list records. 
  // It checks all the flags, and if they have been used, it adds them to "recordStrings". 
  // At the end, it takes all of those strings and creates a database query and then
  // sends that query to the fetchRecord function. 
  if list {
    if custom != "" {
      fmt.Println("Date: ", timeStr)
      record, err := db.Query(custom)
      fetchRecord(db, record, err)
      exit(db, 0)
    }

    // recordStrings collects the sql phrases for each different flag. 
    var recordStrings []string

    // This is the beginning of all queries to the database. I always want every column 
    // returned. So if no options are set, this is sent to fetchRecords all by itself.
    // Otherwise, everything else is added onto this string.
    baseString := "SELECT * FROM bales"

    // Group is -g flag
    if group != "" {
      groupString := fmt.Sprint("Animalgroup='" + group + "'")
      recordStrings = append(recordStrings, groupString)
    }

    if date != "" {
      dateString := fmt.Sprint("date='"+date+"'")
      recordStrings = append(recordStrings, dateString)
    }

    // round is -r flag and square is -s flag
    if round && !square {
      baleString := fmt.Sprint("TypeOfBale='round'")
      recordStrings = append(recordStrings, baleString)
    } else if square && !round {
      baleString := fmt.Sprint("TypeOfBale='square'")
      recordStrings = append(recordStrings, baleString)
    } else if round && square {
      fmt.Println("Can't use -s and -r together. How can a bale be square and round?")
      exit(db, 1)
    }

    // year is -y flag
    if year != "" {
      contains := strings.Contains(year, "-")

      // This handles if you have a range of years. must be written as i.e. 2010-2015
      if contains {
        years := strings.Split(year, "-")
        // Lets the user know that the year must be 4 digits, instead of just returning an empty database.
        if len(years[0]) != 4 || len(years[1]) != 4 {
          fmt.Println("Your year appears to be entered wrong. Make sure year contains exactly 4 digits. ie 2022")
          exit(db, 1)
        }
        yearString := "(strftime('%Y', date) between '" + string(years[0]) + "' and '" + string(years[1]) + "')"
        recordStrings = append(recordStrings, yearString)
      // This handles single year 
      } else {
        // Lets the user know that the year must be 4 digits, instead of just returning an empty database.
        if len(year) != 4 {
          fmt.Println("Your year appears to be entered wrong. Make sure year contains exactly 4 digits. ie 2022")
          exit(db, 1)
        }
        yearString := fmt.Sprint("strftime('%Y', date)='" + year + "'")
        recordStrings = append(recordStrings, yearString)
      }
    }

    // month is -m flag
    if month != "" {
      contains := strings.Contains(month, "-")

      // This handles if you have a range of months. must be written as i.e. 05-10
      if contains {
        months := strings.Split(month, "-")
        // Lets the user know that the month requires a leading 0, instead of just returning an empty database.
        if len(months[0]) != 2 || len(months[1]) != 2 {
          fmt.Println("Your month appears to be wrong. Make sure each month is exactly 2 digits. If its a single digit month, add a leading zero, ie 05.")
          exit(db, 1)
        }
        monthString := "(strftime('%m', date) between '" + string(months[0]) + "' and '" + string(months[1]) + "')"
        recordStrings = append(recordStrings, monthString)
      // This handles single month
      } else {
        // Lets the user know that the month requires a leading 0, instead of just returning an empty database.
        if len(month) != 2 {
          fmt.Println("Your month appears to be wrong. Make sure each month is exactly 2 digits. If its a single digit month, add a leading zero, ie 05.")
          exit(db, 1)
        }
        monthString := fmt.Sprint("strftime('%m', date)='" + month + "'")
        recordStrings = append(recordStrings, monthString)
      }
    }

    // This is the area that puts the sql phrase together and sends it to the fetchRecords
    // function.
    // I set it up to pay attention to three scenarios: No additional phrase, 1 additional
    // phrase, or more than one additional phrase.
    // The phrases are stored in the slice recordStrings
    // If no additional phrases were set, ie no flags were used, sends only the baseString,
    // which returns the entire database.
    if len(recordStrings) == 0 {
      fmt.Println("Date: ", timeStr)
      fmt.Println("SQL Query:", baseString)
      record, err := db.Query(baseString)
      fetchRecord(db, record, err)
      exit(db, 0)
    // If there is one additional phrase, it appends WHERE and the phrase to base string,
    } else if len(recordStrings) == 1 {
      fmt.Println("Date: ", timeStr)
      fullString := fmt.Sprint(baseString + " WHERE " + recordStrings[0])
      fmt.Println("SQL Query:", fullString)
      record, err := db.Query(fullString)
      fetchRecord(db, record, err)
      exit(db, 0)
    // If there are more than one phrase to add, first it combines them with AND, and 
    // then adds that to baseString, with the connecting WHERE as well.
    } else if len(recordStrings) > 1 {
      fmt.Println("Date: ", timeStr)
      combineStrings := strings.Join(recordStrings, " AND ")
      fullString := fmt.Sprint(baseString + " WHERE " + combineStrings)
      fmt.Println("SQL Query:", fullString)
      record, err := db.Query(fullString)
      fetchRecord(db, record, err)
      exit(db, 0)
    }
  }

  // Handles the github push command.
  if push {
    // git add --all
    cmd, stdout := exec.Command("git", "add", "--all"), new(strings.Builder)
    cmd.Dir = dbDir
    cmd.Stdout = stdout
    err := cmd.Run()
    if err != nil {
      fmt.Println("Error executing git add --all:\n", err)
      exit(db, 1)
    }
    fmt.Println(stdout.String())

    // git commit -m 'update bales database'
    cmd, stdout = exec.Command("git", "commit", "-m", "'update bales database'"), new(strings.Builder)
    cmd.Dir = dbDir
    cmd.Stdout = stdout
    err = cmd.Run()
    if err != nil {
      fmt.Println("Error executing git commit -m 'update bales database':\n", err)
      exit(db, 1)
    }
    fmt.Println(stdout.String())

    // git push
    cmd, stdout, stderr := exec.Command("git", "push"), new(strings.Builder), new(strings.Builder)
    cmd.Dir = dbDir
    cmd.Stdout = stdout
    cmd.Stderr = stderr
    err = cmd.Run()
    if err != nil {
      fmt.Println("Error executing git push:\n", err)
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
    cmd.Dir = dbDir
    cmd.Stdout = stdout
    err := cmd.Run()
    if err != nil {
      fmt.Println("Error executing git pull:\n", err)
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
    cmd.Dir = dbDir
    cmd.Stdout = stdout
    err := cmd.Run()
    if err != nil {
      fmt.Println("Error executing git status:\n", err)
      exit(db, 1)
    }
    fmt.Println(stdout.String())

    // exit
    exit(db, 0)
  }

  // This runs if no arguments are specified. Prints help usage.
  flag.Usage()
}
