# Work in Progress (Not functional)

#### Build SWIG from source
1. Change configure.ac to accept Go version 1.10
2. Prepare
```shell
./autogen.sh
```
3. Configure only for Go
```shell
./configure --without-tcl --without-python --without-python3 --without-perl5 --without-octave --without-scilab --without-java --without-javascript --without-gcj --without-android --without-guile --without-mzscheme --without-ruby --without-php5 --without-php --without-ocaml --without-lua --without-csharp --without-pike --without-chicken --without-allegrocl --without-clisp --without-r --without-d
```
4. Make and Make Install
```shell
make && make install
```

#### Build Dlib from source
1. Make
```shell
cmake -DUSE_AVX_INSTRUCTIONS=ON ..
cmake --build . --config Release
ldconfig
```

#### SWIG build command
```shell
swig -go -cgo -c++ -intgosize 64
```

#### Build environment for cgo
```shell
export CGO_CXXFLAGS="-lsass -lstdc++ -lm -std=c++11"
export CGO_LDFLAGS="-ljpeg -lpng -l/usr/local/lib/libdlib.a"
```

#### Go build command
```shell
go build -x -v
```

[Dlib documentation](http://www.dlib.net)

[SWIG documentation](http://www.swig.org/Doc3.0/Go.html)
