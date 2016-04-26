package urlqp

import (
	"fmt"
	"net/url"
	"strconv"
)

// We want to be able to process some query parameters in order. This API
// treats the order of query parameters significantly, using them as a set of
// operations.
func ExampleParse() {
	// Add one, multiply by two, add three.
	str := `add=1&multiply=2&add=3`

	// We should get `5` as the result.
	var expected int64 = 5

	// Here we use the default `net/url` query parser. It parses the parameters
	// into a `map[string][]string`, which is perfectly fine for most uses, but
	// not ours.
	m, _ := url.ParseQuery(str)
	{
		var i int64

		// The order that we range over this map is intentionally undefined. That
		// means we'll end up with one of two results:
		//
		// > add 1, add 3, multiply 2 = 8
		// or
		// > multiply 2, add 1, add 3 = 4
		//
		// This is bad for this particular use case.
		for k, l := range m {
			for _, v := range l {
				n, _ := strconv.ParseInt(v, 10, 32)

				switch k {
				case "add":
					i += n
				case "multiply":
					i *= n
				}
			}
		}

		// This is going to print `false`.
		fmt.Printf("equal: %v\n", i == expected)
	}

	// This is using the `urlqp` parser. It preserves the order of the
	// parameters as it just returns a list of key/value pairs.
	q, _ := Parse(str)
	{
		var i int64

		// Iterating through this list happens in a defined order, which is the
		// order the parameters were specified in the string above.
		for _, p := range q {
			k, v := p[0], p[1]

			n, _ := strconv.ParseInt(v, 10, 32)

			switch k {
			case "add":
				i += n
			case "multiply":
				i *= n
			}
		}

		// This will print `true`
		fmt.Printf("equal: %v\n", i == expected)
	}

	// Output:
	//
	// equal: false
	// equal: true
}
