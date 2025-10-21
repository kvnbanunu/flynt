"use client";

import React from "react";
import * as z from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { Controller, useForm } from "react-hook-form";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "../ui/card";
import { Field, FieldError, FieldGroup, FieldLabel } from "../ui/field";
import { Input } from "../ui/input";
import { useAuth } from "@/contexts/AuthContext";
import { RegisterRequest } from "@/types/req";
import { Button } from "../ui/button";
import { Spinner } from "../ui/spinner";

const registerSchema = z.object({
  username: z.string().trim().min(5).max(50),
  name: z.string().min(5).max(50),
  password: z.string().trim().min(8).max(50),
  email: z.email().toLowerCase(),
});

export const RegisterForm: React.FC = () => {
  const { loading, register } = useAuth();

  const form = useForm<z.infer<typeof registerSchema>>({
    resolver: zodResolver(registerSchema),
    mode: "onTouched",
    defaultValues: {
      username: "",
      name: "",
      password: "",
      email: "",
    },
  });

  const onSubmit = async (data: z.infer<typeof registerSchema>) => {
    const req: RegisterRequest = data;

    await register(req);
  };

  return (
    <Card>
      <CardHeader>
        <CardTitle>Register</CardTitle>
      </CardHeader>
      <CardContent>
        <form id="form-register" onSubmit={form.handleSubmit(onSubmit)}>
          <FieldGroup>
            <Controller
              name="username"
              control={form.control}
              render={({ field, fieldState }) => (
                <Field data-invalid={fieldState.invalid}>
                  <FieldLabel htmlFor="form-register-username">
                    Username
                  </FieldLabel>
                  <Input
                    {...field}
                    id="form-register-username"
                    aria-invalid={fieldState.invalid}
                    placeholder="Flynt"
                    autoComplete="off"
                  />
                  {fieldState.invalid && (
                    <FieldError errors={[fieldState.error]} />
                  )}
                </Field>
              )}
            />
            <Controller
              name="name"
              control={form.control}
              render={({ field, fieldState }) => (
                <Field data-invalid={fieldState.invalid}>
                  <FieldLabel htmlFor="form-register-name">Name</FieldLabel>
                  <Input
                    {...field}
                    id="form-register-name"
                    aria-invalid={fieldState.invalid}
                    placeholder="Your Name"
                    autoComplete="off"
                  />
                  {fieldState.invalid && (
                    <FieldError errors={[fieldState.error]} />
                  )}
                </Field>
              )}
            />
            <Controller
              name="password"
              control={form.control}
              render={({ field, fieldState }) => (
                <Field data-invalid={fieldState.invalid}>
                  <FieldLabel htmlFor="form-register-password">
                    Password
                  </FieldLabel>
                  <Input
                    {...field}
                    id="form-register-password"
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
                  <FieldLabel htmlFor="form-register-email">Email</FieldLabel>
                  <Input
                    {...field}
                    id="form-register-email"
                    aria-invalid={fieldState.invalid}
                    placeholder="example@flyntapp.io"
                    autoComplete="off"
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
            type="button"
            variant="secondary"
            onClick={() => form.reset()}
          >
            Reset
          </Button>
          <Button type="submit" form="form-register">
            {loading && <Spinner />}
            Submit
          </Button>
        </Field>
      </CardFooter>
    </Card>
  );
};
