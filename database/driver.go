package database

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/jcelliott/lumber"
	"github.com/satyajitnayk/json-database/logger"
)

type Driver struct {
	mutex   sync.Mutex
	mutexes map[string]*sync.Mutex
	dir     string
	log     logger.Logger
}

// Options contains optional configurations for the database.
type Options struct {
	logger.Logger
}

func New(dir string, options *Options) (*Driver, error) {
	dir = filepath.Clean(dir)

	opts := Options{}

	if options != nil {
		opts = *options
	}

	if opts.Logger == nil {
		opts.Logger = lumber.NewConsoleLogger(lumber.INFO)
	}
	// Append "collections/" to the base directory
	dir = filepath.Join(dir, "collections")

	driver := Driver{
		dir:     dir,
		mutexes: make(map[string]*sync.Mutex),
		log:     opts.Logger,
	}

	if _, err := os.Stat(dir); err == nil {
		opts.Logger.Debug("Using '%s' (database already exists)\n", dir)
		return &driver, nil
	}

	// careate db
	opts.Logger.Debug("Creating the database at '%s' ...\n", dir)
	// 0755 - folder access level
	return &driver, os.MkdirAll(dir, 0755)
}
