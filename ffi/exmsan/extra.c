#include <stdint.h>
#include <stdlib.h>

typedef struct { uintptr_t x, y; char b; } T;

extern void F(T);

void CF(int x) {
	T *t = malloc(sizeof(T));
	t->x = x;
	t->y = x;
	t->b = (char)x;
	F(*t);
}
