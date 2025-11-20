"use client";
import React, {useCallback, useEffect, useState} from "react";
import { Get, Post } from "@/lib/api";

interface ProfileProps {
    user: Models.User;
}

interface UpdateProfileRequest {
    name?: string;
    bio?: string;
    img_url?: string;
}

export const Profile: React.FC<ProfileProps> = ({ user }) => {
    const [displayName, setDisplayName] = useState<string>(user.username || "");
    const [bio, setBio] = useState<string>(user.bio || "");
    const [avatarUrl, setAvatarUrl] = useState<string>(user.img_url || "");
    const [editing, setEditing] = useState<boolean>(false);
    const [loading, setLoading] = useState<boolean>(true);
    const [message, setMessage] = useState<string | null>(null);
    const [update, setUpdate] = useState<boolean>(false);
    const [error, setError] = useState<string | null>(null);

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        setMessage(null);

        const body: UpdateProfileRequest = {
            name: displayName,
            bio: bio,
            img_url: avatarUrl,
        };

        const res = await Post<Models.User, UpdateProfileRequest>(
            `/profile`,
            body
        );

        if (res.success) {
            setMessage("Profile updated successfully!");
            setEditing(false);
        } else {
            setMessage(`Error: ${res.error.message}`);
            setError(res.error.message);
        }
    };

    const fetchProfile = useCallback(async () => {
        const res = await Get<Models.User>(`/profile`);
        if (res.success && res.data) {
            setDisplayName(res.data.name || "");
            setBio(res.data.bio || "");
            setAvatarUrl(res.data.img_url || "");
        } else {
            // cant figure out what the res.error.message issue is for : setError(res.error.message)
            setError("Failed to fetch profile");
        }
        setLoading(false);
        setUpdate(false);
    }, []);

    useEffect(() => {
        fetchProfile();
    }, [update, fetchProfile]);

    if (loading) {
        return <div>Loading your Profile</div>;
    }

    if (error) {
        return <div>Error: {error}</div>;
    }

    return (
        <div className="m-2 p-2 border-2  rounded-sm">
            <h2 className="text-xl  mb-2">Profile</h2>

            {/*{message && <div className="mb-2 text-center text-sm text-green-600">{message}</div>}*/}

            {!editing ? (
                <div>
                    <div className="flex items-center gap-4 mb-3">
                        {avatarUrl ? (
                            <img
                                src={avatarUrl}
                                alt="avatar"
                                className="w-16 h-16 rounded-full border"
                            />
                        ) : (
                            <div className="w-16 h-16 rounded-full bg-gray-200 flex items-center justify-center">
                                ?
                            </div>
                        )}
                        <div>
                            <h3 className="text-xl">{displayName || user.email}</h3>
                            <p className="text-gray-600 text-sm">{bio || "No bio yet."}</p>
                        </div>
                    </div>
                    <button
                        className="rounded-lg p-2 text-center bg-blue-300 cursor-pointer w-full"
                        onClick={() => setEditing(true)}
                    >
                        Edit Profile
                    </button>
                </div>
            ) : (
                <form onSubmit={handleSubmit}>
                    <div className="mb-2">
                        <label className="block text-sm font-medium">Display Name</label>
                        <input
                            type="text"
                            value={displayName}
                            onChange={(e) => setDisplayName(e.target.value)}
                            className="border-2 rounded-md w-full p-1"
                        />
                    </div>

                    <div className="mb-2">
                        <label className="block text-sm font-medium">Bio</label>
                        <textarea
                            value={bio}
                            onChange={(e) => setBio(e.target.value)}
                            className="border-2 rounded-md w-full p-1"
                        />
                    </div>

                    <div className="mb-2">
                        <label className="block text-sm font-medium">Avatar URL</label>
                        <input
                            type="text"
                            value={avatarUrl}
                            onChange={(e) => setAvatarUrl(e.target.value)}
                            className="border-2 rounded-md w-full p-1"
                        />
                    </div>

                    <div className="flex justify-between gap-2 mt-2">
                        <button
                            type="submit"
                            className="rounded-lg p-2 text-center bg-green-300 w-full cursor-pointer"
                        >
                            Save
                        </button>
                        <button
                            type="button"
                            className="rounded-lg p-2 text-center bg-red-300 w-full cursor-pointer"
                            onClick={() => setEditing(false)}
                        >
                            Cancel
                        </button>
                    </div>
                </form>
            )}
        </div>
    );
};
