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
  name?: string;
  password?: string;
  email?: string;
  img_url?: string;
  bio?: string;
  timezone?: string;
}

export interface CreateFyreRequest {
  title: string;
  streak_count: number | 0;
  active_days: string | "1111111";
}

export interface UpdateFyreRequest {
  title?: string;
  streak_count?: number;
  bonfyre_id?: number;
  active_days?: string;
}
