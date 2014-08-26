package main

import (
  "net/http"
  "image"
  "strings"
  "image/color"
  "image/png"
  "os"
  "fmt"
  "github.com/lucasb-eyer/go-colorful"
)

type GeneratedImage struct {
  C color.Color
}
func (GeneratedImage) Bounds() (image.Rectangle) {
  return image.Rectangle{image.Pt(0,0), image.Pt(1,1)}
}
func (GeneratedImage) ColorModel() (color.Model) {
  return color.RGBAModel
}

func (self GeneratedImage) At(_, _ int) (color.Color) {
  return self.C
}


func handler(w http.ResponseWriter, r *http.Request) {
  c, _ := colorful.Hex(fmt.Sprintf("#%s", strings.Trim(r.URL.Path, "/")))
  img := GeneratedImage{C: c}
  w.Header().Set("Cache-control", "public, max-age=259200")
  png.Encode(w, img)
}

func main() {
  http.HandleFunc("/", handler)
  err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)

  if err != nil {
    panic(err)
  }
}
