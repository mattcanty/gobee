package config

// Config the configuration definition
type Config struct {
	Apple `yaml:"apple"`
}

// Apple the configuration for Apple OS
type Apple struct {
	Dock `yaml:"dock"`
}

// Dock the configuration for Apple Dock
type Dock struct {
	TileSize      int `yaml:"tile-size"`
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
