DIRS = "questions/Algorithms/0001.\ Two\ Sum" "questions/Algorithms/0002.\ Add\ Two\ Numbers" "questions/Algorithms/0003.\ Longest\ Substring\ Without\ Repeating\ Characters" "questions/Algorithms/0004.\ Median\ of\ Two\ Sorted\ Arrays" "questions/Algorithms/0006.\ ZigZag\ Conversion" "questions/Algorithms/0007.\ Reverse\ Integer" "questions/Algorithms/0011.\ Container\ With\ Most\ Water" "questions/Algorithms/0012.\ Integer\ to\ Roman" "questions/Algorithms/0046.\ Permutations" "questions/Algorithms/0066.\ Plus\ One" "questions/Algorithms/0292.\ Nim\ Game" "questions/Algorithms/0319.\ Bulb\ Switcher" "questions/Algorithms/0458.\ Poor\ Pigs" "questions/Algorithms/0877.\ Stone\ Game" "questions/Algorithms/0976.\ Largest\ Perimeter\ Triangle"
run: ${DIRS}
${DIRS}:
	make -C $@ golang
