#include "basiclib.h"



static str temp_str[];

int main() {
    int target = 0;
    for(;;) {
        switch (target) {
        case 0:
            /* line 10 */
            print_str("I am test 1");
            print_char('\n');
            /* line 20 */
            exit(0);
        case -1:
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