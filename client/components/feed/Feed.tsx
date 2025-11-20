"use client"

import { Get } from "@/lib/api"
import { FullPost } from "@/types/req";
import React, { useEffect, useState } from "react"
import { Card, CardAction, CardContent, CardHeader, CardTitle } from "../ui/card";
import { Avatar, AvatarFallback, AvatarImage } from "../ui/avatar";
import { ScrollArea } from "../ui/scroll-area";

export const Feed: React.FC = () => {
  const [posts, setPosts] = useState<FullPost[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState<boolean>(false);

  useEffect(() => {
    fetchPosts();
  }, [])

  const fetchPosts = async () => {
    setLoading(true);
    const res = await Get<FullPost[]>("/socialpost")
    if (res.success) {
      setError(null);
      setPosts(res.data);
    } else {
      setError(res.error.message);
    }
    setLoading(false);
  }

  return (
    <ScrollArea className="w-full">
      {posts && posts.map((post) => (
        <React.Fragment key={post.id}>
          <PostCard post={post} />
        </React.Fragment>
      ))}
    </ScrollArea>
  )
}

const PostCard: React.FC<{ post: FullPost }> = ({ post }) => {
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
        <CardAction className="pt-2">
          {post.streak_count} 🔥
        </CardAction>
      </CardHeader>
      <CardContent>
        <div className="flex flex-col">
          <div>
            {post.username}{post.content}{post.streak_count}!
          </div>
          <div>
            {post.title}: {post.streak_count} days
          </div>
        </div>
      </CardContent>
    </Card>
  )
}
