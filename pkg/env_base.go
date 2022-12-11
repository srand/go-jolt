package jolt

var EnvBase = &Environment{
	Attributes: map[string]Attribute{
		"Outdir": NewStringAttribute("build.dir"),
	},
}
