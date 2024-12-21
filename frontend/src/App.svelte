<script lang="ts">
	import Nav from "./lib/Nav.svelte";
	import AnimatedBackground from "./lib/AnimatedBackground.svelte";
	import { Link, Route, Router } from "svelte-routing";
	import Home from "./lib/Home.svelte";
	import AddGame from "./lib/form/AddGame.svelte";
</script>

<main>
  <Router>
    <Nav />
    <AnimatedBackground />
    <Route path="/" component={Home} />
    <Route path="/create/games" component={AddGame} />
  </Router>
</main>

<style>
</style>
