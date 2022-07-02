import math
import re

COUNT = 100000000
CHAR = "x"
PATTERN = re.compile(f'^(?!({CHAR}{CHAR}+)\\1+$|^{CHAR}$)')

def is_prime_regex(num):
    chars = CHAR * num
    if bool(PATTERN.search(chars)):
        return True
    else:
        return False

def is_prime_python(num):
    if num == 1:
        return False
    if num == 2 or num == 3:
        return True
    if num % 2 == 0:
        return False

    root = math.floor(math.sqrt(num))
    for p in range(3, (root + 1), 2):
        if num % p == 0:
            return False

    return True

if __name__ == '__main__':
    for i in range(1, COUNT):
        is_prime = is_prime_regex(i)
        message = f"{i} is prime" if is_prime else f"{i} is not prime"
        print(message)
