import { jwtDecode } from "jwt-decode";
export function isAuthenticated() {
  const token = localStorage.getItem("token");
  if (token === null || token === "") {
    return false;
  }

  const decoded = jwtDecode(token);

  if (decoded.exp === undefined) {
    return false;
  }

  return decoded.exp > Math.floor(Date.now() / 1000);
}
