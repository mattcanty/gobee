package defaults

import (
	"os/exec"
	"strconv"
)

// Delete runs command `defaults delete`
func Delete(location string, property string) error {
	cmd := exec.Command("defaults", "delete", location, property)

	return cmd.Run()
}

// WriteString runs command `defaults write` accepting a string value
func WriteString(location string, property string, value string) error {
	cmd := exec.Command("defaults", "write", location, property, "-string", value)

	return cmd.Run()
}

// WriteInt runs command `defaults write` accepting an integer value
func WriteInt(location string, property string, value int) error {
	cmd := exec.Command("defaults", "write", location, property, "-int", strconv.Itoa(value))

	return cmd.Run()
}

// ArrayAdd runs command `defaults write` with `-array-add`
func ArrayAdd(location string, property string, value string) error {
	cmd := exec.Command("defaults", "write", location, property, "-array-add", value)

	return cmd.Run()
}
