#ifdef DLIB_JPEG_SUPPORT

#include "dlib/dlib/array2d.h"
#include "dlib/dlib/pixel.h"
#include "dlib/dlib/dir_nav.h"
#include "dlib/dlib/image_loader/jpeg_loader.h"
#include "dlib/dlib/image_loader/image_loader.h"
#include <stdio.h>
#include <jpeglib.h>
#include <sstream>
#include <setjmp.h>
#include <sys/stat.h>
#include <unistd.h>
#include <fcntl.h>
#include <stdlib.h>
#include <iostream>

namespace dlib {
//  template <
//      typename image_type
//      >
//  void load_jpeg (
//      image_type& image,
//      const std::string& file_name
//  )
//  {
//      jpeg_loader(file_name).get_image(image);
//  }

// Developed on the base of jpeg_loader.h (version 19.10)
// ----------------------------------------------------------------------------------------
class image_memory_loader : noncopyable
{
public:

    image_memory_loader( unsigned char* jpg_buffer, long jpg_size );

    //~image_memory_loader();
    bool is_gray() const;
    bool is_rgb() const;
    bool is_rgba() const;

    template<typename T>
    void get_image( T& t_) const
    {
#ifndef DLIB_JPEG_SUPPORT
        /* !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
            You are getting this error because you are trying to use the jpeg_loader
            object but you haven't defined DLIB_JPEG_SUPPORT.  You must do so to use
            this object.   You must also make sure you set your build environment
            to link against the libjpeg library.
        !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!*/
        COMPILE_TIME_ASSERT(sizeof(T) == 0);
#endif
        image_view<T> t(t_);
        t.set_size( height_, width_ );
        for ( unsigned n = 0; n < height_;n++ )
        {
            const unsigned char* v = get_row( n );
            for ( unsigned m = 0; m < width_;m++ )
            {
                if ( is_gray() )
                {
                    unsigned char p = v[m];
                    assign_pixel( t[n][m], p );
                }
                else if ( is_rgba() ) {
                    rgb_alpha_pixel p;
                    p.red = v[m*4];
                    p.green = v[m*4+1];
                    p.blue = v[m*4+2];
                    p.alpha = v[m*4+3];
                    assign_pixel( t[n][m], p );
                }
                else // if ( is_rgb() )
                {
                    rgb_pixel p;
                    p.red = v[m*3];
                    p.green = v[m*3+1];
                    p.blue = v[m*3+2];
                    assign_pixel( t[n][m], p );
                }
            }
        }
    }

private:
    const unsigned char* get_row( unsigned long i ) const
    {
        return &data[i*width_*output_components_];
    }

    void read_image( unsigned char* jpg_buffer, long jpg_size );
    unsigned long height_;
    unsigned long width_;
    unsigned long output_components_;
    std::vector<unsigned char> data;
};

// ----------------------------------------------------------------------------------------



// Developed on the base of jpeg_loader.cpp (version 19.10)
// ----------------------------------------------------------------------------------------

    image_memory_loader::
    image_memory_loader( unsigned char* jpg_buffer, long jpg_size ) : height_( 0 ), width_( 0 ), output_components_(0)
    {
        read_image( jpg_buffer, jpg_size );
    }

// ----------------------------------------------------------------------------------------

    bool image_memory_loader::is_gray() const
    {
        return (output_components_ == 1);
    }

// ----------------------------------------------------------------------------------------

    bool image_memory_loader::is_rgb() const
    {
        return (output_components_ == 3);
    }

// ----------------------------------------------------------------------------------------

    bool image_memory_loader::is_rgba() const
    {
        return (output_components_ == 4);
    }

// ----------------------------------------------------------------------------------------

    struct image_memory_error_mgr
    {
        jpeg_error_mgr pub;    /* "public" fields */
        jmp_buf setjmp_buffer;  /* for return to caller */
    };

