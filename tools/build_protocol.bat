@echo off
chcp 65001 >nul
@echo "-----------fix package name(本地化)------------------"

py  .\amend.py
timeout 1
echo build go protocol file...
cd ../internal
cd protocol

:: 检查gofile目录是否存在
if not exist "gofile" (
    echo 创建目录: %CD%/gofile
    mkdir gofile
) else (
    echo 目录已存在: %CD%/gofile
)

:: 检查jsfile目录是否存在
if not exist "jsfile" (
    echo 创建目录: %CD%/jsfile
    mkdir jsfile
) else (
    echo 目录已存在: %CD%/jsfile
)

rem echo %CD%

for /f "delims=" %%i in ('dir /A:-D /B ') do (
echo %%i -- GoFile jsFile
protoc -I./ --go_out=./gofile/ --go_opt=paths=source_relative %%i
protoc --js_out=import_style=commonjs,binary:./jsfile/ %%i
)

echo build proto to "go and js" complete!

@echo off

cd ../../tools
node onekey.js
rem timeout /t 2
echo 开始 3s 倒计时...
choice /c  abcdQ /n /t 3 /d a /m "如需暂停,请按Q键终止,否则退出"
if %errorlevel%==5 goto stopTerminal
goto :eof
:stopTerminal
pause


