package main

import (
	"errors"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/jithinkunjachan/nasserver/backend/pkg/command"
	"github.com/jithinkunjachan/nasserver/backend/pkg/executor"
	"github.com/jithinkunjachan/nasserver/backend/pkg/render"
	"github.com/jithinkunjachan/nasserver/backend/pkg/ws"
)

var isBusy = false
var mutex = sync.Mutex{}

func main() {
	log.Println("starting server")
	r := gin.Default()

	tmpls, err := render.NewRender()
	if err != nil {
		log.Fatalf("%v", err)
	}

	webSckt := ws.NewWs(tmpls)

	r.GET("/", func(ctx *gin.Context) {
		tmpls.ExecuteTemplate(ctx.Writer, "index", nil)
	})

	r.GET("/disk", func(ctx *gin.Context) {
		err := Busy(true)
		if err != nil {
			ctx.AbortWithStatus(http.StatusLocked)
			return
		}
		defer Busy(false)
		tmpls.ExecuteTemplate(ctx.Writer, "disk", nil)
	})

	r.GET("/snapraid", func(ctx *gin.Context) {
		err := Busy(true)
		if err != nil {
			ctx.AbortWithStatus(http.StatusLocked)
			return
		}
		defer Busy(false)
		tmpls.ExecuteTemplate(ctx.Writer, "snapraid", nil)
	})

	r.GET("/system", func(ctx *gin.Context) {
		err := Busy(true)
		if err != nil {
			ctx.AbortWithStatus(http.StatusLocked)
			return
		}
		defer Busy(false)
		tmpls.ExecuteTemplate(ctx.Writer, "system", nil)
	})

	r.GET("/ws", webSckt.Handle)

	r.GET("/shutdown", func(ctx *gin.Context) {
		err := Busy(true)
		if err != nil {
			ctx.AbortWithStatus(http.StatusLocked)
			return
		}
		defer Busy(false)
		shutdown := &command.Sudo{
			Ws:   webSckt,
			Cmd:  "sudo",
			Args: []string{"-S", "--", "shutdown", "now"},
		}
		executor.Exec(shutdown)
	})

	r.GET("/reboot", func(ctx *gin.Context) {
		err := Busy(true)
		if err != nil {
			ctx.AbortWithStatus(http.StatusLocked)
			return
		}
		defer Busy(false)
		shutdown := &command.Sudo{
			Ws:   webSckt,
			Cmd:  "sudo",
			Args: []string{"-S", "--", "reboot"},
		}
		executor.Exec(shutdown)
	})

	r.GET("/update", func(ctx *gin.Context) {
		err := Busy(true)
		if err != nil {
			ctx.AbortWithStatus(http.StatusLocked)
			return
		}
		defer Busy(false)
		shutdown := &command.Sudo{
			Ws:   webSckt,
			Cmd:  "sudo",
			Args: []string{"-S", "--", "apt", "update"},
		}
		executor.Exec(shutdown)
	})
	r.GET("/upgrade", func(ctx *gin.Context) {
		err := Busy(true)
		if err != nil {
			ctx.AbortWithStatus(http.StatusLocked)
			return
		}
		defer Busy(false)
		shutdown := &command.Sudo{
			Ws:   webSckt,
			Cmd:  "sudo",
			Args: []string{"-S", "--", "apt", "-y", "upgrade"},
		}
		executor.Exec(shutdown)
	})

	r.GET("/lsblk", func(ctx *gin.Context) {
		err := Busy(true)
		if err != nil {
			ctx.AbortWithStatus(http.StatusLocked)
			return
		}
		defer Busy(false)
		lsblk := &command.Lsblk{
			Ws: webSckt,
		}
		executor.Exec(lsblk)
	})

	r.GET("/blkid", func(ctx *gin.Context) {
		err := Busy(true)
		if err != nil {
			ctx.AbortWithStatus(http.StatusLocked)
			return
		}
		defer Busy(false)
		blkid := &command.Blkid{
			Ws: webSckt,
		}
		executor.Exec(blkid)
	})

	r.GET("/snapraid/sync", func(ctx *gin.Context) {
		err := Busy(true)
		if err != nil {
			ctx.AbortWithStatus(http.StatusLocked)
			return
		}
		defer Busy(false)
		snpRaid := &command.SnapRaid{
			Ws:   webSckt,
			Cmd:  "snapraid",
			Args: []string{"sync"},
		}
		executor.Exec(snpRaid)
	})
	r.GET("/snapraid/status", func(ctx *gin.Context) {
		err := Busy(true)
		if err != nil {
			ctx.AbortWithStatus(http.StatusLocked)
			return
		}
		defer Busy(false)
		snpRaid := &command.SnapRaid{
			Ws:   webSckt,
			Cmd:  "snapraid",
			Args: []string{"status"},
		}
		executor.Exec(snpRaid)
	})

	r.GET("/snapraid/diff", func(ctx *gin.Context) {
		err := Busy(true)
		if err != nil {
			ctx.AbortWithStatus(http.StatusLocked)
			return
		}
		defer Busy(false)
		snpRaid := &command.SnapRaid{
			Ws:   webSckt,
			Cmd:  "snapraid",
			Args: []string{"diff"},
		}
		executor.Exec(snpRaid)
	})

	r.GET("/snapraid/scrub", func(ctx *gin.Context) {
		err := Busy(true)
		if err != nil {
			ctx.AbortWithStatus(http.StatusLocked)
			return
		}
		defer Busy(false)
		snpRaid := &command.SnapRaid{
			Ws:   webSckt,
			Cmd:  "snapraid",
			Args: []string{"scrub"},
		}
		executor.Exec(snpRaid)
	})

	r.GET("/snapraid/list", func(ctx *gin.Context) {
		err := Busy(true)
		if err != nil {
			ctx.AbortWithStatus(http.StatusLocked)
			return
		}
		defer Busy(false)
		snpRaid := &command.SnapRaid{
			Ws:   webSckt,
			Cmd:  "snapraid",
			Args: []string{"list"},
		}
		executor.Exec(snpRaid)
	})

	r.GET("/snapraid/dup", func(ctx *gin.Context) {
		err := Busy(true)
		if err != nil {
			ctx.AbortWithStatus(http.StatusLocked)
			return
		}
		defer Busy(false)
		snpRaid := &command.SnapRaid{
			Ws:   webSckt,
			Cmd:  "snapraid",
			Args: []string{"dup"},
		}
		executor.Exec(snpRaid)
	})

	r.GET("/snapraid/check", func(ctx *gin.Context) {
		err := Busy(true)
		if err != nil {
			ctx.AbortWithStatus(http.StatusLocked)
			return
		}
		defer Busy(false)
		snpRaid := &command.SnapRaid{
			Ws:   webSckt,
			Cmd:  "snapraid",
			Args: []string{"check"},
		}
		executor.Exec(snpRaid)
	})

	r.GET("/snapraid/touch", func(ctx *gin.Context) {
		err := Busy(true)
		if err != nil {
			ctx.AbortWithStatus(http.StatusLocked)
			return
		}
		defer Busy(false)
		snpRaid := &command.SnapRaid{
			Ws:   webSckt,
			Cmd:  "snapraid",
			Args: []string{"touch"},
		}
		executor.Exec(snpRaid)
	})

	r.GET("/snapraid/smart", func(ctx *gin.Context) {
		err := Busy(true)
		if err != nil {
			ctx.AbortWithStatus(http.StatusLocked)
			return
		}
		defer Busy(false)
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

func Busy(b bool) error {
	if b == isBusy {
		return errors.New("server is working please wait")
	}
	mutex.Lock()
	isBusy = b
	mutex.Unlock()
	return nil
}
