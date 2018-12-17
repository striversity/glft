# Seciton 06 - Exercise 04 : Password Generator

Write a complete golang application to generate unique passwords.

## Requirements

1. If using a Western keyboard, use the following character classes in your password:
    * letters: a-z and A-Z
    * numbers: 0-9
    * special characters: -=~!@#$%^&*()_+[]\{}|;':",./<>?
2. User may specify password length using '-l' option, else default is 16.
    * TIP: Use the standard _flag_ pacakge.
3. MUST use _select_ statement to pick character from character class
4. Prefer more _letters_ to _numbers_, and more _numbers_ to _special characters_.
    * HINT: Use case statements to increase probability of one character set over another.