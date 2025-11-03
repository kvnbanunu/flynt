"use client";

import { Popup } from "@/components/popup/Popup";
import { RemoveFyre } from "@/components/examples/removefyre" 

export default function Temp() {
  return (
    <div>
      pop
      <Popup 
        description="This is an example"
        button = {<RemoveFyre fyre_id={1} onSuccessHandler={() => {alert("BUTTON WORKS")}}/>}
      />
    </div>
  );
}
