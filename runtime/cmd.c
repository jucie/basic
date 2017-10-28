
#include "common.h"

void nop() { // CMD
	// do nothing
}

void dup() { // CMD
	if (sp <= stackTop){
		fail("Stack overflow");
	}
	unsigned val = *sp--;
	*sp = val;
}

void drop() { // CMD
	if (sp >= stackBottom){
		fail("Stack underflow");
	}
	sp++;
}

void swap() { // CMD
	unsigned val = sp[0];
	sp[0] = sp[1];
	sp[1] = val;
}

void push() { // CMD
}

const Command codeTab[] = {
	nop,
	dup,
	drop,
	swap,
	push,
};
