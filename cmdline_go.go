package cmdline_go

import (
	"fmt"
	"os"
	"strings"
)

const version = "1.1.2"

type CommandLine struct {
	Debug        bool
	Version      bool
	Help         bool
	FileName     string
	InputText    string
	HelpText     string
	VersionText  string
	LogToConsole bool   // new flag to log output to console
	LogFileName  string // new flag to specify log file name
}

func (c *CommandLine) Parse(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("No arguments provided")
	}
	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch strings.ToLower(arg) {
		case "-d", "-debug":
			c.Debug = true
		case "-v", "-version":
			c.Version = true
		case "-h", "-help":
			c.Help = true
		case "-i":
			if i+1 < len(args) {
				c.InputText = args[i+1]
				return nil
			} else {
				return fmt.Errorf("-i requires an input string argument")
			}
		case "-f":
			if i+1 < len(args) {
				c.FileName = args[i+1]
				return nil
			} else {
				return fmt.Errorf("-f requires a filename argument")
			}
		case "-o":
			if c.LogFileName == "" {
				return fmt.Errorf("-o requires a log filename argument")
			}
			c.LogToConsole = true
		case "-l":
			if i+1 < len(args) {
				c.LogFileName = args[i+1]
				return nil
			} else {
				return fmt.Errorf("-l requires a log filename argument")
			}
		}
	}
	return nil
}

func (c *CommandLine) PrintHelp() {
	if c.HelpText == "" {
		c.HelpText = "Usage: cmdline_go [OPTIONS]\n\nOptions:\n  -d, -debug        Set DEBUG flag true\n  -v, -version      Print version number\n  -h, -help         Print this table\n  -f FILENAME       Print report about file\n  -i INPUT_STRING   Process an input string\n  -o LOG_TO_CONSOLE Log output to console\n  -l LOG_FILENAME   Save output to log file\n"
	}

	fmt.Println(c.HelpText) // print unique help message
}

func (c *CommandLine) PrintVersion() {
	if !c.Version {
		str := fmt.Sprintf("cmdline_go -- Version: %s\n", version)
		c.VersionText = str
	}
	fmt.Println(c.VersionText)
}

func (c *CommandLine) PrintReport() {
	fmt.Printf("Report for file %s\n", c.FileName)
}

func (c *CommandLine) Process() error {
	if c.Help {
		c.PrintHelp()
		os.Exit(0)
	}
	if c.Version {
		c.PrintVersion()
		os.Exit(0)
	}
	if c.FileName != "" {
		c.PrintReport()
		os.Exit(0)
	}

	// new code to log output to file
	if c.LogToConsole {
		if c.LogFileName == "" {
			return fmt.Errorf("-o requires a log filename argument")
		}
		logFile, err := os.Create(c.LogFileName)
		if err != nil {
			return fmt.Errorf("Error creating log file: %v", err)
		}
		defer logFile.Close()
		if c.Debug {
			fmt.Fprintf(os.Stderr, "Logging to file%s\n", c.LogFileName)
		}
		fmt.Fprintln(logFile, "Log output:")
		fmt.Fprintln(logFile, "-----------")
		fmt.Fprintln(logFile, c.InputText)
		fmt.Fprintln(logFile, "-----------")
	}

	return nil
}
