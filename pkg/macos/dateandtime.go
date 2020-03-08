package macos

// DateAndTime configuration for MacOS Date & Time
type DateAndTime struct {
	Clock struct {
		Show            bool   `yaml:"show"`
		Type            string `yaml:"type"`
		Format          string `yaml:"format"`
		FlashSeparators bool   `yaml:"flashSeparators"`
	}
}

// GetChanges gets the changes to be applied
func (config DateAndTime) GetChanges() []string {
	var returnVal []string

	return returnVal
}

// ApplyChanges applies those changes
func (config DateAndTime) ApplyChanges() []string {
	var returnVal []string

	return returnVal
}

// ConfigureDateAndTime - Configures the Date & Time on MacOS
// func ConfigureDateAndTime(dateAndTimeConf config.DateAndTime) error {
// 	err := macoshelpers.RunDefaultsWriteIntCmd("com.apple.menuextra.clock", "FlashDateSeparators", boolToInt(dateAndTimeConf.Clock.FlashSeparators))

// 	if err != nil {
// 		return err
// 	}

// 	err = macoshelpers.RunDefaultsWriteStringCmd("com.apple.menuextra.clock", "DateFormat", dateAndTimeConf.Clock.Format)

// 	if err != nil {
// 		return err
// 	}

// 	return killSystemUIServer()
// }

// func killSystemUIServer() error {
// 	return exec.Command("killall", "SystemUIServer").Run()
// }
