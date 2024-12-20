<script lang="ts">
    import { onMount } from "svelte";
	import Nav from "./lib/Nav.svelte";
    import AnimatedBackground from "./lib/AnimatedBackground.svelte";

	const colors = [
    "--svg-white-01",
    "--svg-white-02",
    "--svg-white-03",
    "--svg-white-04",
    "--svg-white-05",
    "--svg-white-06",
    "--svg-blue",
    "--svg-white"
  ];

  // Utility function to shuffle an array
  function shuffle(array: string[]) {
    for (let i = array.length - 1; i > 0; i--) {
      const j = Math.floor(Math.random() * (i + 1));
      [array[i], array[j]] = [array[j], array[i]];
    }
    return array;
  }

//   onMount(() => {
//     const squares = document.querySelectorAll(".square");

//     squares.forEach((square) => {
//       // Shuffle colors for each square
//       const squareColors = shuffle([...colors]);
//       let index = 0;

//       setInterval(() => {
//         const colorVar = squareColors[index];
//         const color = getComputedStyle(document.documentElement).getPropertyValue(colorVar).trim();
//         (square as SVGElement).style.fill = color;

//         index = (index + 1) % squareColors.length;
//       }, 5000); // Change color every second
//     });
//   });


</script>

<main>
	<Nav />
	<!-- <div class="bg-svg"></div> -->
	<AnimatedBackground />
</main>

<style>
    /* .bg-svg {
        width: 100vw;
        height: 80vh;
        background-image: url('./assets/pattern-randomized.svg');
        background-size: cover;
        background-repeat: no-repeat;
    } */
</style>
