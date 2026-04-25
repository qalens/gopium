# gopium

`gopium` is a Go Appium client for `github.com/qalens/gopium`.

It is designed around modern Appium and WebDriver behavior:

- W3C WebDriver first
- `appium:` extension capabilities
- optional `appium:options` grouping
- compatibility with older JSONWP/MJSONWP-style session and element payloads

The API intentionally follows Appium client conventions from the Java client, but uses Go semantics:

- constructor functions instead of overloaded classes
- fluent option builders for capabilities
- explicit `context.Context` on network calls
- small, composable structs instead of inheritance-heavy APIs

## Install

```bash
go get github.com/qalens/gopium
```

## Quick Start

```go
package main

import (
	"context"
	"log"

	"github.com/qalens/gopium"
)

func main() {
	ctx := context.Background()

	options := gopium.NewUiAutomator2Options().
		SetDeviceName("Android Emulator").
		SetUDID("emulator-5554").
		SetApp("/path/to/app.apk")

	driver, err := gopium.NewDriver(ctx, "http://127.0.0.1:4723", options)
	if err != nil {
		log.Fatal(err)
	}
	defer driver.Quit(ctx)

	login, err := driver.FindElement(ctx, gopium.AccessibilityID("login"))
	if err != nil {
		log.Fatal(err)
	}

	if err := login.Click(ctx); err != nil {
		log.Fatal(err)
	}
}
```

## Options

Generic options:

```go
options := gopium.NewBaseOptions().
	SetPlatformName("Android").
	SetAutomationName("UiAutomator2").
	SetCapability("webSocketUrl", true).
	SetAppiumCapability("deviceName", "Pixel 9")
```

Grouped Appium options:

```go
options := gopium.NewBaseOptions().
	WithAppiumOptions().
	SetPlatformName("iOS").
	SetAutomationName("XCUITest").
	SetDeviceName("iPhone 16")
```

Driver-specific helpers:

- `gopium.NewUiAutomator2Options()`
- `gopium.NewXCUITestOptions()`
- `gopium.NewAndroidOptions()`
- `gopium.NewIOSOptions()`
- `gopium.NewAndroidDriver()`
- `gopium.NewIOSDriver()`

## Device Farm Providers

The package includes a provider abstraction so the same test setup can target local Appium, BrowserStack, Sauce Labs, TestingBot, or other clouds with minimal changes.

```go
provider := gopium.NewBrowserStackProvider(
	os.Getenv("BROWSERSTACK_USERNAME"),
	os.Getenv("BROWSERSTACK_ACCESS_KEY"),
).SetVendorOptions(map[string]any{
	"projectName": "Checkout",
	"buildName":   "build-142",
	"sessionName": "login flow",
})

options := gopium.NewXCUITestOptions().
	SetDeviceName("iPhone 16").
	SetPlatformVersion("18").
	SetApp("bs://<app-id>")

driver, err := gopium.NewProviderDriver(ctx, provider, options)
```

Built-in providers:

- `NewBrowserStackProvider`
- `NewSauceLabsProvider`
- `NewLambdaTestProvider`
- `NewTestingBotProvider`
- `NewCloudProvider` for custom or internal grids

Provider capabilities follow the vendor’s own namespaced convention:

- BrowserStack: `bstack:options`
- Sauce Labs: `sauce:options`
- TestingBot: `tb:options`
- LambdaTest: `LT:Options`

The same provider can also add Appium-style cloud extension objects like `browserstack:appiumOptions` or `sauce:appiumOptions` when you want one API for selecting Appium server or driver/plugin versions across clouds.

## Supported Client Surface

- session creation and teardown
- session status and capabilities
- page source, title, URL, screenshot
- navigation commands
- element lookup and nested element lookup
- click, clear, send keys, text, attribute, property, rect
- Appium settings
- timeouts
- script execution, including `mobile:` commands via `ExecuteMobile`
- alerts
- contexts
- W3C actions
- cookies
- orientation and rotation
- app lifecycle commands
- app install/remove/activate/terminate/state
- clipboard get/set helpers
- file push/pull helpers
- screen recording start/stop
- device helpers like current activity/package, keyboard visibility, notifications, key codes
- Appium event logging helpers
- generic `SessionCommand`, `AppiumCommand`, and `ExecuteCDP` escape hatches for commands not wrapped yet

## Protocol Compatibility

New session requests include both:

- W3C `capabilities`
- legacy `desiredCapabilities`

Response parsing supports:

- W3C session responses
- legacy session responses
- W3C element ids
- legacy `ELEMENT` ids

This makes the client work cleanly with current Appium servers while still tolerating older protocol shapes where needed.
