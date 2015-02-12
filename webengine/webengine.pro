TEMPLATE = lib
CONFIG  += dll release
CONFIG  -= embed_manifest_exe embed_manifest_dll
QT      += webengine
TARGET   = go_webengine

DESTDIR = $${PWD}
INCLUDEPATH += .

SOURCES += ./all.cpp

HEADERS += ./cpp/webengine.h
SOURCES += ./cpp/webengine.cpp

DEF_FILE+= ./webengine.def
