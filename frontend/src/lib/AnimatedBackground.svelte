<script lang="ts">
    import { onMount, onDestroy } from "svelte";

    let canvas: HTMLCanvasElement;
    const whiteColors = [
        "--svg-white-01",
        "--svg-white-02",
        "--svg-white-03",
        "--svg-white-04",
        "--svg-white-05",
        "--svg-white-06",
    ];

    const blueColor = "--svg-blue";

    let intervalId: number;
    let animationId: number;
    const transitionDuration = 2000; // 2 seconds

    const maxCount = 4800;
    const squarePerLine = 80;
    const rectWidth = 6;
    const rectHeight = 6;
    const initialX = 3;
    const initialY = 3;
    const offset = 10;

    let currentColors: { r: number; g: number; b: number }[] = [];
    let targetColors: { r: number; g: number; b: number }[] = [];
    let startTime: number;

    const interpolateColor = (
        start: { r: number; g: number; b: number },
        end: { r: number; g: number; b: number },
        factor: number,
    ) => {
        return {
            r: Math.round(start.r + (end.r - start.r) * factor),
            g: Math.round(start.g + (end.g - start.g) * factor),
            b: Math.round(start.b + (end.b - start.b) * factor),
        };
    };

    const parseRgbString = (rgbString: string) => {
        const [r, g, b] = rgbString
            .replace(/[^\d,]/g, "") // Remove non-numeric characters
            .split(",")
            .map(Number);
        return { r, g, b };
    };

    const getRandomColor = () => {
        const probability = Math.random();
        if (probability < 0.01) { // 5% chance for blue
            const resolved = getComputedStyle(document.documentElement)
                .getPropertyValue(blueColor)
                .trim();
            return parseRgbString(resolved);
        } else { // 95% chance for white
            const colorVar = whiteColors[Math.floor(Math.random() * whiteColors.length)];
            const resolved = getComputedStyle(document.documentElement)
                .getPropertyValue(colorVar)
                .trim();
            return parseRgbString(resolved);
        }
    };

    const drawSquares = (
        context: CanvasRenderingContext2D,
        colorsArray: { r: number; g: number; b: number }[],
    ) => {
        let count = 0;

        // Clear the canvas before drawing
        context.clearRect(0, 0, canvas.width, canvas.height);

        // Draw multiple rectangles with interpolated colors
        while (count < maxCount) {
            const rectX = initialX + (count % squarePerLine) * offset;
            const rectY = initialY + Math.floor(count / squarePerLine) * offset;

            const color = colorsArray[count];
            context.fillStyle = `rgb(${color.r}, ${color.g}, ${color.b})`;

            context.fillRect(rectX, rectY, rectWidth, rectHeight);
            count++;
        }
    };

    const startTransition = (context: CanvasRenderingContext2D) => {
        // Assign a random target color for each square
        targetColors = currentColors.map(() => getRandomColor());

        const startColors = currentColors;
        const endColors = targetColors;

        startTime = performance.now();

        const animate = (time: number) => {
            const elapsed = time - startTime;
            const factor = Math.min(elapsed / transitionDuration, 1);

            const interpolatedColors = startColors.map((start, index) => {
                const end = endColors[index];
                return interpolateColor(start, end, factor);
            });

            drawSquares(context, interpolatedColors);

            if (factor < 1) {
                animationId = requestAnimationFrame(animate);
            } else {
                currentColors = targetColors;
            }
        };

        animationId = requestAnimationFrame(animate);
    };

    onMount(() => {
        if (canvas) {
            const context = canvas.getContext("2d");
            if (!context) {
                console.error("Canvas context not available");
                return;
            }

            // Initialize currentColors with random colors
            currentColors = Array.from({ length: maxCount }, () => getRandomColor());

            drawSquares(context, currentColors);

            // Update colors every 10 seconds with transition
            intervalId = window.setInterval(() => {
                startTransition(context);
            }, 10000);
        }
    });

    onDestroy(() => {
        if (intervalId) {
            clearInterval(intervalId);
        }
        if (animationId) {
            cancelAnimationFrame(animationId);
        }
    });
</script>

<canvas bind:this={canvas} width="800" height="600"></canvas>

<style>
    canvas {
        position: fixed;
        width: 100%;
        height: auto;
        background-color: var(--svg-white);
        z-index: -1;
    }
</style>