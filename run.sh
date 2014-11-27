#!/bin/sh

echo "setting go path "
currentPath=`pwd`
export GOPATH="${currentPath}"
export PATH="$PATH:$GOPATH/bin"

echo "you should run getPackage.sh first otherwise go server will be fail"
echo "setted go path"

echo "when you shut down this go server you "
echo "should copy: export GOPATH=${currentPath}"
echo "         and export PATH=\$PATH:\$GOPATH/bin"
echo "to you terminal"

echo "starting go server......"
go run ${currentPath}/start.go

