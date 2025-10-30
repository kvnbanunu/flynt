"use client";
import { useEffect, useRef } from "react";

export default function FlameAnimation() {
  const canvasRef = useRef<HTMLCanvasElement | null>(null);

  useEffect(() => {
    const canvas = canvasRef.current;
    if (!canvas) return;
    const ctx = canvas.getContext("2d");
    if (!ctx) return;

    ctx.translate(0, canvas.height);
    ctx.scale(23, -23);

    let fps = 7;
    let interval = 1000 / fps;
    let prev = Date.now();

    const y = [2, 1, 0, 0, 0, 0, 1, 2];
    const max = [7, 9, 11, 13, 13, 11, 9, 7];
    const min = [4, 7, 8, 10, 10, 8, 7, 4];

    function flame() {
      const now = Date.now();
      const dif = now - prev;

      if (dif > interval) {
        if (!ctx || !canvas) return;

        prev = now;
        ctx.clearRect(0, 0, canvas.width, canvas.height);

        ctx.strokeStyle = "#d14234";
        let i = 0;
        for (let x = 4; x < 12; x++) {
          const a = Math.random() * (max[i] - min[i] + 1) + min[i];
          ctx.beginPath();
          ctx.moveTo(x + 0.5, y[i++]);
          ctx.lineTo(x + 0.5, a);
          ctx.stroke();
        }

        ctx.strokeStyle = "#f2a55f";
        let j = 1;
        for (let x = 5; x < 11; x++) {
          const a =
            Math.random() * (max[j] - 5 - (min[j] - 5) + 1) + (min[j] - 5);
          ctx.beginPath();
          ctx.moveTo(x + 0.5, y[j++] + 1);
          ctx.lineTo(x + 0.5, a);
          ctx.stroke();
        }

        ctx.strokeStyle = "#e8dec5";
        let k = 3;
        for (let x = 7; x < 9; x++) {
          const a =
            Math.random() * (max[k] - 9 - (min[k] - 9) + 1) + (min[k] - 9);
          ctx.beginPath();
          ctx.moveTo(x + 0.5, y[k++]);
          ctx.lineTo(x + 0.5, a);
          ctx.stroke();
        }
      }
      requestAnimationFrame(flame);
    }

    flame();
  }, []);

  return (
    <canvas
      ref={canvasRef}
      id="c"
      width={360}
      height={360}
      className="w-55 h-55"
    />
  );
}
