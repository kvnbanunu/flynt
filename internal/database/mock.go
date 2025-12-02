package database

type MockDB struct {
	// init
	InitDBFunc func(path string) error
	CloseFunc  func() error

	// health
	CheckHealthFunc func() error

	// seed
	SeedDataFunc func() error

	// account
	ValidateLoginFunc func(req AccountLoginRequest) (*User, error)

	// bonfyre
	GetBonfyreFunc        func(id int) (*Bonfyre, error)
	GetBonfyreMembersFunc func(bonfyreID int) ([]BonfyreMember, error)
	JoinBonfyreFunc       func(req BonfyreRequest, userID int) error

	// friend
	AddFriendFunc         func(req UpdateFriendRequest) error
	AcceptFriendFunc      func(req UpdateFriendRequest) error
	BlockFriendFunc       func(req UpdateFriendRequest) error
	DeleteFriendFunc      func(req UpdateFriendRequest) error
	GetFriendsListFunc    func(id int) ([]FriendsListItem, error)
	GetNonFriendsListFunc func(id int) ([]User, error)

	// fyre
	GetAllCategoriesFunc   func() ([]Category, error)
	CreateFyreFunc         func(req CreateFyreRequest, id int) (*Fyre, error)
	GetFyreByIDFunc        func(id int) (*Fyre, error)
	MapFyreGoalsFunc       func(fyres []Fyre, userID int) ([]FullFyre, error)
	GetAllUserFyresFunc    func(id int) ([]Fyre, error)
	GetAllFriendsFyresFunc func(id int) ([]Fyre, error)
	UpdateFyreFunc         func(id int, req UpdateFyreRequest) (*Fyre, error)
	ResetChecksFunc        func(ids []int) ([]Fyre, error)
	CheckFyreFunc          func(req CheckFyreRequest, id int) (*Fyre, error)
	GetFyreTotalFunc       func(id int) (*FyreTotalResponse, error)
	UpdateFyreTotalFunc    func(id, amount int, increment bool) (*FyreTotalResponse, error)
	DeleteFyreFunc         func(id int) error

	// goal
	CreateGoalFunc  func(req CreateGoalRequest) (*Goal, error)
	GetGoalByIDFunc func(fyreID int) (*Goal, error)
	UpdateGoalFunc  func(fyreID int, req UpdateGoalRequest) (*Goal, error)
	DeleteGoalFunc  func(fyreID int) error

	// socialpost
	GetAllPostsFunc func(id int) ([]FullPost, error)
	LikePostFunc    func(id int) error

	// user
	CreateUserFunc          func(req CreateUserRequest) (*User, error)
	GetUserByIDFunc         func(id int) (*User, error)
	GetUserByEmailFunc      func(email string) (*User, error)
	GetAllUsersFunc         func() ([]User, error)
	GetAllUsersRedactedFunc func(id int) ([]User, error)
	UpdateUserFunc          func(id int, req UpdateUserRequest) (*User, error)
	DeleteUserFunc          func(id int) error
}

var _ DBInterface = (*MockDB)(nil)

func NewMockDB() *MockDB {
	return &MockDB{}
}

// init
func (db *MockDB) InitDB(path string) error {
	if db.InitDBFunc != nil {
		return db.InitDBFunc(path)
	}
	return nil
}

func (db *MockDB) Close() error {
	if db.CloseFunc != nil {
		return db.CloseFunc()
	}
	return nil
}

// health
func (db *MockDB) CheckHealth() error {
	if db.CheckHealthFunc != nil {
		return db.CheckHealthFunc()
	}
	return nil
}

// seed
func (db *MockDB) SeedData() error {
	if db.SeedDataFunc != nil {
		return db.SeedDataFunc()
	}
	return nil
}

// account
func (db *MockDB) ValidateLogin(req AccountLoginRequest) (*User, error) {
	if db.ValidateLoginFunc != nil {
		return db.ValidateLoginFunc(req)
	}
	return nil, nil
}

// bonfyre
func (db *MockDB) GetBonfyre(id int) (*Bonfyre, error) {
	if db.GetBonfyreFunc != nil {
		return db.GetBonfyreFunc(id)
	}
	return nil, nil
}

func (db *MockDB) GetBonfyreMembers(bonfyreID int) ([]BonfyreMember, error) {
	if db.GetBonfyreMembersFunc != nil {
		return db.GetBonfyreMembersFunc(bonfyreID)
	}
	return nil, nil
}

func (db *MockDB) JoinBonfyre(req BonfyreRequest, userID int) error {
	if db.JoinBonfyreFunc != nil {
		return db.JoinBonfyreFunc(req, userID)
	}
	return nil
}

// friend
func (db *MockDB) AddFriend(req UpdateFriendRequest) error {
	if db.AddFriendFunc != nil {
		return db.AddFriendFunc(req)
	}
	return nil
}

func (db *MockDB) AcceptFriend(req UpdateFriendRequest) error {
	if db.AcceptFriendFunc != nil {
		return db.AcceptFriendFunc(req)
	}
	return nil
}

func (db *MockDB) BlockFriend(req UpdateFriendRequest) error {
	if db.BlockFriendFunc != nil {
		return db.BlockFriendFunc(req)
	}
	return nil
}

func (db *MockDB) DeleteFriend(req UpdateFriendRequest) error {
	if db.DeleteFriendFunc != nil {
		return db.DeleteFriendFunc(req)
	}
	return nil
}

func (db *MockDB) GetFriendsList(id int) ([]FriendsListItem, error) {
	if db.GetFriendsListFunc != nil {
		return db.GetFriendsListFunc(id)
	}
	return nil, nil
}

