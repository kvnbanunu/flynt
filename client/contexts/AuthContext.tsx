"use client";

import { Get, Post } from "@/lib/api";
import { useRouter } from "next/router";
import { createContext, ReactNode, useContext, useEffect, useState } from "react";
import { LoginRequest, RegisterRequest } from "@/types/req";

const AuthContext = createContext({});

interface AuthProviderProps {
  children: ReactNode;
}

export const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
  const [user, setUser] = useState<Models.User | null>(null);
  const [loading, setLoading] = useState(true);
  const router = useRouter();

  useEffect(() => {
    const checkUser = async () => {
      const res = await Get<Models.User>("/user");
      if (res.success) {
        setUser(res.data);
      } else {
        setUser(null);
      }
      setLoading(false);
    };
    checkUser();
  }, []);

  const login = async (credentials: LoginRequest): Promise<Boolean> => {
    setLoading(true);
    const res = await Post<Models.User, LoginRequest>("/account/login", credentials);
    if (res.success) {
      setUser(res.data);
      // set other stuff here
      router.push('/'); // send to home
    }
    setLoading(false);
    return res.success;
  }

  const register = async (credentials: RegisterRequest): Promise<Boolean> => {
    setLoading(true);
    const res = await Post<Models.User, RegisterRequest>("/account/register", credentials);
    if (res.success) {
      setUser(res.data);
      // set other stuff here
      router.push('/'); // send to home
    }
    setLoading(false);
    return res.success;
  }

  const value = {
    user,
    loading,
    login,
    register,
    isAuthenticated: !!user,
  };

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
};

export const useAuth = () => useContext(AuthContext);

export default AuthContext;
