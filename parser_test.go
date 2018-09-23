package simpleformatter

/**
    This file is part of Simple Formatter.

    Simple Formatter is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    Simple Formatter is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
	along with Simple Formatter.  If not, see <https://www.gnu.org/licenses/>.
**/
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

func TestParserErrorSectionNameNotKnown(t *testing.T) {
	defaultSectionDefinition := SectionDefinition{}
	sections := []*SectionDefinition{&SectionDefinition{SectionName: "END", FieldNames: names}}

	err := NewParsedFormatWithSections("hello BEGIN{User {user}} this", &defaultSectionDefinition, sections...)
	if err == nil {
		t.Error("no error returned")
		return
	}

	if _, ok := err.(ErrSectionNameNotKnown); !ok {
		t.Error("unexpected error returned", err)
		return
	}

	if err.Error() != "the section name BEGIN is not known" {
		t.Error("unexpected error string")
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

func TestParserIntFieldInSecion(t *testing.T) {

	defaultSectionDefinition := SectionDefinition{}
	sections := []*SectionDefinition{&SectionDefinition{SectionName: "BEGIN", FieldNames: names}}

	err := NewParsedFormatWithSections("hello BEGIN{User {user}} this", &defaultSectionDefinition, sections...)
	if err != nil {
		t.Error(err)
		return
	}

	v := sections[0].GetFormattedValue([]interface{}{123})
	assert.Equal(t, "User 123", v)
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
