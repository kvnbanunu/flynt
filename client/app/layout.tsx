import type { Metadata } from "next";
import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";
import { AuthProvider } from "@/contexts/AuthContext";
import { Toaster } from "@/components/ui/sonner";
import { DesktopSidebar } from "@/components/sidebar/Sidebar";
import { MobileNavbar } from "@/components/navbar/Navbar";
import { SidebarProvider } from "@/components/ui/sidebar";
import { MainContainer } from "@/components/MainContainer";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "Flynt",
  description: "Use Flynt to start your next Fyre streak",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" className="dark">
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased`}
      >
        <AuthProvider>
          <SidebarProvider>
            <DesktopSidebar />
            <Toaster />
            <div className="flex flex-col md:flex-row w-screen md:w-full h-screen">
              <MainContainer>
                {children}
              </MainContainer>
              <MobileNavbar />
            </div>
          </SidebarProvider>
        </AuthProvider>
      </body>
    </html>
  );
}
