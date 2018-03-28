package main

import(
  "detector"
  "log"
)

func main() {
  log.Printf("Example Create Rectangle")
  rect := detector.Centered_rect__SWIG_0(1,1,5,10)

  log.Printf("Example Caclulate Rectangle Area")
  log.Println(detector.Area(rect))
}
