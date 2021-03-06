package utils

import (
	"github.com/galaco/kero/framework/graphics"
	"image"
	"image/png"
	"os"
)

// DumpLightmap exports the loaded lightmap texture atlas as a JPG
func DumpLightmap(name string, im graphics.Texture) {
	img := image.NewRGBA(image.Rect(0, 0, im.Width(), im.Height()))
	copy(img.Pix, im.Image())

	outfile, _ := os.Create("./" + name + ".jpg")
	defer outfile.Close()
	png.Encode(outfile, img)
}
