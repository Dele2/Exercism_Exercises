import sys

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
    cipher = letters[key:] + letters[:key]
    message = ""

    for letter in text:
        if letter.isupper() and letter.lower() in letters:
            message += cipher[letters.find(letter.lower())].upper()
        elif letter in letters:
            message += cipher[letters.find(letter)]
        else:
            message += letter

    return message