"use client";

import React from "react";
import { ScrollArea } from "../ui/scroll-area";
import { useAuth } from "@/contexts/AuthContext";
import { FyreCard } from "./FyreCard";
import { AddFyre } from "./AddFyre";

export const FyreArea: React.FC = () => {
  const { fyres } = useAuth();
  return (
    <ScrollArea className="h-full w-full">
      {fyres &&
        fyres.map((fyre) => (
          <React.Fragment key={fyre.id}>
            <FyreCard fyre={fyre} />
          </React.Fragment>
        ))}
      <AddFyre />
    </ScrollArea>
  );
};
