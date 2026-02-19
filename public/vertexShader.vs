#version 460

struct transform {
    vec2 size;
    vec2 location;
};

layout (location = 0) in vec2 vertex;
layout (location = 1) in vec2 uvMap;
//layout (location = 2) in int id;

layout(std140,binding = 0) uniform WorldBlock {
    vec2 viewsize;
};
layout(std430,binding = 1) buffer TransformBlock {
    transform list[];
};
out vec2 uv;

void main() {
    vec2 pixelLocation = vertex * list[gl_BaseInstance].size + list[gl_BaseInstance].location + (viewsize / 2);
    vec2 finalLocation = (pixelLocation / viewsize) * 2 - 1.0;
    gl_Position = vec4(finalLocation,0,1);
    uv = uvMap;
}