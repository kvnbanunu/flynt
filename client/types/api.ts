// Place any api related types here

// Generic response
export interface ApiResponse<T> {
  data: T;
  message: string;
}

export interface ApiError {
  statusCode: number;
  message: string;
}

export type Result<T, E> = { success: true; data: T} | { success: false; error: E }
