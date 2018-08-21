package tpl

import (
	"github.com/zcong1993/badge/libs"
	"html/template"
)

var normalizeFloat64 = libs.NormalizeFloat64

func add(a, b float64) float64 {
	return normalizeFloat64(a + b)
}

func ftbfunc1(a float64) float64 {
	return normalizeFloat64(a / 2 * 10)
}

func ftbfunc2(a float64) float64 {
	return normalizeFloat64((a - 24) * 10)
}

func ftbfunc3(a, b float64) float64 {
	return normalizeFloat64((a + b/2) * 10)
}

// TPL is struct of tpl engine
type TPL struct {
	DefaultTpl     *template.Template
	FlatTpl        *template.Template
	ForTheBadgeTpl *template.Template
}

// NewTPL is constructor for TPL
func NewTPL() *TPL {
	tpl := &TPL{}
	funcs := template.FuncMap{
		"add":              add,
		"ftbfunc1":         ftbfunc1,
		"ftbfunc2":         ftbfunc2,
		"ftbfunc3":         ftbfunc3,
		"normalizeFloat64": normalizeFloat64,
	}
	dft := template.Must(template.New("default").Funcs(funcs).Parse(DefaultTpl))
	flat := template.Must(template.New("flat").Funcs(funcs).Parse(FlatTpl))
	ftb := template.Must(template.New("for-the-badge").Funcs(funcs).Parse(ForTheBadgeTpl))

	tpl.DefaultTpl = dft
	tpl.FlatTpl = flat
	tpl.ForTheBadgeTpl = ftb

	return tpl
}
