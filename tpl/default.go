package tpl

// DefaultTpl is default style tpl
var DefaultTpl = `
    <svg width="{{ .Width }}" height="20" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
      <linearGradient id="a" x2="0" y2="100%">
        <stop offset="0" stop-color="#EEE" stop-opacity=".1"/>
        <stop offset="1" stop-opacity=".1"/>
      </linearGradient>
      <mask id="m"><rect width="{{ .Width }}" height="20" rx="3" fill="#FFF"/></mask>
      <g mask="url(#m)">
        <rect width="{{ .SbRectWidth }}" height="20" fill="#555"/>
        <rect x="{{ .SbRectWidth }}" width="{{ .StRectWidth }}" height="20" fill="#{{ .Color }}"/>
        <rect width="{{ .Width }}" height="20" fill="url(#a)"/>
      </g>
      <g fill="#fff" text-anchor="start" font-family="Verdana,DejaVu Sans,sans-serif" font-size="11">
        <text x="6.5" y="14.8" textLength="{{ .SbTextWidth }}" fill="#000" opacity="0.25">{{ .Subject }}</text>
        <text x="5.5" y="13.8" textLength="{{ .SbTextWidth }}">{{ .Subject }}</text>
        <text x="{{ add .SbRectWidth 5.5 }}" y="14.8" fill="#000" opacity="0.25" textLength="{{ .StTextWidth }}">{{ .Status }}</text>
        <text x="{{ add .SbRectWidth 4.5 }}" y="13.8" textLength="{{ .StTextWidth }}">{{ .Status }}</text>
      </g>
    </svg>
`
