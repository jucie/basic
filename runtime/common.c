
#include "common.h"

unsigned *sp, *stackBottom, *stackTop;
byte *pc, *codeBottom, *codeTop;

void fail(const char *msg) {
	fprintf(stderr, "\n%s\n", msg);
	exit(1);
}

void *allocMem(size_t num, size_t size) {
	void *result = calloc(num, size);
	if (!result) {
		fail("Memory allocation failure");
	}
	return result;
}
