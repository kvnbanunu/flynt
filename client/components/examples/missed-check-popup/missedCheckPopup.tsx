"use client";
import React from "react";
import { Card } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Put } from "@/lib/api";

interface MissedCheckPopupProps {
    fyres: Models.Fyre[];
    onResolve: (id: number) => void;
}

const MissedCheckPopup: React.FC<MissedCheckPopupProps> = ({ fyres, onResolve }) => {
    console.log("INSIDE PROP: fyres:", fyres);
    if (fyres.length === 0) return null;
    return (
        <div className="absolute top-4 left-1/2 transform -translate-x-1/2 z-50 w-full max-w-md flex flex-col gap-3">
            {fyres.map((f, index) => (
                <Card
                    key={f.id}
                    className="bg-slate-800 text-white p-4 rounded-2xl shadow-xl animate-fadeIn"
                >
                    <h3 className="text-lg font-semibold mb-2">
                        You missed &quot;{f.title}&quot; yesterday
                    </h3>
                    <p className="text-sm mb-3 text-gray-300">
                        Would you like to keep your streak or reset it?
                    </p>
                    <div className="flex justify-end gap-2">
                        <Button
                            variant="secondary"
                            onClick={() => handleResolve(f.id, true, onResolve)}
                        >
                            Keep Streak
                        </Button>
                        <Button
                            variant="destructive"
                            onClick={() => handleResolve(f.id, false, onResolve)}
                        >
                            Reset
                        </Button>
                    </div>
                    <div className="text-xs text-gray-500 mt-1">
                        Fyre {index + 1} of {fyres.length}
                    </div>
                </Card>
            ))}
        </div>
    );
};

async function handleResolve(
    id: number,
    keep: boolean,
    onResolve: (id: number) => void
) {
    // optional: send to backend if needed
    try {
        const res = await Put(`/fyre/missed/${id}`, { keepStreak: keep });
        if (!res.success) console.warn("Failed backend update", res.error);
    } catch (err) {
        console.error("Network error updating missed fyre", err);
    }

    // always remove from popup immediately
    onResolve(id);
}

export default MissedCheckPopup;
