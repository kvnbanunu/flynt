"use client";
import React, { useState, ReactNode } from "react";
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
import { Delete, Get, Put } from "@/lib/api";
import {
  BonfyreMember,
  CheckFyreRequest,
  UpdateFyreRequest,
} from "@/types/req";
import { toast } from "sonner";
import { GoalSection } from "../goals/GoalSection";
import { useAuth } from "@/contexts/AuthContext";
import { Item, ItemActions, ItemContent, ItemTitle } from "../ui/item";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "../ui/table";
import { ActiveDays } from "./ActiveDays";
import { FyreBadges } from "./FyreBadges";
import { CategorySelector } from "./CategorySelector";

export const FyreCard: React.FC<{
  fyre: Models.Fyre;
  goals?: Models.Goal[];
}> = ({ fyre, goals }) => {
  const { checkUser, categories } = useAuth();
  const [isOpen, setIsOpen] = useState<boolean>(false);
  const [isChecked, setIsChecked] = useState<boolean>(fyre.is_checked);
  const [currentFyre, setCurrentFyre] = useState<Models.Fyre>(fyre);
  const [currentStreak, setCurrentStreak] = useState<number>(fyre.streak_count);
  const [activeDays, setActiveDays] = useState<string>(fyre.active_days);
  const [isPrivate, setIsPrivate] = useState<boolean>(fyre.is_private);
  const [error, setError] = useState<string | null>(null);
  const [changes, setChanges] = useState<Set<string>>(new Set<string>());
  const [categoryId, setCategoryId] = useState<number>(fyre.category_id);
  const [bonfyreMembers, setBonfyreMembers] = useState<BonfyreMember[]>([]);

  const reset = () => {
    setActiveDays(currentFyre.active_days);
    setIsPrivate(currentFyre.is_private);
    const temp = changes;
    temp.clear();
    setChanges(temp);
    setError(null);
  };

  const fetchBonfyreMembers = async () => {
    if (fyre.bonfyre_id) {
      const res = await Get<BonfyreMember[]>(
        `/fyre/bonfyre/${fyre.bonfyre_id}`,
      );
      if (res.success) {
        setBonfyreMembers(res.data);
      } else {
        toast(res.error.message);
      }
    }
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

  const removeFyre = async () => {
    const res = await Delete(`/fyre/${fyre.id}`);
    if (res.success) {
      checkUser();
      toast("Fyre removed");
    } else {
      toast(res.error.message);
    }
  }

  const onOpen = async () => {
    if (isOpen) {
      setIsOpen(false);
      return;
    }
    await fetchBonfyreMembers();
    setIsOpen(true);
  };

  return (
    <Collapsible open={isOpen} onOpenChange={onOpen} className="">
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
          <div className="flex items-center justify-between">
            <div className="flex flex-wrap gap-4 items-center">
              <ActiveDays
                active={activeDays}
                isOpen={isOpen}
                onChange={changeDays}
              />
              <FyreBadges fyre={fyre} goals={goals} onOpen={onOpen} />
            </div>
            <div className="flex flex-wrap gap-4 items-center">
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
            <div className="flex flex-col mt-4">
              <GoalSection fyreId={fyre.id} goal={goals && goals[0]} />

              {fyre.bonfyre_id && (
                <BonfyreCard
                  total={fyre.bonfyre_total ?? 0}
                  members={bonfyreMembers}
                />
              )}

              <FyreEditField title="Category">
                <CategorySelector
                  allCategories={categories}
                  selectedId={categoryId}
                  onChange={(id) => {
                    setCategoryId(id);
                    const newChanges = new Set(changes);
                    newChanges.add("category_id");
                    setChanges(newChanges);
                  }}
                />
              </FyreEditField>

              <FyreEditField title="Private">
                <Checkbox
                  checked={isPrivate}
                  onCheckedChange={(value) => {
                    setIsPrivate(!!value);
                    changePrivate();
                  }}
                />
              </FyreEditField>
              <div className="grid grid-cols-2 justify-stretch gap-4">
                <Button variant="destructive" size="lg" onClick={removeFyre}>
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

const FyreEditField: React.FC<{ title: string; children?: ReactNode }> = ({
  title,
  children,
}) => {
  return (
    <Item className="justify-stretch">
      <ItemContent className="flex flex-row justify-between">
        <ItemTitle>{title}</ItemTitle>
        <ItemActions>{children}</ItemActions>
      </ItemContent>
    </Item>
  );
};

const BonfyreCard: React.FC<{ total: number; members: BonfyreMember[] }> = ({
  total,
  members,
}) => {
  return (
    <Item variant="outline" className="mt-4">
      <ItemContent>
        <div className="flex flex-row justify-between">
          <ItemTitle className="text-lg">Bonfyre Members</ItemTitle>
        </div>
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Rank</TableHead>
              <TableHead>Username</TableHead>
              <TableHead className="text-right">Total: {total} 🔥</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {members.map((m, index) => (
              <TableRow key={index}>
                <TableCell>{index + 1}</TableCell>
                <TableCell>{m.username}</TableCell>
                <TableCell className="text-right">
                  {m.streak_count} 🔥
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </ItemContent>
    </Item>
  );
};
