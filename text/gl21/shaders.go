package gl21

var (
	glyphVertexShader = []byte(`uniform mat3 transform;

attribute vec4 pos;

varying vec2 bc;

void main() {
	bc = pos.zw;
	gl_Position = vec4(transform * vec3(pos.xy, 1.0), 0.0).xywz;
}
`)

	glyphFragmentShader = []byte(`uniform vec4 colour;

varying vec2 bc;

void main() {
	if (bc.x * bc.x - bc.y > 0.0) {
		discard;
	}
	gl_FragColor = colour * (gl_FrontFacing ? 16.0 / 255.0 : 1.0 / 255.0);
}
`)
)
