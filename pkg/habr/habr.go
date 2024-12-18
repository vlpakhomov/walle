package habr

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/tebeka/selenium"
	"math/rand/v2"
	"time"
	"walle/pkg/jscript"
	"walle/pkg/null"
)

type Habr struct {
	webDriver selenium.WebDriver
	steps     int
	urls      []string
}

func New(webDriver selenium.WebDriver, steps int, urls []string) *Habr {
	return &Habr{
		webDriver: webDriver,
		steps:     steps,
		urls:      urls,
	}
}

func (h *Habr) processArticle(url string) (null.Null[string], error) {
	if err := h.webDriver.Get(url); err != nil {
		return null.Null[string]{}, errors.Wrap(err, "web driver get article fail")
	}

	// TODO: realistic time sleep
	time.Sleep(time.Second * 10)

	if _, err := h.webDriver.ExecuteScript(jscript.ScrollDown, nil); err != nil {
		return null.Null[string]{}, errors.Wrap(err, "execute scroll script fail")
	}

	// TODO: realistic time sleep
	time.Sleep(time.Second * 3)

	if _, err := h.webDriver.ExecuteScript(jscript.SmoothScrollUp, nil); err != nil {
		return null.Null[string]{}, errors.Wrap(err, "execute scroll script fail")
	}

	err := h.webDriver.WaitWithTimeout(func(webDriver selenium.WebDriver) (bool, error) {
		webElement, _ := h.webDriver.FindElement(selenium.ByCSSSelector, ".tm-tabs__tab-link.tm-tabs__tab-link.tm-tabs__tab-link_slim")

		if webElement != nil {
			return webElement.IsDisplayed()
		}

		return false, nil
	}, 10*time.Second)
	if err != nil {
		return null.Null[string]{}, errors.Wrap(err, "wait with timeout fail")
	}

	webElements, err := h.webDriver.FindElements(selenium.ByCSSSelector, ".tm-tabs__tab-link.tm-tabs__tab-link.tm-tabs__tab-link_slim")
	if err != nil {
		return null.Null[string]{}, errors.Wrap(err, "find element by class name fail")
	}

	if _, err := h.webDriver.ExecuteScript(jscript.ClickOnFirstArgument, []interface{}{webElements[1]}); err != nil {
		return null.Null[string]{}, errors.Wrap(err, "wait with timeout fail")
	}

	err = h.webDriver.WaitWithTimeout(func(webDriver selenium.WebDriver) (bool, error) {
		webElement, _ := h.webDriver.FindElement(selenium.ByClassName, "tm-article-card-list")

		if webElement != nil {
			return webElement.IsDisplayed()
		}

		return false, nil
	}, 10*time.Second)
	if err != nil {
		return null.Null[string]{}, errors.Wrap(err, "wait with timeout fail")
	}

	webElements, err = h.webDriver.FindElements(selenium.ByClassName, "tm-bordered-card")
	if err != nil {
		return null.Null[string]{}, errors.Wrap(err, "find element by class name fail")
	}

	webElement, err := webElements[rand.IntN(len(webElements))].FindElement(selenium.ByClassName, "tm-title__link")
	if err != nil {
		return null.Null[string]{}, errors.Wrap(err, "find element by class name fail")
	}

	newUrl, err := webElement.GetAttribute("href")
	if err != nil {
		return null.Null[string]{}, errors.Wrap(err, "find attribute by name fail")
	}

	return *null.New(newUrl), nil
}

func (h *Habr) Exec() error {
	url := *null.New(h.urls[rand.IntN(len(h.urls))])
	for range h.steps {
		fmt.Println(url)
		newUrl, err := h.processArticle(*url.GetValue())
		if err != nil {
			return err
		}

		url = newUrl
	}

	return nil
}
