import { api } from "@/lib/axios";

export async function getDaftarKontrak() {
  const { data } = await api.get("/v1/kontrak");

  return data.data.kontrak;
}
