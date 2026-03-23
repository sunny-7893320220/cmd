def len_of_list(list_values):
    print(len(list_values))

# lenth = [1,2,3,4,5,6,7,8,9,10]

# lenth = ["sai", "krishna", "jamanjyothi"]

def lenth_of_list(list_values):
    for i in list_values:
        print(i, end= " ")
    
# lenth_of_list(lenth)

def finding_factorial(num):
    factorial = 1
    for i in range(1, num+1):
        factorial *= i
    print(factorial)

# finding_factorial(int(input("Enter the number: ")))


# writting a function to convert from USDC to INR

def usdc_to_inr(usd):
    inr = usd * 100
    print(inr)

usdc_to_inr(int(input("Enter the amount in USD: ")))