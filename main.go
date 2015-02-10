package main

import (
	"fmt"
	"github.com/chai2010/qml"
	"github.com/maran/go-webengine-windows/webengine"
	"os"
	"runtime"
)

func run() error {
	webengine.Initialize()
	engine := qml.NewEngine()
	engine.On("quit", func() { os.Exit(0) })

	component, err := engine.LoadString("hello.qml", qmlHello)
	if err != nil {
		return err
	}
	window := component.CreateWindow(nil)
	window.Set("height", 640)
	window.Set("width", 860)
	window.Show()
	window.Wait()

	return nil
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if err := qml.Run(run); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

const qmlHello = `
import QtQuick 2.0
import QtQuick.Controls 1.0;
import QtQuick.Controls.Styles 1.0
import QtQuick.Layouts 1.0;
import QtWebEngine 1.0
//import QtWebEngine.experimental 1.0
import QtQuick.Window 2.0;

Rectangle {
	id: window
	anchors.fill: parent
	color: "#00000000"

	property var title: ""
	property var menuItem
	property var hideUrl: true

	property alias url: webview.url
	property alias windowTitle: webview.title
	property alias webView: webview

	property var cleanPath: false
	property var open: function(url) {
		if(!window.cleanPath) {
			var uri = url;
			if(!/.*\:\/\/.*/.test(uri)) {
				uri = "http://" + uri;
			}

			window.cleanPath = true;

			webview.url = uri;

			//uriNav.text = uri.text.replace(/(^https?\:\/\/(?:www\.)?)([a-zA-Z0-9_\-]*\.\w{2,3})(.*)/, "$1$2<span style='color:#CCC'>$3</span>");
			uriNav.text = uri;
		} else {
			// Prevent inf loop.
			window.cleanPath = false;
		}
	}

	Component.onCompleted: {
		webview.url = "http://google.com"
	}

	Item {
		objectName: "root"
		id: root
		anchors.fill: parent
		state: "inspectorShown"

		RowLayout {
			id: navBar
			height: 40
			anchors {
				left: parent.left
				right: parent.right
				leftMargin: 7
			}

			TextField {
				anchors {
					leftMargin: 10
					rightMargin: 10
				}
				width: 400
				text: webview.url;
				id: uriNav
				y: parent.height / 2 - this.height / 2

				Keys.onReturnPressed: {
					webview.url = this.text;
				}
			}
		}

		// Border
		Rectangle {
			id: divider
			anchors {
				left: parent.left
				right: parent.right
				top: navBar.bottom
			}
			z: -1
			height: 1
			color: "#CCCCCC"
		}

		WebEngineView {
			objectName: "webView"
			id: webview
			anchors {
				left: parent.left
				right: parent.right
				bottom: parent.bottom
				top: divider.bottom
			}

			onLoadingChanged: {
				if (loadRequest.status == WebEngineView.LoadSucceededStatus) {
					webview.runJavaScript("document.title", function(pageTitle) {
						menuItem.title = pageTitle;
					});
				}
			}
			onJavaScriptConsoleMessage: {
				console.log(sourceID + ":" + lineNumber + ":" + JSON.stringify(message));
			}
		}

		Rectangle {
			id: sizeGrip
			color: "gray"
			visible: false
			height: 10
			anchors {
				left: root.left
				right: root.right
			}
			y: Math.round(root.height * 2 / 3)

			MouseArea {
				anchors.fill: parent
				drag.target: sizeGrip
				drag.minimumY: 0
				drag.maximumY: root.height
				drag.axis: Drag.YAxis
			}
		}

		WebEngineView {
			id: inspector
			visible: false
			anchors {
				left: root.left
				right: root.right
				top: sizeGrip.bottom
				bottom: root.bottom
			}

		}

		states: [
			State {
				name: "inspectorShown"
				PropertyChanges {
					target: inspector
				}
			}
		]
	}
}
`
