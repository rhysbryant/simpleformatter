package simpleformatter

import (
	"bytes"
	"errors"
)

const (
	fieldNameStart    = '{'
	fieldNameEnd      = '}'
	fieldArgSeperator = ':'
)

type Field interface {
	GetValue(values []interface{}) string
}

type SimpleFormatParser struct {
	fields []Field
}

func (f *SimpleFormatParser) appendConstantField(value string) {

	f.fields = append(f.fields, &ConstantField{value})
}

func (f *SimpleFormatParser) appendMappedField(index int, arg string) {
	f.fields = append(f.fields, &mappedField{index, arg})
}

func getFieldIndex(fieldNames []string, findName string) int {
	for index, name := range fieldNames {
		if name == findName {
			return index
		}
	}
	return -1
}

func NewParsedFormat(fieldNames []string, format string) (*SimpleFormatParser, error) {
	parsedFieldFormat := SimpleFormatParser{}
	i := 0
	lastTokenPos := 0
	openFieldNameTagCount := 0
	argSeperatorIndex := 0

	for i < len(format) {
		if format[i] == fieldNameStart {
			if openFieldNameTagCount > 0 { //cannot have more then one open brace
				return nil, errors.New(string(fieldNameStart) + "unexpected here")
			}
			openFieldNameTagCount++

			value := format[lastTokenPos:i]
			parsedFieldFormat.appendConstantField(value)
			lastTokenPos = i

		} else if format[i] == fieldArgSeperator && openFieldNameTagCount > 0 {
			argSeperatorIndex = i
		} else if format[i] == fieldNameEnd && openFieldNameTagCount > 0 {
			var arg, value string
			var endIndex = lastTokenPos + 1
			if argSeperatorIndex > 0 {
				endIndex = argSeperatorIndex
				value = format[lastTokenPos+1 : endIndex]
				arg = format[endIndex+1 : i]
			} else {
				value = format[lastTokenPos+1 : i]
			}

			if index := getFieldIndex(fieldNames, value); index == -1 {
				return nil, errors.New("field " + value + " not find")
			} else {
				parsedFieldFormat.appendMappedField(index, arg)
			}

			lastTokenPos = i + 1
			openFieldNameTagCount--
			argSeperatorIndex = 0
		}
		i++
	}

	if openFieldNameTagCount > 0 { //woops the user didn't close the brace
		return nil, errors.New("missing closing " + string(fieldNameEnd))
	} else if lastTokenPos != i {
		value := format[lastTokenPos:i]
		parsedFieldFormat.appendConstantField(value)
	}
	return &parsedFieldFormat, nil
}

func (f *SimpleFormatParser) GetFormattedValue(values []interface{}) string {
	var buffer bytes.Buffer
	for _, field := range f.fields {
		buffer.WriteString(field.GetValue(values))
	}
	return buffer.String()
}
