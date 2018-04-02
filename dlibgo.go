/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.12
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: dlibgo.i

package dlibgo

/*
#define intgo swig_intgo
typedef void *swig_voidp;

#include <stdint.h>


typedef long long intgo;
typedef unsigned long long uintgo;



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;


typedef long long swig_type_1;
typedef long long swig_type_2;
typedef long long swig_type_3;
typedef long long swig_type_4;
typedef long long swig_type_5;
typedef long long swig_type_6;
typedef long long swig_type_7;
typedef long long swig_type_8;
typedef long long swig_type_9;
typedef long long swig_type_10;
typedef long long swig_type_11;
typedef long long swig_type_12;
typedef long long swig_type_13;
typedef long long swig_type_14;
typedef long long swig_type_15;
typedef long long swig_type_16;
typedef long long swig_type_17;
typedef long long swig_type_18;
typedef long long swig_type_19;
typedef long long swig_type_20;
typedef long long swig_type_21;
typedef long long swig_type_22;
typedef long long swig_type_23;
typedef long long swig_type_24;
typedef long long swig_type_25;
typedef long long swig_type_26;
typedef long long swig_type_27;
typedef long long swig_type_28;
typedef long long swig_type_29;
typedef long long swig_type_30;
typedef long long swig_type_31;
typedef long long swig_type_32;
typedef long long swig_type_33;
typedef long long swig_type_34;
typedef long long swig_type_35;
typedef long long swig_type_36;
typedef long long swig_type_37;
typedef long long swig_type_38;
typedef long long swig_type_39;
typedef long long swig_type_40;
typedef long long swig_type_41;
typedef long long swig_type_42;
typedef long long swig_type_43;
typedef long long swig_type_44;
typedef long long swig_type_45;
typedef long long swig_type_46;
typedef long long swig_type_47;
typedef _gostring_ swig_type_48;
typedef _gostring_ swig_type_49;
typedef long long swig_type_50;
typedef long long swig_type_51;
typedef long long swig_type_52;
typedef _gostring_ swig_type_53;
typedef _gostring_ swig_type_54;
extern void _wrap_Swig_free_dlibgo_63b7eba619e1d075(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_dlibgo_63b7eba619e1d075(swig_intgo arg1);
extern uintptr_t _wrap_new_rectangle__SWIG_0_dlibgo_63b7eba619e1d075(swig_type_1 arg1, swig_type_2 arg2, swig_type_3 arg3, swig_type_4 arg4);
extern uintptr_t _wrap_new_rectangle__SWIG_1_dlibgo_63b7eba619e1d075(swig_type_5 arg1, swig_type_6 arg2);
extern uintptr_t _wrap_new_rectangle__SWIG_2_dlibgo_63b7eba619e1d075(uintptr_t arg1);
extern uintptr_t _wrap_new_rectangle__SWIG_3_dlibgo_63b7eba619e1d075(uintptr_t arg1, uintptr_t arg2);
extern uintptr_t _wrap_new_rectangle__SWIG_5_dlibgo_63b7eba619e1d075(void);
extern swig_type_7 _wrap_rectangle_top__SWIG_0_dlibgo_63b7eba619e1d075(uintptr_t arg1);
extern swig_voidp _wrap_rectangle_top__SWIG_1_dlibgo_63b7eba619e1d075(uintptr_t arg1);
extern void _wrap_rectangle_set_top_dlibgo_63b7eba619e1d075(uintptr_t arg1, swig_type_8 arg2);
extern swig_type_9 _wrap_rectangle_left__SWIG_0_dlibgo_63b7eba619e1d075(uintptr_t arg1);
extern swig_voidp _wrap_rectangle_left__SWIG_1_dlibgo_63b7eba619e1d075(uintptr_t arg1);
extern void _wrap_rectangle_set_left_dlibgo_63b7eba619e1d075(uintptr_t arg1, swig_type_10 arg2);
extern swig_type_11 _wrap_rectangle_right__SWIG_0_dlibgo_63b7eba619e1d075(uintptr_t arg1);
extern swig_voidp _wrap_rectangle_right__SWIG_1_dlibgo_63b7eba619e1d075(uintptr_t arg1);
extern void _wrap_rectangle_set_right_dlibgo_63b7eba619e1d075(uintptr_t arg1, swig_type_12 arg2);
extern swig_type_13 _wrap_rectangle_bottom__SWIG_0_dlibgo_63b7eba619e1d075(uintptr_t arg1);
extern swig_voidp _wrap_rectangle_bottom__SWIG_1_dlibgo_63b7eba619e1d075(uintptr_t arg1);
extern void _wrap_rectangle_set_bottom_dlibgo_63b7eba619e1d075(uintptr_t arg1, swig_type_14 arg2);
extern uintptr_t _wrap_rectangle_tl_corner_dlibgo_63b7eba619e1d075(uintptr_t arg1);
extern uintptr_t _wrap_rectangle_bl_corner_dlibgo_63b7eba619e1d075(uintptr_t arg1);
extern uintptr_t _wrap_rectangle_tr_corner_dlibgo_63b7eba619e1d075(uintptr_t arg1);
extern uintptr_t _wrap_rectangle_br_corner_dlibgo_63b7eba619e1d075(uintptr_t arg1);
extern swig_type_15 _wrap_rectangle_width_dlibgo_63b7eba619e1d075(uintptr_t arg1);
extern swig_type_16 _wrap_rectangle_height_dlibgo_63b7eba619e1d075(uintptr_t arg1);
extern swig_type_17 _wrap_rectangle_area_dlibgo_63b7eba619e1d075(uintptr_t arg1);
extern _Bool _wrap_rectangle_is_empty_dlibgo_63b7eba619e1d075(uintptr_t arg1);
extern uintptr_t _wrap_rectangle_intersect_dlibgo_63b7eba619e1d075(uintptr_t arg1, uintptr_t arg2);
extern _Bool _wrap_rectangle_contains__SWIG_0_dlibgo_63b7eba619e1d075(uintptr_t arg1, uintptr_t arg2);
extern _Bool _wrap_rectangle_contains__SWIG_1_dlibgo_63b7eba619e1d075(uintptr_t arg1, swig_type_18 arg2, swig_type_19 arg3);
extern _Bool _wrap_rectangle_contains__SWIG_2_dlibgo_63b7eba619e1d075(uintptr_t arg1, uintptr_t arg2);
extern void _wrap_delete_rectangle_dlibgo_63b7eba619e1d075(uintptr_t arg1);
extern void _wrap_serialize_dlibgo_63b7eba619e1d075(uintptr_t arg1, uintptr_t arg2);
extern void _wrap_deserialize_dlibgo_63b7eba619e1d075(uintptr_t arg1, uintptr_t arg2);
extern uintptr_t _wrap_centered_rect__SWIG_0_dlibgo_63b7eba619e1d075(swig_type_20 arg1, swig_type_21 arg2, swig_type_22 arg3, swig_type_23 arg4);
extern uintptr_t _wrap_intersect_dlibgo_63b7eba619e1d075(uintptr_t arg1, uintptr_t arg2);
extern swig_type_24 _wrap_area_dlibgo_63b7eba619e1d075(uintptr_t arg1);
extern uintptr_t _wrap_center_dlibgo_63b7eba619e1d075(uintptr_t arg1);
extern uintptr_t _wrap_dcenter_dlibgo_63b7eba619e1d075(uintptr_t arg1);
extern swig_type_25 _wrap_distance_to_rect_edge_dlibgo_63b7eba619e1d075(uintptr_t arg1, uintptr_t arg2);
extern uintptr_t _wrap_nearest_point_dlibgo_63b7eba619e1d075(uintptr_t arg1, uintptr_t arg2);
extern swig_type_26 _wrap_nearest_rect_dlibgo_63b7eba619e1d075(uintptr_t arg1, uintptr_t arg2);
extern void _wrap_clip_line_to_rectangle_dlibgo_63b7eba619e1d075(uintptr_t arg1, uintptr_t arg2, uintptr_t arg3);
extern uintptr_t _wrap_centered_rect__SWIG_1_dlibgo_63b7eba619e1d075(uintptr_t arg1, swig_type_27 arg2, swig_type_28 arg3);
extern uintptr_t _wrap_centered_rect__SWIG_2_dlibgo_63b7eba619e1d075(uintptr_t arg1, swig_type_29 arg2, swig_type_30 arg3);
extern uintptr_t _wrap_shrink_rect__SWIG_0_dlibgo_63b7eba619e1d075(uintptr_t arg1, swig_type_31 arg2);
extern uintptr_t _wrap_grow_rect__SWIG_0_dlibgo_63b7eba619e1d075(uintptr_t arg1, swig_type_32 arg2);
extern uintptr_t _wrap_shrink_rect__SWIG_1_dlibgo_63b7eba619e1d075(uintptr_t arg1, swig_type_33 arg2, swig_type_34 arg3);
extern uintptr_t _wrap_grow_rect__SWIG_1_dlibgo_63b7eba619e1d075(uintptr_t arg1, swig_type_35 arg2, swig_type_36 arg3);
extern uintptr_t _wrap_translate_rect__SWIG_0_dlibgo_63b7eba619e1d075(uintptr_t arg1, uintptr_t arg2);
extern uintptr_t _wrap_translate_rect__SWIG_1_dlibgo_63b7eba619e1d075(uintptr_t arg1, swig_type_37 arg2, swig_type_38 arg3);
extern uintptr_t _wrap_resize_rect_dlibgo_63b7eba619e1d075(uintptr_t arg1, swig_type_39 arg2, swig_type_40 arg3);
extern uintptr_t _wrap_resize_rect_width_dlibgo_63b7eba619e1d075(uintptr_t arg1, swig_type_41 arg2);
extern uintptr_t _wrap_resize_rect_height_dlibgo_63b7eba619e1d075(uintptr_t arg1, swig_type_42 arg2);
extern uintptr_t _wrap_move_rect__SWIG_0_dlibgo_63b7eba619e1d075(uintptr_t arg1, uintptr_t arg2);
extern uintptr_t _wrap_move_rect__SWIG_1_dlibgo_63b7eba619e1d075(uintptr_t arg1, swig_type_43 arg2, swig_type_44 arg3);
extern uintptr_t _wrap_set_rect_area_dlibgo_63b7eba619e1d075(uintptr_t arg1, swig_type_45 arg2);
extern uintptr_t _wrap_set_aspect_ratio_dlibgo_63b7eba619e1d075(uintptr_t arg1, double arg2);
extern uintptr_t _wrap_new_ShapeObjects_dlibgo_63b7eba619e1d075(swig_intgo arg1, uintptr_t arg2);
extern uintptr_t _wrap_ShapeObjects_GetRect_dlibgo_63b7eba619e1d075(uintptr_t arg1);
extern swig_intgo _wrap_ShapeObjects_GetSize_dlibgo_63b7eba619e1d075(uintptr_t arg1);
extern void _wrap_ShapeObjects_SetPoint_dlibgo_63b7eba619e1d075(uintptr_t arg1, swig_intgo arg2, uintptr_t arg3);
extern uintptr_t _wrap_ShapeObjects_GetPoint_dlibgo_63b7eba619e1d075(uintptr_t arg1, swig_intgo arg2);
extern swig_type_46 _wrap_ShapeObjects_X_dlibgo_63b7eba619e1d075(uintptr_t arg1, swig_intgo arg2);
extern swig_type_47 _wrap_ShapeObjects_Y_dlibgo_63b7eba619e1d075(uintptr_t arg1, swig_intgo arg2);
extern void _wrap_delete_ShapeObjects_dlibgo_63b7eba619e1d075(uintptr_t arg1);
extern uintptr_t _wrap_LoadShapePredictor_dlibgo_63b7eba619e1d075(swig_type_48 arg1);
extern uintptr_t _wrap_UseShapePredictor_dlibgo_63b7eba619e1d075(uintptr_t arg1, swig_type_49 arg2, uintptr_t arg3);
extern uintptr_t _wrap_UseShapePredictorByte_dlibgo_63b7eba619e1d075(uintptr_t arg1, swig_voidp arg2, swig_type_50 arg3, uintptr_t arg4);
extern swig_type_51 _wrap_TestMemoryFile_dlibgo_63b7eba619e1d075(swig_voidp arg1, swig_type_52 arg2);
extern uintptr_t _wrap_DetectObjectsRect_dlibgo_63b7eba619e1d075(swig_type_53 arg1, swig_type_54 arg2, uintptr_t arg3);
#undef intgo
*/
// #cgo CXXFLAGS: -lsass -lstdc++ -lm -std=c++11 -ljpeg -lpng -llapack -lblas
// #cgo LDFLAGS: -ljpeg -lpng -llapack -lblas
import "C"

