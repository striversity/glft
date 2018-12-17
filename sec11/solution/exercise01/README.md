# Section 11 - Exercise 01 : Page Statistics

Write a complete HTTP client application in golang that will fetch a page and print some stats about the page.

## TODO 1 - Fetch the page and determine how many types of HTML tags are used alone with the page size

There are many HTML tags, but a few to get started are:

| Tag name  | tag      |
| --------- | -------- |
| Button    | <button> |
| Div       | <div>    |
| Form      | <form>   |
| Image     | <img>    |
| Input     | <input>  |
| Link      | <a>      |
| Paragraph | <p>      |

See https://www.w3schools.com/tags/default.asp

* TIP: Consider using _bytes.Count()_ to find nubmer of occurrences.