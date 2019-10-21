#ifdef _MSC_VER
#define _CRT_SECURE_NO_WARNINGS 1
#endif

#include "basiclib.h"

#include <stdarg.h>
#include <stdlib.h>
#include <string.h>

static void ops(const char *format, ...) {
    va_list ap;
    va_start(ap, format);
    fputc('\n', stderr);
    vfprintf(stderr, format, ap);
    fputc('\n', stderr);
    va_end(ap);
    abort();
}

static void *realloc_mem(void *ptr, size_t size) {
    void *result = realloc(ptr, size);
    if (!result) {
        ops("Out of memory");
    }
    return result;
}

static void dim(arr *a, size_t elem_size, int argcnt, va_list list) {
    va_list ap;
    size_t total, multiplier;
    int i;
    size_t *p;

    ap = list;
    multiplier = 1;
    for (i = 0; i < argcnt; i++) {
        size_t size = va_arg(ap, size_t);
        if (size < 1) {
            ops("DIMension must be greater than zero");
        }
        multiplier *= size +1;
    }
    total = argcnt * sizeof(size_t) + multiplier * elem_size;

    p = *a = realloc_mem(*a, total);
    memset(p, 0, total);

    ap = list;
    for (i = 0; i < argcnt; i++) {
        size_t size = va_arg(ap, size_t);
        p[i] = size +1;
    }
}

void dim_num(arr *a, int argcnt, ...) {
    va_list ap;
    va_start(ap, argcnt);
    dim(a, sizeof(num), argcnt, ap);
    va_end(ap);
}

void dim_str(arr *a, int argcnt, ...) {
    va_list ap;
    va_start(ap, argcnt);
    dim(a, sizeof(str), argcnt, ap);
    va_end(ap);
}


static void *element_in_array(arr *a, size_t elem_size, int argcnt, va_list ap) {
    size_t *p = *a;
    size_t offset, multiplier;
    int i;

    offset = 0;
    multiplier = 1;
    for (i = 0; i < argcnt; i++, p++) {
        size_t pos = va_arg(ap, size_t);
        if (pos >= *p) {
            ops("Index out of bounds for dimension %d: %u. It should be a value from 0 up to %u", i+1, pos, *p -1);
        }
        offset += multiplier * pos;
        multiplier *= *p;
    }
    offset *= elem_size;
    return ((char*)p)+offset;
}

num *num_in_array(arr *a, int argcnt, ...) {
    num *result;
    va_list ap;

    if (*a == NULL) { /* Did the programmer forget to DIMension this array? */
        dim_num(a, 1, 10); /* No problem. We can DIM it! It's a BASIC thing */
    }

    va_start(ap, argcnt);
    result = (num*)element_in_array(a, sizeof(num), argcnt, ap);
    va_end(ap);
    return result;
}

str *str_in_array(arr *a, int argcnt, ...) {
    str *result;
    va_list ap;

    if (*a == NULL) { /* Did the programmer forget to DIMension this array? */
        dim_str(a, 1, 10); /* No problem. We can DIM it! It's a BASIC thing */
    }

    va_start(ap, argcnt);
    result = (str*)element_in_array(a, sizeof(str), argcnt, ap);
    va_end(ap);
    return result;
}

void let_num(num *dst, num src) {
    *dst = src;
}

void let_str(str *dst, str src) {
    if (!src) {
        free(*dst);
        *dst = NULL;
        return;
    }
    size_t size = strlen(src)+1;
    *dst = realloc_mem(*dst, size);
    strcpy(*dst, src);
}

static char input_buffer[4 * 1024], *input_ptr;

void input_to_buffer() {
    char *p;

    print_str("? ");
    if (!fgets(input_buffer, sizeof input_buffer, stdin)) {
        ops("INPUT failed");
    }

    p = strchr(input_buffer, '\n');
    if (p) {
        *p = '\0';
    }
    input_ptr = input_buffer;
}

void input_num(num *dst) {
    char *p;
    if (!input_ptr || *input_ptr == '\0') {
        *dst = 0;
    }

    *dst = (num)atof(input_ptr);
    p = strchr(input_ptr, ',');
    if (p) {
        input_ptr = p +1;
    } else {
        input_ptr = NULL;
    }
}

void input_str(str *dst) {
    char *p;
    size_t size;

    if (!input_ptr || *input_ptr == '\0') {
        free(*dst);
        *dst = NULL;
    }

    p = strchr(input_ptr, ',');
    if (p) {
        size = p - input_ptr;
        *dst = realloc_mem(*dst, size +1);
        memcpy(*dst, input_ptr, size);
        (*dst)[size] = '\0';
        input_ptr = p+1;
    } else {
        size = strlen(input_ptr);
        *dst = realloc_mem(*dst, size +1);
        strcpy(*dst, input_ptr);
        input_ptr = NULL; /* input is depleted */
    }
}

static int current_column;

void print_char(char c) {
    switch (c) {
    case '\n':
        current_column = 0;
        break;
    case '\t':
    {
        int pos = ((current_column / 16) + 1) * 16;
        while (current_column < pos) {
            print_char(' ');
        }
    }
    return;
    }
    if (c == '\n') {
        current_column = 0;
    } else {
        current_column++;
    }
    putchar(c);
}

void print_num(num val) {
    char buffer[64], *p = buffer;
    if (val == (int)val) {
        sprintf(buffer, " %d ", (int)val);
    } else {
        sprintf(buffer, " %f ", val);
    }
    while (*p) {
        print_char(*p++);
    }
}

void print_str(str val) {
    if (!val) {
        return;
    }
    for (; *val; val++) {
        print_char(*val);
    }
}

void TAB_void(num val) {
    int column = (int)val;
    if (column <= current_column) {
        print_char('\n');
    }
    while (current_column < column) {
        print_char(' ');
    }
}

