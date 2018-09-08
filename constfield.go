package simpleformatter

type ConstantField struct {
	value string
}

func (field *ConstantField) GetValue(values []interface{}) string {
	return field.value
}
