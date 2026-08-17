package templates

import "embed"

//go:embed all:go all:python all:cpp all:js
var FS embed.FS
