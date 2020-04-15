package model

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
)

var (
	// ErrInvalidType is returned when the client sent wrong template type.
	ErrInvalidType = errors.New("invalid template type")
	// ErrJSONUnmarshal is returned when JSON message can't be unmarshal.
	ErrJSONUnmarshal = errors.New("JSON unmarshal fail")
)

/**
 * description: 执行模板并返回输出
 */
func ExecTmpl(tasktmpl TaskTmpl, varMap string) (content string, e error) {
	t, e := template.New(tasktmpl.Name).Parse(tasktmpl.Content)
	var b bytes.Buffer

	m := make(map[string]string)
	e = json.Unmarshal([]byte(varMap), &m)
	if e != nil {
		return content, ErrJSONUnmarshal
	}

	if e != nil {
		return content, fmt.Errorf("json unmarshal fail, %s", e.Error())
	}

	e = t.Execute(&b, m)
	if e != nil {
		return content, fmt.Errorf("execute template fail, %s", e.Error())
	}
	return b.String(), nil
}

// VarType defines variable's type.
type VarType int

const (
	// String means variable is a string.
	String VarType = iota + 1
	// Int means variable is an integer.
	Int
	// Float means variable is a float.
	Float
)

// VarFormType defines variable's type in a submitted form.
const (
	// FormString means variable is an input item of string
	FormString = iota + 1
	// FormNumber means variable is an input item of number
	FormNumber
	// Percentage means variable is an input item of percentage range
	FormPercentage
	// FormSelect means variable is an select item
	FormSelect
	// FormDatetime means variable is an datetime picker item
	FormDatetime
)
