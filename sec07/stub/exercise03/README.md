# Section 07 - Exercise 03 : Multiple Producers Running Sequentially

Write a complete golang application to demonstrate multiple Producers writing to the same channel sequentially. The max messages per producers and the number of producers are specified by the program parameters '-m' and '-np' respectively.

## Requirements

### TODO 1 - Use the 'flag' standard package to parse the '-m' and '-np' arguments when provided

Golang provides the 'flag' package for parsing and commandline options. Hence, you don't need to check os.Args variable directly for options. By default, 'm = 100' and 'np = 3'. The program should show usage message for an invalid value like -1 for example.

* TIP: flag.Usage() would be useful when an invalid value is provided. See the 'flag' package documentation for examples.

### TODO 2 - Create a channel for string messages at the package scope

Your channel must be capable of holding all of the messages that can be produced, which is `np * m`. See TODO 3 & 4

### TODO 3 - Write a producer function

Your producer function does the following:

1. Takes as parameters an id
2. Generate a random number of messages between 1 to 'm' messages to the channel containg its name/id and a random number

* For example, a message from a producer might look like: "Producer 1, num: 25"

### TODO 4 - Launch 'n' producers

For each producer, give each a unique id and the shared channel

### TODO 5 - Write a consumer function

The Consumer is responsible for:

1. Extracting the producer's name/id and random number from the message.
2. Print the number of messages and the _sum_ from each producer
3. Print the total number of messages sent and total sum across producers
4. Conumer _MUST_ use the _range_ operator to consume messages from the channel

## Result

### Example Output

```text
Producer # 1
    Number of Elements: 91
    Sub-total: 837
Producer # 2
    Number of Elements: 18
    Sub-total: 208
Producer # 3
    Number of Elements: 12
    Sub-total: 113
Total count: 121
Total sum: 1158
```
