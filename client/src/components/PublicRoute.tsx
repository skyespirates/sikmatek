import { isAuthenticated } from "@/lib/auth";
import { Navigate, Outlet } from "react-router";

const PublicRoute = () => {
  if (isAuthenticated()) {
    return <Navigate to="/" />;
  }
  return <Outlet />;
};

export default PublicRoute;
