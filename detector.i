/* File : example.i */
%module detector

%rename($ignore, %$isclass) ""; // Only ignore all classes
%rename("%s") Rectangle; // Unignore
%rename("%s") Vector; // Unignore
%rename("%s") Point; // Unignore

%include "/usr/local/include/dlib/geometry/rectangle.h"
%include "/usr/local/include/dlib/geometry/vector.h"


%{
#include "/usr/local/include/dlib/geometry/rectangle.h"
#include "/usr/local/include/dlib/geometry/vector.h"

typedef dlib::vector<long,2> point;

%}


/* Tell SWIG about it */
typedef dlib::vector<long,2> point;
