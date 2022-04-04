package client

import (
	"testing"
)

var registerAmisStage RegisterAmisStage

func init() {
	registerAmisStage = *NewRegisterAmisStage()
}

func TestNewRegisterAmisStage(t *testing.T) {
	if registerAmisStage.Type != RegisterAmisStageType {
		t.Fatalf("Bake stage type should be %s, not \"%s\"", RegisterAmisStageType, registerAmisStage.Type)
	}
}

func TestRegisterAmisStageGetType(t *testing.T) {
	if registerAmisStage.GetType() != RegisterAmisStageType {
		t.Fatalf("Bake stage GetType() should be %s, not \"%s\"", RegisterAmisStageType, registerAmisStage.GetType())
	}
	if registerAmisStage.Type != RegisterAmisStageType {
		t.Fatalf("Bake stage Type should be %s, not \"%s\"", RegisterAmisStageType, registerAmisStage.Type)
	}
}
