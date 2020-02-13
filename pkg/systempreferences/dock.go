package systempreferences

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"

	"github.com/mattcanty/gobee/pkg/config"
)

// ConfigureDock - Configures the Dock on MacOS
func ConfigureDock(dockConf config.Dock) {
	dockNamespace := "com.apple.dock"

	defaultsWriteInt(dockNamespace, "tilesize", dockConf.TileSize)
	defaultsWriteInt(dockNamespace, "magnification", boolToInt(dockConf.Magnification.Enabled))
	defaultsWriteInt(dockNamespace, "largesize", dockConf.Magnification.Size)
	defaultsWriteString(dockNamespace, "orientation", dockConf.Position)
	defaultsWriteString(dockNamespace, "mineffect", dockConf.MinimiseEffect)
	defaultsWriteString("NSGlobalDomain", "AppleWindowTabbingMode", dockConf.PreferTabs)
	defaultsWriteString("NSGlobalDomain", "AppleActionOnDoubleClick", dockConf.DoubleClickTitleTo)
	defaultsWriteInt(dockNamespace, "minimize-to-application", boolToInt(dockConf.MinimiseToAppIcon))
	defaultsWriteInt(dockNamespace, "launchanim", boolToInt(dockConf.AnimateOpening))
	defaultsWriteInt(dockNamespace, "autohide", boolToInt(dockConf.AutoHide))
	defaultsWriteInt(dockNamespace, "show-process-indicators", boolToInt(dockConf.ShowOpenIndicator))
	defaultsWriteInt(dockNamespace, "show-recents", boolToInt(dockConf.ShowRecent))

	setDockApps(dockConf.Apps)

	killDock()
}

func setDockApps(apps []string) {
	defaultsDelete("com.apple.dock", "persistent-apps")

	for _, app := range apps {
		appPath, err := getAppPath(app)

		if err != nil {
			log.Fatalf("getAppPath failed for %s\n", err)
		}

		xmlApp := fmt.Sprintf("<dict><key>tile-data</key><dict><key>file-data</key><dict><key>_CFURLString</key><string>%s</string><key>_CFURLStringType</key><integer>0</integer></dict></dict></dict>", appPath)
		defaultsArrayAdd("com.apple.dock", "persistent-apps", xmlApp)
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

func defaultsDelete(location string, property string) {
	cmd := exec.Command("defaults", "delete", location, property)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

func defaultsWriteString(location string, property string, value string) {
	cmd := exec.Command("defaults", "write", location, property, "-string", value)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

func defaultsWriteInt(location string, property string, value int) {
	cmd := exec.Command("defaults", "write", location, property, "-int", strconv.Itoa(value))

	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

func defaultsArrayAdd(location string, property string, value string) {
	cmd := exec.Command("defaults", "write", location, property, "-array-add", value)

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
