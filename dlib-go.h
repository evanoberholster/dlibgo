#include "dlib/dlib/image_processing/frontal_face_detector.h"
#include "dlib/dlib/image_processing.h"
#include "dlib/dlib/image_io.h"
#include "dlib/dlib/image_loader/jpeg_loader.h"
#include "dlib/dlib/image_loader/png_loader.h"
#include <iostream>

using namespace dlib;
using namespace std;

dlib::shape_predictor LoadShapePredictor(const std::string& dfilename) {
    dlib::shape_predictor sp;
    deserialize(dfilename) >> sp;
    return sp;
};

dlib::full_object_detection UseShapePredictor(const dlib::shape_predictor& sp, const std::string& ifilename, const dlib::rectangle rect) {
    dlib::array2d<dlib::rgb_pixel> img;
    dlib::load_image(img, ifilename);

    dlib::full_object_detection shape = sp(img, rect);
    return shape;
};

dlib::full_object_detection DetectObjectsRect(const std::string& dfilename, const std::string& ifilename, const dlib::rectangle rect) {
    dlib::frontal_face_detector detector = dlib::get_frontal_face_detector();
    dlib::shape_predictor sp;

    deserialize(dfilename) >> sp;
    // dlib::deserialize(dfilename) >> sp;

    dlib::array2d<dlib::rgb_pixel> img;
    dlib::load_image(img, ifilename);

    std::vector<dlib::rectangle> dets = detector(img);

    dlib::full_object_detection shape = sp(img, rect);
    return shape;
};
