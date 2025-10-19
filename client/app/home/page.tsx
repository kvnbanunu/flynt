// app/page.tsx
"use client";

import { motion } from "framer-motion";
import { Button } from "@/components/ui/button";

export default function welcome() {
  return (
    <main className="min-h-screen flex flex-col items-center justify-center bg-gradient-to-b from-gray-50 to-gray-200 text-gray-900 p-6">
      <motion.h1
        initial={{ opacity: 0, y: -20 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.6 }}
        className="text-4xl md:text-6xl font-extrabold text-center mb-4"
      >
        Flynt
      </motion.h1>

      <motion.div
        initial={{ scale: 0.9, opacity: 0 }}
        animate={{ scale: 1, opacity: 1 }}
        transition={{ delay: 0.6 }}
        className="flex flex-col gap-1"
      >
        <Button
          onClick={() => alert("nothing here")}
          className="rounded-2xl text-lg"
        >
          Signup
        </Button>

        <Button
          onClick={() => alert("nothing here")}
          className="rounded-2xl text-lg"
        >
          Login
        </Button>
      </motion.div>

      <footer className="mt-12 text-sm text-gray-500">
        Flynting it
      </footer>
    </main>
  );
}
