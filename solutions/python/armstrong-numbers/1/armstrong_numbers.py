
def is_armstrong_number(number):
    # Convert number to  a string number
    str_num = str(number)
    
    # Get length of string number
    size = len(str_num)
    
    # Convert string number to a list of numbers
    lst = list(str_num)
    
    # Calculate Armstrong number
    lst = [int(n)**size for n in lst]
    
    return sum(lst) == number