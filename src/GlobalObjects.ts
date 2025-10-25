import { wl_display } from "./objects/wl_display.ts";
import { wl_output } from "./objects/wl_output.ts";
import { wl_seat } from "./objects/wl_seat.ts";
import { wl_shm } from "./objects/wl_shm.ts";
import { wl_keyboard } from "./objects/wl_keyboard.ts";
import { wl_compositor } from "./objects/wl_compositor.ts";
import { zxdg_decoration_manager_v1 } from "./objects/zxdg_decoration_manager_v1.ts";
import { wl_pointer } from "./protocols/wayland.xml.ts";
import { wl_subcompositor } from "./objects/wl_subcompositor.ts";
import { xdg_wm_base } from "./objects/xdg_wm_base.ts";
import { wl_data_device_manager } from "./objects/wl_data_device_manager.ts";
import { pointer } from "./objects/wl_pointer.ts";
import { zwp_xwayland_keyboard_grab_manager_v1 } from "./objects/zwp_xwayland_keyboard_grab_manager_v1.ts";
import { xwayland_shell_v1 } from "./objects/xwayland_shell_v1.ts";
import { wl_touch } from "./objects/wl_touch.ts";
import { zwp_linux_dmabuf_v1 } from "./objects/zwp_linux_dmabuf_v1.ts";
export enum Global_Ids {
  wl_display = 1,
  wl_compositor = 0xff00_000,
  wl_subcompositor,
  wl_output,
  wl_seat,
  wl_shm,
  xdg_wm_base,
  wl_data_device_manager,
  wl_keyboard,
  wl_pointer,
  zwp_xwayland_keyboard_grab_manager_v1,
  xwayland_shell_v1,
  wl_data_device,
  wl_touch,
  zxdg_decoration_manager_v1,
  zwp_linux_dmabuf_v1,
}
const seat = wl_seat.make();
const globals = {
  [1]: wl_display.make(),
  [Global_Ids.wl_output]: wl_output.make(),
  [Global_Ids.wl_seat]: seat,
  [Global_Ids.wl_shm]: wl_shm.make(),
  [Global_Ids.wl_compositor]: wl_compositor.make(),
  [Global_Ids.wl_subcompositor]: wl_subcompositor.make(),
  [Global_Ids.xdg_wm_base]: xdg_wm_base.make(),
  [Global_Ids.wl_data_device_manager]: wl_data_device_manager.make(),
  [Global_Ids.wl_keyboard]: wl_keyboard.make(),
  [Global_Ids.wl_pointer]: new wl_pointer(pointer),
  [Global_Ids.zwp_xwayland_keyboard_grab_manager_v1]:
    zwp_xwayland_keyboard_grab_manager_v1.make(),
  [Global_Ids.xwayland_shell_v1]: xwayland_shell_v1.make(),
  [Global_Ids.wl_data_device]: seat,
  [Global_Ids.wl_touch]: wl_touch.make(),
  [Global_Ids.zxdg_decoration_manager_v1]: zxdg_decoration_manager_v1.make(),
  [Global_Ids.zwp_linux_dmabuf_v1]: zwp_linux_dmabuf_v1.make(),
};

export class GlobalObjects {
  objects: typeof globals & {
    [key: number]: (typeof globals)[keyof typeof globals] | undefined;
  } = globals;

  constructor() {}
}

export const advertised_global_objects_names = [
  {
    name: "wl_compositor",
    id: Global_Ids.wl_compositor,
    version: 6,
  },
  /**
   * Turning off the wl_subcompositor will turn off
   * decorations. Any other side effects??? Looks like
   * GameScope has it turned off, so maybe we could do that
   * too.
   *
   * some programs will crash if wl_subcompositor is not
   * advertised.
   */
  { name: "wl_subcompositor", id: Global_Ids.wl_subcompositor, version: 1 },
  { name: "wl_output", id: Global_Ids.wl_output, version: 5 },

  { name: "wl_seat", id: Global_Ids.wl_seat, version: 10 },
  { name: "wl_shm", id: Global_Ids.wl_shm, version: 2 },
  { name: "xdg_wm_base", id: Global_Ids.xdg_wm_base, version: 6 },
  {
    name: "wl_data_device_manager",
    id: Global_Ids.wl_data_device_manager,
    version: 3,
  },

  {
    name: "zxdg_decoration_manager_v1",
    id: Global_Ids.zxdg_decoration_manager_v1,
    version: 1,
  },
  /**
   * @TODO only advertise these to Xwayland clients
   */
  {
    name: "zwp_xwayland_keyboard_grab_manager_v1",
    id: Global_Ids.zwp_xwayland_keyboard_grab_manager_v1,
    version: 1,
  },
  {
    name: "xwayland_shell_v1",
    id: Global_Ids.xwayland_shell_v1,
    version: 1,
  },
  /**
   * @TODO should this be advertised?
   */
  {
    name: "zwp_linux_dmabuf_v1",
    id: Global_Ids.zwp_linux_dmabuf_v1,
    /**
     * @TODO Latest version is 5, but I don't want to implement 
     * linux_dmabuf_feedback_v1 just yet. So use version 3
     * which doesn't require it.
     */
    version: 3,
  }
];

export const global_objects = new GlobalObjects();
