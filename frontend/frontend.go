package frontend

import (
	"embed"
	"io/fs"
)

//go:embed build/**
var files embed.FS

func Files() fs.FS {
	return files
}
