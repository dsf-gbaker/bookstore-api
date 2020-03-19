package interfaces

import (
	"restapi/dtos"
)

// IValidator derp
type IValidator interface {
	IsValid(w *Book)
	IsValid(w *Author)
}
