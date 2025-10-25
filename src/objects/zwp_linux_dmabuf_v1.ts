import { auto_release } from "../auto_release.ts";
import {
  zwp_linux_dmabuf_v1_delegate as d,
  zwp_linux_dmabuf_v1 as w,
  zwp_linux_dmabuf_feedback_v1 as feedback_w,
  wl_surface as wl_surface_w
  
} from "../protocols/wayland.xml.ts";
import { Wayland_Client } from "../Wayland_Client.ts";

import { Object_ID } from "../wayland_types.ts";
import { zwp_linux_buffer_params_v1 } from "./zwp_linux_buffer_params_v1.ts";
import { zwp_linux_dmabuf_feedback_v1 } from "./zwp_linux_dmabuf_feedback_v1.ts";

export class zwp_linux_dmabuf_v1 implements d {
  /**
   * @TODO should this be auto_release?
   */
  zwp_linux_dmabuf_v1_destroy: d["zwp_linux_dmabuf_v1_destroy"] = auto_release;
  zwp_linux_dmabuf_v1_create_params: d["zwp_linux_dmabuf_v1_create_params"] = (
    s,
    object_id,
    params_id
  ) => {
    s.add_object(params_id, zwp_linux_buffer_params_v1.make());
  };
  zwp_linux_dmabuf_v1_get_default_feedback: d["zwp_linux_dmabuf_v1_get_default_feedback"] =
    (s, object_id, id) => {
      s.add_object(id, zwp_linux_dmabuf_feedback_v1.make(null));
      this.send_feedback(s, id, null);
    };
  zwp_linux_dmabuf_v1_get_surface_feedback: d["zwp_linux_dmabuf_v1_get_surface_feedback"] =
    (s, object_id, id, surface) => {
      s.add_object(id, zwp_linux_dmabuf_feedback_v1.make(surface));
      this.send_feedback(s, id, surface);
    };

  send_feedback = (s: Wayland_Client, feedback: Object_ID<feedback_w>, surface: Object_ID<wl_surface_w> | null) => {
    /**
     * @TODO Implement proper feedback based on the surface
     */
    feedback_w.done(s, feedback);
  }
  zwp_linux_dmabuf_v1_on_bind: d["zwp_linux_dmabuf_v1_on_bind"] = (
    s,
    name,
    interface_,
    new_id,
    version_number
  ) => {
    console.log("On bind zwp_linux_dmabuf_v1 version", version_number);
    /**
     * In version 4 and above the compositor must not send the format event on
     * bind, but in earlier versions it should.
     */
    if (version_number >= 4) {
      return;
    }
    // Example formats, more can be added as needed
    /**
     * @TODO actually check what formats the system supports
     */
    const formats = [
      0x34325241, // 'AR24' - ARGB8888
      0x34325258, // 'XR24' - XRGB8888
      0x30323449, // 'I420' - I420
    ];
    for (const format of formats) {
      w.format(s, new_id, format);
    }

    // Send the modifier event for each format with some example modifiers
    const example_modifiers = [
      { hi: 0x0, lo: 0x0 }, // Linear modifier
      { hi: 0x0, lo: 0x1 }, // Example modifier
    ];
    for (const format of formats) {
      for (const mod of example_modifiers) {
        w.modifier(s, version_number, new_id, format, mod.hi, mod.lo);
      }
    }
  };
  static make(): w {
    return new w(new zwp_linux_dmabuf_v1());
  }
}
