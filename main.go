package main

import(
  "detector"
  "log"
)

func main() {
  log.Printf("Hello World")

  rect := detector.Centered_rect__SWIG_0(1,1,5,5)

  log.Println(detector.Area(rect))
}
