package model

// Model is a regular domain model, and has custom it's own method set and
// struct tags.
type Model struct {
	Foo string `validate:"len=10"`
	Bar int    `validate:"gt=0"`
}
