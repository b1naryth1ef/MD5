# MD5

This is a library for loading (and possibly eventually rendering/etc) MD5 animations (.md5mesh .md5anim). It currently only supports loading (not saving) of MD5 files.

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