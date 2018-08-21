package tpl

// ForTheBadgeTpl is template for for-the-badge style
var ForTheBadgeTpl = `
<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="{{ normalizeFloat64 .Width }}"
  height="28">
  <g shape-rendering="crispEdges">
    <path fill="#555" d="M0 0h{{ normalizeFloat64 .SbRectWidth }}v28H0z"
    />
    <path fill="#{{ .Color }}" d="M{{ normalizeFloat64 .SbRectWidth }} 0h{{ normalizeFloat64 .StRectWidth }}v28H{{ normalizeFloat64 .SbRectWidth }}z"
    />
  </g>
  <g fill="#fff" text-anchor="middle" font-family="DejaVu Sans,Verdana,Geneva,sans-serif" font-size="100">
    <text x="{{ ftbfunc1 .SbRectWidth }}" y="175" transform="scale(.1)"
      textLength="{{ ftbfunc2 .SbRectWidth }}">{{ .Subject }}</text>
    <text x="{{ ftbfunc3 .SbRectWidth .StRectWidth }}" y="175" font-weight="bold" transform="scale(.1)"
      textLength="{{ ftbfunc2 .StRectWidth }}">{{ .Status }}</text>
  </g>
</svg>
`
