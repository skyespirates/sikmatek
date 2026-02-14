import { api } from "@/lib/axios";

export async function getLimitList() {
  const { data } = await api.get("/v1/limits");
  return data.data.limit;
}

export async function createLimit(requested_limit: number) {
  const { data } = await api.post("/v1/pengajuan-limit", {
    requested_limit: requested_limit,
  });
  return data;
}
