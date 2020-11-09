# LeetCode

LeetCode Solutions.

### Solutions file rules

- Each problem is organized according to the title number.
- Each problem folder has a description file which one named `readme.md`.
- There can be multiple solutions for the same language in each folder, ranked according to the degree of algorithm optimization, and the file name prefix is the language name. For example: the performance of `c1.c` is better than the performance of `c2.c`.
- Each problem solution must contain a `main` function that contains all the test cases that can be thought of, making it easy to run offline. The solution function is written separately as a separate function.

### How to test problem's solution offline

- Entering a language environment requires only `make $language`. For example, the environment that enters Python is `make python`. It may take some time to enter a certain environment for the first time. The `Docker` image will be compiled first.
- All files in this directory will be mapped to the `/app` directory of `Container`.
- The environment of `Container` may be `Debian` or `Alpine`.

### Solutions' list

|Problem|C|C++|Go|Java|JS|PHP|Python|Rust|SQL|Bash|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
|[0001](https://leetcode.com/problems/two-sum)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0001.%20Two%20Sum/c1.c)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0001.%20Two%20Sum/cpp1.cc)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0001.%20Two%20Sum/go1/go1.go)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0001.%20Two%20Sum/java1.java)|||[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0001.%20Two%20Sum/python1.py)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0001.%20Two%20Sum/rust1.rs)|-|-|
|[0002](https://leetcode.com/problems/add-two-numbers)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0002.%20Add%20Two%20Numbers/c1.c)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0002.%20Add%20Two%20Numbers/cpp1.cc)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0002.%20Add%20Two%20Numbers/go1/go1.go)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0002.%20Add%20Two%20Numbers/java1.java)||[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0002.%20Add%20Two%20Numbers/php1.php)|||-|-|
|[0003](https://leetcode.com/problems/longest-substring-without-repeating-characters)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0003.%20Longest%20Substring%20Without%20Repeating%20Characters/c1.c)||[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0003.%20Longest%20Substring%20Without%20Repeating%20Characters/go1/go1.go) [2](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0003.%20Longest%20Substring%20Without%20Repeating%20Characters/go2/go2.go)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0003.%20Longest%20Substring%20Without%20Repeating%20Characters/java1.java) [2](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0003.%20Longest%20Substring%20Without%20Repeating%20Characters/java2.java) [3](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0003.%20Longest%20Substring%20Without%20Repeating%20Characters/java3.java) [4](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0003.%20Longest%20Substring%20Without%20Repeating%20Characters/java4.java)|||||-|-|
|[0004](https://leetcode.com/problems/median-of-two-sorted-arrays)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0004.%20Median%20of%20Two%20Sorted%20Arrays/c1.c)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0004.%20Median%20of%20Two%20Sorted%20Arrays/cpp1.cc)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0004.%20Median%20of%20Two%20Sorted%20Arrays/go1/go1.go) [2](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0004.%20Median%20of%20Two%20Sorted%20Arrays/go2/go2.go)|||[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0004.%20Median%20of%20Two%20Sorted%20Arrays/php1.php) [2](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0004.%20Median%20of%20Two%20Sorted%20Arrays/php2.php)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0004.%20Median%20of%20Two%20Sorted%20Arrays/python1.py) [2](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0004.%20Median%20of%20Two%20Sorted%20Arrays/python2.py)||-|-|
|[0006](https://leetcode.com/problems/zigzag-conversion)|||[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0006.%20ZigZag%20Conversion/go1/go1.go)|||[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0006.%20ZigZag%20Conversion/php1.php)|||-|-|
|[0007](https://leetcode.com/problems/reverse-integer)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0007.%20Reverse%20Integer/c1.c)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0007.%20Reverse%20Integer/cpp1.cc)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0007.%20Reverse%20Integer/go1/go1.go)|||||[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0007.%20Reverse%20Integer/rust1.rs)|-|-|
|[0011](https://leetcode.com/problems/container-with-most-water)|||[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0011.%20Container%20With%20Most%20Water/go1/go1.go)||||||-|-|
|[0012](https://leetcode.com/problems/integer-to-roman)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0012.%20Integer%20to%20Roman/c1.c)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0012.%20Integer%20to%20Roman/cpp1.cc)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0012.%20Integer%20to%20Roman/go1/go1.go)||[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0012.%20Integer%20to%20Roman/js1.js)|||[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0012.%20Integer%20to%20Roman/rust1.rs)|-|-|
|[0013](https://leetcode.com/problems/roman-to-integer)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0013.%20Roman%20to%20Integer/c1.c)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0013.%20Roman%20to%20Integer/cpp1.cc)|||||||-|-|
|[0021](https://leetcode.com/problems/merge-two-sorted-lists)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0021.%20Merge%20Two%20Sorted%20Lists/c1.c)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0021.%20Merge%20Two%20Sorted%20Lists/cpp1.cc)|||||||-|-|
|[0037](https://leetcode.com/problems/sudoku-solver)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0037.%20Sudoku%20Solver/c1.c)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0037.%20Sudoku%20Solver/cpp1.cc)||||||[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0037.%20Sudoku%20Solver/rust1.rs)|-|-|
|[0048](https://leetcode.com/problems/rotate-image)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0048.%20Rotate%20Image/c1.c)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0048.%20Rotate%20Image/cpp1.cc)|||||||-|-|
|[0050](https://leetcode.com/problems/powx-n)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0050.%20Pow(x,%20n)/c1.c)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0050.%20Pow(x,%20n)/cpp1.cc)||[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0050.%20Pow(x,%20n)/java1.java)|||||-|-|
|[0061](https://leetcode.com/problems/rotate-list)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0061.%20Rotate%20List/c1.c)||||||||-|-|
|[0066](https://leetcode.com/problems/plus-one)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0066.%20Plus%20One/c1.c)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0066.%20Plus%20One/cpp1.cc)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0066.%20Plus%20One/go1/go1.go)||||[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0066.%20Plus%20One/python1.py)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0066.%20Plus%20One/rust1.rs)|-|-|
|[0067](https://leetcode.com/problems/add-binary)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0067.%20Add%20Binary/c1.c)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0067.%20Add%20Binary/cpp1.cc)|||||||-|-|
|[0069](https://leetcode.com/problems/sqrtx)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0069.%20Sqrt(x)/c1.c)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0069.%20Sqrt(x)/cpp1.cc)|||||||-|-|
|[0070](https://leetcode.com/problems/climbing-stairs)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0070.%20Climbing%20Stairs/c1.c) [2](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0070.%20Climbing%20Stairs/c2.c)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0070.%20Climbing%20Stairs/cpp1.cc) [2](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0070.%20Climbing%20Stairs/cpp2.cc)||||||[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0070.%20Climbing%20Stairs/rust1.rs)|-|-|
|[0089](https://leetcode.com/problems/gray-code)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0089.%20Gray%20Code/c1.c)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0089.%20Gray%20Code/cpp1.cc)|||||||-|-|
|[0100](https://leetcode.com/problems/same-tree)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0100.%20Same%20Tree/c1.c)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0100.%20Same%20Tree/cpp1.cc)|||||||-|-|
|[0101](https://leetcode.com/problems/symmetric-tree)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0101.%20Symmetric%20Tree/c1.c)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0101.%20Symmetric%20Tree/cpp1.cc)|||||||-|-|
|[0121](https://leetcode.com/problems/best-time-to-buy-and-sell-stock)||[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0121.%20Best%20Time%20to%20Buy%20and%20Sell%20Stock/cpp1.cc)|||||[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0121.%20Best%20Time%20to%20Buy%20and%20Sell%20Stock/python1.py)||-|-|
|[0136](https://leetcode.com/problems/single-number)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0136.%20Single%20Number/c1.c)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0136.%20Single%20Number/cpp1.cc)|||||||-|-|
|[0169](https://leetcode.com/problems/majority-element)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0169.%20Majority%20Element/c1.c) [2](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0169.%20Majority%20Element/c2.c)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0169.%20Majority%20Element/cpp1.cc)||||||[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0169.%20Majority%20Element/rust1.rs)|-|-|
|[0192](https://leetcode.com/problems/word-frequency)|-|-|-|-|-|-|-|-|-|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Shell/0192.%20Word%20Frequency/bash1.sh)|
|[0193](https://leetcode.com/problems/valid-phone-numbers)|-|-|-|-|-|-|-|-|-|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Shell/0193.%20Valid%20Phone%20Numbers/bash1.sh) [2](https://github.com/6leetcode/6leetcode/blob/master/questions/Shell/0193.%20Valid%20Phone%20Numbers/bash2.sh)|
|[0194](https://leetcode.com/problems/transpose-file)|-|-|-|-|-|-|-|-|-|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Shell/0194.%20Transpose%20File/bash1.sh)|
|[0195](https://leetcode.com/problems/tenth-line)|-|-|-|-|-|-|-|-|-|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Shell/0195.%20Tenth%20Line/bash1.sh) [2](https://github.com/6leetcode/6leetcode/blob/master/questions/Shell/0195.%20Tenth%20Line/bash2.sh)|
|[0226](https://leetcode.com/problems/invert-binary-tree)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0226.%20Invert%20Binary%20Tree/c1.c)||||||||-|-|
|[0231](https://leetcode.com/problems/power-of-two)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0231.%20Power%20of%20Two/c1.c) [2](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0231.%20Power%20of%20Two/c2.c)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0231.%20Power%20of%20Two/cpp1.cc) [2](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0231.%20Power%20of%20Two/cpp2.cc)|||||||-|-|
|[0233](https://leetcode.com/problems/number-of-digit-one)|||||||[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0233.%20Number%20of%20Digit%20One/python1.py)||-|-|
|[0283](https://leetcode.com/problems/move-zeroes)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0283.%20Move%20Zeroes/c1.c)||||||||-|-|
|[0319](https://leetcode.com/problems/bulb-switcher)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0319.%20Bulb%20Switcher/c1.c)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0319.%20Bulb%20Switcher/cpp1.cc)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0319.%20Bulb%20Switcher/go1/go1.go)||[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0319.%20Bulb%20Switcher/js1.js)||[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0319.%20Bulb%20Switcher/python1.py)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0319.%20Bulb%20Switcher/rust1.rs)|-|-|
|[0326](https://leetcode.com/problems/power-of-three)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0326.%20Power%20of%20Three/c1.c)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0326.%20Power%20of%20Three/cpp1.cc)|||||||-|-|
|[0344](https://leetcode.com/problems/reverse-string)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0344.%20Reverse%20String/c1.c)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0344.%20Reverse%20String/cpp1.cc)|||||[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0344.%20Reverse%20String/python1.py)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0344.%20Reverse%20String/rust1.rs)|-|-|
|[0458](https://leetcode.com/problems/poor-pigs)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0458.%20Poor%20Pigs/c1.c)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0458.%20Poor%20Pigs/cpp1.cc)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0458.%20Poor%20Pigs/go1/go1.go)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0458.%20Poor%20Pigs/java1.java)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0458.%20Poor%20Pigs/js1.js)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0458.%20Poor%20Pigs/php1.php)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0458.%20Poor%20Pigs/python1.py)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0458.%20Poor%20Pigs/rust1.rs)|-|-|
|[0498](https://leetcode.com/problems/diagonal-traverse)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0498.%20Diagonal%20Traverse/c1.c)||||||||-|-|
|[0679](https://leetcode.com/problems/24-game)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0679.%2024%20Game/c1.c)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0679.%2024%20Game/cpp1.cc)||||||[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0679.%2024%20Game/rust1.rs)|-|-|
|[0696](https://leetcode.com/problems/count-binary-substrings)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0696.%20Count%20Binary%20Substrings/c1.c)|[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0696.%20Count%20Binary%20Substrings/cpp1.cc)||||||[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0696.%20Count%20Binary%20Substrings/rust1.rs)|-|-|
|[0973](https://leetcode.com/problems/k-closest-points-to-origin)||[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/0973.%20K%20Closest%20Points%20to%20Origin/cpp1.cc)|||||||-|-|
|[1100](https://leetcode.com/problems/find-k-length-substrings-with-no-repeated-characters)||[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/1100.%20Find%20K-Length%20Substrings%20With%20No%20Repeated%20Characters/cpp1.cc)|||||||-|-|
|[1106](https://leetcode.com/problems/parsing-a-boolean-expression)||[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/1106.%20Parsing%20A%20Boolean%20Expression/cpp1.cc)|||||||-|-|
|[1165](https://leetcode.com/problems/single-row-keyboard)||[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/1165.%20Single-Row%20Keyboard/cpp1.cc)|||||||-|-|
|[1309](https://leetcode.com/problems/decrypt-string-from-alphabet-to-integer-mapping)||[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/1309.%20Decrypt%20String%20from%20Alphabet%20to%20Integer%20Mapping/cpp1.cc)|||||||-|-|
|[1358](https://leetcode.com/problems/number-of-substrings-containing-all-three-characters)||[1](https://github.com/6leetcode/6leetcode/blob/master/questions/Algorithms/1358.%20Number%20of%20Substrings%20Containing%20All%20Three%20Characters/cpp1.cc)|||||||-|-|