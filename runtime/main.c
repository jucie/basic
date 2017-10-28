// main.c : Defines the entry point for the console application.
#include "common.h"

extern const Command codeTab[];

void applyFixupTable(FILE *in, unsigned f(unsigned)) {
	int cnt;
	if (fread(&cnt, 1, sizeof cnt, in) < sizeof cnt) {
		fail("error reading fixup counter");
	}
	int i;
	for (i = 0; i < cnt; i++){
		long offset;
		if (fread(&offset, 1, sizeof offset, in) < sizeof offset) {
			fail("error reading fixup");
		}
		unsigned *p = (unsigned *)(codeBottom + offset);
		if (p < (unsigned*)codeBottom || p >= (unsigned*)codeTop){
			fail("fixup out of range");
		}
		*p = f(*p);
	}
}

unsigned fixFuncAddress(unsigned val){
	return (unsigned) codeTab[-val];
}

unsigned fixOffset(unsigned val){
	return val + (unsigned)codeBottom;
}

void applyFixups(FILE *in) {
	applyFixupTable(in, fixFuncAddress);
	applyFixupTable(in, fixOffset);
}

void loadProgram(const char *exePath){
	FILE *in = fopen(exePath, "rb");
	if (!in) {
		fail("Couldn't open program file");
	}
	struct {
		long offset;
		char name[4];
	} tail;
	fseek(in, -sizeof tail, SEEK_END);
	long end = ftell(in);
	fread(&tail, sizeof tail, 1, in);
	if (memcmp(tail.name, "TAIL", sizeof tail)) {
		fail("Missing tail");
	}
	fseek(in, tail.offset, SEEK_SET);
	size_t version;
	if (fread(&version, sizeof version, 1, in) < 1) {
		fail("error reading version");
	}
	if (version != 1) {
		fail("unsupported version");
	}
	size_t dataSize;
	if (fread(&dataSize, sizeof dataSize, 1, in) < 1) {
		fail("error reading size of data block");
	}
	size_t stackSize;
	if (fread(&stackSize, sizeof stackSize, 1, in) < 1) {
		fail("error reading size of stack");
	}
	size_t codeSize;
	if (fread(&codeSize, sizeof codeSize, 1, in) < 1) {
		fail("error reading fixup offset");
	}

	size_t size = codeSize + dataSize + stackSize;
	pc = (byte*)allocMem(size, 1);
	codeBottom = pc;
	codeTop = pc + codeSize;
	stackBottom = (unsigned*)pc + size;
	stackTop = (unsigned*)pc + size + dataSize;
	sp = stackBottom;
	
	if (fread(pc, 1, codeSize, in) < codeSize) {
		fail("error reading program");
	}

	applyFixups(in);
	fclose(in);
}

int main(int argc, const char **argv){
	loadProgram(argc > 1? argv[1] : argv[0]);
	for (;;) {
		Command cmd = GET(Command);
		cmd();
	}
    return 0;
}

