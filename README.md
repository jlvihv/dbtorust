# dbtogo

这是一个将数据库中存在的表转换为 rust 结构体的命令行工具，用 go 语言写的，考虑以后使用 rust 重写，现在对 rust 还不是太熟

暂时只支持 mysql 和 mariadb

## 安装

```bash
go install github.com/jlvihv/dbtorust@latest
```

## 使用

使用前，必须在 config.toml 配置文件中配置数据库连接信息：

默认情况下，程序会读取 `~/.config/dbtogo/config.toml` 处的配置文件

可以使用 --config 选项自定义配置文件位置

配置示例：config.toml

```toml
[db]
ip = "127.0.0.1"
port = "3306"
username = "root"
password = "your_password"
charset = "utf8mb4"
```

支持三种输出方式：
1. 输出到标准输出 (默认)
2. 输出到剪贴板
3. 输出到文件

使用时，必须在命令中指定数据库名和表名：

使用示例：

输出到命令行
```bash
dbtogo --db 数据库名 --table 表名
```

输出到系统剪贴板
```bash
dbtogo --db 数据库名 --table 表名 --clip
```

输出到文件
```bash
dbtogo --db 数据库名 --table 表名 --file struct.txt
```

可以同时指定多个表名，用逗号隔开，不要加空格
```bash
dbtogo --db 数据库名 --table 表名1,表名2
```

如果一定想在多个表名之间加空格，那就把它们放在一对引号里
```bash
dbtogo --db 数据库名 --table "表名1, 表名2"
```

指定配置文件位置
```bash
dbtogo --db 数据库名 --table 表名 --config ~/dbtogo.toml
```

更多信息，使用 -h 选项查看

## 备注

- 实现过程参考了煎鱼大佬的代码 https://github.com/go-programming-tour-book/tour
- 输出到剪贴板并不总是可以成功的，在 linux 下，你可能需要安装 xclip 或 xsel 命令，具体请看这里 https://github.com/atotto/clipboard

## License

MIT
