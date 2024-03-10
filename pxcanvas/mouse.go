package pxcanvas

import (
	"pixl/apptype"
	"pixl/pxcanvas/brush"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

func (pxCanvas *PxCanvas) Scrolled(ev *fyne.ScrollEvent) {
	pxCanvas.Scale(int(ev.Scrolled.DY))
	pxCanvas.Refresh()
}

func (pxCanvas *PxCanvas) MouseMoved(ev *desktop.MouseEvent) {
	if x, y := pxCanvas.MouseToCanvasXY(ev); x != nil && y != nil {
		brush.TryBrush(pxCanvas.AppState, pxCanvas, ev)
		cursor := brush.Cursor(pxCanvas.PxCanvasConfig, apptype.BrushType(pxCanvas.AppState.BrushType), ev, *x, *y)
		pxCanvas.Renderer.SetCursor(cursor)
	} else {
		pxCanvas.Renderer.SetCursor(make([]fyne.CanvasObject, 0))
	}
	pxCanvas.TryPan(pxCanvas.MouseState.PreviousCoord, ev)
	pxCanvas.Refresh()
	pxCanvas.MouseState.PreviousCoord = &ev.PointEvent
}

func (pxCanvas *PxCanvas) MouseIn(ev *desktop.MouseEvent) {}

func (pxCanvas *PxCanvas) MouseOut() {}

func (pxCanvas *PxCanvas) MouseDown(ev *desktop.MouseEvent) {
	brush.TryBrush(pxCanvas.AppState, pxCanvas, ev)
}

func (pxCanvas *PxCanvas) MouseUp(ev *desktop.MouseEvent) {}
