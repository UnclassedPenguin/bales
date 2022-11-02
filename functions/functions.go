package functions

import (
  "fmt"
  "os"
  "regexp"
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
func CheckDate(date string) {
  // Use regexp to check date to make sure it is a valid yyyy-mm-dd date
  dateCheck, err := regexp.MatchString("^\\d{4}-\\d{2}-\\d{2}$", date)
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
}

// Function to use for debugging or things
func DebugFunction() {
  fmt.Println("Nothing here for now...")
  os.Exit(0)
}

// For flag -v. Print version info
func PrintVersion() {
  fmt.Println("UnclassedPenguin Bale Tracker")
  fmt.Println("v0.3.2")
  os.Exit(0)
}
