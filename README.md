# 文件重命名工具

这是一个支持多种重命名操作的命令行工具。

## 功能特性

支持以下三种重命名操作：

1. **replace.ext** - 替换文件扩展名
2. **add.ext** - 为没有扩展名的文件添加扩展名
3. **order.name** - 按数字顺序重命名文件

## 使用方法

### 基本语法

```bash
go run main.go -dir <目录路径> -action <动作> -params <参数>
```

### 参数说明

- `-dir`: 要处理的目录路径（必需）
- `-action`: 重命名动作（必需）
  - `replace.ext`: 替换扩展名
  - `add.ext`: 添加扩展名
  - `order.name`: 按数字顺序重命名
- `-params`: 动作参数，用逗号分隔（可选）

## 使用示例

### 1. 替换文件扩展名

将目录中所有 `.asc` 文件重命名为 `.md` 文件：

```bash
go run main.go -dir "/path/to/directory" -action "replace.ext" -params "asc,md"
```

### 2. 添加文件扩展名
为目录中没有扩展名的文件添加 `.txt` 扩展名：
```bash
go run main.go -dir "/path/to/directory" -action "add.ext" -params "txt"
```

### 3. 按数字顺序重命名
将目录中的所有文件按数字顺序重命名（从1开始）：
```bash
go run main.go -dir "/path/to/directory" -action "order.name"
```

从10开始重命名：
```bash
go run main.go -dir "/path/to/directory" -action "order.name" -params "10"
```

## 注意事项

- 工具会递归处理指定目录下的所有文件
- 重命名操作会保持文件的扩展名不变（除了replace.ext操作）
- 如果目标文件名已存在，重命名操作可能会失败
- 建议在重要文件上使用前先备份

## 编译

编译为可执行文件：

```bash
go build -o rename main.go
```

然后可以直接运行：
```bash
./rename -dir "/path/to/directory" -action "replace.ext" -params "asc,md"
``` 
