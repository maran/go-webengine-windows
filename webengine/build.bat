setlocal

qmake
nmake clean
nmake release
echo 'Copy the dll to your QT Bin folder'
