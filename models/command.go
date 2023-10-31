package models

// Type Attribute values are : IN, OUT, ROUNDROBIN, DEFAULT
type Command struct {
	Type    string
	Patient *Patient
}
