"use client";
import { Post } from "@/lib/api";
import React, { useState } from "react";
import SigninForm from "./signin";
import { FyreList } from "./fyrelist";
import FriendsList from "./friendlist/FriendsList";
import { RegisterForm } from "./Register";
import { Profile } from "./profile";

interface LoginData {
  type: string;
  email: string;
  password: string;
}

export const ExampleHome: React.FC = () => {
  const [isloggedIn, setIsLoggedIn] = useState<boolean>(false);
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [currentUser, setCurrentUser] = useState<Models.User | null>(null);
  const [selectedTab, setSelectedTab] = useState<string>("login");
  const [profileOpened, setProfileOpened] = useState<boolean>(false);

  const onLogin = async (email: string, password: string) => {
    setLoading(true);
    const loginData: LoginData = { type: "email", email: email, password: password };
    const res = await Post<Models.User, LoginData>(
        "/account/login",
        loginData,
    );

    if (res.success) {
      const user: Models.User = res.data;
      setCurrentUser(user);
      setIsLoggedIn(true);
    } else {
      setError(res.error.message);
    }
    setLoading(false);
  };

  const onRegister = (user: Models.User) => {
    setError(null);
    setCurrentUser(user);
    setIsLoggedIn(true);
    setLoading(false);
  }

  if (loading) {
    return <div>Logging in...</div>;
  }

  if (!isloggedIn) {
    return (
        <div>
          {error && <div>{error}</div>}
          {selectedTab == "login" &&
              <div>
                <div>
                  <button
                      className="rounded-xl p-2 text-center bg-gray-300"
                      disabled
                  >Login</button>
                  <button
                      className="rounded-xl p-2 text-center bg-green-300 cursor-pointer"
                      onClick={() => setSelectedTab("register")}
                  >Register</button>
                </div>
                <SigninForm onSubmitHandler={onLogin} />
              </div>
          }

          {selectedTab == "register" &&
              <div>
                <div>
                  <button
                      className="rounded-xl p-2 text-center bg-green-300 cursor-pointer"
                      onClick={() => setSelectedTab("login")}
                  >Login</button>
                  <button
                      className="rounded-xl p-2 text-center bg-gray-300"
                      disabled
                  >Register</button>
                </div>
                <RegisterForm onSuccessHandler={onRegister} />
              </div>
          }
        </div>
    );
  }

  if (isloggedIn && currentUser) {
    return (
        <main>
          <div className="flex gap-4">
            {/*<FriendsList id={currentUser.id} />*/}
            <FriendsList />
            <FyreList user={currentUser} />

            <div className="flex flex-col items-center mt-2">
              <button
                  className="rounded-lg p-2 bg-blue-300 cursor-pointer"
                  onClick={() => setProfileOpened(!profileOpened)}
              >
                {profileOpened ? "Close Profile" : "View Profile"}
              </button>

              {profileOpened && (
                  <div className="mt-0">
                    <Profile user={currentUser} />
                  </div>
              )}
            </div>
          </div>
        </main>
    );
  }
};
