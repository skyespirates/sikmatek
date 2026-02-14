import { jwtDecode } from "jwt-decode";
import type { JwtPayload } from "jwt-decode";

interface User extends JwtPayload {
  id: number;
  email: string;
  role_id: number;
  consumer_id: number;
  is_verified: boolean;
}

const tokenKey = "token";

export function isAuthenticated() {
  const token = localStorage.getItem(tokenKey);
  if (token === null || token === "") {
    return false;
  }

  const decoded = jwtDecode(token);

  if (decoded.exp === undefined) {
    return false;
  }

  return decoded.exp > Math.floor(Date.now() / 1000);
}

export function logout() {
  localStorage.removeItem(tokenKey);
  window.location.href = "/login";
}

export function getUser(): User | null {
  const token = localStorage.getItem(tokenKey);
  if (!token) {
    return null;
  }

  try {
    return jwtDecode<User>(token);
  } catch (error) {
    console.error("Invalid token:", error);
    return null;
  }
}
