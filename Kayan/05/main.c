#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_NUM_ELEMENTS    10
#define NUM_OF_IDS          781
#define ROW_PROD_ID         8
#define ROW_INFO_SIZE       7
#define COL_INFO_SIZE       3
#define SEAT_ROWS   128
#define SEAT_COLS   8
/*
#define FRONT   'F'
#define BACK    'B'
#define RIGHT   'R'
#define LEFT    'L'
*/
typedef struct {
    int row;
    int col;
    int bp_id;
}BoardingPass;


int getRowID(int row, int col){
    int product = row * ROW_PROD_ID + col;
    return product;
}



BoardingPass getBoardingPassInfo(char* data){
    BoardingPass b_pass;


    int upper = SEAT_ROWS-1;
    int lower = 0;

    int i = 0;
    //int row;
    //int col;
    int temp;
    int pos;
    while (i < ROW_INFO_SIZE){
        if ( data[i] == 'F'){
            //same lower
            temp = (upper - lower +1 )/2 + lower - 1;
            upper = temp;
            pos = upper;
        }
        if ( data[i] == 'B'){
            //same upper
            temp = ((upper - lower + 1)/2) + lower;
            lower = temp;
            pos = lower;
        }
        i++;
    }
    b_pass.row = pos;

    upper = 7;
    lower = 0;
    while (i < ROW_INFO_SIZE+3){
        if ( data[i] == 'F' || data[i] == 'R'){
            //same lower
            temp = (upper - lower +1 )/2 + lower - 1;
            upper = temp;
            pos = upper;
        }
        if ( data[i] == 'B' || data[i] == 'L'){
            //same upper
            temp = ((upper - lower + 1)/2) + lower;
            lower = temp;
            pos = lower;
        }
        i++;
    }
    b_pass.col = pos;

    b_pass.id = getRowID(b_pass.row, b_pass.col);
    return b_pass;
}




int main()
{
    printf ("-------------DAY 05-------------");
    FILE *filep;
    char input[MAX_NUM_ELEMENTS];
    char* filename = "test_input.txt";
    filep = fopen(filename, "r");
    if (filep == NULL){
        printf("Could not open file %s\n",filename);
    }


    BoardingPass boarding_pass[NUM_OF_IDS];
    int max_id = 0;
    int i = 0;
    while (fgets(input, MAX_NUM_ELEMENTS , filep) != NULL){
        printf(" input: %s  - ", input);
        boarding_pass[i] = getBoardingPassInfo(input);
        i++;
        if (max_id < boarding_pass[i].id){
            max_id = boarding_pass[i].id;
            }
        printf("id: %d\n", boarding_pass[i].id);
    }

    printf("max id: %d", max_id);
    fclose(filep);
    return 0;
}
