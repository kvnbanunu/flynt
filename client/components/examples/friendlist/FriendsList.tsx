"use client";
import { Get } from "@/lib/api";
import { CS_ENV } from "@/lib/utils";
import React, { useEffect, useState } from "react";

interface FriendsListItem {
  id: number;
  name: string;
  status: string;
}

export const FriendsList: React.FC<{ id: number }> = (props) => {
  const { id } = props;
  const [friends, setFriends] = useState<FriendsListItem[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);
  const [update, setUpdate] = useState<boolean>(false);

  const fetchFriends = async () => {
    const res = await Get<FriendsListItem[]>(`${CS_ENV.api_url}/api/friend/${id}`);
    if (res.success) {
      setFriends(res.data);
    } else {
      setError(res.error.message);
    }
    setLoading(false);
    setUpdate(false);
  };

  useEffect(() => {
    fetchFriends();
  }, [update]);

  if (loading) {
    return <div>Loading Friends List...</div>;
  }

  return (
    <div>
      <h1>Friends List</h1>
      {error && <div className="bg-red-300">Error: {error}</div>}
      <ul>{friends && friends.map((f, index) => (
        <li key={index}>
          {f.id}  {f.name}  {f.status}
        </li>
      ))}</ul>
    </div>
  );
};

export default FriendsList;
