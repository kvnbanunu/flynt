package database

type FullPost struct {
	ID          int      `db:"id" json:"id"`
	UserID      int      `db:"user_id" json:"user_id"`
	FyreID      int      `db:"fyre_id" json:"fyre_id"`
	Type        PostType `db:"type" json:"type"`
	Content     string   `db:"content" json:"content"`
	Username    string   `db:"username" json:"username"`
	ImgURL      *string  `db:"img_url" json:"img_url,omitempty"`
	Title       string   `db:"title" json:"title"`
	StreakCount int      `db:"streak_count" json:"streak_count"`
	Likes       int      `db:"likes" json:"likes"`
}

func (db *DB) GetAllPosts(id int) ([]FullPost, error) {
	query := `SELECT s.id, s.user_id, s.fyre_id, s.type, s.content,
	u.username, u.img_url, f.title, f.streak_count, f.likes, friend.statuts
	FROM social_post s
	JOIN user u ON s.user_id = u.id
	JOIN fyre f ON s.fyre_id = f.id
	JOIN friend ON friend.user_id_1 = ? AND friend.user_id_2 = u.id
	ORDER BY s.id ASC
	`

	var posts []FullPost
	err := db.Select(&posts, query, id)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (db *DB) LikePost(id int) (error) {
	query := `UPDATE social_post
	SET likes = likes + 1
	WHERE id = ?
	RETURNING *
	`

	var post SocialPost
	err := db.Get(&post, query, id)
	if err != nil {
		return err
	}

	query = `UPDATE fyre
	SET likes = ?
	WHERE id = ?
	RETURNING *
	`

	var fyre Fyre
	err = db.Get(&fyre, query, post.Likes, post.FyreID)
	if err != nil {
		return err
	}

	return nil
}
