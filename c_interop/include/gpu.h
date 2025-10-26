#pragma once

#define GL_GLEXT_PROTOTYPES
#include <GL/gl.h>
#define EGL_EGLEXT_PROTOTYPES
#include <EGL/egl.h>
#include <EGL/eglext.h>


extern PFNEGLEXPORTDMABUFIMAGEQUERYMESAPROC eglExportDMABUFImageQueryMESA2;
extern PFNEGLEXPORTDMABUFIMAGEMESAPROC eglExportDMABUFImageMESA2;