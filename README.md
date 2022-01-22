# 网课搜题 QQ 机器人插件
> 网课搜题机器人插件，可安装到 [mirai-class-notice](https://github.com/PBK-B/mirai-class-notice) 后台

### 编译
```
$ sudo chmod +x ./build.sh
$ bash build.sh ./
```

### 项目结构
```
.
├── README.md
├── build             插件包编译输出目录
├── build.sh          构建脚本
├── go.mod
├── go.sum
├── main.go           插件源文件
├── manifest.json     插件描述文件
└── view.json         插件视图文件
```