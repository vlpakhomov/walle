package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"os"
	"walle/config"
	"walle/internal/app"
	"walle/pkg/jscript"
	"walle/pkg/youtube"
)

func main() {
	cfg := config.Load("./config/config.json")

	logger := zerolog.New(os.Stdout)

	service, err := selenium.NewChromeDriverService("chromedriver", 4444)
	if err != nil {
		logger.Err(err).Msg("create headless chrome instance fail")
	}
	defer service.Stop()

	caps := selenium.Capabilities{}

	caps.AddChrome(chrome.Capabilities{
		Args: []string{
			//"--ignore-certificate-errors",
			//"--ignore-ssl-errors",
			//"--ignore-certificate-errors-spki-list",
			//"--disable-gpu",
			"--disable-blink-features=AutomationControlled",
			"--start-maximized",
			fmt.Sprintf("--proxy-server=socks5://%v:%v", cfg.Proxy.Host, cfg.Proxy.Port),
			fmt.Sprintf("--user-agent=%v", cfg.UserAgent),
		}})

	webDriver, err := selenium.NewRemote(caps, "")
	if err != nil {
		logger.Err(err).Msg("create browser client fail")
	}

	if _, err = webDriver.ExecuteScript(jscript.SetNavigatorWebDriver, nil); err != nil {
		logger.Err(err).Msg("execute set navigator webdriver script fail")
	}

	if _, err = webDriver.ExecuteScript(jscript.SetLanguages, nil); err != nil {
		logger.Err(err).Msg("execute set languages script fail")
	}

	var deps []app.Dep

	//habr := habr.New(webDriver, 5, cfg.Deps.Habr.URLs)
	youtube := youtube.New(webDriver, 5, cfg.Deps.YouTube.URLs)
	deps = append(deps, youtube)

	var walle app.App
	walle = app.New(&cfg, &logger, webDriver, deps)

	walle.Start()
	defer walle.Stop()
}
