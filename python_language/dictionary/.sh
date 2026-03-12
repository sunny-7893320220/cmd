In the python dictionary are like the map in the go

Dictionary are the built in data values that stores key:values paires

Dictionary are unorders and mutable(changable) and don't allow duplicate keys

syntax 

dic = {
    "name": "saikrishna",
    "age": 22,
    "city": "hyderabad"
}

print(dic)


In the dictionary nesting is posiable {which mean dictionary inside the dictionary}


dic = {
    "name": "saikrishna",
    "age": 22,
    "city": "hyderabad",
    "subject": ["python", "java", "c++"],
    "tuples": ("dict", "set", "list", "tuple"),
    "dict": {
        "name": "krishna",
        "age": 23,
        "city": "hyderabad"
    }
}

print(dic)