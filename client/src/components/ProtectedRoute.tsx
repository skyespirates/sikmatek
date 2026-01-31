import { isAuthenticated } from "@/lib/auth";
import { Navigate, Outlet } from "react-router";

const ProtectedRoute = () => {
  if (!isAuthenticated()) {
    return <Navigate to="/login" replace />;
  }
  return <Outlet />;
};

export default ProtectedRoute;
