#include "dlib/dlib/image_processing/frontal_face_detector.h"
#include "dlib/dlib/image_processing.h"
#include "dlib/dlib/image_io.h"
#include "dlib/dlib/image_loader/jpeg_loader.h"
#include "dlib/dlib/image_loader/png_loader.h"
#include <iostream>

using namespace dlib;
using namespace std;

class ShapeObjects {
  std::vector<dlib::vector<long,2>> points;
  public:
    dlib::rectangle rect;
    int sz;
    ShapeObjects(int _sz) {
      points.resize(_sz);
      sz = _sz;
    }
    void SetRectangle(dlib::rectangle _rect) {
      rect = _rect;
    }
    void SetPoint(int index, dlib::vector<long,2> val) {
      points[index] = val;
    }
    dlib::vector<long,2> GetPoint(int index) {
      return points[index];
    }
    long GetXCoord(int index) {
      return points[index].x();
    }
    long GetYCoord(int index) {
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
    dlib::load_image(img, ifilename);

    dlib::full_object_detection shape = sp(img, rect);

    ShapeObjects so(shape.num_parts());
    so.SetRectangle(shape.get_rect());

    for (int i = 0; i < shape.num_parts(); ++i)
    {
        so.SetPoint(i, shape.part(i));
    }
    return so;
};

dlib::full_object_detection DetectObjectsRect(const std::string& dfilename, const std::string& ifilename, const dlib::rectangle rect) {
    dlib::frontal_face_detector detector = dlib::get_frontal_face_detector();
    dlib::shape_predictor sp;

    deserialize(dfilename) >> sp;

    dlib::array2d<dlib::rgb_pixel> img;
    dlib::load_image(img, ifilename);

    std::vector<dlib::rectangle> dets = detector(img);

    dlib::full_object_detection shape = sp(img, rect);
    return shape;
};
