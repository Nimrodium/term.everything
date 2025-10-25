import { auto_release } from "../auto_release.ts";
import {
  zwp_linux_dmabuf_feedback_v1_delegate as d,
  zwp_linux_dmabuf_feedback_v1 as w,
  wl_surface
} from "../protocols/wayland.xml.ts";
import { Wayland_Client } from "../Wayland_Client.ts";
import { Object_ID } from "../wayland_types.ts";

export class zwp_linux_dmabuf_feedback_v1 implements d {
  zwp_linux_dmabuf_feedback_v1_destroy: d["zwp_linux_dmabuf_feedback_v1_destroy"] = auto_release;
  zwp_linux_dmabuf_feedback_v1_on_bind: d["zwp_linux_dmabuf_feedback_v1_on_bind"] =
    (s, name, interface_, new_id, version_number) => {
    };
  constructor(public surface: Object_ID<wl_surface> | null) {
  }
  static make(surface_id: Object_ID<wl_surface> | null): w {
    return new w(new zwp_linux_dmabuf_feedback_v1(surface_id));
  }
}
