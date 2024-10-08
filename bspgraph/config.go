package bspgraph

import (
	"github.com/hashicorp/go-multierror"
	"github.com/rezaAmiri123/edatV2/bspgraph/message"
	"golang.org/x/xerrors"
)

// GraphConfig encapsulates the configuration options for creating graphs.
type GraphConfig struct {
	// QueueFactory is used by the graph to create message queue instances
	// for each vertex that is added to the graph. If not specified, the
	// default in-memory queue will be used instead.
	QueueFactory message.QueueFactory

	// ComputeFn is the compute function that will be invoked for each graph
	// vertex when executing a superstep. A valid ComputeFnnc instance is
	// required for the config to be valid.
	ComputeFn ComputeFnnc

	// ComputeWorkers specifies the number of workers to use for invoking
	// the registered ComputeFnnc when executing each superstep. If not
	// specified, a single worker will be used.
	ComputeWorkers int
}

// validate checks whether a graph configuration is valid and sets the default
// values where required.
func (g *GraphConfig) validate() error {
	var err error
	if g.QueueFactory == nil {
		g.QueueFactory = message.NewInMemoryQueue
	}
	if g.ComputeWorkers <= 0 {
		g.ComputeWorkers = 1
	}

	if g.ComputeFn == nil {
		err = multierror.Append(err, xerrors.New("compute function not specified"))
	}

	return err
}
