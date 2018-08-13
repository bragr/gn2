import json
import random
rnd = random.SystemRandom()

OUTPUT="invert.json"
SIZE=2500

inputs = []
outputs = []

for _ in range(SIZE):
    x = rnd.random()
    inputs.append([x])
    outputs.append([1-x])

with open(OUTPUT, 'w') as f:
    f.write(json.dumps({"Input": inputs, "Output": outputs}))
