package youtube

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/tebeka/selenium"
	"math/rand/v2"
	"time"
	"walle/pkg/jscript"
	"walle/pkg/null"
)

type YouTube struct {
	webDriver selenium.WebDriver
	steps     int
	urls      []string
}

func New(webDriver selenium.WebDriver, steps int, urls []string) *YouTube {
	return &YouTube{
		webDriver: webDriver,
		steps:     steps,
		urls:      urls,
	}
}

func (y *YouTube) processVideo(url string, first bool) (null.Null[string], error) {
	if err := y.webDriver.Get(url); err != nil {
		return null.Null[string]{}, errors.Wrap(err, "web driver get article fail")
	}

	if first {
		fmt.Println("first")
		if err := y.webDriver.WaitWithTimeout(func(webDriver selenium.WebDriver) (bool, error) {
			webElement, _ := y.webDriver.FindElement(selenium.ByCSSSelector, ".yt-spec-button-shape-next.yt-spec-button-shape-next--filled.yt-spec-button-shape-next--mono.yt-spec-button-shape-next--size-m.yt-spec-button-shape-next--enable-backdrop-filter-experiment")

			if webElement != nil {
				return webElement.IsDisplayed()
			}

			return false, nil
		}, 10*time.Second); err != nil {
			return null.Null[string]{}, errors.Wrap(err, "wait with timeout fail")
		}

		webElements, err := y.webDriver.FindElements(selenium.ByCSSSelector, ".yt-spec-button-shape-next.yt-spec-button-shape-next--filled.yt-spec-button-shape-next--mono.yt-spec-button-shape-next--size-m.yt-spec-button-shape-next--enable-backdrop-filter-experiment")
		if err != nil {
			return null.Null[string]{}, errors.Wrap(err, "find element by class name fail")
		}

		if _, err := y.webDriver.ExecuteScript(jscript.ClickOnFirstArgument, []interface{}{webElements[2]}); err != nil {
			return null.Null[string]{}, errors.Wrap(err, "execute click on first argument script fail")
		}
	}

	if err := y.webDriver.WaitWithTimeout(func(webDriver selenium.WebDriver) (bool, error) {
		webElement, _ := y.webDriver.FindElement(selenium.ByCSSSelector, ".ytp-play-button.ytp-button")

		if webElement != nil {
			return webElement.IsDisplayed()
		}

		return false, nil
	}, 10*time.Second); err != nil {
		return null.Null[string]{}, errors.Wrap(err, "wait with timeout fail")
	}

	webElement, err := y.webDriver.FindElement(selenium.ByCSSSelector, ".ytp-play-button.ytp-button")
	if err != nil {
		return null.Null[string]{}, errors.Wrap(err, "find element by class name fail")
	}

	if _, err := y.webDriver.ExecuteScript(jscript.ClickOnFirstArgument, []interface{}{webElement}); err != nil {
		return null.Null[string]{}, errors.Wrap(err, "execute click on first argument script fail")
	}

	fmt.Println("play")
	time.Sleep(time.Second * 15)

	webElement, err = y.webDriver.FindElement(selenium.ByCSSSelector, ".ytp-play-button.ytp-button")
	if err != nil {
		return null.Null[string]{}, errors.Wrap(err, "find element by class name fail")
	}

	if _, err := y.webDriver.ExecuteScript(jscript.ClickOnFirstArgument, []interface{}{webElement}); err != nil {
		return null.Null[string]{}, errors.Wrap(err, "execute click on first argument script fail")
	}

	webElements, err := webElement.FindElements(selenium.ByTagName, "ytd-compact-video-renderer")
	if err != nil {
		return null.Null[string]{}, errors.Wrap(err, "find elements by tag name fail")
	}

	fmt.Println(len(webElements))
	webElement, err = webElements[rand.IntN(len(webElements))].FindElement(selenium.ByTagName, "#thumbnail")
	if err != nil {
		return null.Null[string]{}, errors.Wrap(err, "find element by tag name fail")
	}

	fmt.Println(webElement.GetAttribute("href"))

	return null.Null[string]{}, nil
}

func (y *YouTube) Exec() error {
	url := *null.New(y.urls[rand.IntN(len(y.urls))])
	first := true
	for range y.steps {
		fmt.Println(url)
		newUrl, err := y.processVideo(*url.GetValue(), first)
		if err != nil {
			return err
		}

		first = false

		url = newUrl
	}

	return nil
}
