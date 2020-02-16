package macoshelpers

import (
	"os/exec"
	"strconv"
)

// RunDefaultsDeleteCmd runs command `defaults delete`
func RunDefaultsDeleteCmd(location string, property string) error {
	return exec.Command("defaults", "delete", location, property).Run()
}

// RunDefaultsWriteStringCmd runs command `defaults write` accepting a string value
func RunDefaultsWriteStringCmd(location string, property string, value string) error {
	return exec.Command("defaults", "write", location, property, "-string", value).Run()
}

// RunDefaultsWriteIntCmd runs command `defaults write` accepting an integer value
func RunDefaultsWriteIntCmd(location string, property string, value int) error {
	return exec.Command("defaults", "write", location, property, "-int", strconv.Itoa(value)).Run()
}

// RunDefaultsArrayAddCmd runs command `defaults write` with `-array-add`
func RunDefaultsArrayAddCmd(location string, property string, value string) error {
	return exec.Command("defaults", "write", location, property, "-array-add", value).Run()
}
