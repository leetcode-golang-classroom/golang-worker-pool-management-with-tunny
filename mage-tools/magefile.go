//go:build mage
// +build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var Default = Build

// clean the build binary
func Clean() error {
	return sh.Rm("bin")
}

// Creates the binary in the current directory.
func Build() error {
	mg.Deps(Clean)
	if err := sh.Run("go", "mod", "download"); err != nil {
		return err
	}
	return sh.Run("go", "build", "-o", "./bin/go-routine-simple", "./cmd/simple/main.go")
}

// Create the binary in current directory for with-state sample
func Build_state() error {
	mg.Deps(Clean)
	if err := sh.Run("go", "mod", "download"); err != nil {
		return err
	}
	return sh.Run("go", "build", "-o", "./bin/go-state-routine", "./cmd/with-state/main.go")
}

// start the app
func Launch() error {
	mg.Deps(Build)
	err := sh.RunV("./bin/go-routine-simple")
	if err != nil {
		return err
	}
	return nil
}

// start the state app
func Launch_state() error {
	mg.Deps(Build_state)
	err := sh.RunV("./bin/go-state-routine")
	if err != nil {
		return err
	}
	return nil
}

// run the test
func Test() error {
	err := sh.RunV("go", "test", "-v", "./...")
	if err != nil {
		return err
	}
	return nil
}
