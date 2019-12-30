// +build js

//
package main

import (
	iotmakerPlatformIDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	coordinateManager "github.com/helmutkemper/iotmaker.platform.coordinate"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/Html"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/canvas"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/document"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/eventMouse"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/factoryBrowserDocument"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/factoryBrowserImage"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/factoryBrowserStage"
	webBrowserMouse "github.com/helmutkemper/iotmaker.platform.webbrowser/mouse"
	"github.com/helmutkemper/iotmaker.platform/abstractType/draw"
	"github.com/helmutkemper/iotmaker.platform/factoryImage"
	"github.com/helmutkemper/iotmaker.platform/mouse"
)

var (
	density                                   = 1.0
	densityManager coordinateManager.IDensity = &coordinateManager.Density{}
	bx2                                       = &draw.BasicBox{}
	stage                                     = canvas.Stage{}
	gradientFilter iotmakerPlatformIDraw.IFilterGradientInterface
)

var html iotmakerPlatformIDraw.IHtml
var browserDocument document.Document
var imgSpace interface{}
var imgPlayer interface{}

func prepareDataBeforeRun() {

	html = &Html.Html{}
	browserDocument = factoryBrowserDocument.NewDocument()
	stage = factoryBrowserStage.NewStage(
		browserDocument,
		"stage",
		800,
		600,
		density,
		densityManager,
	)

	imgSpace = factoryBrowserImage.NewImage(
		html,
		browserDocument.SelfDocument,
		map[string]interface{}{
			"id":  "spacecraft",
			"src": "./small.png",
		},
		true,
		false,
	)

	imgPlayer = factoryBrowserImage.NewImage(
		html,
		browserDocument.SelfDocument,
		map[string]interface{}{
			"id":  "player",
			"src": "./player_big.png",
		},
		true,
		false,
	)
}

func main() {

	done := make(chan struct{}, 0)
	prepareDataBeforeRun()

	i := factoryImage.NewImage(
		&stage.Canvas,
		&stage.ScratchPad,
		imgSpace,
		-100,
		-100,
		29,
		50,
		density,
		densityManager,
	)
	i.SetDraggable(true)
	stage.Add(i.Draw)

	browserDocument.AddEventListener(eventMouse.KMouseMove, webBrowserMouse.SetMouseMoveManager(mouse.ManagerMouseMove))
	//mouse.AddFunctionPointer("bBox2", bx2.GetCollisionBox, bateu)

	<-done
}
