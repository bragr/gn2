# This is lame "vision" but it proves the point
import json
import sys
PIXELS = 5*5
LETTERS = 26
OUTPUT = 'vision.json'

# This is either horrible, or genius. Why not both?
alphabet = {
'a': 
"""
.XXX.
X...X
XXXXX
X...X
X...X
""",
'b':  
"""
XXXX.
X...X
XXXX.
X...X
XXXX.
""",
'c':  
"""
.XXXX
X....
X....
X....
.XXXX
""",
'd': 
"""
XXXX.
X...X
X...X
X...X
XXXX.
""",
'e':  
"""
XXXXX
X....
XXXX.
X....
XXXXX
""",
'f':  
"""
XXXXX
X....
XXXX.
X....
X....
""",
'g':  
"""
.XXXX
X....
X..XX
X...X
.XXXX
""",
'h': 
"""
X...X
X...X
XXXXX
X...X
X...X
""",
'i': 
"""
.XXX.
..X..
..X..
..X..
.XXX.
""",
'j':  
"""
XXXXX
....X
....x
X...X
.XXX.
""",
'k': 
"""
.X..X
.X.X.
.XX..
.X.X.
.X..X
""",
'l':  
"""
X....
X....
X....
X....
XXXXX
""",
'm': 
"""
X...X
XX.XX
X.X.X
X...X
X...X
""",
'n': 
"""
.X..X
.XX.X
.X.XX
.X..X
.X..X
""",
'o': 
"""
.XXX.
X...X
X...X
X...X
.XXX.
""",
'p':  
"""
XXXX.
X...X
XXXX.
X....
X....
""",
'q':  
"""
.XXX.
X...X
X...X
.XXX.
....X
""",
'r': 
"""
XXXX.
X...X
XXXX.
X..X.
X...X
""",
's': 
"""
.XXXX
X....
.XXX.
....X
XXXX.
""",
't': 
"""
XXXXX
..X..
..X..
..X..
..X..
""",
'u': 
"""
X...X
X...X
X...X
X...X
.XXX.
""",
'v': 
"""
X...X
X...X
.X.X.
.X.X.
..X..
""",
'w': 
"""
X...X
X...X
X.X.X
X.X.X
.X.X.
""",
'x': 
"""
X...X
.X.X.
..X..
.X.X.
X...X
""",
'y': 
"""
X...X
.X.X.
..X..
..X..
..X..
""",
'z': 
"""
XXXXX
....X
..XX.
X....
XXXXX
"""
}

if len(alphabet) != LETTERS:
    print("Cool alphabet bro, but...")
    sys.exit(u)

letters = sorted(alphabet.keys())
print("alphabet: {}".format(letters))

inputs = []
outputs = []

for i in xrange(0, len(letters)):
    letter = letters[i]
    print letter
    pixels = alphabet[letter]
    print pixels.strip()
    pixels = pixels.replace('\n','')
    print pixels
    if len(pixels) != PIXELS:
        print("Cool pixels bro")
        sys.exit(1)

    output = [0.0] * LETTERS
    output[i] = 1.0
    outputs.append(output)

    single_input = []
    for pixel in pixels:
        if pixel == 'X':
            single_input.append(1.0)
        else:
            single_input.append(0.0)

    if len(single_input) != PIXELS:
        print("I don't know what you did son")
        sys.exit(1)
    inputs.append(single_input)

with open(OUTPUT, 'w') as f:
    f.write(json.dumps({"Input": inputs, "Output": outputs}))
