### [Find the Maximum Achievable Number](https://leetcode.com/problems/find-the-maximum-achievable-number)

<p>Given two integers, <code>num</code> and <code>t</code>. A <strong>number is achievable</strong> if it can become equal to <code>num</code> after applying the following operation:</p>

<ul>
	<li>Increase or decrease the number by <code>1</code>, and simultaneously increase or decrease <code>num</code> by <code>1</code>.</li>
</ul>

<p>Return the <strong>maximum achievable number</strong> after applying the operation at most <code>t</code> times.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">num = 4, t = 1</span></p>

<p><strong>Output:</strong> <span class="example-io">6</span></p>

<p><strong>Explanation:</strong></p>

<p>Apply the following operation once to make the maximum achievable number equal to <code>num</code>:</p>

<ul>
	<li>Decrease the maximum achievable number by 1, and increase <code>num</code> by 1.</li>
</ul>
</div>

<p><strong class="example">Example 2:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">num = 3, t = 2</span></p>

<p><strong>Output:</strong> <span class="example-io">7</span></p>

<p><strong>Explanation:</strong></p>

<p>Apply the following operation twice to make the maximum achievable number equal to <code>num</code>:</p>

<ul>
	<li>Decrease the maximum achievable number by 1, and increase <code>num</code> by 1.</li>
</ul>
</div>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= num, t&nbsp;&lt;= 50</code></li>
</ul>
