
def main(total=0, past=0, current=1, limit=4_000_000):
    next = past + current
    if next % 2 == 0:
        total += next
    if next < limit:
        return main(total, current, next, limit)
    return total


if __name__ == '__main__':
    t = main()
    print(t)
