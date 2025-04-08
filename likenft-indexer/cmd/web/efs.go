package web

import (
	"embed"

	_ "github.com/a-h/templ"
	_ "github.com/a-h/templ/runtime"
)

//go:embed "assets"
var Files embed.FS
