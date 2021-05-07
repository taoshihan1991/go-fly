package static

import "embed"

//go:embed templates/*
var TemplatesEmbed embed.FS

//go:embed js/*
var JsEmbed embed.FS
