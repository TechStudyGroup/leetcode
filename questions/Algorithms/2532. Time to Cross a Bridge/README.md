### [Time to Cross a Bridge](https://leetcode.com/problems/time-to-cross-a-bridge)

<p>There are <code>k</code> workers who want to move <code>n</code> boxes from the right (old) warehouse to the left (new) warehouse. You are given the two integers <code>n</code> and <code>k</code>, and a 2D integer array <code>time</code> of size <code>k x 4</code> where <code>time[i] = [right<sub>i</sub>, pick<sub>i</sub>, left<sub>i</sub>, put<sub>i</sub>]</code>.</p>

<p>The warehouses are separated by a river and connected by a bridge. Initially, all <code>k</code> workers are waiting on the left side of the bridge. To move the boxes, the <code>i<sup>th</sup></code> worker can do the following:</p>

<ul>
	<li>Cross the bridge to the right side in <code>right<sub>i</sub></code> minutes.</li>
	<li>Pick a box from the right warehouse in <code>pick<sub>i</sub></code> minutes.</li>
	<li>Cross the bridge to the left side in <code>left<sub>i</sub></code> minutes.</li>
	<li>Put the box into the left warehouse in <code>put<sub>i</sub></code> minutes.</li>
</ul>

<p>The <code>i<sup>th</sup></code> worker is <strong>less efficient</strong> than the j<code><sup>th</sup></code> worker if either condition is met:</p>

<ul>
	<li><code>left<sub>i</sub> + right<sub>i</sub> &gt; left<sub>j</sub> + right<sub>j</sub></code></li>
	<li><code>left<sub>i</sub> + right<sub>i</sub> == left<sub>j</sub> + right<sub>j</sub></code> and <code>i &gt; j</code></li>
</ul>

<p>The following rules regulate the movement of the workers through the bridge:</p>

<ul>
	<li>Only one worker can use the bridge at a time.</li>
	<li>When the bridge is unused prioritize the <strong>least efficient</strong> worker on the right side to cross. If there are no workers on the right side, prioritize the <strong>least efficient</strong> worker on the left side to cross.</li>
	<li>If enough workers have already been dispatched from the left side to pick up all the remaining boxes, <strong>no more</strong> workers will be sent from the left side.</li>
</ul>

<p>Return the <strong>elapsed minutes</strong> at which the last box reaches the <strong>left side of the bridge</strong>.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">n = 1, k = 3, time = [[1,1,2,1],[1,1,3,1],[1,1,4,1]]</span></p>

<p><strong>Output:</strong> <span class="example-io">6</span></p>

<p><strong>Explanation:</strong></p>

<pre>
From 0 to 1 minutes: worker 2 crosses the bridge to the right.
From 1 to 2 minutes: worker 2 picks up a box from the right warehouse.
From 2 to 6 minutes: worker 2 crosses the bridge to the left.
From 6 to 7 minutes: worker 2 puts a box at the left warehouse.
The whole process ends after 7 minutes. We return 6 because the problem asks for the instance of time at which the last worker reaches the left side of the bridge.
</pre>
</div>

<p><strong class="example">Example 2:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">n = 3, k = 2, time = [[1,9,1,8],[10,10,10,10]]</span></p>

<p><strong>Output:</strong> <span class="example-io">50</span></p>

<p><strong>Explanation:</strong></p>

<pre>
From 0  to 10: worker 1 crosses the bridge to the right.
From 10 to 20: worker 1 picks up a box from the right warehouse.
From 10 to 11: worker 0 crosses the bridge to the right.
From 11 to 20: worker 0 picks up a box from the right warehouse.
From 20 to 30: worker 1 crosses the bridge to the left.
From 30 to 40: worker 1 puts a box at the left warehouse.
From 30 to 31: worker 0 crosses the bridge to the left.
From 31 to 39: worker 0 puts a box at the left warehouse.
From 39 to 40: worker 0 crosses the bridge to the right.
From 40 to 49: worker 0 picks up a box from the right warehouse.
From 49 to 50: worker 0 crosses the bridge to the left.
</pre>
</div>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= n, k &lt;= 10<sup>4</sup></code></li>
	<li><code>time.length == k</code></li>
	<li><code>time[i].length == 4</code></li>
	<li><code>1 &lt;= leftToRight<sub>i</sub>, pickOld<sub>i</sub>, rightToLeft<sub>i</sub>, putNew<sub>i</sub> &lt;= 1000</code></li>
</ul>
