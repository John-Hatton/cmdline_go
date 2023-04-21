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

**Note: This is only good for one session of the terminal. To set permenantly, do the following:

    Open the Start menu and search for "Environment Variables".
    Click on "Edit the system environment variables".
    In the System Properties dialog box, click on the "Environment Variables" button.
    Under "User variables", click the "New" button.
    In the "Variable name" field, enter "GOPATH".
    In the "Variable value" field, enter the path to your Go workspace (e.g. "C:\Users\YourUsername\go").
    Click "OK" to save the changes.

--- 

## Setting Up Go

Next you will need to create src, pkg, and bin directories in your new GOPATH folder.

    mkdir src
    mkdir pkg
    mkdir bin

Now we are ready to compile and run the code.

---

## Import Project

To Import the project, be sure git is installed. If not, see the link below:
    https://gitforwindows.org/

**Note: Be sure to navigate to your GOPATH with:

    cd $env:GOPATH

Then cd into the newly created src folder. If it's not there, remember you need to create the
folders from above:
    
    mkdir src
    mkdir pkg
    mkdir bin 


Once you're in your src folder, you've installed Go, and Git is installed, do the following:

    git clone git@github.com:John-Hatton/cmdline_go.git

The project has successfully been imported!

---

## Testing

There are a few test cases that cover the input parameters. They are as follows:

-d or -debug
    

Set DEBUG flag true

 

-v or -version
    

Print version number (must be constant within program) and die

 

-h or -help
    

Print this table and die

 

-f FILENAME
    

Two-part parameter, -f means there must be a filename, and print report and die

---


To run tests with go simply do the following:

    go test


This will run the tests, and you should get the following message:

PASS
ok      cmdline 0.322s

---

## Building for Debug Purposes

If you want to open the executable, say perhaps with GDB, simply do the following:

    go build -gcflags "-N -l" -o cmdline_go cmdline_go.go