import "unsafe"
import _ "runtime/cgo"
import "sync"


type _ unsafe.Pointer



var Swig_escape_always_false bool
var Swig_escape_val interface{}


type _swig_fnptr *byte
type _swig_memberptr *byte


type _ sync.Mutex

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_dlibgo_63b7eba619e1d075(C.swig_intgo(_swig_i_0)))
	return swig_r
}

type SwigcptrRectangle uintptr

func (p SwigcptrRectangle) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrRectangle) SwigIsRectangle() {
}

func NewRectangle__SWIG_0(arg1 int64, arg2 int64, arg3 int64, arg4 int64) (_swig_ret Rectangle) {
	var swig_r Rectangle
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	swig_r = (Rectangle)(SwigcptrRectangle(C._wrap_new_rectangle__SWIG_0_dlibgo_63b7eba619e1d075(C.swig_type_1(_swig_i_0), C.swig_type_2(_swig_i_1), C.swig_type_3(_swig_i_2), C.swig_type_4(_swig_i_3))))
	return swig_r
}

func NewRectangle__SWIG_1(arg1 uint64, arg2 uint64) (_swig_ret Rectangle) {
	var swig_r Rectangle
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (Rectangle)(SwigcptrRectangle(C._wrap_new_rectangle__SWIG_1_dlibgo_63b7eba619e1d075(C.swig_type_5(_swig_i_0), C.swig_type_6(_swig_i_1))))
	return swig_r
}

