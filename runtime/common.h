#ifndef COMMON_H
#define COMMON_H

#include <stdio.h>
#include <stdlib.h>

typedef unsigned char byte;
typedef void(*Command)(void);

extern unsigned *sp, *stackBottom, *stackTop;
extern byte *pc, *codeBottom, *codeTop;
#define GET(type) *((type*)pc); pc += sizeof(type);

void fail(const char *msg);
void *allocMem(size_t num, size_t size);

#endif // COMMON_H