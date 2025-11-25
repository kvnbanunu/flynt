"use client";

import { Button } from "../ui/button";
import { MenuIcon } from "lucide-react";
import { useState } from "react";
import {
  FriendFyre,
  FriendRequest,
  FriendsListItem,
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
import { DeleteBody, Get, Post, Put } from "@/lib/api";
import { toast } from "sonner";
import { Spinner } from "../ui/spinner";
import React from "react";
import { ApiError, Result } from "@/types/api";
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

interface FriendCardProps {
  friend: FriendsListItem;
  callback: () => void;
}

export const FriendCard: React.FC<FriendCardProps> = ({ friend, callback }) => {
  const img_url = friend.img_url || "/default_profile.jpg";
  const addFlag = friend.status === "pending";
  const [fyres, setFyres] = useState<FriendFyre[]>([]);

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
      if (callback) callback();
    } else {
      toast(res.error.message);
    }
  };

  const fetchFyres = async () => {
    if (fyres.length > 0) {
      return;
    }
    const res = await Get<FriendFyre[]>(`/fyre/user/${friend.id}`);
    if (res.success) {
      setFyres(res.data);
    } else {
      toast(res.error.message);
    }
  };

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
              <TableHead className="text-right">Streak</TableHead>
              <TableHead className="text-right">Status</TableHead>
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
                  <TableCell className="text-right">
                    {fyre.streak_count} 🔥
                  </TableCell>
                  <TableCell className="text-center">
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
