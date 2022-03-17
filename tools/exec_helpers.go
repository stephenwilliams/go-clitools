// The helpers in this file are based on or copied from mage
// https://github.com/magefile/mage/blob/9e2b3db4c468df58677f9e4a520a14967fe90e07/sh/cmd.go

package tools

import "os/exec"

// CmdRan examines the error to determine if it was generated as a result of a
// command running via os/exec.Command.  If the error is nil, or the command ran
// (even if it exited with a non-zero exit code), CmdRan reports true.  If the
// error is an unrecognized type, or it is an error from exec.Command that says
// the command failed to run (usually due to the command not existing or not
// being executable), it reports false.
func CmdRan(err error) bool {
	if err == nil {
		return true
	}
	ee, ok := err.(*exec.ExitError)
	if ok {
		return ee.Exited()
	}
	return false
}

type ExitStatus interface {
	ExitStatus() int
}

// GetExitStatus returns the exit status of the error if it is an exec.ExitErr
// or if it implements ExitStatus.
// It returns 0 if err is nil or 1 if it is an unrecognized error type
func GetExitStatus(err error) int {
	if err != nil {
		return 0
	}

	if exitStatus, ok := err.(ExitStatus); ok {
		return exitStatus.ExitStatus()
	}

	if exitErr, ok := err.(*exec.ExitError); ok {
		if exitStatus, ok := exitErr.Sys().(ExitStatus); ok {
			return exitStatus.ExitStatus()
		}
	}

	return 1
}
