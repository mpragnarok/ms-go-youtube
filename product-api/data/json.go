package data

import (
	"encoding/json"
	"io"
)

// FromJSON deserializes the object from JSON string
// in a io.Reader to giver interface
func FromJSON(i interface{}, r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(i)
}

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocatio and the overheads of the service
//
// https://golang.org/pkg/encoding/json/#NewEncoder
func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}
