"use client";

import { Get } from "@/lib/api";
import { FullPost } from "@/types/req";
import React, { useEffect, useState } from "react";
import {
  Card,
  CardAction,
  CardContent,
  CardHeader,
  CardTitle,
} from "../ui/card";
import { Avatar, AvatarFallback, AvatarImage } from "../ui/avatar";
import { ScrollArea } from "../ui/scroll-area";
import { Button } from "../ui/button";
import { EllipsisVerticalIcon, FlameKindlingIcon } from "lucide-react";

export const Feed: React.FC = () => {
  const [posts, setPosts] = useState<FullPost[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState<boolean>(false);

  useEffect(() => {
    fetchPosts();
  }, []);

  const fetchPosts = async () => {
    setLoading(true);
    const res = await Get<FullPost[]>("/socialpost");
    if (res.success) {
      setError(null);
      setPosts(res.data);
    } else {
      setError(res.error.message);
    }
    setLoading(false);
  };

  const likePost = async (id: number) => {
    setLoading(true);
    await Get<null>(`/socialpost/like/${id}`)
    fetchPosts();
    setLoading(false);
  }

  return (
    <ScrollArea className="w-full">
      {posts &&
        posts.map((post) => (
          <React.Fragment key={post.id}>
            <PostCard post={post} callback={likePost} />
          </React.Fragment>
        ))}
    </ScrollArea>
  );
};

const PostCard: React.FC<{ post: FullPost, callback: (id:number) => void }> = ({ post, callback }) => {
  const img_url = post.img_url || "/default_profile.jpg";
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
          <Button size="icon-lg" className="rounded-full">
            <EllipsisVerticalIcon />
          </Button>
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
                onClick={() => callback(post.id)}
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

const LikerAvatars: React.FC<{ likes: number }> = ({ likes }) => {
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
      ))
      }
    </div>
  );
};
