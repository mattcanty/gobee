package config

// Config the configuration definition
type Config struct {
	MacOS `yaml:"macOS"`
}

// MacOS the configuration for Mac OS
type MacOS struct {
	Dock `yaml:"dock"`
}

// Dock the configuration for MacOS Dock
type Dock struct {
	Apps          []string `yaml:"apps"`
	TileSize      int      `yaml:"tile-size"`
	Magnification struct {
		Enabled bool `yaml:"enabled"`
		Size    int  `yaml:"size"`
	} `yaml:"magnification"`
	Position           string `yaml:"position"`
	MinimiseEffect     string `yaml:"minimise-effect"`
	PreferTabs         string `yaml:"prefer-tabs"`
	DoubleClickTitleTo string `yaml:"double-click-title-to"`
	MinimiseToAppIcon  bool   `yaml:"minimise-to-app-icon"`
	AnimateOpening     bool   `yaml:"animate-opening"`
	AutoHide           bool   `yaml:"auto-hide"`
	ShowOpenIndicator  bool   `yaml:"show-open-indicator"`
	ShowRecent         bool   `yaml:"show-recent"`
}
