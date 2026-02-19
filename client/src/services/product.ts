import { api } from "@/lib/axios";

export async function getProductList() {
  const { data } = await api.get("/api/v1/products");

  return data.products;
}
