深分页（deep pagination）在 MySQL 的应对策略

Keyset Pagination（也叫 Seek / 游标分页）
原理：不使用 OFFSET，而用已知的“最后一条记录的排序键”做 WHERE 过滤（比如自增主键、时间戳或组合索引），做索引查找，效率非常高。

示例（倒序按 created_at + id 唯一确定顺序）

```sql
-- 首次请求（第一页）
SELECT id, title, created_at
FROM posts
WHERE status = 'published'
ORDER BY created_at DESC, id DESC
LIMIT 20;

-- 之后请求（把上次最后一条的 (created_at_last, id_last) 作为游标）
SELECT id, title, created_at
FROM posts
WHERE status = 'published'
  AND (created_at < '2025-11-11 12:34:56'
       OR (created_at = '2025-11-11 12:34:56' AND id < 12345))
ORDER BY created_at DESC, id DESC
LIMIT 20;
```
也可用行比较（MySQL 支持）：
```sql
... AND (created_at, id) < ('2025-11-11 12:34:56', 12345)
```

优点：速度快、稳定，适合实时滚动（如“下一页/上一页”或无限滚动）。
缺点：无法直接跳到任意第 N 页（比如用户输入页码 10000），需要“游标 / token”支持或预计算索引。

实践建议：
	•	对排序字段建立覆盖索引（如 INDEX(status, created_at, id)），避免回表。
	•	用 opaque cursor（例如 base64 编码的 created_at|id）在前后端传递，而不是明文 id。

https://www.cnblogs.com/pDJJq/p/18876878/how-to-do-pagination-z1bs5pm

keyset(seek) pagination vs offset pagination
Keyset 就是指前一页最后一条的 (created_at, id), 用这一组 key 来决定下一页从哪里继续。
Seek 表示“从某个位置开始查找”



  
