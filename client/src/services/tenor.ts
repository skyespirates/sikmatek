import { api } from "@/lib/axios";

export async function getTenors() {
  const response = await api.get("/api/v1/tenors");
  return response.data;
}
