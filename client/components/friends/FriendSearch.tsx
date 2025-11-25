"use client";

import { Button } from "../ui/button";
import { useState } from "react";
import { FriendRequest, FriendsUserListItem } from "@/types/req";
import { Item, ItemContent, ItemMedia, ItemTitle } from "../ui/item";
import { Avatar, AvatarFallback, AvatarImage } from "../ui/avatar";
import { Post } from "@/lib/api";
import { toast } from "sonner";
import { Spinner } from "../ui/spinner";
import React from "react";

interface FriendSearchCardProps {
  user: FriendsUserListItem;
  callback: () => void;
}

export const FriendSearchCard: React.FC<FriendSearchCardProps> = ({
  user,
  callback,
}) => {
  const img_url = user.img_url || "/default_profile.jpg";
  const [loading, setLoading] = useState<boolean>(false);

  const addFriend = async () => {
    setLoading(true);
    const req: FriendRequest = {
      type: "addfriend",
      user_id_2: user.id,
    };
    const res = await Post<null, FriendRequest>("/friend", req);
    if (res.success) {
      toast("Friend request sent!");
      if (callback) callback();
    } else {
      toast(res.error.message);
    }
    setLoading(false);
  };

  return (
    <Item className="bg-fcontainer2 rounded-xl mb-4">
      <ItemMedia>
        <Avatar className="rounded-full">
          <AvatarImage src={img_url} alt={`@${user.username}`} />
          <AvatarFallback>{user.username.slice(0, 2)}</AvatarFallback>
        </Avatar>
      </ItemMedia>
      <ItemContent className="flex-row justify-between">
        <ItemTitle>{user.username}</ItemTitle>
        <Button onClick={addFriend}>
          {loading && <Spinner />}
          Add Friend
        </Button>
      </ItemContent>
    </Item>
  );
};
