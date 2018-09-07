package gorillalog

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/1dustindavis/gorilla/pkg/config"
)

// Store the original values, before we override
var (
	origVerbose     = config.Verbose
	origDebug       = config.Debug
	origProgramData = config.GorillaData
)

func restoreVerbose() {
	config.Verbose = origVerbose
}

func restoreDebug() {
	config.Verbose = origVerbose
}

// TestNewLog tests the creation of the log and it's directory
func TestNewLog(t *testing.T) {
	// Set up a place for test data
	tmpDir := filepath.Join(os.Getenv("TMPDIR"), "gorillalog")
	config.GorillaData = tmpDir

	// Clean up when we are done
	defer func() {
		// Clean up
		config.GorillaData = origProgramData
		os.RemoveAll(tmpDir)
	}()

	// Run the function
	NewLog()

	// Check values
	logDir := tmpDir
	logFile := filepath.Join(tmpDir, "gorilla.log")
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		fmt.Println(err)
		t.Errorf("Log Directory not created: %s", logDir)
	}
	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		t.Errorf("Log File not created: %s", logFile)
	}
}

// TestDebug tests that debug logs properly
func TestDebug(t *testing.T) {
	// Set the output
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
		restoreDebug()
	}()

	// Set up what we want
	prefix := "DEBUG: "
	now := time.Now().Format("2006/01/02 15:04:05 ")
	logString := "Debug String!"
	expected := fmt.Sprint(prefix, now, logString)

	// Run the function
	config.Debug = true
	Debug(logString)

	result := strings.TrimSpace(buf.String())

	// Check out result
	if have, want := result, expected; have != want {
		t.Errorf("-----\nhave %s\nwant %s\n-----", have, want)
	}
}

// ExampleDebugOff tests the output of a log sent to DEBUG while config.Debug is false
func ExampleDebugOff() {
	// Set up what we expect
	logString := "Debug String!"

	// Run the function without debug
	config.Debug = false
	defer restoreDebug()
	Debug(logString)
	// Output:
}

// ExampleDebugOn tests the output of a log sent to DEBUG while config.Debug is true
func ExampleDebugOn() {
	// Set up what we expect
	logString := "Debug String!"

	// Run the function with debug
	config.Debug = true
	defer restoreDebug()

	Debug(logString)
	// Output:
	// Debug String!
}

// TestInfo tests the formatting of a log sent to INFO
func TestInfo(t *testing.T) {
	// Set the output
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	// Set up what we want
	prefix := "INFO: "
	now := time.Now().Format("2006/01/02 15:04:05 ")
	logString := "Info String!"
	expected := fmt.Sprint(prefix, now, logString)

	// Run the function
	Info(logString)

	result := strings.TrimSpace(buf.String())

	// Check out result
	if have, want := result, expected; have != want {
		t.Errorf("-----\nhave %s\nwant %s\n-----", have, want)
	}
}

// ExampleInfoVerboseOff tests the output of a log sent to INFO while config.Verbose is false
func ExampleInfoVerboseOff() {
	// Set up what we expect
	logString := "Info String!"

	// Run the function without verbose
	config.Verbose = false
	defer restoreVerbose()

	Info(logString)
	// Output:
}

// ExampleInfoVerboseOn tests the output of a log sent to INFO while config.Verbose is true
func ExampleInfoVerboseOn() {
	// Set up what we expect
	logString := "Info String!"

	// Run the function with verbose
	config.Verbose = true
	defer restoreVerbose()

	Info(logString)
	// Output:
	// Info String!
}

// TestWarn tests the formatting of a log sent to WARN
func TestWarn(t *testing.T) {
	// Set the output
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	// Set up what we want
	prefix := "WARN: "
	now := time.Now().Format("2006/01/02 15:04:05 ")
	logString := "Warn String!"
	expected := fmt.Sprint(prefix, now, logString)

	// Run the function
	Warn(logString)

	result := strings.TrimSpace(buf.String())

	// Check out result
	if have, want := result, expected; have != want {
		t.Errorf("-----\nhave %s\nwant %s\n-----", have, want)
	}
}

// ExampleWarn tests the output of a log sent to WARN
func ExampleWarn() {
	// Set up what we expect
	logString := "Warn String!"

	// Run the function
	Warn(logString)
	// Output:
	// Warn String!
}

// TestError tests the formatting of a log sent to ERROR
func TestError(t *testing.T) {
	// Set the output
	var buf bytes.Buffer
	log.SetOutput(&buf)
	// Prepare to recover from a panic
	defer func() {
		log.SetOutput(os.Stderr)
		if r := recover(); r == nil {
			t.Errorf("Error didnt panic")
		}
	}()

	// Set up what we want
	prefix := "ERROR: "
	now := time.Now().Format("2006/01/02 15:04:05 ")
	logString := "Error String!"
	expected := fmt.Sprint(prefix, now, logString)

	// Run the function
	Error(logString)

	result := strings.TrimSpace(buf.String())

	// Check out result
	if have, want := result, expected; have != want {
		t.Errorf("-----\nhave %s\nwant %s\n-----", have, want)
	}
}

// ExampleError tests the output of a log sent to ERROR
func ExampleError() {
	// Set up what we expect
	logString := "Error String!"

	// Prepare to recover from a panic
	defer func() {
		recover()
	}()

	// Run the function
	Error(logString)
	// Output:
}
