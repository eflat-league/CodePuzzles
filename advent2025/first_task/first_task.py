
# position - current position as an int
# direction - L (left) or R (right)
# steps - an int between 0 and 99
def rotate(position, direction, steps):
    # first add 100 so we can modulo after regardless of whether it's left or right
    position = position + 100
    if direction == 'L':
        position = position - steps
    elif direction == 'R':
        position = position + steps

    return position % 100

# reads input file and returns a list of tuples
# each tuple has L or R direction character as first element
# and number of steps to rotate as second element
# - as part of processing the input, we modulo number of steps by 100 since the dial is 0-99
def read_input(filepath):
    lines = []
    try:
        with open("./puzzle_input.txt", 'r') as f:
            lines = [line.strip() for line in f]
    except FileNotFoundError:
        print(f"Error: The title '{filepath}' was not found.")

    rotationSteps = []
    for line in lines:
        direction = line[0]
        if direction not in ['L', 'R']:
            print(f"Error: unrecognized direction {direction}")
            exit(0)

        # input contains numbers > 99, we can just modulo them
        steps = int(line[1:]) % 100

        rotationSteps.append((direction, steps))

    return rotationSteps


def main():
    print("Hello world!")
    rotationSteps = read_input("./puzzle_input.txt")
    # rotate and count number of times we landed on 0
    # starting position is 50
    position = 50
    zeros = 0
    for step in rotationSteps:
        position = rotate(position, step[0], step[1])
        if position == 0:
            zeros += 1

    print(f"Answer is: {zeros}")

if __name__ == "__main__":
    main()