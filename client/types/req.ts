export type LoginType = "username" | "email";

export interface LoginRequest {
  type: LoginType;
  username?: string;
  email?: string;
  password: string;
}

export interface RegisterRequest {
  username: string;
  name: string;
  password: string;
  email: string;
  timezone: string;
}

export interface UpdateUserRequest {
  username?: string;
  name?: string;
  current_password?: string;
  new_password?: string;
  email?: string;
  img_url?: string;
  bio?: string;
  timezone?: string;
}

export interface CreateFyreRequest {
  title: string;
  streak_count: number | 0;
  active_days: string | "1111111";
  category_id: number | 0;
}

export interface CheckFyreRequest {
  id: number; // fyre_id
  increment: boolean;
}

export interface UpdateFyreRequest {
  title?: string;
  streak_count?: number;
  bonfyre_id?: number;
  active_days?: string;
  is_private?: boolean;
  category_id?: number;
}

export interface JoinBonfyreRequest {
  fyre_id: number;
}

export interface CreateGoalRequest {
  fyre_id: number;
  description: string;
  goal_type_id: number;
  data: string;
}

export interface UpdateGoalRequest {
  description?: string;
  goal_type_id?: number;
  data?: string;
}

export interface FriendRequest {
  type: string;
  user_id_2: number;
}

export interface FriendsListItem {
  id: number;
  username: string;
  img_url?: string;
  status: string;
}

export interface FriendsUserListItem {
  id: number;
  username: string;
  name: string;
  img_url?: string;
  bio?: string;
}
export interface FullPost {
  id: number;
  user_id: number;
  fyre_id: number;
  type: "dailycheck" | "milestone";
  content: string;
  username: string;
  img_url?: string;
  title: string;
  streak_count: number;
  likes: number;
  status?: string;
}
