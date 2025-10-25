import { auto_release } from "../auto_release.ts";
import {
  zwp_linux_buffer_params_v1_delegate as d,
  zwp_linux_buffer_params_v1 as w,
  wl_shm_pool as pool_w,
  wl_buffer as wl_buffer_w,
} from "../protocols/wayland.xml.ts";
import { Wayland_Client } from "../Wayland_Client.ts";
import { File_Descriptor, Object_ID } from "../wayland_types.ts";
import { wl_shm } from "./wl_shm.ts";
import { wl_shm_pool } from "./wl_shm_pool.ts";

interface Plane {
  fd: File_Descriptor;
  plane_idx: number;
  offset: number;
  stride: number;
  modifier_hi: number;
  modifier_lo: number;
}

let fake_pool_ids = 0xf000_000 as Object_ID<pool_w>;
let buffer_ids = 0xf100_000 as Object_ID<wl_buffer_w>;

export class zwp_linux_buffer_params_v1 implements d {
  zwp_linux_buffer_params_v1_destroy: d["zwp_linux_buffer_params_v1_destroy"] =
    auto_release;

  planes: Plane[] = [];

  zwp_linux_buffer_params_v1_add: d["zwp_linux_buffer_params_v1_add"] = (
    s,
    object_id,
    fd,
    plane_idx,
    offset,
    stride,
    modifier_hi,
    modifier_lo
  ) => {
    this.planes.push({
      fd,
      plane_idx,
      offset,
      stride,
      modifier_hi,
      modifier_lo,
    });
  };
  zwp_linux_buffer_params_v1_create: d["zwp_linux_buffer_params_v1_create"] = (
    s,
    object_id,
    width,
    height,
    format,
    flags
  ) => {
    const buffer_id = buffer_ids;
    buffer_ids = (buffer_ids + 1) as Object_ID<wl_buffer_w>;

    this.zwp_linux_buffer_params_v1_create_immed(
      s,
      object_id,
      buffer_id,
      width,
      height,
      format,
      flags
    );

    w.created(s, object_id, buffer_id);
    
  };
  zwp_linux_buffer_params_v1_create_immed: d["zwp_linux_buffer_params_v1_create_immed"] =
    (s, object_id, buffer_id, width, height, format, flags) => {
      /**
       * @TODO not sure how to handle buffers with planes
       * that are not all in one fd
       */
      console.log(`Creating dmabuf buffer ${buffer_id} ${width}x${height} format=${format} flags=${flags} with ${this.planes.length} planes`);
      if (this.planes.length === 0) {
        console.log("No planes specified for dmabuf buffer");
        w.failed(s, object_id);
        return;
      }
      /**
       * @TODO handl multiple fds properly
       */
      const fd = this.planes[0].fd;

      for (const plane of this.planes) {
        if (plane.fd === fd) {
          continue;
        }
        console.log(
          "Multiple fds specified for dmabuf buffer, not supported yet"
        );
        w.failed(s, object_id);
        return;
      }
      const pool_id = fake_pool_ids;
      fake_pool_ids = (fake_pool_ids + 1) as Object_ID<pool_w>;

      const buffer_pool = wl_shm_pool.make(
        s,
        pool_id,
        fd,
        this.find_buffer_size(height)
      );
      
      s.add_object(pool_id, buffer_pool);
      /**
       * @TODO what do do with multiple planes?
       */
      buffer_pool.delegate.wl_shm_pool_create_buffer(
        s,
        pool_id,
        buffer_id,
        0,
        width,
        height,
        this.planes[0].stride,
        format
      )
      /**
       * Because the client doesn't know about the pool object
       * we need to destroy it here. It will be destroyed after
       * all buffers created from it are destroyed.
       */
      buffer_pool.delegate.wl_shm_pool_destroy(s, pool_id);
    };

  find_buffer_size = (height: number) =>
    this.planes.reduce((max, plane) => {
      const end = plane.offset + height * plane.stride;
      return end > max ? end : max;
    }, 0);

  zwp_linux_buffer_params_v1_on_bind: d["zwp_linux_buffer_params_v1_on_bind"] =
    (s, name, interface_, new_id, version_number) => {};
  constructor() {}
  static make(): w {
    return new w(new zwp_linux_buffer_params_v1());
  }
}
