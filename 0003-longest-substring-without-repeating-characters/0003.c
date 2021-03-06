#include <stdio.h>
#include <string.h>

void init_alpha(int alpha[]){
    for(int i= 0; i < 255; i ++){
        alpha[i] = 0;
    }
}

void print_al(int alpa[]){
	printf("alpa ======= ");
    for(int i= 0; i < 26; i ++){
		printf("%d ", alpa[i]);
	}
	printf("  ===\n");
}

int lengthOfLongestSubstring(char * s){
    int max = 0;
    int alpha[255];
    init_alpha(alpha);

    int str_len = strlen(s);
    int start_ptr = 0;
	int now_ptr = 0;
	int now_num = 0;
	int now_max = 0;
    while(start_ptr <= str_len - 1){
		now_max = 0;
		now_ptr = start_ptr;

		while(now_ptr <= str_len - 1){
			now_num = (unsigned char)s[now_ptr];
			if (alpha[now_num] == 0){
				alpha[now_num] = 1;
				now_max += 1;
			} else{
				break;
			}
			now_ptr += 1;
    	}

		if (now_max > max){
			max = now_max;
		}
		start_ptr += 1;
		init_alpha(alpha);
	}
	return max;
}

int lengthOfLongestSubstringV2(char * s){
	
	
}

int main(){
	char* s = " ";
	printf("string %s max_substring %d\n", s, lengthOfLongestSubstring(s));

	char* s2 = "";
	printf("string %s max_substring %d\n", s, lengthOfLongestSubstring(s2));


	return 1;
}
