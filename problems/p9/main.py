
"""
Ron Nathaniel
July 26 2020
"""

def isPythTriplet (*triplet):
    try:
        return triplet[0]**2 + triplet[1]**2 == triplet[2]**2
    except:
        return False

def isSumReached (*triplet):
    return sum(triplet) == 1000


def main ():
    for i in range(2, 1000):
        for j in range(1, i):
            for k in range(j):
                if isPythTriplet(k, j, i) and isSumReached(k, j, i):
                    return k, j, i

if __name__ == '__main__':
    t = main()
    print(t)
