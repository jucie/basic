# basic
A compiler for BASIC programming language.

This tool has been tested against every program from the 70's bestseller "BASIC Computer Games" by David H. Ahl. More information can be found here:

https://www.atariarchives.org/basicgames

The source files for the programs in the book are available in the "pool" folder for your delight.
Please keep in mind that those programs are very antique and extremely simple for today standards.

To be able to compile such antique source code, this compiler understands what is in David H. Ahl's own words "the gold standard of microcomputer BASICs: MITS Altair 8K BASIC, Rev. 4.0 (ix)."

By default this BASIC compiler reads from standard input and writes to standard output.
That behaviour may be modified by specifying paths for input and output files. E.g:

basic -in civilwar.bas -out civilwar.c

For generating the final executable you may use your C compiler of choice. Any compiler from C89 and later will do.
Be sure to have the basiclib files from "lib" available in the same folder as the main source file. 
Then, invoke the C compiler like that:

gcc civilwar.c basiclib.c

The generated code has been tested on half a dozen distinct C compilers and works fine.
