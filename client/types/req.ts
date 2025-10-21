export type LoginType = "username" | "email"

export interface LoginRequest {
  type: LoginType;
  username?: string | null;
  email?: string | null;
  password: string;
}

export interface RegisterRequest {
  username: string;
  name: string;
  email: string;
  password: string;
}

export interface UpdateFyreRequest {
  title?: string;
  streak_count?: number;
  bonfyre_id?: number;
  active_days?: string;
}
