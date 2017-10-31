package texture_test

import (
	"lengine/resources/texture"
	"lengine/sdlHelper"
	"testing"
)

func TestLoadTexture(t *testing.T) {
	sdlHelper.Window()

	tex1 := texture.Get("testImage.png")

	if tex1.ID() == 0 {
		t.Error("GLID is 0 for tex1")
	}

	tex2 := texture.Get("testImage2.png")
	if tex2.ID() == 0 {
		t.Error("GLID is 0 for tex2")
	}

	if tex2.ID() == tex1.ID() {
		t.Error("Tex1 and 2 have the same GL-ID!")
	}
	if tex2.Data() == tex1.Data() {
		t.Error("Tex1 and 2 have the same data pointer!")
	}
}
