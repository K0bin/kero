package entity

import (
	"github.com/galaco/kero/framework/entity"
)

// PropDynamicOverride
type PropDynamicOverride struct {
	entity.Entity
	PropRenderableBase
}

// Classname
func (entity PropDynamicOverride) Classname() string {
	return "prop_dynamic_override"
}
