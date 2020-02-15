package systempreferences

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/mattcanty/gobee/pkg/config"
	macoshelpers "github.com/mattcanty/gobee/pkg/macos/helpers"
)

// ConfigureDock - Configures the Dock on MacOS
func ConfigureDock(dockConf config.Dock) {
	dockNamespace := "com.apple.dock"

	macoshelpers.RunDefaultsWriteIntCmd(dockNamespace, "tilesize", dockConf.TileSize)
	macoshelpers.RunDefaultsWriteIntCmd(dockNamespace, "magnification", boolToInt(dockConf.Magnification.Enabled))
	macoshelpers.RunDefaultsWriteIntCmd(dockNamespace, "largesize", dockConf.Magnification.Size)
	macoshelpers.RunDefaultsWriteStringCmd(dockNamespace, "orientation", dockConf.Position)
	macoshelpers.RunDefaultsWriteStringCmd(dockNamespace, "mineffect", dockConf.MinimiseEffect)
	macoshelpers.RunDefaultsWriteStringCmd("NSGlobalDomain", "AppleWindowTabbingMode", dockConf.PreferTabs)
	macoshelpers.RunDefaultsWriteStringCmd("NSGlobalDomain", "AppleActionOnDoubleClick", dockConf.DoubleClickTitleTo)
	macoshelpers.RunDefaultsWriteIntCmd(dockNamespace, "minimize-to-application", boolToInt(dockConf.MinimiseToAppIcon))
	macoshelpers.RunDefaultsWriteIntCmd(dockNamespace, "launchanim", boolToInt(dockConf.AnimateOpening))
	macoshelpers.RunDefaultsWriteIntCmd(dockNamespace, "autohide", boolToInt(dockConf.AutoHide))
	macoshelpers.RunDefaultsWriteIntCmd(dockNamespace, "show-process-indicators", boolToInt(dockConf.ShowOpenIndicator))
	macoshelpers.RunDefaultsWriteIntCmd(dockNamespace, "show-recents", boolToInt(dockConf.ShowRecent))

	setDockApps(dockConf.Apps)

	killDock()
}

func setDockApps(apps []string) {
	macoshelpers.RunDefaultsDeleteCmd("com.apple.dock", "persistent-apps")

	for _, app := range apps {
		appPath, err := getAppPath(app)

		if err != nil {
			log.Fatalf("getAppPath failed for %s\n", err)
		}

		xmlApp := fmt.Sprintf("<dict><key>tile-data</key><dict><key>file-data</key><dict><key>_CFURLString</key><string>%s</string><key>_CFURLStringType</key><integer>0</integer></dict></dict></dict>", appPath)
		macoshelpers.RunDefaultsArrayAddCmd("com.apple.dock", "persistent-apps", xmlApp)
	}
}

func getAppPath(appName string) (string, error) {
	cmd := exec.Command("mdfind", "-name", appName)

	out, err := cmd.Output()

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

func killDock() {
	cmd := exec.Command("killall", "Dock")

	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

func boolToInt(b bool) int {
	var bitSetVar int
	if b {
		bitSetVar = 1
	}

	return bitSetVar
}
