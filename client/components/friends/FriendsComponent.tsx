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
import { UserCheck, UserRoundSearch, Users } from "lucide-react";
import { Label } from "../ui/label";
import { Input } from "../ui/input";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "../ui/tabs";
import { useState } from "react";

export const FriendsComponent: React.FC = () => {
  const { friends, users, loading, error, fetchFriends, fetchUsers } =
    useFriends();
  const [tab, setTab] = useState<string>("friendslist");
  const [search, setSearch] = useState<string>("");

  if (loading) {
    return <div>Loading...</div>;
  }

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
          <TabsContent value="friendslist">
            {friends.map((f) => (
              <div key={f.id}>{f.username}</div>
            ))}
          </TabsContent>
          <TabsContent value="requests">Requests</TabsContent>
        </CardContent>
        <CardFooter className="relative justify-self-end">
          <Label htmlFor="search" className="sr-only">
            Search
          </Label>
          <Input
            id="search"
            placeholder="Add new friend"
            className="pl-8"
            type="search"
          />
          <UserRoundSearch className="pointer-events-none absolute left-8.5 bottom-0.5 size-4 -translate-y-1/2 opacity-50 select-none" />
        </CardFooter>
      </Card>
    </Tabs>
  );
};
