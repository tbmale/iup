// A group of functions to draw in a IupCanvas or a IupBackgroundBox. 
// They are simple functions designed to help the drawing of custom controls based on these two controls. 
// It is NOT a complete set of drawing functions, for that you should still use another toolkit like CD. 
// Internally IupDraw uses several drawing APIs: GDI, Direct2D, GDI+ (in Windows); X11 (in Motif), GDK and Cairo (in GTK). 
// All drivers are double buffered, so drawing occurs off-screen and the final result is displayed when IupDrawEnd is called only.

package iupdraw

import (
	"unsafe"

	"github.com/tbmale/iup"
)
/*
#cgo CFLAGS: -I./../include
#cgo LDFLAGS: -L${SRCDIR}/../lib -L${SRCDIR}/../lib/cd -L${SRCDIR}/../lib/im
#cgo LDFLAGS: -liupcontrols -liupcd -liup
#cgo LDFLAGS: -lcd -lfreetype6 -lz
#cgo LDFLAGS: -lgdi32 -luser32 -lcomdlg32 -lcomctl32 -luuid -loleaut32 -lole32

#include <stdlib.h>

#include <iup.h>
#include <iupcontrols.h>
*/
import "C"


// void IupDrawBegin(Ihandle* ih);
func DrawBegin(ih Ihandle){
	C.IupDrawBegin(ih.ptr())
}

// void IupDrawEnd(Ihandle* ih);
func DrawEnd(ih Ihandle){
	C.IupDrawEnd(ih.ptr())
}

// /* all functions can be called only between calls to Begin and End */

// void IupDrawSetClipRect(Ihandle* ih, int x1, int y1, int x2, int y2);
func DrawSetClipRect(ih Ihandle, x1, y1, x2, y2 int){
	C.IupDrawSetClipRect(ih.ptr(),C.int(x1), C.int(y1), C.int(x2), C.int(y2))
}

// void IupDrawGetClipRect(Ihandle* ih, int *x1, int *y1, int *x2, int *y2);
func DrawGetClipRect(ih Ihandle) (x1, y1, x2, y2 int){
	C.IupDrawGetClipRect(ih.ptr(), (*C.int)(unsafe.Pointer(&x1)), (*C.int)(unsafe.Pointer(&y1)), (*C.int)(unsafe.Pointer(&x2)), (*C.int)(unsafe.Pointer(&y2)))
	return
}

// void IupDrawResetClip(Ihandle* ih);
func DrawResetClip(ih Ihandle){
	C.IupDrawResetClip(ih.ptr())
}

// /* color controlled by the attribute DRAWCOLOR */
// /* line style or fill controlled by the attribute DRAWSTYLE */

// void IupDrawParentBackground(Ihandle* ih);
func DrawParentBackground(ih Ihandle){
	C.IupDrawParentBackground(ih.ptr())
}

// void IupDrawLine(Ihandle* ih, int x1, int y1, int x2, int y2);
func DrawLine(ih Ihandle, x1, y1, x2, y2 int){
	C.IupDrawLine(ih.ptr(), C.int(x1), C.int(y1), C.int(x2), C.int(y2))
}

// void IupDrawRectangle(Ihandle* ih, int x1, int y1, int x2, int y2);
func DrawRectangle(ih Ihandle, x1, y1, x2, y2 int){
	C.IupDrawRectangle(ih.ptr(), C.int(x1), C.int(y1), C.int(x2), C.int(y2))
}

// void IupDrawArc(Ihandle* ih, int x1, int y1, int x2, int y2, double a1, double a2);
func DrawArc(ih Ihandle, x1, y1, x2, y2 int, ang1, ang2 ){
	C.IupDrawArc(ih.ptr(), C.int(x1), C.int(y1), C.int(x2), C.int(y2), C.double(ang1), C.double(ang2))
}

// void IupDrawPolygon(Ihandle* ih, int* points, int count);
// Draws a polygon. Coordinates are stored in the array in the sequence: x1, y1, x2, y2, ...
func DrawPolygon(ih Ihandle, points []int, count int){
	if points.len() != count{
		panic(fmt.Errorf("points array size different from count"))
	}
	C.IupDrawPolygon(ih.ptr(), (*C.int)(unsafe.Pointer(&points[0])), C.int(count))
}

// void IupDrawText(Ihandle* ih, const char* text, int len, int x, int y, int w, int h);
func DrawText(ih Ihandle, text string, length, x, y, w, h int){
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))
	if length != -1 && len(text) != length + 1{
		panic(fmt.Errorf("text length differs from lenght parameter and lenght parameter is not -1"))
	}
	C.IupDrawText(ih.ptr(), c_text, C.int(length), C.int(x), C.int(y), C.int(w), C.int(h))
}

// void IupDrawImage(Ihandle* ih, const char* name, int x, int y, int w, int h);
func DrawImage(ih Ihandle, name string, x, y, w, h int){
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	C.IupDrawImage(ih.ptr(), c_name, C.int(x), C.int(y), C.int(w), C.int(h))
}

// void IupDrawSelectRect(Ihandle* ih, int x1, int y1, int x2, int y2);
func DrawSelectRect(ih Ihandle, x1, y1, x2, y2 int){
	C.IupDrawSelectRect(ih.ptr(), C.int(x1), C.int(y1), C.int(x2), C.int(y2))
}

// void IupDrawFocusRect(Ihandle* ih, int x1, int y1, int x2, int y2);
func DrawFocusRect(ih Ihandle, x1, y1, x2, y2 int){
	C.IupDrawFocusRect(ih.ptr(), C.int(x1), C.int(y1), C.int(x2), C.int(y2))
}

// void IupDrawGetSize(Ihandle* ih, int *w, int *h);
func DrawGetSize(ih Ihandle) (w, h int){
	C.IupDrawGetSize(ih.ptr(), (*C.int)(unsafe.Pointer(&w)), (*C.int)(unsafe.Pointer(&h)))
	return
}

// void IupDrawGetTextSize(Ihandle* ih, const char* text, int len, int *w, int *h);
func DrawGetTextSize(ih Ihandle, text string, length) (w, h int){
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))
	if length != -1 && len(text) != length + 1{
		panic(fmt.Errorf("text length differs from lenght parameter and lenght parameter is not -1"))
	}
	C.IupDrawGetTextSize(ih.ptr(), c_text, C.int(length), (*C.int)(unsafe.Pointer(&w)), (*C.int)(unsafe.Pointer(&h)))
	return
}

// void IupDrawGetImageInfo(const char* name, int *w, int *h, int *bpp);
func DrawGetImageInfo(name string) (w, h, bpp int){
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	C.IupDrawGetImageInfo(ih.ptr(), c_name, (*C.int)(unsafe.Pointer(&w)), (*C.int)(unsafe.Pointer(&h)), (*C.int)(unsafe.Pointer(&bpp)))
	return
}
