"use client";

import { Get, Post, Put } from "@/lib/api";
import { FriendRequest, FullPost, BonfyreRequest } from "@/types/req";
import React from "react";
import {
  Card,
  CardAction,
  CardContent,
  CardHeader,
  CardTitle,
} from "../ui/card";
import { Avatar, AvatarFallback, AvatarImage } from "../ui/avatar";
import { Button } from "../ui/button";
import { EllipsisVerticalIcon, FlameKindlingIcon } from "lucide-react";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuTrigger,
} from "../ui/dropdown-menu";
import { useAuth } from "@/contexts/AuthContext";
import { toast } from "sonner";
import { PostDialog } from "./PostDialog";
import { LikerAvatars } from "./LikerAvatars";

interface PostCardProps {
  post: FullPost;
  fetchPosts: () => void;
}

export const PostCard: React.FC<PostCardProps> = ({ post, fetchPosts }) => {
  const { user, fetchFyres } = useAuth();
  const img_url = post.img_url || "/default_profile.jpg";

  const likePost = async (id: number) => {
    await Get<null>(`/socialpost/like/${id}`);
    fetchPosts();
  };

  const joinBonfyre = async () => {
    const req: BonfyreRequest = {
      fyre_id: post.fyre_id,
      bonfyre_id: post.bonfyre_id,
    };
    const res = await Post<null, BonfyreRequest>("/fyre/bonfyre", req);
    if (res.success) {
      toast("Successfully joined bonfyre!");
      fetchFyres();
    } else {
      toast(res.error.message);
    }
  };

  const addFriend = async () => {
    if (post.status === "friends") {
      toast(`You are already friends with ${post.username}!`);
      return;
    }
    const req: FriendRequest = {
      type: "addfriend",
      user_id_2: post.user_id,
    };
    const res = await Post<null, FriendRequest>("/friend", req);
    if (res.success) {
      toast("Friend request sent!");
    } else {
      toast(res.error.message);
    }
  };

  const blockUser = async () => {
    const req: FriendRequest = {
      type: "blockfriend",
      user_id_2: post.user_id,
    };
    const res = await Put<null, FriendRequest>("/friend", req);
    if (res.success) {
      toast(`${post.username} is now blocked.`);
      fetchPosts();
    } else {
      toast(res.error.message);
    }
  };

  return (
    <Card className="mb-4 pt-0">
      <CardHeader className="py-2 bg-primary text-primary-foreground rounded-t-xl">
        <CardTitle className="flex flex-row pt-1 items-center gap-4">
          <Avatar>
            <AvatarImage src={img_url} alt={`@${post.username}`} />
            <AvatarFallback>{post.username.slice(0, 2)}</AvatarFallback>
          </Avatar>
          {post.username}
        </CardTitle>
        <CardAction className="flex flex-row items-center justify-end">
          {post.streak_count} 🔥
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button variant="ghost" size="icon-lg" className="rounded-full">
                <EllipsisVerticalIcon />
              </Button>
            </DropdownMenuTrigger>
            {user && post.user_id !== user.id && (
              <DropdownMenuContent align="end">
                <DropdownMenuGroup>
                  <PostDialog
                    trigger={"Join BonFyre"}
                    title={"Join this BonFyre?"}
                    callback={joinBonfyre}
                  >
                    User: {post.username}
                    <br />
                    Fyre: {post.title}
                  </PostDialog>
                  <PostDialog
                    trigger={"Add friend"}
                    title={`Add ${post.username} as a friend?`}
                    callback={addFriend}
                  ></PostDialog>
                  <PostDialog
                    trigger={"Block user"}
                    title={`Are you sure you want to block ${post.username}?`}
                    callback={blockUser}
                  ></PostDialog>
                </DropdownMenuGroup>
              </DropdownMenuContent>
            )}
          </DropdownMenu>
        </CardAction>
      </CardHeader>
      <CardContent>
        <div className="flex justify-between">
          <div className="flex flex-col">
            <div>
              {post.username}
              {post.content}
              {post.streak_count}!
            </div>
            <div>
              {post.title}: {post.streak_count} days
            </div>
          </div>
          <div className="flex items-center gap-2">
            <LikerAvatars likes={post.likes} />
            <div className="flex items-center gap-2 justify-end">
              {post.likes}
              <Button
                size="icon-lg"
                className="rounded-full"
                onClick={() => likePost(post.id)}
              >
                <FlameKindlingIcon />
              </Button>
            </div>
          </div>
        </div>
      </CardContent>
    </Card>
  );
};
