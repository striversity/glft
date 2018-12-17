# Section 11 - Lab 01 : Remote Command Runner

Write a golang TCP/IP server and client applications, which will work together to run a command sent by the client to the server.

- HINT: Start with the program from Section 10 Lab 1.

## TODO 1 - Write a TCP/IP server applcation which reads a command to run from a client

When a client connects to the server, the client must sendas a JSON string describing the command the server must run. The format of the json string is: '{"Name":"cmd","Args":[arg0, arg1, ...]}'. Additionally, the follow type can be use to encode/decode JSON string with _encoding/json_ standard package. The server uses the _os/exec_ package to run the command, copying all further data the client sents to the 'command' input and any output back to the client.

```go
type Command struct{
    Name string
    Args []string
}
```

## TODO 2 - Write a TCP/IP client application which sends a command to the server to execute

When the client start, it uses the commandline argument specified to construct a command to send to the server. By default, the client tries to connect to "localhost:12345", but if the '-s' option is specified, the specified server address is used. NOTE: Use the standard _flag_ package to easily handle program options. _TIP_: To get the remaining commandline arguments, use _flag.Args()_, dont' use _os.Args_ if you are using the _flag_ package.

### REQUIREMENTS

1. Use _flag_ package
2. The client must be able to send input to the remote command being run by the server.
