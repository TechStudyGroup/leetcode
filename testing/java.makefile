DIRS = "questions/Algorithms/0001.\ Two\ Sum" "questions/Algorithms/0002.\ Add\ Two\ Numbers" "questions/Algorithms/0003.\ Longest\ Substring\ Without\ Repeating\ Characters" "questions/Algorithms/0020.\ Valid\ Parentheses" "questions/Algorithms/0043.\ Multiply\ Strings" "questions/Algorithms/0050.\ Pow(x,\ n)" "questions/Algorithms/0097.\ Interleaving\ String" "questions/Algorithms/0204.\ Count\ Primes" "questions/Algorithms/0292.\ Nim\ Game" "questions/Algorithms/0458.\ Poor\ Pigs" "questions/Algorithms/0514.\ Freedom\ Trail" "questions/Algorithms/0877.\ Stone\ Game"
run: ${DIRS}
${DIRS}:
	make -C $@ java
