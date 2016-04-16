package urlqp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	a := assert.New(t)

	if v, err := Parse(``); a.NoError(err) {
		a.Nil(v)
	}

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

	{
		_, err := Parse(`%`)
		a.Error(err)
	}

	{
		_, err := Parse(`a=%`)
		a.Error(err)
	}
}

func TestString(t *testing.T) {
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

func TestGet(t *testing.T) {
	a := assert.New(t)

	v, err := Parse(`a=b&c=d&e=f&e=g`)
	a.NoError(err)

	a.Equal("b", v.Get("a"))
	a.Equal("d", v.Get("c"))
	a.Equal("f", v.Get("e"))
	a.Empty(v.Get("x"))
}

func TestAll(t *testing.T) {
	a := assert.New(t)

	v, err := Parse(`a=b&c=d&e=f&e=g`)
	a.NoError(err)

	a.Equal([]string{"b"}, v.All("a"))
	a.Equal([]string{"d"}, v.All("c"))
	a.Equal([]string{"f", "g"}, v.All("e"))
	a.Nil(v.All("x"))
}

func TestFilter(t *testing.T) {
	a := assert.New(t)

	v, err := Parse(`a=b&c=d&e=f&e=g`)
	a.NoError(err)

	a.Equal(Values{
		Pair{"a", "b"},
	}, v.Filter("a"))

	a.Equal(Values{
		Pair{"a", "b"},
		Pair{"c", "d"},
	}, v.Filter("a", "c"))

	a.Equal(Values{
		Pair{"e", "f"},
		Pair{"e", "g"},
	}, v.Filter("e"))
}
