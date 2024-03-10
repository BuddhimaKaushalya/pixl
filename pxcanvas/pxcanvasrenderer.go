package pxcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type PxCanvasRenderer struct {
	pxCanvas     *PxCanvas
	CanvasImage  *canvas.Image // Changed field name to CanvasImage
	CanvasBorder []canvas.Line // Changed field name to CanvasBorder
	CanvaCursor  []fyne.CanvasObject
}

func (renderer *PxCanvasRenderer) SetCursor(objects []fyne.CanvasObject) {
	renderer.CanvaCursor = objects
}

func (renderer *PxCanvasRenderer) MinSize() fyne.Size {
	return renderer.pxCanvas.DrawingArea
}

func (renderer *PxCanvasRenderer) Objects() []fyne.CanvasObject { // Changed method name to Objects
	objects := make([]fyne.CanvasObject, 0, 5)
	for i := 0; i < len(renderer.CanvasBorder); i++ {
		objects = append(objects, &renderer.CanvasBorder[i])
	}
	objects = append(objects, renderer.CanvasImage)
	objects = append(objects, renderer.CanvaCursor...)
	return objects
}

func (renderer *PxCanvasRenderer) Destroy() {

}

func (renderer *PxCanvasRenderer) Layout(size fyne.Size) {
	renderer.LayoutCanvas(size)
	renderer.LayoutBorder(size) // Corrected method name
}

func (renderer *PxCanvasRenderer) Refresh() {
	if renderer.pxCanvas.ReloadImage {
		renderer.CanvasImage = canvas.NewImageFromImage(renderer.pxCanvas.PixlData)
		renderer.CanvasImage.ScaleMode = canvas.ImageScalePixels
		renderer.CanvasImage.FillMode = canvas.ImageFillContain
		renderer.pxCanvas.ReloadImage = false
	}
	renderer.Layout(renderer.pxCanvas.Size())
	canvas.Refresh(renderer.CanvasImage)
}

func (renderer *PxCanvasRenderer) LayoutCanvas(size fyne.Size) {
	imgPxWidth := renderer.pxCanvas.PxCols
	imgPxHeight := renderer.pxCanvas.PxRows
	pxSize := renderer.pxCanvas.PxSize
	renderer.CanvasImage.Move(renderer.pxCanvas.CanvasOffset)
	renderer.CanvasImage.Resize(fyne.NewSize(float32(imgPxWidth*pxSize), float32(imgPxHeight*pxSize)))
}

func (renderer *PxCanvasRenderer) LayoutBorder(size fyne.Size) { // Added missing method
	offset := renderer.pxCanvas.CanvasOffset
	imgHeight := renderer.CanvasImage.Size().Height
	imgWidth := renderer.CanvasImage.Size().Width

	left := &renderer.CanvasBorder[0]
	left.Position1 = fyne.NewPos(float32(offset.X), float32(offset.Y))
	left.Position2 = fyne.NewPos(float32(offset.X), float32(offset.Y+imgHeight))

	top := &renderer.CanvasBorder[1]
	top.Position1 = fyne.NewPos(float32(offset.X), float32(offset.Y))
	top.Position2 = fyne.NewPos(float32(offset.X+imgWidth), float32(offset.Y))

	right := &renderer.CanvasBorder[2]
	right.Position1 = fyne.NewPos(float32(offset.X+imgWidth), float32(offset.Y))
	right.Position2 = fyne.NewPos(float32(offset.X+imgWidth), float32(offset.Y+imgHeight))

	bottom := &renderer.CanvasBorder[3]
	bottom.Position1 = fyne.NewPos(float32(offset.X), float32(offset.Y+imgHeight))
	bottom.Position2 = fyne.NewPos(float32(offset.X+imgWidth), float32(offset.Y+imgHeight))
}
