#version 460

layout (location = 0) in vec2 shape;
layout(std140,binding = 0) uniform 
TransformBlock {
    vec4 transform;
};

void main() {
    vec2 size = transform.xy;
    vec2 location = transform.zw;
    vec2 finalLocation = shape * size + location;
    gl_Position = vec4(finalLocation,0,1);
}