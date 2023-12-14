package ui

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:generate yarn
//go:generate yarn build
//go:embed all:dist
var content embed.FS

func Dist() http.FileSystem {
	dist, err := fs.Sub(content, "dist")

	if err != nil {
		panic(err)
	}

	return http.FS(dist)
}
