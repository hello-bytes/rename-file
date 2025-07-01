package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type RenameAction struct {
	Action string
	Params []string
}

func main() {
	// 定义命令行参数
	var (
		dirPath = flag.String("dir", "", "要处理的目录路径")
		action  = flag.String("action", "", "重命名动作: replace.ext, add.ext, order.name")
		params  = flag.String("params", "", "动作参数，用逗号分隔")
	)

	flag.Parse()

	// 检查必需参数
	if *dirPath == "" {
		log.Fatal("请指定目录路径: -dir <path>")
	}

	if *action == "" {
		log.Fatal("请指定重命名动作: -action <action>")
	}

	// 解析参数
	var actionParams []string
	if *params != "" {
		actionParams = strings.Split(*params, ",")
	}

	renameAction := RenameAction{
		Action: *action,
		Params: actionParams,
	}

	// 执行重命名操作
	err := renameResource(*dirPath, renameAction)
	if err != nil {
		log.Fatal("重命名失败:", err)
	}

	fmt.Println("重命名操作完成!")
}

func renameResource(src string, action RenameAction) error {
	var files []string

	// 收集所有文件
	err := filepath.Walk(src, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !fi.IsDir() {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return err
	}

	// 根据动作类型执行不同的重命名操作
	switch action.Action {
	case "replace.ext":
		return replaceExtension(files, action.Params)
	case "add.ext":
		return addExtension(files, action.Params)
	case "order.name":
		return orderByName(files, action.Params)
	default:
		return fmt.Errorf("不支持的动作: %s", action.Action)
	}
}

// replace.ext: 替换文件扩展名
// 参数: [原扩展名, 新扩展名]
func replaceExtension(files []string, params []string) error {
	if len(params) < 2 {
		return fmt.Errorf("replace.ext 需要两个参数: 原扩展名, 新扩展名")
	}

	oldExt := params[0]
	newExt := params[1]

	// 确保扩展名以.开头
	if !strings.HasPrefix(oldExt, ".") {
		oldExt = "." + oldExt
	}
	if !strings.HasPrefix(newExt, ".") {
		newExt = "." + newExt
	}

	for _, file := range files {
		if strings.HasSuffix(file, oldExt) {
			newPath := strings.ReplaceAll(file, oldExt, newExt)
			if err := os.Rename(file, newPath); err != nil {
				log.Printf("重命名失败 %s -> %s: %v", file, newPath, err)
			} else {
				fmt.Printf("重命名: %s -> %s\n", filepath.Base(file), filepath.Base(newPath))
			}
		}
	}

	return nil
}

// add.ext: 为没有扩展名的文件添加扩展名
// 参数: [扩展名]
func addExtension(files []string, params []string) error {
	if len(params) < 1 {
		return fmt.Errorf("add.ext 需要一个参数: 扩展名")
	}

	ext := params[0]
	if !strings.HasPrefix(ext, ".") {
		ext = "." + ext
	}

	for _, file := range files {
		// 检查文件是否已经有扩展名
		if filepath.Ext(file) == "" {
			newPath := file + ext
			if err := os.Rename(file, newPath); err != nil {
				log.Printf("重命名失败 %s -> %s: %v", file, newPath, err)
			} else {
				fmt.Printf("添加扩展名: %s -> %s\n", filepath.Base(file), filepath.Base(newPath))
			}
		}
	}

	return nil
}

// order.name: 按数字顺序重命名文件
// 参数: [起始数字(可选，默认为1)]
func orderByName(files []string, params []string) error {
	startNum := 1
	if len(params) >= 1 {
		if num, err := strconv.Atoi(params[0]); err == nil {
			startNum = num
		}
	}

	// 按文件名排序
	sort.Strings(files)

	for i, file := range files {
		dir := filepath.Dir(file)
		ext := filepath.Ext(file)
		name := filepath.Base(file)

		// 移除扩展名
		if ext != "" {
			name = strings.TrimSuffix(name, ext)
		}

		// 生成新的文件名
		newName := fmt.Sprintf("%d%s", startNum+i, ext)
		newPath := filepath.Join(dir, newName)

		if err := os.Rename(file, newPath); err != nil {
			log.Printf("重命名失败 %s -> %s: %v", file, newPath, err)
		} else {
			fmt.Printf("重命名: %s -> %s\n", filepath.Base(file), newName)
		}
	}

	return nil
}
