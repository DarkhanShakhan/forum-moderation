package reactions

const (
	upsertPostReactionStmt = `INSERT INTO post_reactions(post_id, user_id, reaction) 
								VALUES(?, ?, ?)
								ON CONFLICT (post_id, user_ud)
								DO UPDATE SET reaction=excluded.reaction WHERE reaction != excluded.reaction`
	deletePostReactionStmt    = `DELETE FROM post_reactions WHERE post_id = ? AND user_id = ?`
	upsertCommentReactionStmt = `INSERT INTO comment_reactions(comment_id, user_id, reaction) 
								VALUES(?, ?, ?)
								ON CONFLICT (comment_id, user_ud)
								DO UPDATE SET reaction=excluded.reaction WHERE reaction != excluded.reaction`
	deleteCommentReactionStmt = `DELETE FROM comment_reactions WHERE comment_id = ? AND user_id = ?`
)
