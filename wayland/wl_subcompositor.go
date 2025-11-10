package wayland

import (
	"github.com/mmulet/term.everything/wayland/protocols"
)

type WlSubcompositor struct{}

func (sc *WlSubcompositor) WlSubcompositor_destroy(
	_ protocols.ClientState,
	_ protocols.ObjectID[protocols.WlSubcompositor],
) bool {
	return true
}

func (sc *WlSubcompositor) WlSubcompositor_get_subsurface(
	s protocols.ClientState,
	object_id protocols.ObjectID[protocols.WlSubcompositor],
	id protocols.ObjectID[protocols.WlSubsurface],
	surface_id protocols.ObjectID[protocols.WlSurface],
	parent_surface_id protocols.ObjectID[protocols.WlSurface],
) {

	surface := GetWlSurfaceObject(s, surface_id)
	if surface == nil {
		SendError(
			s,
			object_id,
			protocols.WlSubcompositorError_enum_bad_surface,
			"surface not found",
		)
		return
	}

	if surface.Role == nil {
		surface.Role = &SurfaceRoleSubSurface{Data: nil}
	}

	roleSub, ok := surface.Role.(*SurfaceRoleSubSurface)
	if !ok {
		SendError(
			s,
			object_id,
			protocols.WlSubcompositorError_enum_bad_surface,
			"surface has different role instead of sub_surface",
		)
		return
	}

	if roleSub.HasData() {
		SendError(
			s,
			object_id,
			protocols.WlSubcompositorError_enum_bad_surface,
			"surface already is a subsurface",
		)
		return
	}

	if surface_id == parent_surface_id {
		SendError(
			s,
			object_id,
			protocols.WlSubcompositorError_enum_bad_parent,
			"parent == surface",
		)
		return
	}

	if s.FindDescendantSurface(surface_id, parent_surface_id) {
		SendError(
			s,
			object_id,
			protocols.WlSubcompositorError_enum_bad_parent,
			"parent is a descendant of surface",
		)
		return
	}

	parent_surface := GetWlSurfaceObject(s, parent_surface_id)
	if parent_surface == nil {
		SendError(
			s,
			object_id,
			protocols.WlSubcompositorError_enum_bad_parent,
			"parent not found",
		)
		return
	}

	roleSub.Data = &id
	if parent_surface.PendingUpdate.AddSubSurface == nil {
		parent_surface.PendingUpdate.AddSubSurface = []protocols.ObjectID[protocols.WlSurface]{}
	}
	parent_surface.PendingUpdate.AddSubSurface = append(
		parent_surface.PendingUpdate.AddSubSurface,
		surface_id,
	)

	RegisterRoleToSurface(s, id, surface_id)
	AddObject(s, id, MakeWlSubsurface(parent_surface_id))
}

func (sc *WlSubcompositor) OnBind(
	_ protocols.ClientState,
	_ protocols.AnyObjectID,
	_ string,
	_ protocols.AnyObjectID,
	_ uint32,
) {
}

func MakeWlSubcompositor() *protocols.WlSubcompositor {
	return &protocols.WlSubcompositor{
		Delegate: &WlSubcompositor{},
	}
}
