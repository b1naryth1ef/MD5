package md5

import (
	"log"
	"strconv"
	"strings"
)

type MD5Model struct {
}

func parseLots(s string) []float64 {
	l := strings.Split(s, " ")
	res := make([]float64, 0)
	for _, val := range l {
		if (val != " ") && (val != ")") && (val != "(") && (val != "") {
			v, _ := strconv.ParseFloat(val, 64)
			res = append(res, v)
		}
	}
	return res
}

func (m *MD5Model) Load(file string) {
}

func main() {
	m := LoadAnimation("idle2.md5anim")
	log.Printf("%d, %d, %f", m.Version, m.NumFrames, m.AnimTime)
	log.Printf("Len Hier: %d", len(m.Hierarchys))
	log.Printf("Len Bounds: %d", len(m.Bounds))
	log.Printf("Len BF: %d", len(m.BaseFrames))
	log.Printf("Len Frames: %d", len(m.Frames))
}
