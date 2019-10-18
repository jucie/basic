#include "basiclib.h"

static num A_num;
static str A_str;
static num C_num;
static num L_num;
static num M_num;
static num N_num;
static num Q_num;
static num R_num;
static num X_num;
static num Y_num;
static num Z_num;


static str temp_str[];

int main() {
    int target = 0;
    for(;;) {
        switch (target) {
        case 0:
            /* line 1 */
            TAB_void(33.0000f);
            print_str("DIAMOND");
            print_char('\n');
            /* line 2 */
            TAB_void(15.0000f);
            print_str("CREATIVE COMPUTING  MORRISTOWN, NEW JERSEY");
            print_char('\n');
            /* line 3 */
            print_char('\n');
            print_char('\n');
            print_char('\n');
            /* line 4 */
            print_str("FOR A PRETTY DIAMOND PATTERN,");
            print_char('\n');
        /* line 5 */
        case -1:
            print_str("TYPE IN AN ODD NUMBER BETWEEN 5 AND 21? ");
            input_to_buffer();
            if (!input_num(&R_num)) {
                target = -1;
                break;
            }
            print_char('\n');
            /* line 6 */
            let_num(&Q_num,INT_num(60.0000f/R_num));
            let_str(&A_str,"CC");
            /* line 8 */
            for_num(&L_num,1.0000f,Q_num,1.0f,-2);
        case -2:
            /* line 10 */
            let_num(&X_num,1.0000f);
            let_num(&Y_num,R_num);
            let_num(&Z_num,2.0000f);
        case 20: /* line 20 */
            for_num(&N_num,X_num,Y_num,Z_num,-3);
        case -3:
            /* line 25 */
            TAB_void((R_num-N_num)/2.0000f);
            /* line 28 */
            for_num(&M_num,1.0000f,Q_num,1.0f,-4);
        case -4:
            /* line 29 */
            let_num(&C_num,1.0000f);
            /* line 30 */
            for_num(&A_num,1.0000f,N_num,1.0f,-5);
        case -5:
            /* line 32 */
            if (!(C_num>LEN_num(A_str))) {
                target = -6;
                break;
            }
            print_str("!");
            target = 50;
            break;
        case -6:
            /* line 34 */
            print_str(MID_str(&temp_str[0],A_str,C_num,1.0000f));
            /* line 36 */
            let_num(&C_num,C_num+1.0000f);
        case 50: /* line 50 */
            if (next(&target)) break; /* NEXT A */
            /* line 53 */
            if (M_num==Q_num) {
                target = 60;
                break;
            }
            /* line 55 */
            TAB_void(R_num*M_num+(R_num-N_num)/2.0000f);
            /* line 56 */
            if (next(&target)) break; /* NEXT M */
        case 60: /* line 60 */
            print_char('\n');
            /* line 70 */
            if (next(&target)) break; /* NEXT N */
            /* line 83 */
            if (X_num!=1.0000f) {
                target = 95;
                break;
            }
            /* line 85 */
            let_num(&X_num,R_num-2.0000f);
            let_num(&Y_num,1.0000f);
            let_num(&Z_num,-2.0000f);
            /* line 90 */
            target = 20;
            break;
        case 95: /* line 95 */
            if (next(&target)) break; /* NEXT L */
            /* line 99 */
            exit(0);
        case -7:
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

static str temp_str[1];
