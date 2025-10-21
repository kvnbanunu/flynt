"use client"

import { Delete } from "@/lib/api";
import React, { useState } from "react"

interface RemoveFyreProps {
  onSuccessHandler: () => void;
  fyre_id: number;
}

export const RemoveFyre: React.FC<RemoveFyreProps> = (props) => {
  const {onSuccessHandler, fyre_id} = props

  const [error, setError] = useState<string | null>(null)

  const handleSubmit = async () => {
    const res = await Delete(`/fyre/${fyre_id}`)

    if (res.success) {
      onSuccessHandler()
    } else {
      setError(res.error.message)
    }
  }

  return (
    <div>
      {error && <div className="bg-red-100">{error}</div>}
      <button
        className="rounded-xl px-2 py-1 m-2 text-center bg-red-300 cursor-pointer"
        onClick={handleSubmit}
      >
        Remove
      </button>
    </div>
  )
}
