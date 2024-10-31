### [哈希分割字符串](https://leetcode-cn.com/problems/hash-divided-string)

<p>给你一个长度为 <code>n</code>&nbsp;的字符串 <code>s</code>&nbsp;和一个整数&nbsp;<code>k</code>&nbsp;，<code>n</code>&nbsp;是 <code>k</code>&nbsp;的 <strong>倍数</strong>&nbsp;。你的任务是将字符串 <code>s</code>&nbsp;哈希为一个长度为 <code>n / k</code>&nbsp;的新字符串&nbsp;<code>result</code>&nbsp;。</p>

<p>首先，将&nbsp;<code>s</code>&nbsp;分割成&nbsp;<code>n / k</code>&nbsp;个&nbsp;<strong><span data-keyword="substring-nonempty">子字符串</span></strong>&nbsp;，每个子字符串的长度都为&nbsp;<code>k</code>&nbsp;。然后，将&nbsp;<code>result</code>&nbsp;初始化为一个 <strong>空</strong>&nbsp;字符串。</p>

<p>我们依次从前往后处理每一个 <strong>子字符串</strong>&nbsp;：</p>

<ul>
	<li>一个字符的 <strong>哈希值</strong>&nbsp;是它在 <strong>字母表</strong>&nbsp;中的下标（也就是&nbsp;<code>'a' →<!-- notionvc: d3f8e4c2-23cd-41ad-a14b-101dfe4c5aba --> 0</code>&nbsp;，<code>'b' →<!-- notionvc: d3f8e4c2-23cd-41ad-a14b-101dfe4c5aba --> 1</code>&nbsp;，... ，<code>'z' →<!-- notionvc: d3f8e4c2-23cd-41ad-a14b-101dfe4c5aba --> 25</code>）。</li>
	<li>将子字符串中字母的 <strong>哈希值</strong>&nbsp;求和。</li>
	<li>将和对 26 取余，将结果记为&nbsp;<code>hashedChar</code>&nbsp;。</li>
	<li>找到小写字母表中 <code>hashedChar</code>&nbsp;对应的字符。</li>
	<li>将该字符添加到&nbsp;<code>result</code>&nbsp;的末尾。</li>
</ul>

<p>返回&nbsp;<code>result</code>&nbsp;。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>s = "abcd", k = 2</span></p>

<p><span class="example-io"><b>输出：</b>"bf"</span></p>

<p><b>解释：</b></p>

<p>第一个字符串为&nbsp;<code>"ab"</code>&nbsp;，<code>0 + 1 = 1</code>&nbsp;，<code>1 % 26 = 1</code>&nbsp;，<code>result[0] = 'b'</code>&nbsp;。</p>

<p>第二个字符串为： <code>"cd"</code>&nbsp;，<code>2 + 3 = 5</code>&nbsp;，<code>5 % 26 = 5</code>&nbsp;，<code>result[1] = 'f'</code>&nbsp;。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>s = "mxz", k = 3</span></p>

<p><span class="example-io"><b>输出：</b>"i"</span></p>

<p><b>解释：</b></p>

<p>唯一的子字符串为&nbsp;<code>"mxz"</code>&nbsp;，<code>12 + 23 + 25 = 60</code>&nbsp;，<code>60 % 26 = 8</code>&nbsp;，<code>result[0] = 'i'</code>&nbsp;。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= k &lt;= 100</code></li>
	<li><code>k &lt;= s.length &lt;= 1000</code></li>
	<li><code>s.length</code>&nbsp;能被 <code>k</code>&nbsp;整除。</li>
	<li><code>s</code> 只含有小写英文字母。</li>
</ul>