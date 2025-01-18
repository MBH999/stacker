package types

type Config struct {
	Regions      Regions      `hcl:"regions,block"`
	Environments Environments `hcl:"environments,block"`
	Stacks       Stacks       `hcl:"stacks,block"`
}

type Stacks struct {
	Stack []Stack `hcl:"stack,block"`
}

type Stack struct {
	Name                string   `hcl:"name,label"`
	Tags                []string `hcl:"tags"`
	ExcludeRegions      []string `hcl:"exclude_regions,optional"`
	ExcludeEnvironments []string `hcl:"exclude_environments,optional"`
	ChildStack          []Stack  `hcl:"stack,block"`
}

type Regions struct {
	Config             []string `hcl:"config"`
	CreateRegionStacks bool     `hcl:"create_region_stacks"`
}

type Environments struct {
	Config                  []string `hcl:"config"`
	CreateEnvironmentStacks bool     `hcl:"create_environment_stacks"`
}
