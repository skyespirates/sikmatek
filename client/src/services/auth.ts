import { api } from "@/lib/axios";

export async function Login(data: { email: string; password: string }) {
  const res = await api.post("/v1/auth/login", data);
  return res.data;
}

export async function Register(data: { email: string; password: string }) {
  const res = await api.post("/v1/auth/register", data);
  return res.data;
}
