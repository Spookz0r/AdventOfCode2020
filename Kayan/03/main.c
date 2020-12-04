#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <inttypes.h>

#define LINES           323
#define MAX_COLUMNS     100
#define COLUMNS         32
#define RIGHT3          3
#define DOWN1           1
#define START_POS       0
#define NUM_OF_SLOPES   5


int calc_encounters(int inc_right, int inc_down, char map[LINES][COLUMNS]){
    int counter = 0;
    int row_pos = 0;
    int col_pos = 0;
    int wrap = COLUMNS-1;
    
    while (row_pos < LINES-1 ){
        col_pos = (col_pos + inc_right) % wrap;
        row_pos = row_pos + inc_down;
        if( map[row_pos][col_pos] == '#'){
            counter++;
        }
    }
    
    return counter;
}

int main()
{
    printf("starting\n");
    FILE* filep;
    char* filename = "input.txt";
    filep = fopen(filename, "r");
    if (filep == NULL){
        printf(" Could not open file: %s\n",filename);
    }

    char row[MAX_COLUMNS];
    char map[LINES][COLUMNS];
    int index = 0;
    while (fgets(row, MAX_COLUMNS, filep) != NULL){
        strncpy (map[index], row, COLUMNS);
        map[index][COLUMNS-1] = '\0';
        index++;
    }

    int slopes[NUM_OF_SLOPES][2] = {
            {1, 1},
            {3, 1},
            {5, 1},
            {7, 1},
            {1, 2} };
    
    
    int encounters[NUM_OF_SLOPES]= {};
    uint32_t total = 1;
    uint32_t temp=0;
    int i = 0;
    while (i < 5){
        encounters[i]=calc_encounters(slopes[i][0], slopes[i][1], map);
        printf("Slope: %d encounters: %d\n", i, encounters[i]);
        temp = total * encounters[i];
        total = temp;
        i++;
        
    }
    printf("total: %" PRIu32"\n", total);

    fclose(filep);
    return 0;
}
//