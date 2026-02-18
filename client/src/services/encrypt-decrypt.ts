import { api } from "@/lib/axios";

export async function generateKey() {
  const response = await api.post("/api/generate-key");
  return response.data;
}

export type Payload = {
  key: string;
  text: string;
};

export async function encryptText(payload: Payload) {
  const response = await api.post("/api/encrypt", {
    key: payload.key,
    text: payload.text,
  });
  return response.data;
}

export async function decryptText(payload: Payload) {
  const response = await api.post("/api/decrypt", {
    key: payload.key,
    text: payload.text,
  });
  return response.data;
}
