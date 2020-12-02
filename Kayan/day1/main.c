#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>


#define MAXNUM 7
#define LISTSIZE 200
#define DESIRED_ANS 2020

int day1(void){

    FILE *fp;
    char str[MAXNUM];
    char* filename = "input.txt";

    fp = fopen(filename, "r");
    if (fp == NULL){
        printf("Could not open file %s",filename);
        return 1;
    }

    uint16_t list_vaues[LISTSIZE];
    uint16_t counter = 0;
    while (fgets(str, MAXNUM, fp) != NULL){
        list_vaues[counter] = atoi(str);
        counter++;
    }

    fclose(fp);
    counter = 0;

    uint16_t i;
    uint16_t j;
    uint32_t ans;
    printf("The first puzzle------------------------------\n");
    for (i=0; i < LISTSIZE-1; i++){
        for(j=i+1; j < LISTSIZE; j++ ){
            if ( list_vaues[i]+list_vaues[j] == DESIRED_ANS){
                ans = list_vaues[i]*list_vaues[j];
                printf("1st: %d\n2nd: %d\n", list_vaues[i], list_vaues[j]);
                printf("ANSWER: %d\n", ans);
                break;
            }
        }
    }

    printf("\nThe second puzzle-----------------------------\n");
    uint16_t k;
    for (i=0; i < LISTSIZE-2; i++){
        for (j=i+1; j < LISTSIZE-1; j++ ){
            for (k=j+1; k < LISTSIZE; k++){
                 if ( list_vaues[i]+list_vaues[j]+list_vaues[k] == DESIRED_ANS){
                    ans = list_vaues[i]*list_vaues[j]*list_vaues[k];
                    printf("1st: %d\n2nd: %d\n3rd: %d \n", list_vaues[i], list_vaues[j], list_vaues[k]);
                    printf("ANSWER: %d\n", ans);
                    break;
                }
            }
        }
    }
    return 0;
}


int main()
{
    day1();
    return 0;
}
