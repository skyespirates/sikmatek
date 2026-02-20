import { api } from "@/lib/axios";
import * as z from "zod";
import { formSchema } from "@/pages/Profile";
import { format } from "date-fns";

type Data = z.infer<typeof formSchema>;

export async function getProfile() {
  const { data } = await api.get("/api/v1/consumers");
  return data.data;
}

export async function updateProfile(payload: Data) {
  const { data } = await api.put("/api/v1/consumers", {
    ...payload,
    tanggal_lahir: format(payload.tanggal_lahir, "yyyy-MM-dd"),
  });

  return data;
}

export async function uploadKtp(ktp: FormData) {
  const { data } = await api.put("/api/v1/consumers/upload-ktp", ktp, {
    headers: { "Content-Type": "multipart/form-data" },
  });

  return data;
}

export async function uploadSelfie(selfie: FormData) {
  const { data } = await api.put("/api/v1/consumers/upload-selfie", selfie, {
    headers: { "Content-Type": "multipart/form-data" },
  });

  return data;
}
