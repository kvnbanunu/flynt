package database

// all methods implemented by DB
type DBInterface interface {
	// init
	InitDB(path string) error
	Close() error

	// health
	CheckHealth() error

	// seed
	SeedData() error

	// account
	ValidateLogin(req AccountLoginRequest) (*User, error)

	// bonfyre
	GetBonfyre(id int) (*Bonfyre, error)
	GetBonfyreMembers(bonfyreID int) ([]BonfyreMember, error)
	JoinBonfyre(req BonfyreRequest, userID int) error

	// friend
	AddFriend(req UpdateFriendRequest) error
	AcceptFriend(req UpdateFriendRequest) error
	BlockFriend(req UpdateFriendRequest) error
	DeleteFriend(req UpdateFriendRequest) error
	GetFriendsList(id int) ([]FriendsListItem, error)
	GetNonFriendsList(id int) ([]User, error)

	// fyre
	GetAllCategories() ([]Category, error)
	CreateFyre(req CreateFyreRequest, id int) (*Fyre, error)
	GetFyreByID(id int) (*Fyre, error)
	MapFyreGoals(fyres []Fyre, userID int) ([]FullFyre, error)
	GetAllUserFyres(id int) ([]Fyre, error)
	GetAllFriendsFyres(id int) ([]Fyre, error)
	UpdateFyre(id int, req UpdateFyreRequest) (*Fyre, error)
	ResetChecks(ids []int) ([]Fyre, error)
	CheckFyre(req CheckFyreRequest, id int) (*Fyre, error)
	GetFyreTotal(id int) (*FyreTotalResponse, error)
	UpdateFyreTotal(id, amount int, increment bool) (*FyreTotalResponse, error)
	DeleteFyre(id int) error

	// goal
	CreateGoal(req CreateGoalRequest) (*Goal, error)
	GetGoalByID(fyreID int) (*Goal, error)
	UpdateGoal(fyreID int, req UpdateGoalRequest) (*Goal, error)
	DeleteGoal(fyreID int) error

	// socialpost
	GetAllPosts(id int) ([]FullPost, error)
	LikePost(id int) error

	// user
	CreateUser(req CreateUserRequest) (*User, error)
	GetUserByID(id int) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetAllUsers() ([]User, error)
	GetAllUsersRedacted(id int) ([]User, error)
	UpdateUser(id int, req UpdateUserRequest) (*User, error)
	DeleteUser(id int) error
}

var _ DBInterface = (*DB)(nil)
