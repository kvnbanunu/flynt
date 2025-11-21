"use client";

import React, { useState } from "react";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "../ui/dialog";
import { Button } from "../ui/button";
import z from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { Post } from "@/lib/api";
import { Controller, useForm } from "react-hook-form";
import { Spinner } from "../ui/spinner";
import { Field, FieldError, FieldGroup, FieldLabel } from "../ui/field";
import { Input } from "../ui/input";
import { ButtonGroup } from "../ui/button-group";
import { useAuth } from "@/contexts/AuthContext";

const addSchema = z.object({
  title: z.string().trim().min(5).max(50),
  streak_count: z.number().int().nonnegative(),
  active_days: z.string().length(7),
});

export const AddFyre: React.FC = () => {
  const { checkUser, fetchFyres } = useAuth();
  const [loading, setLoading] = useState<boolean>(false);
  const [open, setOpen] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  const form = useForm<z.infer<typeof addSchema>>({
    resolver: zodResolver(addSchema),
    defaultValues: {
      title: "",
      streak_count: 0,
      active_days: "1111111",
    },
  });

  const onSubmit = async (data: z.infer<typeof addSchema>) => {
    setLoading(true);
    const res = await Post<Models.Fyre, z.infer<typeof addSchema>>(
      "/fyre",
      data,
    );
    if (res.success) {
      setError(null);
      checkUser();
      fetchFyres();
      setOpen(false);
    } else {
      setError(res.error.message);
    }
    setLoading(false);
  };

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <form id="form-addfyre" onSubmit={form.handleSubmit(onSubmit)}>
        <DialogTrigger asChild>
          <Button size="lg" className="w-full">
            Add New Fyre
          </Button>
        </DialogTrigger>
        <DialogContent className="sm:max-w-[425px]">
          <DialogHeader>
            <DialogTitle>Add New Fyre</DialogTitle>
            <DialogDescription>
              This is the start to your new streak!
            </DialogDescription>
          </DialogHeader>
          {error && <div>error</div>}
          <FieldGroup>
            <Controller
              name="title"
              control={form.control}
              render={({ field, fieldState }) => (
                <Field data-invalid={fieldState.invalid}>
                  <FieldLabel htmlFor="form-addfyre-title">Title</FieldLabel>
                  <Input
                    {...field}
                    id="form-addfyre-title"
                    aria-invalid={fieldState.invalid}
                    placeholder="Run a mile"
                    autoComplete="off"
                  />
                  {fieldState.invalid && (
                    <FieldError errors={[fieldState.error]} />
                  )}
                </Field>
              )}
            />
            <Controller
              name="streak_count"
              control={form.control}
              render={({ field, fieldState }) => (
                <Field data-invalid={fieldState.invalid}>
                  <FieldLabel htmlFor="form-addfyre-streak-count">
                    Streak Count
                  </FieldLabel>
                  <Input
                    {...field}
                    id="form-addfyre-streak-count"
                    aria-invalid={fieldState.invalid}
                    placeholder="0"
                    autoComplete="off"
                    type="number"
                    onChange={(e) => field.onChange(e.target.valueAsNumber || 0)}
                  />
                  {fieldState.invalid && (
                    <FieldError errors={[fieldState.error]} />
                  )}
                </Field>
              )}
            />
            <Controller
              name="active_days"
              control={form.control}
              render={({ field, fieldState }) => (
                <Field data-invalid={fieldState.invalid}>
                  <FieldLabel htmlFor="form-addfyre-days-active">
                    Days Active
                  </FieldLabel>

                  <ButtonGroup className="justify-center w-full">
                    {["S", "M", "T", "W", "T", "F", "S"].map((day, index) => (
                      <Button
                        key={index}
                        variant={
                          (field.value || "1111111")[index] === "0"
                            ? "secondary"
                            : "default"
                        }
                        className="px-5.5 border-1"
                        onClick={() => {
                          const currentValue = field.value || "1111111";
                          const valueArray = currentValue.split("");
                          valueArray[index] =
                            valueArray[index] === "0" ? "1" : "0";
                          field.onChange(valueArray.join(""));
                        }}
                      >
                        {day}
                      </Button>
                    ))}
                  </ButtonGroup>
                  {fieldState.invalid && (
                    <FieldError errors={[fieldState.error]} />
                  )}
                </Field>
              )}
            />
          </FieldGroup>
          <DialogFooter>
            <DialogClose asChild>
              <Button variant="secondary">Cancel</Button>
            </DialogClose>
            <Button type="submit" form="form-addfyre">
              {loading && <Spinner />}Save
            </Button>
          </DialogFooter>
        </DialogContent>
      </form>
    </Dialog>
  );
};
