#include "NODE_API_MODULE.h"

#include "Send_Message_And_File_Descriptors.h"

#include "Listen_for_New_Client.h"
#include "Get_Message_and_File_Descriptors.h"
#include "listen_to_wayland.h"
#include "mmap_fd.h"
#include "get_fd.h"
#include "init_draw_state.h"
#include "memcopy_buffer_to_uint8array.h"
#include "draw_desktop.h"
#include "close_wayland_socket.h"
#include "get_socket_path_from_name.h"
#include "get_dmabuf_format.h"
#include "make_offscreen_renderer.h"
//{NEW_INCLUDE} will be added here

Object Init(Env env, Object exports)
{
    exports["send_message_and_file_descriptors"] = Function::New(env, send_message_and_file_descriptors_js);
    exports["get_message_and_file_descriptors"] = Function::New(env, get_message_and_file_descriptors_js);
    exports["listen_for_client"] = Function::New(env, listen_for_client);
    exports["listen_to_wayland_socket"] = Function::New(env, listen_to_wayland_socket_js);
    exports["mmap_shm_pool"] = Function::New(env, mmap_shm_pool_js);
    exports["remap_shm_pool"] = Function::New(env, remap_shm_pool_js);
    exports["unmmap_shm_pool"] = Function::New(env, unmmap_shm_pool_js);
    exports["get_fd"] = Function::New(env, get_fd_js);
    exports["init_draw_state"] = Function::New(env, init_draw_state_js);
    exports["memcopy_buffer_to_uint8array"] = Function::New(env, memcopy_buffer_to_uint8array_js);
    exports["draw_desktop"] = Function::New(env, draw_desktop_js);
    exports["close_wayland_socket"] = Function::New(env, close_wayland_socket_js);
    exports["get_socket_path_from_name"] = Function::New(env, get_socket_path_from_name_js);
    exports["get_dmabuf_format"] = Function::New(env, get_dmabuf_format_js);
    exports["make_offscreen_renderer"] = Function::New(env, make_offscreen_renderer_js);
    //{NEW_FUNC} will be added here

    return exports;
}

NODE_API_MODULE(addon, Init)