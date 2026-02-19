import { api } from "@/lib/axios";

export async function getDaftarKontrak() {
  const { data } = await api.get("/api/v1/kontrak");

  return data.data.kontrak;
}

type KontrakData = {
  product_id: number;
  limit_id: number;
  tenor: number;
};

export async function buatKontrak(payload: KontrakData) {
  const { data } = await api.post("/api/v1/kontrak", payload);

  return data.data.nomor_kontrak;
}
