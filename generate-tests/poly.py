import json
import random
rnd = random.SystemRandom()

OUTPUT="poly.json"
SIZE=5000
SCALING=1.0

inputs = []
outputs = []

for _ in range(SIZE):
    x = rnd.random()*SCALING
    y = rnd.random()*SCALING
    z = x * y**2 + y**x - 2 * (x*y)
    inputs.append([x, y])
    outputs.append([z])

with open(OUTPUT, 'w') as f:
    f.write(json.dumps({"Input": inputs, "Output": outputs}))
