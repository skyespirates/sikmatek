// lib/axios.ts
import axios from "axios";

export const api = axios.create({
  headers: {
    "Content-Type": "application/json",
  },
});

// Optional: add interceptors for auth tokens, logging, etc.
api.interceptors.request.use((config) => {
  // e.g. attach token
  // config.headers.Authorization = `Bearer ${token}`
  if (config.url?.includes("/login") || config.url?.includes("/register")) {
    return config;
  }

  const token = localStorage.getItem("token");
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});
