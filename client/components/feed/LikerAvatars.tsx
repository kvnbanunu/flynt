"use client";

import React from "react";
import { Avatar, AvatarFallback, AvatarImage } from "../ui/avatar";

export const LikerAvatars: React.FC<{ likes: number }> = ({ likes }) => {
  likes = Math.min(likes, 3);
  const img_url = "/default_profile.jpg";
  const arr: number[] = [];
  for (let i = 0; i < likes; i++) {
    arr.push(i);
  }

  return (
    <div className="*:data-[slot=avatar]:ring-background flex -space-x-2 *:data-[slot=avatar]:ring-2">
      {arr.map((index) => (
        <Avatar key={index}>
          <AvatarImage src={img_url} alt={`liker${index}`} />
          <AvatarFallback>{`liker${index}`}</AvatarFallback>
        </Avatar>
      ))}
    </div>
  );
};
