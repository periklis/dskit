package flagext

import "net/url"

// URLValue is a url.URL that can be used as a flag.
type URLValue struct {
	*url.URL
}

// String implements flag.Value
func (v URLValue) String() string {
	if v.URL == nil {
		return ""
	}
	return v.URL.String()
}

// Set implements flag.Value
func (v *URLValue) Set(s string) error {
	u, err := url.Parse(s)
	if err != nil {
		return err
	}
	v.URL = u
	return nil
}
