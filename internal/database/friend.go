package database

import "fmt"

type UpdateFriendRequest struct {
	ID1 int `json:"user_id_1"`
	ID2 int `json:"user_id_2"`
}

type FriendsListItem struct {
	ID     int          `json:"id"`
	Name   string       `json:"name"`
	Status FriendStatus `json:"status"`
}

func (db *DB) AddFriend(req UpdateFriendRequest) error {
	query := `
	INSERT INTO friend (user_id_1, user_id_2, status)
	VALUES
	(?, ?, 'accepted'),
	(?, ?, 'pending')
	ON CONFLICT (user_id_1, user_id_2) DO UPDATE
	SET status = EXCLUDED.status;
	`

	err := db.Get(nil, query, req.ID1, req.ID2, req.ID2, req.ID1)
	if err != nil {
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

	err := db.Get(nil, query, req.ID1, req.ID2, req.ID2, req.ID1)
	if err != nil {
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

	err := db.Get(nil, query, req.ID1, req.ID2)
	if err != nil {
		return fmt.Errorf("Error blocking friend: %w", err)
	}

	return nil
}

func (db *DB) DeleteFriend(req UpdateFriendRequest) error {
	query := `DELETE FROM friend
	WHERE (user_id_1 = ? AND user_id_2 = ?)
	OR (user_id_1 = ? AND user_id_2 = ?)
	`

	err := db.Get(nil, query, req.ID1, req.ID2, req.ID2, req.ID1)
	if err != nil {
		return fmt.Errorf("Error deleting friend: %w", err)
	}

	return nil
}

func (db *DB) GetFriendsList(id int) ([]FriendsListItem, error) {
	query := `
	SELECT (u.id, u.name, f.status) FROM friend f
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
