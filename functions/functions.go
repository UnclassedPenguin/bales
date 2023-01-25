package functions

import (
  "os"
  "fmt"
  "regexp"
  "strconv"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
)

// For flag -i. Should add some more useful (i)nfo here,
// but this is helpful for now.
func PrintInfo() {
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

// A fucntion to check that the date is correct format
func CheckDate(date string) bool {
  // Use regexp to check date to make sure it is a valid yyyy-mm-dd date
  dateCheck, err := regexp.MatchString("^\\d{4}-\\d{2}-\\d{2}$", date)
  if err != nil {
    fmt.Println("Error in dateCheck: ", err)
    os.Exit(1)
  }

  if dateCheck {
    return true
  } else {
    return false
  }
}

// Checks if an entry is a valid year. Hard to limit it. So just checks if it is 4 digits.
// Therefore....9999 would be a valid year. 
func CheckYear(date string) bool {
  yearCheck, err := regexp.MatchString("^\\d{4}$", date)
  if err != nil {
    fmt.Println("Error in CheckYear: \n", err)
  }
  if yearCheck {
    return true
  } else {
    return false
  }
}

// Check if an entry is a valid month. must be 2 digits and between 01-12
func CheckMonth(date string) bool {
  month, err := strconv.Atoi(date)
  if err != nil {
    fmt.Println("Error converting month to int:\n", err)
  }
  monthCheck, err := regexp.MatchString("^\\d{2}$", date)
  if err != nil {
    fmt.Println("Error in CheckMonth: \n", err)
  }
  if monthCheck && month >= 1 && month <= 12 {
    return true
  } else {
    return false
  }
}

// Check if an entry is a valid day. must be 2 digits and between 01-31
func CheckDay(date string) bool {
  day, err := strconv.Atoi(date)
  if err != nil {
    fmt.Println("Error converting day to int:\n", err)
  }
  dayCheck, err := regexp.MatchString("^\\d{2}$", date)
  if err != nil {
    fmt.Println("Error in CheckDay: \n", err)
  }
  if dayCheck && day >= 1 && day <= 31 {
    return true
  } else {
    return false
  }
}

// Function to use for debugging or things
func DebugFunction() {
  fmt.Println("Nothing here for now...")
  os.Exit(0)
}

// For flag -v. Print version info
func PrintVersion() {
  fmt.Println("UnclassedPenguin Bale Tracker")
  fmt.Println("v0.3.7")
  os.Exit(0)
}

// Exits. Obvious,  yeah?
// Closes the database
func Exit(db *sql.DB, status int) {
  db.Close()
  fmt.Printf("bales: exit (%d)\n", status)
  os.Exit(status)
}


