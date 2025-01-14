package mixinprovider

import (
	"fmt"
	"path/filepath"

	"github.com/deislabs/porter/pkg/mixin"
	"github.com/pkg/errors"
)

func (fs *FileSystem) Delete(opts mixin.DeleteOptions) (*mixin.Metadata, error) {
	if opts.Name != "" {
		return fs.deleteByName(opts.Name)
	}

	return nil, errors.New("No mixin name was provided for deletion")
}

func (fs *FileSystem) deleteByName(name string) (*mixin.Metadata, error) {
	mixinsDir, err := fs.GetMixinsDir()
	if err != nil {
		return nil, err
	}
	mixinDir := filepath.Join(mixinsDir, name)
	exists, _ := fs.FileSystem.Exists(mixinDir)
	if exists == true {
		err = fs.FileSystem.RemoveAll(mixinDir)
		if err != nil {
			return nil, errors.Wrapf(err, "could not remove mixin directory %q", mixinDir)
		}

		m := mixin.Metadata{
			Name: name,
			Dir:  mixinsDir,
		}
		return &m, nil
	}

	if fs.Debug {
		fmt.Fprintln(fs.Out, "Unable to find requested mixin.")
	}

	return nil, nil
}
