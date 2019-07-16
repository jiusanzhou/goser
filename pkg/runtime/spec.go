package runtime

// Spec present all things in yaml
type Spec struct {
	Meta Meta `yaml:"_" json:"_"`
}

// Meta presents settings option
type Meta struct {
	// Readable present that the model will marshal to
	// response
	Readable bool
	// Writable present that the model will unmarshal from
	// request
	Writable bool
}
