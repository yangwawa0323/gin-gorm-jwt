package routers

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yangwawa0323/gin-gorm-jwt/utils"
)

type FavFS struct {
	*embed.FS
	*gin.Engine
}

var debug = utils.Debug

func FaviconFS(embeddedFiles *embed.FS) http.FileSystem {
	sub, err := fs.Sub(embeddedFiles, "assets/favicon.ico")

	if err != nil {
		debug(err.Error())
		panic("")
	}
	return http.FS(sub)
}

func FavFSRouter(favFS *FavFS) {
	favFS.Engine.GET("/favicon.ico", func(ctx *gin.Context) {
		ctx.FileFromFS(".", FaviconFS(favFS.FS))
	})
}
