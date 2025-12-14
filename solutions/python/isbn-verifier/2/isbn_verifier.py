import re

def is_valid(isbn):
    isbn = isbn.replace("-","")
    match = re.search(r'^(\d{9})(\d|X)$', isbn)

    if not match :
        return False
    else:
        tail = "10" if match.group(2) == 'X' else match.group(2)
        digits = match.group(1)
        result = sum((10 - i) * int(digit) for i, digit in enumerate(digits))
        return (result + int(tail)) % 11 == 0
