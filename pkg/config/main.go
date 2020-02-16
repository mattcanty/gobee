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
	TileSize      int      `yaml:"tileSize"`
	Magnification struct {
		Enabled bool `yaml:"enabled"`
		Size    int  `yaml:"size"`
	} `yaml:"magnification"`
	Position           string `yaml:"position"`
	MinimiseEffect     string `yaml:"minimiseEffect"`
	PreferTabs         string `yaml:"preferTabs"`
	DoubleClickTitleTo string `yaml:"doubleClickTitleTo"`
	MinimiseToAppIcon  bool   `yaml:"minimiseToAppIcon"`
	AnimateOpening     bool   `yaml:"animateOpening"`
	AutoHide           bool   `yaml:"autoHide"`
	ShowOpenIndicator  bool   `yaml:"showOpenIndicator"`
	ShowRecent         bool   `yaml:"showRecent"`
}
