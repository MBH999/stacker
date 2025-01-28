package types

type DecodedStacks struct {
	DecodedStack []DecodedStack
}

type DecodedStack struct {
	Name                   string
	Path                          string
	ParentPath                   string
	Tags                          []string
	Description                   string
	Region                             string
	Environment                   string
	ExcludedEnvironments []string
	ExcludedRegions           []string
}
