"use client";

import React from "react";
import { ScrollArea } from "../ui/scroll-area";
import { useAuth } from "@/contexts/AuthContext";
import { FyreCard } from "./FyreCard";
import { AddFyre } from "./AddFyre";
import { FyreTotal } from "../flame/FyreTotal";

export const FyreArea: React.FC = () => {
  const { fyres } = useAuth();
  return (
    <ScrollArea className="w-full">
      <div className="flex flex-col gap-4 w-full">
      <FyreTotal />
      {fyres &&
        fyres.map((fyre) => (
          <React.Fragment key={fyre.fyre.id}>
            <FyreCard fyre={fyre.fyre} goals={fyre.goals} />
          </React.Fragment>
        ))}
      <AddFyre />
      </div>
    </ScrollArea>
  );
};
