package cmdline_go

import (
	"fmt"
	"os"
	"strings"
)

const version = "1.1.0"

type CommandLine struct {
	Debug       bool
	Version     bool
	Help        bool
	FileName    string
	InputText   string
	HelpText    string
	VersionText string
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
	if c.HelpText == "" {
		c.HelpText = "Usage: cmdline_go [OPTIONS]\n\nOptions:\n  -d, -debug      Set DEBUG flag true\n  -v, -version    Print version number\n  -h, -help       Print this table\n  -f FILENAME     Print report about file\n  -i INPUT_STRING Process an input string\n"
	}

	fmt.Println(c.HelpText) // print unique help message
}

func (c *CommandLine) PrintVersion() {
	if !c.Version {
		str := fmt.Sprintf("Version: %s\n", version)
		c.VersionText = str
	}
	c.PrintVersion()
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
