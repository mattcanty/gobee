package systempreferences

import (
	"log"
	"os/exec"
	"strconv"

	"github.com/mattcanty/gobee/pkg/config"
)

// ConfigureDock - Configures the Dock on MacOS
func ConfigureDock(dockConf config.Dock) {
	defaultsWriteInt("com.apple.dock", "tilesize", dockConf.TileSize)
	defaultsWriteInt("com.apple.dock", "magnification", boolToInt(dockConf.Magnification.Enabled))
	defaultsWriteInt("com.apple.dock", "largesize", dockConf.Magnification.Size)
	defaultsWriteString("com.apple.dock", "orientation", dockConf.Position)
	defaultsWriteString("com.apple.dock", "mineffect", dockConf.MinimiseEffect)
	defaultsWriteString("NSGlobalDomain", "AppleWindowTabbingMode", dockConf.PreferTabs)
	defaultsWriteString("NSGlobalDomain", "AppleActionOnDoubleClick", dockConf.DoubleClickTitleTo)
	defaultsWriteInt("com.apple.dock", "minimize-to-application", boolToInt(dockConf.MinimiseToAppIcon))
	defaultsWriteInt("com.apple.dock", "launchanim", boolToInt(dockConf.AnimateOpening))
	defaultsWriteInt("com.apple.dock", "autohide", boolToInt(dockConf.AutoHide))
	defaultsWriteInt("com.apple.dock", "show-process-indicators", boolToInt(dockConf.ShowOpenIndicator))
	defaultsWriteInt("com.apple.dock", "show-recents", boolToInt(dockConf.ShowRecent))

	killDock()
}

func killDock() {
	cmd := exec.Command("killall", "Dock")

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

func boolToInt(b bool) int {
	var bitSetVar int
	if b {
		bitSetVar = 1
	}

	return bitSetVar
}
