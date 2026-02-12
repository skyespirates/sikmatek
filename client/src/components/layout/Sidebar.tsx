import { NavLink, Link } from "react-router-dom";
import {
  LayoutDashboard,
  Gauge,
  FileText,
  Menu,
  X,
  Rocket,
} from "lucide-react";
import { Button } from "@/components/ui/button";
import { Sheet, SheetContent, SheetTrigger } from "@/components/ui/sheet";
import { cn } from "@/lib/utils";

const Sidebar = ({ className }: { className: string }) => {
  const navItems = [
    { name: "Dashboard", path: "/", icon: LayoutDashboard },
    { name: "Limits", path: "/limits", icon: Gauge },
    { name: "Contracts", path: "/contracts", icon: FileText },
  ];

  const NavContent = () => (
    <div className="flex flex-col h-full py-6">
      <div className="px-6 mb-8 flex items-center gap-2">
        <div className="h-8 w-8 rounded-lg bg-primary flex items-center justify-center">
          <Rocket className="h-5 w-5 text-white" />
        </div>
        <span className="text-xl font-heading font-bold bg-clip-text text-transparent bg-gradient-to-r from-primary to-blue-600">
          Krediyam
        </span>
      </div>

      <nav className="flex-1 px-4 space-y-2">
        {navItems.map((item) => (
          <NavLink
            key={item.path}
            to={item.path}
            className={({ isActive }) =>
              cn(
                "flex items-center gap-3 px-4 py-3 rounded-lg text-sm font-medium transition-all duration-200",
                isActive
                  ? "bg-primary text-white shadow-lg shadow-primary/25"
                  : "text-muted-foreground hover:bg-slate-100 hover:text-slate-900"
              )
            }
          >
            <item.icon className="h-5 w-5" />
            {item.name}
          </NavLink>
        ))}
      </nav>

      <div className="px-6 mt-auto">
        <div className="bg-slate-50 rounded-xl p-4 border border-slate-100">
          <h4 className="font-semibold text-sm mb-1">Upgrade Plan</h4>
          <p className="text-xs text-muted-foreground mb-3">
            Unlock more limits.
          </p>
          <Button size="sm" className="w-full text-xs" variant="outline">
            View Plans
          </Button>
        </div>
      </div>
    </div>
  );

  return (
    <>
      {/* Desktop Sidebar */}
      <aside
        className={cn(
          "hidden md:flex flex-col w-64 fixed inset-y-0 left-0 z-40 bg-white/70 backdrop-blur-xl border-r border-slate-200/60 shadow-[4px_0_24px_rgba(0,0,0,0.02)]",
          className
        )}
      >
        <NavContent />
      </aside>

      {/* Mobile Sidebar */}
      <Sheet>
        <SheetTrigger asChild>
          <Button
            variant="ghost"
            size="icon"
            className="md:hidden fixed top-4 left-4 z-50"
          >
            <Menu className="h-6 w-6" />
          </Button>
        </SheetTrigger>
        <SheetContent
          side="left"
          className="p-0 w-64 border-r-0 bg-white/95 backdrop-blur-xl"
        >
          <NavContent />
        </SheetContent>
      </Sheet>
    </>
  );
};

export default Sidebar;
