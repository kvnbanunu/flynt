"use client";
import { Post } from "@/lib/api";
import React, { useState } from "react";

interface AddFyreProps {
  onSuccessHandler: () => void;
}

interface AddFyreRequest {
  title: string;
  streak_count: number;
}

export const AddFyre: React.FC<AddFyreProps> = (props) => {
  const { onSuccessHandler } = props;
  const [collapsed, setCollapsed] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);
  const [inputTitle, setInputTitle] = useState<string>("");
  const [inputStreak, setInputStreak] = useState<number>(0);

  const handleTitleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setInputTitle(event.target.value);
  };

  const handleStreakChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const input = event.target.value;
    const num = input === "" ? 0 : Number(input);
    setInputStreak(num);
  };

  const reset = () => {
    setError(null);
    setInputTitle("");
    setInputStreak(0);
    setCollapsed(true);
  };

  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    const fyre: AddFyreRequest = {
      title: inputTitle,
      streak_count: inputStreak,
    };
    const res = await Post<Models.Fyre, AddFyreRequest>(
      "/fyre",
      fyre,
    );

    if (res.success) {
      reset();
      onSuccessHandler();
    } else {
      setError(res.error.message);
    }
  };

  if (collapsed) {
    return (
      <div className="text-center my-1">
        <button
          className="rounded-full p-2 text-center bg-green-300 w-full cursor-pointer"
          onClick={() => setCollapsed(false)}
        >
          Add a new Fyre
        </button>
      </div>
    );
  }

  return (
    <div className="p-2 border-2 rounded-sm">
      <div>
        <button
          className="rounded-full px-2 mr-2 align-middle text-center bg-red-300 aspect-square"
          onClick={reset}
        >
          x
        </button>
        Add a new Fyre
      </div>
      {error && <div className="bg-red-100 w-full">{error}</div>}
      <form onSubmit={handleSubmit}>
        <div className="my-1">
          <label className="mr-2" htmlFor="titleInput">
            Title
          </label>
          <input
            type="text"
            id="titleInput"
            value={inputTitle}
            onChange={handleTitleChange}
            className="border-2 rounded-xs w-full"
          />
        </div>
        <div className="my-1">
          <label className="mr-2" htmlFor="streakInput">
            Streak Count
          </label>
          <input
            type="number"
            id="streakInput"
            value={inputStreak}
            onChange={handleStreakChange}
            className="border-2 rounded-xs w-full"
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
    </div>
  );
};

export default AddFyre;
