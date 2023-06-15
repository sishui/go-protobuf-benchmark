package pb

import (
	"testing"

	"github.com/gogo/protobuf/proto"
)

var jerry = &Friend{
	Uid:     9528,
	Name:    "Jerry",
	Comment: []byte("I am Jerry"),
	Career:  Career_CAREER_DEATH_KNIGHT,
	Race:    Classes_CLASSES_NIGHT_ELF,
}

var spike = &Friend{
	Uid:     9529,
	Name:    "Spike",
	Comment: []byte("I am Spike"),
	Career:  Career_CAREER_WARRIOR,
	Race:    Classes_CLASSES_DRACTHYR,
}

var tom = &User{
	Uid:     9527,
	Name:    "Tom",
	Age:     20,
	Height:  170,
	Weight:  65,
	Comment: []byte("I am Tom"),
	Career:  Career_CAREER_HUNTER,
	Race:    Classes_CLASSES_ORC,
	Friends: []*Friend{
		jerry,
		spike,
	},
	Props: map[string]int32{
		"HEALTH":          200000,
		"STRENGTH":        100,
		"STAMINA":         300,
		"CRITICAL_STRIKE": 20,
		"VERSATILITY":     100,
	},
}

func BenchmarkSmallerMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(jerry)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSmallerUnmarshal(b *testing.B) {
	data, err := jerry.Marshal()
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		friend := new(Friend)
		err := friend.Unmarshal(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkLargerMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := tom.Marshal()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkLargerUnmarshal(b *testing.B) {
	data, err := tom.Marshal()
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		user := new(User)
		err := user.Unmarshal(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSmallerClone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = proto.Clone(jerry)
	}
}

func BenchmarkLargerClone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = proto.Clone(tom)
	}
}
