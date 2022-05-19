# zcat 1.dz > 1.dict.txt
# f = open(fname, 'rb')
# f.seek(618037)
# f.read(119).decode()

def decode_char(char):
    if "A" <= char <= "Z":
        return ord(char) - 65
    if "a" <= char <= "z":
        return ord(char) - 71
    if "0" <= char <= "9":
        return ord(char) + 4
    if char == '+':
        return 62
    if char == '/':
        return 63
    raise ValueError


def decode_str(s):
    result = 0
    for i, char in enumerate(s[::-1]):
        result += decode_char(char) * 64 ** i
    return result


print(decode_str("CW41"))
print(decode_str("B3"))
