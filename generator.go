package main

import (
	crand "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
)

type Generator interface {
	Generate() (string, error)
}

type Uuidv1Generator struct {
}

func (g *Uuidv1Generator) Generate() (string, error) {
	u, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

type Uuidv4Generator struct {
}

func (g *Uuidv4Generator) Generate() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

type UlidGenerator struct {
	t time.Time
}

func (g *UlidGenerator) Generate() (string, error) {
	s, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	entropy := ulid.Monotonic(rand.New(rand.NewSource(s.Int64())), 0)
	return ulid.MustNew(ulid.Timestamp(g.t), entropy).String(), nil
}

func GetGenerator(format string) (Generator, error) {
	switch format {
	case "uuidv1":
		return &Uuidv1Generator{}, nil
	case "uuidv4":
		return &Uuidv4Generator{}, nil
	case "ulid":
		return &UlidGenerator{t: time.Now()}, nil
	default:
		return nil, fmt.Errorf("no support format: %s", format)
	}
}
