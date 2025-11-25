"use client";

import React, { useState } from "react";
import { Button } from "../ui/button";
import { Input } from "../ui/input";
import {
  Select,
  SelectTrigger,
  SelectValue,
  SelectContent,
  SelectItem,
} from "../ui/select";
import { Post, Delete } from "@/lib/api";
import { toast } from "sonner";
import type { CreateGoalRequest } from "@/types/req";
import { useAuth } from "@/contexts/AuthContext";
import { Item, ItemActions, ItemContent, ItemTitle } from "../ui/item";

export const GoalSection: React.FC<{ fyreId: number; goal?: Models.Goal }> = ({
  fyreId,
  goal,
}) => {
  const { fetchFyres } = useAuth();
  const [goalType, setGoalType] = useState<string>("");
  const [goalData, setGoalData] = useState<string>("");

  const saveGoal = async () => {
    if (!goalType || !goalData) {
      toast.error("Please select a goal type and enter a value");
      return;
    }

    const req: CreateGoalRequest = {
      fyre_id: fyreId,
      goal_type_id: parseInt(goalType),
      description:
        goalType == "1" ? "Reach target date" : "Reach target streak count",
      data: goalData,
    };

    const res = await Post<Models.Goal, CreateGoalRequest>("/goal", req);
    if (res.success) {
      toast.success("Goal saved!");
      fetchFyres();
    } else {
      toast.error(res.error.message);
    }
  };

  const deleteGoal = async () => {
    const res = await Delete(`/goal/${fyreId}`);
    if (res.success) {
      toast.success("Goal removed");
      setGoalData("");
      setGoalType("");
      fetchFyres();
    } else {
      toast.error(res.error.message);
    }
  };

  return (
    <Item variant="outline">
      <ItemContent className="flex flex-row justify-between">
        <div className="flex flex-col gap-3">
          <ItemTitle className="text-lg">Goal</ItemTitle>
          {goal ? (
            <div className="flex flex-col gap-3">
              <p>
                <strong>Type:</strong>{" "}
                {goal.goal_type_id === 1 ? "Until Date" : "Number of Days"}
              </p>
              <p>
                <strong>Value:</strong>{" "}
                {goal.goal_type_id === 1
                  ? goal.data
                    ? new Date(goal.data).toLocaleDateString()
                    : "-"
                  : (goal.data ?? "-")}{" "}
                days
              </p>
            </div>
          ) : (
            <div className="flex flex-col gap-3">
              <Select value={goalType} onValueChange={setGoalType}>
                <SelectTrigger>
                  <SelectValue placeholder="Select goal type" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="1">Until Date</SelectItem>
                  <SelectItem value="2">Number of Days</SelectItem>
                </SelectContent>
              </Select>

              {goalType === "1" && (
                <Input
                  type="date"
                  value={goalData}
                  onChange={(e) => setGoalData(e.target.value)}
                />
              )}
              {goalType === "2" && (
                <Input
                  type="number"
                  placeholder="Number of days"
                  value={goalData}
                  onChange={(e) => setGoalData(e.target.value)}
                />
              )}
            </div>
          )}
        </div>

        <ItemActions>
          {goal ? (
            <Button
              variant="destructive"
              size="sm"
              onClick={deleteGoal}
              className="w-fit"
            >
              Remove Goal
            </Button>
          ) : (
            <Button size="sm" onClick={saveGoal} className="w-fit">
              Save Goal
            </Button>
          )}
        </ItemActions>
      </ItemContent>
    </Item>
  );
};
