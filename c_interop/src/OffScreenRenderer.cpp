#include "OffScreenRenderer.h"

#include <stdexcept>
OffScreenRenderer::OffScreenRenderer() {
    // Constructor implementation
    eglBindAPI(EGL_OPENGL_API);
    EGLDisplay display = eglGetDisplay(EGL_DEFAULT_DISPLAY);
    if (display == EGL_NO_DISPLAY) {
      throw std::runtime_error("Failed to get EGL display");
    }
    if (!eglInitialize(display, nullptr, nullptr)) {
      throw std::runtime_error("Failed to initialize EGL");
    }
    EGLint config_attribs[] = {
      EGL_SURFACE_TYPE, EGL_PBUFFER_BIT,
      EGL_BLUE_SIZE, 8,
      EGL_GREEN_SIZE, 8,
      EGL_RED_SIZE, 8,
      EGL_ALPHA_SIZE, 8,
      EGL_RENDERABLE_TYPE, EGL_OPENGL_BIT,
      EGL_NONE
    };
    EGLConfig config;
    EGLint num_configs;
    if (!eglChooseConfig(display, config_attribs, &config, 1, &num_configs) || num_configs == 0) {
      throw std::runtime_error("Failed to choose EGL config");
    }
    EGLint pbuffer_attribs[] = {
      EGL_WIDTH, 800,
      EGL_HEIGHT, 600,
      EGL_NONE,
    };
    EGLSurface surface = eglCreatePbufferSurface(display, config, pbuffer_attribs);
    if (surface == EGL_NO_SURFACE) {
      throw std::runtime_error("Failed to create EGL pbuffer surface");
    }

    EGLContext context = eglCreateContext(display, config, EGL_NO_CONTEXT, nullptr);
    if (context == EGL_NO_CONTEXT) {
      throw std::runtime_error("Failed to create EGL context");
    }
    if (!eglMakeCurrent(display, surface, surface, context)) {
      throw std::runtime_error("Failed to make EGL context current");
    }

    this->display = display;
    this->surface = surface;
    this->context = context;


}

OffScreenRenderer::~OffScreenRenderer() {
    // Destructor implementation
    eglMakeCurrent(display, EGL_NO_SURFACE, EGL_NO_SURFACE, EGL_NO_CONTEXT);
    eglDestroyContext(display, context);
    eglDestroySurface(display, surface);
    eglTerminate(display);
}