package runtime

// Runtime a the main factory to deal with all
type Runtime struct {
	models   map[string]*Model
	services map[string]*Service
}

// Model presents struct of model like message in proto
type Model struct {
	Name string
}

// Service presents struct of service like service in proto
type Service struct {
	Name string
}

// Load data from a spec
func (r *Runtime) Load(spec *Spec) error {

	return nil
}

// New init a runtime
func New() *Runtime {
	r := &Runtime{
		models:   map[string]*Model{},
		services: map[string]*Service{},
	}

	return r
}
