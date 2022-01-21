#!/bin/bash
###
 # @Author: Bin
 # @Date: 2021-10-02
 # @FilePath: /mirai-plugin-checkclass/build.sh
### 

cachePath="./cache" # 缓存目录
buildPath="./build" # 构建输出目录
buildPackagePath="$buildPath/plugin.tar" # 构建包输出目录
codePath="./main.go"
manifestPath="./manifest.json"

# 检查项目是否为插件包
[ ! -f "$manifestPath" ] && echo -e "\033[31m编译失败，请检查该项目是否属于插件目录。\033[0m" && exit 1
[ ! -f "$codePath" ] && echo -e "\033[31m编译失败，该项目不存在代码文件 main.go\033[0m" && exit 1
 
# 判断缓存文件夹和构建文件夹是否存在
[ -d "$cachePath/" ] && rm -rf "$cachePath/" # 清理缓存文件夹
[ -f "$buildPackagePath" ] && rm -rf "$buildPackagePath" # 清理构建包

# 创建缓存目录和构建输出目录
mkdir -p "$cachePath/"
[ ! -d "$buildPath/" ] && mkdir -p "$buildPath/"

echo -e "开始构建插件包！🦕 "

# 拷贝包文件
/bin/cp "$manifestPath" "$cachePath/manifest.json"
/bin/cp ./view.json "$cachePath/view.json"

# 编译插件代码
go build -o="$cachePath/build.so" -buildmode=plugin $codePath
echo -e "插件编译成功，正在打包…"

# 压缩插件包文件
bale_log=$(cd "$cachePath/" && tar -cf "../$buildPackagePath" * && cd -)
echo -e "\033[32m插件包打包成功！路径：$buildPackagePath\033[0m"

# 清理缓存
[ -d "$cachePath/" ] && rm -rf "$cachePath/" # 清理缓存文件夹
