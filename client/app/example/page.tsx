import { ExampleHome } from "@/components/examples/examplehome";

import { MainContainer } from "@/components/MainContainer";
import { MobileNavbar } from "@/components/navbar/Navbar";
import { DesktopSidebar } from "@/components/sidebar/Sidebar";
import { SidebarProvider, SidebarTrigger } from "@/components/ui/sidebar"

export default function Examples() {
  return (
    <div className="font-sans grid grid-rows-[20px_1fr_20px] items-center justify-items-center min-h-screen p-8 pb-20 gap-16 sm:p-20">
      
      <SidebarProvider>
        <DesktopSidebar />
        {/* <main className="flex flex-col gap-[32px] row-start-2 items-center sm:items-start"> */}
        <MainContainer>
          <ExampleHome />
          
        </MainContainer>
        <MobileNavbar />
        {/* </main> */}
      </SidebarProvider>
      
    </div>
  );
}
