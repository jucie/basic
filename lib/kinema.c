#include "basiclib.h"

static num A_num;
static num G_num;
static num Q_num;
static num T_num;
static num V_num;


static str temp_str[];

int main() {
    int target = 0;
    for(;;) {
        switch (target) {
        case 0:
            /* line 10 */
            TAB_void(33.0000f);
            print_str("KINEMA");
            print_char('\n');
            /* line 20 */
            TAB_void(15.0000f);
            print_str("CREATIVE COMPUTING  MORRISTOWN, NEW JERSEY");
            print_char('\n');
            /* line 30 */
            print_char('\n');
            print_char('\n');
            print_char('\n');
        case 100: /* line 100 */
            print_char('\n');
            /* line 105 */
            print_char('\n');
            /* line 106 */
            let_num(&Q_num,0.0000f);
            /* line 110 */
            let_num(&V_num,5.0000f+INT_num(35.0000f*RND_num(1.0000f)));
            /* line 111 */
            print_str("A BALL IS THROWN UPWARDS AT");
            print_num(V_num);
            print_str("METERS PER SECOND.");
            print_char('\n');
            /* line 112 */
            print_char('\n');
            /* line 115 */
            let_num(&A_num,0.0500f*(num)pow(V_num,2.0000f));
            /* line 116 */
            print_str("HOW HIGH WILL IT GO (IN METERS)");
            /* line 117 */
            push_address(-1);
            target = 500;
            break;
        case -1:
            /* line 120 */
            let_num(&A_num,V_num/5.0000f);
            /* line 122 */
            print_str("HOW LONG UNTIL IT RETURNS (IN SECONDS)");
            /* line 124 */
            push_address(-2);
            target = 500;
            break;
        case -2:
            /* line 130 */
            let_num(&T_num,1.0000f+INT_num(2.0000f*V_num*RND_num(1.0000f))/10.0000f);
            /* line 132 */
            let_num(&A_num,V_num-10.0000f*T_num);
            /* line 134 */
            print_str("WHAT WILL ITS VELOCITY BE AFTER");
            print_num(T_num);
            print_str("SECONDS");
            /* line 136 */
            push_address(-3);
            target = 500;
            break;
        case -3:
            /* line 140 */
            print_char('\n');
            /* line 150 */
            print_num(Q_num);
            print_str("RIGHT OUT OF 3.");
            /* line 160 */
            if (Q_num<2.0000f) {
                target = 100;
                break;
            }
            /* line 170 */
            print_str("  NOT BAD.");
            print_char('\n');
            /* line 180 */
            target = 100;
            break;
        case 500: /* line 500 */
        case -4:
            input_to_buffer();
            if (!input_num(&G_num)) {
                target = -4;
                break;
            }
            /* line 502 */
            if (ABS_num((G_num-A_num)/A_num)<0.1500f) {
                target = 510;
                break;
            }
            /* line 504 */
            print_str("NOT EVEN CLOSE....");
            print_char('\n');
            /* line 506 */
            target = 512;
            break;
        case 510: /* line 510 */
            print_str("CLOSE ENOUGH.");
            print_char('\n');
            /* line 511 */
            let_num(&Q_num,Q_num+1.0000f);
        case 512: /* line 512 */
            print_str("CORRECT ANSWER IS ");
            print_num(A_num);
            print_char('\n');
            /* line 520 */
            print_char('\n');
            /* line 530 */
            pop_address(&target);
            break;
            /* line 999 */
            exit(0);
        case -5:
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
