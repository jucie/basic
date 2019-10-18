#include "basiclib.h"

static num fn_A(num Z_num);

static num L_num;
static num X_num;
static num Y1_num;
static num Y_num;
static num Z_num;


static str temp_str[];

int main() {
    int target = 0;
    for(;;) {
        switch (target) {
        case 0:
            /* line 1 */
            TAB_void(32.0000f);
            print_str("3D PLOT");
            print_char('\n');
            /* line 2 */
            TAB_void(15.0000f);
            print_str("CREATIVE COMPUTING  MORRISTOWN, NEW JERSEY");
            print_char('\n');
            /* line 3 */
            print_char('\n');
            print_char('\n');
            print_char('\n');
            /* line 5 */
            /* line 100 */
            print_char('\n');
            /* line 110 */
            for_num(&X_num,-30.0000f,30.0000f,1.5000f,-1);
        case -1:
            /* line 120 */
            let_num(&L_num,0.0000f);
            /* line 130 */
            let_num(&Y1_num,5.0000f*INT_num(SQR_num(900.0000f-X_num*X_num)/5.0000f));
            /* line 140 */
            for_num(&Y_num,Y1_num,-Y1_num,-5.0000f,-2);
        case -2:
            /* line 150 */
            let_num(&Z_num,INT_num(25.0000f+fn_A(SQR_num(X_num*X_num+Y_num*Y_num))-0.7000f*Y_num));
            /* line 160 */
            if (Z_num<=L_num) {
                target = 190;
                break;
            }
            /* line 170 */
            let_num(&L_num,Z_num);
            /* line 180 */
            TAB_void(Z_num);
            print_str("*");
        case 190: /* line 190 */
            if (next(&target)) break; /* NEXT Y */
            /* line 200 */
            print_char('\n');
            /* line 210 */
            if (next(&target)) break; /* NEXT X */
            /* line 300 */
            exit(0);
        case -3:
            exit(0);
        default:
            fprintf(stderr, "Undefined target line %d", target);
            abort();
        }
    }
} /* main */

static num fn_A(num Z_num) {
    Z_num=Z_num;
    return 30.0000f*EXP_num(-Z_num*Z_num/100.0000f);
}

const size_t data_area_for_str_cnt=0;
const str data_area_for_str[0]= {
};

const size_t data_area_for_num_cnt=0;
const num data_area_for_num[0]= {
};

static str temp_str[0];
