package shaders

const VertexShaderSource = `
	#version 410 core

	layout(location = 0) in vec3 inPosition;

	out vec3 FragPos;

	void main()
	{
	    gl_Position = vec4(inPosition, 1.0);
	    FragPos = inPosition;
	}
`

const FragmentShaderSource = `
	#version 410 core

	out vec4 FragColor;
	in vec3 FragPos;

	void main()
	{
	    // Calculate the color based on the fragment's position
	    vec3 color = vec3(abs(FragPos.x), abs(FragPos.y), abs(FragPos.z));
    
	    // Output the color
	    FragColor = vec4(color, 1.0);
	}
`
