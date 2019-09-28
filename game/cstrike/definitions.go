package cstrike

import (
	"github.com/galaco/kero/game/common/entity"
	loader "github.com/galaco/kero/valve/entity"
)

// CounterstrikeSource
type CounterstrikeSource struct{}

// RegisterEntityClasses loads all Game entity classes into the engine.
func (target *CounterstrikeSource) RegisterEntityClasses() {
	loader.RegisterClass(&entity.InfoPlayerStart{})
	loader.RegisterClass(&entity.PropDoorRotating{})
	loader.RegisterClass(&entity.PropDynamic{})
	loader.RegisterClass(&entity.PropDynamicOrnament{})
	loader.RegisterClass(&entity.PropDynamicOverride{})
	loader.RegisterClass(&entity.PropPhysics{})
	loader.RegisterClass(&entity.PropPhysicsMultiplayer{})
	loader.RegisterClass(&entity.PropPhysicsOverride{})
	loader.RegisterClass(&entity.PropRagdoll{})
}
