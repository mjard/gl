#include <GL/glew.h>
#include "gl.h"

void go_glDrawRangeElements(GLenum mode, GLuint start, GLuint end, GLsizei count, GLenum type, const GLvoid *indices) {
    glDrawRangeElements(mode, start, end, count, type, indices);
}

void go_glPrimitiveRestartIndex(GLuint index) {
    glPrimitiveRestartIndex(index);
}
