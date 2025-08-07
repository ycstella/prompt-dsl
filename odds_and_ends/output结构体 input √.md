output结构体 input √

outputspec 结构化 √  特殊的json，eg，hint，自定义

expression 

after fix 定义golang函数 √ 还有点粗糙，文件生成的位置还要考虑

if else，参数传入修改√ 

before不好处理，考虑到before是对数据的前处理里，那么得在input，ouput之后，sys

，user之前，并且要先执行，那么是不是要生成两个不同的文件去执行，一个before，一个after和fix

如果我要使他能完全在before块处理数据，那么应该来说，他需要支持用户自定义参数等，这样的话，g4是不是需要修改

顺序应该是执行解析input，output，before（数据处理）      sys，user（prompt）     after，fix（调模型后）

目标
![image-20250716170629410](C:\Users\16\AppData\Roaming\Typora\typora-user-images\image-20250716170629410.png)

现在

![image-20250716170744790](C:\Users\16\AppData\Roaming\Typora\typora-user-images\image-20250716170744790.png)

函数格式化，import构造，package

传入参数应该是模型回复+output，+input




![image-20250717210037602](C:\Users\16\AppData\Roaming\Typora\typora-user-images\image-20250717210037602.png)

![image-20250717210107321](C:\Users\16\AppData\Roaming\Typora\typora-user-images\image-20250717210107321.png)

![image-20250717210153826](C:\Users\16\AppData\Roaming\Typora\typora-user-images\image-20250717210153826.png)