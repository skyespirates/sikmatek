import { api } from "@/lib/axios";

export async function generateKey() {
  const response = await api.post("/generate-key");
  return response.data;
}

export type Payload = {
  key: string;
  text: string;
};

export async function encryptText(payload: Payload) {
  const response = await api.post("/encrypt", {
    key: payload.key,
    text: payload.text,
  });
  return response.data;
}

export async function decryptText(payload: Payload) {
  const response = await api.post("/decrypt", {
    key: payload.key,
    text: payload.text,
  });
  return response.data;
}
