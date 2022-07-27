// Package ui handles the PocketBase Admin frontend embedding.
package ui

import (
	"embed"
)

//go:embed dist
var DistDir embed.FS

// //go:embed dist/index.html
// var indexHTML embed.FS
