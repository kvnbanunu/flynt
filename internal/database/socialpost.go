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
}

func (db *DB) GetAllPosts() ([]FullPost, error) {
	query := `SELECT s.id, s.user_id, s.fyre_id, s.type, s.content,
	u.username, u.img_url, f.title, f.streak_count FROM social_post s
	JOIN user u ON s.user_id = u.id
	JOIN fyre f ON s.fyre_id = f.id
	ORDER BY s.id ASC
	`

	var posts []FullPost
	err := db.Select(&posts, query)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
