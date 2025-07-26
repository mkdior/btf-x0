package database

// We assume that our primary data storing methodology uses hashes to point at
// serialized data. To allow for alternative access, we've implemented index
// maps for specific objects which we wrap over implementations of this
// interface. See the /internal/index folder for the wrappers.

// @TODO(Hamza) ~> Check if we're to change the naming here of the byte-array
// in case it doesn't properly denote our intention of using hashes.

type Database interface {
	Get([32]byte) (DataFormat, error)
	Set(DataFormat) error
	Delete([32]byte) error
}
