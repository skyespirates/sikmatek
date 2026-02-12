import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { BrowserRouter, Routes, Route } from "react-router";
import "./index.css";
import Register from "@/pages/Register.tsx";
import Login from "@/pages/Login.tsx";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import ProtectedRoute from "./components/ProtectedRoute";
import PublicRoute from "./components/PublicRoute";
import EncryptDecrypt from "./pages/EncryptDecrypt";
import NotFound from "./pages/NotFound";
import Dashboard from "./pages/Dashboard";
import Limits from "@/pages/Limit";
import Contracts from "@/pages/Contract";

export const queryClient = new QueryClient();

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <QueryClientProvider client={queryClient}>
      <BrowserRouter>
        <Routes>
          <Route element={<ProtectedRoute />}>
            <Route index element={<Dashboard />} />
            <Route path="limits" element={<Limits />} />
            <Route path="contracts" element={<Contracts />} />
          </Route>
          <Route element={<PublicRoute />}>
            <Route path="register" element={<Register />} />
            <Route path="login" element={<Login />} />
          </Route>
          <Route path="labs" element={<EncryptDecrypt />} />
          <Route path="*" element={<NotFound />} />
        </Routes>
      </BrowserRouter>
    </QueryClientProvider>
  </StrictMode>
);