#define ADDRESS_STACK_SIZE 64
static int address_stack[ADDRESS_STACK_SIZE], *address_stack_ptr;

void push_address(int address) {
    if (!address_stack_ptr) {
        address_stack_ptr = address_stack;
    } else {
        if (address_stack_ptr - address_stack >= ADDRESS_STACK_SIZE) {
            ops("Too many nested GOSUBs");
        }
        address_stack_ptr++;
    }
    *address_stack_ptr = address;
}

void pop_address(int *address) {
    if (!address_stack_ptr) {
        ops("RETURN without GOSUB");
    }
    *address = *address_stack_ptr;
    if (address_stack_ptr == address_stack) {
        address_stack_ptr = NULL;
    } else {
        address_stack_ptr--;
    }
}

static size_t data_area_for_num_index;

void read_num(num *val) {
    extern const size_t data_area_for_num_cnt;
    extern const num data_area_for_num[];

    if (data_area_for_num_index >= data_area_for_num_cnt) {
        ops("READ number past DATA");
    }
    *val = data_area_for_num[data_area_for_num_index++];
}

static size_t data_area_for_str_index;

void read_str(str *val) {
    extern const size_t data_area_for_str_cnt;
    extern const str data_area_for_str[];

    size_t size;
    str src;
    str dst;

    if (data_area_for_str_index >= data_area_for_str_cnt) {
        ops("READ string past DATA");
    }
    src = data_area_for_str[data_area_for_str_index++];
    size = strlen(src);
    dst = *val = realloc_mem(*val, size +1);
    memcpy(dst, src, size);
    dst[size] = '\0';
}

void restore(void) {
    data_area_for_num_index = 0;
    data_area_for_str_index = 0;
}

num ABS_num(num val) {
    return (num)fabs(val);
}

num ASC_num(str val) {
    if (!val) {
        ops("Trying to perform ASC function on an empty string");
    }
    return *val;
}

num ATN_num(num val) {
    return (num)atan(val);
}

str CHR_str(str *dst, num val) {
    str s = *dst = realloc_mem(*dst, 2);
    s[0] = (unsigned char)val;
    s[1] = '\0';
    return *dst;
}

num COS_num(num val) {
    return (num)cos(val);
}

num EXP_num(num val) {
    return (num)exp(val);
}

num INT_num(num val) {
    return (num)(int)val;
}

num LEN_num(str val) {
    if (!val) {
        return 0;
    }
    return (num)strlen(val);
}

num LOG_num(num val) {
    return (num)log(val);
}

str RIGHT_str(str *dst, str s, num length_num) {
    size_t size;
    size_t length = (size_t)length_num;

    if (!s) {
        free(*dst);
        return *dst = NULL;
    }

    size = strlen(s);
    if (size <= length) {
        length = size;
    }

    if (length == 0) {
        free(*dst);
        return *dst = NULL;
    }

    *dst = realloc_mem(*dst, length +1);
    memcpy(*dst, s +size -length, length);
    (*dst)[length] = '\0';
    return *dst;
}

str LEFT_str(str *dst, str s, num length_num) {
    size_t size;
    size_t length = (size_t)length_num;

    if (!s) {
        free(*dst);
        return *dst = NULL;
    }

    size = strlen(s);
    if (size < length) {
        length = size;
    }

    if (length == 0) {
        free(*dst);
        return *dst = NULL;
    }

    *dst = realloc_mem(*dst, length +1);
    memcpy(*dst, s, length);
    (*dst)[length] = '\0';
    return *dst;
}

str MID_str(str *dst, str s, num start_num, num length_num) {
    size_t size;
    size_t start = (size_t)start_num;
    size_t length = (size_t)length_num;

    if (!s) {
        free(*dst);
        return *dst = NULL;
    }

    start--;
    size = strlen(s);
    if (size <= start) {
        free(*dst);
        return *dst = NULL;
    }
    return LEFT_str(dst, s + start, (num)length);
}

num RND_num(num val) {
    val = val;
    double dummy;

    /* see https://stackoverflow.com/questions/13408990/how-to-generate-random-float-number-in-c#13409133 */
    return (num)modf(((num)rand()/(num)(RAND_MAX)), &dummy);
}

num SGN_num(num val) {
    return (num)signbit(val);
}

num SIN_num(num val) {
    return (num)sin(val);
}

num SQR_num(num val) {
    return (num)sqrt(val);
}

str STR_str(str *dst, num val) {
    char buffer[64];
    size_t size = sprintf(buffer, "%f", val);
    *dst = realloc_mem(*dst, size +1);
    memcpy(*dst, buffer, size);
    (*dst)[size] = '\0';
    return *dst;
}

num TAN_num(num val) {
    return (num)tan(val);
}

num VAL_num(str val) {
    return (num)atof(val);
}

str concat_str(str *dst, int argcnt, ...) {
    va_list ap;
    size_t size;
    int i;
    str p;

    va_start(ap, argcnt);
    size = 0;
    for (i = 0; i < argcnt; i++) {
        str s = va_arg(ap, str);
        if (s) {
            size += strlen(s);
        }
    }
    p = *dst = realloc_mem(*dst, size +1);

    va_start(ap, argcnt);
    for (i = 0; i < argcnt; i++) {
        str s = va_arg(ap, str);
        if (s) {
            size = strlen(s);
            memcpy(p, s, size);
            p += size;
        }
    }
    *p = '\0';
    va_end(ap);

    return *dst;
}

int compare_str(str lhs, str rhs) {
    if (!lhs) {
        lhs = "";
    }
    if (!rhs) {
        rhs = "";
    }
    return strcmp(lhs, rhs);
}
