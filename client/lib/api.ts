import HttpStatusCode from "@/types/status";
import { ApiError, ApiResponse, Result } from "@/types/api";

// Generic GET call
export async function Get<T>(url: string): Promise<Result<T, ApiError>> {
  try {
    const res = await fetch(url);

    if (!res.ok) {
      const errData: ApiError = await res.json();
      return { success: false, error: errData };
    }

    const successData: ApiResponse<T> = await res.json();
    return { success: true, data: successData.data };
  } catch (error: unknown) {
    const message = error instanceof Error ? error.message : "An unknown error occurred";
    return {
      success: false,
      error: {
        status_code: "" + HttpStatusCode.INTERNAL_SERVER_ERROR,
        message: message,
      },
    };
  }
}

// Generic PUT/update call
export async function Put<T>(
  url: string,
  content: T,
): Promise<Result<T, ApiError>> {
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

    const successData: ApiResponse<T> = await res.json();
    return { success: true, data: successData.data };
  } catch (error: unknown) {
    const message = error instanceof Error ? error.message : "An unknown error occurred";
    return {
      success: false,
      error: {
        status_code: "" + HttpStatusCode.INTERNAL_SERVER_ERROR,
        message: message,
      },
    };
  }
}

// Generic Post call
export async function Post<T, U>(
  url: string,
  content: U,
): Promise<Result<T, ApiError>> {
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

    const successData: ApiResponse<T> = await res.json();
    return { success: true, data: successData.data };
  } catch (error: unknown) {
    const message = error instanceof Error ? error.message : "An unknown error occurred";
    return {
      success: false,
      error: {
        status_code: "" + HttpStatusCode.INTERNAL_SERVER_ERROR,
        message: message,
      },
    };
  }
}

// Generic DELETE call using query params
export async function Delete(url: string): Promise<Result<null, ApiError>> {
  try {
    const res = await fetch(url, {
      method: "DELETE",
    });

    if (!res.ok) {
      const errData: ApiError = await res.json();
      return { success: false, error: errData };
    }
    return { success: true, data: null };
  } catch (error: unknown) {
    const message = error instanceof Error ? error.message : "An unknown error occurred";
    return {
      success: false,
      error: {
        status_code: "" + HttpStatusCode.INTERNAL_SERVER_ERROR,
        message: message,
      },
    };
  }
}

// Generic DELETE call with body
export async function DeleteBody<T>(
  url: string,
  content: T,
): Promise<Result<null, ApiError>> {
  try {
    const res = await fetch(url, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(content),
    });

    if (!res.ok) {
      const errData: ApiError = await res.json();
      return { success: false, error: errData };
    }
    return { success: true, data: null };
  } catch (error: unknown) {
    const message = error instanceof Error ? error.message : "An unknown error occurred";
    return {
      success: false,
      error: {
        status_code: "" + HttpStatusCode.INTERNAL_SERVER_ERROR,
        message: message,
      },
    };
  }
}
