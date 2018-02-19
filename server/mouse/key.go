package mouse

import (
	"fmt"
	"math/rand"
)

type KeyContainer []string

var animals = []string{"alligator", "ant", "bear", "bee", "bird", "camel", "cat", "cheetah", "chicken", "chimpanzee", "cow", "crocodile", "deer", "dog", "dolphin", "duck", "eagle", "elephant", "fish", "fly", "fox", "frog", "giraffe", "goat", "goldfish", "hamster", "hippopotamus", "horse", "kangaroo", "kitten", "lion", "lobster", "monkey", "octopus", "owl", "panda", "pig", "puppy", "rabbit", "rat", "scorpion", "seal", "shark", "sheep", "snail", "snake", "spider", "squirrel", "tiger", "turtle", "wolf", "zebra"}

func pickOne() string {
	return animals[rand.Intn(len(animals))]
}

func (kc *KeyContainer) Generate() string {
	var candidate string
	for {
		candidate = fmt.Sprintf(`%s-%s-%s-%s`, pickOne(), pickOne(), pickOne(), pickOne())
		if kc.Exist(candidate) == -1 {
			break
		}
	}
	*kc = append(*kc, candidate)
	return candidate
}

func (kc *KeyContainer) Exist(key string) int {
	for index, key2 := range *kc {
		if key2 == key {
			return index
		}
	}
	return -1
}

func (kc *KeyContainer) Remove(index int) {
	*kc = append((*kc)[:index], (*kc)[index+1:]...)
}
