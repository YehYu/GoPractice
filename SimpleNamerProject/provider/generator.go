package provider

import (
	"math/rand"
	"time"
)

type Generator struct {
	rand  *rand.Rand
	Names NamesResource
}

//Name : Get Top 100 Chinese Name
func (generator *Generator) Name() string {
	return generator.lastname() + generator.firstname()
}

func (generator *Generator) lastname() string {
	length := len(generator.Names.LastNames)
	return generator.Names.LastNames[generator.rand.Intn(length)]
}

func (generator *Generator) firstname() string {
	length := len(generator.Names.FirstNames)
	return generator.Names.FirstNames[generator.rand.Intn(length)]
}

func Create() Generator {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return Generator{
		rand: r,
	}
}
