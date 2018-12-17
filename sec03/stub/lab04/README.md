# Section 03 : Lab 03 - _my_at_, A Simplified _at_

Our objective is to write a very simplified version of the Unix/Linux _at_ utility. The _at_ utility is a simple reminder. It take a 'time specification' as a program argument, then prompts for a message which will be displayed after the specified time passes.

[Here is an example of using _at_](https://www.computerhope.com/unix/uat.htm).  There are a number of ways to specify time in _at_, but we will simplify that to:

## Example Usage

```zsh
my_at now + <duration>
```

Where _duration_ is defined in Golang _time_ package [func ParseDuration](https://golang.org/pkg/time/#ParseDuration).

| command | comment |
|---------|---------|
| `my_at now + 2s` | # 2 seconds from now |
| `my_at now + 2m2s5ns` |# 2 minutes 2 seconds and 5 nanoseconds from now |
| `my_at now + 1h` | # 1 hour from now |


## TODO 1

Write a *simplified* _at_ program using Golang.

1. Your program must use the 'os' pacakge 'os.Args' array to validate the correct number of arguments are provided.
2. If the number of inputs are incorrect, call a function to print help with example usage.
3. Use 'time.ParseDuration()' to get the elaspe time and error if the time specified couldn't be parsed correctly.
4. _NOTE:_ We haven't covered error handling yet, so ignore the 2nd return value from 'time.ParseDuration()'.
5. Prompt the user's for the reminder message. *HINT* Use the 'input.PromptString()'.
6. Prints the message after the time has elaspe/passed. *HINT:* Use 'time.Sleep()' to allow the time to pass. See Golang's _time_ pacakge [Sleep function](https://golang.org/pkg/time/#Sleep)

### Example program output:

```zsh
go install -o my_at main.go
my_at now + 2s
```

> Time to make that call
Time to make that call
