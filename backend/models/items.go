package models

//the tags are used for JSON serialization and deserialization.
// For example, the ID field will be represented as "id" in the JSON object, Name as "name", and Type as "type".
type Food struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
}