func NewRectangle__SWIG_2(arg1 Dlib_vector_Sl_long_Sc_2_Sg_) (_swig_ret Rectangle) {
	var swig_r Rectangle
	_swig_i_0 := arg1.Swigcptr()
	swig_r = (Rectangle)(SwigcptrRectangle(C._wrap_new_rectangle__SWIG_2_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func NewRectangle__SWIG_3(arg1 Dlib_vector_Sl_long_Sc_2_Sg_, arg2 Dlib_vector_Sl_long_Sc_2_Sg_) (_swig_ret Rectangle) {
	var swig_r Rectangle
	_swig_i_0 := arg1.Swigcptr()
	_swig_i_1 := arg2.Swigcptr()
	swig_r = (Rectangle)(SwigcptrRectangle(C._wrap_new_rectangle__SWIG_3_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))))
	return swig_r
}

func NewRectangle__SWIG_5() (_swig_ret Rectangle) {
	var swig_r Rectangle
	swig_r = (Rectangle)(SwigcptrRectangle(C._wrap_new_rectangle__SWIG_5_dlibgo_63b7eba619e1d075()))
	return swig_r
}

func NewRectangle(a ...interface{}) Rectangle {
	argc := len(a)
	if argc == 0 {
		return NewRectangle__SWIG_5()
	}
	if argc == 1 {
		return NewRectangle__SWIG_2(a[0].(Dlib_vector_Sl_long_Sc_2_Sg_))
	}
	if argc == 2 {
		if _, ok := a[0].(SwigcptrDlib_vector_Sl_long_Sc_2_Sg_); !ok {
			goto check_3
		}
		if _, ok := a[1].(SwigcptrDlib_vector_Sl_long_Sc_2_Sg_); !ok {
			goto check_3
		}
		return NewRectangle__SWIG_3(a[0].(Dlib_vector_Sl_long_Sc_2_Sg_), a[1].(Dlib_vector_Sl_long_Sc_2_Sg_))
	}
check_3:
	if argc == 2 {
		return NewRectangle__SWIG_1(a[0].(uint64), a[1].(uint64))
	}
	if argc == 4 {
		return NewRectangle__SWIG_0(a[0].(int64), a[1].(int64), a[2].(int64), a[3].(int64))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrRectangle) Top__SWIG_0() (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	swig_r = (int64)(C._wrap_rectangle_top__SWIG_0_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrRectangle) Top__SWIG_1() (_swig_ret *int64) {
	var swig_r *int64
	_swig_i_0 := arg1
	swig_r = (*int64)(C._wrap_rectangle_top__SWIG_1_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (p SwigcptrRectangle) Top(a ...interface{}) interface{} {
	argc := len(a)
	if argc == 0 {
		return p.Top__SWIG_0()
	}
	if argc == 0 {
		return p.Top__SWIG_1()
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrRectangle) Set_top(arg2 int64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_rectangle_set_top_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0), C.swig_type_8(_swig_i_1))
}

func (arg1 SwigcptrRectangle) Left__SWIG_0() (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	swig_r = (int64)(C._wrap_rectangle_left__SWIG_0_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrRectangle) Left__SWIG_1() (_swig_ret *int64) {
	var swig_r *int64
	_swig_i_0 := arg1
	swig_r = (*int64)(C._wrap_rectangle_left__SWIG_1_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (p SwigcptrRectangle) Left(a ...interface{}) interface{} {
	argc := len(a)
	if argc == 0 {
		return p.Left__SWIG_0()
	}
	if argc == 0 {
		return p.Left__SWIG_1()
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrRectangle) Set_left(arg2 int64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_rectangle_set_left_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0), C.swig_type_10(_swig_i_1))
}

func (arg1 SwigcptrRectangle) Right__SWIG_0() (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	swig_r = (int64)(C._wrap_rectangle_right__SWIG_0_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrRectangle) Right__SWIG_1() (_swig_ret *int64) {
	var swig_r *int64
	_swig_i_0 := arg1
	swig_r = (*int64)(C._wrap_rectangle_right__SWIG_1_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (p SwigcptrRectangle) Right(a ...interface{}) interface{} {
	argc := len(a)
	if argc == 0 {
		return p.Right__SWIG_0()
	}
	if argc == 0 {
		return p.Right__SWIG_1()
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrRectangle) Set_right(arg2 int64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_rectangle_set_right_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0), C.swig_type_12(_swig_i_1))
}

func (arg1 SwigcptrRectangle) Bottom__SWIG_0() (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	swig_r = (int64)(C._wrap_rectangle_bottom__SWIG_0_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrRectangle) Bottom__SWIG_1() (_swig_ret *int64) {
	var swig_r *int64
	_swig_i_0 := arg1
	swig_r = (*int64)(C._wrap_rectangle_bottom__SWIG_1_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (p SwigcptrRectangle) Bottom(a ...interface{}) interface{} {
	argc := len(a)
	if argc == 0 {
		return p.Bottom__SWIG_0()
	}
	if argc == 0 {
		return p.Bottom__SWIG_1()
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrRectangle) Set_bottom(arg2 int64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_rectangle_set_bottom_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0), C.swig_type_14(_swig_i_1))
}

func (arg1 SwigcptrRectangle) Tl_corner() (_swig_ret Dlib_vector_Sl_long_Sc_2_Sg_) {
	var swig_r Dlib_vector_Sl_long_Sc_2_Sg_
	_swig_i_0 := arg1
	swig_r = (Dlib_vector_Sl_long_Sc_2_Sg_)(SwigcptrDlib_vector_Sl_long_Sc_2_Sg_(C._wrap_rectangle_tl_corner_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func (arg1 SwigcptrRectangle) Bl_corner() (_swig_ret Dlib_vector_Sl_long_Sc_2_Sg_) {
	var swig_r Dlib_vector_Sl_long_Sc_2_Sg_
	_swig_i_0 := arg1
	swig_r = (Dlib_vector_Sl_long_Sc_2_Sg_)(SwigcptrDlib_vector_Sl_long_Sc_2_Sg_(C._wrap_rectangle_bl_corner_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func (arg1 SwigcptrRectangle) Tr_corner() (_swig_ret Dlib_vector_Sl_long_Sc_2_Sg_) {
	var swig_r Dlib_vector_Sl_long_Sc_2_Sg_
	_swig_i_0 := arg1
	swig_r = (Dlib_vector_Sl_long_Sc_2_Sg_)(SwigcptrDlib_vector_Sl_long_Sc_2_Sg_(C._wrap_rectangle_tr_corner_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func (arg1 SwigcptrRectangle) Br_corner() (_swig_ret Dlib_vector_Sl_long_Sc_2_Sg_) {
	var swig_r Dlib_vector_Sl_long_Sc_2_Sg_
	_swig_i_0 := arg1
	swig_r = (Dlib_vector_Sl_long_Sc_2_Sg_)(SwigcptrDlib_vector_Sl_long_Sc_2_Sg_(C._wrap_rectangle_br_corner_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func (arg1 SwigcptrRectangle) Width() (_swig_ret uint64) {
	var swig_r uint64
	_swig_i_0 := arg1
	swig_r = (uint64)(C._wrap_rectangle_width_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrRectangle) Height() (_swig_ret uint64) {
	var swig_r uint64
	_swig_i_0 := arg1
	swig_r = (uint64)(C._wrap_rectangle_height_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrRectangle) Area() (_swig_ret uint64) {
	var swig_r uint64
	_swig_i_0 := arg1
	swig_r = (uint64)(C._wrap_rectangle_area_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrRectangle) Is_empty() (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	swig_r = (bool)(C._wrap_rectangle_is_empty_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrRectangle) Intersect(arg2 Rectangle) (_swig_ret Rectangle) {
	var swig_r Rectangle
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	swig_r = (Rectangle)(SwigcptrRectangle(C._wrap_rectangle_intersect_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))))
	return swig_r
}

func (arg1 SwigcptrRectangle) Contains__SWIG_0(arg2 Dlib_vector_Sl_long_Sc_2_Sg_) (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	swig_r = (bool)(C._wrap_rectangle_contains__SWIG_0_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1)))
	return swig_r
}

func (arg1 SwigcptrRectangle) Contains__SWIG_1(arg2 int64, arg3 int64) (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	swig_r = (bool)(C._wrap_rectangle_contains__SWIG_1_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0), C.swig_type_18(_swig_i_1), C.swig_type_19(_swig_i_2)))
	return swig_r
}

func (arg1 SwigcptrRectangle) Contains__SWIG_2(arg2 Rectangle) (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	swig_r = (bool)(C._wrap_rectangle_contains__SWIG_2_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1)))
	return swig_r
}

func (p SwigcptrRectangle) Contains(a ...interface{}) bool {
	argc := len(a)
	if argc == 1 {
		if _, ok := a[0].(SwigcptrDlib_vector_Sl_long_Sc_2_Sg_); !ok {
			goto check_1
		}
		return p.Contains__SWIG_0(a[0].(Dlib_vector_Sl_long_Sc_2_Sg_))
	}
check_1:
	if argc == 1 {
		return p.Contains__SWIG_2(a[0].(Rectangle))
	}
	if argc == 2 {
		return p.Contains__SWIG_1(a[0].(int64), a[1].(int64))
	}
	panic("No match for overloaded function call")
}

func DeleteRectangle(arg1 Rectangle) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_rectangle_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0))
}

type Rectangle interface {
	Swigcptr() uintptr
	SwigIsRectangle()
	Top(a ...interface{}) interface{}
	Set_top(arg2 int64)
	Left(a ...interface{}) interface{}
	Set_left(arg2 int64)
	Right(a ...interface{}) interface{}
	Set_right(arg2 int64)
	Bottom(a ...interface{}) interface{}
	Set_bottom(arg2 int64)
	Tl_corner() (_swig_ret Dlib_vector_Sl_long_Sc_2_Sg_)
	Bl_corner() (_swig_ret Dlib_vector_Sl_long_Sc_2_Sg_)
	Tr_corner() (_swig_ret Dlib_vector_Sl_long_Sc_2_Sg_)
	Br_corner() (_swig_ret Dlib_vector_Sl_long_Sc_2_Sg_)
	Width() (_swig_ret uint64)
	Height() (_swig_ret uint64)
	Area() (_swig_ret uint64)
	Is_empty() (_swig_ret bool)
	Intersect(arg2 Rectangle) (_swig_ret Rectangle)
	Contains(a ...interface{}) bool
}

func Serialize(arg1 Rectangle, arg2 Std_ostream) {
	_swig_i_0 := arg1.Swigcptr()
	_swig_i_1 := arg2.Swigcptr()
	C._wrap_serialize_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))
}

func Deserialize(arg1 Rectangle, arg2 Std_istream) {
	_swig_i_0 := arg1.Swigcptr()
	_swig_i_1 := arg2.Swigcptr()
	C._wrap_deserialize_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))
}

func Centered_rect__SWIG_0(arg1 int64, arg2 int64, arg3 uint64, arg4 uint64) (_swig_ret Rectangle) {
	var swig_r Rectangle
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	swig_r = (Rectangle)(SwigcptrRectangle(C._wrap_centered_rect__SWIG_0_dlibgo_63b7eba619e1d075(C.swig_type_20(_swig_i_0), C.swig_type_21(_swig_i_1), C.swig_type_22(_swig_i_2), C.swig_type_23(_swig_i_3))))
	return swig_r
}

func Intersect(arg1 Rectangle, arg2 Rectangle) (_swig_ret Rectangle) {
	var swig_r Rectangle
	_swig_i_0 := arg1.Swigcptr()
	_swig_i_1 := arg2.Swigcptr()
	swig_r = (Rectangle)(SwigcptrRectangle(C._wrap_intersect_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))))
	return swig_r
}

func Area(arg1 Rectangle) (_swig_ret uint64) {
	var swig_r uint64
	_swig_i_0 := arg1.Swigcptr()
	swig_r = (uint64)(C._wrap_area_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func Center(arg1 Rectangle) (_swig_ret Dlib_vector_Sl_long_Sc_2_Sg_) {
	var swig_r Dlib_vector_Sl_long_Sc_2_Sg_
	_swig_i_0 := arg1.Swigcptr()
	swig_r = (Dlib_vector_Sl_long_Sc_2_Sg_)(SwigcptrDlib_vector_Sl_long_Sc_2_Sg_(C._wrap_center_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func Dcenter(arg1 Rectangle) (_swig_ret Dlib_vector_Sl_double_Sc_2_Sg_) {
	var swig_r Dlib_vector_Sl_double_Sc_2_Sg_
	_swig_i_0 := arg1.Swigcptr()
	swig_r = (Dlib_vector_Sl_double_Sc_2_Sg_)(SwigcptrDlib_vector_Sl_double_Sc_2_Sg_(C._wrap_dcenter_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func Distance_to_rect_edge(arg1 Rectangle, arg2 Dlib_vector_Sl_long_Sc_2_Sg_) (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1.Swigcptr()
	_swig_i_1 := arg2.Swigcptr()
	swig_r = (int64)(C._wrap_distance_to_rect_edge_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1)))
	return swig_r
}

func Nearest_point(arg1 Rectangle, arg2 Dlib_vector_Sl_long_Sc_2_Sg_) (_swig_ret Dlib_vector_Sl_long_Sc_2_Sg_) {
	var swig_r Dlib_vector_Sl_long_Sc_2_Sg_
	_swig_i_0 := arg1.Swigcptr()
	_swig_i_1 := arg2.Swigcptr()
	swig_r = (Dlib_vector_Sl_long_Sc_2_Sg_)(SwigcptrDlib_vector_Sl_long_Sc_2_Sg_(C._wrap_nearest_point_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))))
	return swig_r
}

func Nearest_rect(arg1 Std_vector_Sl_dlib_rectangle_Sg_, arg2 Dlib_vector_Sl_long_Sc_2_Sg_) (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1.Swigcptr()
	_swig_i_1 := arg2.Swigcptr()
	swig_r = (int64)(C._wrap_nearest_rect_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1)))
	return swig_r
}

func Clip_line_to_rectangle(arg1 Rectangle, arg2 Dlib_vector_Sl_long_Sc_2_Sg_, arg3 Dlib_vector_Sl_long_Sc_2_Sg_) {
	_swig_i_0 := arg1.Swigcptr()
	_swig_i_1 := arg2.Swigcptr()
	_swig_i_2 := arg3.Swigcptr()
	C._wrap_clip_line_to_rectangle_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1), C.uintptr_t(_swig_i_2))
}

func Centered_rect__SWIG_1(arg1 Dlib_vector_Sl_long_Sc_2_Sg_, arg2 uint64, arg3 uint64) (_swig_ret Rectangle) {
	var swig_r Rectangle
	_swig_i_0 := arg1.Swigcptr()
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	swig_r = (Rectangle)(SwigcptrRectangle(C._wrap_centered_rect__SWIG_1_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0), C.swig_type_27(_swig_i_1), C.swig_type_28(_swig_i_2))))
	return swig_r
}

func Centered_rect__SWIG_2(arg1 Rectangle, arg2 uint64, arg3 uint64) (_swig_ret Rectangle) {
	var swig_r Rectangle
	_swig_i_0 := arg1.Swigcptr()
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	swig_r = (Rectangle)(SwigcptrRectangle(C._wrap_centered_rect__SWIG_2_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0), C.swig_type_29(_swig_i_1), C.swig_type_30(_swig_i_2))))
	return swig_r
}

func Centered_rect(a ...interface{}) Rectangle {
	argc := len(a)
	if argc == 3 {
		if _, ok := a[0].(SwigcptrDlib_vector_Sl_long_Sc_2_Sg_); !ok {
			goto check_1
		}
		return Centered_rect__SWIG_1(a[0].(Dlib_vector_Sl_long_Sc_2_Sg_), a[1].(uint64), a[2].(uint64))
	}
check_1:
	if argc == 3 {
		return Centered_rect__SWIG_2(a[0].(Rectangle), a[1].(uint64), a[2].(uint64))
	}
	if argc == 4 {
		return Centered_rect__SWIG_0(a[0].(int64), a[1].(int64), a[2].(uint64), a[3].(uint64))
	}
	panic("No match for overloaded function call")
}

func Shrink_rect__SWIG_0(arg1 Rectangle, arg2 int64) (_swig_ret Rectangle) {
	var swig_r Rectangle
	_swig_i_0 := arg1.Swigcptr()
	_swig_i_1 := arg2
	swig_r = (Rectangle)(SwigcptrRectangle(C._wrap_shrink_rect__SWIG_0_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0), C.swig_type_31(_swig_i_1))))
	return swig_r
}

func Grow_rect__SWIG_0(arg1 Rectangle, arg2 int64) (_swig_ret Rectangle) {
	var swig_r Rectangle
	_swig_i_0 := arg1.Swigcptr()
	_swig_i_1 := arg2
	swig_r = (Rectangle)(SwigcptrRectangle(C._wrap_grow_rect__SWIG_0_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0), C.swig_type_32(_swig_i_1))))
	return swig_r
}

func Shrink_rect__SWIG_1(arg1 Rectangle, arg2 int64, arg3 int64) (_swig_ret Rectangle) {
	var swig_r Rectangle
	_swig_i_0 := arg1.Swigcptr()
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	swig_r = (Rectangle)(SwigcptrRectangle(C._wrap_shrink_rect__SWIG_1_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0), C.swig_type_33(_swig_i_1), C.swig_type_34(_swig_i_2))))
	return swig_r
}

