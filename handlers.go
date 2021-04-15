package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jaskaranSM/megasdkgo"
)

var megaClient *megasdkgo.MegaClient

func getMegaClient(api_key string) *megasdkgo.MegaClient {
	client := megasdkgo.NewMegaClient(api_key)
	return client
}

func LoginHandler(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	if email == "" || password == "" {
		c.JSON(400, gin.H{
			"login":   "failed",
			"message": "whares da credentials",
		})
		return
	}
	err := megaClient.Login(email, password)
	if err != nil {
		c.JSON(401, gin.H{
			"login":   "failed",
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"login":   "success",
			"message": "start using the service",
		})
	}
}

func AddDownloadHandler(c *gin.Context) {
	dl_link := c.PostForm("link")
	dir := c.PostForm("dir")
	if dir == "" || dl_link == "" {
		c.JSON(400, gin.H{
			"adddl":   "failed",
			"message": "bad input parameters",
		})
		return
	}
	gid, err := megaClient.AddDownload(dl_link, dir)
	if err != nil {
		c.JSON(401, gin.H{
			"adddl":   "failed",
			"message": err.Error(),
			"gid":     "",
			"dir":     dir,
		})
	} else {
		c.JSON(200, gin.H{
			"adddl":   "success",
			"message": "add download succeeded",
			"gid":     gid,
			"dir":     dir,
		})
	}
}

func CancelDownloadHandler(c *gin.Context) {
	gid := c.PostForm("gid")
	if gid == "" {
		c.JSON(400, gin.H{
			"canceldl": "failed",
			"message":  "whares da gid",
		})
		return
	}
	err := megaClient.CancelDownload(gid)
	if err != nil {
		c.JSON(400, gin.H{
			"canceldl": "failed",
			"message":  err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"canceldl": "success",
			"message":  "download canceled",
		})
	}
}

func GetDownloadInfoHandler(c *gin.Context) {
	gid := c.Param("gid")
	if gid == "" {
		c.JSON(401, gin.H{
			"dlinfo":  "failed",
			"message": "whares da gid",
		})
		return
	}
	status := megaClient.GetDownloadInfo(gid)
	if status.Gid == "" {
		c.JSON(404, gin.H{
			"dlinfo":  "failed",
			"message": "no such gid currently exists in downloader",
		})
	} else {
		c.JSON(200, gin.H{
			"dlinfo":           "success",
			"message":          "heres the download info",
			"name":             status.Name,
			"error_code":       status.ErrorCode,
			"error_string":     status.ErrorString,
			"gid":              status.Gid,
			"speed":            status.Speed,
			"completed_length": status.CompletedLength,
			"total_length":     status.TotalLength,
			"state":            status.State,
		})
	}
}
