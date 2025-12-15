def equilateral(sides):
     return sides[0] == sides[1] == sides[2] and is_triangle(sides)

def isosceles(sides):
    if equilateral(sides):
        return True
    return (sides[0] == sides[1] and sides[0] != sides[2] or 
            sides[1] == sides[2] and sides[1] != sides[0] or
            sides[0] == sides[2] and sides[0] != sides[1]) and is_triangle(sides)

def scalene(sides):
    return sides[0] != sides[1] != sides[2] != sides[0] and is_triangle(sides)

def is_triangle(sides):
    return (sides[0] + sides[1] > sides[2] and 
            sides[2] + sides[0] > sides[1] and
            sides[2] + sides[1] > sides[0])