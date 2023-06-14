package pb

import (
	"testing"

	"github.com/cloudwego/fastpb"
	"google.golang.org/protobuf/proto"
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

func Marshal(msg fastpb.Writer) []byte {
	buf := make([]byte, msg.Size())

	msg.FastWrite(buf)
	return buf
}

func Unmarshal(buf []byte, msg fastpb.Reader) error {
	_, err := fastpb.ReadMessage(buf, int8(fastpb.SkipTypeCheck), msg)
	return err
}

func BenchmarkSmallerMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Marshal(jerry)
	}
}

func BenchmarkSmallerUnmarshal(b *testing.B) {
	data := Marshal(jerry)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		friend := new(Friend)
		err := Unmarshal(data, friend)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkLargerMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Marshal(tom)
	}
}

func BenchmarkLargerUnmarshal(b *testing.B) {
	data := Marshal(tom)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		user := new(User)
		err := Unmarshal(data, user)
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
