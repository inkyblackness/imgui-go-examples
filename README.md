# Examples of Dear ImGui for Go

This project contains a set of [Go](https://www.golang.org) examples for [imgui-go](https://github.com/inkyblackness/imgui-go), which is a wrapper for [**Dear ImGui**](https://github.com/ocornut/imgui).

## Layout
The project follows the basic concept of the examples of **Dear ImGui** by separating platform and renderer bindings from the example applications that wire them together in compatible constellations.

* `cmd` contains the main functions of the example applications.
* `pkg` contains the reusable library components
  * `platforms` contains code for mouse/keyboard/gamepad inputs, cursor shape, timing, windowing.
  * `renderers` contains code for creating the main font texture, rendering imgui draw data.
  * `demo` contains the common example code.

## License

The project is available under the terms of the **New BSD License** (see LICENSE file).
