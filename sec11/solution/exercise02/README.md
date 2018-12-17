# Section 11 - Exercise 02 : Web Greeter

Write a Go application which renders a Web page with an input box. Upon submission, the app will use the submitted user name to render a dynamic page with a greeting for the user.

## TODO 1 - Greet the user with a custom message and the server's date and time

If the user enters 'Jane' on the home web page, then the app will reply with something like:

    Hi, Jane. Today is Wed, 28 Nov 2018. The current time is 11:05AM.

- TIP: _github.com/striversity/glft/shared/cal_ provides conventince functions for date and time. Or you can get the same info from the _time_ standard package.
