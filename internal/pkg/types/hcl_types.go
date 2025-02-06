package types

type Config struct {
	Regions      Regions      `hcl:"regions,block"`
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
	ChildStack          []Stack  `hcl:"stack,block"`
}

type Regions struct {
	Region []Region `hcl:"region,block"`
}

type Region struct {
	Name          string   `hcl:"name,label"`
	Tags          []string `hcl:"tags"`
	DeployAsStack bool     `hcl:"deploy_as_stack"`
}
