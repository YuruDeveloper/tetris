#version 460

in vec2 uv;

out vec4 fragmentColor;

uniform sampler2D textureMap;
uniform vec4 colorTint;
void main() {
    fragmentColor = texture(textureMap,uv) * colorTint;
}