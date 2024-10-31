### [转换数组中的每个元素](https://leetcode-cn.com/problems/apply-transform-over-each-element-in-array)

<p>编写一个函数，这个函数接收一个整数数组&nbsp;<code>arr</code> 和一个映射函数&nbsp; <code>fn</code>&nbsp;，通过该映射函数返回一个新的数组。</p>

<p>返回数组的创建语句应为 <code>returnedArray[i] = fn(arr[i], i)</code>&nbsp;。</p>

<p>请你在不使用内置方法&nbsp;<code>Array.map</code>&nbsp;的前提下解决这个问题。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1:</strong></p>

<pre>
<strong>输入：</strong>arr = [1,2,3], fn = function plusone(n) { return n + 1; }
<strong>输出：</strong>[2,3,4]
<strong>解释： </strong>
const newArray = map(arr, plusone); // [2,3,4]
此映射函数返回值是将数组中每个元素的值加 1。
</pre>

<p><strong class="example">示例</strong><strong class="example"> 2:</strong></p>

<pre>
<strong>输入：</strong>arr = [1,2,3], fn = function plusI(n, i) { return n + i; }
<strong>输出：</strong>[1,3,5]
<strong>解释：</strong>此映射函数返回值根据输入数组索引增加每个值。
</pre>

<p><strong class="example">示例&nbsp;3:</strong></p>

<pre>
<strong>输入：</strong>arr = [10,20,30], fn = function constant() { return 42; }
<strong>输出：</strong>[42,42,42]
<strong>解释：</strong>此映射函数返回值恒为 42。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>0 &lt;= arr.length &lt;= 1000</code></li>
	<li><code><font face="monospace">-10<sup>9</sup>&nbsp;&lt;= arr[i] &lt;= 10<sup>9</sup></font></code></li>
	<li><font face="monospace"><code>fn</code> 返回一个数</font></li>
</ul>
<span style="display:block"><span style="height:0px"><span style="position:absolute">​​​​​​</span></span></span>