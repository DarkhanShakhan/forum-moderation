package reactions

import (
	"context"
	"database/sql"

	"github.com/DarkhanShakhan/forum-moderation/internal/domain/entity"
)

type Repository interface {
	UpsertPostReaction(ctx context.Context, reaction *entity.Reaction) error
	// DeletePostReaction(ctx context.Context, reaction *entity.Reaction) error
	UpsertCommentReaction(ctx context.Context, reaction *entity.Reaction) error
	// DeleteCommentReaction(ctx context.Context, reaction *entity.Reaction) error
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) UpsertPostReaction(ctx context.Context, reaction *entity.Reaction) error {
	return r.upsertReaction(ctx, upsertPostReactionStmt, deletePostReactionStmt, reaction)
}

func (r *repository) UpsertCommentReaction(ctx context.Context, reaction *entity.Reaction) error {
	return r.upsertReaction(ctx, upsertCommentReactionStmt, deleteCommentReactionStmt, reaction)
}

func (r *repository) upsertReaction(ctx context.Context, upsertStmt, deleteStmt string, reaction *entity.Reaction) error {
	res, err := r.db.Exec(upsertStmt, reaction.EntityID, reaction.UserID, reaction.Like)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return r.deleteReaction(ctx, deleteStmt, reaction)
	}
	return nil
}
func (r *repository) deleteReaction(ctx context.Context, stmt string, reaction *entity.Reaction) error {
	_, err := r.db.Exec(stmt, reaction.EntityID, reaction.UserID)
	return err
}
