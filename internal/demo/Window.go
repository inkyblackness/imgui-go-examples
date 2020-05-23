package demo

import (
	"fmt"

	"github.com/inkyblackness/imgui-go/v2"
)

var window = struct {
	flags int
}{}

// Show demonstrates most ImGui features that were ported to Go.
// This function tries to recreate the original demo window as closely as possible.
//
// In theory, if both windows would provide the identical functionality, then the wrapper would be complete.
func Show(show *bool) {
	open := show

	imgui.SetNextWindowPosV(imgui.Vec2{X: 650, Y: 20}, imgui.ConditionFirstUseEver, imgui.Vec2{})
	imgui.SetNextWindowSizeV(imgui.Vec2{X: 550, Y: 680}, imgui.ConditionFirstUseEver)

	if !imgui.BeginV("ImGui-Go Demo", open, window.flags) {
		// Early out if the window is collapsed, as an optimization.
		imgui.End()
		return
	}
	imgui.Text(fmt.Sprintf("dear imgui says hello. (%s)", imgui.Version()))

	// TODO: missing imgui.FontSize() to push item width

	imgui.End()
}
