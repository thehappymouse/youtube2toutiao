#!/bin/sh

# 此脚本由 GO 程序调用，传入一个 youtube 地址进行下载
# 依赖于翻墙代理
# GO 程序对输出结果进行验证（验证也可以转移至此脚本内）

#!/bin/sh

youtube-dl $1 --proxy http://127.0.0.1:55218/ -f mp4 --write-thumbnail --write-description