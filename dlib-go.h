#include "dlib/dlib/image_processing/frontal_face_detector.h"
#include "dlib/dlib/image_processing.h"
#include "dlib/dlib/image_io.h"
#include "dlib/dlib/image_loader/jpeg_loader.h"
#include "dlib/dlib/image_loader/png_loader.h"
#include "image_memory.h"
#include <iostream>

using namespace dlib;
using namespace std;

class ShapeObjects {
  int sz;
  std::vector<dlib::vector<long,2>> points;
  dlib::rectangle rect;
  public:
    ShapeObjects(int _sz, dlib::rectangle _rect) {
      points.resize(_sz);
      sz = _sz;
      rect = _rect;
    }
    dlib::rectangle GetRect() {
      return rect;
    }
    int GetSize() {
      return sz;
    }
    void SetPoint(int index, dlib::vector<long,2> val) {
      points[index] = val;
    }
    dlib::vector<long,2> GetPoint(int index) {
      return points[index];
    }
    long X(int index) {
      return points[index].x();
    }
    long Y(int index) {
      return points[index].y();
    }
};

dlib::shape_predictor LoadShapePredictor(const std::string& dfilename) {
    dlib::shape_predictor sp;
    deserialize(dfilename) >> sp;
    return sp;
};

ShapeObjects UseShapePredictor(const dlib::shape_predictor& sp, const std::string& ifilename, const dlib::rectangle rect) {
    dlib::array2d<dlib::rgb_pixel> img;
    dlib::load_jpeg(img, ifilename);

    dlib::full_object_detection shape = sp(img, rect);

    ShapeObjects so(shape.num_parts(), shape.get_rect());

    for (int i = 0; i < shape.num_parts(); ++i)
    {
        so.SetPoint(i, shape.part(i));
    }
    return so;
};

ShapeObjects UseShapePredictorByte(const dlib::shape_predictor& sp, unsigned char* jpg_buffer, long jpg_size, const dlib::rectangle rect) {

    dlib::array2d<dlib::rgb_pixel> img;
    dlib::load_jpeg_mem(img, jpg_buffer, jpg_size);

    dlib::full_object_detection shape = sp(img, rect);

    ShapeObjects so(shape.num_parts(), shape.get_rect());

    for (int i = 0; i < shape.num_parts(); ++i)
    {
        so.SetPoint(i, shape.part(i));
    }

    return so;
};

long TestMemoryFile(unsigned char* jpg_buffer, long jpg_size) {

    dlib::array2d<dlib::rgb_pixel> img;
    dlib::load_jpeg_mem(img, jpg_buffer, jpg_size);

    int fd = open("/go/src/detector/output.jpg", O_CREAT | O_WRONLY, 0666);
    write(fd, jpg_buffer, jpg_size); // Write out all RGB pixel data
    close(fd);
    //free(jpg_buffer);
    return jpg_size;
};

dlib::full_object_detection DetectObjectsRect(const std::string& dfilename, const std::string& ifilename, const dlib::rectangle rect) {
    dlib::frontal_face_detector detector = dlib::get_frontal_face_detector();
    dlib::shape_predictor sp;

    deserialize(dfilename) >> sp;

    dlib::array2d<dlib::rgb_pixel> img;
    dlib::load_jpeg(img, ifilename);

    std::vector<dlib::rectangle> dets = detector(img);

    dlib::full_object_detection shape = sp(img, rect);
    return shape;
};
