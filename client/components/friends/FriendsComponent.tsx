"use client";

import { useFriends } from "@/contexts/FriendsContext";
import {
  Card,
  CardAction,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "../ui/card";
import { Button } from "../ui/button";
import { UserCheck, Users } from "lucide-react";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "../ui/tabs";
import { useState } from "react";
import {
  FriendRequest,
  FriendsListItem,
  FriendsUserListItem,
} from "@/types/req";
import { Item, ItemContent, ItemMedia, ItemTitle } from "../ui/item";
import { Avatar, AvatarFallback, AvatarImage } from "../ui/avatar";
import { Badge } from "../ui/badge";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
} from "../ui/command";
import { Post } from "@/lib/api";
import { toast } from "sonner";
import { Spinner } from "../ui/spinner";
import React from "react";

export const FriendsComponent: React.FC = () => {
  const { friends, users, loading, error, fetchFriends, fetchUsers } =
    useFriends();
  const [tab, setTab] = useState<string>("friendslist");
  const [search, setSearch] = useState<string>("");

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>{error}</div>
  }

  const onAddFriend = () => {
    fetchFriends();
    setSearch("");
  };

  return (
    <Tabs defaultValue="friendslist" className="w-full">
      <Card className="h-full w-full pt-0">
        <CardHeader className="bg-primary text-primary-foreground rounded-t-xl">
          <CardTitle className="pt-3 pb-1">Friends List</CardTitle>
          <CardAction className="pt-0.5">
            <TabsList className="bg-primary">
              <Button
                variant="outline"
                size="icon"
                className="rounded-full"
                asChild
                onClick={() => setTab("friendslist")}
                hidden={tab === "friendslist"}
              >
                <TabsTrigger
                  value="friendslist"
                  className="dark:text-foreground"
                >
                  <Users />
                </TabsTrigger>
              </Button>
              <Button
                variant="outline"
                size="icon"
                className="rounded-full"
                asChild
                onClick={() => setTab("requests")}
                hidden={tab === "requests"}
              >
                <TabsTrigger value="requests" className="dark:text-foreground">
                  <UserCheck />
                </TabsTrigger>
              </Button>
            </TabsList>
          </CardAction>
        </CardHeader>
        <CardContent className="grow">
          {!search && (
            <React.Fragment>
              <TabsContent value="friendslist">
                {friends.map((f) => (
                  <FriendCard key={f.id} friend={f} />
                ))}
              </TabsContent>
              <TabsContent value="requests">Requests</TabsContent>
            </React.Fragment>
          )}
        </CardContent>
        <CardFooter className="relative justify-self-end">
          <Command>
            <CommandEmpty>{search && "No users found."}</CommandEmpty>
            <CommandGroup>
              {search &&
                users.map((u) => (
                  <CommandItem key={u.id} asChild value={u.username}>
                    <FriendSearchCard user={u} callback={onAddFriend} />
                  </CommandItem>
                ))}
            </CommandGroup>
            <CommandInput
              placeholder="Add new friend"
              onValueChange={setSearch}
            />
          </Command>
        </CardFooter>
      </Card>
    </Tabs>
  );
};

const FriendSearchCard: React.FC<{
  user: FriendsUserListItem;
  callback: () => void;
}> = ({ user, callback }) => {
  const img_url = user.img_url || "/default_profile.jpg";
  const [error, setError] = useState<string | null>(null);
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
      setError(null);
      if (callback) callback();
    } else {
      setError(res.error.message);
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
          {error && error}
          {loading && <Spinner />}
          Add Friend
        </Button>
      </ItemContent>
    </Item>
  );
};

const FriendCard: React.FC<{
  friend: FriendsListItem;
}> = ({ friend }) => {
  const img_url = friend.img_url || "/default_profile.jpg";

  return (
    <Item className="bg-fcontainer2 rounded-xl mb-4">
      <ItemMedia>
        <Avatar className="rounded-full">
          <AvatarImage src={img_url} alt={`@${friend.username}`} />
          <AvatarFallback>{friend.username.slice(0, 2)}</AvatarFallback>
        </Avatar>
      </ItemMedia>
      <ItemContent className="flex-row justify-between">
        <ItemTitle>{friend.username}</ItemTitle>
        <Badge>{friend.status}</Badge>
      </ItemContent>
    </Item>
  );
};
