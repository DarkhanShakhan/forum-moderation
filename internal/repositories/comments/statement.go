package comments

const (
	createCommentStmt       = `INSERT INTO comments(post_id, author_id, content, created_at) VALUES (?, ?, ?, ?);`
	getCommentsByPostIDStmt = `SELECT
								c.id,
								c.post_id,
								c.content,
								c.author_id,
								SUM(CASE WHEN r.reaction = true THEN 1 ELSE 0 END) count_likes,
								SUM(CASE WHEN r.reaction = false THEN 1 ELSE 0 END) count_dislikes
							FROM comments c 
							LEFT JOIN comments_reactions r ON c.id=r.post_id
							WHERE c.visible = true AND c.deleted_at = NULL AND c.post_id = ?
							GROUP BY c.id, c.post_id, c.title, c.content, c.author_id
							ORDER BY c.created_at;`
	setVisibleStmt    = `UPDATE comments SET visible = ? WHERE id = ?;`
	deleteCommentStmt = `UPDATE comments SET deleted_at = ?, deleted_by = ?, delete_message=?, delete_category=? WHERE id = ?;`
)
