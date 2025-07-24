package utils

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func MultipleFiles(c *gin.Context) []string {
	// 将文件保存到 .temp/uploads 目录，如果不存在则创建
	dir := Dir(".temp", "uploads")
	// 把上传的文件保存到uploads文件夹，保留原名称，并返回文件路径
	filePaths := []string{}
	for _, file := range c.Request.MultipartForm.File["file"] {
		filePath := dir.Join(file.Filename).String()
		c.SaveUploadedFile(file, filePath)
		filePaths = append(filePaths, filePath)
	}
	return filePaths
}

func MultipleFilesWithTimestamp(c *gin.Context) []string {
	// 类似于 MultipleFiles 但在不改变文件后缀的情况下，在文件名前添加时间戳
	dir := Dir(".temp", "uploads")
	filePaths := []string{}
	for _, file := range c.Request.MultipartForm.File["file"] {
		// 生成时间戳
		timestamp := time.Now().Unix()
		// 生成新文件名
		newFilename := fmt.Sprintf("%d_%s", timestamp, file.Filename)
		// 保存文件
		filePath := dir.Join(newFilename).String()
		c.SaveUploadedFile(file, filePath)
		filePaths = append(filePaths, filePath)
	}
	return filePaths
}
