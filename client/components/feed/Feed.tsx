"use client";

import { Get, Post, Put } from "@/lib/api";
import { FriendRequest, FullPost, BonfyreRequest } from "@/types/req";
import React, { ReactNode, useEffect, useState } from "react";
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
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuTrigger,
} from "../ui/dropdown-menu";
import { useAuth } from "@/contexts/AuthContext";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "../ui/dialog";
import { toast } from "sonner";

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

  return (
    <ScrollArea className="w-full">
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

const PostCard: React.FC<{
  post: FullPost;
  fetchPosts: () => void;
}> = ({ post, fetchPosts }) => {
  const { user, fetchFyres } = useAuth();
  const img_url = post.img_url || "/default_profile.jpg";

  const likePost = async (id: number) => {
    await Get<null>(`/socialpost/like/${id}`);
    fetchPosts();
  };

  const joinBonfyre = async () => {
    const req: BonfyreRequest = { fyre_id: post.fyre_id, bonfyre_id: post.bonfyre_id };
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
      ))}
    </div>
  );
};

const PostDialog: React.FC<{
  trigger: string;
  title: string;
  callback: () => void;
  children?: ReactNode;
}> = ({ trigger, title, callback, children }) => {
  const [isOpen, setIsOpen] = useState<boolean>(false);

  return (
    <Dialog open={isOpen} onOpenChange={setIsOpen}>
      <form>
        <DialogTrigger asChild>
          <Button variant="ghost">{trigger}</Button>
        </DialogTrigger>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>{title}</DialogTitle>
            <DialogDescription className="text-md">
              {children}
            </DialogDescription>
          </DialogHeader>
          <DialogFooter>
            <DialogClose asChild>
              <Button variant="outline">Cancel</Button>
            </DialogClose>
            <DialogClose asChild>
              <Button type="submit" onClick={callback}>
                Confirm
              </Button>
            </DialogClose>
          </DialogFooter>
        </DialogContent>
      </form>
    </Dialog>
  );
};
