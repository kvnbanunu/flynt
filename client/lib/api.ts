import HttpStatusCode from "@/types/status";
import { ApiError, ApiResponse, Result } from "@/types/api";

// Generic GET call
export async function Get<T>(
  url: string,
): Promise<Result<ApiResponse<T>, ApiError>> {
  try {
    const res = await fetch(url);

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
        status_code: "" + HttpStatusCode.INTERNAL_SERVER_ERROR,
        message: error.message,
      },
    };
  }
}

// Generic PUT/update call
export async function Put<T>(
  url: string,
  content: T,
): Promise<Result<ApiResponse<T>, ApiError>> {
  try {
    const res = await fetch(url, {
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
        status_code: "" + HttpStatusCode.INTERNAL_SERVER_ERROR,
        message: error.message,
      },
    };
  }
}

// Generic Post call
export async function Post<T>(
  url: string,
  content: any,
): Promise<Result<ApiResponse<T>, ApiError>> {
  try {
    const res = await fetch(url, {
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

    const data: ApiResponse<T> = await res.json();
    return { success: true, data };
  } catch (error: any) {
    return {
      success: false,
      error: {
        status_code: "" + HttpStatusCode.INTERNAL_SERVER_ERROR,
        message: error.message,
      },
    };
  }
}
