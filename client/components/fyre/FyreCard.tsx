"use client";

import React, { useState } from "react";
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
import { Put } from "@/lib/api";
import { UpdateFyreRequest } from "@/types/req";
import { toast } from "sonner";

export const FyreCard: React.FC<{ fyre: Models.Fyre }> = ({ fyre }) => {
  const [isOpen, setIsOpen] = useState<boolean>(false);
  const [isChecked, setIsChecked] = useState<boolean>(false);
  const [currentFyre, setCurrentFyre] = useState<Models.Fyre>(fyre);
  const [activeDays, setActiveDays] = useState<string>(fyre.active_days);
  const [error, setError] = useState<string | null>(null);
  const [changes, setChanges] = useState<Set<string>>(new Set<string>());

  const reset = () => {
    setActiveDays(currentFyre.active_days);
    const temp = changes;
    temp.clear();
    setChanges(temp);
  };

  const checkFyre = async (checked: boolean) => {
    const newCount = checked
      ? currentFyre.streak_count + 1
      : currentFyre.streak_count - 1;
    const req: UpdateFyreRequest = { streak_count: newCount };
    await fetchFyre(req);
  };

  const fetchFyre = async (req: UpdateFyreRequest) => {
    const res = await Put<Models.Fyre, UpdateFyreRequest>(
      `/fyre/${fyre.id}`,
      req,
    );
    if (res.success) {
      setCurrentFyre(res.data);
      setError(null);
      toast("Fyre updated");
    } else {
      setError(res.error.message);
    }
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
      }
    }
    await fetchFyre(req);
    const newChanges = changes;
    newChanges.clear();
    setChanges(newChanges);
  };

  return (
    <Collapsible open={isOpen} onOpenChange={setIsOpen} className="">
      <Card className="mb-4 pt-0">
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
            <h4>Streak: {currentFyre.streak_count} 🔥</h4>
            <Checkbox
              className="mr-2 size-8 rounded-md"
              checked={isChecked}
              onCheckedChange={(value) => {
                const newValue = !!value;
                setIsChecked(newValue);
                checkFyre(newValue);
              }}
            />
          </div>
          <CollapsibleContent>
            <div className="flex flex-col gap-4 mt-4">
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

const ActiveDays: React.FC<ActiveDaysProps> = ({
  active,
  isOpen,
  onChange,
}) => {
  const days = ["S", "M", "T", "W", "T", "F", "S"];

  if (active === "1111111" && !isOpen) {
    return (
      <Button className="w-44" disabled>
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
          className="px-2 border-1"
          disabled={!isOpen}
          onClick={() => onChange(index)}
        >
          {day}
        </Button>
      ))}
    </ButtonGroup>
  );
};
