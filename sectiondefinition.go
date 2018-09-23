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
const (
	defaultSectionName = "Default"
)

//SectionDefinition holds the definition of a named section and the fields valid within it
type SectionDefinition struct {
	FieldNames  []string
	SectionName string
	SimpleFormatParser
}

type sectionDefinitions []*SectionDefinition

func (s sectionDefinitions) getIndexByName(name string) int {
	for i := 0; i < len(s); i++ {
		if s[i].SectionName == name {
			return i
		}
	}
	return -1
}

func (s SectionDefinition) getFieldIndexByName(name string) int {
	for i := 0; i < len(s.FieldNames); i++ {
		if s.FieldNames[i] == name {
			return i
		}
	}
	return -1
}
