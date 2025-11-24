"use client";

import { useAuth } from "@/contexts/AuthContext";
import React from "react";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "../ui/card";
import { FlameAnimation } from "./Flame";

export const FyreTotal: React.FC = () => {
  const { user } = useAuth();

  if (user) {
    return (
      <Card className="items-center justify-center">
        <CardHeader className="justify-center">
          <CardTitle className="text-5xl">Flynt</CardTitle>
        </CardHeader>
        <CardContent className="flex justify-center items-center">
          <FlameAnimation />
          {/* <div className="bg-red-500 w-[360px] h-[360px]"></div> */}
        </CardContent>
        <CardFooter className="text-3xl font-bold">
          {user.fyre_total}
        </CardFooter>
      </Card>
    );
  }
};
