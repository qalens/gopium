package gopium

import (
	"fmt"
	"strings"
)

func NewNamedProvider(name string, credentials map[string]any) (*CloudProvider, error) {
	switch normalizeProviderName(name) {
	case "browserstack", "bsaccount", "bs":
		username, err := requiredCredential(credentials, "username", "userName", "user")
		if err != nil {
			return nil, fmt.Errorf("browserstack: %w", err)
		}
		accessKey, err := requiredCredential(credentials, "access_key", "accessKey", "key")
		if err != nil {
			return nil, fmt.Errorf("browserstack: %w", err)
		}
		return NewBrowserStackProvider(username, accessKey), nil
	case "saucelabs", "sauce labs", "sauce", "sauceaccount":
		username, err := requiredCredential(credentials, "username", "userName", "user")
		if err != nil {
			return nil, fmt.Errorf("saucelabs: %w", err)
		}
		accessKey, err := requiredCredential(credentials, "access_key", "accessKey", "key")
		if err != nil {
			return nil, fmt.Errorf("saucelabs: %w", err)
		}
		return NewSauceLabsProvider(username, accessKey), nil
	case "lambdatest", "lambda test", "lt", "lambdatestaccount":
		username, err := requiredCredential(credentials, "username", "userName", "user")
		if err != nil {
			return nil, fmt.Errorf("lambdatest: %w", err)
		}
		accessKey, err := requiredCredential(credentials, "access_key", "accessKey", "key")
		if err != nil {
			return nil, fmt.Errorf("lambdatest: %w", err)
		}
		return NewLambdaTestProvider(username, accessKey), nil
	case "testingbot", "testing bot", "tb", "testingbotaccount":
		key, err := requiredCredential(credentials, "key", "access_key", "accessKey")
		if err != nil {
			return nil, fmt.Errorf("testingbot: %w", err)
		}
		secret, err := requiredCredential(credentials, "secret", "access_key", "accessKey")
		if err != nil {
			return nil, fmt.Errorf("testingbot: %w", err)
		}
		return NewTestingBotProvider(key, secret), nil
	default:
		return nil, fmt.Errorf("unsupported provider %q", name)
	}
}

func normalizeProviderName(name string) string {
	name = strings.TrimSpace(strings.ToLower(name))
	replacer := strings.NewReplacer("-", "", "_", "", " ", "")
	return replacer.Replace(name)
}

func requiredCredential(credentials map[string]any, keys ...string) (string, error) {
	for _, key := range keys {
		if value, ok := credentials[key]; ok {
			formatted := strings.TrimSpace(fmt.Sprint(value))
			if formatted != "" {
				return formatted, nil
			}
		}
	}
	return "", fmt.Errorf("missing credential %s", strings.Join(keys, "/"))
}
