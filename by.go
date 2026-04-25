package gopium

type By struct {
	Using string
	Value string
}

func Locator(using, value string) By {
	return By{Using: using, Value: value}
}

func ID(value string) By {
	return Locator("id", value)
}

func XPath(value string) By {
	return Locator("xpath", value)
}

func AccessibilityID(value string) By {
	return Locator("accessibility id", value)
}

func ClassName(value string) By {
	return Locator("class name", value)
}

func Name(value string) By {
	return Locator("name", value)
}

func TagName(value string) By {
	return Locator("tag name", value)
}

func CSSSelector(value string) By {
	return Locator("css selector", value)
}

func LinkText(value string) By {
	return Locator("link text", value)
}

func PartialLinkText(value string) By {
	return Locator("partial link text", value)
}

func AndroidUIAutomator(value string) By {
	return Locator("-android uiautomator", value)
}

func IOSClassChain(value string) By {
	return Locator("-ios class chain", value)
}

func IOSPredicate(value string) By {
	return Locator("-ios predicate string", value)
}

func Image(value string) By {
	return Locator("-image", value)
}

func Custom(using, value string) By {
	return Locator(using, value)
}

func (b By) payload() map[string]any {
	return map[string]any{
		"using": b.Using,
		"value": b.Value,
	}
}
