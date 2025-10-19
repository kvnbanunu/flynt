"use client";

import React from "react";
import { Card, CardAction, CardContent, CardHeader, CardTitle } from "../ui/card";
import { Button } from "../ui/button";

export const FyreCard: React.FC<{fyre: Models.Fyre}> = ({fyre}) => {

  return (
    <Card className="mb-4 pt-0">
      <CardHeader className="bg-primary text-primary-foreground rounded-t-xl">
        <CardTitle className="pt-3 pb-1">{fyre.title}</CardTitle>
        <CardAction className="pt-0.5">
          <Button variant="ghost">+</Button>
        </CardAction>
      </CardHeader>
      <CardContent>
        Streak: {fyre.streak_count} 🔥
      </CardContent>
    </Card>
  )
}
