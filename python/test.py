import numpy as np
from keras.datasets import imdb
a = np.arange(18).reshape(2,3,3)
print('原始数组是：')
print(a)
print('\n')
print('迭代输出元素：')
## order 分隔附 C cloumn F
for x in np.nditer(a,order="C",op_flags=['readwrite']):
    print(x, end=", ")
    x[...]= x * 2
print('\n 修改后的元素')
print(a)
print('\n')
(train_data, train_labels), (test_data, test_labels) = imdb.load_data(
 num_words=10000)
print(train_data.shape,train_labels.shape,test_data.shape,test_labels.shape)
print(train_data[0])
print(train_labels[0])
word_index = imdb.get_word_index() 
reverse_word_index = dict(
 [(value, key) for (key, value) in word_index.items()])
decoded_review = ' '.join(
 [reverse_word_index.get(i - 3, '?') for i in train_data[0]])

# (train_images, train_labels), (test_images, test_labels) = mnist.load_data()
# train_images = train_images.reshape((60000, 28 * 28))
# train_images = train_images.astype('float32') / 255
# test_images = test_images.reshape((10000, 28 * 28))
# test_images = test_images.astype('float32') / 255
