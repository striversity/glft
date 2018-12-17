# Section 5 : Exercise 01 - Largest Cities in the World

## Todo 1

Given a file with Comma-Separated-Values representing of the largest cities in the World, write a Go program that does the following:

1. Read the records form the data file. Data file name is passed as a program argument to the program.
    * Hint: Use os.Args
    * TIP: Examine a few lines of the data file to see what the records look like.
2. For each line read from the file, initialize a struct object.
    * TIP: Write a function which takes a CSV string and returns a object.
    * TIP: You will need to use the packages 'strings' and 'strconv'.
    * HINT: If you are using the 'input.FileReader' object, be sure to check for io.EOF when reading records and handle it accordingly.
3. Print names of the countries with 5 or more of the largest cities in the data set.

## Requirements

1. You must use a Struct type to represent a record from the file.
2. Use more than one files for readability.

## Testing

Sample data provided in *cities.csv* file within the *stub* and *solution* directories. Data was taken from http://www.citymayors.com/features/largest_cities1.html.