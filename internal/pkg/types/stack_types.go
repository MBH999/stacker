package types

type DecodedStacks struct {
	DecodedStack []DecodedStack
}

type DecodedStack struct {
	Path        string
	Tags        []string
	Description string
}
