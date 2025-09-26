import { ENV } from "./utils";

export async function Get(endpoint: string) {
  const data = await fetch(`${ENV.api_url}${endpoint}`);
  return data.json();
}
