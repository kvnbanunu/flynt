import React from "react"
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"
import { describe } from "node:test";

export const Popup: React.FC<{
  description: string;
}> = async (props) => {

  const { description } = props

  return (
    <Dialog>
      <DialogTrigger>Open</DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Are you absolutely sure?</DialogTitle>
          <DialogDescription>
            { description }
          </DialogDescription>
        </DialogHeader>
      </DialogContent>
    </Dialog>
  );
}