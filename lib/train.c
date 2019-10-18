#include "basiclib.h"

static num A_num;
static str A_str;
static num C_num;
static num D_num;
static num E_num;
static num T_num;
static num V_num;


static str temp_str[];

int main() {
    int target = 0;
    for(;;) {
        switch (target) {
        case 0:
            /* line 1 */
            TAB_void(33.0000f);
            print_str("TRAIN");
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
            print_str("TIME - SPEED DISTANCE EXERCISE");
            print_char('\n');
            print_char('\n');
        case 10: /* line 10 */
            let_num(&C_num,INT_num(25.0000f*RND_num(1.0000f))+40.0000f);
            /* line 15 */
            let_num(&D_num,INT_num(15.0000f*RND_num(1.0000f))+5.0000f);
            /* line 20 */
            let_num(&T_num,INT_num(19.0000f*RND_num(1.0000f))+20.0000f);
            /* line 25 */
            print_str(" A CAR TRAVELING");
            print_num(C_num);
            print_str("MPH CAN MAKE A CERTAIN TRIP IN");
            print_char('\n');
            /* line 30 */
            print_num(D_num);
            print_str("HOURS LESS THAN A TRAIN TRAVELING AT");
            print_num(T_num);
            print_str("MPH.");
            print_char('\n');
            /* line 35 */
            print_str("HOW LONG DOES THE TRIP TAKE BY CAR");
        /* line 40 */
        case -1:
            input_to_buffer();
            if (!input_num(&A_num)) {
                target = -1;
                break;
            }
            /* line 45 */
            let_num(&V_num,D_num*T_num/(C_num-T_num));
            /* line 50 */
            let_num(&E_num,INT_num(ABS_num((V_num-A_num)*100.0000f/A_num)+0.5000f));
            /* line 55 */
            if (E_num>5.0000f) {
                target = 70;
                break;
            }
            /* line 60 */
            print_str("GOOD! ANSWER WITHIN");
            print_num(E_num);
            print_str("PERCENT.");
            print_char('\n');
            /* line 65 */
            target = 80;
            break;
        case 70: /* line 70 */
            print_str("SORRY.  YOU WERE OFF BY");
            print_num(E_num);
            print_str("PERCENT.");
            print_char('\n');
        case 80: /* line 80 */
            print_str("CORRECT ANSWER IS");
            print_num(V_num);
            print_str("HOURS.");
            print_char('\n');
            /* line 90 */
            print_char('\n');
            /* line 95 */
            print_str("ANOTHER PROBLEM (YES OR NO)");
        /* line 100 */
        case -2:
            input_to_buffer();
            if (!input_str(&A_str)) {
                target = -2;
                break;
            }
            /* line 105 */
            print_char('\n');
            /* line 110 */
            if (strcmp(A_str,"YES")==0) {
                target = 10;
                break;
            }
            /* line 999 */
            exit(0);
        case -3:
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