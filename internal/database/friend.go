package database

import "fmt"

type UpdateFriendRequest struct {
	ID1 int `db:"user_id_1" json:"user_id_1"`
	ID2 int `db:"user_id_2" json:"user_id_2"`
}

type FriendsListItem struct {
	ID       int          `db:"id" json:"id"`
	Username string       `db:"username" json:"username"`
	ImgURL   *string      `db:"img_url" json:"img_url"`
	Status   FriendStatus `db:"status" json:"status"`
}

func (db *DB) AddFriend(req UpdateFriendRequest) error {
	query := `
	INSERT INTO friend (user_id_1, user_id_2, status)
	VALUES
	(?, ?, 'sent'),
	(?, ?, 'pending')
	ON CONFLICT (user_id_1, user_id_2) DO UPDATE
	SET status = EXCLUDED.status
	`

	result, err := db.Exec(query, req.ID1, req.ID2, req.ID2, req.ID1)
	if err != nil {
		return fmt.Errorf("Error creating friend request: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Error parsing rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("Error creating friend request: %w", err)
	}

	return nil
}

func (db *DB) AcceptFriend(req UpdateFriendRequest) error {
	query := `
	UPDATE friend
	SET status = 'friends'
	WHERE (user_id_1 = ? AND user_id_2 = ?)
	OR (user_id_1 = ? AND user_id_2 = ?)
	`

	result, err := db.Exec(query, req.ID1, req.ID2, req.ID2, req.ID1)
	if err != nil {
		return fmt.Errorf("Error accepting friend request: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Error parsing rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("Error accepting friend request: %w", err)
	}

	return nil
}

func (db *DB) BlockFriend(req UpdateFriendRequest) error {
	query := `
	INSERT INTO friend (user_id_1, user_id_2, status)
	VALUES (?, ?, 'blocked')
	ON CONFLICT (user_id_1, user_id_2) DO UPDATE
	SET status = 'blocked'
	`

	result, err := db.Exec(query, req.ID1, req.ID2)
	if err != nil {
		return fmt.Errorf("Error blocking friend: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Error parsing rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("Error blocking friend: %w", err)
	}

	return nil
}

func (db *DB) DeleteFriend(req UpdateFriendRequest) error {
	query := `DELETE FROM friend
	WHERE (user_id_1 = ? AND user_id_2 = ?)
	OR (user_id_1 = ? AND user_id_2 = ?)
	`

	result, err := db.Exec(query, req.ID1, req.ID2, req.ID2, req.ID1)
	if err != nil {
		return fmt.Errorf("Error deleting friend: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Error parsing rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("Error deleting friend: %w", err)
	}

	return nil
}

func (db *DB) GetFriendsList(id int) ([]FriendsListItem, error) {
	query := `
	SELECT u.id, u.username, u.img_url, f.status FROM friend f
	JOIN user u ON u.id = f.user_id_2
	WHERE f.user_id_1 = ?
	ORDER BY u.name ASC
	`

	var friends []FriendsListItem
	err := db.Select(&friends, query, id)
	if err != nil {
		return nil, fmt.Errorf("Failed to get friends list: %w", err)
	}

	return friends, nil
}
