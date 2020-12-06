#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#define MAX_NUM_ELEMENTS 50
#define TRUE        1
#define FALSE       0
#define LETTERS     26



int isEmpty(char* data){
    if(data[0] == '\n'||data[0] == '\0'){
        return TRUE;
    }
    return FALSE;
}

int isEol( char  val){
    if (val == '\n' || val == '\0'){
        return TRUE;
    }
    return FALSE;
}

void resetArray(int *array){

    for (int i = 0; i < LETTERS; i++){
        array[i]=0;
    }
}

int main()
{
    printf ("-------------DAY 06-------------\n");
    FILE *filep;
    char input[MAX_NUM_ELEMENTS];
    char* filename = "input.txt";
    filep = fopen(filename, "r");
    if (filep == NULL){
        printf("Could not open file %s\n ",filename);
    }


    int sum =0;
    int row_info[LETTERS = {0}; //a-97, z-122
    int new_row[LETTERS]={0};
    int count=0;
    int persons=0;
    int yes_answers =0;
    while (fgets(input, MAX_NUM_ELEMENTS , filep) != NULL){
        if ( isEmpty(input) == FALSE){
            persons++;
            int idx = 0;
            int check_idx;

            while ( isEol(input[idx]) == FALSE ){
                check_idx = input[idx]-'a';
                new_row[check_idx]++;
                if (  row_info[check_idx] != 1 ) {
                    row_info[check_idx] = 1;
                    count++;
                }


                idx++;
            }
        }
        else{
            resetArray(row_info);
            yes_answers =0;

            for (int l=0; l<LETTERS; l++){
                if(new_row[l] == persons){
                    yes_answers++;
                }
            }
            sum = sum + yes_answers;
            persons = 0;
            resetArray(new_row);
        }

        printf("\n");

    }
    yes_answers =0;
    for (int l=0; l<LETTERS; l++){
        if(new_row[l] == persons){
            yes_answers++;
            printf("ans: %d\n", yes_answers);
        }
    }
            sum = sum + yes_answers;
    printf("count: %d\n", count);
    printf("sum: %d\n", sum);
    fclose(filep);
    return 0;
}

