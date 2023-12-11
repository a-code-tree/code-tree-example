import numpy as np

a = np.arange(18).reshape(2, 3,3)
print('原始数组是：')
print(a)
print('\n')
print('迭代输出元素：')
## order 分隔附 C cloumn F
for x in np.nditer(a,order="c",op_flags=['readwrite']):
    print(x, end=", ")
    x[...]= x * 2
print('\n 修改后的元素')
print(a)
print('\n')
