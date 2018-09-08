package simpleformatter

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
