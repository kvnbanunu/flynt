import HttpStatusCode from "@/types/status";
import { ApiError, ApiResponse, Result } from "@/types/api";
import { CS_ENV, SS_ENV } from "./utils";


// Generic GET call
// <T> is the expected response type
export async function Get<T>(url: string): Promise<Result<T, ApiError>> {
  try {
    const res = await fetch((CS_ENV.api_url ?? SS_ENV.api_url) + url, {
      credentials: "include",
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

// Generic PUT/update call
// <T> is the type you are updating and receiving
export async function Put<T>(
  url: string,
  content: T,
): Promise<Result<T, ApiError>> {
  try {
    const res = await fetch((CS_ENV.api_url ?? SS_ENV.api_url) + url, {
      method: "PUT",
      credentials: "include",
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
// <T> is the expected response type
// <U> is the type of content you are sending
export async function Post<T, U>(
  url: string,
  content: U,
): Promise<Result<T, ApiError>> {
  try {
    const res = await fetch((CS_ENV.api_url ?? SS_ENV.api_url) + url, {
      method: "POST",
      credentials: "include",
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
    const res = await fetch((CS_ENV.api_url ?? SS_ENV.api_url) + url, {
      method: "DELETE",
      credentials: "include",
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
// <T> is the type you are deleting
export async function DeleteBody<T>(
  url: string,
  content: T,
): Promise<Result<null, ApiError>> {
  try {
    const res = await fetch((CS_ENV.api_url ?? SS_ENV.api_url) + url, {
      method: "DELETE",
      credentials: "include",
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
