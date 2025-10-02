import HttpStatusCode from "@/types/status";
import { ENV } from "./utils";
import { ApiError, ApiResponse, Result } from "@/types/api";

// Generic GET call
export async function Get<T>(
  endpoint: string,
): Promise<Result<ApiResponse<T>, ApiError>> {
  try {
    const res = await fetch(`${ENV.api_url}${endpoint}`);

    if (!res.ok) {
      const errData: ApiError = await res.json();
      return { success: false, error: errData };
    }

    const data: ApiResponse<T> = await res.json();
    return { success: true, data };
  } catch (error: any) {
    return {
      success: false,
      error: {
        statusCode: HttpStatusCode.INTERNAL_SERVER_ERROR,
        message: error.message,
      },
    };
  }
}

// Generic PUT/update call
export async function Put<T>(
  endpoint: string,
  content: T,
): Promise<Result<ApiResponse<T>, ApiError>> {
  try {
    const res = await fetch(`${ENV.api_url}${endpoint}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(content),
    });

    if (!res.ok) {
      const errData: ApiError = await res.json();
      return { success: false, error: errData };
    }

    const data: ApiResponse<T> = await res.json();
    return { success: true, data };
  } catch (error: any) {
    return {
      success: false,
      error: {
        statusCode: HttpStatusCode.INTERNAL_SERVER_ERROR,
        message: error.message,
      },
    };
  }
}

// Generic Post call
export async function Post<T, R>(
  endpoint: string,
  content: T,
): Promise<Result<ApiResponse<R>, ApiError>> {
  try {
    const res = await fetch(`${ENV.api_url}${endpoint}`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(content),
    });

    if (!res.ok) {
      const errData: ApiError = await res.json();
      return { success: false, error: errData };
    }

    const data: ApiResponse<R> = await res.json();
    return { success: true, data };
  } catch (error: any) {
    return {
      success: false,
      error: {
        statusCode: HttpStatusCode.INTERNAL_SERVER_ERROR,
        message: error.message,
      },
    };
  }
}
