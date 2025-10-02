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
    <form onSubmit={handleSubmit}>
      <label htmlFor="emailInput">Email</label>
      <input
        type="text"
        id="emailInput"
        value={inputEmail}
        onChange={handleEmailChange}
      />
      <label htmlFor="passwordInput">Password</label>
      <input
        type="password"
        id="passwordInput"
        value={inputPassword}
        onChange={handlePasswordChange}
      />
      <button type="submit">Submit</button>
    </form>
  );
};

export default SigninForm;
