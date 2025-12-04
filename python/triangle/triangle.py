def equilateral(sides):
    a, b, c = sides
    # Check if all sides are positive and form a valid triangle
    if a <= 0 or b <= 0 or c <= 0:
        return False
    if a + b <= c or b + c <= a or a + c <= b:
        return False
    # Check if all sides are equal
    return a == b == c


def isosceles(sides):
    a, b, c = sides
    # Check if all sides are positive and form a valid triangle
    if a <= 0 or b <= 0 or c <= 0:
        return False
    if a + b <= c or b + c <= a or a + c <= b:
        return False
    # Check if at least two sides are equal
    return a == b or b == c or a == c


def scalene(sides):
    a, b, c = sides
    # Check if all sides are positive and form a valid triangle
    if a <= 0 or b <= 0 or c <= 0:
        return False
    if a + b <= c or b + c <= a or a + c <= b:
        return False
    # Check if all sides are different
    return a != b and b != c and a != c
