#!/bin/sh

export LANG=zh_CN.UTF-8
export LC_CTYPE=zh_CN.UTF-8

# 此脚本由 GO 程序调用，传入一个 youtube 地址进行下载
# 依赖于翻墙代理
# GO 程序对输出结果进行验证（验证也可以转移至此脚本内）
# 命令参考 https://blog.csdn.net/weixin_43278670/article/details/109466413

# 2020.11.01.1
echo pwd `pwd`
youtube-dl $1 -f "best[ext=mp4]" --write-description --write-thumbnail --proxy http://127.0.0.1:1087/
# youtube-dl $1 -f "best[ext=mp4]" --write-description --write-thumbnail --proxy http://127.0.0.1:1087/ 2>&1