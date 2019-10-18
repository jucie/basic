#ifndef BASICLIB_H
#define BASICLIB_H

#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <math.h>

typedef float num;
typedef unsigned char* str;
typedef size_t* arr;

void dim_num(arr *a, int argcnt, ...);
void dim_str(arr *a, int argcnt, ...);

num *num_in_array(arr *a, int argcnt, ...);
str *str_in_array(arr *a, int argcnt, ...);

void let_num(num *dst, num src);
void let_str(str *dst, str src);

void for_num(num *var, num start, num end, num step, int target);
int next(int *target);

void input_to_buffer();
int input_num(num *dst);
int input_str(str *dst);

void print_num(num val);
void print_str(str val);
void print_char(char c);

void push_address(int address);
void pop_address(int *address);

void read_num(num *val);
void read_str(str *val);
void restore(void);

num ABS_num(num val);
num ASC_num(str val);
num ATN_num(num val);
str CHR_str(str *dst, num val);
num COS_num(num val);
num EXP_num(num val);
num INT_num(num val);
str LEFT_str(str *dst, str s, num length);
num LEN_num(str val);
num LOG_num(num val);
str MID_str(str *dst, str s, num start, num length);
str RIGHT_str(str *dst, str s, num length);
num RND_num(num val);
num SGN_num(num val);
num SIN_num(num val);
num SQR_num(num val);
str STR_str(str *dst, num val);
void TAB_void(num val);
num TAN_num(num val);
num VAL_num(str val);

#endif /* BASICLIB_H */
