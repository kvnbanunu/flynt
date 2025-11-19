"use client";

import { FlameAnimation } from "@/components/flame/Flame";
import { Button } from "@/components/ui/button";
import { motion } from "framer-motion";
import { usePathname, useRouter } from "next/navigation";

export default function LandingPage() {
  const router = useRouter();
  const pathname = usePathname();

  return (
    <div key={pathname} className="flex flex-col items-center justify-center">
      <motion.h1
        className="text-5xl font-bold mb-4"
        initial={{ y: -50, opacity: 0 }}
        animate={{ y: 0, opacity: 1 }}
        transition={{ delay: 0.2 }}
      >
        Flynt
      </motion.h1>

      <motion.p
        className="text-lg max-w-md"
        initial={{ opacity: 0 }}
        animate={{ opacity: 1 }}
        transition={{ delay: 0.3 }}
      >
        Ignite your fyre today
      </motion.p>

      <motion.div
        initial={{ filter: "blur(20px)", opacity: 0 }}
        animate={{ filter: "blur(0px)", opacity: 1 }}
        transition={{ duration: 0.5, ease: "easeOut" }}
      >
        <FlameAnimation />
      </motion.div>

      <motion.div
        className="flex space-x-4"
        initial={{ opacity: 0, y: 20 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ delay: 0.5 }}
      >
        <Button
          onClick={() => router.push("/login")}
          className=" px-15 py-10 my-7 rounded-xl text-2xl"
        >
          Login
        </Button>
      </motion.div>
    </div>
  );
}
