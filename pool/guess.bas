1 PRINT TAB(33);"GUESS"
2 PRINT TAB(15);"CREATIVE COMPUTING  MORRISTOWN, NEW JERSEY"
3 PRINT:PRINT:PRINT
4 PRINT "THIS IS A NUMBER GUESSING GAME. I'LL THINK"
5 PRINT "OF A NUMBER BETWEEN 1 AND ANY LIMIT YOU WANT."
6 PRINT "THEN YOU HAVE TO GUESS WHAT IT IS."
7 PRINT
8 PRINT "WHAT LIMIT DO YOU WANT";
9 INPUT L
10 PRINT
11 L1=INT(LOG(L)/LOG(2))+1
12 PRINT "I'M THINKING OF A NUMBER BETWEEN 1 AND";L
13 G=1
14 PRINT "NOW YOU TRY TO GUESS WHAT IT IS."
15 M=INT(L*RND(1)+1)
20 INPUT N
21 IF N>0 THEN 25
22 GOSUB 70
23 GOTO 1
25 IF N=M THEN 50
30 G=G+1
31 IF N>M THEN 40
32 PRINT "TOO LOW. TRY A BIGGER ANSWER."
33 GOTO 20
40 PRINT "TOO HIGH. TRY A SMALLER ANSWER."
42 GOTO 20
50 PRINT "THAT'S IT! YOU GOT IT IN";G;"TRIES."
52 IF G<L1 THEN 58
54 IF G=L1 THEN 60
56 PRINT "YOU SHOULD HAVE BEEN ABLE TO GET IT IN ONLY";L1
57 GOTO 65
58 PRINT "VERY ";
60 PRINT "GOOD."
65 GOSUB 70
66 GOTO 12
70 FOR H=1 TO 5
71 PRINT
72 NEXT H
73 RETURN
99 END
