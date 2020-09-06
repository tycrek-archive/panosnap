@echo off
set GOOS=windows
set GOARCH=amd64
set CGO_ENABLED=1
rem set CC=C:\Mingw\bin\gcc.exe
rem set CXX=C:\Mingw\bin\g++.exe
go build -buildmode=exe panosnap.go