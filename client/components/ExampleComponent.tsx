import { Get } from "@/lib/api";
import { ApiResponse } from "@/types/api";
import React from "react";

export const ExampleComponent: React.FC<{
    title: string;
}> = async (props) => {
    const { title } = props;
    const res: ApiResponse<Models.User> = await Get("api/users");

    const users: Models.User[] = res.data

    return (
        <>
            <h1>{title}</h1>
            <ul>
                {users && users.map((result, index) => {
                    return (
                        <li key={index}>
                            {result.id}: {result.name}
                        </li>
                    );
                })}
            </ul>
        </>
    );
};
