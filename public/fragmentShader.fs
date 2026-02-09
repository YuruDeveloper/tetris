#version 460

in vec2 uv;

out vec4 fragmentColor;

uniform sampler2D textureMap;
void main() {
    fragmentColor = texture(textureMap,uv);
}