package graphics

import (
	"github.com/galaco/gosigl"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

func Init() error {
	return gl.Init()
}

func Viewport(x, y, width, height int32) {
	gl.Viewport(x, y, width, height)
}

func ClearColor(r, g, b, a float32) {
	gl.ClearColor(r, g, b, a)
}

func ClearAll() {
	Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func Clear(mask uint32) {
	gl.Clear(mask)
}

func UploadTexture(texture *Texture) uint32 {
	return uint32(gosigl.CreateTexture2D(
		gosigl.TextureSlot(0),
		texture.Width(),
		texture.Height(),
		texture.colour,
		gosigl.PixelFormat(texture.Format()),
		false))
}

func BindTexture(id uint32) {
	gosigl.BindTexture2D(gosigl.TextureSlot(0), gosigl.TextureBindingId(id))
}

// textureFormatFromVtfFormat swap vtf format to openGL format
func textureFormatFromVtfFormat(vtfFormat uint32) uint32 {
	switch vtfFormat {
	case 0:
		return gl.RGBA
	case 2:
		return gl.RGB
	case 3:
		return gl.BGR
	case 12:
		return gl.BGRA
	case 13:
		return gl.COMPRESSED_RGB_S3TC_DXT1_EXT
	case 14:
		return gl.COMPRESSED_RGBA_S3TC_DXT3_EXT
	case 15:
		return gl.COMPRESSED_RGBA_S3TC_DXT5_EXT
	default:
		return gl.RGB
	}
}


type GpuMesh *gosigl.VertexObject

func UploadMesh(mesh *Mesh) GpuMesh {
	gpuResource  := gosigl.NewMesh(mesh.Vertices())
	gosigl.CreateVertexAttribute(gpuResource, mesh.UVs(), 2)
	gosigl.CreateVertexAttribute(gpuResource, mesh.Normals(), 3)
	gosigl.CreateVertexAttribute(gpuResource, mesh.Tangents(), 4)
	gosigl.FinishMesh()

	return GpuMesh(gpuResource)
}

func DrawArray(offset int, num int) {
	gosigl.DrawArray(offset, num)
}

func DrawFace(offset int, num int, textureId uint32) {
	BindTexture(textureId)
	DrawArray(offset, num)
}

func BindMesh(mesh *GpuMesh) {
	gosigl.BindMesh(*mesh)
}

func PushMat4(uniform int32, count int, transpose bool, mat mgl32.Mat4) {
	gl.UniformMatrix4fv(uniform, int32(count), transpose, &mat[0])
}