# Section 06 - Exercise 02 : Optimal Timing for Multiple Prime Searchers Running Concurrently

Write a complete golang application which uses multiple Prime Searcher
functions running concurrently to find all the prime numbers between 1 and 100,000.

However, the application _MUST_ end immediately after the _last_ Prime Search finishes.

## TODO 1 - Copy the code written for Section 06 - Exercise 01

## TODO 2 - Use sync.WaitGroup in stead of time.Sleep in main()

## TODO 3 - Run between 3 to 10 Prime Searchers

Check your results:

1. Compile and run your program
2. Count the number of lines in the file, it should be equal to 9,592.
    * TIP
        * Mac/Linux Users: `go build -o prime && ./prime | wc -l`
        * Windows Users: `go build -o prime && ./prime > results.txt && code results.txt`
    * See the entry for number of primes between 1 and 100,000 at https://primes.utm.edu/howmany.html
