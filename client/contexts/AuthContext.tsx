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
import { FullFyre, LoginRequest, RegisterRequest } from "@/types/req";
import { toast } from "sonner";

export interface AuthContextType {
  user?: Models.User | null;
  fyres: FullFyre[];
  categories: Models.Category[];
  loading: boolean;
  checkUser: () => void;
  setUser: (user: Models.User) => void;
  login: (credentials: LoginRequest) => Promise<Boolean>;
  register: (credentials: RegisterRequest) => Promise<Boolean>;
  logout: () => Promise<Boolean>;
  fetchFyres: () => void;
  isAuthenticated: boolean;
}

const AuthContext = createContext<AuthContextType>({
  fyres: [],
  categories: [],
  loading: false,
  checkUser: function(): void {
    throw new Error("Function not implemented.");
  },
  setUser: function(_user: Models.User): void {
    throw new Error("Function not implemented.");
  },
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
  const [fyres, setFyres] = useState<FullFyre[]>([]);
  const [categories, setCategories] = useState<Models.Category[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const router = useRouter();

  useEffect(() => {
    checkUser();
  }, []);

  const checkUser = async () => {
    const res = await Get<Models.User>("/user");
    if (res.success) {
      setUser(res.data);
      fetchFyres();
      fetchCategories();
    } else {
      setUser(null);
      setFyres([]);
    }
    setLoading(false);
  };

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
      await fetchCategories();
      router.push("/"); // send to home
    } else {
      toast(res.error.message);
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
      await fetchCategories();
      router.push("/"); // send to home
    } else {
      toast(res.error.message);
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
      router.push("/login");
    } else {
      toast(res.error.message);
    }
    setLoading(false);
    return res.success;
  };

  const fetchFyres = async () => {
    const url = fyres.length === 0 ? "/fyre" : "/fyre/full";
    const res = await Get<FullFyre[]>(url);
    if (res.success) {
      setFyres(res.data);
    } else {
      toast(res.error.message);
    }
  };
  const fetchCategories = async () => {
    const res = await Get<Models.Category[]>("/fyre/categories");
    if (res.success) {
      setCategories(res.data);
    } else {
      toast(res.error.message);
    }
  };

  const value: AuthContextType = {
    user,
    fyres,
    categories,
    loading,
    checkUser,
    setUser,
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
