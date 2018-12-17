# Section 06 - Lab 02 : Word Count (concurrent)

## TODO 1

Write a golang application that count the number of each word in one or more input files.

### Requirements

1. Your program must _use_ concurrency.
   - TIP: You may want to start with a copy of Section 06 - Lab 01.
2. Your program must accept one or more filenames on the commandline as input.
3. After processing each file, your program must print the list of words found in _all_ of the files and how many time that word appeared.

   - Example output:

     | Count | Word         |
     | ----- | ------------ |
     | 8     | Hendrikhovna |
     | 30    | succeeded    |
     | 32    | Russia.      |
     | 2     | worry,       |
     | 1     | carpus       |

4. Your program must print the _time_ taken for word counting as the last output.
   - NOTE: This time need not include the time for printing out the result.
   - TIP: Use time.Now() a the start of the program and time.Since() at the end to get the elapse time.
     - Additionally, see supplemental video for this lab.

## TODO 2 - Compare Runtime

### Requrements

1. Compare the time taken by the concurrent word count application to the iterative verion on the same input.
2. If your concurrent version doesn't run faster for large number of text, consider why not.
