def check_tail(input_str):
    tail = input_str[-1]
    return tail.isdigit() or tail == 'X'

def check_head(input_str):
    head = input_str[:-1]
    head = head.replace("-", "")
    return head.isdigit()


def list_numbers(input_str):
    num_list = []  
    for ch in input_str:
        if ch.isdigit():
            num_list.append(int(ch))
        elif ch == 'X':
            num_list.append(10)
    return num_list
    


def calculate(lst):
    total = 0
    for idx, num in enumerate(lst, start=-len(lst)):
        countdown_idx = abs(idx)
        total += (num * countdown_idx)
    return total

def is_valid(isbn):
    if isbn == "":
        return False
    elif not check_head(isbn) or not check_tail(isbn):
        return False
    else:
        number_list = list_numbers(isbn)
        if len(number_list) == 10:
            return calculate(number_list) % 11 == 0
        else:
            return False