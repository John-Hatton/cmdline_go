package cmdline_go

import (
	"fmt"
	"os"
	"strings"
)

const version = "1.0.8"

type CommandLine struct {
	Debug     bool
	Version   bool
	Help      bool
	FileName  string
	InputText string
	HelpText  string // unique help message
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
		}
	}
	return nil
}

func (c *CommandLine) PrintHelp() {
	fmt.Println(c.HelpText) // print unique help message
	fmt.Println("")
	fmt.Println("Usage: cmdline_go [OPTIONS]")
	fmt.Println("")
	fmt.Println("Options:")
	fmt.Println("  -d, -debug      Set DEBUG flag true")
	fmt.Println("  -v, -version    Print version number")
	fmt.Println("  -h, -help       Print this table")
	fmt.Println("  -f FILENAME     Print report about file")
	fmt.Println("  -i INPUT_STRING Process an input string")
}

func (c *CommandLine) PrintHelpText() {
	fmt.Println(c.HelpText)
}

func (c *CommandLine) PrintVersion() {
	fmt.Printf("Version: %s\n", version)
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
	return nil
}
