"use client";

import { Get } from "@/lib/api";
import { FriendsListItem, FriendsUserListItem } from "@/types/req";
import {
  createContext,
  ReactNode,
  useContext,
  useEffect,
  useState,
} from "react";
import { useAuth } from "./AuthContext";
import { toast } from "sonner";

export interface FriendsContextType {
  friends: FriendsListItem[];
  users: FriendsUserListItem[];
  loading: boolean;
  fetchFriends: () => void;
  fetchUsers: () => void;
}

const FriendsContext = createContext<FriendsContextType>({
  friends: [],
  users: [],
  loading: false,
  fetchFriends: function(): void {
    throw new Error("Function not implemented.");
  },
  fetchUsers: function(): void {
    throw new Error("Function not implemented.");
  },
});

interface FriendsProviderProps {
  children: ReactNode;
}

export const FriendsProvider: React.FC<FriendsProviderProps> = ({
  children,
}) => {
  const [friends, setFriends] = useState<FriendsListItem[]>([]);
  const [users, setUsers] = useState<FriendsUserListItem[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const { isAuthenticated } = useAuth();

  useEffect(() => {
    if (isAuthenticated) {
      setLoading(true);
      get<FriendsListItem[]>("/friend/all", setFriends);
      get<FriendsUserListItem[]>("/friend/non", setUsers);
      setLoading(false);
    }
  }, []);

  const fetchFriends = async () => {
    setLoading(true);
    await get<FriendsListItem[]>("/friend/all", setFriends);
    setLoading(false);
  };

  const fetchUsers = async () => {
    setLoading(true);
    await get<FriendsUserListItem[]>("/friend/non", setUsers);
    setLoading(false);
  };

  async function get<T>(url: string, callback: (t: T) => void) {
    const res = await Get<T>(url);
    if (res.success) {
      if (res.data != null) {
        callback(res.data);
      }
    } else {
      toast(res.error.message);
    }
  }

  const value: FriendsContextType = {
    friends,
    users,
    loading,
    fetchFriends,
    fetchUsers,
  };

  return (
    <FriendsContext.Provider value={value}>{children}</FriendsContext.Provider>
  );
};

export const useFriends = () => useContext(FriendsContext);

export default FriendsContext;
