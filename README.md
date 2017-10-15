
设计说明
selpg可以从输入中选取某些页输出，具体请看以下链接
https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html

使用selpg和测试结果

这里有两个测试文件，test1.txt和test2.txt，test1.txt是页行数固定的文本（为了简单，这里一页两行），text2.txt是有ascii换页字符定界的文本。

$ selpg -s 1 -e 1 input_file
该命令把输入文件的第一页输出到标准输出流中

测试结果如下：
ljx@ljx-X550JX:~/go$ selpg -s 1 -e 1 test1.txt
1
2

$ selpg -s 1 -e 1 < input_file
该命令与示例 1 所做的工作相同，但在本例中，selpg 读取标准输入，而标准输入已被 shell／内核重定向为来自“input_file”
而不是显式命名的文件名参数。输入的第 1 页被写至屏幕。

测试结果如下：
ljx@ljx-X550JX:~/go$ selpg -s 1 -e 1 < test1.txt
1
2

$ other_command | selpg -s 1 -e 5
“other_command”的标准输出被 shell／内核重定向至 selpg 的标准输入。将第 1 页到第 5 页写至 selpg 的标准输出（屏幕）。

测试结果如下：
ljx@ljx-X550JX:~/go$ cat test1.txt | selpg -s 1 -e 5
1
2
3
4
5
6
7
8
9
10

$ selpg -s 1 -e 5 input_file > output_file
selpg 将第 1 页到第 5 页写至标准输出；标准输出被 shell／内核重定向至“output_file”。

测试结果如下：
ljx@ljx-X550JX:~/go$ selpg -s 1 -e 5 test1.txt > output1.txt
ljx@ljx-X550JX:~/go$ cat output1.txt
1
2
3
4
5
6
7
8
9
10

$ selpg -s 1 -e 2 input_file 2> error_file
selpg 将第 1 页到第 2 页写至标准输出（屏幕）；所有的错误消息被 shell／内核重定向至“error_file”。

测试结果如下：
ljx@ljx-X550JX:~/go$ selpg test1.txt 2> error1.txt
ljx@ljx-X550JX:~/go$ cat error1.txt
2017/10/15 20:23:27 selpg: not enough arguments
Usage of selpg:

$ selpg -s 1 -e 2 input_file > output_file 2> error_file
selpg 将第 1 页到第 2 页写至标准输出；标准输出被 shell／内核重定向至“output_file”，
所有的错误消息被 shell／内核重定向至“error_file”。

测试结果如下：
ljx@ljx-X550JX:~/go$ selpg -s 1 -e 2 test1.txt > output2.txt 2> error2.txt
ljx@ljx-X550JX:~/go$ cat output2.txt
1
2
3
4

$ selpg -s 1 -e 2 input_file > output_file 2> /dev/null
selpg 将第 1 页到第 2 页写至标准输出；标准输出被 shell／内核重定向至“output_file”，
所有的错误消息被抛弃

测试结果如下：
ljx@ljx-X550JX:~/go$ selpg -s 1 -e 2 test1.txt > output3.txt 2> /dev/null
ljx@ljx-X550JX:~/go$ cat output3.txt
1
2
3
4

$ selpg -s 1 -e 2 input_file > /dev/null
selpg 将第 1 页到第 2 页写至标准输出，标准输出被丢弃；错误消息在屏幕出现。

$ selpg -s 1 -e 2 input_file | other_command
selpg 的标准输出透明地被 shell／内核重定向，成为“other_command”的标准输入，第 1 页到第 2 页被写至该标准输入。

测试结果如下：
ljx@ljx-X550JX:~/go$ selpg -s 1 -e 2 test1.txt | grep 3
3

$ selpg -s 1 -e 2 input_file 2> error_file | other_command
和上个例子差不多，只是把错误信息重定向到error_file中

$ selpg -s 1 -e 2 -l 5 input_file
-l选项把每页的行数改为5

测试结果如下：
ljx@ljx-X550JX:~/go$ selpg -s 1 -e 2 -l 5 test1.txt
1
2
3
4
5
6
7
8
9
10

$ selpg -s 1 -e 5 -f input_file
-f选项 有换页符定界

测试结果如下：
ljx@ljx-X550JX:~/go$ selpg -s 1 -e 5 -f test2.txt
1
 2
  3
   4
    5

selpg -s 1 -e 2 -d lp1 input_file
第 1 页到第 2 页由管道输送至命令“lp -dlp1”，该命令将使输出在打印机 lp1 上打印。

$ selpg -s10 -e20 input_file > output_file 2>error_file &
在后台运行
