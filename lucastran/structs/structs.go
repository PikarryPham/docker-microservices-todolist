package structs

import "github.com/google/uuid"

type IP struct {
    Query string
}

type HTTPResponse struct {
	Status      int
	Application string
	IP	string
	UUID        uuid.UUID
	Data        string
}
