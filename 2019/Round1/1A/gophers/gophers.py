import sys
t, n, m = map(int, raw_input().split())

def solve():
    primeremainders = {}
    # For each of the 7 primes
    for prime in [2, 3, 5, 7, 11, 13, 17]:
        # Set all windmills to this prime
        print(" ".join([str(prime)]*18))
        sys.stdout.flush()
        # Process response
        blades = map(int, raw_input().split())
        rem = sum(blades)
        # print(blades)
        # print("Average of blades = " + str(avg))
        primeremainders[prime * 18] = rem

    print("Prime remainders:")
    print(primeremainders)

    candidate = primeremainders[17 * 18]
    while candidate <= m:
        # check candidate number for division by other primes
        for check in [2,3,5,7,11,13]:
            if candidate % (check * 18) != primeremainders[check * 18]:
                candidate += 17 * 18
                break
        else:
            # if for loop finished correctly
            # print("Answer is " + str(candidate))
            # print("There are " + str(candidate * 18) + " Gophers")
            print(candidate)
            sys.stdout.flush()
            raw_input() # eat the returned value
            break
    

for __ in xrange(t):
    solve()
