package reslist

import (
	"github.com/scrollodex/dex/dexmodels"
)

// GITConfig stores configuration settings for the provider.
type GITConfig struct {
	URL string
}

// GITHandle is the handle used to refer to GIT.
type GITHandle struct {
	config   GITConfig
	fshandle FSHandle
}

// NewGIT creates a new GIT object.
func NewGIT(c GITConfig) (Databaser, error) {
	db := &GITHandle{
		config: c,
	}

	// TODO(tlim): If directory exists, git pull. else git clone

	// TODO(tlim): NewFS and store in db.fshandle

	return db, nil
}

// CategoryStore stores a category in stable storage.
func (rh GITHandle) CategoryStore(data dexmodels.Category) error {

	// TODO: git pull
	// TODO: rh.fsthandle.CategoryStore(data)
	// TODO: git commit and push

	return nil
}

// LocationStore stores a location in stable storage.
func (rh GITHandle) LocationStore(data dexmodels.Location) error {

	// TODO: git pull
	// TODO: rh.fsthandle.LocationStore(data)
	// TODO: git commit and push

	return nil
}

// EntryStore stores an entry in stable storage.
func (rh GITHandle) EntryStore(data dexmodels.Entry) error {

	// TODO: git pull
	// TODO: rh.fsthandle.LocationStore(data)
	// TODO: git commit and push

	return nil
}

// CategoryList returns a list of all categories.
func (rh GITHandle) CategoryList() ([]dexmodels.Category, error) {
	// TODO: git pull
	// TODO: rh.fsthandle.CategoryList(data)
	return nil, nil
}

// LocationList returns a list of all locations.
func (rh GITHandle) LocationList() ([]dexmodels.Location, error) {
	// TODO: git pull
	// TODO: rh.fsthandle.LocationList(data)
	return nil, nil
}

// EntryList returns a list of all entries.
func (rh GITHandle) EntryList() ([]dexmodels.Entry, error) {
	// TODO: git pull
	// TODO: rh.fsthandle.EntryList(data)
	return nil, nil
}

// CategoryGet gets a single item
func (rh GITHandle) CategoryGet(id int) (*dexmodels.Category, error) {
	// TODO: git pull
	// TODO: rh.fsthandle.CategoryGet(data)
	return nil, nil
}

// LocationGet gets a single item
func (rh GITHandle) LocationGet(id int) (*dexmodels.Location, error) {
	// TODO: git pull
	// TODO: rh.fsthandle.LocationGet(data)
	return nil, nil
}

// EntryGet gets a single item
func (rh GITHandle) EntryGet(id int) (*dexmodels.Entry, error) {
	// TODO: git pull
	// TODO: rh.fsthandle.EntryGet(data)
	return nil, nil
}
