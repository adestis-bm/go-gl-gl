// This file implements GlowGetProcAddress for each platform. The correct
// version is chosen automatically based on build tags:
// windows: WGL
// darwin: CGL
// linux: GLX
// Use of EGL instead of the platform's default (listed above) is made possible
// via the "egl" build tag.
// It is also possible to install your own function outside this package for
// retrieving OpenGL function pointers, to do this see InitWithProcAddrFunc.
package gl

/*
#cgo windows CFLAGS: -DTAG_WINDOWS
#cgo windows LDFLAGS: -lopengl32
#cgo darwin CFLAGS: -DTAG_DARWIN
#cgo darwin LDFLAGS: -framework OpenGL
#cgo linux CFLAGS: -DTAG_LINUX
#cgo linux LDFLAGS: -lGL
#cgo egl CFLAGS: -DTAG_EGL
#cgo egl LDFLAGS: -lEGL
#if defined(TAG_WINDOWS)
	#define WIN32_LEAN_AND_MEAN 1
	#include <windows.h>
	#include <stdlib.h>
	static HMODULE ogl32dll = NULL;
	void* GlowGetProcAddress(const char* name) {
		void* pf = wglGetProcAddress((LPCSTR)name);
		if(pf) {
			return pf;
		}
		if(ogl32dll == NULL) {
			ogl32dll = LoadLibraryA("opengl32.dll");
		}
		return GetProcAddress(ogl32dll, (LPCSTR)name);
	}
#elif defined(TAG_DARWIN)
	#include <stdlib.h>
	#include <dlfcn.h>
	void* GlowGetProcAddress(const char* name) {
		return dlsym(RTLD_DEFAULT, name);
	}
#elif defined(TAG_LINUX)
	#include <stdlib.h>
	#include <GL/glx.h>
	void* GlowGetProcAddress(const char* name) {
		return glXGetProcAddress(name);
	}
#elif defined(TAG_EGL)
	#include <stdlib.h>
	#include <EGL/egl.h>
	void* GlowGetProcAddress(const char* name) {
		return eglGetProcAddress(name);
	}
#endif
*/
import "C"
import "unsafe"

func getProcAddress(namea string) unsafe.Pointer {
	cname := C.CString(namea)
	defer C.free(unsafe.Pointer(cname))
	return C.GlowGetProcAddress(cname)
}
