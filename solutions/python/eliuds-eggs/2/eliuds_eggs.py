"""
A function that takes a decimal number, converts
into a binary number and returns the counted occurances of 1's.
"""
def egg_count(display_value):
    """
    Parameters:
    -display_value  (string)

    Returns:
    - (integer) The total number of 1's
    """
    lst = []
    while display_value > 0:
        lst.append(display_value % 2)
        display_value //= 2

    return lst.count(1)

print(egg_count (89))