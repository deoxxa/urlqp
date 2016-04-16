package urlqp // import "fknsrs.biz/p/urlqp"

import (
	"net/url"
	"strings"
)

type Pair [2]string
type Values []Pair

func (v Values) Get(k string) string {
	for _, p := range v {
		if p[0] == k {
			return p[1]
		}
	}

	return ""
}

func (v Values) All(k string) []string {
	var a []string

	for _, p := range v {
		if p[0] == k {
			a = append(a, p[1])
		}
	}

	return a
}

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

func (v Values) String() string {
	a := make([]string, len(v))

	for i, p := range v {
		a[i] = url.QueryEscape(p[0]) + "=" + url.QueryEscape(p[1])
	}

	return strings.Join(a, "&")
}

func Parse(s string) (Values, error) {
	s = strings.TrimPrefix(s, "?")

	if s == "" {
		return nil, nil
	}

	a := strings.Split(s, "&")

	r := make(Values, len(a))

	for i, p := range a {
		b := strings.SplitN(p, "=", 2)

		k, err := url.QueryUnescape(b[0])
		if err != nil {
			return nil, err
		}

		v := ""
		if len(b) > 1 {
			d, err := url.QueryUnescape(b[1])
			if err != nil {
				return nil, err
			}

			v = d
		}

		r[i] = Pair{k, v}
	}

	return r, nil
}
