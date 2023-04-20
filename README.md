# cmdline 

### A Go Package read and process command line parameters

---


## Introduction

This is a package written in Go, which takes a string, parses the input, sets the appropriate 
attributes in the cmdline struct, and then we use this to setup our future projects, so they all
have a cmdline struct as a member attribute.

---

## Installing Go:

Go can be installed in Windows, Linux, or Mac. Simply go to:

    https://go.dev/dl/

Then select your version, and install as usual.

You will need to then set the GOPATH environment variable. 

To do so in Powershell (in Windows)

    mkdir C:\Users\YourUsername\go
    $env:GOPATH = "C:\Users\YourUsername\go"

To verify it's done, do the following:

    echo $env:GOPATH

--- 

## Setting Up Go

Next you will need to create src, pkg, and bin directories in your new GOROOT folder.

    mkdir src
    mkdir pkg
    mkdir bin

Now we are ready to compile and run the code.

---

## Import Project

To Import the project, be sure git is installed. If not, see the link below:
    https://gitforwindows.org/

In the Powershell terminal run the following:








