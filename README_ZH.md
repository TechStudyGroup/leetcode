# LeetCode

LeetCode 题解

## 文件规则

- 每个题目按照题目序号编排。
- 每个题目文件夹下有相关的题目说明，文件名为 `readme.md`。
- 每个文件夹中同一种语言可以有多重解决方案，按照算法优化程度排名，文件名前缀为语言名字。例如：`c1.c` 的性能是比 `c2.c` 的性能更优的。
- 每个题目里边都要含有 `main` 函数，含有可以想到的所有测试案例，方便线下运行。题解函数，单独写成一个独立的函数。

## 本地测试

- 进入某种语言的环境仅需要 `make $language`，例如进入 python 的环境就是 `make python`，第一次进入某一种环境可能会需要一些时间，首先会编译这个 `Docker` 镜像。
- 本目录的所有文件将会映射到 `Container` 的 `/app` 目录下。
- `Container` 的环境可能是 `Debian` 或者 `Alpine`。

## Solutions' list

<details>

  <summary>Algorithms</summary>

|Index|Difficulty|C|C++|Go|Java|TypeScript|PHP|Python|Rust|SQL|Bash|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
|[0001](https://leetcode.cn/problems/two-sum)|Easy|:white_check_mark:|:white_check_mark:|:white_check_mark:|:white_check_mark:|:white_check_mark:|:white_check_mark:|:white_check_mark:|:white_check_mark:|||
|[0002](https://leetcode.cn/problems/add-two-numbers)|Medium|:white_check_mark:|:white_check_mark:|:white_check_mark:|:white_check_mark:||:white_check_mark:|||||
|[0003](https://leetcode.cn/problems/longest-substring-without-repeating-characters)|Medium|:white_check_mark:||:white_check_mark:|:white_check_mark:|||||||
|[0004](https://leetcode.cn/problems/median-of-two-sorted-arrays)|Hard|:white_check_mark:|:white_check_mark:|:white_check_mark:|||:white_check_mark:|:white_check_mark:||||
|[0005](https://leetcode.cn/problems/longest-palindromic-substring)|Medium|:white_check_mark:||||||||||
|[0006](https://leetcode.cn/problems/zigzag-conversion)|Medium|:white_check_mark:||:white_check_mark:||||||||
|[0007](https://leetcode.cn/problems/reverse-integer)|Medium|:white_check_mark:|:white_check_mark:|:white_check_mark:|||||:white_check_mark:|||
|[0008](https://leetcode.cn/problems/string-to-integer-atoi)|Medium|:white_check_mark:||||||||||
|[0009](https://leetcode.cn/problems/palindrome-number)|Easy|:white_check_mark:||||||||||
|[0010](https://leetcode.cn/problems/regular-expression-matching)|Hard|:white_check_mark:||||||||||
|[0011](https://leetcode.cn/problems/container-with-most-water)|Medium|:white_check_mark:||:white_check_mark:||||||||
|[0012](https://leetcode.cn/problems/integer-to-roman)|Medium|:white_check_mark:|:white_check_mark:|:white_check_mark:|||||:white_check_mark:|||
|[0013](https://leetcode.cn/problems/roman-to-integer)|Easy|:white_check_mark:|:white_check_mark:|||||||||
|[0014](https://leetcode.cn/problems/longest-common-prefix)|Easy|:white_check_mark:||||||||||
|[0015](https://leetcode.cn/problems/3sum)|Medium|:white_check_mark:|:white_check_mark:|:white_check_mark:||||||||
|[0016](https://leetcode.cn/problems/3sum-closest)|Medium|:white_check_mark:|:white_check_mark:|||||||||
|[0017](https://leetcode.cn/problems/letter-combinations-of-a-phone-number)|Medium|:white_check_mark:||||||||||
|[0018](https://leetcode.cn/problems/4sum)|Medium|:white_check_mark:|:white_check_mark:|||||||||
|[0019](https://leetcode.cn/problems/remove-nth-node-from-end-of-list)|Medium|:white_check_mark:||||||||||
|[0020](https://leetcode.cn/problems/valid-parentheses)|Easy|:white_check_mark:|:white_check_mark:||:white_check_mark:|||:white_check_mark:||||
|[0021](https://leetcode.cn/problems/merge-two-sorted-lists)|Easy|:white_check_mark:|:white_check_mark:|||||||||
|[0022](https://leetcode.cn/problems/generate-parentheses)|Medium||:white_check_mark:|||||||||
|[0023](https://leetcode.cn/problems/merge-k-sorted-lists)|Hard||:white_check_mark:|||||||||
|[0024](https://leetcode.cn/problems/swap-nodes-in-pairs)|Medium||:white_check_mark:|||||||||
|[0025](https://leetcode.cn/problems/reverse-nodes-in-k-group)|Hard||:white_check_mark:|||||||||
|[0026](https://leetcode.cn/problems/remove-duplicates-from-sorted-array)|Easy||:white_check_mark:|||||||||
|[0027](https://leetcode.cn/problems/remove-element)|Easy||:white_check_mark:|||||||||
|[0028](https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string)|Easy|||||||||||
|[0029](https://leetcode.cn/problems/divide-two-integers)|Medium|||||||||||
|[0030](https://leetcode.cn/problems/substring-with-concatenation-of-all-words)|Hard|||||||||||
|[0031](https://leetcode.cn/problems/next-permutation)|Medium|:white_check_mark:|:white_check_mark:|||||||||
|[0032](https://leetcode.cn/problems/longest-valid-parentheses)|Hard|||||||||||
|[0033](https://leetcode.cn/problems/search-in-rotated-sorted-array)|Medium|||||||||||
|[0034](https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array)|Medium|||||||||||
|[0035](https://leetcode.cn/problems/search-insert-position)|Easy|||||||||||
|[0036](https://leetcode.cn/problems/valid-sudoku)|Medium|:white_check_mark:||||||||||
|[0037](https://leetcode.cn/problems/sudoku-solver)|Hard|:white_check_mark:|:white_check_mark:||||||:white_check_mark:|||
|[0038](https://leetcode.cn/problems/count-and-say)|Medium||:white_check_mark:|||||||||
|[0039](https://leetcode.cn/problems/combination-sum)|Medium|||||||||||
|[0040](https://leetcode.cn/problems/combination-sum-ii)|Medium|||||||||||
|[0041](https://leetcode.cn/problems/first-missing-positive)|Hard|||||||||||
|[0042](https://leetcode.cn/problems/trapping-rain-water)|Hard|||||||||||
|[0043](https://leetcode.cn/problems/multiply-strings)|Medium|:white_check_mark:|:white_check_mark:||:white_check_mark:|||||||
|[0044](https://leetcode.cn/problems/wildcard-matching)|Hard|||||||||||
|[0045](https://leetcode.cn/problems/jump-game-ii)|Medium|||||||||||
|[0046](https://leetcode.cn/problems/permutations)|Medium|:white_check_mark:|:white_check_mark:|:white_check_mark:||||||||
|[0047](https://leetcode.cn/problems/permutations-ii)|Medium||:white_check_mark:|||||||||
|[0048](https://leetcode.cn/problems/rotate-image)|Medium|:white_check_mark:|:white_check_mark:|||||||||
|[0049](https://leetcode.cn/problems/group-anagrams)|Medium|||||||||||
|[0050](https://leetcode.cn/problems/powx-n)|Medium|:white_check_mark:|:white_check_mark:||:white_check_mark:|||||||
|[0051](https://leetcode.cn/problems/n-queens)|Hard|||||||||||
|[0052](https://leetcode.cn/problems/n-queens-ii)|Hard|||||||||||
|[0053](https://leetcode.cn/problems/maximum-subarray)|Medium|:white_check_mark:|:white_check_mark:|||||||||
|[0054](https://leetcode.cn/problems/spiral-matrix)|Medium||:white_check_mark:|||||||||
|[0055](https://leetcode.cn/problems/jump-game)|Medium||:white_check_mark:|||||||||
|[0056](https://leetcode.cn/problems/merge-intervals)|Medium||:white_check_mark:|||||||||
|[0057](https://leetcode.cn/problems/insert-interval)|Medium|||||||||||
|[0058](https://leetcode.cn/problems/length-of-last-word)|Easy||:white_check_mark:|||||||||
|[0059](https://leetcode.cn/problems/spiral-matrix-ii)|Medium|||||||||||
|[0060](https://leetcode.cn/problems/permutation-sequence)|Hard|||||||||||
|[0061](https://leetcode.cn/problems/rotate-list)|Medium|:white_check_mark:||||||||||
|[0062](https://leetcode.cn/problems/unique-paths)|Medium|||||||||||
|[0063](https://leetcode.cn/problems/unique-paths-ii)|Medium|||||||||||
|[0064](https://leetcode.cn/problems/minimum-path-sum)|Medium|||||||||||
|[0065](https://leetcode.cn/problems/valid-number)|Hard|||||||||||
|[0066](https://leetcode.cn/problems/plus-one)|Easy|:white_check_mark:|:white_check_mark:|:white_check_mark:||||:white_check_mark:|:white_check_mark:|||
|[0067](https://leetcode.cn/problems/add-binary)|Easy|:white_check_mark:|:white_check_mark:|||||||||
|[0068](https://leetcode.cn/problems/text-justification)|Hard|||||||||||
|[0069](https://leetcode.cn/problems/sqrtx)|Easy|:white_check_mark:|:white_check_mark:|||||||||
|[0070](https://leetcode.cn/problems/climbing-stairs)|Easy|:white_check_mark:|:white_check_mark:||||||:white_check_mark:|||
|[0071](https://leetcode.cn/problems/simplify-path)|Medium|||||||||||
|[0072](https://leetcode.cn/problems/edit-distance)|Medium|||||||||||
|[0073](https://leetcode.cn/problems/set-matrix-zeroes)|Medium|||||||||||
|[0074](https://leetcode.cn/problems/search-a-2d-matrix)|Medium||:white_check_mark:|||||||||
|[0075](https://leetcode.cn/problems/sort-colors)|Medium||:white_check_mark:|||||||||
|[0076](https://leetcode.cn/problems/minimum-window-substring)|Hard|||||||||||
|[0077](https://leetcode.cn/problems/combinations)|Medium|||||||||||
|[0078](https://leetcode.cn/problems/subsets)|Medium|||||||||||
|[0079](https://leetcode.cn/problems/word-search)|Medium|||||||||||
|[0080](https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii)|Medium|||||||||||
|[0081](https://leetcode.cn/problems/search-in-rotated-sorted-array-ii)|Medium|||||||||||
|[0082](https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii)|Medium|||||||||||
|[0083](https://leetcode.cn/problems/remove-duplicates-from-sorted-list)|Easy|||||||||||
|[0084](https://leetcode.cn/problems/largest-rectangle-in-histogram)|Hard|||||||||||
|[0085](https://leetcode.cn/problems/maximal-rectangle)|Hard|||||||||||
|[0086](https://leetcode.cn/problems/partition-list)|Medium|||||||||||
|[0087](https://leetcode.cn/problems/scramble-string)|Hard|||||||||||
|[0088](https://leetcode.cn/problems/merge-sorted-array)|Easy|||||||||||
|[0089](https://leetcode.cn/problems/gray-code)|Medium|:white_check_mark:|:white_check_mark:|||||||||
|[0090](https://leetcode.cn/problems/subsets-ii)|Medium|||||||||||
|[0091](https://leetcode.cn/problems/decode-ways)|Medium||:white_check_mark:|||||||||
|[0092](https://leetcode.cn/problems/reverse-linked-list-ii)|Medium||:white_check_mark:|||||||||
|[0093](https://leetcode.cn/problems/restore-ip-addresses)|Medium||:white_check_mark:|||||||||
|[0094](https://leetcode.cn/problems/binary-tree-inorder-traversal)|Easy||:white_check_mark:|||||||||
|[0095](https://leetcode.cn/problems/unique-binary-search-trees-ii)|Medium|||||||||||
|[0096](https://leetcode.cn/problems/unique-binary-search-trees)|Medium|||||||||||
|[0097](https://leetcode.cn/problems/interleaving-string)|Medium|:white_check_mark:|:white_check_mark:|:white_check_mark:|:white_check_mark:|||||||
|[0098](https://leetcode.cn/problems/validate-binary-search-tree)|Medium|||||||||||
|[0099](https://leetcode.cn/problems/recover-binary-search-tree)|Medium|||||||||||
|[0100](https://leetcode.cn/problems/same-tree)|Easy|:white_check_mark:|:white_check_mark:|||||||||
|[0101](https://leetcode.cn/problems/symmetric-tree)|Easy|:white_check_mark:|:white_check_mark:|||||||||
|[0102](https://leetcode.cn/problems/binary-tree-level-order-traversal)|Medium|||||||||||
|[0103](https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal)|Medium|||||||||||
|[0104](https://leetcode.cn/problems/maximum-depth-of-binary-tree)|Easy|||||||||||
|[0105](https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal)|Medium|||||||||||
|[0106](https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal)|Medium|||||||||||
|[0107](https://leetcode.cn/problems/binary-tree-level-order-traversal-ii)|Medium|||||||||||
|[0108](https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree)|Easy|||||||||||
|[0109](https://leetcode.cn/problems/convert-sorted-list-to-binary-search-tree)|Medium|||||||||||
|[0110](https://leetcode.cn/problems/balanced-binary-tree)|Easy|||||||||||
|[0111](https://leetcode.cn/problems/minimum-depth-of-binary-tree)|Easy|||||||||||
|[0112](https://leetcode.cn/problems/path-sum)|Easy|||||||||||
|[0113](https://leetcode.cn/problems/path-sum-ii)|Medium|||||||||||
|[0114](https://leetcode.cn/problems/flatten-binary-tree-to-linked-list)|Medium|||||||||||
|[0115](https://leetcode.cn/problems/distinct-subsequences)|Hard|||||||||||
|[0116](https://leetcode.cn/problems/populating-next-right-pointers-in-each-node)|Medium|||||||||||
|[0117](https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii)|Medium|||||||||||
|[0118](https://leetcode.cn/problems/pascals-triangle)|Easy||:white_check_mark:|||||||||
|[0119](https://leetcode.cn/problems/pascals-triangle-ii)|Easy|||||||||||
|[0120](https://leetcode.cn/problems/triangle)|Medium|||||||||||
|[0121](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock)|Easy||:white_check_mark:|||||:white_check_mark:||||
|[0122](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii)|Medium|||||||||||
|[0123](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii)|Hard|||||||||||
|[0124](https://leetcode.cn/problems/binary-tree-maximum-path-sum)|Hard|||||||||||
|[0125](https://leetcode.cn/problems/valid-palindrome)|Easy|||||||||||
|[0126](https://leetcode.cn/problems/word-ladder-ii)|Hard|||||||||||
|[0127](https://leetcode.cn/problems/word-ladder)|Hard|||||||||||
|[0128](https://leetcode.cn/problems/longest-consecutive-sequence)|Medium|||||||||||
|[0129](https://leetcode.cn/problems/sum-root-to-leaf-numbers)|Medium|||||||||||
|[0130](https://leetcode.cn/problems/surrounded-regions)|Medium|||||||||||
|[0131](https://leetcode.cn/problems/palindrome-partitioning)|Medium|||||||||||
|[0132](https://leetcode.cn/problems/palindrome-partitioning-ii)|Hard|||||||||||
|[0133](https://leetcode.cn/problems/clone-graph)|Medium|||||||||||
|[0134](https://leetcode.cn/problems/gas-station)|Medium|||||||||||
|[0135](https://leetcode.cn/problems/candy)|Hard|||||||||||
|[0136](https://leetcode.cn/problems/single-number)|Easy|:white_check_mark:|:white_check_mark:|||||||||
|[0137](https://leetcode.cn/problems/single-number-ii)|Medium|:white_check_mark:|:white_check_mark:|||||||||
|[0138](https://leetcode.cn/problems/copy-list-with-random-pointer)|Medium|||||||||||
|[0139](https://leetcode.cn/problems/word-break)|Medium|||||||||||
|[0140](https://leetcode.cn/problems/word-break-ii)|Hard|||||||||||
|[0141](https://leetcode.cn/problems/linked-list-cycle)|Easy|||||||||||
|[0142](https://leetcode.cn/problems/linked-list-cycle-ii)|Medium|||||||||||
|[0143](https://leetcode.cn/problems/reorder-list)|Medium|||||||||||
|[0144](https://leetcode.cn/problems/binary-tree-preorder-traversal)|Easy|||||||||||
|[0145](https://leetcode.cn/problems/binary-tree-postorder-traversal)|Easy|||||||||||
|[0146](https://leetcode.cn/problems/lru-cache)|Medium|||||||||||
|[0147](https://leetcode.cn/problems/insertion-sort-list)|Medium|||||||||||
|[0148](https://leetcode.cn/problems/sort-list)|Medium|||||||||||
|[0149](https://leetcode.cn/problems/max-points-on-a-line)|Hard|||||||||||
|[0150](https://leetcode.cn/problems/evaluate-reverse-polish-notation)|Medium|||||||||||
|[0151](https://leetcode.cn/problems/reverse-words-in-a-string)|Medium|||||||||||
|[0152](https://leetcode.cn/problems/maximum-product-subarray)|Medium|||||||||||
|[0153](https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array)|Medium|||||||||||
|[0154](https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array-ii)|Hard|||||||||||
|[0155](https://leetcode.cn/problems/min-stack)|Medium|||||||||||
|[0156](https://leetcode.cn/problems/binary-tree-upside-down)|Medium|||||||||||
|[0157](https://leetcode.cn/problems/read-n-characters-given-read4)|Easy|||||||||||
|[0158](https://leetcode.cn/problems/read-n-characters-given-read4-ii-call-multiple-times)|Hard|||||||||||
|[0159](https://leetcode.cn/problems/longest-substring-with-at-most-two-distinct-characters)|Medium|||||||||||
|[0160](https://leetcode.cn/problems/intersection-of-two-linked-lists)|Easy|||||||||||
|[0161](https://leetcode.cn/problems/one-edit-distance)|Medium|||||||||||
|[0162](https://leetcode.cn/problems/find-peak-element)|Medium|||||||||||
|[0163](https://leetcode.cn/problems/missing-ranges)|Easy|||||||||||
|[0164](https://leetcode.cn/problems/maximum-gap)|Medium|||||||||||
|[0165](https://leetcode.cn/problems/compare-version-numbers)|Medium|||||||||||
|[0166](https://leetcode.cn/problems/fraction-to-recurring-decimal)|Medium|||||||||||
|[0167](https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted)|Medium|||||||||||
|[0168](https://leetcode.cn/problems/excel-sheet-column-title)|Easy|||||||||||
|[0169](https://leetcode.cn/problems/majority-element)|Easy|:white_check_mark:|:white_check_mark:||||||:white_check_mark:|||
|[0170](https://leetcode.cn/problems/two-sum-iii-data-structure-design)|Easy|||||||||||
|[0171](https://leetcode.cn/problems/excel-sheet-column-number)|Easy|||||||||||
|[0172](https://leetcode.cn/problems/factorial-trailing-zeroes)|Medium|||||||||||
|[0173](https://leetcode.cn/problems/binary-search-tree-iterator)|Medium|||||||||||
|[0174](https://leetcode.cn/problems/dungeon-game)|Hard|||||||||||
|[0179](https://leetcode.cn/problems/largest-number)|Medium|:white_check_mark:|:white_check_mark:|||||||||
|[0186](https://leetcode.cn/problems/reverse-words-in-a-string-ii)|Medium|||||||||||
|[0187](https://leetcode.cn/problems/repeated-dna-sequences)|Medium|||||||||||
|[0188](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv)|Hard|||||||||||
|[0189](https://leetcode.cn/problems/rotate-array)|Medium|||||||||||
|[0190](https://leetcode.cn/problems/reverse-bits)|Easy|||||||||||
|[0191](https://leetcode.cn/problems/number-of-1-bits)|Easy|||||||||||
|[0198](https://leetcode.cn/problems/house-robber)|Medium|||||||||||
|[0199](https://leetcode.cn/problems/binary-tree-right-side-view)|Medium|||||||||||
|[0200](https://leetcode.cn/problems/number-of-islands)|Medium|||||||||||
|[0201](https://leetcode.cn/problems/bitwise-and-of-numbers-range)|Medium|||||||||||
|[0202](https://leetcode.cn/problems/happy-number)|Easy|||||||||||
|[0203](https://leetcode.cn/problems/remove-linked-list-elements)|Easy|||||||||||
|[0204](https://leetcode.cn/problems/count-primes)|Medium|:white_check_mark:|:white_check_mark:||:white_check_mark:|||||||
|[0205](https://leetcode.cn/problems/isomorphic-strings)|Easy|:white_check_mark:||||||||||
|[0206](https://leetcode.cn/problems/reverse-linked-list)|Easy|||||||||||
|[0207](https://leetcode.cn/problems/course-schedule)|Medium|||||||||||
|[0208](https://leetcode.cn/problems/implement-trie-prefix-tree)|Medium|||||||||||
|[0209](https://leetcode.cn/problems/minimum-size-subarray-sum)|Medium|||||||||||
|[0210](https://leetcode.cn/problems/course-schedule-ii)|Medium|||||||||||
|[0211](https://leetcode.cn/problems/design-add-and-search-words-data-structure)|Medium|||||||||||
|[0212](https://leetcode.cn/problems/word-search-ii)|Hard|||||||||||
|[0213](https://leetcode.cn/problems/house-robber-ii)|Medium|||||||||||
|[0214](https://leetcode.cn/problems/shortest-palindrome)|Hard|||||||||||
|[0215](https://leetcode.cn/problems/kth-largest-element-in-an-array)|Medium|||||||||||
|[0216](https://leetcode.cn/problems/combination-sum-iii)|Medium|||||||||||
|[0217](https://leetcode.cn/problems/contains-duplicate)|Easy|||||||||||
|[0218](https://leetcode.cn/problems/the-skyline-problem)|Hard|||||||||||
|[0219](https://leetcode.cn/problems/contains-duplicate-ii)|Easy|||||||||||
|[0220](https://leetcode.cn/problems/contains-duplicate-iii)|Hard|||||||||||
|[0221](https://leetcode.cn/problems/maximal-square)|Medium|||||||||||
|[0222](https://leetcode.cn/problems/count-complete-tree-nodes)|Easy|||||||||||
|[0223](https://leetcode.cn/problems/rectangle-area)|Medium|||||||||||
|[0224](https://leetcode.cn/problems/basic-calculator)|Hard|||||||||||
|[0225](https://leetcode.cn/problems/implement-stack-using-queues)|Easy|||||||||||
|[0226](https://leetcode.cn/problems/invert-binary-tree)|Easy|:white_check_mark:||||||||||
|[0227](https://leetcode.cn/problems/basic-calculator-ii)|Medium|||||||||||
|[0228](https://leetcode.cn/problems/summary-ranges)|Easy|||||||||||
|[0229](https://leetcode.cn/problems/majority-element-ii)|Medium|||||||||||
|[0230](https://leetcode.cn/problems/kth-smallest-element-in-a-bst)|Medium|||||||||||
|[0231](https://leetcode.cn/problems/power-of-two)|Easy|:white_check_mark:|:white_check_mark:|||||||||
|[0232](https://leetcode.cn/problems/implement-queue-using-stacks)|Easy|||||||||||
|[0233](https://leetcode.cn/problems/number-of-digit-one)|Hard|||||||:white_check_mark:||||
|[0234](https://leetcode.cn/problems/palindrome-linked-list)|Easy|||||||||||
|[0235](https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree)|Medium|||||||||||
|[0236](https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree)|Medium|||||||||||
|[0237](https://leetcode.cn/problems/delete-node-in-a-linked-list)|Medium|||||||||||
|[0238](https://leetcode.cn/problems/product-of-array-except-self)|Medium|||||||||||
|[0239](https://leetcode.cn/problems/sliding-window-maximum)|Hard|||||||||||
|[0240](https://leetcode.cn/problems/search-a-2d-matrix-ii)|Medium|||||||||||
|[0241](https://leetcode.cn/problems/different-ways-to-add-parentheses)|Medium|||||||||||
|[0242](https://leetcode.cn/problems/valid-anagram)|Easy|||||||||||
|[0243](https://leetcode.cn/problems/shortest-word-distance)|Easy|||||||||||
|[0244](https://leetcode.cn/problems/shortest-word-distance-ii)|Medium|||||||||||
|[0245](https://leetcode.cn/problems/shortest-word-distance-iii)|Medium|||||||||||
|[0246](https://leetcode.cn/problems/strobogrammatic-number)|Easy|||||||||||
|[0247](https://leetcode.cn/problems/strobogrammatic-number-ii)|Medium|||||||||||
|[0248](https://leetcode.cn/problems/strobogrammatic-number-iii)|Hard|||||||||||
|[0249](https://leetcode.cn/problems/group-shifted-strings)|Medium|||||||||||
|[0250](https://leetcode.cn/problems/count-univalue-subtrees)|Medium|||||||||||
|[0251](https://leetcode.cn/problems/flatten-2d-vector)|Medium|||||||||||
|[0252](https://leetcode.cn/problems/meeting-rooms)|Easy|||||||||||
|[0253](https://leetcode.cn/problems/meeting-rooms-ii)|Medium|||||||||||
|[0254](https://leetcode.cn/problems/factor-combinations)|Medium|||||||||||
|[0255](https://leetcode.cn/problems/verify-preorder-sequence-in-binary-search-tree)|Medium|||||||||||
|[0256](https://leetcode.cn/problems/paint-house)|Medium|||||||||||
|[0257](https://leetcode.cn/problems/binary-tree-paths)|Easy|||||||||||
|[0258](https://leetcode.cn/problems/add-digits)|Easy|||||:white_check_mark:||||||
|[0259](https://leetcode.cn/problems/3sum-smaller)|Medium|||||||||||
|[0260](https://leetcode.cn/problems/single-number-iii)|Medium|||||||||||
|[0261](https://leetcode.cn/problems/graph-valid-tree)|Medium|||||||||||
|[0263](https://leetcode.cn/problems/ugly-number)|Easy|:white_check_mark:|:white_check_mark:|||||||||
|[0264](https://leetcode.cn/problems/ugly-number-ii)|Medium|||||||||||
|[0265](https://leetcode.cn/problems/paint-house-ii)|Hard|||||||||||
|[0266](https://leetcode.cn/problems/palindrome-permutation)|Easy|||||||||||
|[0267](https://leetcode.cn/problems/palindrome-permutation-ii)|Medium|||||||||||
|[0268](https://leetcode.cn/problems/missing-number)|Easy|||||||||||
|[0269](https://leetcode.cn/problems/alien-dictionary)|Hard|||||||||||
|[0270](https://leetcode.cn/problems/closest-binary-search-tree-value)|Easy|||||||||||
|[0271](https://leetcode.cn/problems/encode-and-decode-strings)|Medium|||||||||||
|[0272](https://leetcode.cn/problems/closest-binary-search-tree-value-ii)|Hard|||||||||||
|[0273](https://leetcode.cn/problems/integer-to-english-words)|Hard|||||||||||
|[0274](https://leetcode.cn/problems/h-index)|Medium|||||||||||
|[0275](https://leetcode.cn/problems/h-index-ii)|Medium|||||||||||
|[0276](https://leetcode.cn/problems/paint-fence)|Medium|||||||||||
|[0277](https://leetcode.cn/problems/find-the-celebrity)|Medium|||||||||||
|[0278](https://leetcode.cn/problems/first-bad-version)|Easy|||||||||||
|[0279](https://leetcode.cn/problems/perfect-squares)|Medium|||||||||||
|[0280](https://leetcode.cn/problems/wiggle-sort)|Medium|||||||||||
|[0281](https://leetcode.cn/problems/zigzag-iterator)|Medium|||||||||||
|[0282](https://leetcode.cn/problems/expression-add-operators)|Hard|||||||||||
|[0283](https://leetcode.cn/problems/move-zeroes)|Easy|:white_check_mark:|:white_check_mark:|||||||||
|[0284](https://leetcode.cn/problems/peeking-iterator)|Medium|||||||||||
|[0285](https://leetcode.cn/problems/inorder-successor-in-bst)|Medium|||||||||||
|[0286](https://leetcode.cn/problems/walls-and-gates)|Medium|||||||||||
|[0287](https://leetcode.cn/problems/find-the-duplicate-number)|Medium|||||||||||
|[0288](https://leetcode.cn/problems/unique-word-abbreviation)|Medium|||||||||||
|[0289](https://leetcode.cn/problems/game-of-life)|Medium|||||||||||
|[0290](https://leetcode.cn/problems/word-pattern)|Easy|||||||||||
|[0291](https://leetcode.cn/problems/word-pattern-ii)|Medium|||||||||||
|[0292](https://leetcode.cn/problems/nim-game)|Easy|:white_check_mark:|:white_check_mark:|:white_check_mark:|:white_check_mark:|||:white_check_mark:|:white_check_mark:|||
|[0293](https://leetcode.cn/problems/flip-game)|Easy||:white_check_mark:|||||||||
|[0294](https://leetcode.cn/problems/flip-game-ii)|Medium||:white_check_mark:|||||||||
|[0295](https://leetcode.cn/problems/find-median-from-data-stream)|Hard|||||||||||
|[0296](https://leetcode.cn/problems/best-meeting-point)|Hard|||||||||||
|[0297](https://leetcode.cn/problems/serialize-and-deserialize-binary-tree)|Hard|||||||||||
|[0298](https://leetcode.cn/problems/binary-tree-longest-consecutive-sequence)|Medium|||||||||||
|[0299](https://leetcode.cn/problems/bulls-and-cows)|Medium|||||||||||
|[0300](https://leetcode.cn/problems/longest-increasing-subsequence)|Medium|:white_check_mark:||||||||||
|[0301](https://leetcode.cn/problems/remove-invalid-parentheses)|Hard|||||||||||
|[0302](https://leetcode.cn/problems/smallest-rectangle-enclosing-black-pixels)|Hard|||||||||||
|[0303](https://leetcode.cn/problems/range-sum-query-immutable)|Easy|||||||||||
|[0304](https://leetcode.cn/problems/range-sum-query-2d-immutable)|Medium|||||||||||
|[0305](https://leetcode.cn/problems/number-of-islands-ii)|Hard|||||||||||
|[0306](https://leetcode.cn/problems/additive-number)|Medium|||||||||||
|[0307](https://leetcode.cn/problems/range-sum-query-mutable)|Medium|||||||||||
|[0308](https://leetcode.cn/problems/range-sum-query-2d-mutable)|Hard|||||||||||
|[0309](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown)|Medium|||||||||||
|[0310](https://leetcode.cn/problems/minimum-height-trees)|Medium|||||||||||
|[0311](https://leetcode.cn/problems/sparse-matrix-multiplication)|Medium|||||||||||
|[0312](https://leetcode.cn/problems/burst-balloons)|Hard|||||||||||
|[0313](https://leetcode.cn/problems/super-ugly-number)|Medium|||||||||||
|[0314](https://leetcode.cn/problems/binary-tree-vertical-order-traversal)|Medium|||||||||||
|[0315](https://leetcode.cn/problems/count-of-smaller-numbers-after-self)|Hard||:white_check_mark:|||||||||
|[0316](https://leetcode.cn/problems/remove-duplicate-letters)|Medium|||||||||||
|[0317](https://leetcode.cn/problems/shortest-distance-from-all-buildings)|Hard|||||||||||
|[0318](https://leetcode.cn/problems/maximum-product-of-word-lengths)|Medium|||||||||||
|[0319](https://leetcode.cn/problems/bulb-switcher)|Medium|:white_check_mark:|:white_check_mark:|:white_check_mark:||||:white_check_mark:|:white_check_mark:|||
|[0320](https://leetcode.cn/problems/generalized-abbreviation)|Medium|||||||||||
|[0321](https://leetcode.cn/problems/create-maximum-number)|Hard|||||||||||
|[0322](https://leetcode.cn/problems/coin-change)|Medium|||||||||||
|[0323](https://leetcode.cn/problems/number-of-connected-components-in-an-undirected-graph)|Medium|||||||||||
|[0324](https://leetcode.cn/problems/wiggle-sort-ii)|Medium|||||||||||
|[0325](https://leetcode.cn/problems/maximum-size-subarray-sum-equals-k)|Medium|||||||||||
|[0326](https://leetcode.cn/problems/power-of-three)|Easy|:white_check_mark:|:white_check_mark:|||||||||
|[0327](https://leetcode.cn/problems/count-of-range-sum)|Hard|||||||||||
|[0328](https://leetcode.cn/problems/odd-even-linked-list)|Medium|||||||||||
|[0329](https://leetcode.cn/problems/longest-increasing-path-in-a-matrix)|Hard|||||||||||
|[0330](https://leetcode.cn/problems/patching-array)|Hard|||||||||||
|[0331](https://leetcode.cn/problems/verify-preorder-serialization-of-a-binary-tree)|Medium|||||||||||
|[0332](https://leetcode.cn/problems/reconstruct-itinerary)|Hard|||||||||||
|[0333](https://leetcode.cn/problems/largest-bst-subtree)|Medium|||||||||||
|[0334](https://leetcode.cn/problems/increasing-triplet-subsequence)|Medium|||||||||||
|[0335](https://leetcode.cn/problems/self-crossing)|Hard|||||||||||
|[0336](https://leetcode.cn/problems/palindrome-pairs)|Hard|||||||||||
|[0337](https://leetcode.cn/problems/house-robber-iii)|Medium|||||||||||
|[0338](https://leetcode.cn/problems/counting-bits)|Easy|||||||||||
|[0339](https://leetcode.cn/problems/nested-list-weight-sum)|Medium|||||||||||
|[0340](https://leetcode.cn/problems/longest-substring-with-at-most-k-distinct-characters)|Medium|||||||||||
|[0341](https://leetcode.cn/problems/flatten-nested-list-iterator)|Medium|||||||||||
|[0342](https://leetcode.cn/problems/power-of-four)|Easy|||||||||||
|[0343](https://leetcode.cn/problems/integer-break)|Medium|||||||||||
|[0344](https://leetcode.cn/problems/reverse-string)|Easy|:white_check_mark:|:white_check_mark:|||||:white_check_mark:|:white_check_mark:|||
|[0345](https://leetcode.cn/problems/reverse-vowels-of-a-string)|Easy|||||||||||
|[0346](https://leetcode.cn/problems/moving-average-from-data-stream)|Easy|||||||||||
|[0347](https://leetcode.cn/problems/top-k-frequent-elements)|Medium|||||||||||
|[0348](https://leetcode.cn/problems/design-tic-tac-toe)|Medium|||||||||||
|[0349](https://leetcode.cn/problems/intersection-of-two-arrays)|Easy|||||||||||
|[0350](https://leetcode.cn/problems/intersection-of-two-arrays-ii)|Easy|||||||||||
|[0351](https://leetcode.cn/problems/android-unlock-patterns)|Medium|||||||||||
|[0352](https://leetcode.cn/problems/data-stream-as-disjoint-intervals)|Hard|||||||||||
|[0353](https://leetcode.cn/problems/design-snake-game)|Medium|||||||||||
|[0354](https://leetcode.cn/problems/russian-doll-envelopes)|Hard|||||||||||
|[0355](https://leetcode.cn/problems/design-twitter)|Medium|||||||||||
|[0356](https://leetcode.cn/problems/line-reflection)|Medium|||||||||||
|[0357](https://leetcode.cn/problems/count-numbers-with-unique-digits)|Medium|||||||||||
|[0358](https://leetcode.cn/problems/rearrange-string-k-distance-apart)|Hard|||||||||||
|[0359](https://leetcode.cn/problems/logger-rate-limiter)|Easy|||||||||||
|[0360](https://leetcode.cn/problems/sort-transformed-array)|Medium|||||||||||
|[0361](https://leetcode.cn/problems/bomb-enemy)|Medium|||||||||||
|[0362](https://leetcode.cn/problems/design-hit-counter)|Medium|||||||||||
|[0363](https://leetcode.cn/problems/max-sum-of-rectangle-no-larger-than-k)|Hard|||||||||||
|[0364](https://leetcode.cn/problems/nested-list-weight-sum-ii)|Medium|||||||||||
|[0365](https://leetcode.cn/problems/water-and-jug-problem)|Medium|||||||||||
|[0366](https://leetcode.cn/problems/find-leaves-of-binary-tree)|Medium|||||||||||
|[0367](https://leetcode.cn/problems/valid-perfect-square)|Easy|||||||||||
|[0368](https://leetcode.cn/problems/largest-divisible-subset)|Medium|||||||||||
|[0369](https://leetcode.cn/problems/plus-one-linked-list)|Medium|||||||||||
|[0370](https://leetcode.cn/problems/range-addition)|Medium|||||||||||
|[0371](https://leetcode.cn/problems/sum-of-two-integers)|Medium|||||||||||
|[0372](https://leetcode.cn/problems/super-pow)|Medium|||||||||||
|[0373](https://leetcode.cn/problems/find-k-pairs-with-smallest-sums)|Medium|||||||||||
|[0374](https://leetcode.cn/problems/guess-number-higher-or-lower)|Easy|||||||||||
|[0375](https://leetcode.cn/problems/guess-number-higher-or-lower-ii)|Medium|||||||||||
|[0376](https://leetcode.cn/problems/wiggle-subsequence)|Medium|||||||||||
|[0377](https://leetcode.cn/problems/combination-sum-iv)|Medium|||||||||||
|[0378](https://leetcode.cn/problems/kth-smallest-element-in-a-sorted-matrix)|Medium|||||||||||
|[0379](https://leetcode.cn/problems/design-phone-directory)|Medium|||||||||||
|[0380](https://leetcode.cn/problems/insert-delete-getrandom-o1)|Medium|||||||||||
|[0381](https://leetcode.cn/problems/insert-delete-getrandom-o1-duplicates-allowed)|Hard|||||||||||
|[0382](https://leetcode.cn/problems/linked-list-random-node)|Medium|||||||||||
|[0383](https://leetcode.cn/problems/ransom-note)|Easy|||||||||||
|[0384](https://leetcode.cn/problems/shuffle-an-array)|Medium|||||||||||
|[0385](https://leetcode.cn/problems/mini-parser)|Medium|||||||||||
|[0386](https://leetcode.cn/problems/lexicographical-numbers)|Medium|||||||||||
|[0387](https://leetcode.cn/problems/first-unique-character-in-a-string)|Easy|||||||||||
|[0388](https://leetcode.cn/problems/longest-absolute-file-path)|Medium|||||||||||
|[0389](https://leetcode.cn/problems/find-the-difference)|Easy|||||||||||
|[0390](https://leetcode.cn/problems/elimination-game)|Medium|||||||||||
|[0391](https://leetcode.cn/problems/perfect-rectangle)|Hard|||||||||||
|[0392](https://leetcode.cn/problems/is-subsequence)|Easy|||||||||||
|[0393](https://leetcode.cn/problems/utf-8-validation)|Medium|||||||||||
|[0394](https://leetcode.cn/problems/decode-string)|Medium|||||||||||
|[0395](https://leetcode.cn/problems/longest-substring-with-at-least-k-repeating-characters)|Medium|||||||||||
|[0396](https://leetcode.cn/problems/rotate-function)|Medium|||||||||||
|[0397](https://leetcode.cn/problems/integer-replacement)|Medium|||||||||||
|[0398](https://leetcode.cn/problems/random-pick-index)|Medium|||||||||||
|[0399](https://leetcode.cn/problems/evaluate-division)|Medium|||||||||||
|[0400](https://leetcode.cn/problems/nth-digit)|Medium|||||||||||
|[0401](https://leetcode.cn/problems/binary-watch)|Easy|||||||||||
|[0402](https://leetcode.cn/problems/remove-k-digits)|Medium|||||||||||
|[0403](https://leetcode.cn/problems/frog-jump)|Hard|||||||||||
|[0404](https://leetcode.cn/problems/sum-of-left-leaves)|Easy|||||||||||
|[0405](https://leetcode.cn/problems/convert-a-number-to-hexadecimal)|Easy|||||||||||
|[0406](https://leetcode.cn/problems/queue-reconstruction-by-height)|Medium|||||||||||
|[0407](https://leetcode.cn/problems/trapping-rain-water-ii)|Hard|||||||||||
|[0408](https://leetcode.cn/problems/valid-word-abbreviation)|Easy|||||||||||
|[0409](https://leetcode.cn/problems/longest-palindrome)|Easy|||||||||||
|[0410](https://leetcode.cn/problems/split-array-largest-sum)|Hard|||||||||||
|[0411](https://leetcode.cn/problems/minimum-unique-word-abbreviation)|Hard|||||||||||
|[0412](https://leetcode.cn/problems/fizz-buzz)|Easy|||||||||||
|[0413](https://leetcode.cn/problems/arithmetic-slices)|Medium|||||||||||
|[0414](https://leetcode.cn/problems/third-maximum-number)|Easy|||||||||||
|[0415](https://leetcode.cn/problems/add-strings)|Easy|||||||||||
|[0416](https://leetcode.cn/problems/partition-equal-subset-sum)|Medium|||||||||||
|[0417](https://leetcode.cn/problems/pacific-atlantic-water-flow)|Medium|||||||||||
|[0418](https://leetcode.cn/problems/sentence-screen-fitting)|Medium|||||||||||
|[0419](https://leetcode.cn/problems/battleships-in-a-board)|Medium|||||||||||
|[0420](https://leetcode.cn/problems/strong-password-checker)|Hard|||||||||||
|[0421](https://leetcode.cn/problems/maximum-xor-of-two-numbers-in-an-array)|Medium|||||||||||
|[0422](https://leetcode.cn/problems/valid-word-square)|Easy|||||||||||
|[0423](https://leetcode.cn/problems/reconstruct-original-digits-from-english)|Medium|||||||||||
|[0424](https://leetcode.cn/problems/longest-repeating-character-replacement)|Medium|||||||||||
|[0425](https://leetcode.cn/problems/word-squares)|Hard|||||||||||
|[0432](https://leetcode.cn/problems/all-oone-data-structure)|Hard|||||||||||
|[0433](https://leetcode.cn/problems/minimum-genetic-mutation)|Medium|||||||||||
|[0434](https://leetcode.cn/problems/number-of-segments-in-a-string)|Easy|||||||||||
|[0435](https://leetcode.cn/problems/non-overlapping-intervals)|Medium|||||||||||
|[0436](https://leetcode.cn/problems/find-right-interval)|Medium|||||||||||
|[0437](https://leetcode.cn/problems/path-sum-iii)|Medium|||||||||||
|[0438](https://leetcode.cn/problems/find-all-anagrams-in-a-string)|Medium|||||||||||
|[0439](https://leetcode.cn/problems/ternary-expression-parser)|Medium|||||||||||
|[0440](https://leetcode.cn/problems/k-th-smallest-in-lexicographical-order)|Hard|||||||||||
|[0441](https://leetcode.cn/problems/arranging-coins)|Easy|||||||||||
|[0442](https://leetcode.cn/problems/find-all-duplicates-in-an-array)|Medium|||||||||||
|[0443](https://leetcode.cn/problems/string-compression)|Medium|||||||||||
|[0444](https://leetcode.cn/problems/sequence-reconstruction)|Medium|||||||||||
|[0445](https://leetcode.cn/problems/add-two-numbers-ii)|Medium|||||||||||
|[0446](https://leetcode.cn/problems/arithmetic-slices-ii-subsequence)|Hard|||||||||||
|[0447](https://leetcode.cn/problems/number-of-boomerangs)|Medium|||||||||||
|[0448](https://leetcode.cn/problems/find-all-numbers-disappeared-in-an-array)|Easy|||||||||||
|[0449](https://leetcode.cn/problems/serialize-and-deserialize-bst)|Medium|||||||||||
|[0450](https://leetcode.cn/problems/delete-node-in-a-bst)|Medium|||||||||||
|[0451](https://leetcode.cn/problems/sort-characters-by-frequency)|Medium|||||||||||
|[0452](https://leetcode.cn/problems/minimum-number-of-arrows-to-burst-balloons)|Medium|||||||||||
|[0453](https://leetcode.cn/problems/minimum-moves-to-equal-array-elements)|Medium|||||||||||
|[0454](https://leetcode.cn/problems/4sum-ii)|Medium|||||||||||
|[0455](https://leetcode.cn/problems/assign-cookies)|Easy|||||||||||
|[0456](https://leetcode.cn/problems/132-pattern)|Medium|||||||||||
|[0457](https://leetcode.cn/problems/circular-array-loop)|Medium|||||||||||
|[0458](https://leetcode.cn/problems/poor-pigs)|Hard|:white_check_mark:|:white_check_mark:|:white_check_mark:|:white_check_mark:||:white_check_mark:|:white_check_mark:|:white_check_mark:|||
|[0459](https://leetcode.cn/problems/repeated-substring-pattern)|Easy|||||||||||
|[0460](https://leetcode.cn/problems/lfu-cache)|Hard|||||||||||
|[0461](https://leetcode.cn/problems/hamming-distance)|Easy|:white_check_mark:|:white_check_mark:|:white_check_mark:|:white_check_mark:|||:white_check_mark:||||
|[0462](https://leetcode.cn/problems/minimum-moves-to-equal-array-elements-ii)|Medium|||||||||||
|[0463](https://leetcode.cn/problems/island-perimeter)|Easy|||||||||||
|[0464](https://leetcode.cn/problems/can-i-win)|Medium|||||||||||
|[0465](https://leetcode.cn/problems/optimal-account-balancing)|Hard|||||||||||
|[0466](https://leetcode.cn/problems/count-the-repetitions)|Hard|||||||||||
|[0467](https://leetcode.cn/problems/unique-substrings-in-wraparound-string)|Medium|||||||||||
|[0468](https://leetcode.cn/problems/validate-ip-address)|Medium|||||||||||
|[0469](https://leetcode.cn/problems/convex-polygon)|Medium|||||||||||
|[0471](https://leetcode.cn/problems/encode-string-with-shortest-length)|Hard|||||||||||
|[0472](https://leetcode.cn/problems/concatenated-words)|Hard|||||||||||
|[0473](https://leetcode.cn/problems/matchsticks-to-square)|Medium|||||||||||
|[0474](https://leetcode.cn/problems/ones-and-zeroes)|Medium|||||||||||
|[0475](https://leetcode.cn/problems/heaters)|Medium|||||||||||
|[0476](https://leetcode.cn/problems/number-complement)|Easy|||||||||||
|[0477](https://leetcode.cn/problems/total-hamming-distance)|Medium|||||||||||
|[0479](https://leetcode.cn/problems/largest-palindrome-product)|Hard|||||||||||
|[0480](https://leetcode.cn/problems/sliding-window-median)|Hard|||||||||||
|[0481](https://leetcode.cn/problems/magical-string)|Medium|||||||||||
|[0482](https://leetcode.cn/problems/license-key-formatting)|Easy|||||||||||
|[0483](https://leetcode.cn/problems/smallest-good-base)|Hard|||||||||||
|[0484](https://leetcode.cn/problems/find-permutation)|Medium|||||||||||
|[0485](https://leetcode.cn/problems/max-consecutive-ones)|Easy|||||||||||
|[0486](https://leetcode.cn/problems/predict-the-winner)|Medium|||||||||||
|[0487](https://leetcode.cn/problems/max-consecutive-ones-ii)|Medium|||||||||||
|[0488](https://leetcode.cn/problems/zuma-game)|Hard|||||||||||
|[1643](https://leetcode.cn/problems/kth-smallest-instructions)|Hard|||||||||||
|[0490](https://leetcode.cn/problems/the-maze)|Medium|||||||||||
|[0491](https://leetcode.cn/problems/non-decreasing-subsequences)|Medium|||||||||||
|[0492](https://leetcode.cn/problems/construct-the-rectangle)|Easy|||||||||||
|[0493](https://leetcode.cn/problems/reverse-pairs)|Hard|:white_check_mark:||||||||||
|[0494](https://leetcode.cn/problems/target-sum)|Medium|||||||||||
|[0495](https://leetcode.cn/problems/teemo-attacking)|Easy|||||||||||
|[0496](https://leetcode.cn/problems/next-greater-element-i)|Easy|||||||||||
|[0498](https://leetcode.cn/problems/diagonal-traverse)|Medium|:white_check_mark:||||||||||
|[0499](https://leetcode.cn/problems/the-maze-iii)|Hard|||||||||||
|[0500](https://leetcode.cn/problems/keyboard-row)|Easy|||||||||||
|[0501](https://leetcode.cn/problems/find-mode-in-binary-search-tree)|Easy|||||||||||
|[0502](https://leetcode.cn/problems/ipo)|Hard|||||||||||
|[0503](https://leetcode.cn/problems/next-greater-element-ii)|Medium|||||||||||
|[0504](https://leetcode.cn/problems/base-7)|Easy|||||||||||
|[0505](https://leetcode.cn/problems/the-maze-ii)|Medium|||||||||||
|[0506](https://leetcode.cn/problems/relative-ranks)|Easy|||||||||||
|[0507](https://leetcode.cn/problems/perfect-number)|Easy|||||||||||
|[0508](https://leetcode.cn/problems/most-frequent-subtree-sum)|Medium|||||||||||
|[0510](https://leetcode.cn/problems/inorder-successor-in-bst-ii)|Medium|||||||||||
|[2031](https://leetcode.cn/problems/count-subarrays-with-more-ones-than-zeros)|Medium|||||||||||
|[1059](https://leetcode.cn/problems/all-paths-from-source-lead-to-destination)|Medium|||||||||||
|[2036](https://leetcode.cn/problems/maximum-alternating-subarray-sum)|Medium|||||||||||
|[0513](https://leetcode.cn/problems/find-bottom-left-tree-value)|Medium|||||||||||
|[0514](https://leetcode.cn/problems/freedom-trail)|Hard|:white_check_mark:|:white_check_mark:||:white_check_mark:|||||||
|[0515](https://leetcode.cn/problems/find-largest-value-in-each-tree-row)|Medium|||||||||||
|[0516](https://leetcode.cn/problems/longest-palindromic-subsequence)|Medium|||||||||||
|[0517](https://leetcode.cn/problems/super-washing-machines)|Hard|||||||||||
|[0518](https://leetcode.cn/problems/coin-change-ii)|Medium|||||||||||
|[1983](https://leetcode.cn/problems/widest-pair-of-indices-with-equal-range-sum)|Medium|||||||||||
|[0520](https://leetcode.cn/problems/detect-capital)|Easy|:white_check_mark:|:white_check_mark:|||||||||
|[0521](https://leetcode.cn/problems/longest-uncommon-subsequence-i)|Easy|||||||||||
|[0522](https://leetcode.cn/problems/longest-uncommon-subsequence-ii)|Medium|||||||||||
|[0523](https://leetcode.cn/problems/continuous-subarray-sum)|Medium|||||||||||
|[0524](https://leetcode.cn/problems/longest-word-in-dictionary-through-deleting)|Medium|||||||||||
|[0525](https://leetcode.cn/problems/contiguous-array)|Medium|||||||||||
|[0526](https://leetcode.cn/problems/beautiful-arrangement)|Medium|||||||||||
|[0527](https://leetcode.cn/problems/word-abbreviation)|Hard|||||||||||
|[1721](https://leetcode.cn/problems/swapping-nodes-in-a-linked-list)|Medium|||||||||||
|[0529](https://leetcode.cn/problems/minesweeper)|Medium|||||||||||
|[0530](https://leetcode.cn/problems/minimum-absolute-difference-in-bst)|Easy|||||||||||
|[0531](https://leetcode.cn/problems/lonely-pixel-i)|Medium|||||||||||
|[0532](https://leetcode.cn/problems/k-diff-pairs-in-an-array)|Medium|||||||||||
|[0533](https://leetcode.cn/problems/lonely-pixel-ii)|Medium|||||||||||
|[0535](https://leetcode.cn/problems/encode-and-decode-tinyurl)|Medium|||||||||||
|[0536](https://leetcode.cn/problems/construct-binary-tree-from-string)|Medium|||||||||||
|[0537](https://leetcode.cn/problems/complex-number-multiplication)|Medium|||||||||||
|[0538](https://leetcode.cn/problems/convert-bst-to-greater-tree)|Medium|||||||||||
|[0539](https://leetcode.cn/problems/minimum-time-difference)|Medium|||||||||||
|[0540](https://leetcode.cn/problems/single-element-in-a-sorted-array)|Medium|||||||||||
|[0541](https://leetcode.cn/problems/reverse-string-ii)|Easy|||||||||||
|[0542](https://leetcode.cn/problems/01-matrix)|Medium|||||||||||
|[0543](https://leetcode.cn/problems/diameter-of-binary-tree)|Easy|||||||||||
|[0544](https://leetcode.cn/problems/output-contest-matches)|Medium|||||||||||
|[0545](https://leetcode.cn/problems/boundary-of-binary-tree)|Medium|||||||||||
|[0546](https://leetcode.cn/problems/remove-boxes)|Hard|||||||||||
|[0547](https://leetcode.cn/problems/number-of-provinces)|Medium|||||||||||
|[0548](https://leetcode.cn/problems/split-array-with-equal-sum)|Hard|||||||||||
|[0549](https://leetcode.cn/problems/binary-tree-longest-consecutive-sequence-ii)|Medium|||||||||||
|[1730](https://leetcode.cn/problems/shortest-path-to-get-food)|Medium|||||||||||
|[0551](https://leetcode.cn/problems/student-attendance-record-i)|Easy|||||||||||
|[0552](https://leetcode.cn/problems/student-attendance-record-ii)|Hard|||||||||||
|[0553](https://leetcode.cn/problems/optimal-division)|Medium|||||||||||
|[0554](https://leetcode.cn/problems/brick-wall)|Medium|||||||||||
|[0555](https://leetcode.cn/problems/split-concatenated-strings)|Medium|||||||||||
|[0556](https://leetcode.cn/problems/next-greater-element-iii)|Medium|||||||||||
|[0557](https://leetcode.cn/problems/reverse-words-in-a-string-iii)|Easy|||||||||||
|[0560](https://leetcode.cn/problems/subarray-sum-equals-k)|Medium|||||||||||
|[0561](https://leetcode.cn/problems/array-partition)|Easy|||||||||||
|[0562](https://leetcode.cn/problems/longest-line-of-consecutive-one-in-matrix)|Medium|||||||||||
|[0563](https://leetcode.cn/problems/binary-tree-tilt)|Easy|||||||||||
|[0564](https://leetcode.cn/problems/find-the-closest-palindrome)|Hard|||||||||||
|[0565](https://leetcode.cn/problems/array-nesting)|Medium|||||||||||
|[0566](https://leetcode.cn/problems/reshape-the-matrix)|Easy|||||||||||
|[0567](https://leetcode.cn/problems/permutation-in-string)|Medium|||||||||||
|[0568](https://leetcode.cn/problems/maximum-vacation-days)|Hard|||||||||||
|[0572](https://leetcode.cn/problems/subtree-of-another-tree)|Easy|||||||||||
|[0573](https://leetcode.cn/problems/squirrel-simulation)|Medium|||||||||||
|[0575](https://leetcode.cn/problems/distribute-candies)|Easy|||||||||||
|[0576](https://leetcode.cn/problems/out-of-boundary-paths)|Medium|||||||||||
|[0581](https://leetcode.cn/problems/shortest-unsorted-continuous-subarray)|Medium|||||||||||
|[0582](https://leetcode.cn/problems/kill-process)|Medium|||||||||||
|[0583](https://leetcode.cn/problems/delete-operation-for-two-strings)|Medium|||||||||||
|[0587](https://leetcode.cn/problems/erect-the-fence)|Hard|||||||||||
|[0588](https://leetcode.cn/problems/design-in-memory-file-system)|Hard|||||||||||
|[0591](https://leetcode.cn/problems/tag-validator)|Hard|||||||||||
|[0592](https://leetcode.cn/problems/fraction-addition-and-subtraction)|Medium|||||||||||
|[0593](https://leetcode.cn/problems/valid-square)|Medium|||||||||||
|[0594](https://leetcode.cn/problems/longest-harmonious-subsequence)|Easy|||||||||||
|[0598](https://leetcode.cn/problems/range-addition-ii)|Easy|||||||||||
|[0599](https://leetcode.cn/problems/minimum-index-sum-of-two-lists)|Easy|||||||||||
|[0600](https://leetcode.cn/problems/non-negative-integers-without-consecutive-ones)|Hard|||||||||||
|[0604](https://leetcode.cn/problems/design-compressed-string-iterator)|Easy|||||||||||
|[0605](https://leetcode.cn/problems/can-place-flowers)|Easy|||||||||||
|[0606](https://leetcode.cn/problems/construct-string-from-binary-tree)|Medium|||||||||||
|[0609](https://leetcode.cn/problems/find-duplicate-file-in-system)|Medium|||||||||||
|[0611](https://leetcode.cn/problems/valid-triangle-number)|Medium|||||||||||
|[0616](https://leetcode.cn/problems/add-bold-tag-in-string)|Medium|||||||||||
|[0617](https://leetcode.cn/problems/merge-two-binary-trees)|Easy|||||||||||
|[0621](https://leetcode.cn/problems/task-scheduler)|Medium|||||||||||
|[0623](https://leetcode.cn/problems/add-one-row-to-tree)|Medium|||||||||||
|[0624](https://leetcode.cn/problems/maximum-distance-in-arrays)|Medium|||||||||||
|[0625](https://leetcode.cn/problems/minimum-factorization)|Medium|||||||||||
|[0628](https://leetcode.cn/problems/maximum-product-of-three-numbers)|Easy|||||||||||
|[0629](https://leetcode.cn/problems/k-inverse-pairs-array)|Hard|||||||||||
|[0630](https://leetcode.cn/problems/course-schedule-iii)|Hard|||||||||||
|[0631](https://leetcode.cn/problems/design-excel-sum-formula)|Hard|||||||||||
|[0632](https://leetcode.cn/problems/smallest-range-covering-elements-from-k-lists)|Hard|||||||||||
|[0633](https://leetcode.cn/problems/sum-of-square-numbers)|Medium|||||||||||
|[0634](https://leetcode.cn/problems/find-the-derangement-of-an-array)|Medium|||||||||||
|[0635](https://leetcode.cn/problems/design-log-storage-system)|Medium|||||||||||
|[0636](https://leetcode.cn/problems/exclusive-time-of-functions)|Medium|||||||||||
|[0637](https://leetcode.cn/problems/average-of-levels-in-binary-tree)|Easy|||||||||||
|[0638](https://leetcode.cn/problems/shopping-offers)|Medium|||||||||||
|[0639](https://leetcode.cn/problems/decode-ways-ii)|Hard|||||||||||
|[0640](https://leetcode.cn/problems/solve-the-equation)|Medium|||||||||||
|[0642](https://leetcode.cn/problems/design-search-autocomplete-system)|Hard|||||||||||
|[0643](https://leetcode.cn/problems/maximum-average-subarray-i)|Easy|||||||||||
|[0644](https://leetcode.cn/problems/maximum-average-subarray-ii)|Hard|||||||||||
|[0645](https://leetcode.cn/problems/set-mismatch)|Easy|||||||||||
|[0646](https://leetcode.cn/problems/maximum-length-of-pair-chain)|Medium|||||||||||
|[0647](https://leetcode.cn/problems/palindromic-substrings)|Medium|||||||||||
|[0648](https://leetcode.cn/problems/replace-words)|Medium|||||||||||
|[0649](https://leetcode.cn/problems/dota2-senate)|Medium|||||||||||
|[0650](https://leetcode.cn/problems/2-keys-keyboard)|Medium|||||||||||
|[0651](https://leetcode.cn/problems/4-keys-keyboard)|Medium|||||||||||
|[0652](https://leetcode.cn/problems/find-duplicate-subtrees)|Medium|||||||||||
|[0653](https://leetcode.cn/problems/two-sum-iv-input-is-a-bst)|Easy|||||||||||
|[0654](https://leetcode.cn/problems/maximum-binary-tree)|Medium|||||||||||
|[0655](https://leetcode.cn/problems/print-binary-tree)|Medium|||||||||||
|[0656](https://leetcode.cn/problems/coin-path)|Hard|||||||||||
|[0657](https://leetcode.cn/problems/robot-return-to-origin)|Easy|||||||||||
|[0658](https://leetcode.cn/problems/find-k-closest-elements)|Medium|||||||||||
|[0659](https://leetcode.cn/problems/split-array-into-consecutive-subsequences)|Medium|||||||||||
|[0660](https://leetcode.cn/problems/remove-9)|Hard|||||||||||
|[0661](https://leetcode.cn/problems/image-smoother)|Easy|||||||||||
|[0662](https://leetcode.cn/problems/maximum-width-of-binary-tree)|Medium|||||||||||
|[0663](https://leetcode.cn/problems/equal-tree-partition)|Medium|||||||||||
|[0664](https://leetcode.cn/problems/strange-printer)|Hard|||||||||||
|[0665](https://leetcode.cn/problems/non-decreasing-array)|Medium|||||||||||
|[0666](https://leetcode.cn/problems/path-sum-iv)|Medium|||||||||||
|[0667](https://leetcode.cn/problems/beautiful-arrangement-ii)|Medium|||||||||||
|[0668](https://leetcode.cn/problems/kth-smallest-number-in-multiplication-table)|Hard|||||||||||
|[0669](https://leetcode.cn/problems/trim-a-binary-search-tree)|Medium|||||||||||
|[0670](https://leetcode.cn/problems/maximum-swap)|Medium|||||||||||
|[0671](https://leetcode.cn/problems/second-minimum-node-in-a-binary-tree)|Easy|||||||||||
|[0672](https://leetcode.cn/problems/bulb-switcher-ii)|Medium|||||||||||
|[0673](https://leetcode.cn/problems/number-of-longest-increasing-subsequence)|Medium|||||||||||
|[0674](https://leetcode.cn/problems/longest-continuous-increasing-subsequence)|Easy|||||||||||
|[0675](https://leetcode.cn/problems/cut-off-trees-for-golf-event)|Hard|||||||||||
|[0676](https://leetcode.cn/problems/implement-magic-dictionary)|Medium|||||||||||
|[0677](https://leetcode.cn/problems/map-sum-pairs)|Medium|||||||||||
|[0678](https://leetcode.cn/problems/valid-parenthesis-string)|Medium|||||||||||
|[0679](https://leetcode.cn/problems/24-game)|Hard|:white_check_mark:|:white_check_mark:||||||:white_check_mark:|||
|[0680](https://leetcode.cn/problems/valid-palindrome-ii)|Easy|||||||||||
|[0681](https://leetcode.cn/problems/next-closest-time)|Medium|||||||||||
|[0682](https://leetcode.cn/problems/baseball-game)|Easy|||||||||||
|[0683](https://leetcode.cn/problems/k-empty-slots)|Hard|||||||||||
|[0684](https://leetcode.cn/problems/redundant-connection)|Medium|||||||||||
|[0685](https://leetcode.cn/problems/redundant-connection-ii)|Hard|||||||||||
|[0686](https://leetcode.cn/problems/repeated-string-match)|Medium|||||||||||
|[0687](https://leetcode.cn/problems/longest-univalue-path)|Medium|||||||||||
|[0688](https://leetcode.cn/problems/knight-probability-in-chessboard)|Medium|||||||||||
|[0689](https://leetcode.cn/problems/maximum-sum-of-3-non-overlapping-subarrays)|Hard|||||||||||
|[0690](https://leetcode.cn/problems/employee-importance)|Medium|||||||||||
|[0691](https://leetcode.cn/problems/stickers-to-spell-word)|Hard|||||||||||
|[0692](https://leetcode.cn/problems/top-k-frequent-words)|Medium|||||||||||
|[0693](https://leetcode.cn/problems/binary-number-with-alternating-bits)|Easy|||||||||||
|[0694](https://leetcode.cn/problems/number-of-distinct-islands)|Medium|||||||||||
|[0695](https://leetcode.cn/problems/max-area-of-island)|Medium|||||||||||
|[0696](https://leetcode.cn/problems/count-binary-substrings)|Easy|:white_check_mark:|:white_check_mark:||||||:white_check_mark:|||
|[0697](https://leetcode.cn/problems/degree-of-an-array)|Easy|||||||||||
|[0698](https://leetcode.cn/problems/partition-to-k-equal-sum-subsets)|Medium|||||||||||
|[0699](https://leetcode.cn/problems/falling-squares)|Hard|||||||||||
|[0711](https://leetcode.cn/problems/number-of-distinct-islands-ii)|Hard|||||||||||
|[0712](https://leetcode.cn/problems/minimum-ascii-delete-sum-for-two-strings)|Medium|||||||||||
|[0713](https://leetcode.cn/problems/subarray-product-less-than-k)|Medium|||||||||||
|[0714](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee)|Medium|||||||||||
|[0715](https://leetcode.cn/problems/range-module)|Hard|||||||||||
|[0716](https://leetcode.cn/problems/max-stack)|Hard|||||||||||
|[0717](https://leetcode.cn/problems/1-bit-and-2-bit-characters)|Easy|||||||||||
|[0718](https://leetcode.cn/problems/maximum-length-of-repeated-subarray)|Medium|||||||||||
|[0719](https://leetcode.cn/problems/find-k-th-smallest-pair-distance)|Hard|||||||||||
|[0720](https://leetcode.cn/problems/longest-word-in-dictionary)|Medium|||||||||||
|[0721](https://leetcode.cn/problems/accounts-merge)|Medium|||||||||||
|[0722](https://leetcode.cn/problems/remove-comments)|Medium|||||||||||
|[0723](https://leetcode.cn/problems/candy-crush)|Medium|||||||||||
|[0724](https://leetcode.cn/problems/find-pivot-index)|Easy|||||||||||
|[0725](https://leetcode.cn/problems/split-linked-list-in-parts)|Medium|||||||||||
|[0726](https://leetcode.cn/problems/number-of-atoms)|Hard|||||||||||
|[0727](https://leetcode.cn/problems/minimum-window-subsequence)|Hard|||||||||||
|[0728](https://leetcode.cn/problems/self-dividing-numbers)|Easy|||||||||||
|[0729](https://leetcode.cn/problems/my-calendar-i)|Medium|||||||||||
|[0730](https://leetcode.cn/problems/count-different-palindromic-subsequences)|Hard|||||||||||
|[0731](https://leetcode.cn/problems/my-calendar-ii)|Medium|||||||||||
|[0732](https://leetcode.cn/problems/my-calendar-iii)|Hard|||||||||||
|[0733](https://leetcode.cn/problems/flood-fill)|Easy|||||||||||
|[0734](https://leetcode.cn/problems/sentence-similarity)|Easy|||||||||||
|[0735](https://leetcode.cn/problems/asteroid-collision)|Medium|||||||||||
|[0736](https://leetcode.cn/problems/parse-lisp-expression)|Hard|||||||||||
|[0737](https://leetcode.cn/problems/sentence-similarity-ii)|Medium|||||||||||
|[0738](https://leetcode.cn/problems/monotone-increasing-digits)|Medium|||||||||||
|[0739](https://leetcode.cn/problems/daily-temperatures)|Medium|||||||||||
|[0740](https://leetcode.cn/problems/delete-and-earn)|Medium|||||||||||
|[0741](https://leetcode.cn/problems/cherry-pickup)|Hard|||||||||||
|[0709](https://leetcode.cn/problems/to-lower-case)|Easy|||||||||||
|[0742](https://leetcode.cn/problems/closest-leaf-in-a-binary-tree)|Medium|||||||||||
|[0743](https://leetcode.cn/problems/network-delay-time)|Medium|||||||||||
|[0744](https://leetcode.cn/problems/find-smallest-letter-greater-than-target)|Easy|||||||||||
|[0745](https://leetcode.cn/problems/prefix-and-suffix-search)|Hard|||||||||||
|[0746](https://leetcode.cn/problems/min-cost-climbing-stairs)|Easy|||||||||||
|[0747](https://leetcode.cn/problems/largest-number-at-least-twice-of-others)|Easy|||||||||||
|[0748](https://leetcode.cn/problems/shortest-completing-word)|Easy|||||||||||
|[0749](https://leetcode.cn/problems/contain-virus)|Hard|||||||||||
|[0750](https://leetcode.cn/problems/number-of-corner-rectangles)|Medium|||||||||||
|[0751](https://leetcode.cn/problems/ip-to-cidr)|Medium|||||||||||
|[0752](https://leetcode.cn/problems/open-the-lock)|Medium|||||||||||
|[0753](https://leetcode.cn/problems/cracking-the-safe)|Hard|||||||||||
|[0754](https://leetcode.cn/problems/reach-a-number)|Medium|||||||||||
|[0755](https://leetcode.cn/problems/pour-water)|Medium|||||||||||
|[0756](https://leetcode.cn/problems/pyramid-transition-matrix)|Medium|||||||||||
|[0426](https://leetcode.cn/problems/convert-binary-search-tree-to-sorted-doubly-linked-list)|Medium|||||||||||
|[0757](https://leetcode.cn/problems/set-intersection-size-at-least-two)|Hard|||||||||||
|[0758](https://leetcode.cn/problems/bold-words-in-string)|Medium|||||||||||
|[0759](https://leetcode.cn/problems/employee-free-time)|Hard|||||||||||
|[0760](https://leetcode.cn/problems/find-anagram-mappings)|Easy|||||||||||
|[0761](https://leetcode.cn/problems/special-binary-string)|Hard|||||||||||
|[0429](https://leetcode.cn/problems/n-ary-tree-level-order-traversal)|Medium|||||||||||
|[0428](https://leetcode.cn/problems/serialize-and-deserialize-n-ary-tree)|Hard|||||||||||
|[0430](https://leetcode.cn/problems/flatten-a-multilevel-doubly-linked-list)|Medium|||||||||||
|[0762](https://leetcode.cn/problems/prime-number-of-set-bits-in-binary-representation)|Easy|||||||||||
|[0763](https://leetcode.cn/problems/partition-labels)|Medium|||||||||||
|[0764](https://leetcode.cn/problems/largest-plus-sign)|Medium|||||||||||
|[0765](https://leetcode.cn/problems/couples-holding-hands)|Hard|||||||||||
|[0431](https://leetcode.cn/problems/encode-n-ary-tree-to-binary-tree)|Hard|||||||||||
|[0427](https://leetcode.cn/problems/construct-quad-tree)|Medium|||||||||||
|[0558](https://leetcode.cn/problems/logical-or-of-two-binary-grids-represented-as-quad-trees)|Medium|||||||||||
|[0559](https://leetcode.cn/problems/maximum-depth-of-n-ary-tree)|Easy|||||||||||
|[0589](https://leetcode.cn/problems/n-ary-tree-preorder-traversal)|Easy|||||||||||
|[0590](https://leetcode.cn/problems/n-ary-tree-postorder-traversal)|Easy|||||||||||
|[0766](https://leetcode.cn/problems/toeplitz-matrix)|Easy|||||||||||
|[0767](https://leetcode.cn/problems/reorganize-string)|Medium|||||||||||
|[0768](https://leetcode.cn/problems/max-chunks-to-make-sorted-ii)|Hard|||||||||||
|[0769](https://leetcode.cn/problems/max-chunks-to-make-sorted)|Medium|||||||||||
|[0770](https://leetcode.cn/problems/basic-calculator-iv)|Hard|||||||||||
|[0771](https://leetcode.cn/problems/jewels-and-stones)|Easy|||||||||||
|[0700](https://leetcode.cn/problems/search-in-a-binary-search-tree)|Easy|||||||||||
|[0701](https://leetcode.cn/problems/insert-into-a-binary-search-tree)|Medium|||||||||||
|[0772](https://leetcode.cn/problems/basic-calculator-iii)|Hard|||||||||||
|[0702](https://leetcode.cn/problems/search-in-a-sorted-array-of-unknown-size)|Medium|||||||||||
|[0773](https://leetcode.cn/problems/sliding-puzzle)|Hard|||||||||||
|[0774](https://leetcode.cn/problems/minimize-max-distance-to-gas-station)|Hard|||||||||||
|[0703](https://leetcode.cn/problems/kth-largest-element-in-a-stream)|Easy|||||||||||
|[0775](https://leetcode.cn/problems/global-and-local-inversions)|Medium|||||||||||
|[0776](https://leetcode.cn/problems/split-bst)|Medium|||||||||||
|[0704](https://leetcode.cn/problems/binary-search)|Easy|||||||||||
|[0777](https://leetcode.cn/problems/swap-adjacent-in-lr-string)|Medium|||||||||||
|[0778](https://leetcode.cn/problems/swim-in-rising-water)|Hard|||||||||||
|[0779](https://leetcode.cn/problems/k-th-symbol-in-grammar)|Medium|||||||||||
|[0780](https://leetcode.cn/problems/reaching-points)|Hard|||||||||||
|[0781](https://leetcode.cn/problems/rabbits-in-forest)|Medium|||||||||||
|[0782](https://leetcode.cn/problems/transform-to-chessboard)|Hard|||||||||||
|[0783](https://leetcode.cn/problems/minimum-distance-between-bst-nodes)|Easy|||||||||||
|[0784](https://leetcode.cn/problems/letter-case-permutation)|Medium|||||||||||
|[0785](https://leetcode.cn/problems/is-graph-bipartite)|Medium|||||||||||
|[0786](https://leetcode.cn/problems/k-th-smallest-prime-fraction)|Medium|||||||||||
|[0787](https://leetcode.cn/problems/cheapest-flights-within-k-stops)|Medium|||||||||||
|[0788](https://leetcode.cn/problems/rotated-digits)|Medium|||||||||||
|[0789](https://leetcode.cn/problems/escape-the-ghosts)|Medium|||||||||||
|[0790](https://leetcode.cn/problems/domino-and-tromino-tiling)|Medium|||||||||||
|[0791](https://leetcode.cn/problems/custom-sort-string)|Medium|||||||||||
|[0792](https://leetcode.cn/problems/number-of-matching-subsequences)|Medium|||||||||||
|[0793](https://leetcode.cn/problems/preimage-size-of-factorial-zeroes-function)|Hard|||||||||||
|[0794](https://leetcode.cn/problems/valid-tic-tac-toe-state)|Medium|||||||||||
|[0795](https://leetcode.cn/problems/number-of-subarrays-with-bounded-maximum)|Medium|||||||||||
|[0796](https://leetcode.cn/problems/rotate-string)|Easy|||||||||||
|[0797](https://leetcode.cn/problems/all-paths-from-source-to-target)|Medium|||||||||||
|[0798](https://leetcode.cn/problems/smallest-rotation-with-highest-score)|Hard|||||||||||
|[0799](https://leetcode.cn/problems/champagne-tower)|Medium|||||||||||
|[0705](https://leetcode.cn/problems/design-hashset)|Easy|||||||||||
|[0706](https://leetcode.cn/problems/design-hashmap)|Easy|||||||||||
|[0800](https://leetcode.cn/problems/similar-rgb-color)|Easy|||||||||||
|[0801](https://leetcode.cn/problems/minimum-swaps-to-make-sequences-increasing)|Hard|||||||||||
|[0802](https://leetcode.cn/problems/find-eventual-safe-states)|Medium|||||||||||
|[0803](https://leetcode.cn/problems/bricks-falling-when-hit)|Hard|||||||||||
|[0804](https://leetcode.cn/problems/unique-morse-code-words)|Easy|||||||||||
|[0805](https://leetcode.cn/problems/split-array-with-same-average)|Hard|||||||||||
|[0806](https://leetcode.cn/problems/number-of-lines-to-write-string)|Easy|||||||||||
|[0807](https://leetcode.cn/problems/max-increase-to-keep-city-skyline)|Medium|||||||||||
|[0808](https://leetcode.cn/problems/soup-servings)|Medium|||||||||||
|[0809](https://leetcode.cn/problems/expressive-words)|Medium|||||||||||
|[0810](https://leetcode.cn/problems/chalkboard-xor-game)|Hard|||||||||||
|[0811](https://leetcode.cn/problems/subdomain-visit-count)|Medium|||||||||||
|[0812](https://leetcode.cn/problems/largest-triangle-area)|Easy|||||||||||
|[0813](https://leetcode.cn/problems/largest-sum-of-averages)|Medium|||||||||||
|[0814](https://leetcode.cn/problems/binary-tree-pruning)|Medium|||||||||||
|[0815](https://leetcode.cn/problems/bus-routes)|Hard|||||||||||
|[0816](https://leetcode.cn/problems/ambiguous-coordinates)|Medium|||||||||||
|[0817](https://leetcode.cn/problems/linked-list-components)|Medium|||||||||||
|[0818](https://leetcode.cn/problems/race-car)|Hard|||||||||||
|[0819](https://leetcode.cn/problems/most-common-word)|Easy|||||||||||
|[0707](https://leetcode.cn/problems/design-linked-list)|Medium|||||||||||
|[0820](https://leetcode.cn/problems/short-encoding-of-words)|Medium|||||||||||
|[0821](https://leetcode.cn/problems/shortest-distance-to-a-character)|Easy|||||||||||
|[0822](https://leetcode.cn/problems/card-flipping-game)|Medium|||||||||||
|[0823](https://leetcode.cn/problems/binary-trees-with-factors)|Medium|||||||||||
|[0708](https://leetcode.cn/problems/insert-into-a-sorted-circular-linked-list)|Medium|||||||||||
|[0824](https://leetcode.cn/problems/goat-latin)|Easy|||||||||||
|[0825](https://leetcode.cn/problems/friends-of-appropriate-ages)|Medium|||||||||||
|[0826](https://leetcode.cn/problems/most-profit-assigning-work)|Medium|||||||||||
|[0827](https://leetcode.cn/problems/making-a-large-island)|Hard|||||||||||
|[0828](https://leetcode.cn/problems/count-unique-characters-of-all-substrings-of-a-given-string)|Hard|||||||||||
|[0829](https://leetcode.cn/problems/consecutive-numbers-sum)|Hard|||||||||||
|[0830](https://leetcode.cn/problems/positions-of-large-groups)|Easy|||||||||||
|[0831](https://leetcode.cn/problems/masking-personal-information)|Medium|||||||||||
|[0641](https://leetcode.cn/problems/design-circular-deque)|Medium|||||||||||
|[0622](https://leetcode.cn/problems/design-circular-queue)|Medium|||||||||||
|[0832](https://leetcode.cn/problems/flipping-an-image)|Easy|||||||||||
|[0833](https://leetcode.cn/problems/find-and-replace-in-string)|Medium|||||||||||
|[0834](https://leetcode.cn/problems/sum-of-distances-in-tree)|Hard|||||||||||
|[0835](https://leetcode.cn/problems/image-overlap)|Medium|||||||||||
|[0489](https://leetcode.cn/problems/robot-room-cleaner)|Hard|||||||||||
|[0836](https://leetcode.cn/problems/rectangle-overlap)|Easy|||||||||||
|[0837](https://leetcode.cn/problems/new-21-game)|Medium|||||||||||
|[0838](https://leetcode.cn/problems/push-dominoes)|Medium|||||||||||
|[0839](https://leetcode.cn/problems/similar-string-groups)|Hard|||||||||||
|[0840](https://leetcode.cn/problems/magic-squares-in-grid)|Medium|||||||||||
|[0841](https://leetcode.cn/problems/keys-and-rooms)|Medium|||||||||||
|[0842](https://leetcode.cn/problems/split-array-into-fibonacci-sequence)|Medium|||||||||||
|[0843](https://leetcode.cn/problems/guess-the-word)|Hard|||||||||||
|[0844](https://leetcode.cn/problems/backspace-string-compare)|Easy|||||||||||
|[0845](https://leetcode.cn/problems/longest-mountain-in-array)|Medium|||||||||||
|[0846](https://leetcode.cn/problems/hand-of-straights)|Medium|||||||||||
|[0847](https://leetcode.cn/problems/shortest-path-visiting-all-nodes)|Hard|||||||||||
|[0848](https://leetcode.cn/problems/shifting-letters)|Medium|||||||||||
|[0849](https://leetcode.cn/problems/maximize-distance-to-closest-person)|Medium|||||||||||
|[0850](https://leetcode.cn/problems/rectangle-area-ii)|Hard|||||||||||
|[0851](https://leetcode.cn/problems/loud-and-rich)|Medium|||||||||||
|[0852](https://leetcode.cn/problems/peak-index-in-a-mountain-array)|Medium|||||||||||
|[0853](https://leetcode.cn/problems/car-fleet)|Medium|||||||||||
|[0854](https://leetcode.cn/problems/k-similar-strings)|Hard|||||||||||
|[0855](https://leetcode.cn/problems/exam-room)|Medium|||||||||||
|[0856](https://leetcode.cn/problems/score-of-parentheses)|Medium|||||||||||
|[0857](https://leetcode.cn/problems/minimum-cost-to-hire-k-workers)|Hard|||||||||||
|[0858](https://leetcode.cn/problems/mirror-reflection)|Medium|||||||||||
|[0859](https://leetcode.cn/problems/buddy-strings)|Easy|||||||||||
|[0860](https://leetcode.cn/problems/lemonade-change)|Easy|||||||||||
|[0861](https://leetcode.cn/problems/score-after-flipping-matrix)|Medium|||||||||||
|[0862](https://leetcode.cn/problems/shortest-subarray-with-sum-at-least-k)|Hard|||||||||||
|[0863](https://leetcode.cn/problems/all-nodes-distance-k-in-binary-tree)|Medium|||||||||||
|[0710](https://leetcode.cn/problems/random-pick-with-blacklist)|Hard|||||||||||
|[0864](https://leetcode.cn/problems/shortest-path-to-get-all-keys)|Hard|||||||||||
|[0865](https://leetcode.cn/problems/smallest-subtree-with-all-the-deepest-nodes)|Medium|||||||||||
|[0866](https://leetcode.cn/problems/prime-palindrome)|Medium|||||||||||
|[0867](https://leetcode.cn/problems/transpose-matrix)|Easy|||||||||||
|[0868](https://leetcode.cn/problems/binary-gap)|Easy|||||||||||
|[0869](https://leetcode.cn/problems/reordered-power-of-2)|Medium|||||||||||
|[0870](https://leetcode.cn/problems/advantage-shuffle)|Medium|||||||||||
|[0871](https://leetcode.cn/problems/minimum-number-of-refueling-stops)|Hard|||||||||||
|[0470](https://leetcode.cn/problems/implement-rand10-using-rand7)|Medium|||||||||||
|[0872](https://leetcode.cn/problems/leaf-similar-trees)|Easy|||||||||||
|[0873](https://leetcode.cn/problems/length-of-longest-fibonacci-subsequence)|Medium|||||||||||
|[0874](https://leetcode.cn/problems/walking-robot-simulation)|Medium|||||||||||
|[0875](https://leetcode.cn/problems/koko-eating-bananas)|Medium|||||||||||
|[0876](https://leetcode.cn/problems/middle-of-the-linked-list)|Easy|||||||||||
|[0877](https://leetcode.cn/problems/stone-game)|Medium|:white_check_mark:|:white_check_mark:|:white_check_mark:|:white_check_mark:|||:white_check_mark:|:white_check_mark:|||
|[0878](https://leetcode.cn/problems/nth-magical-number)|Hard|||||||||||
|[0879](https://leetcode.cn/problems/profitable-schemes)|Hard|||||||||||
|[0528](https://leetcode.cn/problems/random-pick-with-weight)|Medium|||||||||||
|[0519](https://leetcode.cn/problems/random-flip-matrix)|Medium|||||||||||
|[0497](https://leetcode.cn/problems/random-point-in-non-overlapping-rectangles)|Medium|||||||||||
|[0478](https://leetcode.cn/problems/generate-random-point-in-a-circle)|Medium|||||||||||
|[0880](https://leetcode.cn/problems/decoded-string-at-index)|Medium|||||||||||
|[0881](https://leetcode.cn/problems/boats-to-save-people)|Medium|||||||||||
|[0882](https://leetcode.cn/problems/reachable-nodes-in-subdivided-graph)|Hard|||||||||||
|[0883](https://leetcode.cn/problems/projection-area-of-3d-shapes)|Easy|||||||||||
|[0884](https://leetcode.cn/problems/uncommon-words-from-two-sentences)|Easy|||||||||||
|[0885](https://leetcode.cn/problems/spiral-matrix-iii)|Medium|||||||||||
|[0886](https://leetcode.cn/problems/possible-bipartition)|Medium|||||||||||
|[0887](https://leetcode.cn/problems/super-egg-drop)|Hard|||||||||||
|[0888](https://leetcode.cn/problems/fair-candy-swap)|Easy|||||||||||
|[0889](https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-postorder-traversal)|Medium|||||||||||
|[0890](https://leetcode.cn/problems/find-and-replace-pattern)|Medium|||||||||||
|[0891](https://leetcode.cn/problems/sum-of-subsequence-widths)|Hard|||||||||||
|[0892](https://leetcode.cn/problems/surface-area-of-3d-shapes)|Easy|||||||||||
|[0893](https://leetcode.cn/problems/groups-of-special-equivalent-strings)|Medium|||||||||||
|[0894](https://leetcode.cn/problems/all-possible-full-binary-trees)|Medium|||||||||||
|[0895](https://leetcode.cn/problems/maximum-frequency-stack)|Hard|||||||||||
|[0896](https://leetcode.cn/problems/monotonic-array)|Easy|||||||||||
|[0897](https://leetcode.cn/problems/increasing-order-search-tree)|Easy|||||||||||
|[0898](https://leetcode.cn/problems/bitwise-ors-of-subarrays)|Medium|||||||||||
|[0899](https://leetcode.cn/problems/orderly-queue)|Hard|||||||||||
|[0900](https://leetcode.cn/problems/rle-iterator)|Medium|||||||||||
|[0901](https://leetcode.cn/problems/online-stock-span)|Medium|||||||||||
|[0902](https://leetcode.cn/problems/numbers-at-most-n-given-digit-set)|Hard|||||||||||
|[0903](https://leetcode.cn/problems/valid-permutations-for-di-sequence)|Hard|||||||||||
|[0904](https://leetcode.cn/problems/fruit-into-baskets)|Medium|||||||||||
|[0905](https://leetcode.cn/problems/sort-array-by-parity)|Easy|||||||||||
|[0906](https://leetcode.cn/problems/super-palindromes)|Hard|||||||||||
|[0907](https://leetcode.cn/problems/sum-of-subarray-minimums)|Medium|||||||||||
|[0908](https://leetcode.cn/problems/smallest-range-i)|Easy|||||||||||
|[0909](https://leetcode.cn/problems/snakes-and-ladders)|Medium|||||||||||
|[0910](https://leetcode.cn/problems/smallest-range-ii)|Medium|||||||||||
|[0911](https://leetcode.cn/problems/online-election)|Medium|||||||||||
|[0912](https://leetcode.cn/problems/sort-an-array)|Medium|||||||||||
|[0913](https://leetcode.cn/problems/cat-and-mouse)|Hard|||||||||||
|[0914](https://leetcode.cn/problems/x-of-a-kind-in-a-deck-of-cards)|Easy|||||||||||
|[0915](https://leetcode.cn/problems/partition-array-into-disjoint-intervals)|Medium|||||||||||
|[0916](https://leetcode.cn/problems/word-subsets)|Medium|||||||||||
|[0917](https://leetcode.cn/problems/reverse-only-letters)|Easy|||||||||||
|[0918](https://leetcode.cn/problems/maximum-sum-circular-subarray)|Medium|||||||||||
|[0919](https://leetcode.cn/problems/complete-binary-tree-inserter)|Medium|||||||||||
|[0920](https://leetcode.cn/problems/number-of-music-playlists)|Hard|||||||||||
|[0921](https://leetcode.cn/problems/minimum-add-to-make-parentheses-valid)|Medium|||||||||||
|[0922](https://leetcode.cn/problems/sort-array-by-parity-ii)|Easy|:white_check_mark:|:white_check_mark:|||||||||
|[0923](https://leetcode.cn/problems/3sum-with-multiplicity)|Medium|||||||||||
|[0924](https://leetcode.cn/problems/minimize-malware-spread)|Hard|||||||||||
|[0925](https://leetcode.cn/problems/long-pressed-name)|Easy|||||||||||
|[0926](https://leetcode.cn/problems/flip-string-to-monotone-increasing)|Medium|||||||||||
|[0927](https://leetcode.cn/problems/three-equal-parts)|Hard|||||||||||
|[0928](https://leetcode.cn/problems/minimize-malware-spread-ii)|Hard|||||||||||
|[0929](https://leetcode.cn/problems/unique-email-addresses)|Easy|||||||||||
|[0930](https://leetcode.cn/problems/binary-subarrays-with-sum)|Medium|||||||||||
|[0931](https://leetcode.cn/problems/minimum-falling-path-sum)|Medium|||||||||||
|[0932](https://leetcode.cn/problems/beautiful-array)|Medium|||||||||||
|[0933](https://leetcode.cn/problems/number-of-recent-calls)|Easy|||||||||||
|[0934](https://leetcode.cn/problems/shortest-bridge)|Medium|||||||||||
|[0935](https://leetcode.cn/problems/knight-dialer)|Medium|||||||||||
|[0936](https://leetcode.cn/problems/stamping-the-sequence)|Hard|||||||||||
|[0937](https://leetcode.cn/problems/reorder-data-in-log-files)|Medium|||||||||||
|[0938](https://leetcode.cn/problems/range-sum-of-bst)|Easy|||||||||||
|[0939](https://leetcode.cn/problems/minimum-area-rectangle)|Medium|||||||||||
|[0940](https://leetcode.cn/problems/distinct-subsequences-ii)|Hard|||||||||||
|[0941](https://leetcode.cn/problems/valid-mountain-array)|Easy|||||||||||
|[0942](https://leetcode.cn/problems/di-string-match)|Easy|||||||||||
|[0943](https://leetcode.cn/problems/find-the-shortest-superstring)|Hard|||||||||||
|[0944](https://leetcode.cn/problems/delete-columns-to-make-sorted)|Easy|||||||||||
|[0945](https://leetcode.cn/problems/minimum-increment-to-make-array-unique)|Medium|||||||||||
|[0946](https://leetcode.cn/problems/validate-stack-sequences)|Medium|||||||||||
|[0947](https://leetcode.cn/problems/most-stones-removed-with-same-row-or-column)|Medium|||||||||||
|[0948](https://leetcode.cn/problems/bag-of-tokens)|Medium|||||||||||
|[0949](https://leetcode.cn/problems/largest-time-for-given-digits)|Medium|||||||||||
|[0950](https://leetcode.cn/problems/reveal-cards-in-increasing-order)|Medium|||||||||||
|[0951](https://leetcode.cn/problems/flip-equivalent-binary-trees)|Medium|||||||||||
|[0952](https://leetcode.cn/problems/largest-component-size-by-common-factor)|Hard|||||||||||
|[0953](https://leetcode.cn/problems/verifying-an-alien-dictionary)|Easy|||||||||||
|[0954](https://leetcode.cn/problems/array-of-doubled-pairs)|Medium|||||||||||
|[0955](https://leetcode.cn/problems/delete-columns-to-make-sorted-ii)|Medium|||||||||||
|[0956](https://leetcode.cn/problems/tallest-billboard)|Hard|||||||||||
|[0957](https://leetcode.cn/problems/prison-cells-after-n-days)|Medium|||||||||||
|[0958](https://leetcode.cn/problems/check-completeness-of-a-binary-tree)|Medium|||||||||||
|[0959](https://leetcode.cn/problems/regions-cut-by-slashes)|Medium|||||||||||
|[0960](https://leetcode.cn/problems/delete-columns-to-make-sorted-iii)|Hard|||||||||||
|[0961](https://leetcode.cn/problems/n-repeated-element-in-size-2n-array)|Easy|||||||||||
|[0962](https://leetcode.cn/problems/maximum-width-ramp)|Medium|||||||||||
|[0963](https://leetcode.cn/problems/minimum-area-rectangle-ii)|Medium|||||||||||
|[0964](https://leetcode.cn/problems/least-operators-to-express-number)|Hard|||||||||||
|[0965](https://leetcode.cn/problems/univalued-binary-tree)|Easy|||||||||||
|[0966](https://leetcode.cn/problems/vowel-spellchecker)|Medium|||||||||||
|[0967](https://leetcode.cn/problems/numbers-with-same-consecutive-differences)|Medium|||||||||||
|[0968](https://leetcode.cn/problems/binary-tree-cameras)|Hard|||||||||||
|[0969](https://leetcode.cn/problems/pancake-sorting)|Medium|||||||||||
|[0970](https://leetcode.cn/problems/powerful-integers)|Medium|||||||||||
|[0971](https://leetcode.cn/problems/flip-binary-tree-to-match-preorder-traversal)|Medium|||||||||||
|[0972](https://leetcode.cn/problems/equal-rational-numbers)|Hard|||||||||||
|[0509](https://leetcode.cn/problems/fibonacci-number)|Easy|||||||||||
|[0973](https://leetcode.cn/problems/k-closest-points-to-origin)|Medium||:white_check_mark:|||||||||
|[0974](https://leetcode.cn/problems/subarray-sums-divisible-by-k)|Medium|||||||||||
|[0975](https://leetcode.cn/problems/odd-even-jump)|Hard|||||||||||
|[0976](https://leetcode.cn/problems/largest-perimeter-triangle)|Easy|:white_check_mark:|:white_check_mark:|:white_check_mark:||||||||
|[0977](https://leetcode.cn/problems/squares-of-a-sorted-array)|Easy|||||||||||
|[0978](https://leetcode.cn/problems/longest-turbulent-subarray)|Medium|||||||||||
|[0979](https://leetcode.cn/problems/distribute-coins-in-binary-tree)|Medium|||||||||||
|[0980](https://leetcode.cn/problems/unique-paths-iii)|Hard|||||||||||
|[0981](https://leetcode.cn/problems/time-based-key-value-store)|Medium|||||||||||
|[0982](https://leetcode.cn/problems/triples-with-bitwise-and-equal-to-zero)|Hard|||||||||||
|[0983](https://leetcode.cn/problems/minimum-cost-for-tickets)|Medium|||||||||||
|[0984](https://leetcode.cn/problems/string-without-aaa-or-bbb)|Medium|||||||||||
|[0985](https://leetcode.cn/problems/sum-of-even-numbers-after-queries)|Medium|||||||||||
|[0986](https://leetcode.cn/problems/interval-list-intersections)|Medium|||||||||||
|[0987](https://leetcode.cn/problems/vertical-order-traversal-of-a-binary-tree)|Hard|||||||||||
|[0988](https://leetcode.cn/problems/smallest-string-starting-from-leaf)|Medium|||||||||||
|[0989](https://leetcode.cn/problems/add-to-array-form-of-integer)|Easy|||||||||||
|[0990](https://leetcode.cn/problems/satisfiability-of-equality-equations)|Medium|||||||||||
|[0991](https://leetcode.cn/problems/broken-calculator)|Medium|||||||||||
|[0992](https://leetcode.cn/problems/subarrays-with-k-different-integers)|Hard|||||||||||
|[0993](https://leetcode.cn/problems/cousins-in-binary-tree)|Easy|||||||||||
|[0994](https://leetcode.cn/problems/rotting-oranges)|Medium|||||||||||
|[0995](https://leetcode.cn/problems/minimum-number-of-k-consecutive-bit-flips)|Hard|||||||||||
|[0996](https://leetcode.cn/problems/number-of-squareful-arrays)|Hard|||||||||||
|[0997](https://leetcode.cn/problems/find-the-town-judge)|Easy|||||||||||
|[0998](https://leetcode.cn/problems/maximum-binary-tree-ii)|Medium|||||||||||
|[0999](https://leetcode.cn/problems/available-captures-for-rook)|Easy|||||||||||
|[1000](https://leetcode.cn/problems/minimum-cost-to-merge-stones)|Hard|||||||||||
|[1001](https://leetcode.cn/problems/grid-illumination)|Hard|||||||||||
|[1002](https://leetcode.cn/problems/find-common-characters)|Easy|||||||||||
|[1003](https://leetcode.cn/problems/check-if-word-is-valid-after-substitutions)|Medium|||||||||||
|[1004](https://leetcode.cn/problems/max-consecutive-ones-iii)|Medium|||||||||||
|[1005](https://leetcode.cn/problems/maximize-sum-of-array-after-k-negations)|Easy|||||||||||
|[1006](https://leetcode.cn/problems/clumsy-factorial)|Medium|||||||||||
|[1007](https://leetcode.cn/problems/minimum-domino-rotations-for-equal-row)|Medium|||||||||||
|[1008](https://leetcode.cn/problems/construct-binary-search-tree-from-preorder-traversal)|Medium|||||||||||
|[1055](https://leetcode.cn/problems/shortest-way-to-form-string)|Medium|||||||||||
|[1057](https://leetcode.cn/problems/campus-bikes)|Medium|||||||||||
|[1058](https://leetcode.cn/problems/minimize-rounding-error-to-meet-target)|Medium|||||||||||
|[1009](https://leetcode.cn/problems/complement-of-base-10-integer)|Easy|||||||||||
|[1010](https://leetcode.cn/problems/pairs-of-songs-with-total-durations-divisible-by-60)|Medium|||||||||||
|[1011](https://leetcode.cn/problems/capacity-to-ship-packages-within-d-days)|Medium|||||||||||
|[1012](https://leetcode.cn/problems/numbers-with-repeated-digits)|Hard|||||||||||
|[1061](https://leetcode.cn/problems/lexicographically-smallest-equivalent-string)|Medium|||||||||||
|[1060](https://leetcode.cn/problems/missing-element-in-sorted-array)|Medium|||||||||||
|[1062](https://leetcode.cn/problems/longest-repeating-substring)|Medium|||||||||||
|[1063](https://leetcode.cn/problems/number-of-valid-subarrays)|Hard|||||||||||
|[1013](https://leetcode.cn/problems/partition-array-into-three-parts-with-equal-sum)|Easy|||||||||||
|[1014](https://leetcode.cn/problems/best-sightseeing-pair)|Medium|||||||||||
|[1015](https://leetcode.cn/problems/smallest-integer-divisible-by-k)|Medium|||||||||||
|[1016](https://leetcode.cn/problems/binary-string-with-substrings-representing-1-to-n)|Medium|||||||||||
|[1064](https://leetcode.cn/problems/fixed-point)|Easy|||||||||||
|[1066](https://leetcode.cn/problems/campus-bikes-ii)|Medium|||||||||||
|[1067](https://leetcode.cn/problems/digit-count-in-range)|Hard|||||||||||
|[1056](https://leetcode.cn/problems/confusing-number)|Easy|||||||||||
|[1017](https://leetcode.cn/problems/convert-to-base-2)|Medium|||||||||||
|[1018](https://leetcode.cn/problems/binary-prefix-divisible-by-5)|Easy|||||||||||
|[1019](https://leetcode.cn/problems/next-greater-node-in-linked-list)|Medium|||||||||||
|[1020](https://leetcode.cn/problems/number-of-enclaves)|Medium|||||||||||
|[1086](https://leetcode.cn/problems/high-five)|Easy|||||||||||
|[1065](https://leetcode.cn/problems/index-pairs-of-a-string)|Easy|||||||||||
|[1087](https://leetcode.cn/problems/brace-expansion)|Medium|||||||||||
|[1088](https://leetcode.cn/problems/confusing-number-ii)|Hard|||||||||||
|[1021](https://leetcode.cn/problems/remove-outermost-parentheses)|Easy|||||||||||
|[1022](https://leetcode.cn/problems/sum-of-root-to-leaf-binary-numbers)|Easy|||||||||||
|[1023](https://leetcode.cn/problems/camelcase-matching)|Medium|||||||||||
|[1024](https://leetcode.cn/problems/video-stitching)|Medium|||||||||||
|[1085](https://leetcode.cn/problems/sum-of-digits-in-the-minimum-number)|Easy|||||||||||
|[1099](https://leetcode.cn/problems/two-sum-less-than-k)|Easy|||||||||||
|[1100](https://leetcode.cn/problems/find-k-length-substrings-with-no-repeated-characters)|Medium||:white_check_mark:|||||||||
|[1101](https://leetcode.cn/problems/the-earliest-moment-when-everyone-become-friends)|Medium|||||||||||
|[1025](https://leetcode.cn/problems/divisor-game)|Easy|||||||||||
|[1027](https://leetcode.cn/problems/longest-arithmetic-subsequence)|Medium|||||||||||
|[1118](https://leetcode.cn/problems/number-of-days-in-a-month)|Easy|||||||||||
|[1119](https://leetcode.cn/problems/remove-vowels-from-a-string)|Easy|||||||||||
|[1134](https://leetcode.cn/problems/armstrong-number)|Easy|||||||||||
|[1120](https://leetcode.cn/problems/maximum-average-subtree)|Medium|||||||||||
|[1026](https://leetcode.cn/problems/maximum-difference-between-node-and-ancestor)|Medium|||||||||||
|[1028](https://leetcode.cn/problems/recover-a-tree-from-preorder-traversal)|Hard|||||||||||
|[1030](https://leetcode.cn/problems/matrix-cells-in-distance-order)|Easy|||||||||||
|[1029](https://leetcode.cn/problems/two-city-scheduling)|Medium|||||||||||
|[1031](https://leetcode.cn/problems/maximum-sum-of-two-non-overlapping-subarrays)|Medium|||||||||||
|[1032](https://leetcode.cn/problems/stream-of-characters)|Hard|||||||||||
|[1133](https://leetcode.cn/problems/largest-unique-number)|Easy|||||||||||
|[1102](https://leetcode.cn/problems/path-with-maximum-minimum-value)|Medium|||||||||||
|[1135](https://leetcode.cn/problems/connecting-cities-with-minimum-cost)|Medium|||||||||||
|[1136](https://leetcode.cn/problems/parallel-courses)|Medium|||||||||||
|[1150](https://leetcode.cn/problems/check-if-a-number-is-majority-element-in-a-sorted-array)|Easy|||||||||||
|[1033](https://leetcode.cn/problems/moving-stones-until-consecutive)|Medium|||||||||||
|[1034](https://leetcode.cn/problems/coloring-a-border)|Medium|||||||||||
|[1035](https://leetcode.cn/problems/uncrossed-lines)|Medium|||||||||||
|[1036](https://leetcode.cn/problems/escape-a-large-maze)|Hard|||||||||||
|[1151](https://leetcode.cn/problems/minimum-swaps-to-group-all-1s-together)|Medium|||||||||||
|[1152](https://leetcode.cn/problems/analyze-user-website-visit-pattern)|Medium|||||||||||
|[1039](https://leetcode.cn/problems/minimum-score-triangulation-of-polygon)|Medium|||||||||||
|[1160](https://leetcode.cn/problems/find-words-that-can-be-formed-by-characters)|Easy|||||||||||
|[1040](https://leetcode.cn/problems/moving-stones-until-consecutive-ii)|Medium|||||||||||
|[1038](https://leetcode.cn/problems/binary-search-tree-to-greater-sum-tree)|Medium|||||||||||
|[1037](https://leetcode.cn/problems/valid-boomerang)|Easy|||||||||||
|[1161](https://leetcode.cn/problems/maximum-level-sum-of-a-binary-tree)|Medium|||||||||||
|[1162](https://leetcode.cn/problems/as-far-from-land-as-possible)|Medium|||||||||||
|[1121](https://leetcode.cn/problems/divide-array-into-increasing-sequences)|Hard|||||||||||
|[1041](https://leetcode.cn/problems/robot-bounded-in-circle)|Medium|||||||||||
|[1042](https://leetcode.cn/problems/flower-planting-with-no-adjacent)|Medium|||||||||||
|[1043](https://leetcode.cn/problems/partition-array-for-maximum-sum)|Medium|||||||||||
|[1044](https://leetcode.cn/problems/longest-duplicate-substring)|Hard|||||||||||
|[1165](https://leetcode.cn/problems/single-row-keyboard)|Easy||:white_check_mark:|||||||||
|[1153](https://leetcode.cn/problems/string-transforms-into-another-string)|Hard|||||||||||
|[1166](https://leetcode.cn/problems/design-file-system)|Medium|||||||||||
|[1167](https://leetcode.cn/problems/minimum-cost-to-connect-sticks)|Medium|||||||||||
|[1046](https://leetcode.cn/problems/last-stone-weight)|Easy|||||||||||
|[1047](https://leetcode.cn/problems/remove-all-adjacent-duplicates-in-string)|Easy|||||||||||
|[1048](https://leetcode.cn/problems/longest-string-chain)|Medium|||||||||||
|[1049](https://leetcode.cn/problems/last-stone-weight-ii)|Medium|||||||||||
|[1180](https://leetcode.cn/problems/count-substrings-with-only-one-distinct-letter)|Easy|||||||||||
|[1181](https://leetcode.cn/problems/before-and-after-puzzle)|Medium|||||||||||
|[1163](https://leetcode.cn/problems/last-substring-in-lexicographical-order)|Hard|||||||||||
|[1182](https://leetcode.cn/problems/shortest-distance-to-target-color)|Medium|||||||||||
|[1051](https://leetcode.cn/problems/height-checker)|Easy|||||||||||
|[1052](https://leetcode.cn/problems/grumpy-bookstore-owner)|Medium|||||||||||
|[1053](https://leetcode.cn/problems/previous-permutation-with-one-swap)|Medium|||||||||||
|[1054](https://leetcode.cn/problems/distant-barcodes)|Medium|||||||||||
|[1196](https://leetcode.cn/problems/how-many-apples-can-you-put-into-the-basket)|Easy|||||||||||
|[1197](https://leetcode.cn/problems/minimum-knight-moves)|Medium|||||||||||
|[1198](https://leetcode.cn/problems/find-smallest-common-element-in-all-rows)|Medium|||||||||||
|[1168](https://leetcode.cn/problems/optimize-water-distribution-in-a-village)|Hard|||||||||||
|[1074](https://leetcode.cn/problems/number-of-submatrices-that-sum-to-target)|Hard|||||||||||
|[1071](https://leetcode.cn/problems/greatest-common-divisor-of-strings)|Easy|||||||||||
|[1072](https://leetcode.cn/problems/flip-columns-for-maximum-number-of-equal-rows)|Medium|||||||||||
|[1073](https://leetcode.cn/problems/adding-two-negabinary-numbers)|Medium|||||||||||
|[1213](https://leetcode.cn/problems/intersection-of-three-sorted-arrays)|Easy|||||||||||
|[1214](https://leetcode.cn/problems/two-sum-bsts)|Medium|||||||||||
|[1215](https://leetcode.cn/problems/stepping-numbers)|Medium|||||||||||
|[1183](https://leetcode.cn/problems/maximum-number-of-ones)|Hard|||||||||||
|[1078](https://leetcode.cn/problems/occurrences-after-bigram)|Easy|||||||||||
|[1080](https://leetcode.cn/problems/insufficient-nodes-in-root-to-leaf-paths)|Medium|||||||||||
|[1081](https://leetcode.cn/problems/smallest-subsequence-of-distinct-characters)|Medium|||||||||||
|[1079](https://leetcode.cn/problems/letter-tile-possibilities)|Medium|||||||||||
|[1228](https://leetcode.cn/problems/missing-number-in-arithmetic-progression)|Easy|||||||||||
|[1229](https://leetcode.cn/problems/meeting-scheduler)|Medium|||||||||||
|[1230](https://leetcode.cn/problems/toss-strange-coins)|Medium|||||||||||
|[1199](https://leetcode.cn/problems/minimum-time-to-build-blocks)|Hard|||||||||||
|[1089](https://leetcode.cn/problems/duplicate-zeros)|Easy|||||||||||
|[1090](https://leetcode.cn/problems/largest-values-from-labels)|Medium|||||||||||
|[1092](https://leetcode.cn/problems/shortest-common-supersequence)|Hard|||||||||||
|[1091](https://leetcode.cn/problems/shortest-path-in-binary-matrix)|Medium|||||||||||
|[1243](https://leetcode.cn/problems/array-transformation)|Easy|||||||||||
|[1244](https://leetcode.cn/problems/design-a-leaderboard)|Medium|||||||||||
|[1245](https://leetcode.cn/problems/tree-diameter)|Medium|||||||||||
|[1216](https://leetcode.cn/problems/valid-palindrome-iii)|Hard|||||||||||
|[1093](https://leetcode.cn/problems/statistics-from-a-large-sample)|Medium|||||||||||
|[1094](https://leetcode.cn/problems/car-pooling)|Medium|||||||||||
|[1095](https://leetcode.cn/problems/find-in-mountain-array)|Hard|||||||||||
|[1096](https://leetcode.cn/problems/brace-expansion-ii)|Hard|||||||||||
|[1256](https://leetcode.cn/problems/encode-number)|Medium|||||||||||
|[1257](https://leetcode.cn/problems/smallest-common-region)|Medium|||||||||||
|[1258](https://leetcode.cn/problems/synonymous-sentences)|Medium|||||||||||
|[1231](https://leetcode.cn/problems/divide-chocolate)|Hard|||||||||||
|[1104](https://leetcode.cn/problems/path-in-zigzag-labelled-binary-tree)|Medium|||||||||||
|[1103](https://leetcode.cn/problems/distribute-candies-to-people)|Easy|||||||||||
|[1105](https://leetcode.cn/problems/filling-bookcase-shelves)|Medium|||||||||||
|[1106](https://leetcode.cn/problems/parsing-a-boolean-expression)|Hard||:white_check_mark:|||||||||
|[1271](https://leetcode.cn/problems/hexspeak)|Easy|||||||||||
|[1272](https://leetcode.cn/problems/remove-interval)|Medium|||||||||||
|[1273](https://leetcode.cn/problems/delete-tree-nodes)|Medium|||||||||||
|[1246](https://leetcode.cn/problems/palindrome-removal)|Hard|||||||||||
|[1108](https://leetcode.cn/problems/defanging-an-ip-address)|Easy|||||||||||
|[1109](https://leetcode.cn/problems/corporate-flight-bookings)|Medium|||||||||||
|[1110](https://leetcode.cn/problems/delete-nodes-and-return-forest)|Medium|||||||||||
|[1111](https://leetcode.cn/problems/maximum-nesting-depth-of-two-valid-parentheses-strings)|Medium|||||||||||
|[1619](https://leetcode.cn/problems/mean-of-array-after-removing-some-elements)|Easy|||||||||||
|[1286](https://leetcode.cn/problems/iterator-for-combination)|Medium|||||||||||
|[1291](https://leetcode.cn/problems/sequential-digits)|Medium|||||||||||
|[1259](https://leetcode.cn/problems/handshakes-that-dont-cross)|Hard|||||||||||
|[1122](https://leetcode.cn/problems/relative-sort-array)|Easy|||||||||||
|[1123](https://leetcode.cn/problems/lowest-common-ancestor-of-deepest-leaves)|Medium|||||||||||
|[1124](https://leetcode.cn/problems/longest-well-performing-interval)|Medium|||||||||||
|[1125](https://leetcode.cn/problems/smallest-sufficient-team)|Hard|||||||||||
|[1287](https://leetcode.cn/problems/element-appearing-more-than-25-in-sorted-array)|Easy|||||||||||
|[1288](https://leetcode.cn/problems/remove-covered-intervals)|Medium|||||||||||
|[1627](https://leetcode.cn/problems/graph-connectivity-with-threshold)|Hard|||||||||||
|[1289](https://leetcode.cn/problems/minimum-falling-path-sum-ii)|Hard|||||||||||
|[1128](https://leetcode.cn/problems/number-of-equivalent-domino-pairs)|Easy|||||||||||
|[1130](https://leetcode.cn/problems/minimum-cost-tree-from-leaf-values)|Medium|||||||||||
|[1129](https://leetcode.cn/problems/shortest-path-with-alternating-colors)|Medium|||||||||||
|[1131](https://leetcode.cn/problems/maximum-of-absolute-value-expression)|Medium|||||||||||
|[1299](https://leetcode.cn/problems/replace-elements-with-greatest-element-on-right-side)|Easy|||||||||||
|[1300](https://leetcode.cn/problems/sum-of-mutated-array-closest-to-target)|Medium|||||||||||
|[1274](https://leetcode.cn/problems/number-of-ships-in-a-rectangle)|Hard|||||||||||
|[1301](https://leetcode.cn/problems/number-of-paths-with-max-score)|Hard|||||||||||
|[1137](https://leetcode.cn/problems/n-th-tribonacci-number)|Easy|||||||||||
|[1138](https://leetcode.cn/problems/alphabet-board-path)|Medium|||||||||||
|[1139](https://leetcode.cn/problems/largest-1-bordered-square)|Medium|||||||||||
|[1140](https://leetcode.cn/problems/stone-game-ii)|Medium|||||||||||
|[1313](https://leetcode.cn/problems/decompress-run-length-encoded-list)|Easy|||||||||||
|[1314](https://leetcode.cn/problems/matrix-block-sum)|Medium|||||||||||
|[1315](https://leetcode.cn/problems/sum-of-nodes-with-even-valued-grandparent)|Medium|||||||||||
|[1316](https://leetcode.cn/problems/distinct-echo-substrings)|Hard|||||||||||
|[1144](https://leetcode.cn/problems/decrease-elements-to-make-array-zigzag)|Medium|||||||||||
|[1145](https://leetcode.cn/problems/binary-tree-coloring-game)|Medium|||||||||||
|[1146](https://leetcode.cn/problems/snapshot-array)|Medium|||||||||||
|[1143](https://leetcode.cn/problems/longest-common-subsequence)|Medium|||||||||||
|[1147](https://leetcode.cn/problems/longest-chunked-palindrome-decomposition)|Hard|||||||||||
|[1328](https://leetcode.cn/problems/break-a-palindrome)|Medium|||||||||||
|[1329](https://leetcode.cn/problems/sort-the-matrix-diagonally)|Medium|||||||||||
|[1302](https://leetcode.cn/problems/deepest-leaves-sum)|Medium|||||||||||
|[1330](https://leetcode.cn/problems/reverse-subarray-to-maximize-array-value)|Hard|||||||||||
|[1331](https://leetcode.cn/problems/rank-transform-of-an-array)|Easy|||||||||||
|[1632](https://leetcode.cn/problems/rank-transform-of-a-matrix)|Hard|||||||||||
|[1154](https://leetcode.cn/problems/day-of-the-year)|Easy|||||||||||
|[1156](https://leetcode.cn/problems/swap-for-longest-repeated-character-substring)|Medium|||||||||||
|[1157](https://leetcode.cn/problems/online-majority-element-in-subarray)|Hard|||||||||||
|[1155](https://leetcode.cn/problems/number-of-dice-rolls-with-target-sum)|Medium|||||||||||
|[1935](https://leetcode.cn/problems/maximum-number-of-words-you-can-type)|Easy|||||||||||
|[1171](https://leetcode.cn/problems/remove-zero-sum-consecutive-nodes-from-linked-list)|Medium|||||||||||
|[1172](https://leetcode.cn/problems/dinner-plate-stacks)|Hard|||||||||||
|[1236](https://leetcode.cn/problems/web-crawler)|Medium|||||||||||
|[1169](https://leetcode.cn/problems/invalid-transactions)|Medium|||||||||||
|[1170](https://leetcode.cn/problems/compare-strings-by-frequency-of-the-smallest-character)|Medium|||||||||||
|[1360](https://leetcode.cn/problems/number-of-days-between-two-dates)|Easy|||||||||||
|[1361](https://leetcode.cn/problems/validate-binary-tree-nodes)|Medium|||||||||||
|[1362](https://leetcode.cn/problems/closest-divisors)|Medium|||||||||||
|[1363](https://leetcode.cn/problems/largest-multiple-of-three)|Hard|||||||||||
|[1175](https://leetcode.cn/problems/prime-arrangements)|Easy|||||||||||
|[1176](https://leetcode.cn/problems/diet-plan-performance)|Easy|||||||||||
|[1177](https://leetcode.cn/problems/can-make-palindrome-from-substring)|Medium|||||||||||
|[1178](https://leetcode.cn/problems/number-of-valid-words-for-each-puzzle)|Hard|||||||||||
|[1507](https://leetcode.cn/problems/reformat-date)|Easy|||||||||||
|[1390](https://leetcode.cn/problems/four-divisors)|Medium|||||||||||
|[1382](https://leetcode.cn/problems/balance-a-binary-search-tree)|Medium|||||||||||
|[1425](https://leetcode.cn/problems/constrained-subsequence-sum)|Hard|||||||||||
|[1184](https://leetcode.cn/problems/distance-between-bus-stops)|Easy|||||||||||
|[1186](https://leetcode.cn/problems/maximum-subarray-sum-with-one-deletion)|Medium|||||||||||
|[1185](https://leetcode.cn/problems/day-of-the-week)|Easy|||||||||||
|[1187](https://leetcode.cn/problems/make-array-strictly-increasing)|Hard|||||||||||
|[1550](https://leetcode.cn/problems/three-consecutive-odds)|Easy|||||||||||
|[2080](https://leetcode.cn/problems/range-frequency-queries)|Medium|||||||||||
|[1954](https://leetcode.cn/problems/minimum-garden-perimeter-to-collect-enough-apples)|Medium|||||||||||
|[1483](https://leetcode.cn/problems/kth-ancestor-of-a-tree-node)|Hard|||||||||||
|[1189](https://leetcode.cn/problems/maximum-number-of-balloons)|Easy|||||||||||
|[1190](https://leetcode.cn/problems/reverse-substrings-between-each-pair-of-parentheses)|Medium|||||||||||
|[1191](https://leetcode.cn/problems/k-concatenation-maximum-sum)|Medium|||||||||||
|[1192](https://leetcode.cn/problems/critical-connections-in-a-network)|Hard|||||||||||
|[1957](https://leetcode.cn/problems/delete-characters-to-make-fancy-string)|Easy|||||||||||
|[2139](https://leetcode.cn/problems/minimum-moves-to-reach-target-score)|Medium|||||||||||
|[1405](https://leetcode.cn/problems/longest-happy-string)|Medium|||||||||||
|[1944](https://leetcode.cn/problems/number-of-visible-people-in-a-queue)|Hard|||||||||||
|[1200](https://leetcode.cn/problems/minimum-absolute-difference)|Easy|||||||||||
|[1201](https://leetcode.cn/problems/ugly-number-iii)|Medium|||||||||||
|[1202](https://leetcode.cn/problems/smallest-string-with-swaps)|Medium|||||||||||
|[1203](https://leetcode.cn/problems/sort-items-by-groups-respecting-dependencies)|Hard|||||||||||
|[2079](https://leetcode.cn/problems/watering-plants)|Medium|||||||||||
|[1895](https://leetcode.cn/problems/largest-magic-square)|Medium|||||||||||
|[2201](https://leetcode.cn/problems/count-artifacts-that-can-be-extracted)|Medium|||||||||||
|[1916](https://leetcode.cn/problems/count-ways-to-build-rooms-in-an-ant-colony)|Hard|||||||||||
|[1207](https://leetcode.cn/problems/unique-number-of-occurrences)|Easy|||||||||||
|[1209](https://leetcode.cn/problems/remove-all-adjacent-duplicates-in-string-ii)|Medium|||||||||||
|[1208](https://leetcode.cn/problems/get-equal-substrings-within-budget)|Medium|||||||||||
|[1210](https://leetcode.cn/problems/minimum-moves-to-reach-target-with-rotations)|Hard|||||||||||
|[1706](https://leetcode.cn/problems/where-will-the-ball-fall)|Medium|||||||||||
|[1514](https://leetcode.cn/problems/path-with-maximum-probability)|Medium|||||||||||
|[1862](https://leetcode.cn/problems/sum-of-floored-pairs)|Hard|||||||||||
|[1217](https://leetcode.cn/problems/minimum-cost-to-move-chips-to-the-same-position)|Easy|||||||||||
|[1218](https://leetcode.cn/problems/longest-arithmetic-subsequence-of-given-difference)|Medium|||||||||||
|[1219](https://leetcode.cn/problems/path-with-maximum-gold)|Medium|||||||||||
|[1220](https://leetcode.cn/problems/count-vowels-permutation)|Hard|||||||||||
|[2191](https://leetcode.cn/problems/sort-the-jumbled-numbers)|Medium|||||||||||
|[2310](https://leetcode.cn/problems/sum-of-numbers-with-units-digit-k)|Medium|||||||||||
|[2226](https://leetcode.cn/problems/maximum-candies-allocated-to-k-children)|Medium|||||||||||
|[1960](https://leetcode.cn/problems/maximum-product-of-the-length-of-two-palindromic-substrings)|Hard|||||||||||
|[1206](https://leetcode.cn/problems/design-skiplist)|Hard|||||||||||
|[1221](https://leetcode.cn/problems/split-a-string-in-balanced-strings)|Easy|||||||||||
|[1222](https://leetcode.cn/problems/queens-that-can-attack-the-king)|Medium|||||||||||
|[1223](https://leetcode.cn/problems/dice-roll-simulation)|Hard|||||||||||
|[1224](https://leetcode.cn/problems/maximum-equal-frequency)|Hard|||||||||||
|[1427](https://leetcode.cn/problems/perform-string-shifts)|Easy|||||||||||
|[2202](https://leetcode.cn/problems/maximize-the-topmost-element-after-k-moves)|Medium|||||||||||
|[2204](https://leetcode.cn/problems/distance-to-a-cycle-in-undirected-graph)|Hard|||||||||||
|[2321](https://leetcode.cn/problems/maximum-score-of-spliced-array)|Hard|||||||||||
|[1232](https://leetcode.cn/problems/check-if-it-is-a-straight-line)|Easy|||||||||||
|[1233](https://leetcode.cn/problems/remove-sub-folders-from-the-filesystem)|Medium|||||||||||
|[1234](https://leetcode.cn/problems/replace-the-substring-for-balanced-string)|Medium|||||||||||
|[1235](https://leetcode.cn/problems/maximum-profit-in-job-scheduling)|Hard|||||||||||
|[2273](https://leetcode.cn/problems/find-resultant-array-after-removing-anagrams)|Easy|||||||||||
|[2225](https://leetcode.cn/problems/find-players-with-zero-or-one-losses)|Medium|||||||||||
|[2216](https://leetcode.cn/problems/minimum-deletions-to-make-array-beautiful)|Medium|||||||||||
|[2193](https://leetcode.cn/problems/minimum-number-of-moves-to-make-palindrome)|Hard|||||||||||
|[1237](https://leetcode.cn/problems/find-positive-integer-solution-for-a-given-equation)|Medium|||||||||||
|[1238](https://leetcode.cn/problems/circular-permutation-in-binary-representation)|Medium|||||||||||
|[1239](https://leetcode.cn/problems/maximum-length-of-a-concatenated-string-with-unique-characters)|Medium|||||||||||
|[1240](https://leetcode.cn/problems/tiling-a-rectangle-with-the-fewest-squares)|Hard|||||||||||
|[1227](https://leetcode.cn/problems/airplane-seat-assignment-probability)|Medium|||||||||||
|[2309](https://leetcode.cn/problems/greatest-english-letter-in-upper-and-lower-case)|Easy|||||||||||
|[1726](https://leetcode.cn/problems/tuple-with-same-product)|Medium|||||||||||
|[1429](https://leetcode.cn/problems/first-unique-number)|Medium|||||||||||
|[1691](https://leetcode.cn/problems/maximum-height-by-stacking-cuboids)|Hard|||||||||||
|[1247](https://leetcode.cn/problems/minimum-swaps-to-make-strings-equal)|Medium|||||||||||
|[1248](https://leetcode.cn/problems/count-number-of-nice-subarrays)|Medium|||||||||||
|[1249](https://leetcode.cn/problems/minimum-remove-to-make-valid-parentheses)|Medium|||||||||||
|[1250](https://leetcode.cn/problems/check-if-it-is-a-good-array)|Hard|||||||||||
|[1428](https://leetcode.cn/problems/leftmost-column-with-at-least-a-one)|Medium|||||||||||
|[2217](https://leetcode.cn/problems/find-palindrome-with-fixed-length)|Medium|||||||||||
|[2312](https://leetcode.cn/problems/selling-pieces-of-wood)|Hard|||||||||||
|[1252](https://leetcode.cn/problems/cells-with-odd-values-in-a-matrix)|Easy|||||||||||
|[1253](https://leetcode.cn/problems/reconstruct-a-2-row-binary-matrix)|Medium|||||||||||
|[1254](https://leetcode.cn/problems/number-of-closed-islands)|Medium|||||||||||
|[1255](https://leetcode.cn/problems/maximum-score-words-formed-by-letters)|Hard|||||||||||
|[2303](https://leetcode.cn/problems/calculate-amount-paid-in-taxes)|Easy|||||||||||
|[2198](https://leetcode.cn/problems/number-of-single-divisor-triplets)|Medium|||||||||||
|[1618](https://leetcode.cn/problems/maximum-font-to-fit-a-sentence-in-a-screen)|Medium|||||||||||
|[2189](https://leetcode.cn/problems/number-of-ways-to-build-house-of-cards)|Medium|||||||||||
|[1260](https://leetcode.cn/problems/shift-2d-grid)|Easy|||||||||||
|[1261](https://leetcode.cn/problems/find-elements-in-a-contaminated-binary-tree)|Medium|||||||||||
|[1262](https://leetcode.cn/problems/greatest-sum-divisible-by-three)|Medium|||||||||||
|[1263](https://leetcode.cn/problems/minimum-moves-to-move-a-box-to-their-target-location)|Hard|||||||||||
|[1426](https://leetcode.cn/problems/counting-elements)|Easy|||||||||||
|[2215](https://leetcode.cn/problems/find-the-difference-of-two-arrays)|Easy|||||||||||
|[2218](https://leetcode.cn/problems/maximum-value-of-k-coins-from-piles)|Hard|||||||||||
|[2304](https://leetcode.cn/problems/minimum-path-cost-in-a-grid)|Medium|||||||||||
|[1266](https://leetcode.cn/problems/minimum-time-visiting-all-points)|Easy|||||||||||
|[1267](https://leetcode.cn/problems/count-servers-that-communicate)|Medium|||||||||||
|[1268](https://leetcode.cn/problems/search-suggestions-system)|Medium|||||||||||
|[1269](https://leetcode.cn/problems/number-of-ways-to-stay-in-the-same-place-after-some-steps)|Hard|||||||||||
|[1275](https://leetcode.cn/problems/find-winner-on-a-tic-tac-toe-game)|Easy|||||||||||
|[1276](https://leetcode.cn/problems/number-of-burgers-with-no-waste-of-ingredients)|Medium|||||||||||
|[1277](https://leetcode.cn/problems/count-square-submatrices-with-all-ones)|Medium|||||||||||
|[1278](https://leetcode.cn/problems/palindrome-partitioning-iii)|Hard|||||||||||
|[1265](https://leetcode.cn/problems/print-immutable-linked-list-in-reverse)|Medium|||||||||||
|[1281](https://leetcode.cn/problems/subtract-the-product-and-sum-of-digits-of-an-integer)|Easy|||||||||||
|[1282](https://leetcode.cn/problems/group-the-people-given-the-group-size-they-belong-to)|Medium|||||||||||
|[1283](https://leetcode.cn/problems/find-the-smallest-divisor-given-a-threshold)|Medium|||||||||||
|[1284](https://leetcode.cn/problems/minimum-number-of-flips-to-convert-binary-matrix-to-zero-matrix)|Hard|||||||||||
|[1290](https://leetcode.cn/problems/convert-binary-number-in-a-linked-list-to-integer)|Easy|||||||||||
|[1292](https://leetcode.cn/problems/maximum-side-length-of-a-square-with-sum-less-than-or-equal-to-threshold)|Medium|||||||||||
|[1293](https://leetcode.cn/problems/shortest-path-in-a-grid-with-obstacles-elimination)|Hard|||||||||||
|[2229](https://leetcode.cn/problems/check-if-an-array-is-consecutive)|Easy|||||||||||
|[2305](https://leetcode.cn/problems/fair-distribution-of-cookies)|Medium|||||||||||
|[2263](https://leetcode.cn/problems/make-array-non-decreasing-or-non-increasing)|Hard|||||||||||
|[1295](https://leetcode.cn/problems/find-numbers-with-even-number-of-digits)|Easy|||||||||||
|[1296](https://leetcode.cn/problems/divide-array-in-sets-of-k-consecutive-numbers)|Medium|||||||||||
|[1297](https://leetcode.cn/problems/maximum-number-of-occurrences-of-a-substring)|Medium|||||||||||
|[1298](https://leetcode.cn/problems/maximum-candies-you-can-get-from-boxes)|Hard|||||||||||
|[1304](https://leetcode.cn/problems/find-n-unique-integers-sum-up-to-zero)|Easy|||||||||||
|[1305](https://leetcode.cn/problems/all-elements-in-two-binary-search-trees)|Medium|||||||||||
|[1306](https://leetcode.cn/problems/jump-game-iii)|Medium|||||||||||
|[1307](https://leetcode.cn/problems/verbal-arithmetic-puzzle)|Hard|||||||||||
|[2269](https://leetcode.cn/problems/find-the-k-beauty-of-a-number)|Easy|||||||||||
|[2192](https://leetcode.cn/problems/all-ancestors-of-a-node-in-a-directed-acyclic-graph)|Medium|||||||||||
|[1430](https://leetcode.cn/problems/check-if-a-string-is-a-valid-sequence-from-root-to-leaves-path-in-a-binary-tree)|Medium|||||||||||
|[2227](https://leetcode.cn/problems/encrypt-and-decrypt-strings)|Hard|||||||||||
|[1309](https://leetcode.cn/problems/decrypt-string-from-alphabet-to-integer-mapping)|Easy||:white_check_mark:|||||||||
|[1310](https://leetcode.cn/problems/xor-queries-of-a-subarray)|Medium|||||||||||
|[1311](https://leetcode.cn/problems/get-watched-videos-by-your-friends)|Medium|||||||||||
|[1312](https://leetcode.cn/problems/minimum-insertion-steps-to-make-a-string-palindrome)|Hard|||||||||||
|[1317](https://leetcode.cn/problems/convert-integer-to-the-sum-of-two-no-zero-integers)|Easy|||||||||||
|[1318](https://leetcode.cn/problems/minimum-flips-to-make-a-or-b-equal-to-c)|Medium|||||||||||
|[1319](https://leetcode.cn/problems/number-of-operations-to-make-network-connected)|Medium|||||||||||
|[1320](https://leetcode.cn/problems/minimum-distance-to-type-a-word-using-two-fingers)|Hard|||||||||||
|[1342](https://leetcode.cn/problems/number-of-steps-to-reduce-a-number-to-zero)|Easy|||||||||||
|[1343](https://leetcode.cn/problems/number-of-sub-arrays-of-size-k-and-average-greater-than-or-equal-to-threshold)|Medium|||||||||||
|[1344](https://leetcode.cn/problems/angle-between-hands-of-a-clock)|Medium|||||||||||
|[1345](https://leetcode.cn/problems/jump-game-iv)|Hard|||||||||||
|[1323](https://leetcode.cn/problems/maximum-69-number)|Easy|||||||||||
|[1324](https://leetcode.cn/problems/print-words-vertically)|Medium|||||||||||
|[1325](https://leetcode.cn/problems/delete-leaves-with-a-given-value)|Medium|||||||||||
|[1326](https://leetcode.cn/problems/minimum-number-of-taps-to-open-to-water-a-garden)|Hard|||||||||||
|[1332](https://leetcode.cn/problems/remove-palindromic-subsequences)|Easy|||||||||||
|[1333](https://leetcode.cn/problems/filter-restaurants-by-vegan-friendly-price-and-distance)|Medium|||||||||||
|[1334](https://leetcode.cn/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance)|Medium|||||||||||
|[1335](https://leetcode.cn/problems/minimum-difficulty-of-a-job-schedule)|Hard|||||||||||
|[1356](https://leetcode.cn/problems/sort-integers-by-the-number-of-1-bits)|Easy|||||||||||
|[1357](https://leetcode.cn/problems/apply-discount-every-n-orders)|Medium|||||||||||
|[1358](https://leetcode.cn/problems/number-of-substrings-containing-all-three-characters)|Medium||:white_check_mark:|||||||||
|[1359](https://leetcode.cn/problems/count-all-valid-pickup-and-delivery-options)|Hard|||||||||||
|[1337](https://leetcode.cn/problems/the-k-weakest-rows-in-a-matrix)|Easy|||||||||||
|[1338](https://leetcode.cn/problems/reduce-array-size-to-the-half)|Medium|||||||||||
|[1339](https://leetcode.cn/problems/maximum-product-of-splitted-binary-tree)|Medium|||||||||||
|[1340](https://leetcode.cn/problems/jump-game-v)|Hard|||||||||||
|[1346](https://leetcode.cn/problems/check-if-n-and-its-double-exist)|Easy|||||||||||
|[1347](https://leetcode.cn/problems/minimum-number-of-steps-to-make-two-strings-anagram)|Medium|||||||||||
|[1348](https://leetcode.cn/problems/tweet-counts-per-frequency)|Medium|||||||||||
|[1349](https://leetcode.cn/problems/maximum-students-taking-exam)|Hard|||||||||||
|[1370](https://leetcode.cn/problems/increasing-decreasing-string)|Easy|||||||||||
|[1371](https://leetcode.cn/problems/find-the-longest-substring-containing-vowels-in-even-counts)|Medium|||||||||||
|[1372](https://leetcode.cn/problems/longest-zigzag-path-in-a-binary-tree)|Medium|||||||||||
|[1373](https://leetcode.cn/problems/maximum-sum-bst-in-binary-tree)|Hard|||||||||||
|[1351](https://leetcode.cn/problems/count-negative-numbers-in-a-sorted-matrix)|Easy|||||||||||
|[1352](https://leetcode.cn/problems/product-of-the-last-k-numbers)|Medium|||||||||||
|[1353](https://leetcode.cn/problems/maximum-number-of-events-that-can-be-attended)|Medium|||||||||||
|[1354](https://leetcode.cn/problems/construct-target-array-with-multiple-sums)|Hard|||||||||||
|[1365](https://leetcode.cn/problems/how-many-numbers-are-smaller-than-the-current-number)|Easy|||||||||||
|[1366](https://leetcode.cn/problems/rank-teams-by-votes)|Medium|||||||||||
|[1367](https://leetcode.cn/problems/linked-list-in-binary-tree)|Medium|||||||||||
|[1368](https://leetcode.cn/problems/minimum-cost-to-make-at-least-one-valid-path-in-a-grid)|Hard|||||||||||
|[1385](https://leetcode.cn/problems/find-the-distance-value-between-two-arrays)|Easy|||||||||||
|[1386](https://leetcode.cn/problems/cinema-seat-allocation)|Medium|||||||||||
|[1387](https://leetcode.cn/problems/sort-integers-by-the-power-value)|Medium|||||||||||
|[1388](https://leetcode.cn/problems/pizza-with-3n-slices)|Hard|||||||||||
|[1374](https://leetcode.cn/problems/generate-a-string-with-characters-that-have-odd-counts)|Easy|||||||||||
|[1375](https://leetcode.cn/problems/number-of-times-binary-string-is-prefix-aligned)|Medium|||||||||||
|[1376](https://leetcode.cn/problems/time-needed-to-inform-all-employees)|Medium|||||||||||
|[1377](https://leetcode.cn/problems/frog-position-after-t-seconds)|Hard|||||||||||
|[1380](https://leetcode.cn/problems/lucky-numbers-in-a-matrix)|Easy|||||||||||
|[1381](https://leetcode.cn/problems/design-a-stack-with-increment-operation)|Medium|||||||||||
|[1379](https://leetcode.cn/problems/find-a-corresponding-node-of-a-binary-tree-in-a-clone-of-that-tree)|Easy|||||||||||
|[1383](https://leetcode.cn/problems/maximum-performance-of-a-team)|Hard|||||||||||
|[1399](https://leetcode.cn/problems/count-largest-group)|Easy|||||||||||
|[1401](https://leetcode.cn/problems/circle-and-rectangle-overlapping)|Medium|||||||||||
|[1400](https://leetcode.cn/problems/construct-k-palindrome-strings)|Medium|||||||||||
|[1402](https://leetcode.cn/problems/reducing-dishes)|Hard|||||||||||
|[1389](https://leetcode.cn/problems/create-target-array-in-the-given-order)|Easy|||||||||||
|[1391](https://leetcode.cn/problems/check-if-there-is-a-valid-path-in-a-grid)|Medium|||||||||||
|[1392](https://leetcode.cn/problems/longest-happy-prefix)|Hard|||||||||||
|[1394](https://leetcode.cn/problems/find-lucky-integer-in-an-array)|Easy|||||||||||
|[1395](https://leetcode.cn/problems/count-number-of-teams)|Medium|||||||||||
|[1396](https://leetcode.cn/problems/design-underground-system)|Medium|||||||||||
|[1397](https://leetcode.cn/problems/find-all-good-strings)|Hard|||||||||||
|[1413](https://leetcode.cn/problems/minimum-value-to-get-positive-step-by-step-sum)|Easy|||||||||||
|[1414](https://leetcode.cn/problems/find-the-minimum-number-of-fibonacci-numbers-whose-sum-is-k)|Medium|||||||||||
|[1415](https://leetcode.cn/problems/the-k-th-lexicographical-string-of-all-happy-strings-of-length-n)|Medium|||||||||||
|[1416](https://leetcode.cn/problems/restore-the-array)|Hard|||||||||||
|[1403](https://leetcode.cn/problems/minimum-subsequence-in-non-increasing-order)|Easy|||||||||||
|[1404](https://leetcode.cn/problems/number-of-steps-to-reduce-a-number-in-binary-representation-to-one)|Medium|||||||||||
|[1406](https://leetcode.cn/problems/stone-game-iii)|Hard|||||||||||
|[1408](https://leetcode.cn/problems/string-matching-in-an-array)|Easy|||||||||||
|[1409](https://leetcode.cn/problems/queries-on-a-permutation-with-key)|Medium|||||||||||
|[1410](https://leetcode.cn/problems/html-entity-parser)|Medium|||||||||||
|[1411](https://leetcode.cn/problems/number-of-ways-to-paint-n-3-grid)|Hard|||||||||||
|[1431](https://leetcode.cn/problems/kids-with-the-greatest-number-of-candies)|Easy|||||||||||
|[1432](https://leetcode.cn/problems/max-difference-you-can-get-from-changing-an-integer)|Medium|||||||||||
|[1433](https://leetcode.cn/problems/check-if-a-string-can-break-another-string)|Medium|||||||||||
|[1434](https://leetcode.cn/problems/number-of-ways-to-wear-different-hats-to-each-other)|Hard|||||||||||
|[1417](https://leetcode.cn/problems/reformat-the-string)|Easy|||||||||||
|[1418](https://leetcode.cn/problems/display-table-of-food-orders-in-a-restaurant)|Medium|||||||||||
|[1419](https://leetcode.cn/problems/minimum-number-of-frogs-croaking)|Medium|||||||||||
|[1420](https://leetcode.cn/problems/build-array-where-you-can-find-the-maximum-exactly-k-comparisons)|Hard|||||||||||
|[1422](https://leetcode.cn/problems/maximum-score-after-splitting-a-string)|Easy|||||||||||
|[1423](https://leetcode.cn/problems/maximum-points-you-can-obtain-from-cards)|Medium|||||||||||
|[1424](https://leetcode.cn/problems/diagonal-traverse-ii)|Medium|||||||||||
|[1446](https://leetcode.cn/problems/consecutive-characters)|Easy|||||||||||
|[1447](https://leetcode.cn/problems/simplified-fractions)|Medium|||||||||||
|[1448](https://leetcode.cn/problems/count-good-nodes-in-binary-tree)|Medium|||||||||||
|[1449](https://leetcode.cn/problems/form-largest-integer-with-digits-that-add-up-to-target)|Hard|||||||||||
|[1436](https://leetcode.cn/problems/destination-city)|Easy|||||||||||
|[1437](https://leetcode.cn/problems/check-if-all-1s-are-at-least-length-k-places-away)|Easy|||||||||||
|[1438](https://leetcode.cn/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit)|Medium|||||||||||
|[1439](https://leetcode.cn/problems/find-the-kth-smallest-sum-of-a-matrix-with-sorted-rows)|Hard|||||||||||
|[1441](https://leetcode.cn/problems/build-an-array-with-stack-operations)|Medium|||||||||||
|[1442](https://leetcode.cn/problems/count-triplets-that-can-form-two-arrays-of-equal-xor)|Medium|||||||||||
|[1443](https://leetcode.cn/problems/minimum-time-to-collect-all-apples-in-a-tree)|Medium|||||||||||
|[1444](https://leetcode.cn/problems/number-of-ways-of-cutting-a-pizza)|Hard|||||||||||
|[1460](https://leetcode.cn/problems/make-two-arrays-equal-by-reversing-subarrays)|Easy|||||||||||
|[1461](https://leetcode.cn/problems/check-if-a-string-contains-all-binary-codes-of-size-k)|Medium|||||||||||
|[1462](https://leetcode.cn/problems/course-schedule-iv)|Medium|||||||||||
|[1463](https://leetcode.cn/problems/cherry-pickup-ii)|Hard|||||||||||
|[1450](https://leetcode.cn/problems/number-of-students-doing-homework-at-a-given-time)|Easy|||||||||||
|[1451](https://leetcode.cn/problems/rearrange-words-in-a-sentence)|Medium|||||||||||
|[1452](https://leetcode.cn/problems/people-whose-list-of-favorite-companies-is-not-a-subset-of-another-list)|Medium|||||||||||
|[1453](https://leetcode.cn/problems/maximum-number-of-darts-inside-of-a-circular-dartboard)|Hard|||||||||||
|[1455](https://leetcode.cn/problems/check-if-a-word-occurs-as-a-prefix-of-any-word-in-a-sentence)|Easy|||||||||||
|[1456](https://leetcode.cn/problems/maximum-number-of-vowels-in-a-substring-of-given-length)|Medium|||||||||||
|[1457](https://leetcode.cn/problems/pseudo-palindromic-paths-in-a-binary-tree)|Medium|||||||||||
|[1458](https://leetcode.cn/problems/max-dot-product-of-two-subsequences)|Hard|||||||||||
|[1475](https://leetcode.cn/problems/final-prices-with-a-special-discount-in-a-shop)|Easy|||||||||||
|[1478](https://leetcode.cn/problems/allocate-mailboxes)|Hard|||||||||||
|[1476](https://leetcode.cn/problems/subrectangle-queries)|Medium|||||||||||
|[1477](https://leetcode.cn/problems/find-two-non-overlapping-sub-arrays-each-with-target-sum)|Medium|||||||||||
|[1464](https://leetcode.cn/problems/maximum-product-of-two-elements-in-an-array)|Easy|||||||||||
|[1465](https://leetcode.cn/problems/maximum-area-of-a-piece-of-cake-after-horizontal-and-vertical-cuts)|Medium|||||||||||
|[1466](https://leetcode.cn/problems/reorder-routes-to-make-all-paths-lead-to-the-city-zero)|Medium|||||||||||
|[1467](https://leetcode.cn/problems/probability-of-a-two-boxes-having-the-same-number-of-distinct-balls)|Hard|||||||||||
|[1470](https://leetcode.cn/problems/shuffle-the-array)|Easy|||||||||||
|[1471](https://leetcode.cn/problems/the-k-strongest-values-in-an-array)|Medium|||||||||||
|[1472](https://leetcode.cn/problems/design-browser-history)|Medium|||||||||||
|[1473](https://leetcode.cn/problems/paint-house-iii)|Hard|||||||||||
|[1491](https://leetcode.cn/problems/average-salary-excluding-the-minimum-and-maximum-salary)|Easy|||||||||||
|[1492](https://leetcode.cn/problems/the-kth-factor-of-n)|Medium|||||||||||
|[1493](https://leetcode.cn/problems/longest-subarray-of-1s-after-deleting-one-element)|Medium|||||||||||
|[1494](https://leetcode.cn/problems/parallel-courses-ii)|Hard|||||||||||
|[1480](https://leetcode.cn/problems/running-sum-of-1d-array)|Easy|||||||||||
|[1481](https://leetcode.cn/problems/least-number-of-unique-integers-after-k-removals)|Medium|||||||||||
|[1482](https://leetcode.cn/problems/minimum-number-of-days-to-make-m-bouquets)|Medium|||||||||||
|[1469](https://leetcode.cn/problems/find-all-the-lonely-nodes)|Easy|||||||||||
|[1486](https://leetcode.cn/problems/xor-operation-in-an-array)|Easy|||||||||||
|[1487](https://leetcode.cn/problems/making-file-names-unique)|Medium|||||||||||
|[1488](https://leetcode.cn/problems/avoid-flood-in-the-city)|Medium|||||||||||
|[1489](https://leetcode.cn/problems/find-critical-and-pseudo-critical-edges-in-minimum-spanning-tree)|Hard|||||||||||
|[1508](https://leetcode.cn/problems/range-sum-of-sorted-subarray-sums)|Medium|||||||||||
|[1509](https://leetcode.cn/problems/minimum-difference-between-largest-and-smallest-value-in-three-moves)|Medium|||||||||||
|[1510](https://leetcode.cn/problems/stone-game-iv)|Hard|||||||||||
|[1474](https://leetcode.cn/problems/delete-n-nodes-after-m-nodes-of-a-linked-list)|Easy|||||||||||
|[1496](https://leetcode.cn/problems/path-crossing)|Easy|||||||||||
|[1497](https://leetcode.cn/problems/check-if-array-pairs-are-divisible-by-k)|Medium|||||||||||
|[1498](https://leetcode.cn/problems/number-of-subsequences-that-satisfy-the-given-sum-condition)|Medium|||||||||||
|[1499](https://leetcode.cn/problems/max-value-of-equation)|Hard|||||||||||
|[1485](https://leetcode.cn/problems/clone-binary-tree-with-random-pointer)|Medium|||||||||||
|[1502](https://leetcode.cn/problems/can-make-arithmetic-progression-from-sequence)|Easy|||||||||||
|[1503](https://leetcode.cn/problems/last-moment-before-all-ants-fall-out-of-a-plank)|Medium|||||||||||
|[1504](https://leetcode.cn/problems/count-submatrices-with-all-ones)|Medium|||||||||||
|[1505](https://leetcode.cn/problems/minimum-possible-integer-after-at-most-k-adjacent-swaps-on-digits)|Hard|||||||||||
|[1523](https://leetcode.cn/problems/count-odd-numbers-in-an-interval-range)|Easy|||||||||||
|[1524](https://leetcode.cn/problems/number-of-sub-arrays-with-odd-sum)|Medium|||||||||||
|[1525](https://leetcode.cn/problems/number-of-good-ways-to-split-a-string)|Medium|||||||||||
|[1526](https://leetcode.cn/problems/minimum-number-of-increments-on-subarrays-to-form-a-target-array)|Hard|||||||||||
|[1490](https://leetcode.cn/problems/clone-n-ary-tree)|Medium|||||||||||
|[1512](https://leetcode.cn/problems/number-of-good-pairs)|Easy|||||||||||
|[1513](https://leetcode.cn/problems/number-of-substrings-with-only-1s)|Medium|||||||||||
|[1531](https://leetcode.cn/problems/string-compression-ii)|Hard|||||||||||
|[1515](https://leetcode.cn/problems/best-position-for-a-service-centre)|Hard|||||||||||
|[1500](https://leetcode.cn/problems/design-a-file-sharing-system)|Medium|||||||||||
|[1518](https://leetcode.cn/problems/water-bottles)|Easy|||||||||||
|[1519](https://leetcode.cn/problems/number-of-nodes-in-the-sub-tree-with-the-same-label)|Medium|||||||||||
|[1520](https://leetcode.cn/problems/maximum-number-of-non-overlapping-substrings)|Hard|||||||||||
|[1521](https://leetcode.cn/problems/find-a-value-of-a-mysterious-function-closest-to-target)|Hard|||||||||||
|[1539](https://leetcode.cn/problems/kth-missing-positive-number)|Easy|||||||||||
|[1540](https://leetcode.cn/problems/can-convert-string-in-k-moves)|Medium|||||||||||
|[1541](https://leetcode.cn/problems/minimum-insertions-to-balance-a-parentheses-string)|Medium|||||||||||
|[1546](https://leetcode.cn/problems/maximum-number-of-non-overlapping-subarrays-with-sum-equals-target)|Medium|||||||||||
|[1506](https://leetcode.cn/problems/find-root-of-n-ary-tree)|Medium|||||||||||
|[1528](https://leetcode.cn/problems/shuffle-string)|Easy|||||||||||
|[1529](https://leetcode.cn/problems/minimum-suffix-flips)|Medium|||||||||||
|[1530](https://leetcode.cn/problems/number-of-good-leaf-nodes-pairs)|Medium|||||||||||
|[1516](https://leetcode.cn/problems/move-sub-tree-of-n-ary-tree)|Hard|||||||||||
|[1534](https://leetcode.cn/problems/count-good-triplets)|Easy|||||||||||
|[1535](https://leetcode.cn/problems/find-the-winner-of-an-array-game)|Medium|||||||||||
|[1536](https://leetcode.cn/problems/minimum-swaps-to-arrange-a-binary-grid)|Medium|||||||||||
|[1537](https://leetcode.cn/problems/get-the-maximum-score)|Hard|||||||||||
|[1556](https://leetcode.cn/problems/thousand-separator)|Easy|||||||||||
|[1557](https://leetcode.cn/problems/minimum-number-of-vertices-to-reach-all-nodes)|Medium|||||||||||
|[1558](https://leetcode.cn/problems/minimum-numbers-of-function-calls-to-make-target-array)|Medium|||||||||||
|[1559](https://leetcode.cn/problems/detect-cycles-in-2d-grid)|Medium|||||||||||
|[1522](https://leetcode.cn/problems/diameter-of-n-ary-tree)|Medium|||||||||||
|[1544](https://leetcode.cn/problems/make-the-string-great)|Easy|||||||||||
|[1545](https://leetcode.cn/problems/find-kth-bit-in-nth-binary-string)|Medium|||||||||||
|[1542](https://leetcode.cn/problems/find-longest-awesome-substring)|Hard|||||||||||
|[1547](https://leetcode.cn/problems/minimum-cost-to-cut-a-stick)|Hard|||||||||||
|[1533](https://leetcode.cn/problems/find-the-index-of-the-large-integer)|Medium|||||||||||
|[1551](https://leetcode.cn/problems/minimum-operations-to-make-array-equal)|Medium|||||||||||
|[1552](https://leetcode.cn/problems/magnetic-force-between-two-balls)|Medium|||||||||||
|[1553](https://leetcode.cn/problems/minimum-number-of-days-to-eat-n-oranges)|Hard|||||||||||
|[1572](https://leetcode.cn/problems/matrix-diagonal-sum)|Easy|||||||||||
|[1573](https://leetcode.cn/problems/number-of-ways-to-split-a-string)|Medium|||||||||||
|[1574](https://leetcode.cn/problems/shortest-subarray-to-be-removed-to-make-array-sorted)|Medium|||||||||||
|[1575](https://leetcode.cn/problems/count-all-possible-routes)|Hard|||||||||||
|[1538](https://leetcode.cn/problems/guess-the-majority-in-a-hidden-array)|Medium|||||||||||
|[1560](https://leetcode.cn/problems/most-visited-sector-in-a-circular-track)|Easy|||||||||||
|[1561](https://leetcode.cn/problems/maximum-number-of-coins-you-can-get)|Medium|||||||||||
|[1562](https://leetcode.cn/problems/find-latest-group-of-size-m)|Medium|||||||||||
|[1563](https://leetcode.cn/problems/stone-game-v)|Hard|||||||||||
|[1548](https://leetcode.cn/problems/the-most-similar-path-in-a-graph)|Hard|||||||||||
|[1566](https://leetcode.cn/problems/detect-pattern-of-length-m-repeated-k-or-more-times)|Easy|||||||||||
|[1567](https://leetcode.cn/problems/maximum-length-of-subarray-with-positive-product)|Medium|||||||||||
|[1568](https://leetcode.cn/problems/minimum-number-of-days-to-disconnect-island)|Hard|||||||||||
|[1569](https://leetcode.cn/problems/number-of-ways-to-reorder-array-to-get-same-bst)|Hard|||||||||||
|[1588](https://leetcode.cn/problems/sum-of-all-odd-length-subarrays)|Easy|||||||||||
|[1590](https://leetcode.cn/problems/make-sum-divisible-by-p)|Medium|||||||||||
|[1589](https://leetcode.cn/problems/maximum-sum-obtained-of-any-permutation)|Medium|||||||||||
|[1591](https://leetcode.cn/problems/strange-printer-ii)|Hard|||||||||||
|[1554](https://leetcode.cn/problems/strings-differ-by-one-character)|Medium|||||||||||
|[1576](https://leetcode.cn/problems/replace-all-s-to-avoid-consecutive-repeating-characters)|Easy|||||||||||
|[1577](https://leetcode.cn/problems/number-of-ways-where-square-of-number-is-equal-to-product-of-two-numbers)|Medium|||||||||||
|[1578](https://leetcode.cn/problems/minimum-time-to-make-rope-colorful)|Medium|||||||||||
|[1579](https://leetcode.cn/problems/remove-max-number-of-edges-to-keep-graph-fully-traversable)|Hard|||||||||||
|[1564](https://leetcode.cn/problems/put-boxes-into-the-warehouse-i)|Medium|||||||||||
|[1582](https://leetcode.cn/problems/special-positions-in-a-binary-matrix)|Easy|||||||||||
|[1583](https://leetcode.cn/problems/count-unhappy-friends)|Medium|||||||||||
|[1584](https://leetcode.cn/problems/min-cost-to-connect-all-points)|Medium|||||||||||
|[1585](https://leetcode.cn/problems/check-if-string-is-transformable-with-substring-sort-operations)|Hard|||||||||||
|[1603](https://leetcode.cn/problems/design-parking-system)|Easy|||||||||||
|[1604](https://leetcode.cn/problems/alert-using-same-key-card-three-or-more-times-in-a-one-hour-period)|Medium|||||||||||
|[1606](https://leetcode.cn/problems/find-servers-that-handled-most-number-of-requests)|Hard|||||||||||
|[1605](https://leetcode.cn/problems/find-valid-matrix-given-row-and-column-sums)|Medium|||||||||||
|[1570](https://leetcode.cn/problems/dot-product-of-two-sparse-vectors)|Medium|||||||||||
|[1592](https://leetcode.cn/problems/rearrange-spaces-between-words)|Easy|||||||||||
|[1593](https://leetcode.cn/problems/split-a-string-into-the-max-number-of-unique-substrings)|Medium|||||||||||
|[1594](https://leetcode.cn/problems/maximum-non-negative-product-in-a-matrix)|Medium|||||||||||
|[1595](https://leetcode.cn/problems/minimum-cost-to-connect-two-groups-of-points)|Hard|||||||||||
|[1580](https://leetcode.cn/problems/put-boxes-into-the-warehouse-ii)|Medium|||||||||||
|[1598](https://leetcode.cn/problems/crawler-log-folder)|Easy|||||||||||
|[1599](https://leetcode.cn/problems/maximum-profit-of-operating-a-centennial-wheel)|Medium|||||||||||
|[1600](https://leetcode.cn/problems/throne-inheritance)|Medium|||||||||||
|[1601](https://leetcode.cn/problems/maximum-number-of-achievable-transfer-requests)|Hard|||||||||||
|[1621](https://leetcode.cn/problems/number-of-sets-of-k-non-overlapping-line-segments)|Medium|||||||||||
|[1620](https://leetcode.cn/problems/coordinate-with-maximum-network-quality)|Medium|||||||||||
|[1728](https://leetcode.cn/problems/cat-and-mouse-ii)|Hard|||||||||||
|[1622](https://leetcode.cn/problems/fancy-sequence)|Hard|||||||||||
|[1586](https://leetcode.cn/problems/binary-search-tree-iterator-ii)|Medium|||||||||||
|[1608](https://leetcode.cn/problems/special-array-with-x-elements-greater-than-or-equal-x)|Easy|||||||||||
|[1609](https://leetcode.cn/problems/even-odd-tree)|Medium|||||||||||
|[1611](https://leetcode.cn/problems/minimum-one-bit-operations-to-make-integers-zero)|Hard|||||||||||
|[1610](https://leetcode.cn/problems/maximum-number-of-visible-points)|Hard|||||||||||
|[1597](https://leetcode.cn/problems/build-binary-expression-tree-from-infix-expression)|Hard|||||||||||
|[1614](https://leetcode.cn/problems/maximum-nesting-depth-of-the-parentheses)|Easy|||||||||||
|[1615](https://leetcode.cn/problems/maximal-network-rank)|Medium|||||||||||
|[1616](https://leetcode.cn/problems/split-two-strings-to-make-palindrome)|Medium|||||||||||
|[1617](https://leetcode.cn/problems/count-subtrees-with-max-distance-between-cities)|Hard|||||||||||
|[1636](https://leetcode.cn/problems/sort-array-by-increasing-frequency)|Easy|||||||||||
|[1637](https://leetcode.cn/problems/widest-vertical-area-between-two-points-containing-no-points)|Easy|||||||||||
|[1638](https://leetcode.cn/problems/count-substrings-that-differ-by-one-character)|Medium|||||||||||
|[1639](https://leetcode.cn/problems/number-of-ways-to-form-a-target-string-given-a-dictionary)|Hard|||||||||||
|[1602](https://leetcode.cn/problems/find-nearest-right-node-in-binary-tree)|Medium|||||||||||
|[1624](https://leetcode.cn/problems/largest-substring-between-two-equal-characters)|Easy|||||||||||
|[1625](https://leetcode.cn/problems/lexicographically-smallest-string-after-applying-operations)|Medium|||||||||||
|[1626](https://leetcode.cn/problems/best-team-with-no-conflicts)|Medium|||||||||||
|[1612](https://leetcode.cn/problems/check-if-two-expression-trees-are-equivalent)|Medium|||||||||||
|[1629](https://leetcode.cn/problems/slowest-key)|Easy|||||||||||
|[1630](https://leetcode.cn/problems/arithmetic-subarrays)|Medium|||||||||||
|[1631](https://leetcode.cn/problems/path-with-minimum-effort)|Medium|||||||||||
|[1652](https://leetcode.cn/problems/defuse-the-bomb)|Easy|||||||||||
|[1653](https://leetcode.cn/problems/minimum-deletions-to-make-string-balanced)|Medium|||||||||||
|[1654](https://leetcode.cn/problems/minimum-jumps-to-reach-home)|Medium|||||||||||
|[1655](https://leetcode.cn/problems/distribute-repeating-integers)|Hard|||||||||||
|[1640](https://leetcode.cn/problems/check-array-formation-through-concatenation)|Easy|||||||||||
|[1641](https://leetcode.cn/problems/count-sorted-vowel-strings)|Medium|||||||||||
|[1642](https://leetcode.cn/problems/furthest-building-you-can-reach)|Medium|||||||||||
|[1668](https://leetcode.cn/problems/maximum-repeating-substring)|Easy|||||||||||
|[1669](https://leetcode.cn/problems/merge-in-between-linked-lists)|Medium|||||||||||
|[1671](https://leetcode.cn/problems/minimum-number-of-removals-to-make-mountain-array)|Hard|||||||||||
|[1670](https://leetcode.cn/problems/design-front-middle-back-queue)|Medium|||||||||||
|[1628](https://leetcode.cn/problems/design-an-expression-tree-with-evaluate-function)|Medium|||||||||||
|[1646](https://leetcode.cn/problems/get-maximum-in-generated-array)|Easy|||||||||||
|[1647](https://leetcode.cn/problems/minimum-deletions-to-make-character-frequencies-unique)|Medium|||||||||||
|[1648](https://leetcode.cn/problems/sell-diminishing-valued-colored-balls)|Medium|||||||||||
|[1649](https://leetcode.cn/problems/create-sorted-array-through-instructions)|Hard|||||||||||
|[1634](https://leetcode.cn/problems/add-two-polynomials-represented-as-linked-lists)|Medium|||||||||||
|[1656](https://leetcode.cn/problems/design-an-ordered-stream)|Easy|||||||||||
|[1658](https://leetcode.cn/problems/minimum-operations-to-reduce-x-to-zero)|Medium|||||||||||
|[1657](https://leetcode.cn/problems/determine-if-two-strings-are-close)|Medium|||||||||||
|[1659](https://leetcode.cn/problems/maximize-grid-happiness)|Hard|||||||||||
|[1644](https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree-ii)|Medium|||||||||||
|[1662](https://leetcode.cn/problems/check-if-two-string-arrays-are-equivalent)|Easy|||||||||||
|[1663](https://leetcode.cn/problems/smallest-string-with-a-given-numeric-value)|Medium|||||||||||
|[1664](https://leetcode.cn/problems/ways-to-make-a-fair-array)|Medium|||||||||||
|[1665](https://leetcode.cn/problems/minimum-initial-energy-to-finish-tasks)|Hard|||||||||||
|[1684](https://leetcode.cn/problems/count-the-number-of-consistent-strings)|Easy|||||||||||
|[1685](https://leetcode.cn/problems/sum-of-absolute-differences-in-a-sorted-array)|Medium|||||||||||
|[1686](https://leetcode.cn/problems/stone-game-vi)|Medium|||||||||||
|[1687](https://leetcode.cn/problems/delivering-boxes-from-storage-to-ports)|Hard|||||||||||
|[1650](https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree-iii)|Medium|||||||||||
|[1672](https://leetcode.cn/problems/richest-customer-wealth)|Easy|||||||||||
|[1673](https://leetcode.cn/problems/find-the-most-competitive-subsequence)|Medium|||||||||||
|[1674](https://leetcode.cn/problems/minimum-moves-to-make-array-complementary)|Medium|||||||||||
|[1675](https://leetcode.cn/problems/minimize-deviation-in-array)|Hard|||||||||||
|[1660](https://leetcode.cn/problems/correct-a-binary-tree)|Medium|||||||||||
|[1678](https://leetcode.cn/problems/goal-parser-interpretation)|Easy|||||||||||
|[1679](https://leetcode.cn/problems/max-number-of-k-sum-pairs)|Medium|||||||||||
|[1681](https://leetcode.cn/problems/minimum-incompatibility)|Hard|||||||||||
|[1680](https://leetcode.cn/problems/concatenation-of-consecutive-binary-numbers)|Medium|||||||||||
|[1700](https://leetcode.cn/problems/number-of-students-unable-to-eat-lunch)|Easy|||||||||||
|[1701](https://leetcode.cn/problems/average-waiting-time)|Medium|||||||||||
|[1702](https://leetcode.cn/problems/maximum-binary-string-after-change)|Medium|||||||||||
|[1703](https://leetcode.cn/problems/minimum-adjacent-swaps-for-k-consecutive-ones)|Hard|||||||||||
|[1688](https://leetcode.cn/problems/count-of-matches-in-tournament)|Easy|||||||||||
|[1689](https://leetcode.cn/problems/partitioning-into-minimum-number-of-deci-binary-numbers)|Medium|||||||||||
|[1690](https://leetcode.cn/problems/stone-game-vii)|Medium|||||||||||
|[1714](https://leetcode.cn/problems/sum-of-special-evenly-spaced-elements-in-array)|Hard|||||||||||
|[1666](https://leetcode.cn/problems/change-the-root-of-a-binary-tree)|Medium|||||||||||
|[1694](https://leetcode.cn/problems/reformat-phone-number)|Easy|||||||||||
|[1695](https://leetcode.cn/problems/maximum-erasure-value)|Medium|||||||||||
|[1696](https://leetcode.cn/problems/jump-game-vi)|Medium|||||||||||
|[1697](https://leetcode.cn/problems/checking-existence-of-edge-length-limited-paths)|Hard|||||||||||
|[1676](https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree-iv)|Medium|||||||||||
|[1716](https://leetcode.cn/problems/calculate-money-in-leetcode-bank)|Easy|||||||||||
|[1717](https://leetcode.cn/problems/maximum-score-from-removing-substrings)|Medium|||||||||||
|[1718](https://leetcode.cn/problems/construct-the-lexicographically-largest-valid-sequence)|Medium|||||||||||
|[1719](https://leetcode.cn/problems/number-of-ways-to-reconstruct-a-tree)|Hard|||||||||||
|[1682](https://leetcode.cn/problems/longest-palindromic-subsequence-ii)|Medium|||||||||||
|[1704](https://leetcode.cn/problems/determine-if-string-halves-are-alike)|Easy|||||||||||
|[1705](https://leetcode.cn/problems/maximum-number-of-eaten-apples)|Medium|||||||||||
|[1723](https://leetcode.cn/problems/find-minimum-time-to-finish-all-jobs)|Hard|||||||||||
|[1707](https://leetcode.cn/problems/maximum-xor-with-an-element-from-array)|Hard|||||||||||
|[1692](https://leetcode.cn/problems/count-ways-to-distribute-candies)|Hard|||||||||||
|[1710](https://leetcode.cn/problems/maximum-units-on-a-truck)|Easy|||||||||||
|[1711](https://leetcode.cn/problems/count-good-meals)|Medium|||||||||||
|[1712](https://leetcode.cn/problems/ways-to-split-array-into-three-subarrays)|Medium|||||||||||
|[1713](https://leetcode.cn/problems/minimum-operations-to-make-a-subsequence)|Hard|||||||||||
|[1732](https://leetcode.cn/problems/find-the-highest-altitude)|Easy|||||||||||
|[1733](https://leetcode.cn/problems/minimum-number-of-people-to-teach)|Medium|||||||||||
|[1734](https://leetcode.cn/problems/decode-xored-permutation)|Medium|||||||||||
|[1735](https://leetcode.cn/problems/count-ways-to-make-array-with-product)|Hard|||||||||||
|[1698](https://leetcode.cn/problems/number-of-distinct-substrings-in-a-string)|Medium|||||||||||
|[1720](https://leetcode.cn/problems/decode-xored-array)|Easy|||||||||||
|[1722](https://leetcode.cn/problems/minimize-hamming-distance-after-swap-operations)|Medium|||||||||||
|[1725](https://leetcode.cn/problems/number-of-rectangles-that-can-form-the-largest-square)|Easy|||||||||||
|[1742](https://leetcode.cn/problems/maximum-number-of-balls-in-a-box)|Easy|||||||||||
|[1727](https://leetcode.cn/problems/largest-submatrix-with-rearrangements)|Medium|||||||||||
|[1708](https://leetcode.cn/problems/largest-subarray-length-k)|Easy|||||||||||
|[1748](https://leetcode.cn/problems/sum-of-unique-elements)|Easy|||||||||||
|[1749](https://leetcode.cn/problems/maximum-absolute-sum-of-any-subarray)|Medium|||||||||||
|[1750](https://leetcode.cn/problems/minimum-length-of-string-after-deleting-similar-ends)|Medium|||||||||||
|[1751](https://leetcode.cn/problems/maximum-number-of-events-that-can-be-attended-ii)|Hard|||||||||||
|[1736](https://leetcode.cn/problems/latest-time-by-replacing-hidden-digits)|Easy|||||||||||
|[1737](https://leetcode.cn/problems/change-minimum-characters-to-satisfy-one-of-three-conditions)|Medium|||||||||||
|[1738](https://leetcode.cn/problems/find-kth-largest-xor-coordinate-value)|Medium|||||||||||
|[1739](https://leetcode.cn/problems/building-boxes)|Hard|||||||||||
|[1724](https://leetcode.cn/problems/checking-existence-of-edge-length-limited-paths-ii)|Hard|||||||||||
|[1743](https://leetcode.cn/problems/restore-the-array-from-adjacent-pairs)|Medium|||||||||||
|[1745](https://leetcode.cn/problems/palindrome-partitioning-iv)|Hard|||||||||||
|[1744](https://leetcode.cn/problems/can-you-eat-your-favorite-candy-on-your-favorite-day)|Medium|||||||||||
|[1763](https://leetcode.cn/problems/longest-nice-substring)|Easy|||||||||||
|[1764](https://leetcode.cn/problems/form-array-by-concatenating-subarrays-of-another-array)|Medium|||||||||||
|[1766](https://leetcode.cn/problems/tree-of-coprimes)|Hard|||||||||||
|[1765](https://leetcode.cn/problems/map-of-highest-peak)|Medium|||||||||||
|[1752](https://leetcode.cn/problems/check-if-array-is-sorted-and-rotated)|Easy|||||||||||
|[1753](https://leetcode.cn/problems/maximum-score-from-removing-stones)|Medium|||||||||||
|[1754](https://leetcode.cn/problems/largest-merge-of-two-strings)|Medium|||||||||||
|[1755](https://leetcode.cn/problems/closest-subsequence-sum)|Hard|||||||||||
|[1740](https://leetcode.cn/problems/find-distance-in-a-binary-tree)|Medium|||||||||||
|[1758](https://leetcode.cn/problems/minimum-changes-to-make-alternating-binary-string)|Easy|||||||||||
|[1759](https://leetcode.cn/problems/count-number-of-homogenous-substrings)|Medium|||||||||||
|[1760](https://leetcode.cn/problems/minimum-limit-of-balls-in-a-bag)|Medium|||||||||||
|[1761](https://leetcode.cn/problems/minimum-degree-of-a-connected-trio-in-a-graph)|Hard|||||||||||
|[1779](https://leetcode.cn/problems/find-nearest-point-that-has-the-same-x-or-y-coordinate)|Easy|||||||||||
|[1780](https://leetcode.cn/problems/check-if-number-is-a-sum-of-powers-of-three)|Medium|||||||||||
|[1781](https://leetcode.cn/problems/sum-of-beauty-of-all-substrings)|Medium|||||||||||
|[1782](https://leetcode.cn/problems/count-pairs-of-nodes)|Hard|||||||||||
|[1746](https://leetcode.cn/problems/maximum-subarray-sum-after-one-operation)|Medium|||||||||||
|[1768](https://leetcode.cn/problems/merge-strings-alternately)|Easy|||||||||||
|[1769](https://leetcode.cn/problems/minimum-number-of-operations-to-move-all-balls-to-each-box)|Medium|||||||||||
|[1770](https://leetcode.cn/problems/maximum-score-from-performing-multiplication-operations)|Hard|||||||||||
|[1771](https://leetcode.cn/problems/maximize-palindrome-length-from-subsequences)|Hard|||||||||||
|[1773](https://leetcode.cn/problems/count-items-matching-a-rule)|Easy|||||||||||
|[1774](https://leetcode.cn/problems/closest-dessert-cost)|Medium|||||||||||
|[1775](https://leetcode.cn/problems/equal-sum-arrays-with-minimum-number-of-operations)|Medium|||||||||||
|[1776](https://leetcode.cn/problems/car-fleet-ii)|Hard|||||||||||
|[1756](https://leetcode.cn/problems/design-most-recently-used-queue)|Medium|||||||||||
|[1796](https://leetcode.cn/problems/second-largest-digit-in-a-string)|Easy|||||||||||
|[1797](https://leetcode.cn/problems/design-authentication-manager)|Medium|||||||||||
|[1799](https://leetcode.cn/problems/maximize-score-after-n-operations)|Hard|||||||||||
|[1803](https://leetcode.cn/problems/count-pairs-with-xor-in-a-range)|Hard|||||||||||
|[1762](https://leetcode.cn/problems/buildings-with-an-ocean-view)|Medium|||||||||||
|[1784](https://leetcode.cn/problems/check-if-binary-string-has-at-most-one-segment-of-ones)|Easy|||||||||||
|[1785](https://leetcode.cn/problems/minimum-elements-to-add-to-form-a-given-sum)|Medium|||||||||||
|[1786](https://leetcode.cn/problems/number-of-restricted-paths-from-first-to-last-node)|Medium|||||||||||
|[1787](https://leetcode.cn/problems/make-the-xor-of-all-segments-equal-to-zero)|Hard|||||||||||
|[1790](https://leetcode.cn/problems/check-if-one-string-swap-can-make-strings-equal)|Easy|||||||||||
|[1791](https://leetcode.cn/problems/find-center-of-star-graph)|Easy|||||||||||
|[1792](https://leetcode.cn/problems/maximum-average-pass-ratio)|Medium|||||||||||
|[1793](https://leetcode.cn/problems/maximum-score-of-a-good-subarray)|Hard|||||||||||
|[1772](https://leetcode.cn/problems/sort-features-by-popularity)|Medium|||||||||||
|[1812](https://leetcode.cn/problems/determine-color-of-a-chessboard-square)|Easy|||||||||||
|[1813](https://leetcode.cn/problems/sentence-similarity-iii)|Medium|||||||||||
|[1815](https://leetcode.cn/problems/maximum-number-of-groups-getting-fresh-donuts)|Hard|||||||||||
|[1814](https://leetcode.cn/problems/count-nice-pairs-in-an-array)|Medium|||||||||||
|[1800](https://leetcode.cn/problems/maximum-ascending-subarray-sum)|Easy|||||||||||
|[1801](https://leetcode.cn/problems/number-of-orders-in-the-backlog)|Medium|||||||||||
|[1802](https://leetcode.cn/problems/maximum-value-at-a-given-index-in-a-bounded-array)|Medium|||||||||||
|[1798](https://leetcode.cn/problems/maximum-number-of-consecutive-values-you-can-make)|Medium|||||||||||
|[1778](https://leetcode.cn/problems/shortest-path-in-a-hidden-grid)|Medium|||||||||||
|[1805](https://leetcode.cn/problems/number-of-different-integers-in-a-string)|Easy|||||||||||
|[1807](https://leetcode.cn/problems/evaluate-the-bracket-pairs-of-a-string)|Medium|||||||||||
|[1806](https://leetcode.cn/problems/minimum-number-of-operations-to-reinitialize-a-permutation)|Medium|||||||||||
|[1808](https://leetcode.cn/problems/maximize-number-of-nice-divisors)|Hard|||||||||||
|[1788](https://leetcode.cn/problems/maximize-the-beauty-of-the-garden)|Hard|||||||||||
|[1827](https://leetcode.cn/problems/minimum-operations-to-make-the-array-increasing)|Easy|||||||||||
|[1828](https://leetcode.cn/problems/queries-on-number-of-points-inside-a-circle)|Medium|||||||||||
|[1829](https://leetcode.cn/problems/maximum-xor-for-each-query)|Medium|||||||||||
|[1830](https://leetcode.cn/problems/minimum-number-of-operations-to-make-string-sorted)|Hard|||||||||||
|[1794](https://leetcode.cn/problems/count-pairs-of-equal-substrings-with-minimum-difference)|Medium|||||||||||
|[1816](https://leetcode.cn/problems/truncate-sentence)|Easy|||||||||||
|[1817](https://leetcode.cn/problems/finding-the-users-active-minutes)|Medium|||||||||||
|[1818](https://leetcode.cn/problems/minimum-absolute-sum-difference)|Medium|||||||||||
|[1819](https://leetcode.cn/problems/number-of-different-subsequences-gcds)|Hard|||||||||||
|[1804](https://leetcode.cn/problems/implement-trie-ii-prefix-tree)|Medium|||||||||||
|[1822](https://leetcode.cn/problems/sign-of-the-product-of-an-array)|Easy|||||||||||
|[1823](https://leetcode.cn/problems/find-the-winner-of-the-circular-game)|Medium|||||||||||
|[1824](https://leetcode.cn/problems/minimum-sideway-jumps)|Medium|||||||||||
|[1825](https://leetcode.cn/problems/finding-mk-average)|Hard|||||||||||
|[1844](https://leetcode.cn/problems/replace-all-digits-with-characters)|Easy|||||||||||
|[1845](https://leetcode.cn/problems/seat-reservation-manager)|Medium|||||||||||
|[1846](https://leetcode.cn/problems/maximum-element-after-decreasing-and-rearranging)|Medium|||||||||||
|[1847](https://leetcode.cn/problems/closest-room)|Hard|||||||||||
|[1810](https://leetcode.cn/problems/minimum-path-cost-in-a-hidden-grid)|Medium|||||||||||
|[1832](https://leetcode.cn/problems/check-if-the-sentence-is-pangram)|Easy|||||||||||
|[1833](https://leetcode.cn/problems/maximum-ice-cream-bars)|Medium|||||||||||
|[1834](https://leetcode.cn/problems/single-threaded-cpu)|Medium|||||||||||
|[1835](https://leetcode.cn/problems/find-xor-sum-of-all-pairs-bitwise-and)|Hard|||||||||||
|[1837](https://leetcode.cn/problems/sum-of-digits-in-base-k)|Easy|||||||||||
|[1838](https://leetcode.cn/problems/frequency-of-the-most-frequent-element)|Medium|||||||||||
|[1839](https://leetcode.cn/problems/longest-substring-of-all-vowels-in-order)|Medium|||||||||||
|[1840](https://leetcode.cn/problems/maximum-building-height)|Hard|||||||||||
|[1820](https://leetcode.cn/problems/maximum-number-of-accepted-invitations)|Medium|||||||||||
|[1859](https://leetcode.cn/problems/sorting-the-sentence)|Easy|||||||||||
|[1860](https://leetcode.cn/problems/incremental-memory-leak)|Medium|||||||||||
|[1861](https://leetcode.cn/problems/rotating-the-box)|Medium|||||||||||
|[1848](https://leetcode.cn/problems/minimum-distance-to-the-target-element)|Easy|||||||||||
|[1849](https://leetcode.cn/problems/splitting-a-string-into-descending-consecutive-values)|Medium|||||||||||
|[1851](https://leetcode.cn/problems/minimum-interval-to-include-each-query)|Hard|||||||||||
|[1850](https://leetcode.cn/problems/minimum-adjacent-swaps-to-reach-the-kth-smallest-number)|Medium|||||||||||
|[1989](https://leetcode.cn/problems/maximum-number-of-people-that-can-be-caught-in-tag)|Medium|||||||||||
|[1826](https://leetcode.cn/problems/faulty-sensor)|Easy|||||||||||
|[1836](https://leetcode.cn/problems/remove-duplicates-from-an-unsorted-linked-list)|Medium|||||||||||
|[1854](https://leetcode.cn/problems/maximum-population-year)|Easy|||||||||||
|[1855](https://leetcode.cn/problems/maximum-distance-between-a-pair-of-values)|Medium|||||||||||
|[1856](https://leetcode.cn/problems/maximum-subarray-min-product)|Medium|||||||||||
|[1857](https://leetcode.cn/problems/largest-color-value-in-a-directed-graph)|Hard|||||||||||
|[1876](https://leetcode.cn/problems/substrings-of-size-three-with-distinct-characters)|Easy|||||||||||
|[1877](https://leetcode.cn/problems/minimize-maximum-pair-sum-in-array)|Medium|||||||||||
|[1879](https://leetcode.cn/problems/minimum-xor-sum-of-two-arrays)|Hard|||||||||||
|[1878](https://leetcode.cn/problems/get-biggest-three-rhombus-sums-in-a-grid)|Medium|||||||||||
|[2046](https://leetcode.cn/problems/sort-linked-list-already-sorted-using-absolute-values)|Medium|||||||||||
|[1863](https://leetcode.cn/problems/sum-of-all-subset-xor-totals)|Easy|||||||||||
|[1864](https://leetcode.cn/problems/minimum-number-of-swaps-to-make-the-binary-string-alternating)|Medium|||||||||||
|[1865](https://leetcode.cn/problems/finding-pairs-with-a-certain-sum)|Medium|||||||||||
|[1866](https://leetcode.cn/problems/number-of-ways-to-rearrange-sticks-with-k-sticks-visible)|Hard|||||||||||
|[1842](https://leetcode.cn/problems/next-palindrome-using-same-digits)|Hard|||||||||||
|[1869](https://leetcode.cn/problems/longer-contiguous-segments-of-ones-than-zeros)|Easy|||||||||||
|[1870](https://leetcode.cn/problems/minimum-speed-to-arrive-on-time)|Medium|||||||||||
|[1871](https://leetcode.cn/problems/jump-game-vii)|Medium|||||||||||
|[1872](https://leetcode.cn/problems/stone-game-viii)|Hard|||||||||||
|[1852](https://leetcode.cn/problems/distinct-numbers-in-each-subarray)|Medium|||||||||||
|[1893](https://leetcode.cn/problems/check-if-all-the-integers-in-a-range-are-covered)|Easy|||||||||||
|[1894](https://leetcode.cn/problems/find-the-student-that-will-replace-the-chalk)|Medium|||||||||||
|[2247](https://leetcode.cn/problems/maximum-cost-of-trip-with-k-highways)|Hard|||||||||||
|[1896](https://leetcode.cn/problems/minimum-cost-to-change-the-final-value-of-expression)|Hard|||||||||||
|[1858](https://leetcode.cn/problems/longest-word-with-all-prefixes)|Medium|||||||||||
|[1880](https://leetcode.cn/problems/check-if-word-equals-summation-of-two-words)|Easy|||||||||||
|[1881](https://leetcode.cn/problems/maximum-value-after-insertion)|Medium|||||||||||
|[1882](https://leetcode.cn/problems/process-tasks-using-servers)|Medium|||||||||||
|[1883](https://leetcode.cn/problems/minimum-skips-to-arrive-at-meeting-on-time)|Hard|||||||||||
|[1886](https://leetcode.cn/problems/determine-whether-matrix-can-be-obtained-by-rotation)|Easy|||||||||||
|[1887](https://leetcode.cn/problems/reduction-operations-to-make-the-array-elements-equal)|Medium|||||||||||
|[1888](https://leetcode.cn/problems/minimum-number-of-flips-to-make-the-binary-string-alternating)|Medium|||||||||||
|[1889](https://leetcode.cn/problems/minimum-space-wasted-from-packaging)|Hard|||||||||||
|[1868](https://leetcode.cn/problems/product-of-two-run-length-encoded-arrays)|Medium|||||||||||
|[1909](https://leetcode.cn/problems/remove-one-element-to-make-the-array-strictly-increasing)|Easy|||||||||||
|[1910](https://leetcode.cn/problems/remove-all-occurrences-of-a-substring)|Medium|||||||||||
|[1911](https://leetcode.cn/problems/maximum-alternating-subsequence-sum)|Medium|||||||||||
|[1912](https://leetcode.cn/problems/design-movie-rental-system)|Hard|||||||||||
|[1897](https://leetcode.cn/problems/redistribute-characters-to-make-all-strings-equal)|Easy|||||||||||
|[1899](https://leetcode.cn/problems/merge-triplets-to-form-target-triplet)|Medium|||||||||||
|[1898](https://leetcode.cn/problems/maximum-number-of-removable-characters)|Medium|||||||||||
|[1900](https://leetcode.cn/problems/the-earliest-and-latest-rounds-where-players-compete)|Hard|||||||||||
|[1874](https://leetcode.cn/problems/minimize-product-sum-of-two-arrays)|Medium|||||||||||
|[1884](https://leetcode.cn/problems/egg-drop-with-2-eggs-and-n-floors)|Medium|||||||||||
|[1903](https://leetcode.cn/problems/largest-odd-number-in-string)|Easy|||||||||||
|[1904](https://leetcode.cn/problems/the-number-of-full-rounds-you-have-played)|Medium|||||||||||
|[1906](https://leetcode.cn/problems/minimum-absolute-difference-queries)|Medium|||||||||||
|[1905](https://leetcode.cn/problems/count-sub-islands)|Medium|||||||||||
|[1885](https://leetcode.cn/problems/count-pairs-in-two-arrays)|Medium|||||||||||
|[1925](https://leetcode.cn/problems/count-square-sum-triples)|Easy|||||||||||
|[1926](https://leetcode.cn/problems/nearest-exit-from-entrance-in-maze)|Medium|||||||||||
|[1927](https://leetcode.cn/problems/sum-game)|Medium|||||||||||
|[1928](https://leetcode.cn/problems/minimum-cost-to-reach-destination-in-time)|Hard|||||||||||
|[1913](https://leetcode.cn/problems/maximum-product-difference-between-two-pairs)|Easy|||||||||||
|[1914](https://leetcode.cn/problems/cyclically-rotating-a-grid)|Medium|||||||||||
|[1915](https://leetcode.cn/problems/number-of-wonderful-substrings)|Medium|||||||||||
|[1891](https://leetcode.cn/problems/cutting-ribbons)|Medium|||||||||||
|[1901](https://leetcode.cn/problems/find-a-peak-element-ii)|Medium|||||||||||
|[1920](https://leetcode.cn/problems/build-array-from-permutation)|Easy|||||||||||
|[1921](https://leetcode.cn/problems/eliminate-maximum-number-of-monsters)|Medium|||||||||||
|[1922](https://leetcode.cn/problems/count-good-numbers)|Medium|||||||||||
|[1923](https://leetcode.cn/problems/longest-common-subpath)|Hard|||||||||||
|[1902](https://leetcode.cn/problems/depth-of-bst-given-insertion-order)|Medium|||||||||||
|[1941](https://leetcode.cn/problems/check-if-all-characters-have-equal-number-of-occurrences)|Easy|||||||||||
|[1942](https://leetcode.cn/problems/the-number-of-the-smallest-unoccupied-chair)|Medium|||||||||||
|[1943](https://leetcode.cn/problems/describe-the-painting)|Medium|||||||||||
|[2297](https://leetcode.cn/problems/jump-game-viii)|Medium|||||||||||
|[1929](https://leetcode.cn/problems/concatenation-of-array)|Easy|||||||||||
|[1930](https://leetcode.cn/problems/unique-length-3-palindromic-subsequences)|Medium|||||||||||
|[1932](https://leetcode.cn/problems/merge-bsts-to-create-single-bst)|Hard|||||||||||
|[1931](https://leetcode.cn/problems/painting-a-grid-with-three-different-colors)|Hard|||||||||||
|[1908](https://leetcode.cn/problems/game-of-nim)|Medium|||||||||||
|[2307](https://leetcode.cn/problems/check-for-contradictions-in-equations)|Hard|||||||||||
|[1936](https://leetcode.cn/problems/add-minimum-number-of-rungs)|Medium|||||||||||
|[1937](https://leetcode.cn/problems/maximum-number-of-points-with-cost)|Medium|||||||||||
|[1938](https://leetcode.cn/problems/maximum-genetic-difference-query)|Hard|||||||||||
|[1918](https://leetcode.cn/problems/kth-smallest-subarray-sum)|Medium|||||||||||
|[1933](https://leetcode.cn/problems/check-if-string-is-decomposable-into-value-equal-substrings)|Easy|||||||||||
|[1940](https://leetcode.cn/problems/longest-common-subsequence-between-sorted-arrays)|Medium|||||||||||
|[1950](https://leetcode.cn/problems/maximum-of-minimum-values-in-all-subarrays)|Medium|||||||||||
|[1956](https://leetcode.cn/problems/minimum-time-for-k-virus-variants-to-spread)|Hard|||||||||||
|[1924](https://leetcode.cn/problems/erect-the-fence-ii)|Hard|||||||||||
|[2021](https://leetcode.cn/problems/brightest-position-on-street)|Medium|||||||||||
|[1945](https://leetcode.cn/problems/sum-of-digits-of-string-after-convert)|Easy|||||||||||
|[1946](https://leetcode.cn/problems/largest-number-after-mutating-substring)|Medium|||||||||||
|[1947](https://leetcode.cn/problems/maximum-compatibility-score-sum)|Medium|||||||||||
|[1948](https://leetcode.cn/problems/delete-duplicate-folders-in-system)|Hard|||||||||||
|[1958](https://leetcode.cn/problems/check-if-move-is-legal)|Medium|||||||||||
|[1959](https://leetcode.cn/problems/minimum-total-space-wasted-with-k-resizing-operations)|Medium|||||||||||
|[2052](https://leetcode.cn/problems/minimum-cost-to-separate-sentence-into-rows)|Medium|||||||||||
|[1952](https://leetcode.cn/problems/three-divisors)|Easy|||||||||||
|[1953](https://leetcode.cn/problems/maximum-number-of-weeks-for-which-you-can-work)|Medium|||||||||||
|[1968](https://leetcode.cn/problems/array-with-elements-not-equal-to-average-of-neighbors)|Medium|||||||||||
|[1955](https://leetcode.cn/problems/count-number-of-special-subsequences)|Hard|||||||||||
|[1974](https://leetcode.cn/problems/minimum-time-to-type-word-using-special-typewriter)|Easy|||||||||||
|[1975](https://leetcode.cn/problems/maximum-matrix-sum)|Medium|||||||||||
|[1976](https://leetcode.cn/problems/number-of-ways-to-arrive-at-destination)|Medium|||||||||||
|[1977](https://leetcode.cn/problems/number-of-ways-to-separate-numbers)|Hard|||||||||||
|[1961](https://leetcode.cn/problems/check-if-string-is-a-prefix-of-array)|Easy|||||||||||
|[1962](https://leetcode.cn/problems/remove-stones-to-minimize-the-total)|Medium|||||||||||
|[1963](https://leetcode.cn/problems/minimum-number-of-swaps-to-make-the-string-balanced)|Medium|||||||||||
|[1964](https://leetcode.cn/problems/find-the-longest-valid-obstacle-course-at-each-position)|Hard|||||||||||
|[1967](https://leetcode.cn/problems/number-of-strings-that-appear-as-substrings-in-word)|Easy|||||||||||
|[1969](https://leetcode.cn/problems/minimum-non-zero-product-of-the-array-elements)|Medium|||||||||||
|[1970](https://leetcode.cn/problems/last-day-where-you-can-still-cross)|Hard|||||||||||
|[1991](https://leetcode.cn/problems/find-the-middle-index-in-array)|Easy|||||||||||
|[1992](https://leetcode.cn/problems/find-all-groups-of-farmland)|Medium|||||||||||
|[1993](https://leetcode.cn/problems/operations-on-tree)|Medium|||||||||||
|[1994](https://leetcode.cn/problems/the-number-of-good-subsets)|Hard|||||||||||
|[1979](https://leetcode.cn/problems/find-greatest-common-divisor-of-array)|Easy|||||||||||
|[1980](https://leetcode.cn/problems/find-unique-binary-string)|Medium|||||||||||
|[1981](https://leetcode.cn/problems/minimize-the-difference-between-target-and-chosen-elements)|Medium|||||||||||
|[1982](https://leetcode.cn/problems/find-array-given-subset-sums)|Hard|||||||||||
|[1966](https://leetcode.cn/problems/binary-searchable-numbers-in-an-unsorted-array)|Medium|||||||||||
|[1984](https://leetcode.cn/problems/minimum-difference-between-highest-and-lowest-of-k-scores)|Easy|||||||||||
|[1985](https://leetcode.cn/problems/find-the-kth-largest-integer-in-the-array)|Medium|||||||||||
|[1986](https://leetcode.cn/problems/minimum-number-of-work-sessions-to-finish-the-tasks)|Medium|||||||||||
|[1987](https://leetcode.cn/problems/number-of-unique-good-subsequences)|Hard|||||||||||
|[2006](https://leetcode.cn/problems/count-number-of-pairs-with-absolute-difference-k)|Easy|||||||||||
|[2007](https://leetcode.cn/problems/find-original-array-from-doubled-array)|Medium|||||||||||
|[2008](https://leetcode.cn/problems/maximum-earnings-from-taxi)|Medium|||||||||||
|[2009](https://leetcode.cn/problems/minimum-number-of-operations-to-make-array-continuous)|Hard|||||||||||
|[1971](https://leetcode.cn/problems/find-if-path-exists-in-graph)|Easy|||||||||||
|[1995](https://leetcode.cn/problems/count-special-quadruplets)|Easy|||||||||||
|[1996](https://leetcode.cn/problems/the-number-of-weak-characters-in-the-game)|Medium|||||||||||
|[1997](https://leetcode.cn/problems/first-day-where-you-have-been-in-all-the-rooms)|Medium|||||||||||
|[1998](https://leetcode.cn/problems/gcd-sort-of-an-array)|Hard|||||||||||
|[1973](https://leetcode.cn/problems/count-nodes-equal-to-sum-of-descendants)|Medium|||||||||||
|[2000](https://leetcode.cn/problems/reverse-prefix-of-word)|Easy|||||||||||
|[2001](https://leetcode.cn/problems/number-of-pairs-of-interchangeable-rectangles)|Medium|||||||||||
|[2002](https://leetcode.cn/problems/maximum-product-of-the-length-of-two-palindromic-subsequences)|Medium|||||||||||
|[2003](https://leetcode.cn/problems/smallest-missing-genetic-value-in-each-subtree)|Hard|||||||||||
|[2022](https://leetcode.cn/problems/convert-1d-array-into-2d-array)|Easy|||||||||||
|[2023](https://leetcode.cn/problems/number-of-pairs-of-strings-with-concatenation-equal-to-target)|Medium|||||||||||
|[2024](https://leetcode.cn/problems/maximize-the-confusion-of-an-exam)|Medium|||||||||||
|[2025](https://leetcode.cn/problems/maximum-number-of-ways-to-partition-an-array)|Hard|||||||||||
|[2011](https://leetcode.cn/problems/final-value-of-variable-after-performing-operations)|Easy|||||||||||
|[2012](https://leetcode.cn/problems/sum-of-beauty-in-the-array)|Medium|||||||||||
|[2013](https://leetcode.cn/problems/detect-squares)|Medium|||||||||||
|[2014](https://leetcode.cn/problems/longest-subsequence-repeated-k-times)|Hard|||||||||||
|[1999](https://leetcode.cn/problems/smallest-greater-multiple-made-of-two-digits)|Medium|||||||||||
|[2015](https://leetcode.cn/problems/average-height-of-buildings-in-each-segment)|Medium|||||||||||
|[2016](https://leetcode.cn/problems/maximum-difference-between-increasing-elements)|Easy|||||||||||
|[2017](https://leetcode.cn/problems/grid-game)|Medium|||||||||||
|[2018](https://leetcode.cn/problems/check-if-word-can-be-placed-in-crossword)|Medium|||||||||||
|[2019](https://leetcode.cn/problems/the-score-of-students-solving-math-expression)|Hard|||||||||||
|[2037](https://leetcode.cn/problems/minimum-number-of-moves-to-seat-everyone)|Easy|||||||||||
|[2038](https://leetcode.cn/problems/remove-colored-pieces-if-both-neighbors-are-the-same-color)|Medium|||||||||||
|[2040](https://leetcode.cn/problems/kth-smallest-product-of-two-sorted-arrays)|Hard|||||||||||
|[2039](https://leetcode.cn/problems/the-time-when-the-network-becomes-idle)|Medium|||||||||||
|[2005](https://leetcode.cn/problems/subtree-removal-game-with-fibonacci-tree)|Hard|||||||||||
|[2027](https://leetcode.cn/problems/minimum-moves-to-convert-string)|Easy|||||||||||
|[2028](https://leetcode.cn/problems/find-missing-observations)|Medium|||||||||||
|[2029](https://leetcode.cn/problems/stone-game-ix)|Medium|||||||||||
|[2030](https://leetcode.cn/problems/smallest-k-length-subsequence-with-occurrences-of-a-letter)|Hard|||||||||||
|[2032](https://leetcode.cn/problems/two-out-of-three)|Easy|||||||||||
|[2033](https://leetcode.cn/problems/minimum-operations-to-make-a-uni-value-grid)|Medium|||||||||||
|[2034](https://leetcode.cn/problems/stock-price-fluctuation)|Medium|||||||||||
|[2035](https://leetcode.cn/problems/partition-array-into-two-arrays-to-minimize-sum-difference)|Hard|||||||||||
|[2053](https://leetcode.cn/problems/kth-distinct-string-in-an-array)|Easy|||||||||||
|[2054](https://leetcode.cn/problems/two-best-non-overlapping-events)|Medium|||||||||||
|[2055](https://leetcode.cn/problems/plates-between-candles)|Medium|||||||||||
|[2056](https://leetcode.cn/problems/number-of-valid-move-combinations-on-chessboard)|Hard|||||||||||
|[2042](https://leetcode.cn/problems/check-if-numbers-are-ascending-in-a-sentence)|Easy|||||||||||
|[2043](https://leetcode.cn/problems/simple-bank-system)|Medium|||||||||||
|[2044](https://leetcode.cn/problems/count-number-of-maximum-bitwise-or-subsets)|Medium|||||||||||
|[2045](https://leetcode.cn/problems/second-minimum-time-to-reach-destination)|Hard|||||||||||
|[2047](https://leetcode.cn/problems/number-of-valid-words-in-a-sentence)|Easy|||||||||||
|[2048](https://leetcode.cn/problems/next-greater-numerically-balanced-number)|Medium|||||||||||
|[2049](https://leetcode.cn/problems/count-nodes-with-the-highest-score)|Medium|||||||||||
|[2050](https://leetcode.cn/problems/parallel-courses-iii)|Hard|||||||||||
|[2068](https://leetcode.cn/problems/check-whether-two-strings-are-almost-equivalent)|Easy|||||||||||
|[2069](https://leetcode.cn/problems/walking-robot-simulation-ii)|Medium|||||||||||
|[2070](https://leetcode.cn/problems/most-beautiful-item-for-each-query)|Medium|||||||||||
|[2071](https://leetcode.cn/problems/maximum-number-of-tasks-you-can-assign)|Hard|||||||||||
|[2057](https://leetcode.cn/problems/smallest-index-with-equal-value)|Easy|||||||||||
|[2058](https://leetcode.cn/problems/find-the-minimum-and-maximum-number-of-nodes-between-critical-points)|Medium|||||||||||
|[2059](https://leetcode.cn/problems/minimum-operations-to-convert-number)|Medium|||||||||||
|[2060](https://leetcode.cn/problems/check-if-an-original-string-exists-given-two-encoded-strings)|Hard|||||||||||
|[2062](https://leetcode.cn/problems/count-vowel-substrings-of-a-string)|Easy|||||||||||
|[2063](https://leetcode.cn/problems/vowels-of-all-substrings)|Medium|||||||||||
|[2064](https://leetcode.cn/problems/minimized-maximum-of-products-distributed-to-any-store)|Medium|||||||||||
|[2065](https://leetcode.cn/problems/maximum-path-quality-of-a-graph)|Hard|||||||||||
|[2085](https://leetcode.cn/problems/count-common-words-with-one-occurrence)|Easy|||||||||||
|[2086](https://leetcode.cn/problems/minimum-number-of-food-buckets-to-feed-the-hamsters)|Medium|||||||||||
|[2087](https://leetcode.cn/problems/minimum-cost-homecoming-of-a-robot-in-a-grid)|Medium|||||||||||
|[2088](https://leetcode.cn/problems/count-fertile-pyramids-in-a-land)|Hard|||||||||||
|[2073](https://leetcode.cn/problems/time-needed-to-buy-tickets)|Easy|||||||||||
|[2074](https://leetcode.cn/problems/reverse-nodes-in-even-length-groups)|Medium|||||||||||
|[2075](https://leetcode.cn/problems/decode-the-slanted-ciphertext)|Medium|||||||||||
|[2076](https://leetcode.cn/problems/process-restricted-friend-requests)|Hard|||||||||||
|[2078](https://leetcode.cn/problems/two-furthest-houses-with-different-colors)|Easy|||||||||||
|[2132](https://leetcode.cn/problems/stamping-the-grid)|Hard|||||||||||
|[2097](https://leetcode.cn/problems/valid-arrangement-of-pairs)|Hard|||||||||||
|[2081](https://leetcode.cn/problems/sum-of-k-mirror-numbers)|Hard|||||||||||
|[2061](https://leetcode.cn/problems/number-of-spaces-cleaning-robot-cleaned)|Medium|||||||||||
|[2099](https://leetcode.cn/problems/find-subsequence-of-length-k-with-the-largest-sum)|Easy|||||||||||
|[2100](https://leetcode.cn/problems/find-good-days-to-rob-the-bank)|Medium|||||||||||
|[2101](https://leetcode.cn/problems/detonate-the-maximum-bombs)|Medium|||||||||||
|[2102](https://leetcode.cn/problems/sequentially-ordinal-rank-tracker)|Hard|||||||||||
|[2067](https://leetcode.cn/problems/number-of-equal-count-substrings)|Medium|||||||||||
|[2089](https://leetcode.cn/problems/find-target-indices-after-sorting-array)|Easy|||||||||||
|[2090](https://leetcode.cn/problems/k-radius-subarray-averages)|Medium|||||||||||
|[2091](https://leetcode.cn/problems/removing-minimum-and-maximum-from-array)|Medium|||||||||||
|[2092](https://leetcode.cn/problems/find-all-people-with-secret)|Hard|||||||||||
|[2094](https://leetcode.cn/problems/finding-3-digit-even-numbers)|Easy|||||||||||
|[2095](https://leetcode.cn/problems/delete-the-middle-node-of-a-linked-list)|Medium|||||||||||
|[2096](https://leetcode.cn/problems/step-by-step-directions-from-a-binary-tree-node-to-another)|Medium|||||||||||
|[2077](https://leetcode.cn/problems/paths-in-maze-that-lead-to-same-room)|Medium|||||||||||
|[2114](https://leetcode.cn/problems/maximum-number-of-words-found-in-sentences)|Easy|||||||||||
|[2115](https://leetcode.cn/problems/find-all-possible-recipes-from-given-supplies)|Medium|||||||||||
|[2116](https://leetcode.cn/problems/check-if-a-parentheses-string-can-be-valid)|Medium|||||||||||
|[2117](https://leetcode.cn/problems/abbreviating-the-product-of-a-range)|Hard|||||||||||
|[2083](https://leetcode.cn/problems/substrings-that-begin-and-end-with-the-same-letter)|Medium|||||||||||
|[2103](https://leetcode.cn/problems/rings-and-rods)|Easy|||||||||||
|[2104](https://leetcode.cn/problems/sum-of-subarray-ranges)|Medium|||||||||||
|[2105](https://leetcode.cn/problems/watering-plants-ii)|Medium|||||||||||
|[2106](https://leetcode.cn/problems/maximum-fruits-harvested-after-at-most-k-steps)|Hard|||||||||||
|[2093](https://leetcode.cn/problems/minimum-cost-to-reach-city-with-discounts)|Medium|||||||||||
|[2108](https://leetcode.cn/problems/find-first-palindromic-string-in-the-array)|Easy|||||||||||
|[2109](https://leetcode.cn/problems/adding-spaces-to-a-string)|Medium|||||||||||
|[2110](https://leetcode.cn/problems/number-of-smooth-descent-periods-of-a-stock)|Medium|||||||||||
|[2111](https://leetcode.cn/problems/minimum-operations-to-make-the-array-k-increasing)|Hard|||||||||||
|[2129](https://leetcode.cn/problems/capitalize-the-title)|Easy|||||||||||
|[2130](https://leetcode.cn/problems/maximum-twin-sum-of-a-linked-list)|Medium|||||||||||
|[2131](https://leetcode.cn/problems/longest-palindrome-by-concatenating-two-letter-words)|Medium|||||||||||
|[2119](https://leetcode.cn/problems/a-number-after-a-double-reversal)|Easy|||||||||||
|[2120](https://leetcode.cn/problems/execution-of-all-suffix-instructions-staying-in-a-grid)|Medium|||||||||||
|[2121](https://leetcode.cn/problems/intervals-between-identical-elements)|Medium|||||||||||
|[2122](https://leetcode.cn/problems/recover-the-original-array)|Hard|||||||||||
|[2098](https://leetcode.cn/problems/subsequence-of-size-k-with-the-largest-even-sum)|Medium|||||||||||
|[2124](https://leetcode.cn/problems/check-if-all-as-appears-before-all-bs)|Easy|||||||||||
|[2125](https://leetcode.cn/problems/number-of-laser-beams-in-a-bank)|Medium|||||||||||
|[2126](https://leetcode.cn/problems/destroying-asteroids)|Medium|||||||||||
|[2127](https://leetcode.cn/problems/maximum-employees-to-be-invited-to-a-meeting)|Hard|||||||||||
|[2107](https://leetcode.cn/problems/number-of-unique-flavors-after-sharing-k-candies)|Medium|||||||||||
|[2144](https://leetcode.cn/problems/minimum-cost-of-buying-candies-with-discount)|Easy|||||||||||
|[2145](https://leetcode.cn/problems/count-the-hidden-sequences)|Medium|||||||||||
|[2146](https://leetcode.cn/problems/k-highest-ranked-items-within-a-price-range)|Medium|||||||||||
|[2147](https://leetcode.cn/problems/number-of-ways-to-divide-a-long-corridor)|Hard|||||||||||
|[2133](https://leetcode.cn/problems/check-if-every-row-and-column-contains-all-numbers)|Easy|||||||||||
|[2134](https://leetcode.cn/problems/minimum-swaps-to-group-all-1s-together-ii)|Medium|||||||||||
|[2135](https://leetcode.cn/problems/count-words-obtained-after-adding-a-letter)|Medium|||||||||||
|[2136](https://leetcode.cn/problems/earliest-possible-day-of-full-bloom)|Hard|||||||||||
|[2113](https://leetcode.cn/problems/elements-in-array-after-removing-and-replacing-elements)|Medium|||||||||||
|[2123](https://leetcode.cn/problems/minimum-operations-to-remove-adjacent-ones-in-matrix)|Hard|||||||||||
|[2138](https://leetcode.cn/problems/divide-a-string-into-groups-of-size-k)|Easy|||||||||||
|[2155](https://leetcode.cn/problems/all-divisions-with-the-highest-score-of-a-binary-array)|Medium|||||||||||
|[2140](https://leetcode.cn/problems/solving-questions-with-brainpower)|Medium|||||||||||
|[2141](https://leetcode.cn/problems/maximum-running-time-of-n-computers)|Hard|||||||||||
|[2160](https://leetcode.cn/problems/minimum-sum-of-four-digit-number-after-splitting-digits)|Easy|||||||||||
|[2161](https://leetcode.cn/problems/partition-array-according-to-given-pivot)|Medium|||||||||||
|[2162](https://leetcode.cn/problems/minimum-cost-to-set-cooking-time)|Medium|||||||||||
|[2163](https://leetcode.cn/problems/minimum-difference-in-sums-after-removal-of-elements)|Hard|||||||||||
|[2128](https://leetcode.cn/problems/remove-all-ones-with-row-and-column-flips)|Medium|||||||||||
|[2148](https://leetcode.cn/problems/count-elements-with-strictly-smaller-and-greater-elements)|Easy|||||||||||
|[2150](https://leetcode.cn/problems/find-all-lonely-numbers-in-the-array)|Medium|||||||||||
|[2149](https://leetcode.cn/problems/rearrange-array-elements-by-sign)|Medium|||||||||||
|[2151](https://leetcode.cn/problems/maximum-good-people-based-on-statements)|Hard|||||||||||
|[2137](https://leetcode.cn/problems/pour-water-between-buckets-to-make-water-levels-equal)|Medium|||||||||||
|[2154](https://leetcode.cn/problems/keep-multiplying-found-values-by-two)|Easy|||||||||||
|[2156](https://leetcode.cn/problems/find-substring-with-given-hash-value)|Hard|||||||||||
|[2157](https://leetcode.cn/problems/groups-of-strings)|Hard|||||||||||
|[2176](https://leetcode.cn/problems/count-equal-and-divisible-pairs-in-an-array)|Easy|||||||||||
|[2177](https://leetcode.cn/problems/find-three-consecutive-integers-that-sum-to-a-given-number)|Medium|||||||||||
|[2178](https://leetcode.cn/problems/maximum-split-of-positive-even-integers)|Medium|||||||||||
|[2179](https://leetcode.cn/problems/count-good-triplets-in-an-array)|Hard|||||||||||
|[2143](https://leetcode.cn/problems/choose-numbers-from-two-arrays-in-range)|Hard|||||||||||
|[2164](https://leetcode.cn/problems/sort-even-and-odd-indices-independently)|Easy|||||||||||
|[2165](https://leetcode.cn/problems/smallest-value-of-the-rearranged-number)|Medium|||||||||||
|[2166](https://leetcode.cn/problems/design-bitset)|Medium|||||||||||
|[2167](https://leetcode.cn/problems/minimum-time-to-remove-all-cars-containing-illegal-goods)|Hard|||||||||||
|[2152](https://leetcode.cn/problems/minimum-number-of-lines-to-cover-points)|Medium|||||||||||
|[2169](https://leetcode.cn/problems/count-operations-to-obtain-zero)|Easy|||||||||||
|[2170](https://leetcode.cn/problems/minimum-operations-to-make-the-array-alternating)|Medium|||||||||||
|[2171](https://leetcode.cn/problems/removing-minimum-number-of-magic-beans)|Medium|||||||||||
|[2172](https://leetcode.cn/problems/maximum-and-sum-of-array)|Hard|||||||||||
|[2185](https://leetcode.cn/problems/counting-words-with-a-given-prefix)|Easy|||||||||||
|[2186](https://leetcode.cn/problems/minimum-number-of-steps-to-make-two-strings-anagram-ii)|Medium|||||||||||
|[2187](https://leetcode.cn/problems/minimum-time-to-complete-trips)|Medium|||||||||||
|[2188](https://leetcode.cn/problems/minimum-time-to-finish-the-race)|Hard|||||||||||
|[2158](https://leetcode.cn/problems/amount-of-new-area-painted-each-day)|Hard|||||||||||
|[2180](https://leetcode.cn/problems/count-integers-with-even-digit-sum)|Easy|||||||||||
|[2181](https://leetcode.cn/problems/merge-nodes-in-between-zeros)|Medium|||||||||||
|[2182](https://leetcode.cn/problems/construct-string-with-repeat-limit)|Medium|||||||||||
|[2183](https://leetcode.cn/problems/count-array-pairs-divisible-by-k)|Hard|||||||||||
|[2168](https://leetcode.cn/problems/unique-substrings-with-equal-digit-frequency)|Medium|||||||||||
|[2194](https://leetcode.cn/problems/cells-in-a-range-on-an-excel-sheet)|Easy|||||||||||
|[2195](https://leetcode.cn/problems/append-k-integers-with-minimal-sum)|Medium|||||||||||
|[2196](https://leetcode.cn/problems/create-binary-tree-from-descriptions)|Medium|||||||||||
|[2197](https://leetcode.cn/problems/replace-non-coprime-numbers-in-array)|Hard|||||||||||
|[2206](https://leetcode.cn/problems/divide-array-into-equal-pairs)|Easy|||||||||||
|[2207](https://leetcode.cn/problems/maximize-number-of-subsequences-in-a-string)|Medium|||||||||||
|[2208](https://leetcode.cn/problems/minimum-operations-to-halve-array-sum)|Medium|||||||||||
|[2209](https://leetcode.cn/problems/minimum-white-tiles-after-covering-with-carpets)|Hard|||||||||||
|[2190](https://leetcode.cn/problems/most-frequent-number-following-key-in-an-array)|Easy|||||||||||
|[2174](https://leetcode.cn/problems/remove-all-ones-with-row-and-column-flips-ii)|Medium|||||||||||
|[2210](https://leetcode.cn/problems/count-hills-and-valleys-in-an-array)|Easy|||||||||||
|[2211](https://leetcode.cn/problems/count-collisions-on-a-road)|Medium|||||||||||
|[2212](https://leetcode.cn/problems/maximum-points-in-an-archery-competition)|Medium|||||||||||
|[2213](https://leetcode.cn/problems/longest-substring-of-one-repeating-character)|Hard|||||||||||
|[2200](https://leetcode.cn/problems/find-all-k-distant-indices-in-an-array)|Easy|||||||||||
|[2203](https://leetcode.cn/problems/minimum-weighted-subgraph-with-the-required-paths)|Hard|||||||||||
|[2184](https://leetcode.cn/problems/number-of-ways-to-build-sturdy-brick-wall)|Medium|||||||||||
|[2220](https://leetcode.cn/problems/minimum-bit-flips-to-convert-number)|Easy|||||||||||
|[2221](https://leetcode.cn/problems/find-triangular-sum-of-an-array)|Medium|||||||||||
|[2222](https://leetcode.cn/problems/number-of-ways-to-select-buildings)|Medium|||||||||||
|[2223](https://leetcode.cn/problems/sum-of-scores-of-built-strings)|Hard|||||||||||
|[2231](https://leetcode.cn/problems/largest-number-after-digit-swaps-by-parity)|Easy|||||||||||
|[2232](https://leetcode.cn/problems/minimize-result-by-adding-parentheses-to-expression)|Medium|||||||||||
|[2233](https://leetcode.cn/problems/maximum-product-after-k-increments)|Medium|||||||||||
|[2234](https://leetcode.cn/problems/maximum-total-beauty-of-the-gardens)|Hard|||||||||||
|[2248](https://leetcode.cn/problems/intersection-of-multiple-arrays)|Easy|||||||||||
|[2249](https://leetcode.cn/problems/count-lattice-points-inside-a-circle)|Medium|||||||||||
|[2250](https://leetcode.cn/problems/count-number-of-rectangles-containing-each-point)|Medium|||||||||||
|[2251](https://leetcode.cn/problems/number-of-flowers-in-full-bloom)|Hard|||||||||||
|[2259](https://leetcode.cn/problems/remove-digit-from-number-to-maximize-result)|Easy|||||||||||
|[2260](https://leetcode.cn/problems/minimum-consecutive-cards-to-pick-up)|Medium|||||||||||
|[2261](https://leetcode.cn/problems/k-divisible-elements-subarrays)|Medium|||||||||||
|[2262](https://leetcode.cn/problems/total-appeal-of-a-string)|Hard|||||||||||
|[2255](https://leetcode.cn/problems/count-prefixes-of-a-given-string)|Easy|||||||||||
|[2256](https://leetcode.cn/problems/minimum-average-difference)|Medium|||||||||||
|[2257](https://leetcode.cn/problems/count-unguarded-cells-in-the-grid)|Medium|||||||||||
|[2258](https://leetcode.cn/problems/escape-the-spreading-fire)|Hard|||||||||||
|[2224](https://leetcode.cn/problems/minimum-number-of-operations-to-convert-time)|Easy|||||||||||
|[2264](https://leetcode.cn/problems/largest-3-same-digit-number-in-string)|Easy|||||||||||
|[2265](https://leetcode.cn/problems/count-nodes-equal-to-average-of-subtree)|Medium|||||||||||
|[2266](https://leetcode.cn/problems/count-number-of-texts)|Medium|||||||||||
|[2267](https://leetcode.cn/problems/check-if-there-is-a-valid-parentheses-string-path)|Hard|||||||||||
|[2239](https://leetcode.cn/problems/find-closest-number-to-zero)|Easy|||||||||||
|[2240](https://leetcode.cn/problems/number-of-ways-to-buy-pens-and-pencils)|Medium|||||||||||
|[2241](https://leetcode.cn/problems/design-an-atm-machine)|Medium|||||||||||
|[2242](https://leetcode.cn/problems/maximum-score-of-a-node-sequence)|Hard|||||||||||
|[2214](https://leetcode.cn/problems/minimum-health-to-beat-game)|Medium|||||||||||
|[2274](https://leetcode.cn/problems/maximum-consecutive-floors-without-special-floors)|Medium|||||||||||
|[2275](https://leetcode.cn/problems/largest-combination-with-bitwise-and-greater-than-zero)|Medium|||||||||||
|[2276](https://leetcode.cn/problems/count-integers-in-intervals)|Hard|||||||||||
|[2270](https://leetcode.cn/problems/number-of-ways-to-split-array)|Medium|||||||||||
|[2271](https://leetcode.cn/problems/maximum-white-tiles-covered-by-a-carpet)|Medium|||||||||||
|[2272](https://leetcode.cn/problems/substring-with-largest-variance)|Hard|||||||||||
|[2243](https://leetcode.cn/problems/calculate-digit-sum-of-a-string)|Easy|||||||||||
|[2244](https://leetcode.cn/problems/minimum-rounds-to-complete-all-tasks)|Medium|||||||||||
|[2245](https://leetcode.cn/problems/maximum-trailing-zeros-in-a-cornered-path)|Medium|||||||||||
|[2246](https://leetcode.cn/problems/longest-path-with-different-adjacent-characters)|Hard|||||||||||
|[2278](https://leetcode.cn/problems/percentage-of-letter-in-string)|Easy|||||||||||
|[2279](https://leetcode.cn/problems/maximum-bags-with-full-capacity-of-rocks)|Medium|||||||||||
|[2280](https://leetcode.cn/problems/minimum-lines-to-represent-a-line-chart)|Medium|||||||||||
|[2281](https://leetcode.cn/problems/sum-of-total-strength-of-wizards)|Hard|||||||||||
|[2219](https://leetcode.cn/problems/maximum-sum-score-of-array)|Medium|||||||||||
|[2287](https://leetcode.cn/problems/rearrange-characters-to-make-target-string)|Easy|||||||||||
|[2288](https://leetcode.cn/problems/apply-discount-to-prices)|Medium|||||||||||
|[2289](https://leetcode.cn/problems/steps-to-make-array-non-decreasing)|Medium|||||||||||
|[2290](https://leetcode.cn/problems/minimum-obstacle-removal-to-reach-corner)|Hard|||||||||||
|[2283](https://leetcode.cn/problems/check-if-number-has-equal-digit-count-and-digit-value)|Easy|||||||||||
|[2284](https://leetcode.cn/problems/sender-with-largest-word-count)|Medium|||||||||||
|[2285](https://leetcode.cn/problems/maximum-total-importance-of-roads)|Medium|||||||||||
|[2286](https://leetcode.cn/problems/booking-concert-tickets-in-groups)|Hard|||||||||||
|[2235](https://leetcode.cn/problems/add-two-integers)|Easy|||||||||||
|[2236](https://leetcode.cn/problems/root-equals-sum-of-children)|Easy|||||||||||
|[2237](https://leetcode.cn/problems/count-positions-on-street-with-required-brightness)|Medium|||||||||||
|[2293](https://leetcode.cn/problems/min-max-game)|Easy|||||||||||
|[2294](https://leetcode.cn/problems/partition-array-such-that-maximum-difference-is-k)|Medium|||||||||||
|[2295](https://leetcode.cn/problems/replace-elements-in-an-array)|Medium|||||||||||
|[2296](https://leetcode.cn/problems/design-a-text-editor)|Hard|||||||||||
|[2306](https://leetcode.cn/problems/naming-a-company)|Hard|||||||||||
|[2299](https://leetcode.cn/problems/strong-password-checker-ii)|Easy|||||||||||
|[2300](https://leetcode.cn/problems/successful-pairs-of-spells-and-potions)|Medium|||||||||||
|[2301](https://leetcode.cn/problems/match-substring-after-replacement)|Hard|||||||||||
|[2302](https://leetcode.cn/problems/count-subarrays-with-score-less-than-k)|Hard|||||||||||
|[2311](https://leetcode.cn/problems/longest-binary-subsequence-less-than-or-equal-to-k)|Medium|||||||||||
|[2254](https://leetcode.cn/problems/design-video-sharing-platform)|Hard|||||||||||
|[2320](https://leetcode.cn/problems/count-number-of-ways-to-place-houses)|Medium|||||||||||
|[2319](https://leetcode.cn/problems/check-if-matrix-is-x-matrix)|Easy|||||||||||
|[2313](https://leetcode.cn/problems/minimum-flips-in-binary-tree-to-get-result)|Hard|||||||||||
|[2322](https://leetcode.cn/problems/minimum-score-after-removals-on-a-tree)|Hard|||||||||||
|[2315](https://leetcode.cn/problems/count-asterisks)|Easy|||||||||||
|[2317](https://leetcode.cn/problems/maximum-xor-after-operations)|Medium|||||||||||
|[2316](https://leetcode.cn/problems/count-unreachable-pairs-of-nodes-in-an-undirected-graph)|Medium|||||||||||
|[2318](https://leetcode.cn/problems/number-of-distinct-roll-sequences)|Hard|||||||||||
|[2268](https://leetcode.cn/problems/minimum-number-of-keypresses)|Medium|||||||||||
|[2325](https://leetcode.cn/problems/decode-the-message)|Easy|||||||||||
|[2327](https://leetcode.cn/problems/number-of-people-aware-of-a-secret)|Medium|||||||||||
|[2328](https://leetcode.cn/problems/number-of-increasing-paths-in-a-grid)|Hard|||||||||||
|[2326](https://leetcode.cn/problems/spiral-matrix-iv)|Medium|||||||||||
|[2335](https://leetcode.cn/problems/minimum-amount-of-time-to-fill-cups)|Easy|||||||||||
|[2336](https://leetcode.cn/problems/smallest-number-in-infinite-set)|Medium|||||||||||
|[2337](https://leetcode.cn/problems/move-pieces-to-obtain-a-string)|Medium|||||||||||
|[2338](https://leetcode.cn/problems/count-the-number-of-ideal-arrays)|Hard|||||||||||
|[2331](https://leetcode.cn/problems/evaluate-boolean-binary-tree)|Easy|||||||||||
|[2332](https://leetcode.cn/problems/the-latest-time-to-catch-a-bus)|Medium|||||||||||
|[2333](https://leetcode.cn/problems/minimum-sum-of-squared-difference)|Medium|||||||||||
|[2334](https://leetcode.cn/problems/subarray-with-elements-greater-than-varying-threshold)|Hard|||||||||||
|[2277](https://leetcode.cn/problems/closest-node-to-path-in-tree)|Hard|||||||||||
|[2341](https://leetcode.cn/problems/maximum-number-of-pairs-in-array)|Easy|||||||||||
|[2343](https://leetcode.cn/problems/query-kth-smallest-trimmed-number)|Medium|||||||||||
|[2344](https://leetcode.cn/problems/minimum-deletions-to-make-array-divisible)|Hard|||||||||||
|[2282](https://leetcode.cn/problems/number-of-people-that-can-be-seen-in-a-grid)|Medium|||||||||||
|[2291](https://leetcode.cn/problems/maximum-profit-from-trading-stocks)|Medium|||||||||||
|[2351](https://leetcode.cn/problems/first-letter-to-appear-twice)|Easy|||||||||||
|[2352](https://leetcode.cn/problems/equal-row-and-column-pairs)|Medium|||||||||||
|[2353](https://leetcode.cn/problems/design-a-food-rating-system)|Medium|||||||||||
|[2354](https://leetcode.cn/problems/number-of-excellent-pairs)|Hard|||||||||||
|[2348](https://leetcode.cn/problems/number-of-zero-filled-subarrays)|Medium|||||||||||
|[2347](https://leetcode.cn/problems/best-poker-hand)|Easy|||||||||||
|[2349](https://leetcode.cn/problems/design-a-number-container-system)|Medium|||||||||||
|[2350](https://leetcode.cn/problems/shortest-impossible-sequence-of-rolls)|Hard|||||||||||
|[2357](https://leetcode.cn/problems/make-array-zero-by-subtracting-equal-amounts)|Easy|||||||||||
|[2358](https://leetcode.cn/problems/maximum-number-of-groups-entering-a-competition)|Medium|||||||||||
|[2359](https://leetcode.cn/problems/find-closest-node-to-given-two-nodes)|Medium|||||||||||
|[2360](https://leetcode.cn/problems/longest-cycle-in-a-graph)|Hard|||||||||||
|[2367](https://leetcode.cn/problems/number-of-arithmetic-triplets)|Easy|||||||||||
|[2369](https://leetcode.cn/problems/check-if-there-is-a-valid-partition-for-the-array)|Medium|||||||||||
|[2370](https://leetcode.cn/problems/longest-ideal-subsequence)|Medium|||||||||||
|[2368](https://leetcode.cn/problems/reachable-nodes-with-restrictions)|Medium|||||||||||
|[2363](https://leetcode.cn/problems/merge-similar-items)|Easy|||||||||||
|[2364](https://leetcode.cn/problems/count-number-of-bad-pairs)|Medium|||||||||||
|[2398](https://leetcode.cn/problems/maximum-number-of-robots-within-budget)|Hard|||||||||||
|[2366](https://leetcode.cn/problems/minimum-replacements-to-sort-the-array)|Hard|||||||||||
|[2373](https://leetcode.cn/problems/largest-local-values-in-a-matrix)|Easy|||||||||||
|[2374](https://leetcode.cn/problems/node-with-highest-edge-score)|Medium|||||||||||
|[2375](https://leetcode.cn/problems/construct-smallest-number-from-di-string)|Medium|||||||||||
|[2376](https://leetcode.cn/problems/count-special-integers)|Hard|||||||||||
|[2323](https://leetcode.cn/problems/find-minimum-time-to-finish-all-jobs-ii)|Medium|||||||||||
|[2383](https://leetcode.cn/problems/minimum-hours-of-training-to-win-a-competition)|Easy|||||||||||
|[2385](https://leetcode.cn/problems/amount-of-time-for-binary-tree-to-be-infected)|Medium|||||||||||
|[2386](https://leetcode.cn/problems/find-the-k-sum-of-an-array)|Hard|||||||||||
|[2379](https://leetcode.cn/problems/minimum-recolors-to-get-k-consecutive-black-blocks)|Easy|||||||||||
|[2380](https://leetcode.cn/problems/time-needed-to-rearrange-a-binary-string)|Medium|||||||||||
|[2381](https://leetcode.cn/problems/shifting-letters-ii)|Medium|||||||||||
|[2382](https://leetcode.cn/problems/maximum-segment-sum-after-removals)|Hard|||||||||||
|[2330](https://leetcode.cn/problems/valid-palindrome-iv)|Medium|||||||||||
|[2389](https://leetcode.cn/problems/longest-subsequence-with-limited-sum)|Easy|||||||||||
|[2390](https://leetcode.cn/problems/removing-stars-from-a-string)|Medium|||||||||||
|[2391](https://leetcode.cn/problems/minimum-amount-of-time-to-collect-garbage)|Medium|||||||||||
|[2392](https://leetcode.cn/problems/build-a-matrix-with-conditions)|Hard|||||||||||
|[2342](https://leetcode.cn/problems/max-sum-of-a-pair-with-equal-sum-of-digits)|Medium|||||||||||
|[2340](https://leetcode.cn/problems/minimum-adjacent-swaps-to-make-a-valid-array)|Medium|||||||||||
|[2384](https://leetcode.cn/problems/largest-palindromic-number)|Medium|||||||||||
|[2399](https://leetcode.cn/problems/check-distances-between-same-letters)|Easy|||||||||||
|[2400](https://leetcode.cn/problems/number-of-ways-to-reach-a-position-after-exactly-k-steps)|Medium|||||||||||
|[2401](https://leetcode.cn/problems/longest-nice-subarray)|Medium|||||||||||
|[2402](https://leetcode.cn/problems/meeting-rooms-iii)|Hard|||||||||||
|[2395](https://leetcode.cn/problems/find-subarrays-with-equal-sum)|Easy|||||||||||
|[2396](https://leetcode.cn/problems/strictly-palindromic-number)|Medium|||||||||||
|[2397](https://leetcode.cn/problems/maximum-rows-covered-by-columns)|Medium|||||||||||
|[2365](https://leetcode.cn/problems/task-scheduler-ii)|Medium|||||||||||
|[2345](https://leetcode.cn/problems/finding-the-number-of-visible-mountains)|Medium|||||||||||
|[2404](https://leetcode.cn/problems/most-frequent-even-element)|Easy|||||||||||
|[2405](https://leetcode.cn/problems/optimal-partition-of-string)|Medium|||||||||||
|[2406](https://leetcode.cn/problems/divide-intervals-into-minimum-number-of-groups)|Medium|||||||||||
|[2459](https://leetcode.cn/problems/sort-array-by-moving-items-to-empty-space)|Hard|||||||||||
|[2355](https://leetcode.cn/problems/maximum-number-of-books-you-can-take)|Hard|||||||||||
|[2413](https://leetcode.cn/problems/smallest-even-multiple)|Easy|||||||||||
|[2414](https://leetcode.cn/problems/length-of-the-longest-alphabetical-continuous-substring)|Medium|||||||||||
|[2415](https://leetcode.cn/problems/reverse-odd-levels-of-binary-tree)|Medium|||||||||||
|[2416](https://leetcode.cn/problems/sum-of-prefix-scores-of-strings)|Hard|||||||||||
|[2409](https://leetcode.cn/problems/count-days-spent-together)|Easy|||||||||||
|[2410](https://leetcode.cn/problems/maximum-matching-of-players-with-trainers)|Medium|||||||||||
|[2411](https://leetcode.cn/problems/smallest-subarrays-with-maximum-bitwise-or)|Medium|||||||||||
|[2412](https://leetcode.cn/problems/minimum-money-required-before-transactions)|Hard|||||||||||
|[2361](https://leetcode.cn/problems/minimum-costs-using-the-train-line)|Hard|||||||||||
|[2418](https://leetcode.cn/problems/sort-the-people)|Easy|||||||||||
|[2419](https://leetcode.cn/problems/longest-subarray-with-maximum-bitwise-and)|Medium|||||||||||
|[2420](https://leetcode.cn/problems/find-all-good-indices)|Medium|||||||||||
|[2421](https://leetcode.cn/problems/number-of-good-paths)|Hard|||||||||||
|[2371](https://leetcode.cn/problems/minimize-maximum-value-in-a-grid)|Hard|||||||||||
|[2427](https://leetcode.cn/problems/number-of-common-factors)|Easy|||||||||||
|[2428](https://leetcode.cn/problems/maximum-sum-of-an-hourglass)|Medium|||||||||||
|[2429](https://leetcode.cn/problems/minimize-xor)|Medium|||||||||||
|[2430](https://leetcode.cn/problems/maximum-deletions-on-a-string)|Hard|||||||||||
|[2522](https://leetcode.cn/problems/partition-string-into-substrings-with-values-at-most-k)|Medium|||||||||||
|[2424](https://leetcode.cn/problems/longest-uploaded-prefix)|Medium|||||||||||
|[2426](https://leetcode.cn/problems/number-of-pairs-satisfying-inequality)|Hard|||||||||||
|[2378](https://leetcode.cn/problems/choose-edges-to-maximize-score-in-a-tree)|Medium|||||||||||
|[2432](https://leetcode.cn/problems/the-employee-that-worked-on-the-longest-task)|Easy|||||||||||
|[2433](https://leetcode.cn/problems/find-the-original-array-of-prefix-xor)|Medium|||||||||||
|[2434](https://leetcode.cn/problems/using-a-robot-to-print-the-lexicographically-smallest-string)|Medium|||||||||||
|[2435](https://leetcode.cn/problems/paths-in-matrix-whose-sum-is-divisible-by-k)|Hard|||||||||||
|[2387](https://leetcode.cn/problems/median-of-a-row-wise-sorted-matrix)|Medium|||||||||||
|[2441](https://leetcode.cn/problems/largest-positive-integer-that-exists-with-its-negative)|Easy|||||||||||
|[2442](https://leetcode.cn/problems/count-number-of-distinct-integers-after-reverse-operations)|Medium|||||||||||
|[2407](https://leetcode.cn/problems/longest-increasing-subsequence-ii)|Hard|||||||||||
|[2444](https://leetcode.cn/problems/count-subarrays-with-fixed-bounds)|Hard|||||||||||
|[2437](https://leetcode.cn/problems/number-of-valid-clock-times)|Easy|||||||||||
|[2438](https://leetcode.cn/problems/range-product-queries-of-powers)|Medium|||||||||||
|[2439](https://leetcode.cn/problems/minimize-maximum-of-array)|Medium|||||||||||
|[2440](https://leetcode.cn/problems/create-components-with-same-value)|Hard|||||||||||
|[2423](https://leetcode.cn/problems/remove-letter-to-equalize-frequency)|Easy|||||||||||
|[2425](https://leetcode.cn/problems/bitwise-xor-of-all-pairings)|Medium|||||||||||
|[2393](https://leetcode.cn/problems/count-strictly-increasing-subarrays)|Medium|||||||||||
|[2446](https://leetcode.cn/problems/determine-if-two-events-have-conflict)|Easy|||||||||||
|[2403](https://leetcode.cn/problems/minimum-time-to-kill-all-monsters)|Hard|||||||||||
|[2448](https://leetcode.cn/problems/minimum-cost-to-make-array-equal)|Hard|||||||||||
|[2449](https://leetcode.cn/problems/minimum-number-of-operations-to-make-arrays-similar)|Hard|||||||||||
|[2647](https://leetcode.cn/problems/color-the-triangle-red)|Hard|||||||||||
|[2443](https://leetcode.cn/problems/sum-of-number-and-its-reverse)|Medium|||||||||||
|[2455](https://leetcode.cn/problems/average-value-of-even-numbers-that-are-divisible-by-three)|Easy|||||||||||
|[2456](https://leetcode.cn/problems/most-popular-video-creator)|Medium|||||||||||
|[2457](https://leetcode.cn/problems/minimum-addition-to-make-integer-beautiful)|Medium|||||||||||
|[2458](https://leetcode.cn/problems/height-of-binary-tree-after-subtree-removal-queries)|Hard|||||||||||
|[2447](https://leetcode.cn/problems/number-of-subarrays-with-gcd-equal-to-k)|Medium|||||||||||
|[2451](https://leetcode.cn/problems/odd-string-difference)|Easy|||||||||||
|[2453](https://leetcode.cn/problems/destroy-sequential-targets)|Medium|||||||||||
|[2454](https://leetcode.cn/problems/next-greater-element-iv)|Hard|||||||||||
|[2452](https://leetcode.cn/problems/words-within-two-edits-of-dictionary)|Medium|||||||||||
|[2460](https://leetcode.cn/problems/apply-operations-to-an-array)|Easy|||||||||||
|[2461](https://leetcode.cn/problems/maximum-sum-of-distinct-subarrays-with-length-k)|Medium|||||||||||
|[2462](https://leetcode.cn/problems/total-cost-to-hire-k-workers)|Medium|||||||||||
|[2463](https://leetcode.cn/problems/minimum-total-distance-traveled)|Hard|||||||||||
|[2408](https://leetcode.cn/problems/design-sql)|Medium|||||||||||
|[2469](https://leetcode.cn/problems/convert-the-temperature)|Easy|||||||||||
|[2470](https://leetcode.cn/problems/number-of-subarrays-with-lcm-equal-to-k)|Medium|||||||||||
|[2471](https://leetcode.cn/problems/minimum-number-of-operations-to-sort-a-binary-tree-by-level)|Medium|||||||||||
|[2472](https://leetcode.cn/problems/maximum-number-of-non-overlapping-palindrome-substrings)|Hard|||||||||||
|[2417](https://leetcode.cn/problems/closest-fair-integer)|Medium|||||||||||
|[2465](https://leetcode.cn/problems/number-of-distinct-averages)|Easy|||||||||||
|[2466](https://leetcode.cn/problems/count-ways-to-build-good-strings)|Medium|||||||||||
|[2468](https://leetcode.cn/problems/split-message-based-on-limit)|Hard|||||||||||
|[2467](https://leetcode.cn/problems/most-profitable-path-in-a-tree)|Medium|||||||||||
|[2422](https://leetcode.cn/problems/merge-operations-to-turn-array-into-a-palindrome)|Medium|||||||||||
|[2475](https://leetcode.cn/problems/number-of-unequal-triplets-in-array)|Easy|||||||||||
|[2476](https://leetcode.cn/problems/closest-nodes-queries-in-a-binary-search-tree)|Medium|||||||||||
|[2477](https://leetcode.cn/problems/minimum-fuel-cost-to-report-to-the-capital)|Medium|||||||||||
|[2478](https://leetcode.cn/problems/number-of-beautiful-partitions)|Hard|||||||||||
|[2431](https://leetcode.cn/problems/maximize-total-tastiness-of-purchased-fruits)|Medium|||||||||||
|[2485](https://leetcode.cn/problems/find-the-pivot-integer)|Easy|||||||||||
|[2486](https://leetcode.cn/problems/append-characters-to-string-to-make-subsequence)|Medium|||||||||||
|[2487](https://leetcode.cn/problems/remove-nodes-from-linked-list)|Medium|||||||||||
|[2488](https://leetcode.cn/problems/count-subarrays-with-median-k)|Hard|||||||||||
|[2481](https://leetcode.cn/problems/minimum-cuts-to-divide-a-circle)|Easy|||||||||||
|[2483](https://leetcode.cn/problems/minimum-penalty-for-a-shop)|Medium|||||||||||
|[2484](https://leetcode.cn/problems/count-palindromic-subsequences)|Hard|||||||||||
|[2436](https://leetcode.cn/problems/minimum-split-into-subarrays-with-gcd-greater-than-one)|Medium|||||||||||
|[2490](https://leetcode.cn/problems/circular-sentence)|Easy|||||||||||
|[2491](https://leetcode.cn/problems/divide-players-into-teams-of-equal-skill)|Medium|||||||||||
|[2492](https://leetcode.cn/problems/minimum-score-of-a-path-between-two-cities)|Medium|||||||||||
|[2493](https://leetcode.cn/problems/divide-nodes-into-the-maximum-number-of-groups)|Hard|||||||||||
|[2445](https://leetcode.cn/problems/number-of-nodes-with-value-one)|Medium|||||||||||
|[2500](https://leetcode.cn/problems/delete-greatest-value-in-each-row)|Easy|||||||||||
|[2501](https://leetcode.cn/problems/longest-square-streak-in-an-array)|Medium|||||||||||
|[2502](https://leetcode.cn/problems/design-memory-allocator)|Medium|||||||||||
|[2503](https://leetcode.cn/problems/maximum-number-of-points-from-grid-queries)|Hard|||||||||||
|[2496](https://leetcode.cn/problems/maximum-value-of-a-string-in-an-array)|Easy|||||||||||
|[2497](https://leetcode.cn/problems/maximum-star-sum-of-a-graph)|Medium|||||||||||
|[2498](https://leetcode.cn/problems/frog-jump-ii)|Medium|||||||||||
|[2499](https://leetcode.cn/problems/minimum-total-cost-to-make-arrays-unequal)|Hard|||||||||||
|[2450](https://leetcode.cn/problems/number-of-distinct-binary-strings-after-applying-operations)|Medium|||||||||||
|[2506](https://leetcode.cn/problems/count-pairs-of-similar-strings)|Easy|||||||||||
|[2507](https://leetcode.cn/problems/smallest-value-after-replacing-with-sum-of-prime-factors)|Medium|||||||||||
|[2508](https://leetcode.cn/problems/add-edges-to-make-degrees-of-all-nodes-even)|Hard|||||||||||
|[2509](https://leetcode.cn/problems/cycle-length-queries-in-a-tree)|Hard|||||||||||
|[2515](https://leetcode.cn/problems/shortest-distance-to-target-string-in-a-circular-array)|Easy|||||||||||
|[2516](https://leetcode.cn/problems/take-k-of-each-character-from-left-and-right)|Medium|||||||||||
|[2517](https://leetcode.cn/problems/maximum-tastiness-of-candy-basket)|Medium|||||||||||
|[2518](https://leetcode.cn/problems/number-of-great-partitions)|Hard|||||||||||
|[2511](https://leetcode.cn/problems/maximum-enemy-forts-that-can-be-captured)|Easy|||||||||||
|[2512](https://leetcode.cn/problems/reward-top-k-students)|Medium|||||||||||
|[2541](https://leetcode.cn/problems/minimum-operations-to-make-array-equal-ii)|Medium|||||||||||
|[2514](https://leetcode.cn/problems/count-anagrams)|Hard|||||||||||
|[2482](https://leetcode.cn/problems/difference-between-ones-and-zeros-in-row-and-column)|Medium|||||||||||
|[2464](https://leetcode.cn/problems/minimum-subarrays-in-a-valid-split)|Medium|||||||||||
|[2520](https://leetcode.cn/problems/count-the-digits-that-divide-a-number)|Easy|||||||||||
|[2521](https://leetcode.cn/problems/distinct-prime-factors-of-product-of-array)|Medium|||||||||||
|[2523](https://leetcode.cn/problems/closest-prime-numbers-in-range)|Medium|||||||||||
|[2473](https://leetcode.cn/problems/minimum-cost-to-buy-apples)|Medium|||||||||||
|[2529](https://leetcode.cn/problems/maximum-count-of-positive-integer-and-negative-integer)|Easy|||||||||||
|[2531](https://leetcode.cn/problems/make-number-of-distinct-characters-equal)|Medium|||||||||||
|[2530](https://leetcode.cn/problems/maximal-score-after-applying-k-operations)|Medium|||||||||||
|[2534](https://leetcode.cn/problems/time-taken-to-cross-the-door)|Hard|||||||||||
|[2528](https://leetcode.cn/problems/maximize-the-minimum-powered-city)|Hard|||||||||||
|[2525](https://leetcode.cn/problems/categorize-box-according-to-criteria)|Easy|||||||||||
|[2526](https://leetcode.cn/problems/find-consecutive-integers-from-a-data-stream)|Medium|||||||||||
|[2527](https://leetcode.cn/problems/find-xor-beauty-of-array)|Medium|||||||||||
|[2479](https://leetcode.cn/problems/maximum-xor-of-two-non-overlapping-subtrees)|Hard|||||||||||
|[2535](https://leetcode.cn/problems/difference-between-element-sum-and-digit-sum-of-an-array)|Easy|||||||||||
|[2536](https://leetcode.cn/problems/increment-submatrices-by-one)|Medium|||||||||||
|[2537](https://leetcode.cn/problems/count-the-number-of-good-subarrays)|Medium|||||||||||
|[2538](https://leetcode.cn/problems/difference-between-maximum-and-minimum-price-sum)|Hard|||||||||||
|[2513](https://leetcode.cn/problems/minimize-the-maximum-of-two-arrays)|Medium|||||||||||
|[2489](https://leetcode.cn/problems/number-of-substrings-with-fixed-ratio)|Medium|||||||||||
|[2544](https://leetcode.cn/problems/alternating-digit-sum)|Easy|||||||||||
|[2545](https://leetcode.cn/problems/sort-the-students-by-their-kth-score)|Medium|||||||||||
|[2546](https://leetcode.cn/problems/apply-bitwise-operations-to-make-strings-equal)|Medium|||||||||||
|[2547](https://leetcode.cn/problems/minimum-cost-to-split-an-array)|Hard|||||||||||
|[2540](https://leetcode.cn/problems/minimum-common-value)|Easy|||||||||||
|[2543](https://leetcode.cn/problems/check-if-point-is-reachable)|Hard|||||||||||
|[2542](https://leetcode.cn/problems/maximum-subsequence-score)|Medium|||||||||||
|[2495](https://leetcode.cn/problems/number-of-subarrays-having-even-product)|Medium|||||||||||
|[2553](https://leetcode.cn/problems/separate-the-digits-in-an-array)|Easy|||||||||||
|[2554](https://leetcode.cn/problems/maximum-number-of-integers-to-choose-from-a-range-i)|Medium|||||||||||
|[2556](https://leetcode.cn/problems/disconnect-path-in-a-binary-matrix-by-at-most-one-flip)|Medium|||||||||||
|[2532](https://leetcode.cn/problems/time-to-cross-a-bridge)|Hard|||||||||||
|[2505](https://leetcode.cn/problems/bitwise-or-of-all-subsequence-sums)|Medium|||||||||||
|[2582](https://leetcode.cn/problems/pass-the-pillow)|Easy|||||||||||
|[2583](https://leetcode.cn/problems/kth-largest-sum-in-a-binary-tree)|Medium|||||||||||
|[2584](https://leetcode.cn/problems/split-the-array-to-make-coprime-products)|Hard|||||||||||
|[2585](https://leetcode.cn/problems/number-of-ways-to-earn-points)|Hard|||||||||||
|[2579](https://leetcode.cn/problems/count-total-number-of-colored-cells)|Medium|||||||||||
|[2578](https://leetcode.cn/problems/split-with-minimum-sum)|Easy|||||||||||
|[2580](https://leetcode.cn/problems/count-ways-to-group-overlapping-ranges)|Medium|||||||||||
|[2581](https://leetcode.cn/problems/count-number-of-possible-root-nodes)|Hard|||||||||||
|[2510](https://leetcode.cn/problems/check-if-there-is-a-path-with-equal-number-of-0s-and-1s)|Medium|||||||||||
|[2586](https://leetcode.cn/problems/count-the-number-of-vowel-strings-in-range)|Easy|||||||||||
|[2587](https://leetcode.cn/problems/rearrange-array-to-maximize-prefix-score)|Medium|||||||||||
|[2588](https://leetcode.cn/problems/count-the-number-of-beautiful-subarrays)|Medium|||||||||||
|[2589](https://leetcode.cn/problems/minimum-time-to-complete-all-tasks)|Hard|||||||||||
|[2519](https://leetcode.cn/problems/count-the-number-of-k-big-indices)|Hard|||||||||||
|[2595](https://leetcode.cn/problems/number-of-even-and-odd-bits)|Easy|||||||||||
|[2598](https://leetcode.cn/problems/smallest-missing-non-negative-integer-after-operations)|Medium|||||||||||
|[2596](https://leetcode.cn/problems/check-knight-tour-configuration)|Medium|||||||||||
|[2591](https://leetcode.cn/problems/distribute-money-to-maximum-children)|Easy|||||||||||
|[2592](https://leetcode.cn/problems/maximize-greatness-of-an-array)|Medium|||||||||||
|[2594](https://leetcode.cn/problems/minimum-time-to-repair-cars)|Medium|||||||||||
|[2604](https://leetcode.cn/problems/minimum-time-to-eat-all-grains)|Hard|||||||||||
|[2524](https://leetcode.cn/problems/maximum-frequency-score-of-a-subarray)|Hard|||||||||||
|[2605](https://leetcode.cn/problems/form-smallest-number-from-two-digit-arrays)|Easy|||||||||||
|[2606](https://leetcode.cn/problems/find-the-substring-with-maximum-cost)|Medium|||||||||||
|[2607](https://leetcode.cn/problems/make-k-subarray-sums-equal)|Medium|||||||||||
|[2608](https://leetcode.cn/problems/shortest-cycle-in-a-graph)|Hard|||||||||||
|[2533](https://leetcode.cn/problems/number-of-good-binary-strings)|Medium|||||||||||
|[2555](https://leetcode.cn/problems/maximize-win-from-two-segments)|Medium|||||||||||
|[2599](https://leetcode.cn/problems/make-the-prefix-sum-non-negative)|Medium|||||||||||
|[2639](https://leetcode.cn/problems/find-the-width-of-columns-of-a-grid)|Easy|||||||||||
|[2640](https://leetcode.cn/problems/find-the-score-of-all-prefixes-of-an-array)|Medium|||||||||||
|[2641](https://leetcode.cn/problems/cousins-in-binary-tree-ii)|Medium|||||||||||
|[2642](https://leetcode.cn/problems/design-graph-with-shortest-path-calculator)|Hard|||||||||||
|[2549](https://leetcode.cn/problems/count-distinct-numbers-on-board)|Easy|||||||||||
|[2550](https://leetcode.cn/problems/count-collisions-of-monkeys-on-a-polygon)|Medium|||||||||||
|[2551](https://leetcode.cn/problems/put-marbles-in-bags)|Hard|||||||||||
|[2552](https://leetcode.cn/problems/count-increasing-quadruplets)|Hard|||||||||||
|[2539](https://leetcode.cn/problems/count-the-number-of-good-subsequences)|Medium|||||||||||
|[2660](https://leetcode.cn/problems/determine-the-winner-of-a-bowling-game)|Easy|||||||||||
|[2661](https://leetcode.cn/problems/first-completely-painted-row-or-column)|Medium|||||||||||
|[2662](https://leetcode.cn/problems/minimum-cost-of-a-path-with-special-roads)|Medium|||||||||||
|[2663](https://leetcode.cn/problems/lexicographically-smallest-beautiful-string)|Hard|||||||||||
|[2590](https://leetcode.cn/problems/design-a-todo-list)|Medium|||||||||||
|[2561](https://leetcode.cn/problems/rearranging-fruits)|Hard|||||||||||
|[2560](https://leetcode.cn/problems/house-robber-iv)|Medium|||||||||||
|[2559](https://leetcode.cn/problems/count-vowel-strings-in-ranges)|Medium|||||||||||
|[2558](https://leetcode.cn/problems/take-gifts-from-the-richest-pile)|Easy|||||||||||
|[2548](https://leetcode.cn/problems/maximum-price-to-fill-a-bag)|Medium|||||||||||
|[2644](https://leetcode.cn/problems/find-the-maximum-divisibility-score)|Easy|||||||||||
|[2593](https://leetcode.cn/problems/find-score-of-an-array-after-marking-all-elements)|Medium|||||||||||
|[2597](https://leetcode.cn/problems/the-number-of-beautiful-subsets)|Medium|||||||||||
|[2617](https://leetcode.cn/problems/minimum-number-of-visited-cells-in-a-grid)|Hard|||||||||||
|[2562](https://leetcode.cn/problems/find-the-array-concatenation-value)|Easy|||||||||||
|[2563](https://leetcode.cn/problems/count-the-number-of-fair-pairs)|Medium|||||||||||
|[2564](https://leetcode.cn/problems/substring-xor-queries)|Medium|||||||||||
|[2565](https://leetcode.cn/problems/subsequence-with-the-minimum-score)|Hard|||||||||||
|[2557](https://leetcode.cn/problems/maximum-number-of-integers-to-choose-from-a-range-ii)|Medium|||||||||||
|[2569](https://leetcode.cn/problems/handling-sum-queries-after-update)|Hard|||||||||||
|[2566](https://leetcode.cn/problems/maximum-difference-by-remapping-a-digit)|Easy|||||||||||
|[2568](https://leetcode.cn/problems/minimum-impossible-or)|Medium|||||||||||
|[2567](https://leetcode.cn/problems/minimum-score-by-changing-two-elements)|Medium|||||||||||
|[2570](https://leetcode.cn/problems/merge-two-2d-arrays-by-summing-values)|Easy|||||||||||
|[2573](https://leetcode.cn/problems/find-the-string-with-lcp)|Hard|||||||||||
|[2572](https://leetcode.cn/problems/count-the-number-of-square-free-subsets)|Medium|||||||||||
|[2571](https://leetcode.cn/problems/minimum-operations-to-reduce-an-integer-to-0)|Medium|||||||||||
|[2577](https://leetcode.cn/problems/minimum-time-to-visit-a-cell-in-a-grid)|Hard|||||||||||
|[2576](https://leetcode.cn/problems/find-the-maximum-number-of-marked-indices)|Medium|||||||||||
|[2575](https://leetcode.cn/problems/find-the-divisibility-array-of-a-string)|Medium|||||||||||
|[2574](https://leetcode.cn/problems/left-and-right-sum-differences)|Easy|||||||||||
|[2600](https://leetcode.cn/problems/k-items-with-the-maximum-sum)|Easy|||||||||||
|[2601](https://leetcode.cn/problems/prime-subtraction-operation)|Medium|||||||||||
|[2603](https://leetcode.cn/problems/collect-coins-in-a-tree)|Hard|||||||||||
|[2602](https://leetcode.cn/problems/minimum-operations-to-make-all-array-elements-equal)|Medium|||||||||||
|[2613](https://leetcode.cn/problems/beautiful-pairs)|Hard|||||||||||
|[2616](https://leetcode.cn/problems/minimize-the-maximum-difference-of-pairs)|Medium|||||||||||
|[2615](https://leetcode.cn/problems/sum-of-distances)|Medium|||||||||||
|[2614](https://leetcode.cn/problems/prime-in-diagonal)|Easy|||||||||||
|[2609](https://leetcode.cn/problems/find-the-longest-balanced-substring-of-a-binary-string)|Easy|||||||||||
|[2610](https://leetcode.cn/problems/convert-an-array-into-a-2d-array-with-conditions)|Medium|||||||||||
|[2611](https://leetcode.cn/problems/mice-and-cheese)|Medium|||||||||||
|[2612](https://leetcode.cn/problems/minimum-reverse-operations)|Hard|||||||||||
|[2678](https://leetcode.cn/problems/number-of-senior-citizens)|Easy|||||||||||
|[2679](https://leetcode.cn/problems/sum-in-a-matrix)|Medium|||||||||||
|[2702](https://leetcode.cn/problems/minimum-operations-to-make-numbers-non-positive)|Hard|||||||||||
|[2680](https://leetcode.cn/problems/maximum-or)|Medium|||||||||||
|[2645](https://leetcode.cn/problems/minimum-additions-to-make-valid-string)|Medium|||||||||||
|[2643](https://leetcode.cn/problems/row-with-maximum-ones)|Easy|||||||||||
|[2638](https://leetcode.cn/problems/count-the-number-of-k-free-subsets)|Medium|||||||||||
|[2646](https://leetcode.cn/problems/minimize-the-total-price-of-the-trips)|Hard|||||||||||
|[2651](https://leetcode.cn/problems/calculate-delayed-arrival-time)|Easy|||||||||||
|[2653](https://leetcode.cn/problems/sliding-subarray-beauty)|Medium|||||||||||
|[2652](https://leetcode.cn/problems/sum-multiples)|Easy|||||||||||
|[2654](https://leetcode.cn/problems/minimum-number-of-operations-to-make-all-array-elements-equal-to-1)|Medium|||||||||||
|[2708](https://leetcode.cn/problems/maximum-strength-of-a-group)|Medium|||||||||||
|[2707](https://leetcode.cn/problems/extra-characters-in-a-string)|Medium|||||||||||
|[2706](https://leetcode.cn/problems/buy-two-chocolates)|Easy|||||||||||
|[2719](https://leetcode.cn/problems/count-of-integers)|Hard|||||||||||
|[2658](https://leetcode.cn/problems/maximum-number-of-fish-in-a-grid)|Medium|||||||||||
|[2659](https://leetcode.cn/problems/make-array-empty)|Hard|||||||||||
|[2657](https://leetcode.cn/problems/find-the-prefix-common-array-of-two-arrays)|Medium|||||||||||
|[2656](https://leetcode.cn/problems/maximum-sum-with-exactly-k-elements)|Easy|||||||||||
|[2670](https://leetcode.cn/problems/find-the-distinct-difference-array)|Easy|||||||||||
|[2671](https://leetcode.cn/problems/frequency-tracker)|Medium|||||||||||
|[2672](https://leetcode.cn/problems/number-of-adjacent-elements-with-the-same-color)|Medium|||||||||||
|[2673](https://leetcode.cn/problems/make-costs-of-paths-equal-in-a-binary-tree)|Medium|||||||||||
|[2681](https://leetcode.cn/problems/power-of-heroes)|Hard|||||||||||
|[2717](https://leetcode.cn/problems/semi-ordered-permutation)|Easy|||||||||||
|[2730](https://leetcode.cn/problems/find-the-longest-semi-repetitive-substring)|Medium|||||||||||
|[2731](https://leetcode.cn/problems/movement-of-robots)|Medium|||||||||||
|[2682](https://leetcode.cn/problems/find-the-losers-of-the-circular-game)|Easy|||||||||||
|[2683](https://leetcode.cn/problems/neighboring-bitwise-xor)|Medium|||||||||||
|[2685](https://leetcode.cn/problems/count-the-number-of-complete-components)|Medium|||||||||||
|[2684](https://leetcode.cn/problems/maximum-number-of-moves-in-a-grid)|Medium|||||||||||
|[2696](https://leetcode.cn/problems/minimum-string-length-after-removing-substrings)|Easy|||||||||||
|[2711](https://leetcode.cn/problems/difference-of-number-of-distinct-values-on-diagonals)|Medium|||||||||||
|[2698](https://leetcode.cn/problems/find-the-punishment-number-of-an-integer)|Medium|||||||||||
|[2699](https://leetcode.cn/problems/modify-graph-edge-weights)|Hard|||||||||||
|[2742](https://leetcode.cn/problems/painting-the-walls)|Hard|||||||||||
|[2735](https://leetcode.cn/problems/collecting-chocolates)|Medium|||||||||||
|[2829](https://leetcode.cn/problems/determine-the-minimum-sum-of-a-k-avoiding-array)|Medium|||||||||||
|[2769](https://leetcode.cn/problems/find-the-maximum-achievable-number)|Easy|||||||||||
|[2655](https://leetcode.cn/problems/find-maximal-uncovered-ranges)|Medium|||||||||||
|[2697](https://leetcode.cn/problems/lexicographically-smallest-palindrome)|Easy|||||||||||
|[2712](https://leetcode.cn/problems/minimum-cost-to-make-all-characters-equal)|Medium|||||||||||
|[2713](https://leetcode.cn/problems/maximum-strictly-increasing-cells-in-a-matrix)|Hard|||||||||||
|[2710](https://leetcode.cn/problems/remove-trailing-zeros-from-a-string)|Easy|||||||||||
|[2664](https://leetcode.cn/problems/the-knights-tour)|Medium|||||||||||
|[2729](https://leetcode.cn/problems/check-if-the-number-is-fascinating)|Easy|||||||||||
|[2716](https://leetcode.cn/problems/minimize-string-length)|Easy|||||||||||
|[2732](https://leetcode.cn/problems/find-a-good-subset-of-the-matrix)|Hard|||||||||||
|[2709](https://leetcode.cn/problems/greatest-common-divisor-traversal)|Hard|||||||||||
|[2734](https://leetcode.cn/problems/lexicographically-smallest-string-after-substring-operation)|Medium|||||||||||
|[2748](https://leetcode.cn/problems/number-of-beautiful-pairs)|Easy|||||||||||
|[2831](https://leetcode.cn/problems/find-the-longest-equal-subarray)|Medium|||||||||||
|[2747](https://leetcode.cn/problems/count-zero-request-servers)|Medium|||||||||||
|[2766](https://leetcode.cn/problems/relocate-marbles)|Medium|||||||||||
|[2674](https://leetcode.cn/problems/split-a-circular-linked-list)|Medium|||||||||||
|[2733](https://leetcode.cn/problems/neither-minimum-nor-maximum)|Easy|||||||||||
|[2749](https://leetcode.cn/problems/minimum-operations-to-make-the-integer-zero)|Medium|||||||||||
|[2718](https://leetcode.cn/problems/sum-of-matrix-after-queries)|Medium|||||||||||
|[2736](https://leetcode.cn/problems/maximum-sum-queries)|Hard|||||||||||
|[2689](https://leetcode.cn/problems/extract-kth-character-from-the-rope-tree)|Easy|||||||||||
|[2778](https://leetcode.cn/problems/sum-of-squares-of-special-elements)|Easy|||||||||||
|[2740](https://leetcode.cn/problems/find-the-value-of-the-partition)|Medium|||||||||||
|[2751](https://leetcode.cn/problems/robot-collisions)|Hard|||||||||||
|[2744](https://leetcode.cn/problems/find-maximum-number-of-string-pairs)|Easy|||||||||||
|[2741](https://leetcode.cn/problems/special-permutations)|Medium|||||||||||
|[2763](https://leetcode.cn/problems/sum-of-imbalance-numbers-of-all-subarrays)|Hard|||||||||||
|[2745](https://leetcode.cn/problems/construct-the-longest-new-string)|Medium|||||||||||
|[2746](https://leetcode.cn/problems/decremental-string-concatenation)|Medium|||||||||||
|[2770](https://leetcode.cn/problems/maximum-number-of-jumps-to-reach-the-last-index)|Medium|||||||||||
|[2799](https://leetcode.cn/problems/count-complete-subarrays-in-an-array)|Medium|||||||||||
|[2739](https://leetcode.cn/problems/total-distance-traveled)|Easy|||||||||||
|[2714](https://leetcode.cn/problems/find-shortest-path-with-k-hops)|Hard|||||||||||
|[2760](https://leetcode.cn/problems/longest-even-odd-subarray-with-threshold)|Easy|||||||||||
|[2750](https://leetcode.cn/problems/ways-to-split-array-into-good-subarrays)|Medium|||||||||||
|[2762](https://leetcode.cn/problems/continuous-subarrays)|Medium|||||||||||
|[2771](https://leetcode.cn/problems/longest-non-decreasing-subarray-from-two-arrays)|Medium|||||||||||
|[2765](https://leetcode.cn/problems/longest-alternating-subarray)|Easy|||||||||||
|[2816](https://leetcode.cn/problems/double-a-number-represented-as-a-linked-list)|Medium|||||||||||
|[2789](https://leetcode.cn/problems/largest-element-in-an-array-after-merge-operations)|Medium|||||||||||
|[2761](https://leetcode.cn/problems/prime-pairs-with-target-sum)|Medium|||||||||||
|[2728](https://leetcode.cn/problems/count-houses-in-a-circular-street)|Easy|||||||||||
|[2798](https://leetcode.cn/problems/number-of-employees-who-met-the-target)|Easy|||||||||||
|[2800](https://leetcode.cn/problems/shortest-string-that-contains-three-strings)|Medium|||||||||||
|[2772](https://leetcode.cn/problems/apply-operations-to-make-all-array-elements-equal-to-zero)|Medium|||||||||||
|[2911](https://leetcode.cn/problems/minimum-changes-to-make-k-semi-palindromes)|Hard|||||||||||
|[2737](https://leetcode.cn/problems/find-the-closest-marked-node)|Medium|||||||||||
|[2788](https://leetcode.cn/problems/split-strings-by-separator)|Easy|||||||||||
|[2787](https://leetcode.cn/problems/ways-to-express-an-integer-as-sum-of-powers)|Medium|||||||||||
|[2767](https://leetcode.cn/problems/partition-string-into-minimum-beautiful-substrings)|Medium|||||||||||
|[2781](https://leetcode.cn/problems/length-of-the-longest-valid-substring)|Hard|||||||||||
|[2810](https://leetcode.cn/problems/faulty-keyboard)|Easy|||||||||||
|[2785](https://leetcode.cn/problems/sort-vowels-in-a-string)|Medium|||||||||||
|[2780](https://leetcode.cn/problems/minimum-index-of-a-valid-split)|Medium|||||||||||
|[2768](https://leetcode.cn/problems/number-of-black-blocks)|Medium|||||||||||
|[2743](https://leetcode.cn/problems/count-substrings-without-repeating-character)|Medium|||||||||||
|[2779](https://leetcode.cn/problems/maximum-beauty-of-an-array-after-applying-operation)|Medium|||||||||||
|[2784](https://leetcode.cn/problems/check-if-array-is-good)|Easy|||||||||||
|[2786](https://leetcode.cn/problems/visit-array-positions-to-maximize-score)|Medium|||||||||||
|[2813](https://leetcode.cn/problems/maximum-elegance-of-a-k-length-subsequence)|Hard|||||||||||
|[2753](https://leetcode.cn/problems/count-houses-in-a-circular-street-ii)|Hard|||||||||||
|[2815](https://leetcode.cn/problems/max-pair-sum-in-an-array)|Easy|||||||||||
|[2807](https://leetcode.cn/problems/insert-greatest-common-divisors-in-linked-list)|Medium|||||||||||
|[2826](https://leetcode.cn/problems/sorting-three-groups)|Medium|||||||||||
|[2791](https://leetcode.cn/problems/count-paths-that-can-form-a-palindrome-in-a-tree)|Hard|||||||||||
|[2812](https://leetcode.cn/problems/find-the-safest-path-in-a-grid)|Medium|||||||||||
|[2845](https://leetcode.cn/problems/count-of-interesting-subarrays)|Medium|||||||||||
|[2811](https://leetcode.cn/problems/check-if-it-is-possible-to-split-array)|Medium|||||||||||
|[2824](https://leetcode.cn/problems/count-pairs-whose-sum-is-less-than-target)|Easy|||||||||||
|[2764](https://leetcode.cn/problems/is-array-a-preorder-of-some-binary-tree)|Medium|||||||||||
|[2790](https://leetcode.cn/problems/maximum-number-of-groups-with-increasing-length)|Hard|||||||||||
|[2808](https://leetcode.cn/problems/minimum-seconds-to-equalize-a-circular-array)|Medium|||||||||||
|[2801](https://leetcode.cn/problems/count-stepping-numbers-in-range)|Hard|||||||||||
|[2773](https://leetcode.cn/problems/height-of-special-binary-tree)|Medium|||||||||||
|[2809](https://leetcode.cn/problems/minimum-time-to-make-array-sum-at-most-x)|Hard|||||||||||
|[2857](https://leetcode.cn/problems/count-pairs-of-points-with-distance-k)|Medium|||||||||||
|[2841](https://leetcode.cn/problems/maximum-sum-of-almost-unique-subarray)|Medium|||||||||||
|[2806](https://leetcode.cn/problems/account-balance-after-rounded-purchase)|Easy|||||||||||
|[2828](https://leetcode.cn/problems/check-if-a-string-is-an-acronym-of-words)|Easy|||||||||||
|[2840](https://leetcode.cn/problems/check-if-strings-can-be-made-equal-with-operations-ii)|Medium|||||||||||
|[2830](https://leetcode.cn/problems/maximize-the-profit-as-the-salesman)|Medium|||||||||||
|[2782](https://leetcode.cn/problems/number-of-unique-categories)|Medium|||||||||||
|[2843](https://leetcode.cn/problems/count-symmetric-integers)|Easy|||||||||||
|[2839](https://leetcode.cn/problems/check-if-strings-can-be-made-equal-with-operations-i)|Easy|||||||||||
|[2817](https://leetcode.cn/problems/minimum-absolute-difference-between-elements-with-constraint)|Medium|||||||||||
|[2818](https://leetcode.cn/problems/apply-operations-to-maximize-score)|Hard|||||||||||
|[2792](https://leetcode.cn/problems/count-nodes-that-are-great-enough)|Hard|||||||||||
|[2827](https://leetcode.cn/problems/number-of-beautiful-integers-in-the-range)|Hard|||||||||||
|[2825](https://leetcode.cn/problems/make-string-a-subsequence-using-cyclic-increments)|Medium|||||||||||
|[2833](https://leetcode.cn/problems/furthest-point-from-origin)|Easy|||||||||||
|[2851](https://leetcode.cn/problems/string-transformation)|Hard|||||||||||
|[2835](https://leetcode.cn/problems/minimum-operations-to-form-subsequence-with-target-sum)|Hard|||||||||||
|[2834](https://leetcode.cn/problems/find-the-minimum-possible-sum-of-a-beautiful-array)|Medium|||||||||||
|[2802](https://leetcode.cn/problems/find-the-k-th-lucky-number)|Medium|||||||||||
|[2906](https://leetcode.cn/problems/construct-product-matrix)|Medium|||||||||||
|[2836](https://leetcode.cn/problems/maximize-value-of-function-in-a-ball-passing-game)|Hard|||||||||||
|[2896](https://leetcode.cn/problems/apply-operations-to-make-two-strings-equal)|Medium|||||||||||
|[2848](https://leetcode.cn/problems/points-that-intersect-with-cars)|Easy|||||||||||
|[2814](https://leetcode.cn/problems/minimum-time-takes-to-reach-destination-without-drowning)|Hard|||||||||||
|[2869](https://leetcode.cn/problems/minimum-operations-to-collect-elements)|Easy|||||||||||
|[2855](https://leetcode.cn/problems/minimum-right-shifts-to-sort-the-array)|Easy|||||||||||
|[2844](https://leetcode.cn/problems/minimum-operations-to-make-a-special-number)|Medium|||||||||||
|[2862](https://leetcode.cn/problems/maximum-element-sum-of-a-complete-subset-of-indices)|Hard|||||||||||
|[2864](https://leetcode.cn/problems/maximum-odd-binary-number)|Easy|||||||||||
|[2849](https://leetcode.cn/problems/determine-if-a-cell-is-reachable-at-a-given-time)|Medium|||||||||||
|[2842](https://leetcode.cn/problems/count-k-subsequences-of-a-string-with-maximum-beauty)|Hard|||||||||||
|[2872](https://leetcode.cn/problems/maximum-number-of-k-divisible-components)|Hard|||||||||||
|[2819](https://leetcode.cn/problems/minimum-relative-loss-after-buying-chocolates)|Hard|||||||||||
|[2846](https://leetcode.cn/problems/minimum-edge-weight-equilibrium-queries-in-a-tree)|Hard|||||||||||
|[2871](https://leetcode.cn/problems/split-array-into-maximum-number-of-subarrays)|Medium|||||||||||
|[2856](https://leetcode.cn/problems/minimum-array-length-after-pair-removals)|Medium|||||||||||
|[2832](https://leetcode.cn/problems/maximal-range-that-each-element-is-maximum-in-it)|Medium|||||||||||
|[2902](https://leetcode.cn/problems/count-of-sub-multisets-with-bounded-sum)|Hard|||||||||||
|[2850](https://leetcode.cn/problems/minimum-moves-to-spread-stones-over-grid)|Medium|||||||||||
|[2859](https://leetcode.cn/problems/sum-of-values-at-indices-with-k-set-bits)|Easy|||||||||||
|[2870](https://leetcode.cn/problems/minimum-number-of-operations-to-make-array-empty)|Medium|||||||||||
|[2861](https://leetcode.cn/problems/maximum-number-of-alloys)|Medium|||||||||||
|[2838](https://leetcode.cn/problems/maximum-coins-heroes-can-collect)|Medium|||||||||||
|[2860](https://leetcode.cn/problems/happy-students)|Medium|||||||||||
|[2858](https://leetcode.cn/problems/minimum-edge-reversals-so-every-node-is-reachable)|Hard|||||||||||
|[2915](https://leetcode.cn/problems/length-of-the-longest-subsequence-that-sums-to-target)|Medium|||||||||||
|[2931](https://leetcode.cn/problems/maximum-spending-after-buying-items)|Hard|||||||||||
|[2847](https://leetcode.cn/problems/smallest-number-with-given-digit-product)|Medium|||||||||||
|[2867](https://leetcode.cn/problems/count-valid-paths-in-a-tree)|Hard|||||||||||
|[2866](https://leetcode.cn/problems/beautiful-towers-ii)|Medium|||||||||||
|[2865](https://leetcode.cn/problems/beautiful-towers-i)|Medium|||||||||||
|[2916](https://leetcode.cn/problems/subarrays-distinct-element-sum-of-squares-ii)|Hard|||||||||||
|[2876](https://leetcode.cn/problems/count-visited-nodes-in-a-directed-graph)|Hard|||||||||||
|[2875](https://leetcode.cn/problems/minimum-size-subarray-in-infinite-array)|Medium|||||||||||
|[2901](https://leetcode.cn/problems/longest-unequal-adjacent-groups-subsequence-ii)|Medium|||||||||||
|[2900](https://leetcode.cn/problems/longest-unequal-adjacent-groups-subsequence-i)|Easy|||||||||||
|[2852](https://leetcode.cn/problems/sum-of-remoteness-of-all-cells)|Medium|||||||||||
|[2904](https://leetcode.cn/problems/shortest-and-lexicographically-smallest-beautiful-string)|Medium|||||||||||
|[2895](https://leetcode.cn/problems/minimum-processing-time)|Medium|||||||||||
|[2874](https://leetcode.cn/problems/maximum-value-of-an-ordered-triplet-ii)|Medium|||||||||||
|[2897](https://leetcode.cn/problems/apply-operations-on-array-to-maximize-sum-of-squares)|Hard|||||||||||
|[2873](https://leetcode.cn/problems/maximum-value-of-an-ordered-triplet-i)|Easy|||||||||||
|[2863](https://leetcode.cn/problems/maximum-length-of-semi-decreasing-subarrays)|Medium|||||||||||
|[2913](https://leetcode.cn/problems/subarrays-distinct-element-sum-of-squares-i)|Easy|||||||||||
|[2899](https://leetcode.cn/problems/last-visited-integers)|Easy|||||||||||
|[2903](https://leetcode.cn/problems/find-indices-with-index-and-value-difference-i)|Easy|||||||||||
|[2910](https://leetcode.cn/problems/minimum-number-of-groups-to-create-a-valid-assignment)|Medium|||||||||||
|[2905](https://leetcode.cn/problems/find-indices-with-index-and-value-difference-ii)|Medium|||||||||||
|[2918](https://leetcode.cn/problems/minimum-equal-sum-of-two-arrays-after-replacing-zeros)|Medium|||||||||||
|[2894](https://leetcode.cn/problems/divisible-and-non-divisible-sums-difference)|Easy|||||||||||
|[2868](https://leetcode.cn/problems/the-wording-game)|Hard|||||||||||
|[2914](https://leetcode.cn/problems/minimum-number-of-changes-to-make-binary-string-beautiful)|Medium|||||||||||
|[2908](https://leetcode.cn/problems/minimum-sum-of-mountain-triplets-i)|Easy|||||||||||
|[2892](https://leetcode.cn/problems/minimizing-array-after-replacing-pairs-with-their-product)|Medium|||||||||||
|[2919](https://leetcode.cn/problems/minimum-increment-operations-to-make-array-beautiful)|Medium|||||||||||
|[2920](https://leetcode.cn/problems/maximum-points-after-collecting-coins-from-all-nodes)|Hard|||||||||||
|[2940](https://leetcode.cn/problems/find-building-where-alice-and-bob-can-meet)|Hard|||||||||||
|[2898](https://leetcode.cn/problems/maximum-linear-stock-score)|Medium|||||||||||
|[2917](https://leetcode.cn/problems/find-the-k-or-of-an-array)|Easy|||||||||||
|[2926](https://leetcode.cn/problems/maximum-balanced-subsequence-sum)|Hard|||||||||||
|[2909](https://leetcode.cn/problems/minimum-sum-of-mountain-triplets-ii)|Medium|||||||||||
|[2907](https://leetcode.cn/problems/maximum-profitable-triplets-with-increasing-prices-i)|Medium|||||||||||
|[2923](https://leetcode.cn/problems/find-champion-i)|Easy|||||||||||
|[2924](https://leetcode.cn/problems/find-champion-ii)|Medium|||||||||||
|[2934](https://leetcode.cn/problems/minimum-operations-to-maximize-last-elements-in-arrays)|Medium|||||||||||
|[2925](https://leetcode.cn/problems/maximum-score-after-applying-operations-on-a-tree)|Medium|||||||||||
|[2939](https://leetcode.cn/problems/maximum-xor-product)|Medium|||||||||||
|[2932](https://leetcode.cn/problems/maximum-strong-pair-xor-i)|Easy|||||||||||
|[2942](https://leetcode.cn/problems/find-words-containing-character)|Easy|||||||||||
|[2938](https://leetcode.cn/problems/separate-black-and-white-balls)|Medium|||||||||||
|[2968](https://leetcode.cn/problems/apply-operations-to-maximize-frequency-score)|Hard|||||||||||
|[2935](https://leetcode.cn/problems/maximum-strong-pair-xor-ii)|Hard|||||||||||
|[2912](https://leetcode.cn/problems/number-of-ways-to-reach-destination-in-the-grid)|Hard|||||||||||
|[2928](https://leetcode.cn/problems/distribute-candies-among-children-i)|Easy|||||||||||
|[2930](https://leetcode.cn/problems/number-of-strings-which-can-be-rearranged-to-contain-substring)|Medium|||||||||||
|[2929](https://leetcode.cn/problems/distribute-candies-among-children-ii)|Medium|||||||||||
|[2933](https://leetcode.cn/problems/high-access-employees)|Medium|||||||||||
|[2983](https://leetcode.cn/problems/palindrome-rearrangement-queries)|Hard|||||||||||
|[2921](https://leetcode.cn/problems/maximum-profitable-triplets-with-increasing-prices-ii)|Hard|||||||||||
|[2956](https://leetcode.cn/problems/find-common-elements-between-two-arrays)|Easy|||||||||||
|[2937](https://leetcode.cn/problems/make-three-strings-equal)|Easy|||||||||||
|[2949](https://leetcode.cn/problems/count-beautiful-substrings-ii)|Hard|||||||||||
|[2944](https://leetcode.cn/problems/minimum-number-of-coins-for-fruits)|Medium|||||||||||
|[2947](https://leetcode.cn/problems/count-beautiful-substrings-i)|Medium|||||||||||
|[2945](https://leetcode.cn/problems/find-maximum-non-decreasing-array-length)|Hard|||||||||||
|[2963](https://leetcode.cn/problems/count-the-number-of-good-partitions)|Hard|||||||||||
|[2962](https://leetcode.cn/problems/count-subarrays-where-max-element-appears-at-least-k-times)|Medium|||||||||||
|[2943](https://leetcode.cn/problems/maximize-area-of-square-hole-in-grid)|Medium|||||||||||
|[2946](https://leetcode.cn/problems/matrix-similarity-after-cyclic-shifts)|Easy|||||||||||
|[2927](https://leetcode.cn/problems/distribute-candies-among-children-iii)|Hard|||||||||||
|[2959](https://leetcode.cn/problems/number-of-possible-sets-of-closing-branches)|Hard|||||||||||
|[2973](https://leetcode.cn/problems/find-number-of-coins-to-place-in-tree-nodes)|Hard|||||||||||
|[2948](https://leetcode.cn/problems/make-lexicographically-smallest-array-by-swapping-elements)|Medium|||||||||||
|[2960](https://leetcode.cn/problems/count-tested-devices-after-test-operations)|Easy|||||||||||
|[2951](https://leetcode.cn/problems/find-the-peaks)|Easy|||||||||||
|[2936](https://leetcode.cn/problems/number-of-equal-numbers-blocks)|Medium|||||||||||
|[2953](https://leetcode.cn/problems/count-complete-substrings)|Hard|||||||||||
|[2954](https://leetcode.cn/problems/count-the-number-of-infection-sequences)|Hard|||||||||||
|[2958](https://leetcode.cn/problems/length-of-longest-subarray-with-at-most-k-frequency)|Medium|||||||||||
|[2974](https://leetcode.cn/problems/minimum-number-game)|Easy|||||||||||
|[2965](https://leetcode.cn/problems/find-missing-and-repeated-values)|Easy|||||||||||
|[3002](https://leetcode.cn/problems/maximum-size-of-a-set-after-removals)|Medium|||||||||||
|[2967](https://leetcode.cn/problems/minimum-cost-to-make-array-equalindromic)|Medium|||||||||||
|[2957](https://leetcode.cn/problems/remove-adjacent-almost-equal-characters)|Medium|||||||||||
|[2952](https://leetcode.cn/problems/minimum-number-of-coins-to-be-added)|Medium|||||||||||
|[2941](https://leetcode.cn/problems/maximum-gcd-sum-of-a-subarray)|Hard|||||||||||
|[3003](https://leetcode.cn/problems/maximize-the-number-of-partitions-after-operations)|Hard|||||||||||
|[2961](https://leetcode.cn/problems/double-modular-exponentiation)|Medium|||||||||||
|[2976](https://leetcode.cn/problems/minimum-cost-to-convert-string-i)|Medium|||||||||||
|[2996](https://leetcode.cn/problems/smallest-missing-integer-greater-than-sequential-prefix-sum)|Easy|||||||||||
|[2950](https://leetcode.cn/problems/number-of-divisible-substrings)|Medium|||||||||||
|[2977](https://leetcode.cn/problems/minimum-cost-to-convert-string-ii)|Hard|||||||||||
|[2998](https://leetcode.cn/problems/minimum-number-of-operations-to-make-x-and-y-equal)|Medium|||||||||||
|[3007](https://leetcode.cn/problems/maximum-number-that-sum-of-the-prices-is-less-than-or-equal-to-k)|Medium|||||||||||
|[2966](https://leetcode.cn/problems/divide-array-into-arrays-with-max-difference)|Medium|||||||||||
|[3005](https://leetcode.cn/problems/count-elements-with-maximum-frequency)|Easy|||||||||||
|[2999](https://leetcode.cn/problems/count-the-number-of-powerful-integers)|Hard|||||||||||
|[3012](https://leetcode.cn/problems/minimize-length-of-array-using-operations)|Medium|||||||||||
|[3006](https://leetcode.cn/problems/find-beautiful-indices-in-the-given-array-i)|Medium|||||||||||
|[2980](https://leetcode.cn/problems/check-if-bitwise-or-has-trailing-zeros)|Easy|||||||||||
|[2955](https://leetcode.cn/problems/number-of-same-end-substrings)|Medium|||||||||||
|[2972](https://leetcode.cn/problems/count-the-number-of-incremovable-subarrays-ii)|Hard|||||||||||
|[2997](https://leetcode.cn/problems/minimum-number-of-operations-to-make-array-xor-equal-to-k)|Medium|||||||||||
|[2975](https://leetcode.cn/problems/maximum-square-area-by-removing-fences-from-a-field)|Medium|||||||||||
|[3000](https://leetcode.cn/problems/maximum-area-of-longest-diagonal-rectangle)|Easy|||||||||||
|[2970](https://leetcode.cn/problems/count-the-number-of-incremovable-subarrays-i)|Easy|||||||||||
|[2964](https://leetcode.cn/problems/number-of-divisible-triplet-sums)|Medium|||||||||||
|[3013](https://leetcode.cn/problems/divide-an-array-into-subarrays-with-minimum-cost-ii)|Hard|||||||||||
|[3022](https://leetcode.cn/problems/minimize-or-of-remaining-elements-using-operations)|Hard|||||||||||
|[2971](https://leetcode.cn/problems/find-polygon-with-the-largest-perimeter)|Medium|||||||||||
|[3010](https://leetcode.cn/problems/divide-an-array-into-subarrays-with-minimum-cost-i)|Easy|||||||||||
|[3207](https://leetcode.cn/problems/maximum-points-after-enemy-battles)|Medium|||||||||||
|[3026](https://leetcode.cn/problems/maximum-good-subarray-sum)|Medium|||||||||||
|[2982](https://leetcode.cn/problems/find-longest-special-substring-that-occurs-thrice-ii)|Medium|||||||||||
|[2981](https://leetcode.cn/problems/find-longest-special-substring-that-occurs-thrice-i)|Medium|||||||||||
|[2969](https://leetcode.cn/problems/minimum-number-of-coins-for-fruits-ii)|Hard|||||||||||
|[3034](https://leetcode.cn/problems/number-of-subarrays-that-match-a-pattern-i)|Medium|||||||||||
|[3001](https://leetcode.cn/problems/minimum-moves-to-capture-the-queen)|Medium|||||||||||
|[3015](https://leetcode.cn/problems/count-the-number-of-houses-at-a-certain-distance-i)|Medium|||||||||||
|[3030](https://leetcode.cn/problems/find-the-grid-of-region-average)|Medium|||||||||||
|[2979](https://leetcode.cn/problems/most-expensive-item-that-can-not-be-bought)|Medium|||||||||||
|[3014](https://leetcode.cn/problems/minimum-number-of-pushes-to-type-word-i)|Easy|||||||||||
|[3016](https://leetcode.cn/problems/minimum-number-of-pushes-to-type-word-ii)|Medium|||||||||||
|[3027](https://leetcode.cn/problems/find-the-number-of-ways-to-place-people-ii)|Hard|||||||||||
|[3025](https://leetcode.cn/problems/find-the-number-of-ways-to-place-people-i)|Medium|||||||||||
|[3021](https://leetcode.cn/problems/alice-and-bob-playing-flower-game)|Medium|||||||||||
|[3049](https://leetcode.cn/problems/earliest-second-to-mark-indices-ii)|Hard|||||||||||
|[3036](https://leetcode.cn/problems/number-of-subarrays-that-match-a-pattern-ii)|Hard|||||||||||
|[3011](https://leetcode.cn/problems/find-if-array-can-be-sorted)|Medium|||||||||||
|[3048](https://leetcode.cn/problems/earliest-second-to-mark-indices-i)|Medium|||||||||||
|[2992](https://leetcode.cn/problems/number-of-self-divisible-permutations)|Medium|||||||||||
|[3031](https://leetcode.cn/problems/minimum-time-to-revert-word-to-initial-state-ii)|Hard|||||||||||
|[3029](https://leetcode.cn/problems/minimum-time-to-revert-word-to-initial-state-i)|Medium|||||||||||
|[3041](https://leetcode.cn/problems/maximize-consecutive-elements-in-an-array-after-modification)|Hard|||||||||||
|[3020](https://leetcode.cn/problems/find-the-maximum-number-of-elements-in-subset)|Medium|||||||||||
|[3008](https://leetcode.cn/problems/find-beautiful-indices-in-the-given-array-ii)|Hard|||||||||||
|[3004](https://leetcode.cn/problems/maximum-subtree-of-the-same-color)|Medium|||||||||||
|[3045](https://leetcode.cn/problems/count-prefix-and-suffix-pairs-ii)|Hard|||||||||||
|[3080](https://leetcode.cn/problems/mark-elements-on-array-by-performing-queries)|Medium|||||||||||
|[3068](https://leetcode.cn/problems/find-the-maximum-sum-of-node-values)|Hard|||||||||||
|[3039](https://leetcode.cn/problems/apply-operations-to-make-string-empty)|Medium|||||||||||
|[3042](https://leetcode.cn/problems/count-prefix-and-suffix-pairs-i)|Easy|||||||||||
|[3017](https://leetcode.cn/problems/count-the-number-of-houses-at-a-certain-distance-ii)|Hard|||||||||||
|[3028](https://leetcode.cn/problems/ant-on-the-boundary)|Easy|||||||||||
|[3019](https://leetcode.cn/problems/number-of-changing-keys)|Easy|||||||||||
|[3077](https://leetcode.cn/problems/maximum-strength-of-k-disjoint-subarrays)|Hard|||||||||||
|[3044](https://leetcode.cn/problems/most-frequent-prime)|Medium|||||||||||
|[3009](https://leetcode.cn/problems/maximum-number-of-intersections-on-the-chart)|Hard|||||||||||
|[3098](https://leetcode.cn/problems/find-the-sum-of-subsequence-powers)|Hard|||||||||||
|[3035](https://leetcode.cn/problems/maximum-palindromes-after-operations)|Medium|||||||||||
|[3040](https://leetcode.cn/problems/maximum-number-of-operations-with-the-same-score-ii)|Medium|||||||||||
|[3038](https://leetcode.cn/problems/maximum-number-of-operations-with-the-same-score-i)|Easy|||||||||||
|[3024](https://leetcode.cn/problems/type-of-triangle)|Easy|||||||||||
|[3018](https://leetcode.cn/problems/maximum-number-of-removal-queries-that-can-be-processed-i)|Hard|||||||||||
|[3046](https://leetcode.cn/problems/split-the-array)|Easy|||||||||||
|[3047](https://leetcode.cn/problems/find-the-largest-area-of-square-inside-two-rectangles)|Medium|||||||||||
|[3067](https://leetcode.cn/problems/count-pairs-of-connectable-servers-in-a-weighted-tree-network)|Medium|||||||||||
|[3086](https://leetcode.cn/problems/minimum-moves-to-pick-k-ones)|Hard|||||||||||
|[3091](https://leetcode.cn/problems/apply-operations-to-make-sum-of-array-greater-than-or-equal-to-k)|Medium|||||||||||
|[3043](https://leetcode.cn/problems/find-the-length-of-the-longest-common-prefix)|Medium|||||||||||
|[3033](https://leetcode.cn/problems/modify-the-matrix)|Easy|||||||||||
|[3065](https://leetcode.cn/problems/minimum-operations-to-exceed-threshold-value-i)|Easy|||||||||||
|[3066](https://leetcode.cn/problems/minimum-operations-to-exceed-threshold-value-ii)|Medium|||||||||||
|[3023](https://leetcode.cn/problems/find-pattern-in-infinite-stream-i)|Medium|||||||||||
|[3074](https://leetcode.cn/problems/apple-redistribution-into-boxes)|Easy|||||||||||
|[3071](https://leetcode.cn/problems/minimum-operations-to-write-the-letter-y-on-a-grid)|Medium|||||||||||
|[3100](https://leetcode.cn/problems/water-bottles-ii)|Medium|||||||||||
|[3084](https://leetcode.cn/problems/count-substrings-starting-and-ending-with-given-character)|Medium|||||||||||
|[3070](https://leetcode.cn/problems/count-submatrices-with-top-left-element-and-sum-less-than-k)|Medium|||||||||||
|[3032](https://leetcode.cn/problems/count-numbers-with-unique-digits-ii)|Easy|||||||||||
|[3102](https://leetcode.cn/problems/minimize-manhattan-distances)|Hard|||||||||||
|[3082](https://leetcode.cn/problems/find-the-sum-of-the-power-of-all-subsequences)|Hard|||||||||||
|[3106](https://leetcode.cn/problems/lexicographically-smallest-string-after-operations-with-constraint)|Medium|||||||||||
|[3069](https://leetcode.cn/problems/distribute-elements-into-two-arrays-i)|Easy|||||||||||
|[3108](https://leetcode.cn/problems/minimum-cost-walk-in-weighted-graph)|Hard|||||||||||
|[3090](https://leetcode.cn/problems/maximum-length-substring-with-two-occurrences)|Easy|||||||||||
|[3072](https://leetcode.cn/problems/distribute-elements-into-two-arrays-ii)|Hard|||||||||||
|[3075](https://leetcode.cn/problems/maximize-happiness-of-selected-children)|Medium|||||||||||
|[3037](https://leetcode.cn/problems/find-pattern-in-infinite-stream-ii)|Hard|||||||||||
|[3083](https://leetcode.cn/problems/existence-of-a-substring-in-a-string-and-its-reverse)|Easy|||||||||||
|[3081](https://leetcode.cn/problems/replace-question-marks-in-string-to-minimize-its-value)|Medium|||||||||||
|[3096](https://leetcode.cn/problems/minimum-levels-to-gain-more-points)|Medium|||||||||||
|[3076](https://leetcode.cn/problems/shortest-uncommon-substring-in-an-array)|Medium|||||||||||
|[3063](https://leetcode.cn/problems/linked-list-frequency)|Easy|||||||||||
|[3085](https://leetcode.cn/problems/minimum-deletions-to-make-string-k-special)|Medium|||||||||||
|[3114](https://leetcode.cn/problems/latest-time-you-can-obtain-after-replacing-characters)|Easy|||||||||||
|[3134](https://leetcode.cn/problems/find-the-median-of-the-uniqueness-array)|Hard|||||||||||
|[3092](https://leetcode.cn/problems/most-frequent-ids)|Medium|||||||||||
|[3117](https://leetcode.cn/problems/minimum-sum-of-values-by-dividing-array)|Hard|||||||||||
|[3079](https://leetcode.cn/problems/find-the-sum-of-encrypted-integers)|Easy|||||||||||
|[3062](https://leetcode.cn/problems/winner-of-the-linked-list-game)|Easy|||||||||||
|[3064](https://leetcode.cn/problems/guess-the-number-using-bitwise-questions-i)|Medium|||||||||||
|[3099](https://leetcode.cn/problems/harshad-number)|Easy|||||||||||
|[3105](https://leetcode.cn/problems/longest-strictly-increasing-or-strictly-decreasing-subarray)|Easy|||||||||||
|[3115](https://leetcode.cn/problems/maximum-prime-difference)|Medium|||||||||||
|[3101](https://leetcode.cn/problems/count-alternating-subarrays)|Medium|||||||||||
|[3116](https://leetcode.cn/problems/kth-smallest-amount-with-single-denomination-combination)|Hard|||||||||||
|[3093](https://leetcode.cn/problems/longest-common-suffix-queries)|Hard|||||||||||
|[3073](https://leetcode.cn/problems/maximum-increasing-triplet-value)|Medium|||||||||||
|[3110](https://leetcode.cn/problems/score-of-a-string)|Easy|||||||||||
|[3097](https://leetcode.cn/problems/shortest-subarray-with-or-at-least-k-ii)|Medium|||||||||||
|[3095](https://leetcode.cn/problems/shortest-subarray-with-or-at-least-k-i)|Easy|||||||||||
|[3113](https://leetcode.cn/problems/find-the-number-of-subarrays-where-boundary-elements-are-maximum)|Hard|||||||||||
|[3147](https://leetcode.cn/problems/taking-maximum-energy-from-the-mystic-dungeon)|Medium|||||||||||
|[3137](https://leetcode.cn/problems/minimum-number-of-operations-to-make-word-k-periodic)|Medium|||||||||||
|[3078](https://leetcode.cn/problems/match-alphanumerical-pattern-in-matrix-i)|Medium|||||||||||
|[3123](https://leetcode.cn/problems/find-edges-in-shortest-paths)|Hard|||||||||||
|[3107](https://leetcode.cn/problems/minimum-operations-to-make-median-of-array-equal-to-k)|Medium|||||||||||
|[3128](https://leetcode.cn/problems/right-triangles)|Medium|||||||||||
|[3112](https://leetcode.cn/problems/minimum-time-to-visit-disappearing-nodes)|Medium|||||||||||
|[3111](https://leetcode.cn/problems/minimum-rectangles-to-cover-points)|Medium|||||||||||
|[3148](https://leetcode.cn/problems/maximum-difference-score-in-a-grid)|Medium|||||||||||
|[3088](https://leetcode.cn/problems/make-string-anti-palindrome)|Hard|||||||||||
|[3133](https://leetcode.cn/problems/minimum-array-end)|Medium|||||||||||
|[3138](https://leetcode.cn/problems/minimum-length-of-anagram-concatenation)|Medium|||||||||||
|[3136](https://leetcode.cn/problems/valid-word)|Easy|||||||||||
|[3131](https://leetcode.cn/problems/find-the-integer-added-to-array-i)|Easy|||||||||||
|[3127](https://leetcode.cn/problems/make-a-square-with-the-same-color)|Easy|||||||||||
|[3132](https://leetcode.cn/problems/find-the-integer-added-to-array-ii)|Medium|||||||||||
|[3094](https://leetcode.cn/problems/guess-the-number-using-bitwise-questions-ii)|Medium|||||||||||
|[3139](https://leetcode.cn/problems/minimum-cost-to-equalize-array)|Hard|||||||||||
|[3144](https://leetcode.cn/problems/minimum-substring-partition-of-equal-character-frequency)|Medium|||||||||||
|[3122](https://leetcode.cn/problems/minimum-number-of-operations-to-satisfy-conditions)|Medium|||||||||||
|[3121](https://leetcode.cn/problems/count-the-number-of-special-characters-ii)|Medium|||||||||||
|[3129](https://leetcode.cn/problems/find-all-possible-stable-binary-arrays-i)|Medium|||||||||||
|[3130](https://leetcode.cn/problems/find-all-possible-stable-binary-arrays-ii)|Hard|||||||||||
|[3120](https://leetcode.cn/problems/count-the-number-of-special-characters-i)|Easy|||||||||||
|[3104](https://leetcode.cn/problems/find-longest-self-contained-substring)|Hard|||||||||||
|[3145](https://leetcode.cn/problems/find-products-of-elements-of-big-array)|Hard|||||||||||
|[3146](https://leetcode.cn/problems/permutation-difference-between-two-strings)|Easy|||||||||||
|[3175](https://leetcode.cn/problems/find-the-first-player-to-win-k-games-in-a-row)|Medium|||||||||||
|[3154](https://leetcode.cn/problems/find-number-of-ways-to-reach-the-k-th-stair)|Hard|||||||||||
|[3142](https://leetcode.cn/problems/check-if-grid-satisfies-conditions)|Easy|||||||||||
|[3153](https://leetcode.cn/problems/sum-of-digit-differences-of-all-pairs)|Medium|||||||||||
|[3109](https://leetcode.cn/problems/find-the-index-of-permutation)|Medium|||||||||||
|[3185](https://leetcode.cn/problems/count-pairs-that-form-a-complete-day-ii)|Medium|||||||||||
|[3143](https://leetcode.cn/problems/maximum-points-inside-the-square)|Medium|||||||||||
|[3159](https://leetcode.cn/problems/find-occurrences-of-an-element-in-an-array)|Medium|||||||||||
|[3184](https://leetcode.cn/problems/count-pairs-that-form-a-complete-day-i)|Easy|||||||||||
|[3179](https://leetcode.cn/problems/find-the-n-th-value-after-k-seconds)|Medium|||||||||||
|[3165](https://leetcode.cn/problems/maximum-sum-of-subsequence-with-non-adjacent-elements)|Hard|||||||||||
|[3119](https://leetcode.cn/problems/maximum-number-of-potholes-that-can-be-fixed)|Medium|||||||||||
|[3168](https://leetcode.cn/problems/minimum-number-of-chairs-in-a-waiting-room)|Easy|||||||||||
|[3152](https://leetcode.cn/problems/special-array-ii)|Medium|||||||||||
|[3158](https://leetcode.cn/problems/find-the-xor-of-numbers-which-appear-twice)|Easy|||||||||||
|[3151](https://leetcode.cn/problems/special-array-i)|Easy|||||||||||
|[3169](https://leetcode.cn/problems/count-days-without-meetings)|Medium|||||||||||
|[3149](https://leetcode.cn/problems/find-the-minimum-cost-array-permutation)|Hard|||||||||||
|[3125](https://leetcode.cn/problems/maximum-number-that-makes-result-of-bitwise-and-zero)|Medium|||||||||||
|[3160](https://leetcode.cn/problems/find-the-number-of-distinct-colors-among-the-balls)|Medium|||||||||||
|[3161](https://leetcode.cn/problems/block-placement-queries)|Hard|||||||||||
|[3171](https://leetcode.cn/problems/find-subarray-with-bitwise-or-closest-to-k)|Hard|||||||||||
|[3186](https://leetcode.cn/problems/maximum-total-damage-with-spell-casting)|Medium|||||||||||
|[3187](https://leetcode.cn/problems/peaks-in-array)|Hard|||||||||||
|[3203](https://leetcode.cn/problems/find-minimum-diameter-after-merging-two-trees)|Hard|||||||||||
|[3135](https://leetcode.cn/problems/equalize-strings-by-adding-or-removing-characters-at-ends)|Medium|||||||||||
|[3180](https://leetcode.cn/problems/maximum-total-reward-using-operations-i)|Medium|||||||||||
|[3181](https://leetcode.cn/problems/maximum-total-reward-using-operations-ii)|Hard|||||||||||
|[3164](https://leetcode.cn/problems/find-the-number-of-good-pairs-ii)|Medium|||||||||||
|[3170](https://leetcode.cn/problems/lexicographically-minimum-string-after-removing-stars)|Medium|||||||||||
|[3162](https://leetcode.cn/problems/find-the-number-of-good-pairs-i)|Easy|||||||||||
|[3174](https://leetcode.cn/problems/clear-digits)|Easy|||||||||||
|[3141](https://leetcode.cn/problems/maximum-hamming-distances)|Hard|||||||||||
|[3178](https://leetcode.cn/problems/find-the-child-who-has-the-ball-after-k-seconds)|Easy|||||||||||
|[3163](https://leetcode.cn/problems/string-compression-iii)|Medium|||||||||||
|[3177](https://leetcode.cn/problems/find-the-maximum-length-of-a-good-subsequence-ii)|Hard|||||||||||
|[3211](https://leetcode.cn/problems/generate-binary-strings-without-adjacent-zeros)|Medium|||||||||||
|[3229](https://leetcode.cn/problems/minimum-operations-to-make-array-equal-to-target)|Hard|||||||||||
|[3223](https://leetcode.cn/problems/minimum-length-of-string-after-operations)|Medium|||||||||||
|[3176](https://leetcode.cn/problems/find-the-maximum-length-of-a-good-subsequence-i)|Medium|||||||||||
|[3155](https://leetcode.cn/problems/maximum-number-of-upgradable-servers)|Medium|||||||||||
|[3197](https://leetcode.cn/problems/find-the-minimum-area-to-cover-all-ones-ii)|Hard|||||||||||
|[3193](https://leetcode.cn/problems/count-the-number-of-inversions)|Hard|||||||||||
|[3195](https://leetcode.cn/problems/find-the-minimum-area-to-cover-all-ones-i)|Medium|||||||||||
|[3227](https://leetcode.cn/problems/vowels-game-in-a-string)|Medium|||||||||||
|[3206](https://leetcode.cn/problems/alternating-groups-i)|Easy|||||||||||
|[3196](https://leetcode.cn/problems/maximize-total-cost-of-alternating-subarrays)|Medium|||||||||||
|[3209](https://leetcode.cn/problems/number-of-subarrays-with-and-value-of-k)|Hard|||||||||||
|[3157](https://leetcode.cn/problems/find-the-level-of-tree-with-minimum-sum)|Medium|||||||||||
|[3210](https://leetcode.cn/problems/find-the-encrypted-string)|Easy|||||||||||
|[3200](https://leetcode.cn/problems/maximum-height-of-a-triangle)|Easy|||||||||||
|[3225](https://leetcode.cn/problems/maximum-score-from-grid-operations)|Hard|||||||||||
|[3194](https://leetcode.cn/problems/minimum-average-of-smallest-and-largest-elements)|Easy|||||||||||
|[3173](https://leetcode.cn/problems/bitwise-or-of-adjacent-elements)|Easy|||||||||||
|[3167](https://leetcode.cn/problems/better-compression-of-string)|Medium|||||||||||
|[3191](https://leetcode.cn/problems/minimum-operations-to-make-binary-array-elements-equal-to-one-i)|Medium|||||||||||
|[3190](https://leetcode.cn/problems/find-minimum-operations-to-make-all-elements-divisible-by-three)|Easy|||||||||||
|[3192](https://leetcode.cn/problems/minimum-operations-to-make-binary-array-elements-equal-to-one-ii)|Medium|||||||||||
|[3235](https://leetcode.cn/problems/check-if-the-rectangle-corner-is-reachable)|Hard|||||||||||
|[3234](https://leetcode.cn/problems/count-the-number-of-substrings-with-dominant-ones)|Medium|||||||||||
|[3253](https://leetcode.cn/problems/construct-string-with-minimum-cost-easy)|Medium|||||||||||
|[3213](https://leetcode.cn/problems/construct-string-with-minimum-cost)|Hard|||||||||||
|[3208](https://leetcode.cn/problems/alternating-groups-ii)|Medium|||||||||||
|[3216](https://leetcode.cn/problems/lexicographically-smallest-string-after-a-swap)|Easy|||||||||||
|[3281](https://leetcode.cn/problems/maximize-score-of-numbers-in-ranges)|Medium|||||||||||
|[3249](https://leetcode.cn/problems/count-the-number-of-good-nodes)|Medium|||||||||||
|[3316](https://leetcode.cn/problems/find-maximum-removals-from-source-string)|Medium|||||||||||
|[3183](https://leetcode.cn/problems/the-number-of-ways-to-make-the-sum)|Medium|||||||||||
|[3201](https://leetcode.cn/problems/find-the-maximum-length-of-valid-subsequence-i)|Medium|||||||||||
|[3202](https://leetcode.cn/problems/find-the-maximum-length-of-valid-subsequence-ii)|Medium|||||||||||
|[3212](https://leetcode.cn/problems/count-submatrices-with-equal-frequency-of-x-and-y)|Medium|||||||||||
|[3228](https://leetcode.cn/problems/maximum-number-of-operations-to-move-ones-to-the-end)|Medium|||||||||||
|[3218](https://leetcode.cn/problems/minimum-cost-for-cutting-cake-i)|Medium|||||||||||
|[3275](https://leetcode.cn/problems/k-th-nearest-obstacle-queries)|Medium|||||||||||
|[3296](https://leetcode.cn/problems/minimum-number-of-seconds-to-make-mountain-height-zero)|Medium|||||||||||
|[3313](https://leetcode.cn/problems/find-the-last-marked-nodes-in-tree)|Hard|||||||||||
|[3224](https://leetcode.cn/problems/minimum-array-changes-to-make-differences-equal)|Medium|||||||||||
|[3329](https://leetcode.cn/problems/count-substrings-with-k-frequency-characters-ii)|Hard|||||||||||
|[3219](https://leetcode.cn/problems/minimum-cost-for-cutting-cake-ii)|Hard|||||||||||
|[3217](https://leetcode.cn/problems/delete-nodes-from-linked-list-present-in-array)|Medium|||||||||||
|[3325](https://leetcode.cn/problems/count-substrings-with-k-frequency-characters-i)|Medium|||||||||||
|[3189](https://leetcode.cn/problems/minimum-moves-to-get-a-peaceful-board)|Medium|||||||||||
|[3199](https://leetcode.cn/problems/count-triplets-with-even-xor-set-bits-i)|Easy|||||||||||
|[3233](https://leetcode.cn/problems/find-the-count-of-numbers-which-are-not-special)|Medium|||||||||||
|[3226](https://leetcode.cn/problems/number-of-bit-changes-to-make-two-integers-equal)|Easy|||||||||||
|[3319](https://leetcode.cn/problems/k-th-largest-perfect-subtree-size-in-binary-tree)|Medium|||||||||||
|[3301](https://leetcode.cn/problems/maximize-the-total-height-of-unique-towers)|Medium|||||||||||
|[3222](https://leetcode.cn/problems/find-the-winning-player-in-coin-game)|Easy|||||||||||
|[3205](https://leetcode.cn/problems/maximum-array-hopping-score-i)|Medium|||||||||||
|[3244](https://leetcode.cn/problems/shortest-distance-after-road-addition-queries-ii)|Hard|||||||||||
|[3232](https://leetcode.cn/problems/find-if-digit-game-can-be-won)|Easy|||||||||||
|[3242](https://leetcode.cn/problems/design-neighbor-sum-service)|Easy|||||||||||
|[3243](https://leetcode.cn/problems/shortest-distance-after-road-addition-queries-i)|Medium|||||||||||
|[3290](https://leetcode.cn/problems/maximum-multiplication-score)|Medium|||||||||||
|[3238](https://leetcode.cn/problems/find-the-number-of-winning-players)|Easy|||||||||||
|[3215](https://leetcode.cn/problems/count-triplets-with-even-xor-set-bits-ii)|Medium|||||||||||
|[3254](https://leetcode.cn/problems/find-the-power-of-k-size-subarrays-i)|Medium|||||||||||
|[3255](https://leetcode.cn/problems/find-the-power-of-k-size-subarrays-ii)|Medium|||||||||||
|[3240](https://leetcode.cn/problems/minimum-number-of-flips-to-make-binary-grid-palindromic-ii)|Medium|||||||||||
|[3259](https://leetcode.cn/problems/maximum-energy-boost-from-two-drinks)|Medium|||||||||||
|[3239](https://leetcode.cn/problems/minimum-number-of-flips-to-make-binary-grid-palindromic-i)|Medium|||||||||||
|[3245](https://leetcode.cn/problems/alternating-groups-iii)|Hard|||||||||||
|[3282](https://leetcode.cn/problems/reach-end-of-array-with-max-score)|Medium|||||||||||
|[3221](https://leetcode.cn/problems/maximum-array-hopping-score-ii)|Medium|||||||||||
|[3273](https://leetcode.cn/problems/minimum-amount-of-damage-dealt-to-bob)|Hard|||||||||||
|[3241](https://leetcode.cn/problems/time-taken-to-mark-all-nodes)|Hard|||||||||||
|[3248](https://leetcode.cn/problems/snake-in-matrix)|Easy|||||||||||
|[3265](https://leetcode.cn/problems/count-almost-equal-pairs-i)|Medium|||||||||||
|[3250](https://leetcode.cn/problems/find-the-count-of-monotonic-pairs-i)|Hard|||||||||||
|[3251](https://leetcode.cn/problems/find-the-count-of-monotonic-pairs-ii)|Hard|||||||||||
|[3247](https://leetcode.cn/problems/number-of-subsequences-with-odd-sum)|Medium|||||||||||
|[3237](https://leetcode.cn/problems/alt-and-tab-simulation)|Medium|||||||||||
|[3271](https://leetcode.cn/problems/hash-divided-string)|Medium|||||||||||
|[3295](https://leetcode.cn/problems/report-spam-message)|Medium|||||||||||
|[3257](https://leetcode.cn/problems/maximum-value-sum-by-placing-three-rooks-ii)|Hard|||||||||||
|[3258](https://leetcode.cn/problems/count-substrings-that-satisfy-k-constraint-i)|Easy|||||||||||
|[3267](https://leetcode.cn/problems/count-almost-equal-pairs-ii)|Hard|||||||||||
|[3231](https://leetcode.cn/problems/minimum-number-of-increasing-subsequence-to-be-removed)|Hard|||||||||||
|[3261](https://leetcode.cn/problems/count-substrings-that-satisfy-k-constraint-ii)|Hard|||||||||||
|[3272](https://leetcode.cn/problems/find-the-count-of-good-integers)|Hard|||||||||||
|[3256](https://leetcode.cn/problems/maximum-value-sum-by-placing-three-rooks-i)|Hard|||||||||||
|[3277](https://leetcode.cn/problems/maximum-xor-score-subarray-queries)|Hard|||||||||||
|[3260](https://leetcode.cn/problems/find-the-largest-palindrome-divisible-by-k)|Hard|||||||||||
|[3274](https://leetcode.cn/problems/check-if-two-chessboard-squares-have-the-same-color)|Easy|||||||||||
|[3264](https://leetcode.cn/problems/final-array-state-after-k-multiplication-operations-i)|Easy|||||||||||
|[3266](https://leetcode.cn/problems/final-array-state-after-k-multiplication-operations-ii)|Hard|||||||||||
|[3292](https://leetcode.cn/problems/minimum-number-of-valid-strings-to-form-target-ii)|Hard|||||||||||
|[3286](https://leetcode.cn/problems/find-a-safe-walk-through-a-grid)|Medium|||||||||||
|[3291](https://leetcode.cn/problems/minimum-number-of-valid-strings-to-form-target-i)|Medium|||||||||||
|[3283](https://leetcode.cn/problems/maximum-number-of-moves-to-kill-all-pawns)|Hard|||||||||||
|[3310](https://leetcode.cn/problems/remove-methods-from-project)|Medium|||||||||||
|[3276](https://leetcode.cn/problems/select-cells-in-grid-with-maximum-score)|Hard|||||||||||
|[3323](https://leetcode.cn/problems/minimize-connected-groups-by-inserting-interval)|Medium|||||||||||
|[3324](https://leetcode.cn/problems/find-the-sequence-of-strings-appeared-on-the-screen)|Medium|||||||||||
|[3280](https://leetcode.cn/problems/convert-date-to-binary)|Easy|||||||||||
|[3270](https://leetcode.cn/problems/find-the-key-of-the-numbers)|Easy|||||||||||
|[3306](https://leetcode.cn/problems/count-of-substrings-containing-every-vowel-and-k-consonants-ii)|Medium|||||||||||
|[3305](https://leetcode.cn/problems/count-of-substrings-containing-every-vowel-and-k-consonants-i)|Medium|||||||||||
|[3288](https://leetcode.cn/problems/length-of-the-longest-increasing-path)|Hard|||||||||||
|[3298](https://leetcode.cn/problems/count-substrings-that-can-be-rearranged-to-contain-a-string-ii)|Hard|||||||||||
|[3297](https://leetcode.cn/problems/count-substrings-that-can-be-rearranged-to-contain-a-string-i)|Medium|||||||||||
|[3287](https://leetcode.cn/problems/find-the-maximum-sequence-value-of-array)|Hard|||||||||||
|[3331](https://leetcode.cn/problems/find-subtree-sizes-after-changes)|Medium|||||||||||
|[3263](https://leetcode.cn/problems/convert-doubly-linked-list-to-array-i)|Easy|||||||||||
|[3311](https://leetcode.cn/problems/construct-2d-grid-matching-graph-layout)|Hard|||||||||||
|[3309](https://leetcode.cn/problems/maximum-possible-number-by-binary-concatenation)|Medium|||||||||||
|[3303](https://leetcode.cn/problems/find-the-occurrence-of-first-almost-equal-substring)|Hard|||||||||||
|[3289](https://leetcode.cn/problems/the-two-sneaky-numbers-of-digitville)|Easy|||||||||||
|[3285](https://leetcode.cn/problems/find-indices-of-stable-mountains)|Easy|||||||||||
|[3312](https://leetcode.cn/problems/sorted-gcd-pair-queries)|Hard|||||||||||
|[3302](https://leetcode.cn/problems/find-the-lexicographically-smallest-valid-sequence)|Medium|||||||||||
|[3269](https://leetcode.cn/problems/constructing-two-increasing-arrays)|Hard|||||||||||
|[3332](https://leetcode.cn/problems/maximum-points-tourist-can-earn)|Medium|||||||||||
|[3320](https://leetcode.cn/problems/count-the-number-of-winning-sequences)|Hard|||||||||||
|[3279](https://leetcode.cn/problems/maximum-total-area-occupied-by-pistons)|Hard|||||||||||
|[3321](https://leetcode.cn/problems/find-x-sum-of-all-k-long-subarrays-ii)|Hard|||||||||||
|[3334](https://leetcode.cn/problems/find-the-maximum-factor-score-of-array)|Medium|||||||||||
|[3304](https://leetcode.cn/problems/find-the-k-th-character-in-string-game-i)|Easy|||||||||||
|[3307](https://leetcode.cn/problems/find-the-k-th-character-in-string-game-ii)|Hard|||||||||||
|[3284](https://leetcode.cn/problems/sum-of-consecutive-subarrays)|Medium|||||||||||
|[3327](https://leetcode.cn/problems/check-if-dfs-strings-are-palindromes)|Hard|||||||||||
|[3317](https://leetcode.cn/problems/find-the-number-of-possible-ways-for-an-event)|Hard|||||||||||
|[3314](https://leetcode.cn/problems/construct-the-minimum-bitwise-array-i)|Easy|||||||||||
|[3300](https://leetcode.cn/problems/minimum-element-after-replacement-with-digit-sum)|Easy|||||||||||
|[3326](https://leetcode.cn/problems/minimum-division-operations-to-make-array-non-decreasing)|Medium|||||||||||
|[3336](https://leetcode.cn/problems/find-the-number-of-subsequences-with-equal-gcd)|Hard|||||||||||
|[3318](https://leetcode.cn/problems/find-x-sum-of-all-k-long-subarrays-i)|Easy|||||||||||
|[3315](https://leetcode.cn/problems/construct-the-minimum-bitwise-array-ii)|Medium|||||||||||
|[3339](https://leetcode.cn/problems/find-the-number-of-k-even-arrays)|Medium|||||||||||
|[3294](https://leetcode.cn/problems/convert-doubly-linked-list-to-array-ii)|Medium|||||||||||
|[3330](https://leetcode.cn/problems/find-the-original-typed-string-i)|Easy|||||||||||
|[3333](https://leetcode.cn/problems/find-the-original-typed-string-ii)|Hard|||||||||||
|[3299](https://leetcode.cn/problems/sum-of-consecutive-subsequences)|Hard|||||||||||
|[3335](https://leetcode.cn/problems/total-characters-in-string-after-transformations-i)|Medium|||||||||||
|[3337](https://leetcode.cn/problems/total-characters-in-string-after-transformations-ii)|Hard|||||||||||
|[LCP 02](https://leetcode.cn/problems/deep-dark-fraction)|Easy|||||||||||
|[LCP 04](https://leetcode.cn/problems/broken-board-dominoes)|Hard|||||||||||
|[LCP 05](https://leetcode.cn/problems/coin-bonus)|Hard|||||||||||
|[LCP 03](https://leetcode.cn/problems/programmable-robot)|Medium|||||||||||
|[LCP 01](https://leetcode.cn/problems/guess-numbers)|Easy|||||||||||
|[LCR 125](https://leetcode.cn/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof)|Easy|||||||||||
|[LCR 126](https://leetcode.cn/problems/fei-bo-na-qi-shu-lie-lcof)|Easy|||||||||||
|[LCR 120](https://leetcode.cn/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof)|Easy|||||||||||
|[LCR 121](https://leetcode.cn/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof)|Medium|||||||||||
|[LCR 127](https://leetcode.cn/problems/qing-wa-tiao-tai-jie-wen-ti-lcof)|Easy|||||||||||
|[LCR 128](https://leetcode.cn/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof)|Easy|||||||||||
|[LCR 129](https://leetcode.cn/problems/ju-zhen-zhong-de-lu-jing-lcof)|Medium|||||||||||
|[LCR 122](https://leetcode.cn/problems/ti-huan-kong-ge-lcof)|Easy|||||||||||
|[LCR 130](https://leetcode.cn/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof)|Medium|||||||||||
|[LCR 123](https://leetcode.cn/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof)|Easy|||||||||||
|[LCR 124](https://leetcode.cn/problems/zhong-jian-er-cha-shu-lcof)|Medium|||||||||||
|[LCR 131](https://leetcode.cn/problems/jian-sheng-zi-lcof)|Medium|||||||||||
|[LCR 132](https://leetcode.cn/problems/jian-sheng-zi-ii-lcof)|Medium|||||||||||
|[LCR 142](https://leetcode.cn/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof)|Easy|||||||||||
|[LCR 143](https://leetcode.cn/problems/shu-de-zi-jie-gou-lcof)|Medium|||||||||||
|[LCR 144](https://leetcode.cn/problems/er-cha-shu-de-jing-xiang-lcof)|Easy|||||||||||
|[LCR 145](https://leetcode.cn/problems/dui-cheng-de-er-cha-shu-lcof)|Easy|||||||||||
|[LCR 138](https://leetcode.cn/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof)|Medium|||||||||||
|[LCR 139](https://leetcode.cn/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof)|Easy|||||||||||
|[LCR 133](https://leetcode.cn/problems/er-jin-zhi-zhong-1de-ge-shu-lcof)|Easy|||||||||||
|[LCR 146](https://leetcode.cn/problems/shun-shi-zhen-da-yin-ju-zhen-lcof)|Easy|||||||||||
|[LCR 140](https://leetcode.cn/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof)|Easy|||||||||||
|[LCR 134](https://leetcode.cn/problems/shu-zhi-de-zheng-shu-ci-fang-lcof)|Medium|||||||||||
|[LCR 135](https://leetcode.cn/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof)|Easy|||||||||||
|[LCR 137](https://leetcode.cn/problems/zheng-ze-biao-da-shi-pi-pei-lcof)|Hard|||||||||||
|[LCR 141](https://leetcode.cn/problems/fan-zhuan-lian-biao-lcof)|Easy|||||||||||
|[LCR 136](https://leetcode.cn/problems/shan-chu-lian-biao-de-jie-dian-lcof)|Easy|||||||||||
|[LCR 154](https://leetcode.cn/problems/fu-za-lian-biao-de-fu-zhi-lcof)|Medium|||||||||||
|[LCR 159](https://leetcode.cn/problems/zui-xiao-de-kge-shu-lcof)|Easy|||||||||||
|[LCR 147](https://leetcode.cn/problems/bao-han-minhan-shu-de-zhan-lcof)|Easy|||||||||||
|[LCR 160](https://leetcode.cn/problems/shu-ju-liu-zhong-de-zhong-wei-shu-lcof)|Hard|||||||||||
|[LCR 161](https://leetcode.cn/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof)|Easy|||||||||||
|[LCR 155](https://leetcode.cn/problems/er-cha-sou-suo-shu-yu-shuang-xiang-lian-biao-lcof)|Medium|||||||||||
|[LCR 148](https://leetcode.cn/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof)|Medium|||||||||||
|[LCR 156](https://leetcode.cn/problems/xu-lie-hua-er-cha-shu-lcof)|Hard|||||||||||
|[LCR 157](https://leetcode.cn/problems/zi-fu-chuan-de-pai-lie-lcof)|Medium|||||||||||
|[LCR 162](https://leetcode.cn/problems/1nzheng-shu-zhong-1chu-xian-de-ci-shu-lcof)|Hard|||||||||||
|[LCR 158](https://leetcode.cn/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof)|Easy|||||||||||
|[LCR 149](https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof)|Medium|||||||||||
|[LCR 150](https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof)|Easy|||||||||||
|[LCR 163](https://leetcode.cn/problems/shu-zi-xu-lie-zhong-mou-yi-wei-de-shu-zi-lcof)|Medium|||||||||||
|[LCR 151](https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof)|Medium|||||||||||
|[LCR 152](https://leetcode.cn/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof)|Medium|||||||||||
|[LCR 169](https://leetcode.cn/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof)|Easy|||||||||||
|[LCR 153](https://leetcode.cn/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof)|Medium|||||||||||
|[LCR 170](https://leetcode.cn/problems/shu-zu-zhong-de-ni-xu-dui-lcof)|Hard|||||||||||
|[LCR 175](https://leetcode.cn/problems/er-cha-shu-de-shen-du-lcof)|Easy|||||||||||
|[LCR 177](https://leetcode.cn/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof)|Medium|||||||||||
|[LCR 178](https://leetcode.cn/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof)|Medium|||||||||||
|[LCR 179](https://leetcode.cn/problems/he-wei-sde-liang-ge-shu-zi-lcof)|Easy|||||||||||
|[LCR 164](https://leetcode.cn/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof)|Medium|||||||||||
|[LCR 180](https://leetcode.cn/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof)|Easy|||||||||||
|[LCR 165](https://leetcode.cn/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof)|Medium|||||||||||
|[LCR 171](https://leetcode.cn/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof)|Easy|||||||||||
|[LCR 166](https://leetcode.cn/problems/li-wu-de-zui-da-jie-zhi-lcof)|Medium|||||||||||
|[LCR 181](https://leetcode.cn/problems/fan-zhuan-dan-ci-shun-xu-lcof)|Easy|||||||||||
|[LCR 172](https://leetcode.cn/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof)|Easy|||||||||||
|[LCR 182](https://leetcode.cn/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof)|Easy|||||||||||
|[LCR 173](https://leetcode.cn/problems/que-shi-de-shu-zi-lcof)|Easy|||||||||||
|[LCR 167](https://leetcode.cn/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof)|Medium|||||||||||
|[LCR 174](https://leetcode.cn/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof)|Easy|||||||||||
|[LCR 168](https://leetcode.cn/problems/chou-shu-lcof)|Medium|||||||||||
|[LCR 190](https://leetcode.cn/problems/bu-yong-jia-jian-cheng-chu-zuo-jia-fa-lcof)|Easy|||||||||||
|[LCR 183](https://leetcode.cn/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof)|Hard|||||||||||
|[LCR 184](https://leetcode.cn/problems/dui-lie-de-zui-da-zhi-lcof)|Medium|||||||||||
|[LCR 191](https://leetcode.cn/problems/gou-jian-cheng-ji-shu-zu-lcof)|Medium|||||||||||
|[LCR 185](https://leetcode.cn/problems/nge-tou-zi-de-dian-shu-lcof)|Medium|||||||||||
|[LCR 192](https://leetcode.cn/problems/ba-zi-fu-chuan-zhuan-huan-cheng-zheng-shu-lcof)|Medium|||||||||||
|[LCR 186](https://leetcode.cn/problems/bu-ke-pai-zhong-de-shun-zi-lcof)|Easy|||||||||||
|[LCR 176](https://leetcode.cn/problems/ping-heng-er-cha-shu-lcof)|Easy|||||||||||
|[LCR 187](https://leetcode.cn/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof)|Easy|||||||||||
|[LCR 188](https://leetcode.cn/problems/gu-piao-de-zui-da-li-run-lcof)|Medium|||||||||||
|[LCR 189](https://leetcode.cn/problems/qiu-12n-lcof)|Medium|||||||||||
|[LCR 193](https://leetcode.cn/problems/er-cha-sou-suo-shu-de-zui-jin-gong-gong-zu-xian-lcof)|Easy|||||||||||
|[LCR 194](https://leetcode.cn/problems/er-cha-shu-de-zui-jin-gong-gong-zu-xian-lcof)|Easy|||||||||||
|[LCP 06](https://leetcode.cn/problems/na-ying-bi)|Easy|||||||||||
|[LCP 08](https://leetcode.cn/problems/ju-qing-hong-fa-shi-jian)|Medium|||||||||||
|[LCP 23](https://leetcode.cn/problems/er94lq)|Medium|||||||||||
|[LCP 14](https://leetcode.cn/problems/qie-fen-shu-zu)|Hard|||||||||||
|[LCP 10](https://leetcode.cn/problems/er-cha-shu-ren-wu-diao-du)|Hard|||||||||||
|[LCP 07](https://leetcode.cn/problems/chuan-di-xin-xi)|Easy|||||||||||
|[LCP 16](https://leetcode.cn/problems/you-le-yuan-de-you-lan-ji-hua)|Hard|||||||||||
|[LCP 11](https://leetcode.cn/problems/qi-wang-ge-shu-tong-ji)|Easy|||||||||||
|[LCP 12](https://leetcode.cn/problems/xiao-zhang-shua-ti-ji-hua)|Medium|||||||||||
|[LCP 15](https://leetcode.cn/problems/you-le-yuan-de-mi-gong)|Hard|||||||||||
|[LCP 34](https://leetcode.cn/problems/er-cha-shu-ran-se-UGC)|Medium|||||||||||
|[LCP 09](https://leetcode.cn/problems/zui-xiao-tiao-yue-ci-shu)|Hard|||||||||||
|[LCP 37](https://leetcode.cn/problems/zui-xiao-ju-xing-mian-ji)|Hard|||||||||||
|[LCP 13](https://leetcode.cn/problems/xun-bao)|Hard|||||||||||
|[LCP 19](https://leetcode.cn/problems/UlBDOe)|Medium|||||||||||
|[LCP 24](https://leetcode.cn/problems/5TxKeK)|Hard|||||||||||
|[LCP 26](https://leetcode.cn/problems/hSRGyL)|Hard|||||||||||
|[LCP 20](https://leetcode.cn/problems/meChtZ)|Hard|||||||||||
|[LCP 21](https://leetcode.cn/problems/Za25hA)|Hard|||||||||||
|[LCP 25](https://leetcode.cn/problems/Uh984O)|Hard|||||||||||
|[LCP 17](https://leetcode.cn/problems/nGK0Fy)|Easy|||||||||||
|[LCP 22](https://leetcode.cn/problems/ccw6C7)|Easy|||||||||||
|[LCP 18](https://leetcode.cn/problems/2vYnGI)|Easy|||||||||||
|[LCP 27](https://leetcode.cn/problems/IQvJ9i)|Hard|||||||||||
|[LCP 28](https://leetcode.cn/problems/4xy4Wx)|Easy|||||||||||
|[LCP 29](https://leetcode.cn/problems/SNJvJP)|Medium|||||||||||
|[LCP 31](https://leetcode.cn/problems/Db3wC1)|Hard|||||||||||
|[LCP 32](https://leetcode.cn/problems/t3fKg1)|Hard|||||||||||
|[LCP 35](https://leetcode.cn/problems/DFPeFJ)|Hard|||||||||||
|[LCP 33](https://leetcode.cn/problems/o8SXZn)|Easy|||||||||||
|[LCP 36](https://leetcode.cn/problems/Up5XYM)|Hard|||||||||||
|[LCP 38](https://leetcode.cn/problems/7rLGCR)|Hard|||||||||||
|[LCP 30](https://leetcode.cn/problems/p0NxJO)|Medium|||||||||||
|[LCS 01](https://leetcode.cn/problems/Ju9Xwi)|Easy|||||||||||
|[LCS 02](https://leetcode.cn/problems/WqXACV)|Easy|||||||||||
|[LCS 03](https://leetcode.cn/problems/YesdPw)|Medium|||||||||||
|[LCR 001](https://leetcode.cn/problems/xoh6Oh)|Easy|||||||||||
|[LCR 081](https://leetcode.cn/problems/Ygoe9J)|Medium|||||||||||
|[LCR 003](https://leetcode.cn/problems/w3tCBm)|Easy|||||||||||
|[LCR 002](https://leetcode.cn/problems/JFETK5)|Easy|||||||||||
|[LCR 082](https://leetcode.cn/problems/4sjJUc)|Medium|||||||||||
|[LCR 004](https://leetcode.cn/problems/WGki4K)|Medium|||||||||||
|[LCR 083](https://leetcode.cn/problems/VvJkup)|Medium|||||||||||
|[LCR 084](https://leetcode.cn/problems/7p8L0Z)|Medium|||||||||||
|[LCR 005](https://leetcode.cn/problems/aseY1I)|Medium|||||||||||
|[LCR 006](https://leetcode.cn/problems/kLl5u1)|Easy|||||||||||
|[LCR 085](https://leetcode.cn/problems/IDBivT)|Medium|||||||||||
|[LCR 007](https://leetcode.cn/problems/1fGaJU)|Medium|||||||||||
|[LCR 086](https://leetcode.cn/problems/M99OJA)|Medium|||||||||||
|[LCR 087](https://leetcode.cn/problems/0on3uN)|Medium|||||||||||
|[LCR 008](https://leetcode.cn/problems/2VG8Kg)|Medium|||||||||||
|[LCR 088](https://leetcode.cn/problems/GzCJIP)|Easy|||||||||||
|[LCR 009](https://leetcode.cn/problems/ZVAVXX)|Medium|||||||||||
|[LCR 089](https://leetcode.cn/problems/Gu0c2T)|Medium|||||||||||
|[LCR 010](https://leetcode.cn/problems/QTMn0o)|Medium|||||||||||
|[LCR 011](https://leetcode.cn/problems/A1NYOS)|Medium|||||||||||
|[LCR 012](https://leetcode.cn/problems/tvdfij)|Easy|||||||||||
|[LCR 013](https://leetcode.cn/problems/O4NDxx)|Medium|||||||||||
|[LCR 014](https://leetcode.cn/problems/MPnaiL)|Medium|||||||||||
|[LCR 015](https://leetcode.cn/problems/VabMRr)|Medium|||||||||||
|[LCR 016](https://leetcode.cn/problems/wtcaE1)|Medium|||||||||||
|[LCR 017](https://leetcode.cn/problems/M1oyTv)|Hard|||||||||||
|[LCR 018](https://leetcode.cn/problems/XltzEq)|Easy|||||||||||
|[LCR 019](https://leetcode.cn/problems/RQku0D)|Easy|||||||||||
|[LCR 020](https://leetcode.cn/problems/a7VOhD)|Medium|||||||||||
|[LCR 021](https://leetcode.cn/problems/SLwz0R)|Medium|||||||||||
|[LCR 022](https://leetcode.cn/problems/c32eOV)|Medium|||||||||||
|[LCR 023](https://leetcode.cn/problems/3u1WK4)|Easy|||||||||||
|[LCR 024](https://leetcode.cn/problems/UHnkqh)|Easy|||||||||||
|[LCR 025](https://leetcode.cn/problems/lMSNwu)|Medium|||||||||||
|[LCR 026](https://leetcode.cn/problems/LGjMqU)|Medium|||||||||||
|[LCR 027](https://leetcode.cn/problems/aMhZSa)|Easy|||||||||||
|[LCR 028](https://leetcode.cn/problems/Qv1Da2)|Medium|||||||||||
|[LCR 029](https://leetcode.cn/problems/4ueAj6)|Medium|||||||||||
|[LCR 090](https://leetcode.cn/problems/PzWKhm)|Medium|||||||||||
|[LCR 030](https://leetcode.cn/problems/FortPu)|Medium|||||||||||
|[LCR 091](https://leetcode.cn/problems/JEj789)|Medium|||||||||||
|[LCR 092](https://leetcode.cn/problems/cyJERH)|Medium|||||||||||
|[LCR 031](https://leetcode.cn/problems/OrIXps)|Medium|||||||||||
|[LCR 093](https://leetcode.cn/problems/Q91FMA)|Medium|||||||||||
|[LCR 094](https://leetcode.cn/problems/omKAoA)|Hard|||||||||||
|[LCR 032](https://leetcode.cn/problems/dKk3P7)|Easy|||||||||||
|[LCR 095](https://leetcode.cn/problems/qJnOS7)|Medium|||||||||||
|[LCR 033](https://leetcode.cn/problems/sfvd7V)|Medium|||||||||||
|[LCR 034](https://leetcode.cn/problems/lwyVBB)|Easy|||||||||||
|[LCR 096](https://leetcode.cn/problems/IY6buf)|Medium|||||||||||
|[LCR 035](https://leetcode.cn/problems/569nqc)|Medium|||||||||||
|[LCR 036](https://leetcode.cn/problems/8Zf90G)|Medium|||||||||||
|[LCR 097](https://leetcode.cn/problems/21dk04)|Hard|||||||||||
|[LCR 037](https://leetcode.cn/problems/XagZNi)|Medium|||||||||||
|[LCR 038](https://leetcode.cn/problems/iIQa4I)|Medium|||||||||||
|[LCR 039](https://leetcode.cn/problems/0ynMMM)|Hard|||||||||||
|[LCR 040](https://leetcode.cn/problems/PLYXKQ)|Hard|||||||||||
|[LCR 099](https://leetcode.cn/problems/0i0mDW)|Medium|||||||||||
|[LCR 100](https://leetcode.cn/problems/IlPe0q)|Medium|||||||||||
|[LCR 101](https://leetcode.cn/problems/NUPfPr)|Easy|||||||||||
|[LCR 102](https://leetcode.cn/problems/YaVDxD)|Medium|||||||||||
|[LCR 103](https://leetcode.cn/problems/gaM7Ch)|Medium|||||||||||
|[LCR 104](https://leetcode.cn/problems/D0F0SV)|Medium|||||||||||
|[LCR 105](https://leetcode.cn/problems/ZL6zAn)|Medium|||||||||||
|[LCR 041](https://leetcode.cn/problems/qIsx9U)|Easy|||||||||||
|[LCR 042](https://leetcode.cn/problems/H8086Q)|Easy|||||||||||
|[LCR 106](https://leetcode.cn/problems/vEAB3K)|Medium|||||||||||
|[LCR 043](https://leetcode.cn/problems/NaqhDT)|Medium|||||||||||
|[LCR 107](https://leetcode.cn/problems/2bCMpM)|Medium|||||||||||
|[LCR 044](https://leetcode.cn/problems/hPov7L)|Medium|||||||||||
|[LCR 045](https://leetcode.cn/problems/LwUNpT)|Medium|||||||||||
|[LCR 046](https://leetcode.cn/problems/WNC0Lk)|Medium|||||||||||
|[LCR 108](https://leetcode.cn/problems/om3reC)|Hard|||||||||||
|[LCR 047](https://leetcode.cn/problems/pOCWxh)|Medium|||||||||||
|[LCR 109](https://leetcode.cn/problems/zlDJc7)|Medium|||||||||||
|[LCR 110](https://leetcode.cn/problems/bP4bmD)|Medium|||||||||||
|[LCR 048](https://leetcode.cn/problems/h54YBf)|Hard|||||||||||
|[LCR 111](https://leetcode.cn/problems/vlzXQL)|Medium|||||||||||
|[LCR 049](https://leetcode.cn/problems/3Etpl5)|Medium|||||||||||
|[LCR 050](https://leetcode.cn/problems/6eUYwP)|Medium|||||||||||
|[LCR 112](https://leetcode.cn/problems/fpTFWP)|Hard|||||||||||
|[LCR 051](https://leetcode.cn/problems/jC7MId)|Hard|||||||||||
|[LCR 113](https://leetcode.cn/problems/QA2IGt)|Medium|||||||||||
|[LCR 052](https://leetcode.cn/problems/NYBBNL)|Easy|||||||||||
|[LCR 114](https://leetcode.cn/problems/Jf1JuT)|Hard|||||||||||
|[LCR 053](https://leetcode.cn/problems/P5rCT8)|Medium|||||||||||
|[LCR 115](https://leetcode.cn/problems/ur2n8P)|Medium|||||||||||
|[LCR 054](https://leetcode.cn/problems/w6cpku)|Medium|||||||||||
|[LCR 055](https://leetcode.cn/problems/kTOapQ)|Medium|||||||||||
|[LCR 117](https://leetcode.cn/problems/H6lPxb)|Hard|||||||||||
|[LCR 118](https://leetcode.cn/problems/7LpjUW)|Medium|||||||||||
|[LCR 056](https://leetcode.cn/problems/opLdQZ)|Easy|||||||||||
|[LCR 119](https://leetcode.cn/problems/WhsWhI)|Medium|||||||||||
|[LCR 057](https://leetcode.cn/problems/7WqeDu)|Medium|||||||||||
|[LCR 058](https://leetcode.cn/problems/fi9suh)|Medium|||||||||||
|[LCR 059](https://leetcode.cn/problems/jBjn9C)|Easy|||||||||||
|[LCR 060](https://leetcode.cn/problems/g5c51o)|Medium|||||||||||
|[LCR 062](https://leetcode.cn/problems/QC3q1f)|Medium|||||||||||
|[LCR 061](https://leetcode.cn/problems/qn8gGX)|Medium|||||||||||
|[LCR 063](https://leetcode.cn/problems/UhWRSj)|Medium|||||||||||
|[LCR 064](https://leetcode.cn/problems/US1pGT)|Medium|||||||||||
|[LCR 065](https://leetcode.cn/problems/iSwD2y)|Medium|||||||||||
|[LCR 066](https://leetcode.cn/problems/z1R5dt)|Medium|||||||||||
|[LCR 067](https://leetcode.cn/problems/ms70jA)|Medium|||||||||||
|[LCR 068](https://leetcode.cn/problems/N6YdxV)|Easy|||||||||||
|[LCR 069](https://leetcode.cn/problems/B1IidL)|Easy|||||||||||
|[LCR 070](https://leetcode.cn/problems/skFtm2)|Medium|||||||||||
|[LCR 071](https://leetcode.cn/problems/cuyjEf)|Medium|||||||||||
|[LCR 072](https://leetcode.cn/problems/jJ0w9p)|Easy|||||||||||
|[LCR 073](https://leetcode.cn/problems/nZZqjQ)|Medium|||||||||||
|[LCR 074](https://leetcode.cn/problems/SsGoHC)|Medium|||||||||||
|[LCR 075](https://leetcode.cn/problems/0H97ZC)|Easy|||||||||||
|[LCR 076](https://leetcode.cn/problems/xx4gT2)|Medium|||||||||||
|[LCR 077](https://leetcode.cn/problems/7WHec2)|Medium|||||||||||
|[LCR 078](https://leetcode.cn/problems/vvXgSW)|Hard|||||||||||
|[LCR 079](https://leetcode.cn/problems/TVdhkn)|Medium|||||||||||
|[LCR 080](https://leetcode.cn/problems/uUsW3B)|Medium|||||||||||
|[LCR 098](https://leetcode.cn/problems/2AoeFn)|Medium|||||||||||
|[LCR 116](https://leetcode.cn/problems/bLyHh0)|Medium|||||||||||
|[LCP 49](https://leetcode.cn/problems/K8GULz)|Hard|||||||||||
|[LCP 47](https://leetcode.cn/problems/oPs9Bm)|Hard|||||||||||
|[LCP 44](https://leetcode.cn/problems/sZ59z6)|Easy|||||||||||
|[LCP 43](https://leetcode.cn/problems/Y1VbOX)|Hard|||||||||||
|[LCP 40](https://leetcode.cn/problems/uOAnQW)|Easy|||||||||||
|[LCP 39](https://leetcode.cn/problems/0jQkd0)|Easy|||||||||||
|[LCP 41](https://leetcode.cn/problems/fHi6rV)|Medium|||||||||||
|[LCP 42](https://leetcode.cn/problems/vFjcfV)|Hard|||||||||||
|[LCP 46](https://leetcode.cn/problems/05ZEDJ)|Medium|||||||||||
|[LCP 48](https://leetcode.cn/problems/fsa7oZ)|Hard|||||||||||
|[LCP 45](https://leetcode.cn/problems/kplEvH)|Medium|||||||||||
|[LCP 50](https://leetcode.cn/problems/WHnhjV)|Easy|||||||||||
|[LCP 51](https://leetcode.cn/problems/UEcfPD)|Easy|||||||||||
|[LCP 52](https://leetcode.cn/problems/QO5KpG)|Medium|||||||||||
|[LCP 53](https://leetcode.cn/problems/EJvmW4)|Hard|||||||||||
|[LCP 54](https://leetcode.cn/problems/s5kipK)|Hard|||||||||||
|[LCP 55](https://leetcode.cn/problems/PTXy4P)|Easy|||||||||||
|[LCP 56](https://leetcode.cn/problems/6UEx57)|Medium|||||||||||
|[LCP 57](https://leetcode.cn/problems/ZbAuEH)|Hard|||||||||||
|[LCP 59](https://leetcode.cn/problems/NfY1m5)|Hard|||||||||||
|[LCP 60](https://leetcode.cn/problems/WInSav)|Hard|||||||||||
|[LCP 58](https://leetcode.cn/problems/De4qBB)|Hard|||||||||||
|[LCP 61](https://leetcode.cn/problems/6CE719)|Easy|||||||||||
|[LCP 62](https://leetcode.cn/problems/D9PW8w)|Medium|||||||||||
|[LCP 65](https://leetcode.cn/problems/3aqs1c)|Hard|||||||||||
|[LCP 64](https://leetcode.cn/problems/U7WvvU)|Medium|||||||||||
|[LCP 63](https://leetcode.cn/problems/EXvqDp)|Medium|||||||||||
|[LCP 66](https://leetcode.cn/problems/600YaG)|Easy|||||||||||
|[LCP 67](https://leetcode.cn/problems/KnLfVT)|Medium|||||||||||
|[LCP 68](https://leetcode.cn/problems/1GxJYY)|Medium|||||||||||
|[LCP 69](https://leetcode.cn/problems/rMeRt2)|Hard|||||||||||
|[LCP 70](https://leetcode.cn/problems/XxZZjK)|Hard|||||||||||
|[LCP 71](https://leetcode.cn/problems/kskhHQ)|Hard|||||||||||
|[LCP 72](https://leetcode.cn/problems/hqCnmP)|Easy|||||||||||
|[LCP 73](https://leetcode.cn/problems/0Zeoeg)|Medium|||||||||||
|[LCP 74](https://leetcode.cn/problems/xepqZ5)|Medium|||||||||||
|[LCP 75](https://leetcode.cn/problems/rdmXM7)|Hard|||||||||||
|[LCP 76](https://leetcode.cn/problems/1ybDKD)|Hard|||||||||||
|[LCP 77](https://leetcode.cn/problems/W2ZX4X)|Easy|||||||||||
|[LCP 78](https://leetcode.cn/problems/Nsibyl)|Medium|||||||||||
|[LCP 79](https://leetcode.cn/problems/kjpLFZ)|Medium|||||||||||
|[LCP 80](https://leetcode.cn/problems/qoQAMX)|Hard|||||||||||
|[LCP 81](https://leetcode.cn/problems/ryfUiz)|Hard|||||||||||
|[LCP 82](https://leetcode.cn/problems/cnHoX6)|Hard|||||||||||

</details>

<details>

  <summary>Database</summary>

|Index|Difficulty|C|C++|Go|Java|TypeScript|PHP|Python|Rust|SQL|Bash|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
|[0175](https://leetcode.cn/problems/combine-two-tables)|Easy|||||||||:white_check_mark:||
|[0176](https://leetcode.cn/problems/second-highest-salary)|Medium|||||||||||
|[0177](https://leetcode.cn/problems/nth-highest-salary)|Medium|||||||||||
|[0178](https://leetcode.cn/problems/rank-scores)|Medium|||||||||||
|[0180](https://leetcode.cn/problems/consecutive-numbers)|Medium|||||||||||
|[0181](https://leetcode.cn/problems/employees-earning-more-than-their-managers)|Easy|||||||||||
|[0182](https://leetcode.cn/problems/duplicate-emails)|Easy|||||||||||
|[0183](https://leetcode.cn/problems/customers-who-never-order)|Easy|||||||||||
|[0184](https://leetcode.cn/problems/department-highest-salary)|Medium|||||||||||
|[0185](https://leetcode.cn/problems/department-top-three-salaries)|Hard|||||||||||
|[0196](https://leetcode.cn/problems/delete-duplicate-emails)|Easy|||||||||||
|[0197](https://leetcode.cn/problems/rising-temperature)|Easy|||||||||||
|[0262](https://leetcode.cn/problems/trips-and-users)|Hard|||||||||||
|[0569](https://leetcode.cn/problems/median-employee-salary)|Hard|||||||||||
|[0570](https://leetcode.cn/problems/managers-with-at-least-5-direct-reports)|Medium|||||||||||
|[0571](https://leetcode.cn/problems/find-median-given-frequency-of-numbers)|Hard|||||||||||
|[0574](https://leetcode.cn/problems/winning-candidate)|Medium|||||||||||
|[0577](https://leetcode.cn/problems/employee-bonus)|Easy|||||||||||
|[0578](https://leetcode.cn/problems/get-highest-answer-rate-question)|Medium|||||||||||
|[0579](https://leetcode.cn/problems/find-cumulative-salary-of-an-employee)|Hard|||||||||||
|[0580](https://leetcode.cn/problems/count-student-number-in-departments)|Medium|||||||||||
|[0584](https://leetcode.cn/problems/find-customer-referee)|Easy|||||||||||
|[0585](https://leetcode.cn/problems/investments-in-2016)|Medium|||||||||||
|[0586](https://leetcode.cn/problems/customer-placing-the-largest-number-of-orders)|Easy|||||||||||
|[0595](https://leetcode.cn/problems/big-countries)|Easy|||||||||||
|[0596](https://leetcode.cn/problems/classes-more-than-5-students)|Easy|||||||||||
|[0597](https://leetcode.cn/problems/friend-requests-i-overall-acceptance-rate)|Easy|||||||||||
|[0601](https://leetcode.cn/problems/human-traffic-of-stadium)|Hard|||||||||||
|[0602](https://leetcode.cn/problems/friend-requests-ii-who-has-the-most-friends)|Medium|||||||||||
|[0603](https://leetcode.cn/problems/consecutive-available-seats)|Easy|||||||||||
|[0607](https://leetcode.cn/problems/sales-person)|Easy|||||||||||
|[0608](https://leetcode.cn/problems/tree-node)|Medium|||||||||||
|[0610](https://leetcode.cn/problems/triangle-judgement)|Easy|||||||||||
|[0612](https://leetcode.cn/problems/shortest-distance-in-a-plane)|Medium|||||||||||
|[0613](https://leetcode.cn/problems/shortest-distance-in-a-line)|Easy|||||||||||
|[0614](https://leetcode.cn/problems/second-degree-follower)|Medium|||||||||||
|[0615](https://leetcode.cn/problems/average-salary-departments-vs-company)|Hard|||||||||||
|[0618](https://leetcode.cn/problems/students-report-by-geography)|Hard|||||||||||
|[0619](https://leetcode.cn/problems/biggest-single-number)|Easy|||||||||||
|[0620](https://leetcode.cn/problems/not-boring-movies)|Easy|||||||||||
|[0626](https://leetcode.cn/problems/exchange-seats)|Medium|||||||||||
|[0627](https://leetcode.cn/problems/swap-salary)|Easy|||||||||||
|[1045](https://leetcode.cn/problems/customers-who-bought-all-products)|Medium|||||||||||
|[1050](https://leetcode.cn/problems/actors-and-directors-who-cooperated-at-least-three-times)|Easy|||||||||||
|[1068](https://leetcode.cn/problems/product-sales-analysis-i)|Easy|||||||||||
|[1069](https://leetcode.cn/problems/product-sales-analysis-ii)|Easy|||||||||||
|[1070](https://leetcode.cn/problems/product-sales-analysis-iii)|Medium|||||||||||
|[1075](https://leetcode.cn/problems/project-employees-i)|Easy|||||||||||
|[1076](https://leetcode.cn/problems/project-employees-ii)|Easy|||||||||||
|[1077](https://leetcode.cn/problems/project-employees-iii)|Medium|||||||||||
|[1082](https://leetcode.cn/problems/sales-analysis-i)|Easy|||||||||||
|[1083](https://leetcode.cn/problems/sales-analysis-ii)|Easy|||||||||||
|[1084](https://leetcode.cn/problems/sales-analysis-iii)|Easy|||||||||||
|[0511](https://leetcode.cn/problems/game-play-analysis-i)|Easy|||||||||||
|[0512](https://leetcode.cn/problems/game-play-analysis-ii)|Easy|||||||||||
|[0534](https://leetcode.cn/problems/game-play-analysis-iii)|Medium|||||||||||
|[0550](https://leetcode.cn/problems/game-play-analysis-iv)|Medium|||||||||||
|[1097](https://leetcode.cn/problems/game-play-analysis-v)|Hard|||||||||||
|[1098](https://leetcode.cn/problems/unpopular-books)|Medium|||||||||||
|[1107](https://leetcode.cn/problems/new-users-daily-count)|Medium|||||||||||
|[1112](https://leetcode.cn/problems/highest-grade-for-each-student)|Medium|||||||||||
|[1113](https://leetcode.cn/problems/reported-posts)|Easy|||||||||||
|[1126](https://leetcode.cn/problems/active-businesses)|Medium|||||||||||
|[1127](https://leetcode.cn/problems/user-purchase-platform)|Hard|||||||||||
|[1132](https://leetcode.cn/problems/reported-posts-ii)|Medium|||||||||||
|[1141](https://leetcode.cn/problems/user-activity-for-the-past-30-days-i)|Easy|||||||||||
|[1142](https://leetcode.cn/problems/user-activity-for-the-past-30-days-ii)|Easy|||||||||||
|[1148](https://leetcode.cn/problems/article-views-i)|Easy|||||||||||
|[1149](https://leetcode.cn/problems/article-views-ii)|Medium|||||||||||
|[1158](https://leetcode.cn/problems/market-analysis-i)|Medium|||||||||||
|[1159](https://leetcode.cn/problems/market-analysis-ii)|Hard|||||||||||
|[1164](https://leetcode.cn/problems/product-price-at-a-given-date)|Medium|||||||||||
|[1173](https://leetcode.cn/problems/immediate-food-delivery-i)|Easy|||||||||||
|[1174](https://leetcode.cn/problems/immediate-food-delivery-ii)|Medium|||||||||||
|[1179](https://leetcode.cn/problems/reformat-department-table)|Easy|||||||||||
|[1193](https://leetcode.cn/problems/monthly-transactions-i)|Medium|||||||||||
|[1194](https://leetcode.cn/problems/tournament-winners)|Hard|||||||||||
|[1204](https://leetcode.cn/problems/last-person-to-fit-in-the-bus)|Medium|||||||||||
|[1205](https://leetcode.cn/problems/monthly-transactions-ii)|Medium|||||||||||
|[1211](https://leetcode.cn/problems/queries-quality-and-percentage)|Easy|||||||||||
|[1212](https://leetcode.cn/problems/team-scores-in-football-tournament)|Medium|||||||||||
|[1225](https://leetcode.cn/problems/report-contiguous-dates)|Hard|||||||||||
|[1241](https://leetcode.cn/problems/number-of-comments-per-post)|Easy|||||||||||
|[1251](https://leetcode.cn/problems/average-selling-price)|Easy|||||||||||
|[1264](https://leetcode.cn/problems/page-recommendations)|Medium|||||||||||
|[1270](https://leetcode.cn/problems/all-people-report-to-the-given-manager)|Medium|||||||||||
|[1280](https://leetcode.cn/problems/students-and-examinations)|Easy|||||||||||
|[1285](https://leetcode.cn/problems/find-the-start-and-end-number-of-continuous-ranges)|Medium|||||||||||
|[1294](https://leetcode.cn/problems/weather-type-in-each-country)|Easy|||||||||||
|[1303](https://leetcode.cn/problems/find-the-team-size)|Easy|||||||||||
|[1308](https://leetcode.cn/problems/running-total-for-different-genders)|Medium|||||||||||
|[1321](https://leetcode.cn/problems/restaurant-growth)|Medium|||||||||||
|[1322](https://leetcode.cn/problems/ads-performance)|Easy|||||||||||
|[1327](https://leetcode.cn/problems/list-the-products-ordered-in-a-period)|Easy|||||||||||
|[1336](https://leetcode.cn/problems/number-of-transactions-per-visit)|Hard|||||||||||
|[1341](https://leetcode.cn/problems/movie-rating)|Medium|||||||||||
|[1350](https://leetcode.cn/problems/students-with-invalid-departments)|Easy|||||||||||
|[1355](https://leetcode.cn/problems/activity-participants)|Medium|||||||||||
|[1364](https://leetcode.cn/problems/number-of-trusted-contacts-of-a-customer)|Medium|||||||||||
|[1369](https://leetcode.cn/problems/get-the-second-most-recent-activity)|Hard|||||||||||
|[1378](https://leetcode.cn/problems/replace-employee-id-with-the-unique-identifier)|Easy|||||||||||
|[1384](https://leetcode.cn/problems/total-sales-amount-by-year)|Hard|||||||||||
|[1393](https://leetcode.cn/problems/capital-gainloss)|Medium|||||||||||
|[1398](https://leetcode.cn/problems/customers-who-bought-products-a-and-b-but-not-c)|Medium|||||||||||
|[1407](https://leetcode.cn/problems/top-travellers)|Easy|||||||||||
|[1412](https://leetcode.cn/problems/find-the-quiet-students-in-all-exams)|Hard|||||||||||
|[1421](https://leetcode.cn/problems/npv-queries)|Easy|||||||||||
|[1435](https://leetcode.cn/problems/create-a-session-bar-chart)|Easy|||||||||||
|[1440](https://leetcode.cn/problems/evaluate-boolean-expression)|Medium|||||||||||
|[1445](https://leetcode.cn/problems/apples-oranges)|Medium|||||||||||
|[1454](https://leetcode.cn/problems/active-users)|Medium|||||||||||
|[1459](https://leetcode.cn/problems/rectangles-area)|Medium|||||||||||
|[1468](https://leetcode.cn/problems/calculate-salaries)|Medium|||||||||||
|[1479](https://leetcode.cn/problems/sales-by-day-of-the-week)|Hard|||||||||||
|[1484](https://leetcode.cn/problems/group-sold-products-by-the-date)|Easy|||||||||||
|[1495](https://leetcode.cn/problems/friendly-movies-streamed-last-month)|Easy|||||||||||
|[1501](https://leetcode.cn/problems/countries-you-can-safely-invest-in)|Medium|||||||||||
|[1511](https://leetcode.cn/problems/customer-order-frequency)|Easy|||||||||||
|[1517](https://leetcode.cn/problems/find-users-with-valid-e-mails)|Easy|||||||||||
|[1527](https://leetcode.cn/problems/patients-with-a-condition)|Easy|||||||||||
|[1532](https://leetcode.cn/problems/the-most-recent-three-orders)|Medium|||||||||||
|[1543](https://leetcode.cn/problems/fix-product-name-format)|Easy|||||||||||
|[1549](https://leetcode.cn/problems/the-most-recent-orders-for-each-product)|Medium|||||||||||
|[1555](https://leetcode.cn/problems/bank-account-summary)|Medium|||||||||||
|[1565](https://leetcode.cn/problems/unique-orders-and-customers-per-month)|Easy|||||||||||
|[1571](https://leetcode.cn/problems/warehouse-manager)|Easy|||||||||||
|[1581](https://leetcode.cn/problems/customer-who-visited-but-did-not-make-any-transactions)|Easy|||||||||||
|[1587](https://leetcode.cn/problems/bank-account-summary-ii)|Easy|||||||||||
|[1596](https://leetcode.cn/problems/the-most-frequently-ordered-products-for-each-customer)|Medium|||||||||||
|[1607](https://leetcode.cn/problems/sellers-with-no-sales)|Easy|||||||||||
|[1613](https://leetcode.cn/problems/find-the-missing-ids)|Medium|||||||||||
|[1623](https://leetcode.cn/problems/all-valid-triplets-that-can-represent-a-country)|Easy|||||||||||
|[1633](https://leetcode.cn/problems/percentage-of-users-attended-a-contest)|Easy|||||||||||
|[1635](https://leetcode.cn/problems/hopper-company-queries-i)|Hard|||||||||||
|[1645](https://leetcode.cn/problems/hopper-company-queries-ii)|Hard|||||||||||
|[1651](https://leetcode.cn/problems/hopper-company-queries-iii)|Hard|||||||||||
|[1661](https://leetcode.cn/problems/average-time-of-process-per-machine)|Easy|||||||||||
|[1667](https://leetcode.cn/problems/fix-names-in-a-table)|Easy|||||||||||
|[1677](https://leetcode.cn/problems/products-worth-over-invoices)|Easy|||||||||||
|[1683](https://leetcode.cn/problems/invalid-tweets)|Easy|||||||||||
|[1693](https://leetcode.cn/problems/daily-leads-and-partners)|Easy|||||||||||
|[1699](https://leetcode.cn/problems/number-of-calls-between-two-persons)|Medium|||||||||||
|[1709](https://leetcode.cn/problems/biggest-window-between-visits)|Medium|||||||||||
|[1715](https://leetcode.cn/problems/count-apples-and-oranges)|Medium|||||||||||
|[1729](https://leetcode.cn/problems/find-followers-count)|Easy|||||||||||
|[1731](https://leetcode.cn/problems/the-number-of-employees-which-report-to-each-employee)|Easy|||||||||||
|[1741](https://leetcode.cn/problems/find-total-time-spent-by-each-employee)|Easy|||||||||||
|[1747](https://leetcode.cn/problems/leetflex-banned-accounts)|Medium|||||||||||
|[1757](https://leetcode.cn/problems/recyclable-and-low-fat-products)|Easy|||||||||||
|[1767](https://leetcode.cn/problems/find-the-subtasks-that-did-not-execute)|Hard|||||||||||
|[1777](https://leetcode.cn/problems/products-price-for-each-store)|Easy|||||||||||
|[1783](https://leetcode.cn/problems/grand-slam-titles)|Medium|||||||||||
|[1789](https://leetcode.cn/problems/primary-department-for-each-employee)|Easy|||||||||||
|[1795](https://leetcode.cn/problems/rearrange-products-table)|Easy|||||||||||
|[1809](https://leetcode.cn/problems/ad-free-sessions)|Easy|||||||||||
|[1811](https://leetcode.cn/problems/find-interview-candidates)|Medium|||||||||||
|[1821](https://leetcode.cn/problems/find-customers-with-positive-revenue-this-year)|Easy|||||||||||
|[1831](https://leetcode.cn/problems/maximum-transaction-each-day)|Medium|||||||||||
|[1841](https://leetcode.cn/problems/league-statistics)|Medium|||||||||||
|[1843](https://leetcode.cn/problems/suspicious-bank-accounts)|Medium|||||||||||
|[1853](https://leetcode.cn/problems/convert-date-format)|Easy|||||||||||
|[1867](https://leetcode.cn/problems/orders-with-maximum-quantity-above-average)|Medium|||||||||||
|[1873](https://leetcode.cn/problems/calculate-special-bonus)|Easy|||||||||||
|[1875](https://leetcode.cn/problems/group-employees-of-the-same-salary)|Medium|||||||||||
|[1890](https://leetcode.cn/problems/the-latest-login-in-2020)|Easy|||||||||||
|[1892](https://leetcode.cn/problems/page-recommendations-ii)|Hard|||||||||||
|[1907](https://leetcode.cn/problems/count-salary-categories)|Medium|||||||||||
|[1917](https://leetcode.cn/problems/leetcodify-friends-recommendations)|Hard|||||||||||
|[1919](https://leetcode.cn/problems/leetcodify-similar-friends)|Hard|||||||||||
|[1934](https://leetcode.cn/problems/confirmation-rate)|Medium|||||||||||
|[1939](https://leetcode.cn/problems/users-that-actively-request-confirmation-messages)|Easy|||||||||||
|[1949](https://leetcode.cn/problems/strong-friendship)|Medium|||||||||||
|[1951](https://leetcode.cn/problems/all-the-pairs-with-the-maximum-number-of-common-followers)|Medium|||||||||||
|[1965](https://leetcode.cn/problems/employees-with-missing-information)|Easy|||||||||||
|[1972](https://leetcode.cn/problems/first-and-last-call-on-the-same-day)|Hard|||||||||||
|[1978](https://leetcode.cn/problems/employees-whose-manager-left-the-company)|Easy|||||||||||
|[1988](https://leetcode.cn/problems/find-cutoff-score-for-each-school)|Medium|||||||||||
|[1990](https://leetcode.cn/problems/count-the-number-of-experiments)|Medium|||||||||||
|[2004](https://leetcode.cn/problems/the-number-of-seniors-and-juniors-to-join-the-company)|Hard|||||||||||
|[2010](https://leetcode.cn/problems/the-number-of-seniors-and-juniors-to-join-the-company-ii)|Hard|||||||||||
|[2020](https://leetcode.cn/problems/number-of-accounts-that-did-not-stream)|Medium|||||||||||
|[2026](https://leetcode.cn/problems/low-quality-problems)|Easy|||||||||||
|[2041](https://leetcode.cn/problems/accepted-candidates-from-the-interviews)|Medium|||||||||||
|[2051](https://leetcode.cn/problems/the-category-of-each-member-in-the-store)|Medium|||||||||||
|[2066](https://leetcode.cn/problems/account-balance)|Medium|||||||||||
|[2072](https://leetcode.cn/problems/the-winner-university)|Easy|||||||||||
|[2082](https://leetcode.cn/problems/the-number-of-rich-customers)|Easy|||||||||||
|[2084](https://leetcode.cn/problems/drop-type-1-orders-for-customers-with-type-0-orders)|Medium|||||||||||
|[2112](https://leetcode.cn/problems/the-airport-with-the-most-traffic)|Medium|||||||||||
|[2118](https://leetcode.cn/problems/build-the-equation)|Hard|||||||||||
|[2142](https://leetcode.cn/problems/the-number-of-passengers-in-each-bus-i)|Medium|||||||||||
|[2153](https://leetcode.cn/problems/the-number-of-passengers-in-each-bus-ii)|Hard|||||||||||
|[2159](https://leetcode.cn/problems/order-two-columns-independently)|Medium|||||||||||
|[2173](https://leetcode.cn/problems/longest-winning-streak)|Hard|||||||||||
|[2175](https://leetcode.cn/problems/the-change-in-global-rankings)|Medium|||||||||||
|[2199](https://leetcode.cn/problems/finding-the-topic-of-each-post)|Hard|||||||||||
|[2205](https://leetcode.cn/problems/the-number-of-users-that-are-eligible-for-discount)|Easy|||||||||||
|[2228](https://leetcode.cn/problems/users-with-two-purchases-within-seven-days)|Medium|||||||||||
|[2230](https://leetcode.cn/problems/the-users-that-are-eligible-for-discount)|Easy|||||||||||
|[2238](https://leetcode.cn/problems/number-of-times-a-driver-was-a-passenger)|Medium|||||||||||
|[2252](https://leetcode.cn/problems/dynamic-pivoting-of-a-table)|Hard|||||||||||
|[2253](https://leetcode.cn/problems/dynamic-unpivoting-of-a-table)|Hard|||||||||||
|[2292](https://leetcode.cn/problems/products-with-three-or-more-orders-in-two-consecutive-years)|Medium|||||||||||
|[2298](https://leetcode.cn/problems/tasks-count-in-the-weekend)|Medium|||||||||||
|[2308](https://leetcode.cn/problems/arrange-table-by-gender)|Medium|||||||||||
|[2314](https://leetcode.cn/problems/the-first-day-of-the-maximum-recorded-degree-in-each-city)|Medium|||||||||||
|[2324](https://leetcode.cn/problems/product-sales-analysis-iv)|Medium|||||||||||
|[2329](https://leetcode.cn/problems/product-sales-analysis-v)|Easy|||||||||||
|[2339](https://leetcode.cn/problems/all-the-matches-of-the-league)|Easy|||||||||||
|[2346](https://leetcode.cn/problems/compute-the-rank-as-a-percentage)|Medium|||||||||||
|[2356](https://leetcode.cn/problems/number-of-unique-subjects-taught-by-each-teacher)|Easy|||||||||||
|[2362](https://leetcode.cn/problems/generate-the-invoice)|Hard|||||||||||
|[2372](https://leetcode.cn/problems/calculate-the-influence-of-each-salesperson)|Medium|||||||||||
|[2377](https://leetcode.cn/problems/sort-the-olympic-table)|Easy|||||||||||
|[2388](https://leetcode.cn/problems/change-null-values-in-a-table-to-the-previous-value)|Medium|||||||||||
|[2394](https://leetcode.cn/problems/employees-with-deductions)|Medium|||||||||||
|[2474](https://leetcode.cn/problems/customers-with-strictly-increasing-purchases)|Hard|||||||||||
|[2480](https://leetcode.cn/problems/form-a-chemical-bond)|Easy|||||||||||
|[2494](https://leetcode.cn/problems/merge-overlapping-events-in-the-same-hall)|Hard|||||||||||
|[2504](https://leetcode.cn/problems/concatenate-the-name-and-the-profession)|Easy|||||||||||
|[2668](https://leetcode.cn/problems/find-latest-salaries)|Easy|||||||||||
|[2669](https://leetcode.cn/problems/count-artist-occurrences-on-spotify-ranking-list)|Easy|||||||||||
|[2686](https://leetcode.cn/problems/immediate-food-delivery-iii)|Medium|||||||||||
|[2687](https://leetcode.cn/problems/bikes-last-time-used)|Easy|||||||||||
|[2688](https://leetcode.cn/problems/find-active-users)|Medium|||||||||||
|[2752](https://leetcode.cn/problems/customers-with-maximum-number-of-transactions-on-consecutive-days)|Hard|||||||||||
|[2701](https://leetcode.cn/problems/consecutive-transactions-with-increasing-amounts)|Hard|||||||||||
|[2720](https://leetcode.cn/problems/popularity-percentage)|Hard|||||||||||
|[2738](https://leetcode.cn/problems/count-occurrences-in-text)|Medium|||||||||||
|[2783](https://leetcode.cn/problems/flight-occupancy-and-waitlist-analysis)|Medium|||||||||||
|[2793](https://leetcode.cn/problems/status-of-flight-tickets)|Hard|||||||||||
|[2820](https://leetcode.cn/problems/election-results)|Medium|||||||||||
|[2837](https://leetcode.cn/problems/total-traveled-distance)|Easy|||||||||||
|[2853](https://leetcode.cn/problems/highest-salaries-difference)|Easy|||||||||||
|[2854](https://leetcode.cn/problems/rolling-average-steps)|Medium|||||||||||
|[2893](https://leetcode.cn/problems/calculate-orders-within-each-interval)|Medium|||||||||||
|[2922](https://leetcode.cn/problems/market-analysis-iii)|Medium|||||||||||
|[2978](https://leetcode.cn/problems/symmetric-coordinates)|Medium|||||||||||
|[3050](https://leetcode.cn/problems/pizza-toppings-cost-analysis)|Medium|||||||||||
|[3051](https://leetcode.cn/problems/find-candidates-for-data-scientist-position)|Easy|||||||||||
|[3052](https://leetcode.cn/problems/maximize-items)|Hard|||||||||||
|[3053](https://leetcode.cn/problems/classifying-triangles-by-lengths)|Easy|||||||||||
|[3054](https://leetcode.cn/problems/binary-tree-nodes)|Medium|||||||||||
|[2984](https://leetcode.cn/problems/find-peak-calling-hours-for-each-city)|Medium|||||||||||
|[2985](https://leetcode.cn/problems/calculate-compressed-mean)|Easy|||||||||||
|[2986](https://leetcode.cn/problems/find-third-transaction)|Medium|||||||||||
|[2987](https://leetcode.cn/problems/find-expensive-cities)|Easy|||||||||||
|[2988](https://leetcode.cn/problems/manager-of-the-largest-department)|Medium|||||||||||
|[2989](https://leetcode.cn/problems/class-performance)|Medium|||||||||||
|[2990](https://leetcode.cn/problems/loan-types)|Easy|||||||||||
|[2991](https://leetcode.cn/problems/top-three-wineries)|Hard|||||||||||
|[2993](https://leetcode.cn/problems/friday-purchases-i)|Medium|||||||||||
|[2994](https://leetcode.cn/problems/friday-purchases-ii)|Hard|||||||||||
|[2995](https://leetcode.cn/problems/viewers-turned-streamers)|Hard|||||||||||
|[3055](https://leetcode.cn/problems/top-percentile-fraud)|Medium|||||||||||
|[3056](https://leetcode.cn/problems/snaps-analysis)|Medium|||||||||||
|[3057](https://leetcode.cn/problems/employees-project-allocation)|Hard|||||||||||
|[3058](https://leetcode.cn/problems/friends-with-no-mutual-friends)|Medium|||||||||||
|[3059](https://leetcode.cn/problems/find-all-unique-email-domains)|Easy|||||||||||
|[3060](https://leetcode.cn/problems/user-activities-within-time-bounds)|Hard|||||||||||
|[3061](https://leetcode.cn/problems/calculate-trapping-rain-water)|Hard|||||||||||
|[3087](https://leetcode.cn/problems/find-trending-hashtags)|Medium|||||||||||
|[3089](https://leetcode.cn/problems/find-bursty-behavior)|Medium|||||||||||
|[3103](https://leetcode.cn/problems/find-trending-hashtags-ii)|Hard|||||||||||
|[3118](https://leetcode.cn/problems/friday-purchase-iii)|Medium|||||||||||
|[3124](https://leetcode.cn/problems/find-longest-calls)|Medium|||||||||||
|[3126](https://leetcode.cn/problems/server-utilization-time)|Medium|||||||||||
|[3140](https://leetcode.cn/problems/consecutive-available-seats-ii)|Medium|||||||||||
|[3150](https://leetcode.cn/problems/invalid-tweets-ii)|Easy|||||||||||
|[3156](https://leetcode.cn/problems/employee-task-duration-and-concurrent-tasks)|Hard|||||||||||
|[3166](https://leetcode.cn/problems/calculate-parking-fees-and-duration)|Medium|||||||||||
|[3172](https://leetcode.cn/problems/second-day-verification)|Easy|||||||||||
|[3182](https://leetcode.cn/problems/find-top-scoring-students)|Medium|||||||||||
|[3188](https://leetcode.cn/problems/find-top-scoring-students-ii)|Hard|||||||||||
|[3198](https://leetcode.cn/problems/find-cities-in-each-state)|Easy|||||||||||
|[3204](https://leetcode.cn/problems/bitwise-user-permissions-analysis)|Medium|||||||||||
|[3214](https://leetcode.cn/problems/year-on-year-growth-rate)|Hard|||||||||||
|[3220](https://leetcode.cn/problems/odd-and-even-transactions)|Medium|||||||||||
|[3230](https://leetcode.cn/problems/customer-purchasing-behavior-analysis)|Medium|||||||||||
|[3236](https://leetcode.cn/problems/ceo-subordinate-hierarchy)|Hard|||||||||||
|[3246](https://leetcode.cn/problems/premier-league-table-ranking)|Easy|||||||||||
|[3252](https://leetcode.cn/problems/premier-league-table-ranking-ii)|Medium|||||||||||
|[3262](https://leetcode.cn/problems/find-overlapping-shifts)|Medium|||||||||||
|[3268](https://leetcode.cn/problems/find-overlapping-shifts-ii)|Hard|||||||||||
|[3278](https://leetcode.cn/problems/find-candidates-for-data-scientist-position-ii)|Medium|||||||||||
|[3293](https://leetcode.cn/problems/calculate-product-final-price)|Medium|||||||||||
|[3308](https://leetcode.cn/problems/find-top-performing-driver)|Medium|||||||||||
|[3322](https://leetcode.cn/problems/premier-league-table-ranking-iii)|Medium|||||||||||
|[3328](https://leetcode.cn/problems/find-cities-in-each-state-ii)|Medium|||||||||||
|[3338](https://leetcode.cn/problems/second-highest-salary-ii)|Medium|||||||||||

</details>

<details>

  <summary>Shell</summary>

|Index|Difficulty|C|C++|Go|Java|TypeScript|PHP|Python|Rust|SQL|Bash|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
|[0192](https://leetcode.cn/problems/word-frequency)|Medium||||||||||:white_check_mark:|
|[0193](https://leetcode.cn/problems/valid-phone-numbers)|Easy||||||||||:white_check_mark:|
|[0194](https://leetcode.cn/problems/transpose-file)|Medium||||||||||:white_check_mark:|
|[0195](https://leetcode.cn/problems/tenth-line)|Easy||||||||||:white_check_mark:|

</details>

<details>

  <summary>Concurrency</summary>

|Index|Difficulty|C|C++|Go|Java|TypeScript|PHP|Python|Rust|SQL|Bash|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
|[1117](https://leetcode.cn/problems/building-h2o)|Medium|||||||||||
|[1115](https://leetcode.cn/problems/print-foobar-alternately)|Medium|||||||||||
|[1114](https://leetcode.cn/problems/print-in-order)|Easy|||||||||||
|[1188](https://leetcode.cn/problems/design-bounded-blocking-queue)|Medium|||||||||||
|[1116](https://leetcode.cn/problems/print-zero-even-odd)|Medium|||||||||||
|[1195](https://leetcode.cn/problems/fizz-buzz-multithreaded)|Medium|||||||||||
|[1226](https://leetcode.cn/problems/the-dining-philosophers)|Medium|||||||||||
|[1242](https://leetcode.cn/problems/web-crawler-multithreaded)|Medium|||||||||||
|[1279](https://leetcode.cn/problems/traffic-light-controlled-intersection)|Easy|||||||||||

</details>

<details>

  <summary>JavaScript</summary>

|Index|Difficulty|C|C++|Go|Java|TypeScript|PHP|Python|Rust|SQL|Bash|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
|[2623](https://leetcode.cn/problems/memoize)|Medium|||||||||||
|[2620](https://leetcode.cn/problems/counter)|Easy|||||||||||
|[2621](https://leetcode.cn/problems/sleep)|Easy|||||||||||
|[2619](https://leetcode.cn/problems/array-prototype-last)|Easy|||||||||||
|[2628](https://leetcode.cn/problems/json-deep-equal)|Medium|||||||||||
|[2632](https://leetcode.cn/problems/curry)|Medium|||||||||||
|[2629](https://leetcode.cn/problems/function-composition)|Easy|||||||||||
|[2631](https://leetcode.cn/problems/group-by)|Medium|||||||||||
|[2627](https://leetcode.cn/problems/debounce)|Medium|||||||||||
|[2630](https://leetcode.cn/problems/memoize-ii)|Hard|||||||||||
|[2633](https://leetcode.cn/problems/convert-object-to-json-string)|Medium|||||||||||
|[2634](https://leetcode.cn/problems/filter-elements-from-array)|Easy|||||||||||
|[2635](https://leetcode.cn/problems/apply-transform-over-each-element-in-array)|Easy|||||||||||
|[2637](https://leetcode.cn/problems/promise-time-limit)|Medium|||||||||||
|[2636](https://leetcode.cn/problems/promise-pool)|Medium|||||||||||
|[2618](https://leetcode.cn/problems/check-if-object-instance-of-class)|Medium|||||||||||
|[2625](https://leetcode.cn/problems/flatten-deeply-nested-array)|Medium|||||||||||
|[2624](https://leetcode.cn/problems/snail-traversal)|Medium|||||||||||
|[2626](https://leetcode.cn/problems/array-reduce-transformation)|Easy|||||||||||
|[2622](https://leetcode.cn/problems/cache-with-time-limit)|Medium|||||||||||
|[2774](https://leetcode.cn/problems/array-upper-bound)|Easy|||||||||||
|[2675](https://leetcode.cn/problems/array-of-objects-to-matrix)|Hard|||||||||||
|[2754](https://leetcode.cn/problems/bind-function-to-context)|Medium|||||||||||
|[2676](https://leetcode.cn/problems/throttle)|Medium|||||||||||
|[2690](https://leetcode.cn/problems/infinite-method-object)|Easy|||||||||||
|[2691](https://leetcode.cn/problems/immutability-helper)|Hard|||||||||||
|[2700](https://leetcode.cn/problems/differences-between-two-objects)|Medium|||||||||||
|[2648](https://leetcode.cn/problems/generate-fibonacci-sequence)|Easy|||||||||||
|[2692](https://leetcode.cn/problems/make-object-immutable)|Medium|||||||||||
|[2775](https://leetcode.cn/problems/undefined-to-null)|Medium|||||||||||
|[2776](https://leetcode.cn/problems/convert-callback-based-function-to-promise-based-function)|Medium|||||||||||
|[2649](https://leetcode.cn/problems/nested-array-generator)|Medium|||||||||||
|[2650](https://leetcode.cn/problems/design-cancellable-function)|Hard|||||||||||
|[2665](https://leetcode.cn/problems/counter-ii)|Easy|||||||||||
|[2693](https://leetcode.cn/problems/call-function-with-custom-context)|Medium|||||||||||
|[2755](https://leetcode.cn/problems/deep-merge-of-two-objects)|Medium|||||||||||
|[2666](https://leetcode.cn/problems/allow-one-function-call)|Easy|||||||||||
|[2694](https://leetcode.cn/problems/event-emitter)|Medium|||||||||||
|[2677](https://leetcode.cn/problems/chunk-array)|Easy|||||||||||
|[2777](https://leetcode.cn/problems/date-range-generator)|Medium|||||||||||
|[2705](https://leetcode.cn/problems/compact-object)|Medium|||||||||||
|[2695](https://leetcode.cn/problems/array-wrapper)|Easy|||||||||||
|[2756](https://leetcode.cn/problems/query-batching)|Hard|||||||||||
|[2721](https://leetcode.cn/problems/execute-asynchronous-functions-in-parallel)|Medium|||||||||||
|[2667](https://leetcode.cn/problems/create-hello-world-function)|Easy|||||||||||
|[2704](https://leetcode.cn/problems/to-be-or-not-to-be)|Easy|||||||||||
|[2757](https://leetcode.cn/problems/generate-circular-array-values)|Medium|||||||||||
|[2703](https://leetcode.cn/problems/return-length-of-arguments-passed)|Easy|||||||||||
|[2715](https://leetcode.cn/problems/timeout-cancellation)|Easy|||||||||||
|[2758](https://leetcode.cn/problems/next-day)|Easy|||||||||||
|[2759](https://leetcode.cn/problems/convert-json-string-to-object)|Hard|||||||||||
|[2722](https://leetcode.cn/problems/join-two-arrays-by-id)|Medium|||||||||||
|[2723](https://leetcode.cn/problems/add-two-promises)|Easy|||||||||||
|[2724](https://leetcode.cn/problems/sort-by)|Easy|||||||||||
|[2794](https://leetcode.cn/problems/create-object-from-two-arrays)|Easy|||||||||||
|[2725](https://leetcode.cn/problems/interval-cancellation)|Easy|||||||||||
|[2726](https://leetcode.cn/problems/calculator-with-method-chaining)|Easy|||||||||||
|[2727](https://leetcode.cn/problems/is-object-empty)|Easy|||||||||||
|[2795](https://leetcode.cn/problems/parallel-execution-of-promises-for-individual-results-retrieval)|Medium|||||||||||
|[2796](https://leetcode.cn/problems/repeat-string)|Easy|||||||||||
|[2797](https://leetcode.cn/problems/partial-function-with-placeholders)|Easy|||||||||||
|[2803](https://leetcode.cn/problems/factorial-generator)|Easy|||||||||||
|[2804](https://leetcode.cn/problems/array-prototype-foreach)|Easy|||||||||||
|[2805](https://leetcode.cn/problems/custom-interval)|Medium|||||||||||
|[2821](https://leetcode.cn/problems/delay-the-resolution-of-each-promise)|Medium|||||||||||
|[2823](https://leetcode.cn/problems/deep-object-filter)|Medium|||||||||||
|[2822](https://leetcode.cn/problems/inversion-of-object)|Easy|||||||||||

</details>

<details>

  <summary>Frontend</summary>

|Index|Difficulty|C|C++|Go|Java|TypeScript|PHP|Python|Rust|SQL|Bash|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|

</details>

<details>

  <summary>pandas</summary>

|Index|Difficulty|C|C++|Go|Java|TypeScript|PHP|Python|Rust|SQL|Bash|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
|[2877](https://leetcode.cn/problems/create-a-dataframe-from-list)|Easy|||||||||||
|[2891](https://leetcode.cn/problems/method-chaining)|Easy|||||||||||
|[2888](https://leetcode.cn/problems/reshape-data-concatenate)|Easy|||||||||||
|[2879](https://leetcode.cn/problems/display-the-first-three-rows)|Easy|||||||||||
|[2881](https://leetcode.cn/problems/create-a-new-column)|Easy|||||||||||
|[2884](https://leetcode.cn/problems/modify-columns)|Easy|||||||||||
|[2885](https://leetcode.cn/problems/rename-columns)|Easy|||||||||||
|[2886](https://leetcode.cn/problems/change-data-type)|Easy|||||||||||
|[2887](https://leetcode.cn/problems/fill-missing-data)|Easy|||||||||||
|[2882](https://leetcode.cn/problems/drop-duplicate-rows)|Easy|||||||||||
|[2889](https://leetcode.cn/problems/reshape-data-pivot)|Easy|||||||||||
|[2890](https://leetcode.cn/problems/reshape-data-melt)|Easy|||||||||||
|[2880](https://leetcode.cn/problems/select-data)|Easy|||||||||||
|[2883](https://leetcode.cn/problems/drop-missing-data)|Easy|||||||||||
|[2878](https://leetcode.cn/problems/get-the-size-of-a-dataframe)|Easy|||||||||||

</details>

<details>

  <summary>LCCI</summary>

|Index|Difficulty|C|C++|Go|Java|TypeScript|PHP|Python|Rust|SQL|Bash|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
|[面试题 01.01](https://leetcode.cn/problems/is-unique-lcci)|Easy|||||||||||
|[面试题 01.02](https://leetcode.cn/problems/check-permutation-lcci)|Easy|||||||||||
|[面试题 01.03](https://leetcode.cn/problems/string-to-url-lcci)|Easy|||||||||||
|[面试题 01.06](https://leetcode.cn/problems/compress-string-lcci)|Easy|||||||||||
|[面试题 01.09](https://leetcode.cn/problems/string-rotation-lcci)|Easy|||||||||||
|[面试题 02.01](https://leetcode.cn/problems/remove-duplicate-node-lcci)|Easy|||||||||||
|[面试题 02.06](https://leetcode.cn/problems/palindrome-linked-list-lcci)|Easy|||||||||||
|[面试题 02.07](https://leetcode.cn/problems/intersection-of-two-linked-lists-lcci)|Easy|||||||||||
|[面试题 02.08](https://leetcode.cn/problems/linked-list-cycle-lcci)|Medium|||||||||||
|[面试题 03.02](https://leetcode.cn/problems/min-stack-lcci)|Easy|||||||||||
|[面试题 03.04](https://leetcode.cn/problems/implement-queue-using-stacks-lcci)|Easy|||||||||||
|[面试题 04.01](https://leetcode.cn/problems/route-between-nodes-lcci)|Medium|||||||||||
|[面试题 03.01](https://leetcode.cn/problems/three-in-one-lcci)|Easy|||||||||||
|[面试题 03.05](https://leetcode.cn/problems/sort-of-stacks-lcci)|Medium|||||||||||
|[面试题 04.02](https://leetcode.cn/problems/minimum-height-tree-lcci)|Easy|||||||||||
|[面试题 04.03](https://leetcode.cn/problems/list-of-depth-lcci)|Medium|||||||||||
|[面试题 04.04](https://leetcode.cn/problems/check-balance-lcci)|Easy|||||||||||
|[面试题 04.05](https://leetcode.cn/problems/legal-binary-search-tree-lcci)|Medium|||||||||||
|[面试题 04.06](https://leetcode.cn/problems/successor-lcci)|Medium|||||||||||
|[面试题 04.08](https://leetcode.cn/problems/first-common-ancestor-lcci)|Medium|||||||||||
|[面试题 05.01](https://leetcode.cn/problems/insert-into-bits-lcci)|Easy|||||||||||
|[面试题 05.06](https://leetcode.cn/problems/convert-integer-lcci)|Easy|||||||||||
|[面试题 05.07](https://leetcode.cn/problems/exchange-lcci)|Easy|||||||||||
|[面试题 05.04](https://leetcode.cn/problems/closed-number-lcci)|Medium|||||||||||
|[面试题 01.04](https://leetcode.cn/problems/palindrome-permutation-lcci)|Easy|||||||||||
|[面试题 01.07](https://leetcode.cn/problems/rotate-matrix-lcci)|Medium|||||||||||
|[面试题 01.08](https://leetcode.cn/problems/zero-matrix-lcci)|Medium|||||||||||
|[面试题 02.03](https://leetcode.cn/problems/delete-middle-node-lcci)|Easy|||||||||||
|[面试题 02.05](https://leetcode.cn/problems/sum-lists-lcci)|Medium|||||||||||
|[面试题 03.03](https://leetcode.cn/problems/stack-of-plates-lcci)|Medium|||||||||||
|[面试题 05.08](https://leetcode.cn/problems/draw-line-lcci)|Medium|||||||||||
|[面试题 08.01](https://leetcode.cn/problems/three-steps-problem-lcci)|Easy|||||||||||
|[面试题 08.04](https://leetcode.cn/problems/power-set-lcci)|Medium|||||||||||
|[面试题 08.05](https://leetcode.cn/problems/recursive-mulitply-lcci)|Medium|||||||||||
|[面试题 08.09](https://leetcode.cn/problems/bracket-lcci)|Medium|||||||||||
|[面试题 08.10](https://leetcode.cn/problems/color-fill-lcci)|Easy|||||||||||
|[面试题 08.13](https://leetcode.cn/problems/pile-box-lcci)|Hard|||||||||||
|[面试题 05.02](https://leetcode.cn/problems/binary-number-to-string-lcci)|Medium|||||||||||
|[面试题 03.06](https://leetcode.cn/problems/animal-shelter-lcci)|Easy|||||||||||
|[面试题 04.10](https://leetcode.cn/problems/check-subtree-lcci)|Medium|||||||||||
|[面试题 05.03](https://leetcode.cn/problems/reverse-bits-lcci)|Easy|||||||||||
|[面试题 08.11](https://leetcode.cn/problems/coin-lcci)|Medium|||||||||||
|[面试题 10.03](https://leetcode.cn/problems/search-rotate-array-lcci)|Medium|||||||||||
|[面试题 08.12](https://leetcode.cn/problems/eight-queens-lcci)|Hard|||||||||||
|[面试题 08.03](https://leetcode.cn/problems/magic-index-lcci)|Easy|||||||||||
|[面试题 08.07](https://leetcode.cn/problems/permutation-i-lcci)|Medium|||||||||||
|[面试题 10.05](https://leetcode.cn/problems/sparse-array-search-lcci)|Easy|||||||||||
|[面试题 16.01](https://leetcode.cn/problems/swap-numbers-lcci)|Medium|||||||||||
|[面试题 16.02](https://leetcode.cn/problems/words-frequency-lcci)|Medium|||||||||||
|[面试题 16.03](https://leetcode.cn/problems/intersection-lcci)|Hard|||||||||||
|[面试题 16.04](https://leetcode.cn/problems/tic-tac-toe-lcci)|Medium|||||||||||
|[面试题 16.06](https://leetcode.cn/problems/smallest-difference-lcci)|Medium|||||||||||
|[面试题 08.08](https://leetcode.cn/problems/permutation-ii-lcci)|Medium|||||||||||
|[面试题 16.07](https://leetcode.cn/problems/maximum-lcci)|Easy|||||||||||
|[面试题 16.09](https://leetcode.cn/problems/operations-lcci)|Medium|||||||||||
|[面试题 16.10](https://leetcode.cn/problems/living-people-lcci)|Medium|||||||||||
|[面试题 16.11](https://leetcode.cn/problems/diving-board-lcci)|Easy|||||||||||
|[面试题 16.13](https://leetcode.cn/problems/bisect-squares-lcci)|Medium|||||||||||
|[面试题 16.14](https://leetcode.cn/problems/best-line-lcci)|Medium|||||||||||
|[面试题 16.15](https://leetcode.cn/problems/master-mind-lcci)|Easy|||||||||||
|[面试题 16.16](https://leetcode.cn/problems/sub-sort-lcci)|Medium|||||||||||
|[面试题 16.17](https://leetcode.cn/problems/contiguous-sequence-lcci)|Easy|||||||||||
|[面试题 16.18](https://leetcode.cn/problems/pattern-matching-lcci)|Medium|||||||||||
|[面试题 16.19](https://leetcode.cn/problems/pond-sizes-lcci)|Medium|||||||||||
|[面试题 01.05](https://leetcode.cn/problems/one-away-lcci)|Medium|||||||||||
|[面试题 02.02](https://leetcode.cn/problems/kth-node-from-end-of-list-lcci)|Easy|||||||||||
|[面试题 02.04](https://leetcode.cn/problems/partition-list-lcci)|Medium|||||||||||
|[面试题 04.12](https://leetcode.cn/problems/paths-with-sum-lcci)|Medium|||||||||||
|[面试题 04.09](https://leetcode.cn/problems/bst-sequences-lcci)|Hard|||||||||||
|[面试题 08.02](https://leetcode.cn/problems/robot-in-a-grid-lcci)|Medium|||||||||||
|[面试题 10.01](https://leetcode.cn/problems/sorted-merge-lcci)|Easy|||||||||||
|[面试题 10.11](https://leetcode.cn/problems/peaks-and-valleys-lcci)|Medium|||||||||||
|[面试题 08.06](https://leetcode.cn/problems/hanota-lcci)|Easy|||||||||||
|[面试题 10.09](https://leetcode.cn/problems/sorted-matrix-search-lcci)|Medium|||||||||||
|[面试题 16.05](https://leetcode.cn/problems/factorial-zeros-lcci)|Easy|||||||||||
|[面试题 16.08](https://leetcode.cn/problems/english-int-lcci)|Hard|||||||||||
|[面试题 17.12](https://leetcode.cn/problems/binode-lcci)|Easy|||||||||||
|[面试题 17.13](https://leetcode.cn/problems/re-space-lcci)|Medium|||||||||||
|[面试题 17.14](https://leetcode.cn/problems/smallest-k-lcci)|Medium|||||||||||
|[面试题 17.15](https://leetcode.cn/problems/longest-word-lcci)|Medium|||||||||||
|[面试题 17.16](https://leetcode.cn/problems/the-masseuse-lcci)|Easy|||||||||||
|[面试题 17.17](https://leetcode.cn/problems/multi-search-lcci)|Medium|||||||||||
|[面试题 17.01](https://leetcode.cn/problems/add-without-plus-lcci)|Easy|:white_check_mark:|:white_check_mark:|||||||||
|[面试题 16.25](https://leetcode.cn/problems/lru-cache-lcci)|Medium|||||||||||
|[面试题 16.26](https://leetcode.cn/problems/calculator-lcci)|Medium|||||||||||
|[面试题 17.20](https://leetcode.cn/problems/continuous-median-lcci)|Hard|||||||||||
|[面试题 17.21](https://leetcode.cn/problems/volume-of-histogram-lcci)|Hard|||||||||||
|[面试题 17.22](https://leetcode.cn/problems/word-transformer-lcci)|Medium|||||||||||
|[面试题 08.14](https://leetcode.cn/problems/boolean-evaluation-lcci)|Medium|||||||||||
|[面试题 17.04](https://leetcode.cn/problems/missing-number-lcci)|Easy|||||||||||
|[面试题 17.05](https://leetcode.cn/problems/find-longest-subarray-lcci)|Medium|||||||||||
|[面试题 17.06](https://leetcode.cn/problems/number-of-2s-in-range-lcci)|Hard|||||||||||
|[面试题 17.07](https://leetcode.cn/problems/baby-names-lcci)|Medium|||||||||||
|[面试题 17.08](https://leetcode.cn/problems/circus-tower-lcci)|Medium|||||||||||
|[面试题 17.09](https://leetcode.cn/problems/get-kth-magic-number-lcci)|Medium|||||||||||
|[面试题 17.10](https://leetcode.cn/problems/find-majority-element-lcci)|Easy|||||||||||
|[面试题 17.11](https://leetcode.cn/problems/find-closest-lcci)|Medium|||||||||||
|[面试题 10.02](https://leetcode.cn/problems/group-anagrams-lcci)|Medium|||||||||||
|[面试题 10.10](https://leetcode.cn/problems/rank-from-stream-lcci)|Medium|||||||||||
|[面试题 16.24](https://leetcode.cn/problems/pairs-with-sum-lcci)|Medium|||||||||||
|[面试题 17.18](https://leetcode.cn/problems/shortest-supersequence-lcci)|Medium|||||||||||
|[面试题 17.19](https://leetcode.cn/problems/missing-two-lcci)|Hard|||||||||||
|[面试题 17.23](https://leetcode.cn/problems/max-black-square-lcci)|Medium|||||||||||
|[面试题 17.24](https://leetcode.cn/problems/max-submatrix-lcci)|Hard|||||||||||
|[面试题 16.20](https://leetcode.cn/problems/t9-lcci)|Medium|||||||||||
|[面试题 16.21](https://leetcode.cn/problems/sum-swap-lcci)|Medium|||||||||||
|[面试题 17.25](https://leetcode.cn/problems/word-rectangle-lcci)|Hard|||||||||||
|[面试题 16.22](https://leetcode.cn/problems/langtons-ant-lcci)|Medium|||||||||||
|[面试题 17.26](https://leetcode.cn/problems/sparse-similarity-lcci)|Hard|||||||||||

</details>
