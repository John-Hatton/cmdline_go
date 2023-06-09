package cmdline_go

import (
	"fmt"
	"os"
)

const version = "1.2.5"

type CommandLine struct {
	Input        bool
	Output       bool
	File         bool
	Log          bool
	Debug        bool
	Version      bool
	Help         bool
	FileName     string
	InputText    string
	OutputText   string
	HelpText     string
	VersionText  string
	LogToConsole bool
	LogFileName  string
}

func (c *CommandLine) Parse(args []string) error {
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-d", "-D", "--debug", "--DEBUG":
			c.Debug = true
		case "-v", "-V", "--version", "--VERSION":
			c.Version = true
		case "-h", "-H", "--help", "--HELP":
			c.Help = true
		case "-f", "-F", "--filename", "--FILENAME":
			i++
			if i >= len(args) {
				return fmt.Errorf("-f requires a filename argument")
			}
			c.File = true
			c.FileName = args[i]
		case "-i", "-I", "--input", "--INPUT":
			i++
			if i >= len(args) {
				return fmt.Errorf("-i requires an input string argument")
			}
			c.Input = true
			c.InputText = args[i]
		case "-o", "-O", "--output", "--OUTPUT":
			if i+1 < len(args) {
				c.LogFileName = args[i+1]
				c.Output = true
				c.LogToConsole = false
				i++
				continue // Skip to the next iteration since we have already consumed the next argument
			}
			// If no argument provided, fall back to logging output to console
			c.LogToConsole = true
			c.Output = true
		case "-l", "-L", "--logfile", "--LOGFILE":
			i++
			if i >= len(args) {
				return fmt.Errorf("-l requires a log filename argument")
			}
			c.Log = true
			c.LogToConsole = true
			c.LogFileName = args[i]
		default:
			return fmt.Errorf("unknown option: %s", args[i])
		}
	}
	return nil
}

func (c *CommandLine) PrintHelp() {
	if c.HelpText == "" {
		c.HelpText = "Usage: cmdline_go [OPTIONS]\n\nOptions:\n  -d, --debug        Set DEBUG flag true\n  -v, --version      Print version number\n  -h, --help         Print this table\n  -f FILENAME        Take a file in as input\n  -i INPUT_STRING    Process an input string\n  -o [LOG_TO_FILE]   Log output to file (default: console)\n  -l LOG_FILENAME    Save output to log file\n"
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
	fmt.Printf("Filename set:  %s\n", c.FileName)
}

func (c *CommandLine) PrintOutput() error {
	if c.LogToConsole {
		if c.LogFileName != "" {
			logFile, err := os.Create(c.LogFileName)
			if err != nil {
				return err
			}
			defer logFile.Close()
			fmt.Fprintln(logFile, c.OutputText)
		} else {
			fmt.Println(c.OutputText)
		}
	} else {
		fmt.Println(c.OutputText)
		if c.Log {
			logFile, err := os.Create(c.LogFileName)
			if err != nil {
				return err
			}
			defer logFile.Close()
			fmt.Fprintln(logFile, c.OutputText)
		}
	}
	return nil
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
	}

	if c.InputText != "" {
		c.OutputText = c.InputText
	}

	if c.OutputText == "" {
		return fmt.Errorf(`This program requires -i for input. See -h for more`)
	}

	err := c.PrintOutput()
	if err != nil {
		return err
	}

	return nil
}
