#include <SFML/Graphics/Transform.h>

float GetTransformAtIndex(sfTransform *t, int i) {
	return t->matrix[i];
}

void SetTransformAtIndex(sfTransform *t, int i, float value) {
	t->matrix[i] = value;
}
