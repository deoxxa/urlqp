package urlqp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	a := assert.New(t)

	if v, err := Parse(`a=b&c=d&e=f`); a.NoError(err) {
		a.Equal(Values{
			Pair{"a", "b"},
			Pair{"c", "d"},
			Pair{"e", "f"},
		}, v)
	}

	if v, err := Parse(`a=b&a=c`); a.NoError(err) {
		a.Equal(Values{
			Pair{"a", "b"},
			Pair{"a", "c"},
		}, v)
	}

	if v, err := Parse(`a=b&a=b`); a.NoError(err) {
		a.Equal(Values{
			Pair{"a", "b"},
			Pair{"a", "b"},
		}, v)
	}

	if v, err := Parse(`?a=b&a=b`); a.NoError(err) {
		a.Equal(Values{
			Pair{"a", "b"},
			Pair{"a", "b"},
		}, v)
	}

	if v, err := Parse(`?x=%20&%20=x`); a.NoError(err) {
		a.Equal(Values{
			Pair{"x", " "},
			Pair{" ", "x"},
		}, v)
	}

	if v, err := Parse(`?x=+&+=x`); a.NoError(err) {
		a.Equal(Values{
			Pair{"x", " "},
			Pair{" ", "x"},
		}, v)
	}

	if v, err := Parse(`?x=%2c&%2c=x`); a.NoError(err) {
		a.Equal(Values{
			Pair{"x", ","},
			Pair{",", "x"},
		}, v)
	}
}

func TestSerialize(t *testing.T) {
	a := assert.New(t)

	a.Equal(`a=b&c=d`, (Values{
		Pair{"a", "b"},
		Pair{"c", "d"},
	}).String())

	a.Equal(`a=b&a=b`, (Values{
		Pair{"a", "b"},
		Pair{"a", "b"},
	}).String())

	a.Equal(`a=+&+=a`, (Values{
		Pair{"a", " "},
		Pair{" ", "a"},
	}).String())

	a.Equal(`a=%2C&%2C=a`, (Values{
		Pair{"a", ","},
		Pair{",", "a"},
	}).String())
}
