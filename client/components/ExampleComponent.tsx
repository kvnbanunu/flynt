import { Get } from "@/lib/api";
import { ApiError, ApiResponse, Result } from "@/types/api";
import React from "react";

export const ExampleComponent: React.FC<{
  title: string;
}> = async (props) => {
  const { title } = props;
  const res = await Get<Models.User[]>("/api/users");
  if (res.success) {
    const users: Models.User[] = res.data.data;

    return (
      <>
        <h1>{title}</h1>
        <ul>
          {users &&
            users.map((result, index) => {
              return (
                <li key={index}>
                  {result.id}: {result.name}
                </li>
              );
            })}
        </ul>
      </>
    );
  } else {
    console.error(`API Error: ${res.error.statusCode} - ${res.error.message}`);
    return <div>{res.error.statusCode}</div>;
  }
};
