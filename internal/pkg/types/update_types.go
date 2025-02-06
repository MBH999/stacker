package types

type UpdateStackConfig struct {
	Stack UpdateStack `hcl:"stack,block"`
}

type UpdateStack struct {
	Name        string   `hcl:"name,attr"`          // required
	Description string   `hcl:"description,attr"`   // required
	Tags        []string `hcl:"tags,optional"`      // optional list of strings
	ID          string   `hcl:"id,optional"`        // optional
	Before      []string `hcl:"before,optional"`    // optional list
	After       []string `hcl:"after,optional"`     // optional list
	Wants       []string `hcl:"wants,optional"`     // optional list
	WantedBy    []string `hcl:"wanted_by,optional"` // optional list; note the snake_case here
	Watch       []string `hcl:"watch,optional"`     // optional list
}

