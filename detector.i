/* File : detector.i */
%module detector

%include "/usr/local/include/dlib/geometry/vector.h"
%include "/usr/local/include/dlib/geometry/rectangle.h"
%include "/usr/local/include/dlib/image_processing/full_object_detection.h"

%include "dlib-go.h"

%{
/* Start with config.h */
#define DLIB_DISABLE_ASSERTS
#define DLIB_NO_GUI_SUPPORT
#define DLIB_JPEG_SUPPORT
#define DLIB_PNG_SUPPORT
#define DLIB_USE_BLAS
#define DLIB_USE_LAPACK
#define DLIB_CHECK_FOR_VERSION_MISMATCH DLIB_VERSION_MISMATCH_CHECK__EXPECTED_VERSION_19_10_0

/* Work around */
extern "C" {
  const int USER_ERROR__inconsistent_build_configuration__see_dlib_faq_1_ = 0;
  const int DLIB_CHECK_FOR_VERSION_MISMATCH = 0;
}

#include "/usr/local/include/dlib/geometry/vector.h"
#include "/usr/local/include/dlib/geometry/rectangle.h"
#include "/usr/local/include/dlib/image_processing/full_object_detection.h"

#include "dlib-go.h"

typedef dlib::vector<long,2> point;
%}

/* Tell SWIG about it */
typedef dlib::vector<long,2> point;
