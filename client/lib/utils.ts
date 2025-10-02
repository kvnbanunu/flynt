import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";
import "dotenv/config";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export const SS_ENV = {
  env: process.env.ENVIRONMENT,
  api_url: process.env.API_URL,
};

export const CS_ENV = {
  env: process.env.NEXT_PUBLIC_ENVIRONMENT,
  api_url: process.env.NEXT_PUBLIC_API_URL,
}
