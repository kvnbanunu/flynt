"use client"

import { Delete } from "@/lib/api";
import { CS_ENV } from "@/lib/utils";
import React, { useState } from "react"

interface RemoveFyreProps {
  onSuccessHandler: () => void;
  fyre_id: number;
}

export const RemoveFyre: React.FC<RemoveFyreProps> = (props) => {
  const {onSuccessHandler, fyre_id} = props

  const [error, setError] = useState<string | null>(null)

  const handleSubmit = async () => {
    const res = await Delete(`${CS_ENV.api_url}/api/fyre/${fyre_id}`)

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
        className="rounded-xl p-2 m-2 text-center bg-red-300"
        onClick={handleSubmit}
      >
        Remove
      </button>
    </div>
  )
}
