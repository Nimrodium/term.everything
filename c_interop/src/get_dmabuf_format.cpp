#include "get_dmabuf_format.h"
#include "gpu.h"
#include "OffScreenRenderer.h"
#include <stdexcept>
/**
 * Make a fake dmabuf image to find out what image format we support.
 */
Value get_dmabuf_format_js(const CallbackInfo &info)
{
  auto env = info.Env();
  auto offscreen_renderer = info[0].As<External<OffScreenRenderer>>().Data();

  GLuint texture;
  glGenTextures(1, &texture);
  glBindTexture(GL_TEXTURE_2D, texture);
  /**
   * @TODO any harm in hardcoding texture data like this?
   */
  const int TEXTURE_DATA_WIDTH = 256;
  const int TEXTURE_DATA_HEIGHT = 256;
  const size_t TEXTURE_DATA_SIZE = TEXTURE_DATA_WIDTH * TEXTURE_DATA_HEIGHT;
  auto texture_data = new uint8_t[TEXTURE_DATA_SIZE * 4];
  glTexImage2D(GL_TEXTURE_2D, 0, GL_RGB, TEXTURE_DATA_WIDTH, TEXTURE_DATA_HEIGHT, 0, GL_RGBA, GL_UNSIGNED_BYTE, NULL);
  glTexSubImage2D(GL_TEXTURE_2D, 0, 0, 0, TEXTURE_DATA_WIDTH, TEXTURE_DATA_HEIGHT, GL_RGBA, GL_UNSIGNED_BYTE, texture_data);
  glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MIN_FILTER, GL_NEAREST);
  glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MAG_FILTER, GL_NEAREST);

  // EGL: Create EGL image from the GL texture
  auto image = eglCreateImage(offscreen_renderer->display,
                              offscreen_renderer->context,
                              EGL_GL_TEXTURE_2D,
                              (EGLClientBuffer)(uint64_t)texture,
                              NULL);
  if (image == EGL_NO_IMAGE)
  {
    throw std::runtime_error("Failed to create EGL image from texture");
  }

  int fourcc;
  EGLuint64KHR modifiers;
  int num_planes;

  auto queried = eglExportDMABUFImageQueryMESA2(offscreen_renderer->display,
                                                     image,
                                                     &fourcc,
                                                     &num_planes,
                                                     &modifiers);
  if (!queried)
  {
    throw std::runtime_error("Failed to query dmabuf image format");
  }
  eglDestroyImage(offscreen_renderer->display, image);
  glDeleteTextures(1, &texture);
  delete[] texture_data;

  auto out =  Object::New(env);
  out.Set("fourcc", Value::From(env, fourcc));
  out.Set("modifier_hi", Value::From(env, (uint32_t)(modifiers >> 32)));
  out.Set("modifier_lo", Value::From(env, (uint32_t)(modifiers & 0xFFFFFFFF)));
  return out;
}
