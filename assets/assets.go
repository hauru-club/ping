// Package assets contains static files required
// for frontend part of ping web application.
package assets

import "embed"

//go:embed index.html
var Index []byte

//go:embed css js
var StaticFiles embed.FS
