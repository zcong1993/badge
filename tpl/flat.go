package tpl

// FlatTpl is tpl for flat style
var FlatTpl = `
      <svg width="{{ .Width }}" height="20" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
        <g>
          <rect width="{{ .SbRectWidth }}" height="20" fill="#555"/>
          <rect x="{{ .SbRectWidth }}" width="{{ .StRectWidth }}" height="20" fill="#{{ .Color }}"/>
        </g>
        <g fill="#fff" text-anchor="start" font-family="Verdana,DejaVu Sans,sans-serif" font-size="11">
          <text x="6.3" y="14.8" textLength="{{ .SbTextWidth }}" fill="#000" opacity="0.1">{{ .Subject }}</text>
          <text x="5.3" y="13.8" textLength="{{ .SbTextWidth }}">{{ .Subject }}</text>
          <text x="{{ add .SbRectWidth 5.5 }}" y="14.8" fill="#000" opacity="0.1" textLength="{{ .StTextWidth }}">{{ .Status }}</text>
          <text x="{{ add .SbRectWidth 4.5 }}" y="13.8" textLength="{{ .StTextWidth }}">{{ .Status }}</text>
        </g>
      </svg>
`
