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

export async function quoteKontrak(kontrak_id: string) {
  const { data } = await api.patch(`/api/v1/kontrak/${kontrak_id}/quote`);

  return data;
}

export async function confirmKontrak(kontrak_id: string) {
  const { data } = await api.patch(`/api/v1/kontrak/${kontrak_id}/confirm`);

  return data;
}

export async function cancelKontrak(kontrak_id: string) {
  const { data } = await api.patch(`/api/v1/kontrak/${kontrak_id}/cancel`);

  return data;
}

export async function activateKontrak(kontrak_id: string) {
  const { data } = await api.patch(`/api/v1/kontrak/${kontrak_id}/activate`);

  return data;
}
