@echo off
chcp 65001
REM 
cd /d "D:\work\promptDSL\web-backend"
go run . "CodeRunner.exe" 
cd /d "D:\work\promptDSL\generated_code"
echo 编译成功，正在运行 CodeRunner.exe...
CodeRunner.exe input.json
