#include <dlib/image_processing/frontal_face_detector.h>
#include <dlib/image_processing.h>
#include <dlib/image_io.h>
#include <iostream>

namespace dlib {


full_object_detection ProcessRect(const std::string dfilename, const std::string ifilename, const rectangle rect) {
    frontal_face_detector detector = get_frontal_face_detector();
    shape_predictor sp;

    deserialize(dfilename) >> sp;

    array2d<rgb_pixel> img;
    load_image(img, ifilename);

    std::vector<rectangle> dets = detector(img);

    full_object_detection shape = sp(img, rect);
    return shape;
};

}
