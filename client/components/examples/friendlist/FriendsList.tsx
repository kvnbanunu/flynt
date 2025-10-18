"use client";
import { DeleteBody, Get, Post, Put } from "@/lib/api";
import React, { useCallback, useEffect, useState } from "react";

interface FriendsListItem {
  id: number;
  name: string;
  status: string;
}

interface FriendRequest {
  type: string;
  user_id_2: number;
}

export const FriendsList: React.FC = () => {
  const [friends, setFriends] = useState<FriendsListItem[]>([]);
  const [users, setUsers] = useState<Models.User[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);
  const [update, setUpdate] = useState<boolean>(false);
  const [selectedTab, setSelectedTab] = useState<string>("friendslist");

  const fetchFriends = useCallback(async () => {
    const res = await Get<FriendsListItem[]>("/friend");
    if (res.success) {
      setFriends(res.data);
    } else {
      setError(res.error.message);
    }
    setLoading(false);
    setUpdate(false);
  }, []);

  const fetchUsers = useCallback(async () => {
    const res = await Get<Models.User[]>("/user/redacted");
    if (res.success) {
      setUsers(res.data);
    } else {
      setError(res.error.message);
    }
  }, []);

  const handleTabSwitch = () => {
    const tab = selectedTab === "friendslist" ? "addfriend" : "friendslist";
    setSelectedTab(tab);
  };

  const onUpdate = () => {
    setLoading(true);
    setSelectedTab("friendslist");
    setUpdate(true);
  };

  useEffect(() => {
    fetchFriends();
    fetchUsers();
  }, [update, fetchFriends, fetchUsers]);

  if (loading) {
    return <div>Loading Friends List...</div>;
  }

  return (
    <div className="p-2 m-2 border-2 rounded-sm">
      {error && <div className="bg-red-300">Error: {error}</div>}

      {selectedTab === "friendslist" && (
        <div>
          <h1 className="text-xl">
            Friends List
            <button
              className="bg-blue-200 mx-1 border-1 rounded-lg px-2 cursor-pointer"
              onClick={handleTabSwitch}
            >
              All Users
            </button>
          </h1>
          <FriendsListList friends={friends} handler={onUpdate} />
        </div>
      )}

      {selectedTab === "addfriend" && (
        <div>
          <h1 className="text-xl">
            <button
              className="bg-blue-200 mr-1 border-1 rounded-lg px-2 cursor-pointer"
              onClick={handleTabSwitch}
            >
              Friends List
            </button>
            All Users
          </h1>
          <UsersList users={users} handler={onUpdate} />
        </div>
      )}
    </div>
  );
};

const FriendsListList: React.FC<{
  friends: FriendsListItem[];
  handler: () => void;
}> = (props) => {
  const { friends, handler } = props;
  const [error, setError] = useState<string | null>(null);

  const handleRemove = async (friendID: number) => {
    const req: FriendRequest = {
      type: "deletefriend",
      user_id_2: friendID,
    };
    const res = await DeleteBody("/friend", req);
    if (res.success) {
      setError(null);
      handler();
    } else {
      setError(res.error.message);
    }
  };

  const handleAccept = async (friendID: number) => {
    const req: FriendRequest = {
      type: "acceptfriend",
      user_id_2: friendID,
    };
    const res = await Put<FriendRequest>("/friend", req);
    if (res.success) {
      setError(null);
      handler();
    } else {
      setError(res.error.message);
    }
  };

  return (
    <ul className="divide-y divide-gray-300">
      {error && <div className="bg-red-300 rounded-sm">{error}</div>}
      {friends &&
        friends.map((f, index) => {
          return (
            <li className="flex gap-4 justify-between" key={index}>
              <div className="w-32">{f.name}</div>
              {f.status === "friends" && (
                <div className="w-20 bg-green-300 text-sm font-semibold text-center rounded-lg px-1 py-1">
                  {f.status}
                </div>
              )}
              {f.status === "sent" && (
                <div className="w-20 bg-blue-300 text-sm font-semibold text-center rounded-lg px-1 py-1">
                  {f.status}
                </div>
              )}
              {f.status === "pending" && (
                <button
                  className="w-20 bg-blue-300 text-sm font-semibold rounded-lg text-center px-1 py-1 cursor-pointer"
                  onClick={() => handleAccept(f.id)}
                >
                  Accept
                </button>
              )}
              {f.status === "blocked" && (
                <div className="w-20 bg-gray-300 text-sm font-semibold rounded-lg text-center px-1 py-1">
                  {f.status}
                </div>
              )}
              <button
                className="w-20 bg-red-300 text-sm font-semibold rounded-lg text-center px-1 py-1 cursor-pointer"
                onClick={() => handleRemove(f.id)}
              >
                Remove
              </button>
            </li>
          );
        })}
    </ul>
  );
};

const UsersList: React.FC<{
  users: Models.User[];
  handler: () => void;
}> = (props) => {
  const { users, handler } = props;
  const [error, setError] = useState<string | null>(null);

  const handleAddFriend = async (friendID: number) => {
    const req: FriendRequest = {
      type: "addfriend",
      user_id_2: friendID,
    };
    const res = await Post("/friend", req);
    if (res.success) {
      setError(null);
      handler();
    } else {
      setError(res.error.message);
    }
  };

  const handleBlock = async (friendID: number) => {
    const res = await Put<FriendRequest>("/friend", {
      type: "blockfriend",
      user_id_2: friendID,
    });
    if (res.success) {
      setError(null);
      handler();
    } else {
      setError(res.error.message);
    }
  };

  return (
    <ul className="w-full mx-auto divide-y divide-gray-300">
      {error && <div className="bg-red-300 rounded-sm">{error}</div>}
      {users &&
        users.map((user, index) => {
          return (
            <li className="flex gap-4 justify-between" key={index}>
              <div className="w-32 content-center">{user.name}</div>
              <button
                className="w-24 bg-green-300 text-sm font-semibold rounded-lg py-1 px-1 cursor-pointer"
                onClick={() => handleAddFriend(user.id)}
              >
                Add Friend
              </button>
              <button
                className="w-16 bg-red-300 text-sm font-semibold rounded-lg py-1 px-1 cursor-pointer"
                onClick={() => handleBlock(user.id)}
              >
                Block
              </button>
            </li>
          );
        })}
    </ul>
  );
};

export default FriendsList;
