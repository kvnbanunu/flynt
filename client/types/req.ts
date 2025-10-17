export interface LoginRequest {
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
