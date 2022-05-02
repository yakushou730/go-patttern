package structural

import (
	"fmt"
	"testing"
)

func TestCompositeSwimmerA(t *testing.T) {
	localSwim := Swim
	swimmer := CompositeSwimmerA{
		MySwim: localSwim,
	}
	swimmer.MyAthlete.Train()
	swimmer.MySwim()
}

func TestShark(t *testing.T) {
	fish := Shark{
		Swim: Swim,
	}
	fish.Eat()
	fish.Swim()
}

func TestCompositeSwimmerB(t *testing.T) {
	swimmer := CompositeSwimmerB{
		Trainer: &Athlete{},
		Swimmer: &SwimmerImpl{},
	}
	swimmer.Train()
	swimmer.Swim()
}

func TestTree(t *testing.T) {
	root := Tree{
		LeafValue: 0,
		Right: &Tree{
			LeafValue: 5,
			Right: &Tree{
				LeafValue: 6,
				Right:     nil,
				Left:      nil,
			},
			Left: nil,
		},
		Left: &Tree{
			LeafValue: 4,
			Right:     nil,
			Left:      nil,
		},
	}
	fmt.Println(root.Right.Right.LeafValue)
}
