import { isAuthenticated } from "@/lib/auth";
import { Navigate } from "react-router";
import DashboardLayout from "./layout/DashboardLayout";

const ProtectedRoute = () => {
  if (!isAuthenticated()) {
    return <Navigate to="/login" replace />;
  }
  return <DashboardLayout />;
};

export default ProtectedRoute;
