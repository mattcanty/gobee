# Gobee

:computer: :honeybee:

[![Actions Status](https://github.com/mattcanty/gobee/workflows/pre-commit/badge.svg)](https://github.com/mattcanty/gobee/actions)

Gobee is intended to alleviate the strain of finding that elusive
Shell/Bash/PowerShell (delete appropriate) command from some random website.

Intended to provide a single configuration file as a definition for one
or many systems you manage.

*Disclaimer: this is one of those nutty ideas that will probably go cold. But
if you want to contribute please feel free to do so!*

## Run like

```shell
go install
gobee -f config-examples/mac-os-catalina.yaml
```

## Config Like

### Mac

* Catalina: [mac-os-catalina.yaml](config-examples/mac-os-catalina.yaml)

Could be as simple as having this one config file to make
all your dreams come true.

```yaml
macOS:
  dock:
    apps:
      - Firefox
      - Visual Studio Code
      - Spotify
      - Slack
      - Terminal
    tile-size: 37
    magnification:
      enabled: true
      size: 49
    position: bottom
    minimise-effect: genie
    prefer-tabs: fullscreen
    double-click-title-to: Maximize
    minimise-to-app-icon: true
    animate-opening: true
    auto-hide: true
    show-open-indicator: true
    show-recent: false
```
