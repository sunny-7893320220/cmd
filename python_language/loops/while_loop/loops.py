# count = 100

# while count >= 1:
#     print(count)
#     count -= 1


# count = int(input("enter the number:"))

# count1 = 1

# while count1 <= 10:
#     print(count * count1)
#     count1 += 1


# list_vaules = [1,4,9,16,25,36,49,64,81,100]

# list_index = len(list_vaules) - 1
# while list_index > -1:
#     print(list_vaules[list_index])
#     list_index -= 1

tuple_list = (1,4,9,16,25,36,49,64,81,100)

tuple_index = 0 

tuple_1 = int(input("enter the number:"))
while tuple_index < len(tuple_list):
    if tuple_list[tuple_index] == tuple_1:
        print(tuple_list[tuple_index])
    tuple_index += 1

