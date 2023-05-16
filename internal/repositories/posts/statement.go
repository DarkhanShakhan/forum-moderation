package posts

const (
	getPostByIDStmt = `SELECT 
							p.id, 
							p.title,
							p.content,
							p.author_id,
							SUM(CASE WHEN r.reaction = true THEN 1 ELSE 0 END) count_likes,
							SUM(CASE WHEN r.reaction = false THEN 1 ELSE 0 END) count_dislikes
						FROM posts p 
						LEFT JOIN posts_reactions r ON p.id=r.post_id 
						WHERE p.id = ? AND p.visible = true AND p.deleted_at = NULL
						GROUP BY p.id, p.title, p.content, p.author_id`
	getPostsStmt = `SELECT
						p.id,
						p.title,
						p.content,
						p.author_id,
						SUM(CASE WHEN r.reaction = true THEN 1 ELSE 0 END) count_likes,
						SUM(CASE WHEN r.reaction = false THEN 1 ELSE 0 END) count_dislikes
					FROM posts p 
					LEFT JOIN posts_reactions r ON p.id=r.post_id
					WHERE p.visible = true AND p.deleted_at = NULL
					GROUP BY p.id, p.title, p.content, p.author_id
					ORDER BY p.created_at DESC`
	getPostsByCategoryStmt = `SELECT
								p.id,
								p.title,
								p.content,
								p.author_id,
								SUM(CASE WHEN r.reaction = true THEN 1 ELSE 0 END) count_likes,
								SUM(CASE WHEN r.reaction = false THEN 1 ELSE 0 END) count_dislikes
							FROM posts p 
							LEFT JOIN posts_reactions r ON p.id=r.post_id
							WHERE p.id in (
								SELECT post_id
								FROM posts_categories
								WHERE category_id = ?
							) AND p.visible = true AND p.deleted_at = NULL
							GROUP BY p.id, p.title, p.content, p.author_id
							ORDER BY p.created_at DESC`
	getPostsByAuthorIDStmt = `SELECT
							p.id,
							p.title,
							p.content,
							p.author_id,
							SUM(CASE WHEN r.reaction = true THEN 1 ELSE 0 END) count_likes,
							SUM(CASE WHEN r.reaction = false THEN 1 ELSE 0 END) count_dislikes
						FROM posts p 
						LEFT JOIN posts_reactions r ON p.id=r.post_id
						WHERE p.author_id = ? AND p.visible = true AND p.deleted_at = NULL
						GROUP BY p.id, p.title, p.content, p.author_id
						ORDER BY p.created_at DESC`
	createPostStmt = `INSERT INTO posts(title, content, author_id, created_at)
						VALUES (?, ?, ?, ?)
						`
	setVisibleStmt = `UPDATE posts SET visible = ? WHERE id = ?`
	deletePostStmt = `UPDATE posts SET deleted_at = ?, deleted_by = ?, delete_message=?, delete_category=? WHERE id = ?`
)
