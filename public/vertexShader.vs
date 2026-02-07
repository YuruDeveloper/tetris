#version 460

layout (location = 0) in vec2 shape;

void main() {
    gl_Position = vec4(shape,0,1);
}