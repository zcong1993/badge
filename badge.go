package badge

import (
	"bytes"
	"github.com/zcong1993/badge/libs"
	"github.com/zcong1993/badge/tpl"
	"strings"
)

const (
	// DEFAULT is default style
	DEFAULT = iota + 1
	// FLAT is flat style
	FLAT
	// FOR THE BADGE
	FOR_THE_BADGE
)

// Input is Badgen input struct
type Input struct {
	Subject string
	Status  string
	Color   string
	Style   int
}

// TplData is tpl data we use in template exec
type TplData struct {
	Width       float64
	SbRectWidth float64
	Color       string
	SbTextWidth float64
	Subject     string
	Status      string
	StTextWidth float64
	StRectWidth float64
}

var templ = tpl.NewTPL()

// Badgen is the api we expose
func Badgen(input Input) ([]byte, error) {
	if input.Style == FOR_THE_BADGE {
		input.Subject = strings.ToUpper(input.Subject)
		input.Status = strings.ToUpper(input.Status)
	}
	sbTextWidth := libs.CacWith(input.Subject)
	stTextWidth := libs.CacWith(input.Status)
	sbRectWidth := sbTextWidth + 10.2
	stRectWidth := stTextWidth + 10
	width := sbRectWidth + stRectWidth
	c, ok := libs.COLORS[input.Color]
	if !ok {
		c = input.Color
	}

	data := TplData{
		Width:       width,
		SbRectWidth: sbRectWidth,
		Color:       c,
		SbTextWidth: sbTextWidth,
		Subject:     input.Subject,
		Status:      input.Status,
		StTextWidth: stTextWidth,
		StRectWidth: stRectWidth,
	}

	var bf bytes.Buffer
	t := templ.DefaultTpl
	if input.Style == FLAT {
		t = templ.FlatTpl
	}
	if input.Style == FOR_THE_BADGE {
		data.SbRectWidth += 10 + float64(len(data.Subject))*1.5
		data.StRectWidth += 10 + float64(len(data.Status))*2
		data.Width = data.SbRectWidth + data.StRectWidth
		if data.Color == "" {
			data.Color = "97CA00"
		}
		t = templ.ForTheBadgeTpl
	}
	err := t.Execute(&bf, data)
	if err != nil {
		return nil, err
	}
	return bf.Bytes(), nil
}
