package sfml

//#include <SFML/Graphics.h>
import "C"

type Rectf struct {
	Left   float32
	Top    float32
	Width  float32
	Height float32
}

type Recti struct {
	Left   int
	Top    int
	Width  int
	Height int
}

func (r Recti) IsEqual(rect Recti) bool {
	return r.Left == rect.Left && r.Top == rect.Top &&
		r.Width == rect.Width && r.Height == rect.Height
}

func (r Rectf) IsEqual(rect Rectf) bool {
	return r.Left == rect.Left && r.Top == rect.Top &&
		r.Width == rect.Width && r.Height == rect.Height
}

func (f Rectf) Contains(point Vector2f) bool {
	rect := cRectf(&f)
	return goBool(C.sfFloatRect_contains(&rect, C.float(point.X), C.float(point.Y)))
}

func (i Recti) Contains(point Vector2i) bool {
	rect := cRecti(&i)
	return goBool(C.sfIntRect_contains(&rect, C.int(point.X), C.int(point.Y)))
}

func (f Rectf) Intersects(other Rectf) (bool, Rectf) {
	var rect3 C.sfFloatRect
	rect1 := cRectf(&f)
	rect2 := cRectf(&other)
	r := goBool(C.sfFloatRect_intersects(&rect1, &rect2, &rect3))
	return r, *goRectf(&rect3)
}

func (i Recti) Intersects(other Recti) (bool, Recti) {
	var rect3 C.sfIntRect
	rect1 := cRecti(&i)
	rect2 := cRecti(&other)
	r := goBool(C.sfIntRect_intersects(&rect1, &rect2, &rect3))
	return r, *goRecti(&rect3)
}

func goRectf(f *C.sfFloatRect) *Rectf {
	return &Rectf{
		float32(f.left),
		float32(f.top),
		float32(f.width),
		float32(f.height),
	}
}

func cRectf(f *Rectf) C.sfFloatRect {
	return C.sfFloatRect{
		C.float(f.Left),
		C.float(f.Top),
		C.float(f.Width),
		C.float(f.Height),
	}
}

func goRecti(i *C.sfIntRect) *Recti {
	return &Recti{
		int(i.left),
		int(i.top),
		int(i.width),
		int(i.height),
	}
}

func cRecti(i *Recti) C.sfIntRect {
	return C.sfIntRect{
		C.int(i.Left),
		C.int(i.Top),
		C.int(i.Width),
		C.int(i.Height),
	}
}
