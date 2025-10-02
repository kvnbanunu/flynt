"use client";

import React, { useState } from "react";

interface SigninFormProps {
  onSubmitHandler: (email: string, password: string) => void;
}

export const SigninForm: React.FC<SigninFormProps> = (props) => {
  const { onSubmitHandler } = props;

  const [inputEmail, setInputEmail] = useState<string>("");
  const [inputPassword, setInputPassword] = useState<string>("");

  const handleEmailChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setInputEmail(event.target.value);
  };

  const handlePasswordChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setInputPassword(event.target.value);
  };

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    onSubmitHandler(inputEmail, inputPassword);
  };

  return (
    <form className="p-2 border-2" onSubmit={handleSubmit}>
      <div className="my-1">
      <label className="mx-2" htmlFor="emailInput">Email</label>
        <input
          type="text"
          id="emailInput"
          value={inputEmail}
          onChange={handleEmailChange}
        />
      </div>
      <div className="my-1">
        <label className="mx-2" htmlFor="passwordInput">Password</label>
        <input
          type="password"
          id="passwordInput"
          value={inputPassword}
          onChange={handlePasswordChange}
        />
      </div>
      <div className="text-center my-1">
        <button className="rounded-full p-2 text-center bg-green-300 w-full" type="submit">Submit</button>
      </div>
    </form>
  );
};

export default SigninForm;
