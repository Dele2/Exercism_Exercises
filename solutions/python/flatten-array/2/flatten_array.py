def flatten(iterable):
    """
    This function generates a one dimension list from multi nested list
    """
    flat_list = []
    for item in iterable:
        if isinstance(item, list):
            flat_list.extend(flatten(item))
        elif item is not None:
            flat_list.append(item)
    return flat_list
