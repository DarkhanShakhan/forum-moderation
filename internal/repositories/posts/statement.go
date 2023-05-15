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
						WHERE p.id = ?
						GROUP BY p.id, p.title, p.content, p.author_id`
)
