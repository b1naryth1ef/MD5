package md5

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type Joint struct {
	Name   string
	Parent int
	PosX   float64
	PosY   float64
	PosZ   float64
	OriX   float64
	OriY   float64
	OriZ   float64
}

type Vert struct {
	Index       int
	S           float64
	T           float64
	StartWeight float64
	CountWeight float64
}

type Tri struct {
	Index     int
	VertIndex []int
}

type Weight struct {
	Index int
	Joint int
	Bias  float64
	PosX  float64
	PosY  float64
	PosZ  float64
}

type Mesh struct {
	NumVerts   int
	NumTris    int
	NumWeights int
	Shader     string
	Verts      map[int]Vert
	Tris       map[int]Tri
	Weights    map[int]Weight
}

func newMesh() Mesh {
	m := Mesh{}
	m.Verts = make(map[int]Vert)
	m.Tris = make(map[int]Tri)
	m.Weights = make(map[int]Weight)
	return m
}

type MD5Mesh struct {
	Version   int
	NumJoints int
	NumMeshes int

	Joints map[int]Joint
	Meshes map[int]Mesh
}

func (m Mesh) Parse(line string) Mesh {
	line = strings.TrimLeft(line, " ")
	if strings.HasPrefix(line, "shader") {
		val := strings.Split(line, " ")
		m.Shader = val[0]
	}
	if strings.HasPrefix(line, "numverts") {
		val := strings.Split(line, " ")
		m.NumVerts, _ = strconv.Atoi(val[1])
	}
	if strings.HasPrefix(line, "numtris") {
		val := strings.Split(line, " ")
		m.NumTris, _ = strconv.Atoi(val[1])
	}
	if strings.HasPrefix(line, "numweights") {
		val := strings.Split(line, " ")
		m.NumWeights, _ = strconv.Atoi(val[1])
	}
	if strings.HasPrefix(line, "vert") {
		vals := parseLots(strings.Join(strings.Split(line, "")[1:], " "))
		m.Verts[len(m.Verts)] = Vert{int(vals[0]), vals[1], vals[2], vals[3], vals[4]}
	}
	if strings.HasPrefix(line, "tri") {
		vals := parseLots(strings.Join(strings.Split(line, "")[1:], " "))
		v := []int{int(vals[1]), int(vals[2]), int(vals[3])}
		m.Tris[len(m.Tris)] = Tri{int(vals[0]), v}
	}
	if strings.HasPrefix(line, "weight") {
		vals := parseLots(strings.Join(strings.Split(line, "")[1:], " "))
		m.Weights[len(m.Weights)] = Weight{int(vals[0]), int(vals[1]), vals[2], vals[3], vals[4], vals[5]}
	}
	return m
}

func LoadMesh(file string) *MD5Mesh {
	var status string

	mesh := MD5Mesh{}

	f, _ := os.Open(file)
	r := bufio.NewReader(f)
	for {
		l, _, er := r.ReadLine()
		li := string(l)
		if er == io.EOF {
			break
		}
		split := strings.SplitN(li, "//", 2)
		line := strings.Replace(split[0], "\t", " ", -1) //Fuck you tabs

		// Begin Null
		if status == "" {
			if strings.HasPrefix(line, "MD5Version") {
				val := strings.Split(line, " ")
				mesh.Version, _ = strconv.Atoi(val[1])
			}

			if strings.HasPrefix(line, "commandline") {
				continue
			}

			if strings.HasPrefix(line, "numJoints") {
				val := strings.Split(line, " ")
				mesh.NumJoints, _ = strconv.Atoi(val[1])
			}

			if strings.HasPrefix(line, "numMeshes") {
				val := strings.Split(line, " ")
				mesh.NumMeshes, _ = strconv.Atoi(val[1])
			}
			if strings.HasPrefix(line, "joints") {
				status = "joints"
				mesh.Joints = make(map[int]Joint)
				continue
			}
			if strings.HasPrefix(line, "mesh") {
				status = "mesh"
				if mesh.Meshes == nil {
					mesh.Meshes = make(map[int]Mesh)
				}
				mesh.Meshes[len(mesh.Meshes)] = newMesh()

			}

		} else {
			if line == "}" {
				status = ""
			}
		}
		// End Null

		if status == "joints" {
			val := strings.Split(line, " ")
			res := parseLots(strings.Join(val[2:], " "))
			mesh.Joints[len(mesh.Joints)] = Joint{strings.Replace(val[1], "\"", "", -1), int(res[0]), res[1], res[2], res[3], res[4], res[5], res[6]}
		}
		if status == "mesh" {
			mesh.Meshes[len(mesh.Meshes)-1] = mesh.Meshes[len(mesh.Meshes)-1].Parse(line)
		}
	}
	return &mesh
}
