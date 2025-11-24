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
import { MenuIcon, UserCheck, Users } from "lucide-react";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "../ui/tabs";
import { useState } from "react";
import {
  FriendFyre,
  FriendRequest,
  FriendsListItem,
  FriendsUserListItem,
  BonfyreRequest,
} from "@/types/req";
import {
  Item,
  ItemActions,
  ItemContent,
  ItemMedia,
  ItemTitle,
} from "../ui/item";
import { Avatar, AvatarFallback, AvatarImage } from "../ui/avatar";
import { Badge } from "../ui/badge";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
} from "../ui/command";
import { DeleteBody, Get, Post, Put } from "@/lib/api";
import { toast } from "sonner";
import { Spinner } from "../ui/spinner";
import React from "react";
import { ApiError, Result } from "@/types/api";
import { ScrollArea } from "../ui/scroll-area";
import { Separator } from "../ui/separator";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "../ui/dialog";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "../ui/table";
import { Checkbox } from "../ui/checkbox";

export const FriendsComponent: React.FC = () => {
  const { friends, users, loading, error, fetchFriends, fetchUsers } =
    useFriends();
  const [tab, setTab] = useState<string>("friendslist");
  const [search, setSearch] = useState<string>("");

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>{error}</div>;
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
  callback: () => void;
}> = ({ friend, callback }) => {
  const img_url = friend.img_url || "/default_profile.jpg";
  const addFlag = friend.status === "pending";
  const [fyres, setFyres] = useState<FriendFyre[]>([]);
  const [error, setError] = useState<string | null>(null);

  const friendRequest = async (type: string) => {
    const req: FriendRequest = {
      type: type,
      user_id_2: friend.id,
    };
    let res: Result<null, ApiError>;
    if (type === "deletefriend") {
      res = await DeleteBody("/friend", req);
    } else {
      res = await Put<null, FriendRequest>("/friend", req);
    }
    if (res.success) {
      if (type === "deletefriend") {
        toast(`You have removed ${friend.username} as a friend.`);
      } else {
        toast(`You are now friends with ${friend.username}!`);
      }
      setError(null);
      if (callback) callback();
    } else {
      setError(res.error.message);
    }
  };

  const fetchFyres = async () => {
    if (fyres.length > 0) {
      return;
    }
    const res = await Get<FriendFyre[]>(`/fyre/user/${friend.id}`);
    if (res.success) {
      setFyres(res.data);
      setError(null);
    } else {
      setError(res.error.message);
    }
  };

  if (error) {
    return <div>{error}</div>;
  }

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
        <ItemActions>
          {addFlag && (
            <Badge
              className="cursor-pointer"
              onClick={() => friendRequest("acceptfriend")}
            >
              Accept Friend
            </Badge>
          )}
          {!addFlag && (
            <Badge>
              {`${friend.status.charAt(0).toUpperCase()}${friend.status.slice(1)}`}
            </Badge>
          )}
          <Badge
            className="cursor-pointer"
            variant="destructive"
            onClick={() => friendRequest("deletefriend")}
          >
            Remove
          </Badge>
          <FriendsFyreList
            friend={friend}
            fyres={fyres}
            fetchFyres={fetchFyres}
          />
        </ItemActions>
      </ItemContent>
    </Item>
  );
};

const FriendsFyreList: React.FC<{
  friend: FriendsListItem;
  fyres: FriendFyre[];
  fetchFyres: () => void;
}> = ({ friend, fyres, fetchFyres }) => {
  const [loading, setLoading] = useState<boolean>(false);
  const [isOpen, setIsOpen] = useState<boolean>(false);

  const onOpen = () => {
    if (isOpen) {
      setIsOpen(false);
      return;
    }
    setIsOpen(true);
    setLoading(true);
    fetchFyres();
    setLoading(false);
  };

  const joinBonfyre = async (fyreID: number, bonfyreID?: number) => {
    const req: BonfyreRequest = { fyre_id: fyreID, bonfyre_id: bonfyreID };
    const res = await Post<null, BonfyreRequest>("/fyre/bonfyre", req);
    if (res.success) {
      toast("Successfully joined bonfyre!");
    } else {
      toast(res.error.message);
    }
  };

  return (
    <Dialog open={isOpen} onOpenChange={onOpen}>
      <DialogTrigger asChild>
        <Button variant="ghost" size="icon-sm" className="rounded-full">
          <MenuIcon />
        </Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>
            {friend.username}
            {`'s Fyre List`}
          </DialogTitle>
        </DialogHeader>
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Title</TableHead>
              <TableHead>Streak</TableHead>
              <TableHead>Status</TableHead>
              <TableHead></TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {loading && <Spinner />}
            {!loading &&
              fyres &&
              fyres.map((fyre) => (
                <TableRow key={fyre.id}>
                  <TableCell>{fyre.title}</TableCell>
                  <TableCell>{fyre.streak_count}</TableCell>
                  <TableCell>
                    <Checkbox disabled checked={fyre.is_checked} />
                  </TableCell>
                  <TableCell>
                    <Button
                      onClick={() => joinBonfyre(fyre.id, fyre.bonfyre_id)}
                    >
                      Join BonFyre
                    </Button>
                  </TableCell>
                </TableRow>
              ))}
          </TableBody>
        </Table>
      </DialogContent>
    </Dialog>
  );
};
