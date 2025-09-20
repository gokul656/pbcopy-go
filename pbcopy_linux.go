//go:build freebsd || linux || netbsd || openbsd || solaris || dragonfly

package pbcopy

/*
#cgo LDFLAGS: -lX11
#include <X11/Xlib.h>
#include <X11/Xatom.h>
#include <stdlib.h>
#include <string.h>

Display *d;
Window w;
char *clipboardText;

char* getClipboardText(Display *d, Window w) {
	Atom clip = XInternAtom(d, "CLIPBOARD", False);
	Atom utf8 = XInternAtom(d, "UTF8_STRING", False);
	Atom prop = XInternAtom(d, "XSEL_DATA", False);

	XConvertSelection(d, clip, utf8, prop, w, CurrentTime);

	XEvent event;
	XNextEvent(d, &event);

	if (event.type == SelectionNotify) {
		if (event.xselection.property) {
			Atom actualType;
			int actualFormat;
			unsigned long nitems, bytesAfter;
			unsigned char *propRet = NULL;

			XGetWindowProperty(d, w, prop, 0, (~0L), False,
								AnyPropertyType, &actualType, &actualFormat,
								&nitems, &bytesAfter, &propRet);

			if (propRet) {
				char *result = strdup((char*)propRet);
				XFree(propRet);
				return result;
			}
		}
	}
	return NULL;
}

void serveClipboard(Display *d, Window w, char *clipboardText) {
    Atom clipboardAtom = XInternAtom(d, "CLIPBOARD", False);
    XSetSelectionOwner(d, clipboardAtom, w, CurrentTime);

    if (XGetSelectionOwner(d, clipboardAtom) == 0) {
        return;
    }

    XEvent e;
    for (;;) {
        XNextEvent(d, &e);
        if (e.type == SelectionRequest) {
            XSelectionRequestEvent *req = &e.xselectionrequest;
            XEvent respond;
            memset(&respond, 0, sizeof(XEvent));
            respond.xselection.type = SelectionNotify;
            respond.xselection.display = req->display;
            respond.xselection.requestor = req->requestor;
            respond.xselection.selection = req->selection;
            respond.xselection.target = req->target;
            respond.xselection.property = req->property;
            respond.xselection.time = req->time;

            if (req->target == XA_STRING || req->target == XInternAtom(d, "UTF8_STRING", False)) {
                XChangeProperty(d, req->requestor, req->property, req->target, 8,
                                PropModeReplace, (unsigned char*)clipboardText, strlen(clipboardText));
            } else {
                respond.xselection.property = None;
            }

            XSendEvent(d, req->requestor, False, 0, &respond);
            XFlush(d);
        }
    }
}

*/
import "C"
import (
	"fmt"
	"unsafe"
)

func readAll() (string, error) {
	d := C.XOpenDisplay(nil)
	if d == nil {
		panic("Cannot open X11 display")
	}
	defer C.XCloseDisplay(d)

	w := C.XCreateSimpleWindow(d, C.XDefaultRootWindow(d), 0, 0, 1, 1, 0, 0, 0)

	txt := C.getClipboardText(d, w)
	if txt != nil {
		fmt.Println("Clipboard:", C.GoString(txt))
		C.free(unsafe.Pointer(txt))
	} else {
		fmt.Println("No clipboard data")
	}
	return "", nil
}

func write(data []byte) (int, error) {
	d := C.XOpenDisplay(nil)
	if d == nil {
		return 0, fmt.Errorf("Cannot open X11 display")
	}
	defer C.XCloseDisplay(d)

	w := C.XCreateSimpleWindow(d, C.XDefaultRootWindow(d), 0, 0, 1, 1, 0, 0, 0)
	text := C.CString(string(data))
	defer C.free(unsafe.Pointer(text))

	clipboardAtom := C.XInternAtom(d, C.CString("CLIPBOARD"), C.False)
	C.XSetSelectionOwner(d, clipboardAtom, w, C.CurrentTime)

	owner := C.XGetSelectionOwner(d, clipboardAtom)
	if owner == 0 {
		return 0, fmt.Errorf("Failed to set clipboard owner")
	}

	fmt.Println("Clipboard set. Press Ctrl+C to quit.")

	C.serveClipboard(d, w, text)
	return 0, nil
}
