package generator

import (
	"time"

	"github.com/sony/sonyflake"
)

// All base62 characters
var (
	alphabet = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	length   = uint64(len(alphabet))
)

type Generator struct {
	sf *sonyflake.Sonyflake
}

func NewGenerator() *Generator {
	return &Generator{
		sf: sonyflake.NewSonyflake(sonyflake.Settings{StartTime: time.Date(2022, 12, 1, 0, 0, 0, 0, time.UTC)}),
	}
}

//Generate make new short UUID
func (g *Generator) Generate() (string, error) {
	id, err := g.sf.NextID()
	if err != nil {
		return "", err
	}
	s := make([]byte, 0, 9)
	for id > 0 {
		s = append(s, alphabet[id%length])
		id = id / length

	}
	return string(s), nil

}