    void image_memory_error_exit (j_common_ptr cinfo)
    {
        /* cinfo->err really points to a jpeg_loader_error_mgr struct, so coerce pointer */
        image_memory_error_mgr* myerr = (image_memory_error_mgr*) cinfo->err;

        /* Return control to the setjmp point */
        longjmp(myerr->setjmp_buffer, 1);
    }

// ----------------------------------------------------------------------------------------

    void image_memory_loader::read_image( unsigned char* jpg_buffer, long jpg_size )
    {
        //if ( Filesize < 100 )
        //{
        //    throw image_load_error("jpeg_loader: invalid image, it is NULL");
        //}
        //FILE *fp = fopen( filename, "rb" );
        //if ( !fp )
        //{
        //    throw image_load_error(std::string("jpeg_loader: unable to open file ") + filename);
        //}

        int rc, i, j;
        //struct stat file_info;
        //unsigned long jpg_size;
        //unsigned char *jpg_buffer;

        // ----- Load File from Disk to Memory-----

      //  rc = stat(filename, &file_info);
      //  if (rc) {
      //    throw image_load_error(std::string("jpeg_loader: unable to open file ") + filename);
      //  }
      //  jpg_size = file_info.st_size;
      //  jpg_buffer = (unsigned char*) malloc(jpg_size + 100);

      //  int fd = open(filename, O_RDONLY);
      //  i = 0;
      //  while (i < jpg_size) {
      //    rc = read(fd, jpg_buffer + i, jpg_size - i);
          //cout << "Input: Read " << rc << " " << jpg_size-i << endl;
      //    i += rc;
      //  }
      //  close(fd);
        // -!---- Load File from Disk to Memory-----

        jpeg_decompress_struct cinfo;
        jpeg_loader_error_mgr jerr;

        cinfo.err = jpeg_std_error(&jerr.pub);

        jerr.pub.error_exit = jpeg_loader_error_exit;

        /* Establish the setjmp return context for my_error_exit to use. */
        if (setjmp(jerr.setjmp_buffer))
        {
            /* If we get here, the JPEG code has signaled an error.
             * We need to clean up the JPEG object, close the input file, and return.
             */
            jpeg_destroy_decompress(&cinfo);
            //fclose(fp);
            throw image_load_error(std::string("jpeg_loader: error while reading "));
        }


        jpeg_create_decompress(&cinfo);

        // Load JPEG from memory
        jpeg_mem_src(&cinfo, jpg_buffer, jpg_size);
        // jpeg_stdio_src(&cinfo, fp);

        jpeg_read_header(&cinfo, TRUE);

        jpeg_start_decompress(&cinfo);

        height_ = cinfo.output_height;
        width_ = cinfo.output_width;
        output_components_ = cinfo.output_components;

        if (output_components_ != 1 &&
            output_components_ != 3 &&
            output_components_ != 4)
        {
            //fclose( fp );
            jpeg_destroy_decompress(&cinfo);
            std::ostringstream sout;
            sout << "jpeg_loader: Unsupported number of colors (" << output_components_ << ") in file ";
            throw image_load_error(sout.str());
        }

        std::vector<unsigned char*> rows;
        rows.resize(height_);

        // size the image buffer
        data.resize(height_*width_*output_components_);

        // setup pointers to each row
        for (unsigned long i = 0; i < rows.size(); ++i)
            rows[i] = &data[i*width_*output_components_];

        // read the data into the buffer
        while (cinfo.output_scanline < cinfo.output_height)
        {
            jpeg_read_scanlines(&cinfo, &rows[cinfo.output_scanline], 100);
        }

        jpeg_finish_decompress(&cinfo);
        jpeg_destroy_decompress(&cinfo);

        //fclose( fp );
    }


// ----------------------------------------------------------------------------------------

  template <
      typename image_type
      >
  void load_jpeg_mem (
      image_type& image,
      unsigned char* jpg_buffer,
      long jpg_size
  )
  {
      image_memory_loader(jpg_buffer, jpg_size).get_image(image);
  }

// ----------------------------------------------------------------------------------------

}
#endif // DLIB_JPEG_SUPPORT
