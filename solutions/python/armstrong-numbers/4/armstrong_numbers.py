
def is_armstrong_number(number):
    """
    Determines armstrong  number.

    Args:
        a (int): number.
   
    Returns:
        boolean: Calculated armstrong number against inputted number.
    """
    # Convert number to  a string number
    str_num = str(number)
    
    # Get length of string number
    size = len(str_num)
    
    # Convert string number to a list of numbers
    lst = list(str_num)
    
    # Calculate Armstrong number   
    return sum([int(digit)**size for digit in lst]) == number