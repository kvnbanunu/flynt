"use client";
import React, { useState, useEffect } from "react";
import {
  Card,
  CardAction,
  CardContent,
  CardHeader,
  CardTitle,
} from "../ui/card";
import { Button } from "../ui/button";
import {
  Collapsible,
  CollapsibleContent,
  CollapsibleTrigger,
} from "../ui/collapsible";
import { ChevronsUpDown } from "lucide-react";
import { Checkbox } from "../ui/checkbox";
import { ButtonGroup } from "../ui/button-group";
import { Get, Put, Delete } from "@/lib/api";
import { CheckFyreRequest, UpdateFyreRequest } from "@/types/req";
import { toast } from "sonner";
import { GoalSection } from "../goals/GoalSection";
import { useAuth } from "@/contexts/AuthContext";
import { Item, ItemActions, ItemContent, ItemTitle } from "../ui/item";
import { Hourglass, Cake } from 'lucide-react';
import { Badge } from "../ui/badge";
import { Button as FyreButton } from "../ui/button";


export const FyreCard: React.FC<{ fyre: Models.Fyre }> = ({ fyre }) => {
  const { checkUser } = useAuth();
  const [isOpen, setIsOpen] = useState<boolean>(false);
  const [isChecked, setIsChecked] = useState<boolean>(fyre.is_checked);
  const [currentFyre, setCurrentFyre] = useState<Models.Fyre>(fyre);
  const [currentStreak, setCurrentStreak] = useState<number>(fyre.streak_count);
  const [activeDays, setActiveDays] = useState<string>(fyre.active_days);
  const [isPrivate, setIsPrivate] = useState<boolean>(fyre.is_private);
  const [error, setError] = useState<string | null>(null);
  const [changes, setChanges] = useState<Set<string>>(new Set<string>());
  const [allCategories, setAllCategories] = useState<Models.Category[]>([]);
  const [categoryId, setCategoryId] = useState<number>(fyre.category_id);

  useEffect(() => {
    const fetchCategories = async () => {
      const res = await Get<Models.Category[]>("/fyre/categories")
      if (res.success) {
        setAllCategories(res.data);
      } else {
        setError(res.error.message);
      }
    };
    fetchCategories();
  }, []);

  const reset = () => {
    setActiveDays(currentFyre.active_days);
    setIsPrivate(currentFyre.is_private);
    const temp = changes;
    temp.clear();
    setChanges(temp);
    setError(null);
  };

  const fetchFyre = async (
    path: string,
    req: UpdateFyreRequest | CheckFyreRequest,
  ) => {
    const res = await Put<Models.Fyre, UpdateFyreRequest | CheckFyreRequest>(
      path,
      req,
    );
    if (res.success) {
      setCurrentFyre(res.data);
      setCurrentStreak(res.data.streak_count); // set new count in case of inconsistency
      setError(null);
      toast("Fyre updated");
    } else {
      setError(res.error.message);
    }
    setIsOpen(false);
  };

  const checkFyre = async (checked: boolean) => {
    const increment = checked;
    // set new streak count prior to fetch for quicker update
    const newStreak: number = checked ? currentStreak + 1 : currentStreak - 1;
    setCurrentStreak(newStreak);
    const req: CheckFyreRequest = { id: fyre.id, increment: increment };
    await fetchFyre("/fyre/check", req);
    checkUser();
  };

  const changePrivate = () => {
    const newChanges = changes;
    newChanges.add("is_private");
    setChanges(newChanges);
  };

  const changeDays = (index: number) => {
    const temp: string = activeDays[index] === "1" ? "0" : "1";
    const result =
      activeDays.slice(0, index) + temp + activeDays.slice(index + 1);
    setActiveDays(result);

    const newChanges = changes;
    newChanges.add("active_days");
    setChanges(newChanges);
  };

  const saveChanges = async () => {
    const req: UpdateFyreRequest = {};
    for (const change of changes) {
      switch (change) {
        case "title":
          req.title = currentFyre.title;
          break;
        case "streak_count":
          req.streak_count = currentFyre.streak_count;
          break;
        // case "bonfyre_id":
        //   req.bonfyre_id = currentFyre.bonfyre_id;
        case "active_days":
          req.active_days = activeDays;
          break;
        case "category_id":
          req.category_id = categoryId;
        case "is_private":
          req.is_private = isPrivate;
          break;
      }
    }
    const newChanges = changes;
    newChanges.clear();
    setChanges(newChanges);
    await fetchFyre(`/fyre/${fyre.id}`, req);
  };

  return (
    <Collapsible open={isOpen} onOpenChange={setIsOpen} className="">
      <Card className="pt-0">
        <CardHeader className="bg-primary text-primary-foreground rounded-t-xl">
          <CardTitle className="pt-3 pb-1">{fyre.title}</CardTitle>
          <CardAction className="pt-0.5">
            <CollapsibleTrigger asChild>
              <Button variant="ghost" size="icon" onClick={reset}>
                <ChevronsUpDown />
                <span className="sr-only">Toggle</span>
              </Button>
            </CollapsibleTrigger>
          </CardAction>
        </CardHeader>
        <CardContent>
          <div className="flex items-center justify-between gap-4">
            <ActiveDays
              active={activeDays}
              isOpen={isOpen}
              onChange={changeDays}
            />

            <DaysRemaining
              streak={currentStreak}
              fyreId={fyre.id}
              onOpen={() => setIsOpen(true)}
              onRemoveGoal={ async () => {
                  const res = await Delete(`/goal/${currentFyre.id}`);
                  if (res.success) {
                    toast.success("Goal removed"); 
                  } else {
                    toast.error(res.error.message);
                  }
                }}
            />      

            <div className="flex gap-4 items-center">
              <h4>{currentStreak} 🔥</h4>
              <Checkbox
                className="mr-2 size-8 rounded-md"
                disabled={isOpen}
                checked={isChecked}
                onCheckedChange={(value) => {
                  const newValue = !!value;
                  setIsChecked(newValue);
                  checkFyre(newValue);
                }}
              />
            </div>
          </div>
          <CollapsibleContent>
            <div className="flex flex-col gap-4 mt-4">
              <GoalSection fyreId={fyre.id} />

              <CategorySelector
                allCategories={allCategories}
                selectedId={categoryId}
                onChange={(id) => {
                  setCategoryId(id);
                  const newChanges = new Set(changes);
                  newChanges.add("category_id");
                  setChanges(newChanges);
                }}
              />

              <Item className="justify-stretch">
                <ItemContent className="flex-row justify-between">
                  <ItemTitle>Private</ItemTitle>
                  <ItemActions>
                    <Checkbox
                      checked={isPrivate}
                      onCheckedChange={(value) => {
                        setIsPrivate(!!value);
                        changePrivate();
                      }}
                    />
                  </ItemActions>
                </ItemContent>
              </Item>
              <div className="grid grid-cols-2 justify-stretch gap-4">
                <Button variant="destructive" size="lg">
                  Remove Fyre
                </Button>
                <Button size="lg" onClick={saveChanges}>
                  Save Changes
                </Button>
              </div>
              {error && <div>error</div>}
            </div>
          </CollapsibleContent>
        </CardContent>
      </Card>
    </Collapsible>
  );
};

