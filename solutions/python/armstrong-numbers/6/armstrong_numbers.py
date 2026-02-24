""" A function providing that checks if a number is a valid Armstrong number."""

def is_armstrong_number(number):
    """
    Determines armstrong  number.

    Args:
        a (int): number.
   
    Returns:
        boolean: Is the number a valid armstrong number.
    """

    # Convert number to  a string number
    str_num = str(number)
    
    # Get length of string number
    size = len(str_num)
    
    # Convert string number to a list of numbers
    lst = list(str_num)
    
    # Calculate Armstrong number   
    return sum([int(digit)**size for digit in lst]) == number