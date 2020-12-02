#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>

#define ELEMENTS    100
#define TRUE        1
#define FALSE       0

typedef struct {
    char* val1;
    char* val2;
    char* character;
    char* password;
}Row_info;

Row_info parse_row_input(char input[]){
    Row_info row;
    row.val1 = strtok(input, "-");
    row.val2 = strtok(NULL, " ");
    row.character = strtok(NULL, ":");
    row.password = strtok(NULL, "\0");

    return row;
}

/** Puzzle 1 **/
uint8_t verify_password_rule1(Row_info row){

    uint16_t char_occurence = 0;
    int pos = 0;
    while(row.password[pos]!='\0'){
        if (row.password[pos] == *row.character){
            char_occurence++;
        }
        pos++;
    }
    if (char_occurence >= atoi(row.val1) &&
        char_occurence <= atoi(row.val2)){
        return TRUE;
    }
    return FALSE;
}

/** Puzzle 2 **/
uint8_t verify_password_rule2(Row_info row){
    int pos1 = atoi(row.val1);
    int pos2 = atoi(row.val2);

    if ( (row.password[pos1]==*row.character && row.password[pos2]!=*row.character) ||
         (row.password[pos1]!= *row.character && row.password[pos2]==*row.character)){
        return TRUE;
    }
    return FALSE;
}

int main()
{
    FILE *fp;
    char input[ELEMENTS];
    char* filename = "input.txt";
    fp = fopen(filename, "r");
    if (fp == NULL){
        printf("Could not open file %s\n",filename);
    }

    Row_info row;
    uint16_t num_correct_pw1 = 0;
    uint16_t num_correct_pw2 = 0;

    while (fgets(input, ELEMENTS , fp) != NULL){
        row = parse_row_input(input);

        if (verify_password_rule1(row) == TRUE){
            num_correct_pw1++;
        }

        if (verify_password_rule2(row) == TRUE){
            num_correct_pw2++;
        }
    }
    fclose(fp);
    printf("1st: Number of correct passwords: %d\n", num_correct_pw1);
    printf("2nd: Number of correct passwords: %d\n\n", num_correct_pw2);

    return 0;
}
