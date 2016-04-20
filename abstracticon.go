package abstracticon

import (
	"crypto"
	"hash"
	"image"
	"image/color"
	"image/draw"

	// Include the default hash algorithm.
	_ "crypto/md5"
)

var (
	// DefaultHash is used when a hash algorithm is not defined.
	DefaultHash = crypto.MD5
)

// Attrs controls how an abstracticon is rendered.
type Attrs struct {
	// Hash is the hash algorithm to be used to generate the image.
	Hash crypto.Hash
	// Multiplier defines how many pixels is used to represent a point.
	Multiplier int
	// Points defines the image dimension.
	Points int
	// Transparent defines a Transparent background.
	Transparent bool
	// If false the image will be mirrored from left to right.
	NotMirrored bool
}

func (a *Attrs) bounds() image.Rectangle {
	s := a.Points * a.Multiplier
	return image.Rect(0, 0, s, s)
}

// Render renders a image from a string.
func Render(seed string, attrs ...Attrs) image.Image {
	return RenderFromBytes([]byte(seed), attrs...)
}

// RenderFromBytes renders a image from a byte slice.
func RenderFromBytes(seed []byte, attrs ...Attrs) image.Image {
	a := Attrs{}
	if attrs != nil {
		a = attrs[len(attrs)-1]
	}

	c, bits := a.gen(seed)
	img := image.NewRGBA(a.bounds())
	if !a.Transparent {
		u := image.Uniform{color.RGBA{255, 255, 255, 255}}
		draw.Draw(img, img.Bounds(), &u, image.ZP, draw.Src)
	}

	draw := func(x, y int) {
		s := a.Multiplier
		rect := image.Rect(x*s, y*s, (x+1)*s, (y+1)*s)
		draw.Draw(img, rect, &image.Uniform{c}, image.ZP, draw.Src)
	}

	X, Y := a.Points, a.Points
	if !a.NotMirrored {
		X = (X + 1) / 2
	}

	for x := 0; x < X; x++ {
		for y := 0; y < Y; y++ {
			if bits.next() {
				draw(x, y)
				if !a.NotMirrored {
					draw(a.Points-x-1, y)
				}
			}
		}
	}
	return img
}

func (a *Attrs) gen(data []byte) (color.Color, bits) {
	// Get hash function
	var fn hash.Hash
	if a.Hash == 0 {
		fn = DefaultHash.New()
	} else {
		fn = a.Hash.New()
	}
	// Hash the data
	fn.Write(data)
	bytes := fn.Sum(nil)
	if len(bytes) < 3 {
		bytes = append(bytes, make([]byte, 3-len(bytes))...)
	}
	// Return a color a bit sequence
	c := color.RGBA{0x7f & bytes[0], 0x7f & bytes[1], 0x7f & bytes[2], 255}
	return c, bits{bytes: bytes[3:]}
}

type bits struct {
	bytes  []byte
	offset int
}

func (b *bits) next() bool {
	bitIdx := b.offset % 8
	byteIdx := (b.offset / 8) % len(b.bytes)
	b.offset++
	return b.bytes[byteIdx]&(1<<uint(bitIdx)) != 0
}
