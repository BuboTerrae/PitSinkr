import AppSidebar from "@/components/AppSidebar";
import { SidebarProvider, SidebarTrigger } from "@/components/ui/sidebar";
import Cookies from "js-cookie";
export default async function Layout({
  children,
}: {
  children: React.ReactNode;
}) {
  const defaultOpen = Cookies.get("sidebar_state") === "true";
  return (
    <>
      <SidebarProvider defaultOpen={defaultOpen}>
        <AppSidebar />
        <main className="flex flex-col flex-1 ">
          <SidebarTrigger />
          {children}
        </main>
      </SidebarProvider>
    </>
  );
}
