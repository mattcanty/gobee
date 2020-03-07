package systempreferences

import (
	"os/exec"

	"github.com/mattcanty/gobee/pkg/config"
	macoshelpers "github.com/mattcanty/gobee/pkg/macos/helpers"
)

// ConfigureDateAndTime - Configures the Date & Time on MacOS
func ConfigureDateAndTime(dateAndTimeConf config.DateAndTime) error {
	err := macoshelpers.RunDefaultsWriteIntCmd("com.apple.menuextra.clock", "FlashDateSeparators", boolToInt(dateAndTimeConf.Clock.FlashSeparators))

	if err != nil {
		return err
	}

	err = macoshelpers.RunDefaultsWriteStringCmd("com.apple.menuextra.clock", "DateFormat", dateAndTimeConf.Clock.Format)

	if err != nil {
		return err
	}

	return killSystemUIServer()
}

func killSystemUIServer() error {
	return exec.Command("killall", "SystemUIServer").Run()
}
