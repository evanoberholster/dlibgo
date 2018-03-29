package main

import(
  "detector"
  "log"
)

func main() {
  log.Printf("Example Create Rectangle")
  rect := detector.Centered_rect__SWIG_0(30,30,300,300)

  log.Printf("Example Caclulate Rectangle Area")
  log.Println(detector.Area(rect))

  shapePred := detector.LoadShapePredictor("/go/src/detector/example/shape_predictor_68_face_landmarks.dat")

  pred := detector.UseShapePredictor(shapePred, "/go/src/detector/example/example.jpg", rect)
  log.Println(pred)

  pred2 := detector.UseShapePredictor(shapePred, "/go/src/detector/example/example.jpg", rect)
  log.Println(pred2)

  // detector filename, image filename, rectangle
  res := detector.DetectObjectsRect("shape_predictor_68_face_landmarks.dat", "/go/src/detector/example/example.jpg", rect)


  log.Println(res)

}
