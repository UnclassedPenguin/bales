# Bales
#### A program for keeping track of bales that you have fed

# Install

To install:

```shell
$ go install github.com/unclassedpenguin/bales@latest
```

Then, you need to create the config file at `~/.config/bales/config.yaml`

### Example ~/.config/bales/config.yaml

```yaml
# Database dir is the directory you want to store your databases in.
# It can be a git repo, but doesn't have to be...
# This directory must be created by the user
DatabaseDir: /home/username/git/databases

# RealDatabase is the legit database
# This will be created if it doesn't exist
RealDatabase: balesDatabase.db

# TestDatabase is a database you can use to test features
# This will be created if it doesn't exist
TestDatabase: balesTestDatabase.db

```

Don't forget to edit the DatabaseDir to where you want to store the databases. You just have to create the folder, the databases will be created if they don't exist.

For further documentation checkout [docs.unclassed.ca/bales](https://docs.unclassed.ca/bales)


## To-do:
  
  - If you use option -year without specifying which year it should show current year
  - option '-m' which should only list current month lists for the current month but of all years
  - If you try to list with month first then year, it doesn't work. Has to be year first then month. ? ie `bales -l -m 09 -year 2022` will give you all entries that match the month, however to properly get only the month and year you have to use `bales -l -year 2022 -m 09`
  - Add "-most" i.e. "bales -most day" to list the day that you fed the most bales. Do maybe day/month/year?
  - If you try to list for the current month, it gives the default "There are no results for that query"
    Should probably change this so it says no entries for current month. Try -all or something.
  - Change checkdate functions so they return an error instead of bool. will need a bit of reformating. 
  - Add "-to". So you can do bales -l -from 2022-08-01 -to 2023-04-01
  - Add license?
  - Add ability to list based on ID number, singular or range (1 or 1-5)
    - Do I need this? I'm sure I wanted it for some reason...I'm not sure I want to put it in though. 
  - Add "-between" to list dates between one and another
  - Add an ability to get average. For square/round. Maybe by group as well. For daily weekly monthly?
    - Think this will be slightly more complicated than I originally thought...  
  - ~~On basic list (-l) it shows only current month. It however, still shows for all years. Need to make it show only current month of current year.~~

#### v0.3.6

  - ~~Possibly add Notes column?~~
    - I tried this. Would basically need me to restart the database. I don't think its worth it. It really messes with the formatting when listing...Just doesn't seem worth it at this point.
  - ~~Change -dateoton and -datentoo to -asc and -dec~~
  - ~~Add "-today" to list only for the current day~~
  - ~~When you mess up the -day flag, it still says something about month~~

#### v0.3.4

  - ~~Add "checkMonth" and "checkYear" Functions that do a regex and make sure it is 2 digits and 4 digits. Use when for -month and -year to confirm they are valid. Month could also be checked to make sure it is between 1 and 12.~~
    - Added "CheckYear", "CheckMonth", and "CheckDay" to functions.
  - ~~Remove sql string by default. Add flag to show it. I don't think most people would care to see what the query is.~~
  - ~~Make default list only list the current month. add -all flag to show entire database~~
    - Do I actually want this? Hmm...
    - I tried to implement this, I didn't like it. Instead I added a flag (-m) that lists only current month. It conflicts with -month, so you can't use them together. 
  - ~~Add Ability to delete groups. Maybe add a comfirmation "Are you sure you want to delete group cows? (y or n)"~~  

#### v.0.3.3  

  - ~~Add ability for "or" to group. So command line would be "-g "cows or sheep"", but it would split it to sql: "animalgroup='cows' OR animalgroup='sheep'"~~
  - ~~Add sort by days (similar to years and months) where it can be a specific day or a range...~~
  - ~~Restructure code so its not all stuck in main~~
    - I split off all of the functions at least. I'm not sure if any of the main function can be split off? 
  - ~~Add datefrom function. Maybe -from. So you can list only from a specific date.~~
  - ~~Add ability to sort by date oldest/newest newest/oldest~~
  - ~~Add check for config file, if not, prompt user to make the config file at \~/.config/bales~~
  - ~~I'd like to add a config file and store the database somewhere separate(referenced by the config file), and then put it on github maybe? Although then the database would be specific to a computer and you would have to worry about backing it up/sharing it on your own...Something to think about.~~
    - I think this would work. Make a separate folder for the database(also a git folder), and make that the folder that bales -push and -pull works on. So it still updates the database, can still sync the database between computers, and not share the private data...
  - ~~Add -date to list function, so you can see on a specific day what you used...After the work I did to rewrite the list function, it should be fairly easy to add more features like this.~~
  - ~~Rewrite the list function. Make it so that you can add things together. If you say group, it tags onto the select * from bales statement "where animalgroup=" and then if you ad more it says "and so and so"...It should work.~~
  - ~~Add check to date if entered to make sure it is actually a date and not some random string...Regex?~~
  - ~~Add ability to send custom sql command to list.~~
  - ~~Add flag to pull/push from git to update database~~
    - Finally got this figured out...
  - ~~Add ability to get entries from date range. Maybe separate by year/month. Individual or range. Have to be year and month....~~
    - This is implemented somewhat, although probably done poorly. Not sure how to organize it...it works, but I'm sure there are lots of weird errors if you don't enter it how it expects. 
  - ~~Look into flag to see if you can add a description for -h. Add more info to help~~
  - ~~Add ability to list only a specific group. i.e. bales -l -g bulls would query everything for just bulls~~
  - ~~Add ability to get a total count for all | square/round | specific group.~~
  - ~~Add ability to get total count for square or round bales~~
