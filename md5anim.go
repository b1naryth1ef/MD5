package md5

func (f Frame) Add(a []float64) {
	f.Vars = append(f.Vars, a)
}

func newFrame(id int) Frame {
	return Frame{id, make([][]float64, 0)}
}

type Hierarchy struct {
	Name       string
	Parent     int
	Flags      int
	StartIndex int
}

type Bound struct {
	MinX float64
	MinY float64
	MinZ float64
	MaxX float64
	MaxY float64
	MaxZ float64
}

type BaseFrame struct {
	PosX float64
	PosY float64
	PosZ float64
	OriX float64
	OriY float64
	OriZ float64
}

type Frame struct {
	Id   int
	Vars [][]float64
}

type MD5Animation struct {
	Version               int
	NumFrames             int
	NumJoints             int
	FrameRate             int
	NumAnimatedComponents int
	AnimDur               float64
	FrameDur              float64
	AnimTime              float64

	// SubObjecst
	Hierarchys map[int]Hierarchy
	Bounds     map[int]Bound
	BaseFrames map[int]BaseFrame
	Frames     map[int]Frame
}

func LoadAnimation(file string) *MD5Animation {
	var status string

	// Setup RE
	quotes_re := regexp.MustCompile("(.*?)")

	anim := MD5Animation{}
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

		// Begin Null Block
		if status == "" {
			// MD5 Version
			if strings.HasPrefix(line, "MD5Version") {
				val := strings.Split(line, " ")
				anim.Version, _ = strconv.Atoi(val[1])
			}

			// Ignore DOOM 3 shit
			if strings.HasPrefix(line, "commandline") {
				continue
			}

			// NumFrames
			if strings.HasPrefix(line, "numFrames") {
				val := strings.Split(line, " ")
				anim.NumFrames, _ = strconv.Atoi(val[1])
			}

			// FrameRate
			if strings.HasPrefix(line, "frameRate") {
				val := strings.Split(line, " ")
				anim.FrameRate, _ = strconv.Atoi(val[1])
			}

			// NumAnimatedComponents
			if strings.HasPrefix(line, "numAnimatedComponents") {
				val := strings.Split(line, " ")
				anim.NumAnimatedComponents, _ = strconv.Atoi(val[1])
			}

			// Blocks
			if strings.HasPrefix(line, "hierarchy") {
				status = "hierarchy"
				anim.Hierarchys = make(map[int]Hierarchy)
				continue
			}
			if strings.HasPrefix(line, "bounds") {
				status = "bounds"
				anim.Bounds = make(map[int]Bound)
				continue
			}
			if strings.HasPrefix(line, "baseframe") {
				status = "baseframe"
				anim.BaseFrames = make(map[int]BaseFrame)
				continue
			}
			if strings.HasPrefix(line, "frame") {
				status = "frame"
				frameid, _ := strconv.Atoi(strings.Split(line, " ")[1])
				if anim.Frames == nil {
					anim.Frames = make(map[int]Frame)
				}
				anim.Frames[frameid] = newFrame(frameid)
			}

		} else {
			if line == "}" {
				status = ""
			}
		}
		// End Null Block

		if status == "hierarchy" {
			sp := strings.Split(line, " ")
			h := Hierarchy{}
			h.Name = quotes_re.FindAllString(sp[0], -1)[0]
			h.Parent, _ = strconv.Atoi(sp[1])
			h.Flags, _ = strconv.Atoi(sp[2])
			h.StartIndex, _ = strconv.Atoi(sp[3])
			anim.Hierarchys[len(anim.Hierarchys)] = h
		}

		if status == "bounds" {
			vals := parseLots(line)
			b := Bound{vals[0], vals[1], vals[2], vals[3], vals[4], vals[5]}
			anim.Bounds[len(anim.Bounds)] = b
		}

		if status == "baseframe" {
			vals := parseLots(line)
			bf := BaseFrame{vals[0], vals[1], vals[2], vals[3], vals[4], vals[5]}
			anim.BaseFrames[len(anim.BaseFrames)] = bf
		}

		if status == "frame" {
			anim.Frames[len(anim.Frames)].Add(parseLots(line))
		}

	}
	return &anim
}
