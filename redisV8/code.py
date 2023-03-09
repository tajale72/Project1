import random

def sort(values):
    length = len(values)
    for time in range(0,length-1):
        for position in range(0, length-time-1):
            if values[position] > values[position+1]:
                temp = values[position]
                values[position] = values[position+1]
                values[position+1] = temp



numbers = []
number = 0
while number < 10:
    value = random.randint(1,100)
    if not (value in numbers):
        numbers.append(value)
        number = number + number + 1


print("before", numbers)

sort(numbers)

print("After", numbers)