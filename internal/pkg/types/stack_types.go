package types

type Config struct {
	Stacks Stacks `hcl:"stacks,block"`
}

type Stacks struct {
	Stack []Stack `hcl:"stack,block"`
}

type Stack struct {
	Name                string   `hcl:"name,label"`
	Tags                []string `hcl:"tags"`
	ExcludeRegions      []string `hcl:"exclude_regions,optional"`
	ExcludeEnvironments []string `hcl:"exclude_environments,optional"`
	Stack               []Stack  `hcl:"stack,block"`
}
