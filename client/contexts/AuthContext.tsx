"use client";

import { Get, Post } from "@/lib/api";
import { useRouter } from "next/navigation";
import {
  createContext,
  ReactNode,
  useContext,
  useEffect,
  useState,
} from "react";
import { LoginRequest, RegisterRequest } from "@/types/req";

export interface AuthContextType {
  user?: Models.User | null;
  error?: string | null;
  fyres: Models.Fyre[];
  loading: boolean;
  login: (credentials: LoginRequest) => Promise<Boolean>;
  register: (credentials: RegisterRequest) => Promise<Boolean>;
  logout: () => Promise<Boolean>;
  fetchFyres: () => void;
  isAuthenticated: boolean;
}

const AuthContext = createContext<AuthContextType>({
  fyres: [],
  loading: false,
  login: function(_credentials: LoginRequest): Promise<Boolean> {
    throw new Error("Function not implemented.");
  },
  register: function(_credentials: RegisterRequest): Promise<Boolean> {
    throw new Error("Function not implemented.");
  },
  logout: function(): Promise<Boolean> {
    throw new Error("Function not implemented.");
  },
  fetchFyres: function(): void {
    throw new Error("Function not implemented.");
  },
  isAuthenticated: false,
});

interface AuthProviderProps {
  children: ReactNode;
}

export const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
  const [user, setUser] = useState<Models.User | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [fyres, setFyres] = useState<Models.Fyre[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const router = useRouter();

  useEffect(() => {
    const checkUser = async () => {
      const res = await Get<Models.User>("/user");
      if (res.success) {
        setUser(res.data);
        fetchFyres();
      } else {
        setUser(null);
        setFyres([]);
      }
      setLoading(false);
    };
    checkUser();
  }, []);

  const login = async (credentials: LoginRequest): Promise<Boolean> => {
    setLoading(true);
    const res = await Post<Models.User, LoginRequest>(
      "/account/login",
      credentials,
    );
    if (res.success) {
      setUser(res.data);
      await fetchFyres();
      // set other stuff here
      router.push("/"); // send to home
    } else {
      setError(res.error.message);
    }
    setLoading(false);
    return res.success;
  };

  const register = async (credentials: RegisterRequest): Promise<Boolean> => {
    setLoading(true);
    const res = await Post<Models.User, RegisterRequest>(
      "/account/register",
      credentials,
    );
    if (res.success) {
      setUser(res.data);
      // set other stuff here
      router.push("/"); // send to home
    } else {
      setError(res.error.message);
    }
    setLoading(false);
    return res.success;
  };

  const logout = async (): Promise<Boolean> => {
    setLoading(true);
    const res = await Post<null, null>("/account/logout", null);
    if (res.success) {
      setUser(null);
      setFyres([]);
      setError(null);
      router.push("/login");
    } else {
      setError(res.error.message);
    }
    setLoading(false);
    return res.success;
  };

  const fetchFyres = async () => {
    const res = await Get<Models.Fyre[]>("/fyre");
    if (res.success) {
      setFyres(res.data);
    } else {
      setError(res.error.message);
    }
  };

  const value: AuthContextType = {
    user,
    error,
    fyres,
    loading,
    login,
    register,
    logout,
    fetchFyres,
    isAuthenticated: !!user,
  };

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
};

export const useAuth = () => useContext(AuthContext);

export default AuthContext;
