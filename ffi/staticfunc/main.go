package main

/*
#include <stdbool.h>
#include <stdio.h>
#include <time.h>

static bool do_we_have_2038() {
	time_t t = time(NULL);
	struct tm tm = *localtime(&t);
	if (tm.tm_year + 1900 == 2038) {
		return true;
	} else {
		return false;
	}
}
*/
import "C"

import "fmt"

func main() {
	fmt.Println(C.do_we_have_2038())
}
