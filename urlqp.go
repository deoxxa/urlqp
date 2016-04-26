package urlqp // import "fknsrs.biz/p/urlqp"

import (
	"net/url"
	"strings"
)

// Pair represents a key/value pair.
type Pair [2]string

// Values represents all the key/value pairs in a set of query parameters.
type Values []Pair

// Get retrieves the value of the first query parameter named "k", or returns
// an empty string.
func (v Values) Get(k string) string {
	for _, p := range v {
		if p[0] == k {
			return p[1]
		}
	}

	return ""
}

// All retrieves the values of all the query parameters named "k".
func (v Values) All(k string) []string {
	var a []string

	for _, p := range v {
		if p[0] == k {
			a = append(a, p[1])
		}
	}

	return a
}

// Filter returns a new set of values, limited to only the specified keys.
func (v Values) Filter(keys ...string) Values {
	var a Values

	for _, p := range v {
		for _, k := range keys {
			if p[0] == k {
				a = append(a, p)
			}
		}
	}

	return a
}

// String serialises this set of query parameters into a format suitable for
// use in a URL.
func (v Values) String() string {
	a := make([]string, len(v))

	for i, p := range v {
		a[i] = url.QueryEscape(p[0]) + "=" + url.QueryEscape(p[1])
	}

	return strings.Join(a, "&")
}

// Parse tries to parse the given string into a set of values. If it fails, it
// will return an error.
func Parse(s string) (Values, error) {
	// We don't need the leading question mark, and it's safe to get rid of it.
	s = strings.TrimPrefix(s, "?")

	// A small optimisation - an empty string means an empty set of parameters.
	if s == "" {
		return nil, nil
	}

	// Split the string on `&`, which is the standard query parameter delimiter.
	// It really is this simple, since `&` has to be escaped if it's contained
	// in any values, so we can't end up splitting "too much" or anything like
	// that.
	a := strings.Split(s, "&")

	// Another small optimisation - allocate the `Values` slice with the exact
	// size we need ahead of time, to avoid copying during `append()` calls.
	r := make(Values, len(a))

	// Iterate through the query parameters one at a time, perserving the order.
	for i, p := range a {
		// Split this query parameter into two parts, at the first `=`. In the
		// case of a construct like `?a&b`, the length of `b` will be 1. It can't
		// be zero, since even an empty string will split to a slice with one
		// element.
		b := strings.SplitN(p, "=", 2)

		// Unescape the query parameter key. If this fails, we return an error.
		k, err := url.QueryUnescape(b[0])
		if err != nil {
			return nil, err
		}

		// Set a default for the value. Empty seems reasonable.
		v := ""

		// If `b` has more than one element, that means the second one will be the
		// parameter value, so we should grab it.
		if len(b) > 1 {
			// Now unescape the query parameter value. If this fails, return an
			// error.
			d, err := url.QueryUnescape(b[1])
			if err != nil {
				return nil, err
			}

			// Nothing failed, so we can keep the parameter value.
			v = d
		}

		// Add this parameter pair to the set of values.
		r[i] = Pair{k, v}
	}

	// Return the values and nil, signalling no error.
	return r, nil
}
