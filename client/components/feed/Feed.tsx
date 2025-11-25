"use client";

import { Get } from "@/lib/api";
import { FullPost } from "@/types/req";
import React, { useEffect, useState } from "react";
import { ScrollArea } from "../ui/scroll-area";
import { toast } from "sonner";
import { Spinner } from "../ui/spinner";
import { PostCard } from "./PostCard";

export const Feed: React.FC = () => {
  const [posts, setPosts] = useState<FullPost[]>([]);
  const [loading, setLoading] = useState<boolean>(false);

  useEffect(() => {
    fetchPosts();
  }, []);

  const fetchPosts = async () => {
    setLoading(true);
    const res = await Get<FullPost[]>("/socialpost");
    if (res.success) {
      setPosts(res.data);
    } else {
      toast(res.error.message);
    }
    setLoading(false);
  };

  return (
    <ScrollArea className="w-full">
      {loading && <Spinner />}
      {posts &&
        posts.map((post) => {
          if (post.status !== "blocked")
            return (
              <React.Fragment key={post.id}>
                <PostCard post={post} fetchPosts={fetchPosts} />
              </React.Fragment>
            );
        })}
    </ScrollArea>
  );
};
