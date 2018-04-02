package main

import(
  "dlibgo"
  "log"
  "time"
  "os"
)

func main() {
  log.Printf("Example Create Rectangle")
  rect := dlibgo.NewRectangle__SWIG_0(240,220,580,570)

  // Load Hog file
  start := time.Now()
  pwd, _ := os.Getwd()
  spModelfile := pwd+"/models/shape_predictor_68_face_landmarks.dat"
  shapePred := dlibgo.LoadShapePredictor(spModelfile)
  log.Print("Shape Predictor Loaded: ", spModelfile)
  elapsed := time.Since(start)
  log.Printf("Loading Hog file took: %s", elapsed)

  start_pred2 := time.Now()

  // Process Image
  imageFile := pwd+"/example/example.jpg"
  pred2 := dlibgo.UseShapePredictor(shapePred, imageFile, rect)
  elapsed_pred2 := time.Since(start_pred2)
  log.Printf("Processing Image took: %s", elapsed_pred2)

  start_points := time.Now()
  newRect := pred2.GetRect()
  log.Print("Array Size: ",pred2.GetSize())
  log.Print(newRect.Width(), newRect.Height())
	for i := 0; i < pred2.GetSize(); i++ {
    log.Print("Coords// ", i, " X: ", pred2.X(i), " Y: ", pred2.Y(i))
	}

  elapsed_points := time.Since(start_points)
  log.Printf("Loading Points took: %s", elapsed_points)

}