func Shrink_rect(a ...interface{}) Rectangle {
	argc := len(a)
	if argc == 2 {
		return Shrink_rect__SWIG_0(a[0].(Rectangle), a[1].(int64))
	}
	if argc == 3 {
		return Shrink_rect__SWIG_1(a[0].(Rectangle), a[1].(int64), a[2].(int64))
	}
	panic("No match for overloaded function call")
}

func Grow_rect__SWIG_1(arg1 Rectangle, arg2 int64, arg3 int64) (_swig_ret Rectangle) {
	var swig_r Rectangle
	_swig_i_0 := arg1.Swigcptr()
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	swig_r = (Rectangle)(SwigcptrRectangle(C._wrap_grow_rect__SWIG_1_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0), C.swig_type_35(_swig_i_1), C.swig_type_36(_swig_i_2))))
	return swig_r
}

func Grow_rect(a ...interface{}) Rectangle {
	argc := len(a)
	if argc == 2 {
		return Grow_rect__SWIG_0(a[0].(Rectangle), a[1].(int64))
	}
	if argc == 3 {
		return Grow_rect__SWIG_1(a[0].(Rectangle), a[1].(int64), a[2].(int64))
	}
	panic("No match for overloaded function call")
}

func Translate_rect__SWIG_0(arg1 Rectangle, arg2 Dlib_vector_Sl_long_Sc_2_Sg_) (_swig_ret Rectangle) {
	var swig_r Rectangle
	_swig_i_0 := arg1.Swigcptr()
	_swig_i_1 := arg2.Swigcptr()
	swig_r = (Rectangle)(SwigcptrRectangle(C._wrap_translate_rect__SWIG_0_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))))
	return swig_r
}

