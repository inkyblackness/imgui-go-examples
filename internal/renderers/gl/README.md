## Generating OpenGL bindings

Instead of depending on the pre-generated binding of [go-gl/gl](https://github.com/go-gl/gl), which is not always up to date, the examples application makes use of the code- & binding-generator [go-gl/glow](https://github.com/go-gl/glow).

To update the bindings with the newest generator, do the following:

1. `go get` or install `github.com/go-gl/glow` locally
1. run binding generator for the necessary versions; For example, within the root directory of `go-gl/glow`:
   ```
   go run . generate -out=...path/to/imgui-go-examples/internal/renderers/gl/v3.2-core/gl -api=gl -version=3.2 -profile=core -xml=./xml/ -tmpl=./tmpl/
   go run . generate -out=...path/to/imgui-go-examples/internal/renderers/gl/v2.1/gl -api=gl -version=2.1 -xml=./xml/ -tmpl=./tmpl/
   ```
