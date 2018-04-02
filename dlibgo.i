/* File : dlibgo.i */
%module dlibgo

%include "std_string.i"
%include "dlib/dlib/geometry/vector.h"
%include "dlib/dlib/geometry/rectangle.h"

// %include "full_object.h"

%include "dlib-go.h"

%rename($ignore, %$isclass) ""; // Only ignore all classes
%rename("%s") Rectangle; // Unignore
%rename("%s") DetectObjectsRect; // Unignore
%rename("%s") ShapeObjects;

%{
/* Start with config.h */
#define DLIB_DISABLE_ASSERTS
#define DLIB_NO_GUI_SUPPORT
#define NO_MAKEFILE
#define DLIB_JPEG_SUPPORT
// #define DLIB_PNG_SUPPORT
#define DLIB_USE_BLAS
#define DLIB_USE_LAPACK

/* Work around */
extern "C" {
  const int USER_ERROR__inconsistent_build_configuration__see_dlib_faq_1_ = 0;
  const int USER_ERROR__inconsistent_build_configuration__see_dlib_faq_2 = 0;
  const int DLIB_CHECK_FOR_VERSION_MISMATCH = 0;
}

//
#include "dlib/dlib/base64/base64_kernel_1.cpp"
#include "dlib/dlib/bigint/bigint_kernel_1.cpp"
#include "dlib/dlib/bigint/bigint_kernel_2.cpp"
#include "dlib/dlib/bit_stream/bit_stream_kernel_1.cpp"
#include "dlib/dlib/entropy_decoder/entropy_decoder_kernel_1.cpp"
#include "dlib/dlib/entropy_decoder/entropy_decoder_kernel_2.cpp"
#include "dlib/dlib/entropy_encoder/entropy_encoder_kernel_1.cpp"
#include "dlib/dlib/entropy_encoder/entropy_encoder_kernel_2.cpp"
#include "dlib/dlib/md5/md5_kernel_1.cpp"
#include "dlib/dlib/tokenizer/tokenizer_kernel_1.cpp"
#include "dlib/dlib/unicode/unicode.cpp"


#include "dlib/dlib/geometry/vector.h"
#include "dlib/dlib/geometry/rectangle.h"
#include "dlib/dlib/image_loader/jpeg_loader.h"
#include "dlib/dlib/image_loader/png_loader.h"

#include "dlib-go.h"

// #include "full_object.h"
%}