func Translate_rect__SWIG_1(arg1 Rectangle, arg2 int64, arg3 int64) (_swig_ret Rectangle) {
	var swig_r Rectangle
	_swig_i_0 := arg1.Swigcptr()
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	swig_r = (Rectangle)(SwigcptrRectangle(C._wrap_translate_rect__SWIG_1_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0), C.swig_type_37(_swig_i_1), C.swig_type_38(_swig_i_2))))
	return swig_r
}

func Translate_rect(a ...interface{}) Rectangle {
	argc := len(a)
	if argc == 2 {
		return Translate_rect__SWIG_0(a[0].(Rectangle), a[1].(Dlib_vector_Sl_long_Sc_2_Sg_))
	}
	if argc == 3 {
		return Translate_rect__SWIG_1(a[0].(Rectangle), a[1].(int64), a[2].(int64))
	}
	panic("No match for overloaded function call")
}

func Resize_rect(arg1 Rectangle, arg2 uint64, arg3 uint64) (_swig_ret Rectangle) {
	var swig_r Rectangle
	_swig_i_0 := arg1.Swigcptr()
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	swig_r = (Rectangle)(SwigcptrRectangle(C._wrap_resize_rect_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0), C.swig_type_39(_swig_i_1), C.swig_type_40(_swig_i_2))))
	return swig_r
}

