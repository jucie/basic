#include "basiclib.h"

static num I_num;
static num N_num;


static str temp_str[];

int main() {
    int target = 0;
    for(;;) {
        switch (target) {
        case 0:
            /* line 1 */
            TAB_void(28.0000f);
            print_str("RUSSIAN ROULETTE");
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
            print_str("THIS IS A GAME OF >>>>>>>>>>RUSSIAN ROULETTE.");
            print_char('\n');
        case 10: /* line 10 */
            print_char('\n');
            print_str("HERE IS A REVOLVER.");
            print_char('\n');
        case 20: /* line 20 */
            print_str("TYPE \'1\' TO SPIN CHAMBER AND PULL TRIGGER.");
            print_char('\n');
            /* line 22 */
            print_str("TYPE \'2\' TO GIVE UP.");
            print_char('\n');
            /* line 23 */
            print_str("GO");
            /* line 25 */
            let_num(&N_num,0.0000f);
        case 30: /* line 30 */
        case -1:
            input_to_buffer();
            if (!input_num(&I_num)) {
                target = -1;
                break;
            }
            /* line 31 */
            if (I_num!=2.0000f) {
                target = 35;
                break;
            }
            /* line 32 */
            print_str("     CHICKEN!!!!!");
            print_char('\n');
            /* line 33 */
            target = 72;
            break;
        case 35: /* line 35 */
            let_num(&N_num,N_num+1.0000f);
            /* line 40 */
            if (RND_num(1.0000f)>0.8333f) {
                target = 70;
                break;
            }
            /* line 45 */
            if (N_num>10.0000f) {
                target = 80;
                break;
            }
            /* line 50 */
            print_str("- CLICK -");
            print_char('\n');
            /* line 60 */
            print_char('\n');
            target = 30;
            break;
        case 70: /* line 70 */
            print_str("     BANG!!!!!   YOU\'RE DEAD!");
            print_char('\n');
            /* line 71 */
            print_str("CONDOLENCES WILL BE SENT TO YOUR RELATIVES.");
            print_char('\n');
        case 72: /* line 72 */
            print_char('\n');
            print_char('\n');
            print_char('\n');
            /* line 75 */
            print_str("...NEXT VICTIM...");
            print_char('\n');
            target = 20;
            break;
        case 80: /* line 80 */
            print_str("YOU WIN!!!!!");
            print_char('\n');
            /* line 85 */
            print_str("LET SOMEONE ELSE BLOW HIS BRAINS OUT.");
            print_char('\n');
            /* line 90 */
            target = 10;
            break;
            /* line 99 */
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
