import { api } from "@/lib/axios";

export async function getLimitList() {
  const { data } = await api.get("/api/v1/limits");
  return data.data.limit;
}

export async function createLimit(requested_limit: number) {
  const { data } = await api.post("/api/v1/pengajuan-limit", {
    requested_limit: requested_limit,
  });
  return data;
}

export async function getApprovedLimits() {
  const { data } = await api.get("/api/v1/limits/approved");
  return data.limits;
}

export async function approveLimit(limit_id: number) {
  const { data } = await api.patch(
    `/api/v1/pengajuan-limit/${limit_id}/approve`
  );
  return data;
}

export async function rejectLimit(limit_id: number) {
  const { data } = await api.patch(
    `/api/v1/pengajuan-limit/${limit_id}/reject`
  );
  return data;
}
