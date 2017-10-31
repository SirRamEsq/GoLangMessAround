package texture

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/png"
	"lengine/logger"
	"log"
	"os"
	"strconv"

	"github.com/go-gl/gl/v3.3-core/gl"
)

var (
	gBoundTexture uint32
)

//Texture represents an openGlTexture
type Texture struct {
	w int
	h int

	data *image.RGBA
	glID uint32
	name string
}

func (t *Texture) String() string {
	returnValue := "\r\nTexture " + t.name + ":"
	returnValue += "\r\n Width:  " + strconv.Itoa(t.w)
	returnValue += "\r\n Height: " + strconv.Itoa(t.h)
	returnValue += "\r\n GL-ID:  " + strconv.FormatUint(uint64(t.glID), 10)
	returnValue += fmt.Sprintf("\r\n Data*:  %p", t.data)
	returnValue += "\r\n___"
	return returnValue
}

//Width returns the texture width
func (t *Texture) Width() int {
	return t.w
}

//Height returns the texture height
func (t *Texture) Height() int {
	return t.h
}

//Data Returns a pointer to the data that makes up the Texture
func (t *Texture) Data() *image.RGBA {
	return t.data
}

//ID returns the texture's openglID
func (t *Texture) ID() uint32 {
	return t.glID
}

//Bind will set the texture to the currently bound opengl context
func (t *Texture) Bind() {
	id := t.ID()
	if id != gBoundTexture {
		gl.BindTexture(gl.TEXTURE_2D, id)
	}
}

//New Creates a new texture
func newTexture(fileName string) *Texture {
	imgFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("texture %q not found on disk: %v\n", fileName, err)
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		panic(err)
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		panic("unsupported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	var texID uint32
	gl.Enable(gl.TEXTURE_2D)
	gl.GenTextures(1, &texID)

	if texID == 0 {
		panic("Texture ID is 0!")
	}

	gl.BindTexture(gl.TEXTURE_2D, texID)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))

	texture := Texture{w: rgba.Rect.Size().X,
		h:    rgba.Rect.Size().Y,
		glID: texID,
		data: rgba,
		name: fileName}

	glError := gl.GetError()
	if glError != 0 {
		logger.Critical(fmt.Sprintf("GL Error: %d", glError))
	}
	return &texture
}

//////////////////
//TextureManager//
//////////////////
type textureManager struct {
	textures map[string]*Texture
}

//manager handles texture creation/storage/retreival
var manager textureManager

//Contains returns true if a texture is already loaded
func Contains(name string) bool {
	_, ok := manager.textures[name]
	return ok
}

//Get returns a texture from a string path or null if impossible to load
func Get(name string) *Texture {
	//If textue exists, return it
	tex, ok := manager.textures[name]
	if ok {
		return tex
	}

	//If textue doesn't exist , try to create it
	tex = newTexture(name)
	manager.textures[name] = tex

	//will be null if New() fails
	return tex
}

//init initializes the texture manager singleton
func init() {
	manager.textures = make(map[string]*Texture)
}
