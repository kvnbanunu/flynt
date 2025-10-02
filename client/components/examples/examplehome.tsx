"use client";
import { Post } from "@/lib/api";
import React, { useState } from "react";
import SigninForm from "./signin";
import { FyreList } from "./fyrelist";

interface LoginData {
  email: string;
  password: string;
}

export const ExampleHome: React.FC = () => {
  const [isloggedIn, setIsLoggedIn] = useState<boolean>(false);
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [currentUser, setCurrentUser] = useState<Models.User | null>(null);

  const onLogin = async (email: string, password: string) => {
    setLoading(true);
    const res = await Post<LoginData, Models.User>(`/api/account/login`, {
      email,
      password,
    });

    if (res.success) {
      const user: Models.User = res.data.data;
      setCurrentUser(user);
      setIsLoggedIn(true);
    } else {
      setError(res.error.message);
    }
    setLoading(false);
  };

  if (loading) {
    return <div>Logging in...</div>;
  }

  if (!isloggedIn) {
    return (
      <div>
        {error && <div>{error}</div>}
        <SigninForm onSubmitHandler={onLogin} />
      </div>
    );
  }

  if (isloggedIn && currentUser) {
    return <FyreList user={currentUser} />;
  }
};
