#!/bin/bash

PROJECT_ROOT=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)  # 自动获取项目根目录（脚本所在目录）
Appname="gin-redis-shell"                                           # 应用名称
MAIN_FILE="$PROJECT_ROOT/main.go"                           # 入口文件路径
BUILD_DIR="$PROJECT_ROOT"                               # 编译产物目录
LOG_FILE="$PROJECT_ROOT/logs/app.log"                       # 日志文件路径
MODULE_NAME="gin-redis-shell"                     # go.mod 中的模块名
help_run(){
    echo "start 开始项目"
    echo "build 编译项目"
    echo "restart 重启项目"
    echo "stop 停止项目"
    echo "status 查看项目状态"
    echo "help 帮助"

}
build(){
    echo "开始编译项目..."
    if [ -d "$BUILD_DIR" ]; then
        rm -rf "$BUILD_DIR"
    fi
    mkdir "$BUILD_DIR"
    go build -o "$BUILD_DIR/$Appname" "$MAIN_FILE"
    if [ $? -eq 0 ]; then
        echo "编译成功，生成文件: $BUILD_DIR/$Appname"
    else
        echo "编译失败"
    fi
}
start(){
    if pgrep -x "$Appname" > /dev/null; then
        echo "$Appname 已经在运行中"
    else
        echo "正在启动 $Appname ..."                            
        "$BUILD_DIR/$Appname" &
        if [ $? -eq 0 ]; then
            echo "$Appname 启动成功"
        else
            echo "$Appname 启动失败"            
        fi
    fi
}
stop(){                 
    if pgrep -x "$Appname" > /dev/null; then
        echo "正在停止 $Appname ..."
        pkill -x "$Appname"
        if [ $? -eq 0 ]; then
            echo "$Appname 停止成功"
        else
            echo "$Appname 停止失败"
        fi
    else
        echo "$Appname 未在运行中"
    fi
}                   

status(){
    if pgrep -x "$Appname" > /dev/null; then
        echo "$Appname 正在运行中"
    else
        echo "$Appname 未在运行中"
    fi
}
restart(){
    stop
    start
}      
         
if [ $# -eq 0 ]; then
    echo "usage: ./manager.sh {build|start|restart|stop|status}"
    exit 1
fi

case $1 in
    build)
        build
        ;;
    start)
        start
        ;;
    stop)
        stop
        ;;
    restart)
        restart
        ;;
    status)
        status
        ;;
    help|*)
        help_run
        ;;
esac    

exit 0
