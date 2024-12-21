<script lang="ts">
  import Nav from "./lib/Nav.svelte";
  import AnimatedBackground from "./lib/AnimatedBackground.svelte";
  import { Route, Router } from "svelte-routing";
  import Home from "./lib/Home.svelte";
  import AddGame from "./lib/form/AddGame.svelte";
	import NotFound from "./lib/NotFound.svelte";
</script>

<main>
  <Router>
    <Nav />
    <AnimatedBackground />
    <Route path="/" component={Home} />
    <Route path="/create/games" component={AddGame} />
    <Route path="*" component={NotFound} />
  </Router>
</main>

<style>
</style>