import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { BrowserRouter, Routes, Route } from "react-router";
import "./index.css";
import Homepage from "@/pages/Homepage.tsx";
import Register from "@/pages/Register.tsx";
import Login from "@/pages/Login.tsx";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import ProtectedRoute from "./components/ProtectedRoute";
import PublicRoute from "./components/PublicRoute";
import EncryptDecrypt from "./pages/EncryptDecrypt";

export const queryClient = new QueryClient();

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <QueryClientProvider client={queryClient}>
      <BrowserRouter>
        <Routes>
          <Route element={<ProtectedRoute />}>
            <Route index element={<Homepage />} />
          </Route>
          <Route element={<PublicRoute />}>
            <Route path="register" element={<Register />} />
            <Route path="login" element={<Login />} />
          </Route>
          <Route path="labs" element={<EncryptDecrypt />} />
        </Routes>
      </BrowserRouter>
    </QueryClientProvider>
  </StrictMode>
);
