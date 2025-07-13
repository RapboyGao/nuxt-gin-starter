package utils

import (
	"os"

	"github.com/arduino/go-paths-helper"
)

// Dir 根据传入的目录名称依次在当前工作目录下创建目录，并返回最终目录的路径。
// 参数 names 为要创建的目录名称列表（可变参数）。
// 返回值为最终创建的目录的 *paths.Path 对象。
func Dir(names ...string) *paths.Path {
	workingDir, _ := os.Getwd()
	dir := paths.New(workingDir)
	dir.ToAbs()
	for _, name := range names {
		dir = dir.Join(name)
		dir.Mkdir()
	}
	return dir
}
