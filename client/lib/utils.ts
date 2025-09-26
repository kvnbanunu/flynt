import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";
import "dotenv/config";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export const ENV = {
  env: process.env.ENVIRONMENT,
  api_url: process.env.API_URL,
};
