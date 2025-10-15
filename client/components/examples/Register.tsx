"use client";

import { Post } from "@/lib/api";
import { CS_ENV } from "@/lib/utils";
import React, { useState } from "react";

interface RegisterProps {
  onSuccessHandler: (user: Models.User) => void
}

interface RegisterRequest {
  username: string;
  name: string;
  password: string;
  email: string;
}

export const RegisterForm: React.FC<RegisterProps> = (props) => {
  const { onSuccessHandler } = props;

  const [inputUsername, setInputUsername] = useState<string>("")
  const [inputName, setInputName] = useState<string>("")
  const [inputEmail, setInputEmail] = useState<string>("")
  const [inputPassword, setInputPassword] = useState<string>("")
  const [error, setError] = useState<string | null>(null)
  const [loading, setLoading] = useState<boolean>(false)

  const handleUsernameChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setInputUsername(event.target.value);
  }
  const handleNameChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setInputName(event.target.value);
  }
  const handleEmailChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setInputEmail(event.target.value);
  }
  const handlePasswordChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setInputPassword(event.target.value);
  }

  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    setLoading(true);

    const registerData: RegisterRequest = {
      username: inputUsername,
      name: inputName,
      password: inputPassword,
      email: inputEmail
    }
    const res = await Post<Models.User, RegisterRequest>(`${CS_ENV.api_url}/api/account/register`, registerData);

    if (res.success) {
      const user: Models.User = res.data;
      onSuccessHandler(user);
    } else {
      setError(res.error.message);
    }
    setLoading(false);
  };

  if (loading) {
    return <div>Registering...</div>;
  }

  return (
    <form className="p-2 border-2" onSubmit={handleSubmit}>
      {error && <div>{error}</div>}
      <div className="my-1">
        <label className="mx-2" htmlFor="username">
          Username
        </label>
        <input
          id="username"
          type="text"
          value={inputUsername}
          onChange={handleUsernameChange}
          required
        />
      </div>
      <div className="my-1">
        <label className="mx-2" htmlFor="name">
          Name
        </label>
        <input
          id="name"
          type="text"
          value={inputName}
          onChange={handleNameChange}
          required
        />
      </div>
      <div className="my-1">
        <label className="mx-2" htmlFor="email">
          Email
        </label>
        <input
          id="email"
          type="email"
          value={inputEmail}
          onChange={handleEmailChange}
          required
        />
      </div>
      <div className="my-1">
        <label className="mx-2" htmlFor="password">
          Password
        </label>
        <input
          id="password"
          type="password"
          value={inputPassword}
          onChange={handlePasswordChange}
          required
        />
      </div>
      <div className="text-center my-1">
        <button
          className="rounded-xl p-2 text-center bg-green-300 w-full cursor-pointer"
          type="submit"
        >
          Submit
        </button>
      </div>
    </form>
  )
}
