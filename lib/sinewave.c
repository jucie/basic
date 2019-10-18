#include "basiclib.h"

static num A_num;
static num B_num;
static num T_num;


static str temp_str[];

int main() {
    int target = 0;
    for(;;) {
        switch (target) {
        case 0:
            /* line 10 */
            TAB_void(30.0000f);
            print_str("SINE WAVE");
            print_char('\n');
            /* line 20 */
            TAB_void(15.0000f);
            print_str("CREATIVE COMPUTING  MORRISTOWN, NEW JERSEY");
            print_char('\n');
            /* line 30 */
            print_char('\n');
            print_char('\n');
            print_char('\n');
            print_char('\n');
            print_char('\n');
            /* line 40 */
            /*REMARKABLE PROGRAM BY DAVID AHL*/
            /* line 50 */
            let_num(&B_num,0.0000f);
            /* line 100 */
            /*  START LONG LOOP*/
            /* line 110 */
            for_num(&T_num,0.0000f,40.0000f,0.2500f,-1);
        case -1:
            /* line 120 */
            let_num(&A_num,INT_num(26.0000f+25.0000f*SIN_num(T_num)));
            /* line 130 */
            TAB_void(A_num);
            /* line 140 */
            if (B_num==1.0000f) {
                target = 180;
                break;
            }
            /* line 150 */
            print_str("CREATIVE");
            print_char('\n');
            /* line 160 */
            let_num(&B_num,1.0000f);
            /* line 170 */
            target = 200;
            break;
        case 180: /* line 180 */
            print_str("COMPUTING");
            print_char('\n');
            /* line 190 */
            let_num(&B_num,0.0000f);
        case 200: /* line 200 */
            if (next(&target)) break; /* NEXT T */
            /* line 999 */
            exit(0);
        case -2:
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
