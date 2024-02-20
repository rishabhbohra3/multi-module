package dto

type ActionDefination struct {
	ADFVersion string   `yaml:"adf_version"`
	Actions    []Action `yaml:"actions"`
}

type Action struct {
	Name        string    `yaml:"name"`
	Version     string    `yaml:"version"`
	Author      string    `yaml:"author"`
	Category    string    `yaml:"category"`
	Description string    `yaml:"description"`
	Namespace   string    `yaml:"namespace"`
	CommandName string    `yaml:"command_name"`
	Inputs      []Input   `yaml:"inputs"`
	Outputs     []Output  `yaml:"outputs"`
	Prompt      string    `yaml:"prompt"`
	Contexts    []Context `yaml:"contexts"`
}

type Input struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Required    bool   `yaml:"required"`
	Default     string `yaml:"default"`
}

type Output struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

type Context struct {
	CodeContext struct {
		Enabled                bool   `yaml:"enabled"`
		AdditionalInstructions string `yaml:"additional_instructions"`
	} `yaml:"codeContext"`
	RunContext struct {
		Enabled                bool   `yaml:"enabled"`
		AdditionalInstructions string `yaml:"additional_instructions"`
	} `yaml:"runContext"`
}
