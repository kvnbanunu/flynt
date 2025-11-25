"use client";

import { useAuth } from "@/contexts/AuthContext";
import React, { useState } from "react";
import {
  Card,
  CardAction,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "../ui/card";
import { Button } from "../ui/button";
import * as z from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { Controller, useForm } from "react-hook-form";
import { FieldGroup, Field, FieldLabel, FieldError } from "../ui/field";
import { Input } from "../ui/input";
import { Spinner } from "../ui/spinner";
import { UpdateUserRequest } from "@/types/req";
import { Put } from "@/lib/api";
import { Textarea } from "../ui/textarea";
import { toast } from "sonner";
import { ScrollArea } from "../ui/scroll-area";
import { TimezoneSelector } from "../timezone/TimezoneSelector";

const editSchema = z.object({
  name: z.string().min(5).max(50).optional().or(z.literal("")),
  current_password: z
    .string()
    .trim()
    .min(8)
    .max(50)
    .optional()
    .or(z.literal("")),
  new_password: z.string().trim().min(8).max(50).optional().or(z.literal("")),
  email: z.email().toLowerCase().optional().or(z.literal("")),
  bio: z.string().min(0).max(500).optional().or(z.literal("")),
  timezone: z.string().trim().min(5).max(50).optional().or(z.literal("")),
});

export const Profile: React.FC = () => {
  const { user, setUser, logout, isAuthenticated } = useAuth();
  const [loading, setLoading] = useState<boolean>(false);

  const form = useForm<z.infer<typeof editSchema>>({
    resolver: zodResolver(editSchema),
    mode: "onSubmit",
    defaultValues: {
      name: user?.name ?? "",
      current_password: "",
      new_password: "",
      email: user?.email ?? "",
      bio: user?.bio ?? "",
      timezone: user?.timezone ?? "",
    },
  });

  if (user === null || !isAuthenticated) {
    return;
  }

  const onSubmit = async (data: z.infer<typeof editSchema>) => {
    if (!user) return;
    setLoading(true);
    const req: UpdateUserRequest = data;

    req.name = req.name === user.name ? undefined : req.name;
    req.current_password =
      req.current_password === "" ? undefined : req.current_password;
    req.new_password = req.new_password === "" ? undefined : req.new_password;
    req.email = req.email === user.email ? undefined : req.email;
    req.bio = req.bio === user.bio ? undefined : req.bio;
    req.timezone = req.timezone === user.timezone ? undefined : req.timezone;

    const res = await Put<Models.User, UpdateUserRequest>("/user", req);
    if (res.success) {
      setUser(res.data);
      toast("Successfully updated profile!");
    } else {
      toast(res.error.message);
    }
    setLoading(false);
  };

  if (user && isAuthenticated) {
    return (
      <ScrollArea className="w-full">
        <Card className="w-full">
          <CardHeader>
            <CardTitle className="text-3xl">Profile</CardTitle>
            <CardAction>
              <Button variant="default" onClick={logout}>
                Logout
              </Button>
            </CardAction>
          </CardHeader>
          <CardContent>
            <form id="form-profile" onSubmit={form.handleSubmit(onSubmit)}>
              <Field className="mb-6">
                <FieldLabel htmlFor="form-profile-username">
                  Username
                </FieldLabel>
                <Input
                  id="form-profile-username"
                  placeholder={user.username}
                  autoComplete="off"
                  readOnly
                  className="pointer-events-none"
                />
              </Field>
              <FieldGroup>
                <Controller
                  name="name"
                  control={form.control}
                  render={({ field, fieldState }) => (
                    <Field data-invalid={fieldState.invalid}>
                      <FieldLabel htmlFor="form-profile-name">Name</FieldLabel>
                      <Input
                        {...field}
                        id="form-profile-name"
                        aria-invalid={fieldState.invalid}
                        placeholder={user.name}
                        autoComplete="off"
                      />
                      {fieldState.invalid && (
                        <FieldError errors={[fieldState.error]} />
                      )}
                    </Field>
                  )}
                />
                <Controller
                  name="current_password"
                  control={form.control}
                  render={({ field, fieldState }) => (
                    <Field data-invalid={fieldState.invalid}>
                      <FieldLabel htmlFor="form-profile-current-password">
                        Current Password
                      </FieldLabel>
                      <Input
                        {...field}
                        id="form-profile-current-password"
                        type="password"
                        aria-invalid={fieldState.invalid}
                        placeholder="********"
                        autoComplete="off"
                      />
                      {fieldState.invalid && (
                        <FieldError errors={[fieldState.error]} />
                      )}
                    </Field>
                  )}
                />
                <Controller
                  name="new_password"
                  control={form.control}
                  render={({ field, fieldState }) => (
                    <Field data-invalid={fieldState.invalid}>
                      <FieldLabel htmlFor="form-profile-new-password">
                        New Password
                      </FieldLabel>
                      <Input
                        {...field}
                        id="form-profile-new-password"
                        type="password"
                        aria-invalid={fieldState.invalid}
                        placeholder="********"
                        autoComplete="off"
                      />
                      {fieldState.invalid && (
                        <FieldError errors={[fieldState.error]} />
                      )}
                    </Field>
                  )}
                />
                <Controller
                  name="email"
                  control={form.control}
                  render={({ field, fieldState }) => (
                    <Field data-invalid={fieldState.invalid}>
                      <FieldLabel htmlFor="form-profile-email">
                        Email
                      </FieldLabel>
                      <Input
                        {...field}
                        id="form-profile-email"
                        type="email"
                        aria-invalid={fieldState.invalid}
                        placeholder={user.email}
                        autoComplete="off"
                      />
                      {fieldState.invalid && (
                        <FieldError errors={[fieldState.error]} />
                      )}
                    </Field>
                  )}
                />
                <Controller
                  name="bio"
                  control={form.control}
                  render={({ field, fieldState }) => (
                    <Field data-invalid={fieldState.invalid}>
                      <FieldLabel htmlFor="form-profile-bio">Bio</FieldLabel>
                      <Textarea
                        {...field}
                        id="form-profile-bio"
                        aria-invalid={fieldState.invalid}
                        placeholder={user.bio ? user.bio : ""}
                        autoComplete="off"
                      />
                      {fieldState.invalid && (
                        <FieldError errors={[fieldState.error]} />
                      )}
                    </Field>
                  )}
                />
                <Controller
                  name="timezone"
                  control={form.control}
                  render={({ field, fieldState }) => (
                    <Field data-invalid={fieldState.invalid}>
                      <FieldLabel htmlFor="form-profile-timezone">
                        Timezone
                      </FieldLabel>
                      <TimezoneSelector
                        value={field.value}
                        onValueChange={field.onChange}
                        placeholder={user.timezone}
                      />
                      {fieldState.invalid && (
                        <FieldError errors={[fieldState.error]} />
                      )}
                    </Field>
                  )}
                />
              </FieldGroup>
            </form>
          </CardContent>
          <CardFooter>
            <Field orientation="horizontal">
              <Button
                type="reset"
                variant="secondary"
                onClick={() => form.reset()}
              >
                Reset
              </Button>
              <Button type="submit" form="form-profile">
                {loading && <Spinner />}
                Save
              </Button>
            </Field>
          </CardFooter>
        </Card>
      </ScrollArea>
    );
  }
};
