#ifndef BASICLIB_H
#define BASICLIB_H

#include <stdlib.h>
#include <stdio.h>
#include <math.h>
#include <time.h>

typedef double num;
typedef unsigned char* str;
typedef size_t* arr;

void dim_num(arr *a, int argcnt, ...);
void dim_str(arr *a, int argcnt, ...);

num *num_in_array(arr *a, int argcnt, ...);
str *str_in_array(arr *a, int argcnt, ...);

#define let_num(dst,src) (*(dst) = (src))
void let_str(str *dst, str src);

void input_to_buffer();
void input_num(num *dst);
void input_str(str *dst);

void print_num(num val);
void print_str(str val);
void print_char(char c);

void push_address(int address);
void pop_address(int *address);

void read_num(num *val);
void read_str(str *val);
void restore(void);

#define ABS_num(val) ((num)fabs(val))
num ASC_num(str val);
#define ATN_num(val) ((num)atan(val))
str CHR_str(str *dst, num val);
#define COS_num(val) ((num)cos(val))
#define EXP_num(val) ((num)exp(val))
#define INT_num(val) ((num)(int)(val))
str LEFT_str(str *dst, str s, num length);
num LEN_num(str val);
#define LOG_num(val) ((num)log(val))
str MID_str(str *dst, str s, num start, num length);
str RIGHT_str(str *dst, str s, num length);
num RND_num(num val);
#define SGN_num(val) ((num)signbit(val))
#define SIN_num(val) ((num)sin(val))
#define SQR_num(val) ((num)sqrt(val))
str STR_str(str *dst, num val);
void TAB_void(num val);
#define TAN_num(val) ((num)tan(val))
#define VAL_num(val) ((num)atof(val))

str concat_str(str *dst, int argcnt, ...);
int compare_str(str lhs, str rhs);

#endif /* BASICLIB_H */
