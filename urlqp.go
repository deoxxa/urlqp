package urlqp

import (
	"net/url"
	"strings"
)

type Pair [2]string
type Values []Pair

func Parse(s string) (Values, error) {
	s = strings.TrimPrefix(s, "?")

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

func Serialize(v Values) string {
	a := make([]string, len(v))

	for i, p := range v {
		a[i] = url.QueryEscape(p[0]) + "=" + url.QueryEscape(p[1])
	}

	return strings.Join(a, "&")
}
