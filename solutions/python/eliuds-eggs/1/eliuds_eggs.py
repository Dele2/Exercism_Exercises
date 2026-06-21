def egg_count(display_value):
    lst = []
    while display_value > 0:
        lst.append(display_value % 2)
        display_value //= 2

    return lst.count(1)