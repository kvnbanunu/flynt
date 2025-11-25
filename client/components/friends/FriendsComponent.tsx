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
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
} from "../ui/command";
import React from "react";
import { ScrollArea } from "../ui/scroll-area";
import { Separator } from "../ui/separator";
import { FriendSearchCard } from "./FriendSearch";
import { FriendCard } from "./FriendCard";

export const FriendsComponent: React.FC = () => {
  const { friends, users, loading, fetchFriends, fetchUsers } =
    useFriends();
  const [tab, setTab] = useState<string>("friendslist");
  const [search, setSearch] = useState<string>("");

  if (loading) {
    return <div>Loading...</div>;
  }

  const onFriendRequest = () => {
    fetchFriends();
    fetchUsers();
    setSearch("");
  };

  return (
    <ScrollArea className="w-full">
      <Tabs defaultValue="friendslist" className="w-full">
        <Card className="w-full pt-0">
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
                  <TabsTrigger
                    value="requests"
                    className="dark:text-foreground"
                  >
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
                  {friends &&
                    friends.map((f) => {
                      if (f.status === "friends")
                        return (
                          <FriendCard
                            key={f.id}
                            friend={f}
                            callback={onFriendRequest}
                          />
                        );
                    })}
                </TabsContent>
                <TabsContent value="requests">
                  Requests
                  <Separator className="mb-4 mt-2" />
                  {friends &&
                    friends.map((f) => {
                      if (f.status !== "friends")
                        return (
                          <FriendCard
                            key={f.id}
                            friend={f}
                            callback={onFriendRequest}
                          />
                        );
                    })}
                </TabsContent>
              </React.Fragment>
            )}
          </CardContent>
          <CardFooter className="relative justify-self-end">
            <Command>
              <CommandEmpty>{search && "No users found."}</CommandEmpty>
              <CommandGroup>
                {search &&
                  users &&
                  users.map((u) => (
                    <CommandItem key={u.id} asChild value={u.username}>
                      <FriendSearchCard user={u} callback={onFriendRequest} />
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
    </ScrollArea>
  );
};
