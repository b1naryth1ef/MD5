# MD5

This is a library for loading (and possibly eventually rendering/etc) MD5 animations (.md5mesh .md5anim). It currently only supports loading (not saving) of MD5 files. It was built to match the MD5Mesh/Anim specifications (you can view those over at http://tfc.duke.free.fr/coding/md5-specs-en.html). The library exposes too main functions, `LoadAnimation(s string)` and `LoadMesh(s string)` which both take a file name/path and return an MD5Animation or MD5Mesh accordingly. 


```go
package main

import "github.com/b1naryth1ef/MD5"
import "log"

func main() {
    m := md5.LoadAnimation("idle2.md5anim")
    log.Printf("%d, %d, %f", m.Version, m.NumFrames, m.AnimTime)
    log.Printf("Len Hier: %d", len(m.Hierarchys))
    log.Printf("Len Bounds: %d", len(m.Bounds))
    log.Printf("Len BF: %d", len(m.BaseFrames))
    log.Printf("Len Frames: %d", len(m.Frames))
}
```