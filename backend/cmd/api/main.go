package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jithinkunjachan/nasserver/backend/pkg/command"
	"github.com/jithinkunjachan/nasserver/backend/pkg/executor"
	"github.com/jithinkunjachan/nasserver/backend/pkg/ws"
)

func main() {
	log.Println("starting server")
	r := gin.Default()

	webSckt := ws.NewWs()

	r.GET("/ws", webSckt.Handle)

	r.GET("/lsblk", func(ctx *gin.Context) {
		lsblk := &command.Lsblk{
			Ws: webSckt,
		}
		executor.Exec(lsblk)
	})

	r.GET("/blkid", func(ctx *gin.Context) {
		blkid := &command.Blkid{
			Ws: webSckt,
		}
		executor.Exec(blkid)
	})

	r.GET("/snapraid/sync", func(ctx *gin.Context) {
		snpRaid := &command.SnapRaid{
			Ws:   webSckt,
			Cmd:  "snapraid",
			Args: []string{"sync"},
		}
		executor.Exec(snpRaid)
	})
	r.GET("/snapraid/status", func(ctx *gin.Context) {
		snpRaid := &command.SnapRaid{
			Ws:   webSckt,
			Cmd:  "snapraid",
			Args: []string{"status"},
		}
		executor.Exec(snpRaid)
	})

	r.GET("/snapraid/diff", func(ctx *gin.Context) {
		snpRaid := &command.SnapRaid{
			Ws:   webSckt,
			Cmd:  "snapraid",
			Args: []string{"diff"},
		}
		executor.Exec(snpRaid)
	})

	r.GET("/snapraid/scrub", func(ctx *gin.Context) {
		snpRaid := &command.SnapRaid{
			Ws:   webSckt,
			Cmd:  "snapraid",
			Args: []string{"scrub"},
		}
		executor.Exec(snpRaid)
	})

	r.GET("/snapraid/list", func(ctx *gin.Context) {
		snpRaid := &command.SnapRaid{
			Ws:   webSckt,
			Cmd:  "snapraid",
			Args: []string{"list"},
		}
		executor.Exec(snpRaid)
	})

	r.GET("/snapraid/dup", func(ctx *gin.Context) {
		snpRaid := &command.SnapRaid{
			Ws:   webSckt,
			Cmd:  "snapraid",
			Args: []string{"dup"},
		}
		executor.Exec(snpRaid)
	})

	r.GET("/snapraid/check", func(ctx *gin.Context) {
		snpRaid := &command.SnapRaid{
			Ws:   webSckt,
			Cmd:  "snapraid",
			Args: []string{"check"},
		}
		executor.Exec(snpRaid)
	})

	r.GET("/snapraid/smart", func(ctx *gin.Context) {
		snpRaid := &command.SnapRaid{
			Ws:   webSckt,
			Cmd:  "sudo",
			Args: []string{"-S", "--", "snapraid", "smart"},
		}
		executor.Exec(snpRaid)
	})

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]string{
			"STATUS": "UP",
		})
	})

	r.Run(":8081")
}
