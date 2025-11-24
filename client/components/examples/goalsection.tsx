"use client";

import React, { useEffect, useState } from "react";
import { Card, CardContent, CardHeader, CardTitle } from "../ui/card";
import { Button } from "../ui/button";
import { Input } from "../ui/input";
import { Select, SelectTrigger, SelectValue, SelectContent, SelectItem } from "../ui/select";
import { Get, Post, Delete } from "@/lib/api";
import { toast } from "sonner";
import type { CreateGoalRequest, DeleteGoalRequest } from "@/types/req";

export const GoalSection: React.FC<{ fyreId: number }> = ({ fyreId }) => {
  const [goal, setGoal] = useState<Models.Goal | null>(null);
  const [loading, setLoading] = useState(true);
  const [goalType, setGoalType] = useState<string>("");
  const [goalData, setGoalData] = useState<string>("");
  const [error, setError] = useState<string | null>(null);

  const fetchGoal = async () => {
  setLoading(true);
  try {
    const res = await Get<Models.Goal>(`/fyre/${fyreId}/goal`);
    if (res.success) {
      setGoal(res.data);
      setGoalType(res.data.goal_type_id.toString());
      setGoalData(res.data.data ?? "");
    } else if (res.error.message.includes("not found")) {
      setGoal(null);
    } else {
      setError(res.error.message);
    }
    } catch (err: any) {
      setError(err.message || "Failed to fetch goal");
    } finally {
      setLoading(false);
    }
  }; 

  const saveGoal = async () => {
    if (!goalType || !goalData) {
      toast.error("Please select a goal type and enter a value");
      return;
    }

    const req: CreateGoalRequest = {
      fyre_id: fyreId,
      goal_type_id: parseInt(goalType),
      description:
        goalType == "1"
          ? "Reach target date"
          : "Reach target streak count",
      data: goalData,
    };

    const res = await Post<Models.Goal, any>(`/fyre/${fyreId}/goal`, req);
    if (res.success) {
      toast.success("Goal saved!");
      setGoal(res.data);
    } else {
      toast.error(res.error.message);
    }
  };

  const deleteGoal = async () => {
    const res = await Delete(`/fyre/${fyreId}/goal`);
    if (res.success) {
      toast.success("Goal removed");
      setGoal(null);
      setGoalData("");
      setGoalType("");
    } else {
      toast.error(res.error.message);
    }
  };

  useEffect(() => {
    fetchGoal();
  }, [fyreId]);

  if (loading) return <div>Loading goal...</div>;
  if (error) return <div className="text-red-500">{error}</div>;

  return (
    <Card className="border">
      <CardHeader>
        <CardTitle className="text-md">Goal</CardTitle>
      </CardHeader>
      <CardContent>
        {goal ? (
          <div className="flex flex-col gap-3">
            <p>
              <strong>Type:</strong> {goal.goal_type_id === 1 ? "Until Date" : "Number of Days"}
            </p>
            <p>
              <strong>Value:</strong>{" "}
              {goal.goal_type_id === 1
                ? goal.data ? new Date(goal.data).toLocaleDateString() : "-"
                : goal.data ?? "-"} days
            </p>
            <Button
              variant="destructive"
              size="sm"
              onClick={deleteGoal}
              className="w-fit"
            >
              Remove Goal
            </Button>
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

            <Button size="sm" onClick={saveGoal} className="w-fit">
              Save Goal
            </Button>
          </div>
        )}
      </CardContent>
    </Card>
  );
};
