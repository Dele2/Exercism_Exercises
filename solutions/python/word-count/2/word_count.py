def count_words(sentence):

    # Remove punctuation and split text
    word_list = sentence.replace(',',' ').replace('-',' ').replace('_',' ').split()

    # initializing bad characters list
    bad_chars = [".", ":", "!", "&","@","$","%","^"]
    clean_words = []

    for word in word_list:
        # If the text starts with or ends with an apostrophe, remove it.
        while word[0] == "'" or word[-1] == "'":
            if word[0].startswith("'"):
                word = word[1:]
            if word[-1].endswith("'"):
                word = word[:-1]
        # Using replace() to remove bad characters
        for i in bad_chars:
            if i in word:
                word = word.replace(i,"")
        clean_words.append(word)
    

    # Place the list of words into a dictionary to count each word
    word_dict = {}
    for word in clean_words:
        word_dict.setdefault(word.lower(), 0)
        word_dict[word.lower()] += 1

    return word_dict