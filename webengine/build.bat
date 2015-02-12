@echo off

setlocal
qmake
nmake clean
nmake release
echo '-----------------------------------'
echo 'Copy the dll to your QT Bin folder'
