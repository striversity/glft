# Section 10 - Exercise 05 : Custom Greeter

Write a golang application which prompts a user for their name. If the user was seen before, greet them with a welcome back message. If this is a new user, greet them as a new user and add their name to the db.

## TODO 1 - Prompt user for username

Check user name against file db. If user is new, add them to the file db and greet them with: 'Welcome first time user!' followed by the current day and time. If the user exist in the db, then greet them with 'Welcome back!', followed by teh current day and time. For example:

>     What is your name: jane
>     Hi, Jane
>     Welcome first time user!
>     Today is Fri, 23 Nov 2018. The current time is 11:46AM.

## REQUIREMENTS

1. User name must be properly formatted regardless of how it was entered
2. Create the db file if it doesn't exist, or add user to an existing file
   - TIP: Consider bufio.NewReader() for reading in usernames form the db file. Though bufio.NewScanner() works just as well too.
3. Print some useful information for the user, like the current date and time as a minimum.
   - TIP: See 'time' package or the helper functions in package 'github.com/striversity/glft/shared/cal'.
