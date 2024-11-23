// Package ui adding external files
package ui

import "embed"

//go:embed "html" "static"
var Files embed.FS
