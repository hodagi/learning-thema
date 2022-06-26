package example

import (
	"embed"

	"cuelang.org/go/cue"
	"github.com/grafana/thema"
	"github.com/grafana/thema/load"
)

//go:embed ship.cue cue.mod/module.cue
var modFS embed.FS

// file=>raw data=>cue file
func loadLineage(lib thema.Library) (cue.Value, error) {
	inst, err := load.InstancesWithThema(modFS, ".")
	if err != nil {
		return cue.Value{}, err
	}

	val := lib.Context().BuildInstance(inst)
	return val.LookupPath(cue.MakePath(cue.Str("lin"))), nil
}

// ShipLineage constructs a Go handle representing the Ship lineage.
func ShipLineage(lib thema.Library, opts ...thema.BindOption) (thema.Lineage, error) {
	linval, err := loadLineage(lib)
	if err != nil {
		return nil, err
	}
	return thema.BindLineage(linval, lib, opts...)
}

var _ thema.LineageFactory = ShipLineage // Ensure our factory fulfills the type
