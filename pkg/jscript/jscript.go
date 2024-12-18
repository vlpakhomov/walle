package jscript

const (
	SmoothScrollUp        = `window.scrollTo({ top: 0, behavior: 'smooth' })`
	ScrollDown            = `window.scrollTo({ top: document.body.scrollHeight, behavior: 'smooth' });`
	ClickOnFirstArgument  = `arguments[0].click();`
	SetLanguages          = `Object.defineProperty(navigator, 'languages', { value: ['en-US', 'en'] });`
	SetNavigatorWebDriver = `Object.defineProperty(navigator, 'webdriver', {get: () => undefined});`
)
