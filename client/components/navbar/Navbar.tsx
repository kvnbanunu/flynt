import Link from "next/link";
import {
    NavigationMenu,
    NavigationMenuList,
    NavigationMenuItem,
    NavigationMenuLink,
} from "@/components/ui/navigation-menu";

// menu items
const items = [
  {
    title: "HOME",
    url: "#",
  },
  {
    title: "FEED",
    url: "#",
  },
  {
    title: "FRIEND",
    url: "#",
  },
  {
    title: "PROFILE",
    url: "#",
  },
]

export function MobileNavbar() {
    return (

        <NavigationMenu>
            <NavigationMenuList className="flex md:hidden fixed bottom-0 left-0 z-50 w-full bg-background ">
                <div className="flex gap-8 rounded-2xl px-6 py-4 mb-5 bg-sidebar">


                    {items.map((item) => (
                        <NavigationMenuItem key={item.title}>
                            <NavigationMenuLink asChild>
                                <Link
                                    href={item.url}
                                    className="w-14 h-14 bg-ftrim text-center items-center justify-center"
                                >
                                    {item.title}
                                </Link>
                            </NavigationMenuLink>
                        </NavigationMenuItem>
                    ))}
                </div>
            </NavigationMenuList>
        </NavigationMenu>
    );
};