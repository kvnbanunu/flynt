// Place any api related types here

// Generic response
export interface ApiResponse<T> {
  data: T;
  message: string;
}
