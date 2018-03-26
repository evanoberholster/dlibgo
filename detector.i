/* File : example.i */
%module detector
%include "/dlib-19.4/dlib/geometry/rectangle.h"
%include "/dlib-19.4/dlib/geometry/vector.h"
%include "/dlib-19.4/dlib/array2d.h"
%include "/dlib-19.4/dlib/geometry/vector_abstract.h"


%{
#include "/dlib-19.4/dlib/image_io.h"
#include "/dlib-19.4/dlib/geometry/rectangle.h"
#include "/dlib-19.4/dlib/geometry/vector_abstract.h"
#include "/dlib-19.4/dlib/image_processing/generic_image.h"
%}

%template(DRectangle) dlib::vector<dlib::vector<double, 3l> >;


/* Tell SWIG about it */
typedef dlib::vector<long,2> point;
