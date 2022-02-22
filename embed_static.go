package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-contrib/static"
)

// https://github.com/gin-contrib/static/issues/19

type embedFS struct {
	http.FileSystem
}

func (e embedFS) Exists(prefix string, path string) bool {
	_, err := e.Open(path)
	return err == nil
}

func EmbedDir(fsEmbed embed.FS, targetPath string) static.ServeFileSystem {
	fsys, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFS{FileSystem: http.FS(fsys)}
}
