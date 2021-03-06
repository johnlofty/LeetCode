#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>


void print_merged_array(int* head, int size){
    printf("Array ");
    for (int i = 0; i < size; i ++){
        printf(" %d ", *(head + i));
    }
    printf(" =\n");
}

double findMedianSortedArrays(int* nums1, int nums1Size, int* nums2, int nums2Size){
    bool is_odd = true;
    int sum = nums1Size + nums2Size;
    if (sum % 2 == 0){
        is_odd = false;
    }
    int* head = malloc(sizeof(int) * sum);
    int idx = 0;
    int idx_1 = 0;
    int idx_2 = 0;
    bool num1_end = false;
    bool num2_end = false;
    while (idx < sum){
        if (idx_1 >= nums1Size){
            num1_end = true;
        }
        if (idx_2 >= nums2Size){
            num2_end = true;
        }

        if (num1_end){
            memcpy(head + idx, nums2 + idx_2, sizeof(int));
            idx_2++;
        }else if (num2_end){
            memcpy(head + idx, nums1 + idx_1, sizeof(int));
            idx_1++;
        }else{
            if (*(nums1 + idx_1) < *(nums2+idx_2)){
                memcpy(head +idx, nums1 + idx_1, sizeof(int));
                idx_1++;
            }else{
                memcpy(head + idx, nums2 + idx_2, sizeof(int));
                idx_2++;
            }
        }
        idx++;
    }

    print_merged_array(head, sum);

    int median = sum / 2;
    if (is_odd){
        return (double)(*(head + median));
    }else{
        return (*(head + median) + *(head + median-1))/2.0;
    }
}

void test1(){
    int num1[] = {1, 3};
    int num2[] = {2};
    double ret = findMedianSortedArrays((int *)num1, 2, (int *)num2, 1);
    printf("test2 %lf\n", ret);
}

void test2(){
    int num1[] = {1, 2};
    int num2[] = {3, 4};
    double ret = findMedianSortedArrays((int *)num1, 2, (int *)num2, 2);
    printf("test2 %lf\n", ret);
}

void test3(){
    int num1[] = {0, 0};
    int num2[] = {0, 0};
    double ret = findMedianSortedArrays((int *)num1, 2, (int *)num2, 2);
    printf("test3 %lf\n", ret);
}

void test4(){
    int num1[] = {};
    int num2[] = {1};
    double ret = findMedianSortedArrays((int *)num1, 0, (int *)num2, 1);
    printf("test3 %lf\n", ret);
}
void test5(){
    int num1[] = {2};
    int num2[] = {};
    double ret = findMedianSortedArrays((int *)num1, 1, (int *)num2, 0);
    printf("test3 %lf\n", ret);
}

int main(){
   
    test1();
    test2();
    test3();
    test4();
    test5();
	return 1;
}
