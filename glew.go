package gl
// #cgo darwin LDFLAGS: -framework OpenGL -lGLEW
// #cgo windows LDFLAGS: -lglew32 -lopengl32
// #cgo linux LDFLAGS: -lGLEW -lGL
// #include "glew_stub.h"
import "C"

func PrimitiveRestartIndex(index uint) {
    C.go_glPrimitiveRestartIndex(C.GLuint(index))
}

func DrawRangeElements(mode GLenum, start uint, end uint, count int, typ GLenum, indices interface{}) {
    C.go_glDrawRangeElements(C.GLenum(mode), C.GLuint(start), C.GLuint(end), C.GLsizei(count), C.GLenum(typ), ptr(indices));
}

