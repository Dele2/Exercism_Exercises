def rotate(text, key):
    letters = "abcdefghijklmnopqrstuvwxyz"
    cipher = letters[key:] + letters[:key]
    message = ""

    for ch in text:
        if ch.isupper() and ch.lower() in letters:
            message += cipher[letters.find(ch.lower())].upper()
        elif ch in letters:
            message += cipher[letters.find(ch)]
        else:
            message += ch

    return message
