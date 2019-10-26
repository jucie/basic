# basic
A compiler for BASIC programming language.

This tool understands what is in David H. Ahl's own words "the gold standard of microcomputer BASICs: MITS Altair 8K BASIC, Rev. 4.0 (ix)."
It has been tested against every program from Ahl's bestseller "BASIC Computer Games". More information can be found here:

https://www.atariarchives.org/basicgames

Source files from the book are available in bcg.zip for your convenience.
Please keep in mind that those programs are extremely simple if compared to today standards. They were written to be used with TTY terminals.

By default this BASIC compiler reads from standard input and writes to standard output.
That behaviour may be modified by specifying paths for input and output files, like this:

basic -in civilwar.bas -out civilwar.c

For generating the final executable you may use your C compiler of choice. Any compiler supporting C89 or later will do.
Be sure to have the basiclib files from "lib" available in the same folder as the main source file. 
Then, invoke the C compiler like that:

gcc civilwar.c basiclib.c

The generated code has been tested on half a dozen distinct C compilers and works fine.

I would like to thank:

.David Ahl for the awesome book

.Lyle Kopnicky for inspiration (he wrote a BASIC interpreter)

.Atari Archives for the already typed program files

