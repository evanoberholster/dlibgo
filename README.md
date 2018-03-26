# Work in Progress (Not functional)

#### Build SWIG from source
1.  - Change configure.ac to accept Go1.10
2. Prepare
```
./autogen.sh
```
3. Configure
```
./configure --without-tcl --without-python --without-python3 --without-perl5 --without-octave --without-scilab --without-java --without-javascript --without-gcj --without-android --without-guile --without-mzscheme --without-ruby --without-php5 --without-php --without-ocaml --without-lua --without-csharp --without-pike --without-chicken --without-allegrocl --without-clisp --without-r --without-d
```
4. Make and Make Install
```
make && make install
```


#### Build environment
```
export CGO_CXXFLAGS="-lsass -lstdc++ -lm -std=c++11"
```

#### SWIG build command
```
swig -go -cgo -c++ -intgosize 64
```

#### Go build command
```
go build -x -v
```

[Dlib documentation](http://www.dlib.net)
[SWIG documentation](http://www.swig.org/Doc3.0/Go.html)
