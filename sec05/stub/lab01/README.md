# Section 5 : Lab 01 - People Database Stats

## Todo 1

Given a file with Comma-Separated-Values represensenting information about individuals, write a Go program that does the following:

1. Read the records form the data file. Data file name is passed as a program argument to the program.
    * Hint: Use os.Args
    * TIP: Examine a few lines of the data file to see what the records look like.
2. For each line read from the file, initialize a Person struct object.
    * TIP: Write a function which takes a CSV string and returns a Person object.
    * TIP: You will need to use the packages 'strings' and 'strconv'.
    * HINT: If using the 'input.FileReader' object, be sure to check for io.EOF when reading records and handle it occordingly.
3. Compute the following stats by gender:
    1. total number of records
    2. min salary
    3. max slary
    4. average salary

## Requriements

1. You must use a Struct type to represent a record from the file.
    * HINT: Your program may be eaiser and cleaner if you user an additional struct for stats data. See the solution for one example.
2. Use more than one files for readablity.
    * TIP: Consider breaking your program into packages.
    * TIP: If you are writing a package:
        * use `go install` to install it
        * use `go doc` in the pacakge dir to check import and usage

## Testing

Sample data provided in *people.csv* file within the *stub* and *solution* directories. Data was generated at https://www.mockaroo.com/.