package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_CommandLine_NoArgs(t *testing.T) {
	// Test case 1: No parameters
	c := CommandLine{}
	args1 := []string{}
	err := c.Parse(args1)
	expectedErr := fmt.Errorf("No arguments provided")
	if err.Error() != expectedErr.Error() {
		t.Errorf("Test case 1 failed. Expected error: %v, but got: %v", expectedErr, err)
	}
	expected1 := CommandLine{Debug: false, Version: false, Help: false, FileName: ""}
	if !reflect.DeepEqual(c, expected1) {
		t.Errorf("Test case 1 failed. Expected: %v, but got: %v", expected1, c)
	}
}

func Test_CommandLine_Debug(t *testing.T) {
	// Test case 2: -d or -debug or -D or -DEBUG flag
	c := CommandLine{}
	args2 := []string{"-d"}
	err := c.Parse(args2)
	if err != nil {
		t.Errorf("Test case 2 failed. Expected no error, but got: %v", err)
	}
	expected2 := CommandLine{Debug: true, Version: false, Help: false, FileName: ""}
	if !reflect.DeepEqual(c, expected2) {
		t.Errorf("Test case 2 failed. Expected: %v, but got: %v", expected2, c)
	}

	c = CommandLine{}
	args3 := []string{"-debug"}
	err = c.Parse(args3)
	if err != nil {
		t.Errorf("Test case 3 failed. Expected no error, but got: %v", err)
	}
	if !reflect.DeepEqual(c, expected2) {
		t.Errorf("Test case 3 failed. Expected: %v, but got: %v", expected2, c)
	}

	c = CommandLine{}
	args4 := []string{"-D"}
	err = c.Parse(args4)
	if err != nil {
		t.Errorf("Test case 4 failed. Expected no error, but got: %v", err)
	}
	if !reflect.DeepEqual(c, expected2) {
		t.Errorf("Test case 4 failed. Expected: %v, but got: %v", expected2, c)
	}

	c = CommandLine{}
	args5 := []string{"-DEBUG"}
	err = c.Parse(args5)
	if err != nil {
		t.Errorf("Test case 5 failed. Expected no error, but got: %v", err)
	}
	if !reflect.DeepEqual(c, expected2) {
		t.Errorf("Test case 5 failed. Expected: %v, but got: %v", expected2, c)
	}
}

func Test_CommandLine_Version(t *testing.T) {
	// Test case 3: -v or -debug or -D or -VERSION flag
	cmd := CommandLine{}
	args2 := []string{"-v"}
	err := cmd.Parse(args2)
	if err != nil {
		t.Errorf("Test case 2 failed. Expected no error, but got: %v", err)
	}
	expected2 := CommandLine{Debug: false, Version: true, Help: false, FileName: ""}
	if !reflect.DeepEqual(cmd, expected2) {
		t.Errorf("Test case 2 failed. Expected: %v, but got: %v", expected2, cmd)
	}

	cmd = CommandLine{}
	args3 := []string{"-version"}
	err = cmd.Parse(args3)
	if err != nil {
		t.Errorf("Test case 3 failed. Expected no error, but got: %v", err)
	}
	if !reflect.DeepEqual(cmd, expected2) {
		t.Errorf("Test case 3 failed. Expected: %v, but got: %v", expected2, cmd)
	}

	cmd = CommandLine{}
	args4 := []string{"-V"}
	err = cmd.Parse(args4)
	if err != nil {
		t.Errorf("Test case 4 failed. Expected no error, but got: %v", err)
	}
	if !reflect.DeepEqual(cmd, expected2) {
		t.Errorf("Test case 4 failed. Expected: %v, but got: %v", expected2, cmd)
	}

	cmd = CommandLine{}
	args5 := []string{"-VERSION"}
	err = cmd.Parse(args5)
	if err != nil {
		t.Errorf("Test case 5 failed. Expected no error, but got: %v", err)
	}
	if !reflect.DeepEqual(cmd, expected2) {
		t.Errorf("Test case 5 failed. Expected: %v, but got: %v", expected2, cmd)
	}
}

func Test_CommandLine_Help(t *testing.T) {
	// Test case 3: -v or -debug or -D or -VERSION flag
	cmd := CommandLine{}
	args2 := []string{"-h"}
	err := cmd.Parse(args2)
	if err != nil {
		t.Errorf("Test case 2 failed. Expected no error, but got: %v", err)
	}
	expected2 := CommandLine{Debug: false, Version: false, Help: true, FileName: ""}
	if !reflect.DeepEqual(cmd, expected2) {
		t.Errorf("Test case 2 failed. Expected: %v, but got: %v", expected2, cmd)
	}

	cmd = CommandLine{}
	args3 := []string{"-help"}
	err = cmd.Parse(args3)
	if err != nil {
		t.Errorf("Test case 3 failed. Expected no error, but got: %v", err)
	}
	if !reflect.DeepEqual(cmd, expected2) {
		t.Errorf("Test case 3 failed. Expected: %v, but got: %v", expected2, cmd)
	}

	cmd = CommandLine{}
	args4 := []string{"-H"}
	err = cmd.Parse(args4)
	if err != nil {
		t.Errorf("Test case 4 failed. Expected no error, but got: %v", err)
	}
	if !reflect.DeepEqual(cmd, expected2) {
		t.Errorf("Test case 4 failed. Expected: %v, but got: %v", expected2, cmd)
	}

	cmd = CommandLine{}
	args5 := []string{"-HELP"}
	err = cmd.Parse(args5)
	if err != nil {
		t.Errorf("Test case 5 failed. Expected no error, but got: %v", err)
	}
	if !reflect.DeepEqual(cmd, expected2) {
		t.Errorf("Test case 5 failed. Expected: %v, but got: %v", expected2, cmd)
	}
}

func Test_CommandLine_FileNameFlag(t *testing.T) {
	// Test case 6: FileName flag is set
	cmd := CommandLine{}
	args6 := []string{"-f", "test.txt"}
	err := cmd.Parse(args6)
	if err != nil {
		t.Errorf("Test case 6 failed. Expected no error, but got: %v", err)
	}
	expected6 := CommandLine{Debug: false, Version: false, Help: false, FileName: "test.txt"}
	if !reflect.DeepEqual(cmd, expected6) {
		t.Errorf("Test case 6 failed. Expected: %v, but got: %v", expected6, cmd)
	}
}
