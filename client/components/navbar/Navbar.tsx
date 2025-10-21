"use client";
import Link from "next/link";
import {
  NavigationMenu,
  NavigationMenuList,
  NavigationMenuItem,
  NavigationMenuLink,
} from "@/components/ui/navigation-menu";
import React from "react";
import { Home, Inbox, Sticker, Users } from "lucide-react";
import { useAuth } from "@/contexts/AuthContext";
import { useIsMobile } from "@/hooks/use-mobile";

// menu items
const items = [
  {
    title: "HOME",
    url: "#",
    icon: Home,
  },
  {
    title: "FEED",
    url: "#",
    icon: Inbox,
  },
  {
    title: "FRIEND",
    url: "#",
    icon: Users,
  },
  {
    title: "PROFILE",
    url: "#",
    icon: Sticker,
  },
];

export const MobileNavbar: React.FC = () => {
  const isMobile = useIsMobile();
  const { isAuthenticated } = useAuth();

  if (!isAuthenticated) {
    return null;
  }

  return (
    <NavigationMenu
      viewport={isMobile}
      className="md:hidden sticky bottom-0 bg-background"
    >
      <NavigationMenuList className="bg-background w-screen">
        <div className="flex gap-8 rounded-2xl px-6 py-4 my-4 mx-4 bg-sidebar justify-evenly w-full">
        {items.map((item) => (
          <NavigationMenuItem key={item.title}>
            <NavigationMenuLink asChild className="rounded-full">
              <Link
                href={item.url}
                className="w-14 h-14 bg-ftrim text-center items-center justify-center"
              >
                <item.icon />
                <span className="sr-only">{item.title}</span>
              </Link>
            </NavigationMenuLink>
          </NavigationMenuItem>
        ))}
        </div>
      </NavigationMenuList>
    </NavigationMenu>
  );
};
