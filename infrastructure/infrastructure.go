package infrastructure

import (
	"github.com/hashicorp/otto/appfile"
)

// Factory is a factory function for creating infrastructures.
type Factory func() (Infrastructure, error)

// Infrastructure is an interface that must be implemented by each
// infrastructure type with a method of creating it.
type Infrastructure interface {
	Compile(*Context) (*CompileResult, error)
	Flavors() []string
}

// Context is the context for operations on infrastructures. Some of
// the fields in this struct are only available for certain operations.
type Context struct {
	// Dir is the directory that the compilation is allowed to write to
	// for persistant storage of data.
	Dir string

	// The infrastructure configuration itself from the Appfile. This includes
	// the flavor of the infrastructure we want to launch.
	Infra *appfile.Infrastructure
}

// CompileResult is the structure containing compilation result values.
type CompileResult struct {
}
