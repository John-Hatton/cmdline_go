package cmdline_go

import (
	"fmt"
	"os"
	"strings"
)

const version = "1.0.0"

type CommandLine struct {
	Debug    bool
	Version  bool
	Help     bool
	FileName string
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
		case "-j":
			fmt.Println("John says hi!")
			os.Exit(0)
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
	fmt.Println("Usage: cmdline [OPTIONS]")
	fmt.Println("")
	fmt.Println("Options:")
	fmt.Println("  -d, -debug  Set DEBUG flag true")
	fmt.Println("  -v, -version  Print version number")
	fmt.Println("  -h, -help  Print this table")
	fmt.Println("  -f FILENAME  Print report about File")
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
