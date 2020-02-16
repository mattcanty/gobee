package systempreferences

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/mattcanty/gobee/pkg/config"
	macoshelpers "github.com/mattcanty/gobee/pkg/macos/helpers"
)

// ConfigureDock - Configures the Dock on MacOS
func ConfigureDock(dockConf config.Dock) error {
	err := macoshelpers.RunDefaultsWriteIntCmd("com.apple.dock", "tilesize", dockConf.TileSize)

	if err != nil {
		return err
	}

	err = macoshelpers.RunDefaultsWriteIntCmd("com.apple.dock", "magnification", boolToInt(dockConf.Magnification.Enabled))

	if err != nil {
		return err
	}

	err = macoshelpers.RunDefaultsWriteIntCmd("com.apple.dock", "largesize", dockConf.Magnification.Size)

	if err != nil {
		return err
	}

	err = macoshelpers.RunDefaultsWriteStringCmd("com.apple.dock", "orientation", dockConf.Position)

	if err != nil {
		return err
	}

	err = macoshelpers.RunDefaultsWriteStringCmd("com.apple.dock", "mineffect", dockConf.MinimiseEffect)

	if err != nil {
		return err
	}

	err = macoshelpers.RunDefaultsWriteStringCmd("NSGlobalDomain", "AppleWindowTabbingMode", dockConf.PreferTabs)

	if err != nil {
		return err
	}

	err = macoshelpers.RunDefaultsWriteStringCmd("NSGlobalDomain", "AppleActionOnDoubleClick", dockConf.DoubleClickTitleTo)

	if err != nil {
		return err
	}

	err = macoshelpers.RunDefaultsWriteIntCmd("com.apple.dock", "minimize-to-application", boolToInt(dockConf.MinimiseToAppIcon))

	if err != nil {
		return err
	}

	err = macoshelpers.RunDefaultsWriteIntCmd("com.apple.dock", "launchanim", boolToInt(dockConf.AnimateOpening))

	if err != nil {
		return err
	}

	err = macoshelpers.RunDefaultsWriteIntCmd("com.apple.dock", "autohide", boolToInt(dockConf.AutoHide))

	if err != nil {
		return err
	}

	err = macoshelpers.RunDefaultsWriteIntCmd("com.apple.dock", "show-process-indicators", boolToInt(dockConf.ShowOpenIndicator))

	if err != nil {
		return err
	}

	err = macoshelpers.RunDefaultsWriteIntCmd("com.apple.dock", "show-recents", boolToInt(dockConf.ShowRecent))

	if err != nil {
		return err
	}

	err = setDockApps(dockConf.Apps)

	if err != nil {
		return err
	}

	return killDock()
}

func setDockApps(apps []string) error {
	macoshelpers.RunDefaultsDeleteCmd("com.apple.dock", "persistent-apps")

	for _, app := range apps {
		appPath, err := getAppPath(app)

		if err != nil {
			return err
		}

		xmlApp := fmt.Sprintf("<dict><key>tile-data</key><dict><key>file-data</key><dict><key>_CFURLString</key><string>%s</string><key>_CFURLStringType</key><integer>0</integer></dict></dict></dict>", appPath)
		return macoshelpers.RunDefaultsArrayAddCmd("com.apple.dock", "persistent-apps", xmlApp)
	}

	return nil
}

func getAppPath(appName string) (string, error) {
	out, err := exec.Command("mdfind", "-name", appName).Output()

	if err != nil {
		return "", err
	}

	lines := bytes.Split(out, []byte("\n"))

	for _, line := range lines {
		text := string(line)

		if strings.HasPrefix(text, "/") && strings.HasSuffix(text, ".app") {
			return text, nil
		}
	}

	return "", fmt.Errorf("Could not find path for app %s", appName)
}

func killDock() error {
	return exec.Command("killall", "Dock").Run()
}

func boolToInt(booleanValue bool) int {
	if booleanValue {
		return 1
	}

	return 0
}