func Resize_rect_width(arg1 Rectangle, arg2 uint64) (_swig_ret Rectangle) {
	var swig_r Rectangle
	_swig_i_0 := arg1.Swigcptr()
	_swig_i_1 := arg2
	swig_r = (Rectangle)(SwigcptrRectangle(C._wrap_resize_rect_width_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0), C.swig_type_41(_swig_i_1))))
	return swig_r
}

func Resize_rect_height(arg1 Rectangle, arg2 uint64) (_swig_ret Rectangle) {
	var swig_r Rectangle
	_swig_i_0 := arg1.Swigcptr()
	_swig_i_1 := arg2
	swig_r = (Rectangle)(SwigcptrRectangle(C._wrap_resize_rect_height_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0), C.swig_type_42(_swig_i_1))))
	return swig_r
}

func Move_rect__SWIG_0(arg1 Rectangle, arg2 Dlib_vector_Sl_long_Sc_2_Sg_) (_swig_ret Rectangle) {
	var swig_r Rectangle
	_swig_i_0 := arg1.Swigcptr()
	_swig_i_1 := arg2.Swigcptr()
	swig_r = (Rectangle)(SwigcptrRectangle(C._wrap_move_rect__SWIG_0_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))))
	return swig_r
}

func Move_rect__SWIG_1(arg1 Rectangle, arg2 int64, arg3 int64) (_swig_ret Rectangle) {
	var swig_r Rectangle
	_swig_i_0 := arg1.Swigcptr()
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	swig_r = (Rectangle)(SwigcptrRectangle(C._wrap_move_rect__SWIG_1_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0), C.swig_type_43(_swig_i_1), C.swig_type_44(_swig_i_2))))
	return swig_r
}

