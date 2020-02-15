package macoshelpers

import (
	"os/exec"
	"strconv"
)

// RunDefaultsDeleteCmd runs command `defaults delete`
func RunDefaultsDeleteCmd(location string, property string) error {
	cmd := exec.Command("defaults", "delete", location, property)

	return cmd.Run()
}

// RunDefaultsWriteStringCmd runs command `defaults write` accepting a string value
func RunDefaultsWriteStringCmd(location string, property string, value string) error {
	cmd := exec.Command("defaults", "write", location, property, "-string", value)

	return cmd.Run()
}

// RunDefaultsWriteIntCmd runs command `defaults write` accepting an integer value
func RunDefaultsWriteIntCmd(location string, property string, value int) error {
	cmd := exec.Command("defaults", "write", location, property, "-int", strconv.Itoa(value))

	return cmd.Run()
}

// RunDefaultsArrayAddCmd runs command `defaults write` with `-array-add`
func RunDefaultsArrayAddCmd(location string, property string, value string) error {
	cmd := exec.Command("defaults", "write", location, property, "-array-add", value)

	return cmd.Run()
}
