#version 460

layout (location = 0) in vec2 vertex;
layout (location = 1) in vec2 uvMap;

layout(std140,binding = 0) uniform 
TransformBlock {
    vec4 transform;
};

out vec2 uv;

void main() {
    vec2 size = transform.xy;
    vec2 location = transform.zw;
    vec2 finalLocation = vertex * size + location;
    gl_Position = vec4(finalLocation,0,1);
    uv = uvMap;
}