func Move_rect(a ...interface{}) Rectangle {
	argc := len(a)
	if argc == 2 {
		return Move_rect__SWIG_0(a[0].(Rectangle), a[1].(Dlib_vector_Sl_long_Sc_2_Sg_))
	}
	if argc == 3 {
		return Move_rect__SWIG_1(a[0].(Rectangle), a[1].(int64), a[2].(int64))
	}
	panic("No match for overloaded function call")
}

func Set_rect_area(arg1 Rectangle, arg2 uint64) (_swig_ret Rectangle) {
	var swig_r Rectangle
	_swig_i_0 := arg1.Swigcptr()
	_swig_i_1 := arg2
	swig_r = (Rectangle)(SwigcptrRectangle(C._wrap_set_rect_area_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0), C.swig_type_45(_swig_i_1))))
	return swig_r
}

func Set_aspect_ratio(arg1 Rectangle, arg2 float64) (_swig_ret Rectangle) {
	var swig_r Rectangle
	_swig_i_0 := arg1.Swigcptr()
	_swig_i_1 := arg2
	swig_r = (Rectangle)(SwigcptrRectangle(C._wrap_set_aspect_ratio_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0), C.double(_swig_i_1))))
	return swig_r
}

type SwigcptrShapeObjects uintptr

func (p SwigcptrShapeObjects) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrShapeObjects) SwigIsShapeObjects() {
}

func NewShapeObjects(arg1 int, arg2 Rectangle) (_swig_ret ShapeObjects) {
	var swig_r ShapeObjects
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	swig_r = (ShapeObjects)(SwigcptrShapeObjects(C._wrap_new_ShapeObjects_dlibgo_63b7eba619e1d075(C.swig_intgo(_swig_i_0), C.uintptr_t(_swig_i_1))))
	return swig_r
}

func (arg1 SwigcptrShapeObjects) GetRect() (_swig_ret Rectangle) {
	var swig_r Rectangle
	_swig_i_0 := arg1
	swig_r = (Rectangle)(SwigcptrRectangle(C._wrap_ShapeObjects_GetRect_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func (arg1 SwigcptrShapeObjects) GetSize() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_ShapeObjects_GetSize_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrShapeObjects) SetPoint(arg2 int, arg3 Dlib_vector_Sl_long_Sc_2_Sg_) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3.Swigcptr()
	C._wrap_ShapeObjects_SetPoint_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1), C.uintptr_t(_swig_i_2))
}

