package templates

import "embed"

// Files contains the project templates shipped with the arche binary.
// The all: prefix includes dotfiles such as .env.example.
//
//go:embed all:flask
var Files embed.FS
