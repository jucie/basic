# basic
A compiler for BASIC programming language. It generates C code that you submit to the C compiler for your particular platform.

This tool understands what is in David H. Ahl's own words "the gold standard of microcomputer BASICs: MITS Altair 8K BASIC, Rev. 4.0 (ix)."
It has been tested against the programs from Ahl's bestseller "BASIC Computer Games". More information can be found here:

https://www.atariarchives.org/basicgames

You can download the entire set of games as a tarball or a ZIP archive:

http://vintage-basic.net/downloads/bcg.tar.gz

http://vintage-basic.net/downloads/bcg.zip

Please keep in mind that those programs are extremely simple if compared to today standards. They were written to be used with TTY terminals. There are a couple programs in that compressed folders that were written in newer BASIC dialects. Those programs won't work.

To clone and generate this compiler:

go get -u github.com/jucie/basic

To generate the C code for a game run the command like this:

basic game.bas game.c

The lib folder has some support files to be compiled with the generated C code. Be sure to keep them in the same folder as the game C file. Then use youc C compiler of choice to generate the executable binary:

gcc game.c basiclib.c

The programs have been tested on half a dozen distinct C compilers in both Windows and Linux.

I would like to thank:

.David Ahl for the awesome work collecting all those programs and promoting BASIC

.Lyle Kopnicky for inspiration (he wrote a BASIC interpreter. Check it out at http://vintage-basic.net )

.Atari Archives for the already typed program files

