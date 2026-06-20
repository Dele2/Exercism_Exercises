
def rotate(text, key):
    """
    Returns a string where letters from the input text have been shifted
    forward by the value of the key.

    Parameters:
    -text   (string)
    -key    (int)

    Returns:
    (string) Text where the letters have been rotated.
    """
    letters = "abcdefghijklmnopqrstuvwxyz"
    message = ""

    for letter in text:
        if letter.isupper() and letter.lower() in letters:
            message += letters[(letters.find(letter.lower()) + key) % 26].upper()
        elif letter in letters:
            message += letters[(letters.find(letter) + key) % 26]
        else:
            message += letter

    return message