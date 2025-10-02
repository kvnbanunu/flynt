"use client"

import React, { useState } from "react";

export const SigninForm: React.FC = async () => {
  const [inputEmail, setInputEmail] = useState<string>("")

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setInputEmail(event.target.value)
  }
  
  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
  }

  return (
    <form onSubmit={handleSubmit}>
    <label htmlFor="email">Email:</label>
    <input
      type="text"
      id="email"
      value={inputEmail}
      onChange={handleChange}
    />
    <button type="submit">Submit</button>
    </form>
  );
};

export default SigninForm;
