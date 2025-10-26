#include "make_offscreen_renderer.h"
#include "OffScreenRenderer.h"

  Value make_offscreen_renderer_js(const CallbackInfo &info){
    auto env = info.Env();
    auto renderer = External<OffScreenRenderer>::New(
      env, new OffScreenRenderer(),
      [](Napi::Env env, OffScreenRenderer *data)
      { delete data; });
    return renderer;
  }
  
