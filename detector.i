/* File : example.i */
%module detector
%include "dlib/dlib/geometry/rectangle.h"
%include "dlib/dlib/geometry/vector.h"
%include "dlib/dlib/array2d.h"
%include "dlib/dlib/geometry/vector_abstract.h"


%{
#include "dlib/dlib/image_io.h"
#include "dlib/dlib/geometry/rectangle.h"
#include "dlib/dlib/geometry/vector_abstract.h"
#include "dlib/dlib/image_processing/generic_image.h"
%}

%template(DRectangle) dlib::vector<dlib::vector<double, 3l> >;



/* Tell SWIG about it */
typedef dlib::vector<long,2> point;
