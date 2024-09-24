with open("input_data.txt", 'w') as f:
    data = " ".join(str(x) for x in [1, 2, 3] * 1000)
    f.write(data)
