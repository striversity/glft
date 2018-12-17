# Section 10 - Exercise 03 : myTree - Simple 'tree' program

Write a golang application 'myTree' which counts the *TOTAL* number of files and directory for the speicfied paths. If a path is not a directory, print a simply error message and continue. myTree is a simplified version of the Unix/Linux 'tree' utility. 'tree' shows information about file and directories, and possibly a 'tree-like' view of a specified path.

For example, `tree` of this directory shows this:
  .
  ├── README.md
  └── mytree
      └── main.go
  
  1 directory, 2 files

But we will be writing a much more simple version, it only prints the number of directories and files. So our 'mytree' on this exercise dir shows this:

  1 directories, 2 files

We could have listed the name of directory and file as we iterate over the list, but simply didn't.

## REQUIRMENTS

1. Must allow 0 or more directories to be specified on the commandline
2. Use the current directory, represented by '.', if paths are not provided
