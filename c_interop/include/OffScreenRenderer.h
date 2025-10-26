#pragma once

#include "gpu.h"

/**
 * Renders to an off-screen buffer using OpenGL and EGL.
 */
class OffScreenRenderer
{
public:

    EGLDisplay display;
    EGLSurface surface;
    EGLContext context;

    OffScreenRenderer();
    ~OffScreenRenderer();
};