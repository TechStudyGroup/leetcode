### [完全子集的最大元素和](https://leetcode-cn.com/problems/maximum-element-sum-of-a-complete-subset-of-indices)

<p>给你一个下标从 <strong>1</strong> 开始、由 <code>n</code> 个整数组成的数组。你需要从&nbsp;<code>nums</code>&nbsp;选择一个&nbsp;<strong>完全集</strong>，其中每对元素下标的乘积都是一个 <span data-keyword="perfect-square">完全平方数</span>，例如选择&nbsp;<code>a<sub>i</sub></code>&nbsp;和&nbsp;<code>a<sub>j</sub></code>&nbsp;，<code>i * j</code>&nbsp;一定是完全平方数。</p>

<p>返回&nbsp;<strong>完全子集</strong> 所能取到的 <strong>最大元素和</strong> 。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">nums = [8,7,3,5,7,2,4,9]</span></p>

<p><strong>输出：</strong><span class="example-io">16</span></p>

<p><strong>解释：</strong></p>

<p>我们选择下标为 2 和 8 的元素，并且&nbsp;<code>2 * 8</code>&nbsp;是一个完全平方数。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [8,10,3,8,1,13,7,9,4]</span></p>

<p><span class="example-io"><b>输出：</b>20</span></p>

<p><strong>解释：</strong></p>

<p>我们选择下标为 1, 4, 9 的元素。<code>1 * 4</code>, <code>1 * 9</code>, <code>4 * 9</code>&nbsp;是完全平方数。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n == nums.length &lt;= 10<sup>4</sup></code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>9</sup></code></li>
</ul>