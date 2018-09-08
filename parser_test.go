package simpleformatter

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var names = []string{"user", "other"}

func TestParserErrorOpenBrace(t *testing.T) {
	_, err := NewParsedFormat(names, "hello {")
	if err == nil {
		t.Error("no error returned")
	}
}

func TestParserErrorFieldNotFound(t *testing.T) {
	_, err := NewParsedFormat(names, "hello {notfound}")
	if err == nil {
		t.Error("no error returned")
	}
}

func TestParserErrorMoreThenOneOpenBrace(t *testing.T) {
	_, err := NewParsedFormat(names, "hello {user{other}}")
	if err == nil {
		t.Error("no error returned")
	}
}

func TestParserOneField(t *testing.T) {
	j, err := NewParsedFormat(names, "hello {user}")
	if err != nil {
		t.Error(err)
		return
	}
	v := j.GetFormattedValue([]interface{}{"world", time.Now()})
	assert.Equal(t, "hello world", v)
}

func BenchmarkParserOneField(b *testing.B) {
	j, err := NewParsedFormat(names, "hello {user} ")
	if err != nil {
		b.Error(err)
		return
	}
	t := time.Now()
	for i := 0; i < b.N; i++ {
		j.GetFormattedValue([]interface{}{"world", t})
	}
}

func TestParserDateField(t *testing.T) {
	j, err := NewParsedFormat(names, "hello {other:d} this")
	if err != nil {
		t.Error(err)
		return
	}
	tim := time.Now()
	v := j.GetFormattedValue([]interface{}{"world", tim})
	assert.Equal(t, fmt.Sprintf("hello %d this", tim.Day()), v)
}

func TestParserIntField(t *testing.T) {
	j, err := NewParsedFormat(names, "hello {user} this")
	if err != nil {
		t.Error(err)
		return
	}

	v := j.GetFormattedValue([]interface{}{123})
	assert.Equal(t, "hello 123 this", v)
}

func TestParserUnsupportedField(t *testing.T) {
	j, err := NewParsedFormat(names, "hello {user} this")
	if err != nil {
		t.Error(err)
		return
	}

	v := j.GetFormattedValue([]interface{}{1.2})
	assert.Equal(t, "hello default1.2 this", v)
}
