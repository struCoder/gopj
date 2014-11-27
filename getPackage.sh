#!/bin/sh

echo "正在获取依赖package...."

currentPath=`pwd`
export GOPATH="${currentPath}"
export PATH="$PATH:${GOPATH}/bin"

go get -u -v github.com/hoisie/web
go get -u -v github.com/go-sql-driver/mysql
go get -u -v github.com/gorilla/sessions


echo "获取依赖package成功"
echo "请运行 run.sh"