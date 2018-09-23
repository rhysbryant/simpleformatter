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
	"strconv"
	"time"

	"github.com/vjeantet/jodaTime"
)

type mappedField struct {
	fieldIndex int
	arg        string
}

func (field *mappedField) GetValue(values []interface{}) string {
	switch t := values[field.fieldIndex].(type) {
	case string:
		return t
	case int:
		return strconv.Itoa(t)
	case time.Time:
		return jodaTime.Format(field.arg, t)
	default:
		return "default" + fmt.Sprintf("%+v", t)
	}
}
