import sys

t = int(input())

def solve():
    # Return the largest set of words
    # that is a set of pairs that rhyme
    # Read number of words
    n = int(input())
    # print("number of words = %d" % n)
    words = []
    for _ in range(n):
        words.append(input())
    # print("words: %s" % words)
    return group(1, words)

def group(idx, wordlist):
    # Return (2x) the number of rhyming pairs contained
    # in this wordlist, starting rhyming from index len-idx
    if len(wordlist) <= 1:
        return 0
    elif len(wordlist) <= 3 and idx > 1:
        return 2 # Best we can get with 2 or 3 matching words


    letterset = {}
    for word in wordlist:
        # append to set
        letter = word[len(word) - idx ]
        if letter in letterset:
            letterset[letter].append(word)
        else:
            letterset[letter] = [word]

    # Do we need to handle the special case where two words are indentical?
    # What about when one word is shorter than idx letters?

    # print("index: %d, wordlist: %s" % (idx, wordlist))
    # print(letterset)
    sum = 0
    for l in letterset:
        sum += group(idx+1, letterset[l])
    if idx > 1:
        return 2 + sum
    else:
        return sum

for c in range(1, t+1):
    print("Case #%d: %s" % (c, solve()))
