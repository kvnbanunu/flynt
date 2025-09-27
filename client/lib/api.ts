import HttpStatusCode from "@/types/status";
import { ENV } from "./utils";
import { ApiError, ApiResponse, Result } from "@/types/api";

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
