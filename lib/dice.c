#include "basiclib.h"

static num A_num;
static num B_num;
static arr F_num_array_var;
static num Q_num;
static num R_num;
static num S_num;
static num V_num;
static num X_num;
static str Z_str;

static num *F_num_array(num index0) {
    return num_in_array(&F_num_array_var,1,(size_t)index0);
}

static str temp_str[];

int main() {
    int target = 0;
    for(;;) {
        switch (target) {
        case 0:
            /* line 2 */
            TAB_void(34.0000f);
            print_str("DICE");
            print_char('\n');
            /* line 4 */
            TAB_void(15.0000f);
            print_str("CREATIVE COMPUTING  MORRISTOWN, NEW JERSEY");
            print_char('\n');
            /* line 6 */
            print_char('\n');
            print_char('\n');
            print_char('\n');
            /* line 10 */
            dim_num(&F_num_array_var,1,(size_t)12.0000f);
            /* line 20 */
            /*  DANNY FREIDUS*/
            /* line 30 */
            print_str("THIS PROGRAM SIMULATES THE ROLLING OF A");
            print_char('\n');
            /* line 40 */
            print_str("PAIR OF DICE.");
            print_char('\n');
            /* line 50 */
            print_str("YOU ENTER THE NUMBER OF TIMES YOU WANT THE COMPUTER TO");
            print_char('\n');
            /* line 60 */
            print_str("\'ROLL\' THE DICE.  WATCH OUT, VERY LARGE NUMBERS TAKE");
            print_char('\n');
            /* line 70 */
            print_str("A LONG TIME.  IN PARTICULAR, NUMBERS OVER 5000.");
            print_char('\n');
        case 80: /* line 80 */
            for_num(&Q_num,1.0000f,12.0000f,1.0f,-1);
        case -1:
            /* line 90 */
            let_num(F_num_array(Q_num),0.0000f);
            /* line 100 */
            if (next(&target)) break; /* NEXT Q */
            /* line 110 */
            print_char('\n');
            print_str("HOW MANY ROLLS");
        /* line 120 */
        case -2:
            input_to_buffer();
            if (!input_num(&X_num)) {
                target = -2;
                break;
            }
            /* line 130 */
            for_num(&S_num,1.0000f,X_num,1.0f,-3);
        case -3:
            /* line 140 */
            let_num(&A_num,INT_num(6.0000f*RND_num(1.0000f)+1.0000f));
            /* line 150 */
            let_num(&B_num,INT_num(6.0000f*RND_num(1.0000f)+1.0000f));
            /* line 160 */
            let_num(&R_num,A_num+B_num);
            /* line 170 */
            let_num(F_num_array(R_num),*F_num_array(R_num)+1.0000f);
            /* line 180 */
            if (next(&target)) break; /* NEXT S */
            /* line 185 */
            print_char('\n');
            /* line 190 */
            print_str("TOTAL SPOTS");
            print_char('	');
            print_str("NUMBER OF TIMES");
            print_char('\n');
            /* line 200 */
            for_num(&V_num,2.0000f,12.0000f,1.0f,-4);
        case -4:
            /* line 210 */
            print_num(V_num);
            print_char('	');
            print_num(*F_num_array(V_num));
            print_char('\n');
            /* line 220 */
            if (next(&target)) break; /* NEXT V */
            /* line 221 */
            print_char('\n');
            /* line 222 */
            print_char('\n');
            print_str("TRY AGAIN");
        /* line 223 */
        case -5:
            input_to_buffer();
            if (!input_str(&Z_str)) {
                target = -5;
                break;
            }
            /* line 224 */
            if (strcmp(Z_str,"YES")==0) {
                target = 80;
                break;
            }
            /* line 240 */
            exit(0);
        case -6:
            exit(0);
        default:
            fprintf(stderr, "Undefined target line %d", target);
            abort();
        }
    }
} /* main */

const size_t data_area_for_str_cnt=0;
const str data_area_for_str[0]= {
};

const size_t data_area_for_num_cnt=0;
const num data_area_for_num[0]= {
};

static str temp_str[0];
