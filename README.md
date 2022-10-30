# Bales
#### A program for keeping track of bales that you have fed

## Current Status

Not currently really ready for others to use. If you really wanted to, you can git clone this repository, create the  config.yaml at ~/.config/bales/config.yaml and then `go build main.go`. If it were me, I'd `go build -o bales main.go` then you can put bales in your path somewhere. bales -h will give you help usage.

### Example ~/.config/bales/config.yaml

```yaml
# Database dir is the directory you want to store your databases in.
# It can be a git repo, but doesn't have to be...
DatabaseDir: /home/username/git/databases

# RealDatabase is the legit database
RealDatabase: balesDatabase.db

# TestDatabase is a database you can use to test features
TestDatabase: balesTestDatabase.db

```


## To-do:
  - Add ability to sort by date oldest/newest newest/oldest
  - Add check for config file, if not, prompt user to make the config file at ~/.config/bales
  - Restructure code so its not all stuck in main
  - make default list only list the current month. add -all flag to show entire database
  - Add ability for "or" to group. So command line would be "-g "cows or sheep"", but it would split it to sql: "animalgroup='cows' OR animalgroup='sheep'"
  - add datefrom function. Maybe -from. So you can list only from a specific date.
  - ~~I'd like to add a config file and store the database somewhere separate(referenced by the config file), and then put it on github maybe? Although then the database would be specific to a computer and you would have to worry about backing it up/sharing it on your own...Something to think about.~~
    - I think this would work. Make a separate folder for the database(also a git folder), and make that the folder that bales -push and -pull works on. So it still updates the database, can still sync the database between computers, and not share the private data...
  - ~~add -date to list function, so you can see on a specific day what you used...After the work I did to rewrite the list function, it should be fairly easy to add more features like this.~~
  - ~~Rewrite the list function. Make it so that you can add things together. If you say group, it tags onto the select * from bales statement "where animalgroup=" and then if you ad more it says "and so and so"...It should work.~~
  - Add an ability to get average. For square/round. Maybe by group as well. For daily weekly monthly?
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
