package Crawler

import (
	"github.com/tebeka/selenium"
)

func SeleniumWebDriver() (selenium.WebDriver, error) {
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome", "Args": "--headless --start-maximized"})
	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		return nil, err
	}

	return driver, nil
}
