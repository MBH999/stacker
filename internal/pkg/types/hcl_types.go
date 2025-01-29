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
	Description         string   `hcl:"description,optional"`
	ExcludeRegions      []string `hcl:"exclude_regions,optional"`
	ExcludeEnvironments []string `hcl:"exclude_environments,optional"`
	ChildStack          []Stack  `hcl:"stack,block"`
}

type Regions struct {
	Region []Region `hcl:"region,block"`
	// Config             []string `hcl:"config"`
	// CreateRegionStacks bool     `hcl:"create_region_stacks"`
}

type Region struct {
	Name          string   `hcl:"name,label"`
	Tags          []string `hcl:"tags"`
	DeployAsStack bool     `hcl:"deploy_as_stack"`
}

type Environments struct {
	Environment []Environment `hcl:"environment,block"`
	// Config                  []string `hcl:"config"`
	// CreateEnvironmentStacks bool     `hcl:"create_environment_stacks"`
}

type Environment struct {
	Name          string   `hcl:"name,label"`
	Tags          []string `hcl:"tags"`
	DeployAsStack bool     `hcl:"deploy_as_stack"`
}
