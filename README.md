# basic
A compiler for BASIC programming language.

This tool understands what is in David H. Ahl's own words "the gold standard of microcomputer BASICs: MITS Altair 8K BASIC, Rev. 4.0 (ix)."
It has been tested against every program from Ahl's bestseller "BASIC Computer Games". More information can be found here:

https://www.atariarchives.org/basicgames

Please keep in mind that those programs are extremely simple if compared to today standards. They were written to be used with TTY terminals.

To clone and generate the compiler:

go get -u github.com/jucie/basic

You can download the entire set of games as a tarball or a ZIP archive:

http://vintage-basic.net/downloads/bcg.tar.gz

http://vintage-basic.net/downloads/bcg.zip

To generate the C code for a game run the BASIC compiler like that:

basic game.bas game.c

The lib folder has some support files to be compiled with the generated C code. Be sure to have then in the folder as the game C file.
Then use youc C compiler of choice to generate the executable binary:

gcc game.c basiclib.c

The programs have been tested on half a dozen distinct C compilers in both Windows and Linux.

I would like to thank:

.David Ahl for the awesome work collecting all those programs and promoting BASIC

.Lyle Kopnicky for inspiration (he wrote a BASIC interpreter. Check it out at http://vintage-basic.net )

.Atari Archives for the already typed program files

