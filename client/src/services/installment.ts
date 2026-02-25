import { api } from "@/lib/axios";

export async function getInstallmentList(nomor_kontrak: string) {
  const { data } = await api.get(`/api/v1/installments/list/${nomor_kontrak}`);
  return data.data.installments;
}
