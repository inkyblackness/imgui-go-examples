package main

import (
	"fmt"
	"os"

	"github.com/inkyblackness/imgui-go"

	"github.com/inkyblackness/imgui-go-examples/pkg/demo"
	"github.com/inkyblackness/imgui-go-examples/pkg/platforms"
	"github.com/inkyblackness/imgui-go-examples/pkg/renderers"
)

func main() {
	context := imgui.CreateContext(nil)
	defer context.Destroy()
	io := imgui.CurrentIO()

	platform, err := platforms.NewSDL(io, platforms.SDLClientAPIOpenGL2)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}
	defer platform.Dispose()

	renderer, err := renderers.NewOpenGL2(io)
	if err != nil {
		os.Exit(-1)
	}

	demo.Run(platform, renderer)
}