interface ActiveDaysProps {
  active: string;
  isOpen: boolean;
  onChange: (index: number) => void;
}

export const ActiveDays: React.FC<ActiveDaysProps> = ({
  active,
  isOpen,
  onChange,
}) => {
  const days = ["S", "M", "T", "W", "T", "F", "S"];

  if (active === "1111111" && !isOpen) {
    return (
      <Button className="w-39 sm:w-46 md:w-60" disabled>
        Everyday
      </Button>
    );
  }

  return (
    <ButtonGroup>
      {days.map((day, index) => (
        <Button
          key={index}
          variant={active[index] === "0" ? "secondary" : "default"}
          className="px-1.5 sm:px-2 md:px-3 border-1"
          disabled={!isOpen}
          onClick={() => onChange(index)}
        >
          {day}
        </Button>
      ))}
    </ButtonGroup>
  );
};

interface CategorySelectorProps {
  allCategories: Models.Category[];
  selectedId: number;
  onChange: (id: number) => void;
}

const CategorySelector: React.FC<CategorySelectorProps> = ({
  allCategories,
  selectedId,
  onChange,
}) => {
  return (
    <div className="flex flex-col gap-2">
      <label className="text-sm font-medium text-muted-foreground">
        Category
      </label>
      <select
        className="border rounded-md p-2"
        value={selectedId ?? ""}
        onChange={(e) => onChange(Number(e.target.value))}
      >
        <option value="">Select a category</option>
        {allCategories.map((cat) => (
          <option key={cat.id} value={cat.id}>
            {cat.name}
          </option>
        ))}
      </select>
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

  console.log(goal.data);

  return 0;
}

interface DaysRemainingProps {
  streak: number;
  fyreId: number; 
  onOpen: () => void;
  onRemoveGoal: () => void;
}

const DaysRemaining: React.FC<DaysRemainingProps> = ({
  streak,
  fyreId,
  onOpen,
  onRemoveGoal,
}) => {
  const [goal, setGoal] = useState<Models.Goal>();
  const [loading, setLoading] = useState(true);
  const FyreBadge: React.FC<{ children: ReactNode }> = ({ children }) => {
    return <Badge className="rounded-full aspect-square">{children}</Badge>;
  };

  useEffect(() => {
    async function fetchGoal() {
      const res = await Get<Models.Goal>(`/goal/${fyreId}`);
      if (res.success) {
        setGoal(res.data);
      } else {
        setError(res.error.message);
      }

      setLoading(false);
    };
    fetchGoal();
  }, [fyreId]);

  if (loading) return <div>Loading...</div>;
  if (!goal) return <div>No goal found.</div>;
  
  const daysLeft = getDaysRemaining(goal, streak);
  return (
    <div>
      {daysLeft !== 0 && (
        <div>         
          <FyreBadge key={null}>
            <Hourglass /> {daysLeft} Days Left
          </FyreBadge>
        </div>
      )}

      {daysLeft === 0 && (
        <div className="mt-2 text-red-500 font-semibold">
          <FyreBadge key={null}>
            <Cake />
          </FyreBadge>
          <FyreButton
            variant="link"
            onClick={() => {
              onOpen()
              onRemoveGoal()
            }}
          >
            Create a new goal?
          </FyreButton>
        </div>
      )}
    </div> 
  );
};
