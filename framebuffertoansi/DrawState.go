package framebuffertoansi

type DrawState interface {
	DrawDesktop(texturePixels []byte, width, height uint32, statusLine *string) (int, int)
}
