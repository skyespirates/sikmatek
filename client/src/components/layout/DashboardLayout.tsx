import { Outlet } from "react-router-dom";
import Sidebar from "./Sidebar";
import Header from "./Header";
import { Toaster } from "@/components/ui/sonner";

const DashboardLayout = () => {
  return (
    <div className="min-h-screen bg-slate-50/50">
      <Sidebar className={""} />
      <div className="flex flex-col min-h-screen">
        <Header />
        <main className="flex-1 p-6 md:p-8 md:pl-[18rem] pt-6 animate-fade-in">
          <div className="max-w-7xl mx-auto w-full space-y-8">
            <Outlet />
          </div>
        </main>
        <Toaster />
      </div>
    </div>
  );
};

export default DashboardLayout;