func (db *MockDB) GetNonFriendsList(id int) ([]User, error) {
	if db.GetNonFriendsListFunc != nil {
		return db.GetNonFriendsListFunc(id)
	}
	return nil, nil
}

// fyre
func (db *MockDB) GetAllCategories() ([]Category, error) {
	if db.GetAllCategoriesFunc != nil {
		return db.GetAllCategoriesFunc()
	}
	return nil, nil
}

func (db *MockDB) CreateFyre(req CreateFyreRequest, id int) (*Fyre, error) {
	if db.CreateFyreFunc != nil {
		return db.CreateFyreFunc(req, id)
	}
	return nil, nil
}

func (db *MockDB) GetFyreByID(id int) (*Fyre, error) {
	if db.GetFyreByIDFunc != nil {
		return db.GetFyreByIDFunc(id)
	}
	return nil, nil
}

func (db *MockDB) MapFyreGoals(fyres []Fyre, userID int) ([]FullFyre, error) {
	if db.MapFyreGoalsFunc != nil {
		return db.MapFyreGoalsFunc(fyres, userID)
	}
	return nil, nil
}

func (db *MockDB) GetAllUserFyres(id int) ([]Fyre, error) {
	if db.GetAllUserFyresFunc != nil {
		return db.GetAllUserFyresFunc(id)
	}
	return nil, nil
}

func (db *MockDB) GetAllFriendsFyres(id int) ([]Fyre, error) {
	if db.GetAllFriendsFyresFunc != nil {
		return db.GetAllFriendsFyresFunc(id)
	}
	return nil, nil
}

func (db *MockDB) UpdateFyre(id int, req UpdateFyreRequest) (*Fyre, error) {
	if db.UpdateFyreFunc != nil {
		return db.UpdateFyreFunc(id, req)
	}
	return nil, nil
}

func (db *MockDB) ResetChecks(ids []int) ([]Fyre, error) {
	if db.ResetChecksFunc != nil {
		return db.ResetChecksFunc(ids)
	}
	return nil, nil
}

func (db *MockDB) CheckFyre(req CheckFyreRequest, id int) (*Fyre, error) {
	if db.CheckFyreFunc != nil {
		return db.CheckFyreFunc(req, id)
	}
	return nil, nil
}

func (db *MockDB) GetFyreTotal(id int) (*FyreTotalResponse, error) {
	if db.GetFyreTotalFunc != nil {
		return db.GetFyreTotalFunc(id)
	}
	return nil, nil
}

func (db *MockDB) UpdateFyreTotal(id, amount int, increment bool) (*FyreTotalResponse, error) {
	if db.UpdateFyreTotalFunc != nil {
		return db.UpdateFyreTotalFunc(id, amount, increment)
	}
	return nil, nil
}

func (db *MockDB) DeleteFyre(id int) error {
	if db.DeleteFyreFunc != nil {
		return db.DeleteFyreFunc(id)
	}
	return nil
}

// goal
func (db *MockDB) CreateGoal(req CreateGoalRequest) (*Goal, error) {
	if db.CreateGoalFunc != nil {
		return db.CreateGoalFunc(req)
	}
	return nil, nil
}

func (db *MockDB) GetGoalByID(fyreID int) (*Goal, error) {
	if db.GetGoalByIDFunc != nil {
		return db.GetGoalByIDFunc(fyreID)
	}
	return nil, nil
}

func (db *MockDB) UpdateGoal(fyreID int, req UpdateGoalRequest) (*Goal, error) {
	if db.UpdateGoalFunc != nil {
		return db.UpdateGoalFunc(fyreID, req)
	}
	return nil, nil
}

func (db *MockDB) DeleteGoal(fyreID int) error {
	if db.DeleteGoalFunc != nil {
		return db.DeleteGoalFunc(fyreID)
	}
	return nil
}

// socialpost
func (db *MockDB) GetAllPosts(id int) ([]FullPost, error) {
	if db.GetAllPostsFunc != nil {
		return db.GetAllPostsFunc(id)
	}
	return nil, nil
}

func (db *MockDB) LikePost(id int) error {
	if db.LikePostFunc != nil {
		return db.LikePostFunc(id)
	}
	return nil
}

// user
func (db *MockDB) CreateUser(req CreateUserRequest) (*User, error) {
	if db.CreateUserFunc != nil {
		return db.CreateUserFunc(req)
	}
	return nil, nil
}

func (db *MockDB) GetUserByID(id int) (*User, error) {
	if db.GetUserByIDFunc != nil {
		return db.GetUserByIDFunc(id)
	}
	return nil, nil
}

func (db *MockDB) GetUserByEmail(email string) (*User, error) {
	if db.GetUserByEmailFunc != nil {
		return db.GetUserByEmailFunc(email)
	}
	return nil, nil
}

func (db *MockDB) GetAllUsers() ([]User, error) {
	if db.GetAllUsersFunc != nil {
		return db.GetAllUsersFunc()
	}
	return nil, nil
}

func (db *MockDB) GetAllUsersRedacted(id int) ([]User, error) {
	if db.GetAllUsersRedactedFunc != nil {
		return db.GetAllUsersRedactedFunc(id)
	}
	return nil, nil
}

func (db *MockDB) UpdateUser(id int, req UpdateUserRequest) (*User, error) {
	if db.UpdateUserFunc != nil {
		return db.UpdateUserFunc(id, req)
	}
	return nil, nil
}

func (db *MockDB) DeleteUser(id int) error {
	if db.DeleteUserFunc != nil {
		return db.DeleteUserFunc(id)
	}
	return nil
}
