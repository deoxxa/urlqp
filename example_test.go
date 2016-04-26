package urlqp

import (
	"fmt"
	"net/url"
	"strconv"
)

func ExampleParse() {
	str := `add=1&multiply=2&add=3`

	var expected int64 = 5

	m, _ := url.ParseQuery(str)
	{
		var i int64
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

		fmt.Printf("equal: %v\n", i == expected)
	}

	q, _ := Parse(str)
	{
		var i int64
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

		fmt.Printf("equal: %v\n", i == expected)
	}

	// Output:
	//
	// equal: false
	// equal: true
}
