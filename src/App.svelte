<script>
  import { Router, Link, Route } from "svelte-routing";
	import Home from "./page/Home.svelte";
	import Nav from "./lib/Nav.svelte";
	import Liste from "./page/Liste.svelte";

  export let url = "";
</script>

<Router {url}>
  <Nav />
  <div>
    <Route path="/"><Home /></Route>
    <Route path="/contact"><Liste /></Route>
  </div>
</Router>