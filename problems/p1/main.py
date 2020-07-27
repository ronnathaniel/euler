
"""
Ron Nathaniel
July 26 2020
"""

def main(limit=1000):
    total = 0
    for i in range(limit):
        if i % 3 == 0 or i % 5 == 0:
            total += i
    return total


if __name__ == '__main__':
    t = main()
    print(t)