func (arg1 SwigcptrShapeObjects) GetPoint(arg2 int) (_swig_ret Dlib_vector_Sl_long_Sc_2_Sg_) {
	var swig_r Dlib_vector_Sl_long_Sc_2_Sg_
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (Dlib_vector_Sl_long_Sc_2_Sg_)(SwigcptrDlib_vector_Sl_long_Sc_2_Sg_(C._wrap_ShapeObjects_GetPoint_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))))
	return swig_r
}

func (arg1 SwigcptrShapeObjects) X(arg2 int) (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (int64)(C._wrap_ShapeObjects_X_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1)))
	return swig_r
}

func (arg1 SwigcptrShapeObjects) Y(arg2 int) (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (int64)(C._wrap_ShapeObjects_Y_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1)))
	return swig_r
}

func DeleteShapeObjects(arg1 ShapeObjects) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_ShapeObjects_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0))
}

type ShapeObjects interface {
	Swigcptr() uintptr
	SwigIsShapeObjects()
	GetRect() (_swig_ret Rectangle)
	GetSize() (_swig_ret int)
	SetPoint(arg2 int, arg3 Dlib_vector_Sl_long_Sc_2_Sg_)
	GetPoint(arg2 int) (_swig_ret Dlib_vector_Sl_long_Sc_2_Sg_)
	X(arg2 int) (_swig_ret int64)
	Y(arg2 int) (_swig_ret int64)
}

func LoadShapePredictor(arg1 string) (_swig_ret Dlib_shape_predictor) {
	var swig_r Dlib_shape_predictor
	_swig_i_0 := arg1
	swig_r = (Dlib_shape_predictor)(SwigcptrDlib_shape_predictor(C._wrap_LoadShapePredictor_dlibgo_63b7eba619e1d075(*(*C.swig_type_48)(unsafe.Pointer(&_swig_i_0)))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	return swig_r
}

func UseShapePredictor(arg1 Dlib_shape_predictor, arg2 string, arg3 Rectangle) (_swig_ret ShapeObjects) {
	var swig_r ShapeObjects
	_swig_i_0 := arg1.Swigcptr()
	_swig_i_1 := arg2
	_swig_i_2 := arg3.Swigcptr()
	swig_r = (ShapeObjects)(SwigcptrShapeObjects(C._wrap_UseShapePredictor_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0), *(*C.swig_type_49)(unsafe.Pointer(&_swig_i_1)), C.uintptr_t(_swig_i_2))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func UseShapePredictorByte(arg1 Dlib_shape_predictor, arg2 *byte, arg3 int64, arg4 Rectangle) (_swig_ret ShapeObjects) {
	var swig_r ShapeObjects
	_swig_i_0 := arg1.Swigcptr()
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4.Swigcptr()
	swig_r = (ShapeObjects)(SwigcptrShapeObjects(C._wrap_UseShapePredictorByte_dlibgo_63b7eba619e1d075(C.uintptr_t(_swig_i_0), C.swig_voidp(_swig_i_1), C.swig_type_50(_swig_i_2), C.uintptr_t(_swig_i_3))))
	return swig_r
}

func TestMemoryFile(arg1 *byte, arg2 int64) (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (int64)(C._wrap_TestMemoryFile_dlibgo_63b7eba619e1d075(C.swig_voidp(_swig_i_0), C.swig_type_52(_swig_i_1)))
	return swig_r
}

func DetectObjectsRect(arg1 string, arg2 string, arg3 Rectangle) (_swig_ret Dlib_full_object_detection) {
	var swig_r Dlib_full_object_detection
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3.Swigcptr()
	swig_r = (Dlib_full_object_detection)(SwigcptrDlib_full_object_detection(C._wrap_DetectObjectsRect_dlibgo_63b7eba619e1d075(*(*C.swig_type_53)(unsafe.Pointer(&_swig_i_0)), *(*C.swig_type_54)(unsafe.Pointer(&_swig_i_1)), C.uintptr_t(_swig_i_2))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}


type SwigcptrDlib_vector_Sl_double_Sc_2_Sg_ uintptr
type Dlib_vector_Sl_double_Sc_2_Sg_ interface {
	Swigcptr() uintptr;
}
func (p SwigcptrDlib_vector_Sl_double_Sc_2_Sg_) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrDlib_shape_predictor uintptr
type Dlib_shape_predictor interface {
	Swigcptr() uintptr;
}
func (p SwigcptrDlib_shape_predictor) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrStd_vector_Sl_dlib_rectangle_Sg_ uintptr
type Std_vector_Sl_dlib_rectangle_Sg_ interface {
	Swigcptr() uintptr;
}
func (p SwigcptrStd_vector_Sl_dlib_rectangle_Sg_) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrDlib_vector_Sl_long_Sc_2_Sg_ uintptr
type Dlib_vector_Sl_long_Sc_2_Sg_ interface {
	Swigcptr() uintptr;
}
func (p SwigcptrDlib_vector_Sl_long_Sc_2_Sg_) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrStd_ostream uintptr
type Std_ostream interface {
	Swigcptr() uintptr;
}
func (p SwigcptrStd_ostream) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrStd_istream uintptr
type Std_istream interface {
	Swigcptr() uintptr;
}
func (p SwigcptrStd_istream) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrDlib_full_object_detection uintptr
type Dlib_full_object_detection interface {
	Swigcptr() uintptr;
}
func (p SwigcptrDlib_full_object_detection) Swigcptr() uintptr {
	return uintptr(p)
}