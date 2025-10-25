import { Global_Ids } from "../GlobalObjects.ts";
import {
  wl_registry_delegate as d,
  wl_registry as w,
} from "../protocols/wayland.xml.ts";
import { Object_ID, version } from "../wayland_types.ts";
import { wayland_debug_time_only } from "../debug.ts" with { type: "macro" };

export class wl_registry implements d {
  wl_registry_bind: d["wl_registry_bind"] = (
    s,
    _object_id,
    name,
    id_interface,
    id_version,
    id_id
  ) => {
    const object = s.get_object(name as Object_ID);
    s.add_object(id_id, object);
    const version: version = id_version as version;

    switch (name) {
      case Global_Ids.wl_shm:
      case Global_Ids.wl_seat:
      case Global_Ids.wl_output:
      case Global_Ids.wl_keyboard:
      case Global_Ids.wl_pointer:
      case Global_Ids.wl_touch:
      case Global_Ids.wl_data_device:
      case Global_Ids.zwp_xwayland_keyboard_grab_manager_v1:
      case Global_Ids.zwp_linux_dmabuf_v1:
        s.add_global_bind(name, id_id, version);
        // const set = s.global_binds.get(name) ?? new Set();
        // set.add(id_id);
        // s.global_binds.set(name, set);
        break;
    }
    try {
      if (wayland_debug_time_only()) {
        console.log(
          `client#${s.client_socket} ${id_interface}_on_bind(name:`,
          name,
          ",interface:",
          id_interface,
          ",id:",
          id_id,
          ",version:",
          id_version,
          `)`
        );
      }

      (object?.delegate as any)[`${id_interface}_on_bind`](
        s,
        name,
        id_interface,
        id_id,
        version
      );
    } catch (e) {
      debugger;
      console.log(
        `error in wl_registry_bind. name, ${name}, id_interface, ${id_interface}
    id_version, ${id_version}, id_id, ${id_id}`,
        e
      );
    }
  };
  wl_registry_on_bind: d["wl_registry_on_bind"] = (
    _s,
    _name,
    _id,
    _version
  ) => {};
  constructor() {}
  static make(): w {
    return new w(new wl_registry());
  }
}
