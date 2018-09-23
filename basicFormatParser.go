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
	"bytes"
	"errors"
	"regexp"
	"strings"
)

const (
	fieldNameStart    = '{'
	fieldNameEnd      = '}'
	fieldArgSeperator = ':'
)

var (
	//ErrExpectedClosingBrace returned if a block {...} is unclosed
	ErrExpectedClosingBrace = errors.New("expected }")
	validFieldsectionName   = regexp.MustCompile("^\\w+$")
)

//ErrSectionNameNotKnown returned when the section name can not be found during parsing
type ErrSectionNameNotKnown struct {
	name string
}

func (err ErrSectionNameNotKnown) Error() string {
	return "the section name " + err.name + " is not known"
}

type field interface {
	GetValue(values []interface{}) string
}

//SimpleFormatParser stores the compiled format
type SimpleFormatParser struct {
	fields []field
}

func (f *SimpleFormatParser) appendConstantField(value string) {

	f.fields = append(f.fields, &constantField{value})
}

func (f *SimpleFormatParser) appendMappedField(index int, arg string) {
	f.fields = append(f.fields, &mappedField{index, arg})
}

func findClosingBrace(str string) (int, error) {
	openCount := 1
	for i := 0; i < len(str); i++ {
		if str[i] == fieldNameStart {
			openCount++
		} else if str[i] == fieldNameEnd {
			openCount--
		}

		if openCount == 0 {
			return i, nil
		}
	}

	return 0, ErrExpectedClosingBrace

}

func getFields(format string, currentSection *SectionDefinition, sections sectionDefinitions) error {
	lastSpacePos := 0
	i := 0

	lastTokenPos := 0

	for i < len(format) {
		switch format[i] {
		case ' ':
			lastSpacePos = i + 1
		case fieldNameStart:
			sectionName := format[lastSpacePos:i]
			secionIndex := -1
			var endOfConstant = i
			if sections != nil && validFieldsectionName.MatchString(sectionName) {
				secionIndex = sections.getIndexByName(sectionName)
				if secionIndex == -1 {
					return ErrSectionNameNotKnown{sectionName}
				}
				endOfConstant = lastSpacePos
			}
			value := format[lastTokenPos:endOfConstant]
			currentSection.appendConstantField(value)
			lastTokenPos = i

			blockEnd, err := findClosingBrace(format[i+1:])
			if err != nil {
				return err
			}
			innerBlockText := format[i+1 : i+blockEnd+1]
			lastTokenPos = i + blockEnd + 2

			if secionIndex == -1 {
				value := innerBlockText
				arg := ""

				parts := strings.SplitN(value, ":", 2)
				if len(parts) == 2 {
					value = parts[0]
					arg = parts[1]
				}

				var index int
				if index = currentSection.getFieldIndexByName(value); index == -1 {
					return errors.New("field " + value + " not find")
				}
				currentSection.appendMappedField(index, arg)

			} else {
				if err := getFields(innerBlockText, sections[secionIndex], nil); err != nil {
					return err
				}
			}
			i = lastTokenPos - 1
		}
		i++
	}

	if lastTokenPos != i {
		value := format[lastTokenPos:i]
		currentSection.appendConstantField(value)
	}

	return nil
}

//NewParsedFormat parse non section baeed format string for example {user.name}
func NewParsedFormat(fieldNames []string, format string) (*SimpleFormatParser, error) {
	defaultSection := SectionDefinition{FieldNames: fieldNames}

	if err := getFields(format, &defaultSection, nil); err != nil {
		return nil, err
	}

	return &defaultSection.SimpleFormatParser, nil
}

//NewParsedFormatWithSections parse a format string that includes sections for example userdetails{ {user.name} }
func NewParsedFormatWithSections(format string, defaultSection *SectionDefinition, sections ...*SectionDefinition) error {

	if err := getFields(format, defaultSection, sections); err != nil {
		return err
	}

	return nil
}

//GetFormattedValue returns the formatted string for the given value
func (f *SimpleFormatParser) GetFormattedValue(values []interface{}) string {
	var buffer bytes.Buffer
	for _, field := range f.fields {
		buffer.WriteString(field.GetValue(values))
	}
	return buffer.String()
}
