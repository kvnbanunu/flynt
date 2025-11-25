"use client";
import React, { ReactNode } from "react";
import { Button } from "../ui/button";
import { Calendar, FlameKindling, Tally5 } from "lucide-react";
import { Delete } from "@/lib/api";
import { toast } from "sonner";
import { useAuth } from "@/contexts/AuthContext";
import { Hourglass, Cake } from "lucide-react";
import { Badge } from "../ui/badge";

export const FyreBadge: React.FC<{ children: ReactNode }> = ({ children }) => {
  return <Badge className="rounded-full aspect-square">{children}</Badge>;
};

interface FyreBadgeProps {
  fyre: Models.Fyre;
  goals?: Models.Goal[];
  onOpen: () => void;
}

export const FyreBadges: React.FC<FyreBadgeProps> = ({
  fyre,
  goals,
  onOpen,
}) => {
  return (
    <div className="flex gap-1 justify-start items-center">
      {fyre.bonfyre_id && (
        <FyreBadge key={fyre.bonfyre_id}>
          <FlameKindling />
        </FyreBadge>
      )}
      {goals &&
        goals.map((goal) => (
          <FyreBadge key={goal.goal_type_id}>
            {goal.goal_type_id === 1 ? <Calendar /> : <Tally5 />}
          </FyreBadge>
        ))}
      {goals && goals.length > 0 && (
        <DaysRemaining fyre={fyre} goal={goals[0]} onOpen={onOpen} />
      )}
    </div>
  );
};

function getDaysRemaining(goal: Models.Goal, streak: number): number {
  if (!goal || !goal.data) return 0;

  // DATE GOAL
  if (goal.goal_type_id === 1) {
    const today = new Date();
    const target = new Date(goal.data);

    const diff = target.getTime() - today.getTime();
    return Math.max(Math.ceil(diff / (1000 * 60 * 60 * 24)), 0);
  }

  // DAYS GOAL
  if (goal.goal_type_id === 2) {
    const goalDays = parseInt(goal.data, 10);
    return Math.max(goalDays - streak, 0);
  }

  return 0;
}

interface DaysRemainingProps {
  fyre: Models.Fyre;
  goal: Models.Goal;
  onOpen: () => void;
}

export const DaysRemaining: React.FC<DaysRemainingProps> = ({
  fyre,
  goal,
  onOpen,
}) => {
  const { fetchFyres } = useAuth();
  const daysLeft = getDaysRemaining(goal, fyre.streak_count);

  const removeGoal = async () => {
    const res = await Delete(`/goal/${fyre.id}`);
    if (res.success) {
      toast.success("Goal removed");
      fetchFyres();
    } else {
      toast.error(res.error.message);
    }
  };

  const preOpen = async () => {
    await removeGoal();
    onOpen();
  };

  return (
    <div>
      {daysLeft !== 0 && (
        <div>
          <Badge className="rounded-full py-1.5">
            <Hourglass /> {daysLeft} Days Left
          </Badge>
        </div>
      )}

      {daysLeft === 0 && (
        <div className="flex items-center">
          <FyreBadge>
            <Cake className="text-red-500 font-semibold" />
          </FyreBadge>
          <Button
            variant="link"
            className="text-red-500 font-semibold"
            onClick={preOpen}
          >
            Goal completed!
            <br />
            Create a new goal?
          </Button>
        </div>
      )}
    </div>
  );
};
